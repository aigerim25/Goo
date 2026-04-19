package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetRateSuccess(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
	"base": "USD",
	"target": "EUR",
	"rate": 0.9
	}`))
	}))
	defer server.Close()
	service := NewExchangeService(server.URL)

	rate, err := service.GetRate("USD", "EUR")
	assert.NoError(t, err)
	assert.Equal(t, 0.9, rate)
}
func TestGetRateAPIFail(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{
	"error": "invalid currency pair"}`))
	}))
	defer server.Close()
	service := NewExchangeService(server.URL)
	rate, err := service.GetRate("CCC", "DDD")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid currency pair")
	assert.Equal(t, 0.0, rate)
}
func TestGetRateInvalidJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`invalid json`))
	}))
	defer server.Close()
	service := NewExchangeService(server.URL)
	_, err := service.GetRate("USD", "EUR")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "decode error")
}
func TestGetRateTimeout(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(6 * time.Second)
	}))
	defer server.Close()

	service := NewExchangeService(server.URL)

	_, err := service.GetRate("USD", "EUR")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "network error")
}
func TestGetRateEmptyBody(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	service := NewExchangeService(server.URL)

	_, err := service.GetRate("USD", "EUR")

	assert.Error(t, err)
}
func TestGetRateInternalServerError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`server error`))
	}))
	defer server.Close()

	service := NewExchangeService(server.URL)

	_, err := service.GetRate("USD", "EUR")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "decode error")
}

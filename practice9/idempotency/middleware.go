package idempotency

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"
)

type Record struct {
	Status       string
	StatusCode   int
	ResponseBody []byte
}

var storage = make(map[string]Record)
var mu sync.Mutex

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.Header.Get("Idempotency-Key")
		if key == "" {
			http.Error(w, "missing idempotency key", http.StatusBadRequest)
			return
		}
		mu.Lock()

		record, exists := storage[key]
		if exists && record.Status == "processing" {
			mu.Unlock()
			http.Error(w, "request is already processing", http.StatusConflict)
			return
		}
		if exists && record.Status == "completed" {
			mu.Unlock()
			w.WriteHeader(record.StatusCode)
			w.Write(record.ResponseBody)
			return
		}
		storage[key] = Record{
			Status: "processing",
		}
		mu.Unlock()

		recorder := &responseRecorder{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}
		next.ServeHTTP(recorder, r)
		mu.Lock()
		storage[key] = Record{
			Status:       "completed",
			StatusCode:   recorder.statusCode,
			ResponseBody: recorder.body,
		}
		mu.Unlock()
	})
}

type responseRecorder struct {
	http.ResponseWriter
	statusCode int
	body       []byte
}

func (r *responseRecorder) WriteHeader(statusCode int) {
	r.statusCode = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}
func (r *responseRecorder) Write(body []byte) (int, error) {
	r.body = append(r.body, body...)
	return r.ResponseWriter.Write(body)
}
func PaymentHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(2 * time.Second)

	response := map[string]interface{}{
		"status":         "paid",
		"amount":         1000,
		"transaction_id": "uuid-888888",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
}

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"practice9/filtering"
	"time"
)

func main() { // main function for simulation of a production situation
	attemptCounter := 0

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attemptCounter++
		if attemptCounter <= 3 {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"status": "success"})
	}))
	defer server.Close()

	client := &http.Client{}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req, _ := http.NewRequest("POST", server.URL, nil)
	maxRetries := 5

	_, err := filtering.ExecuteWithLogs(ctx, client, req, maxRetries)
	if err != nil {
		fmt.Println("Final error:", err)
		return
	}
}

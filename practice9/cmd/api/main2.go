package main

import (
	"fmt"
	"net/http"
	"practice9/idempotency"
	"time"
)

func main() { // main function for simulating a double-click attack
	http.Handle("/pay", idempotency.Middleware(http.HandlerFunc(idempotency.PaymentHandler)))

	go func() {
		fmt.Println("Starting server at port 8080")
		http.ListenAndServe(":8080", nil)
	}()
	time.Sleep(500 * time.Millisecond)
	key := "same-payment-key"
	url := "http://localhost:8080/pay"

	for i := 1; i <= 7; i++ {
		go func(requestNumber int) {
			req, _ := http.NewRequest("POST", url, nil)
			req.Header.Set("Idempotency-Key", key)

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				fmt.Println("Request", requestNumber, "error:", err)
				return
			}
			defer resp.Body.Close()
			fmt.Println("Request", requestNumber, "status:", resp.StatusCode)
		}(i)
	}
	time.Sleep(4 * time.Second)
	req, _ := http.NewRequest("POST", url, nil)
	req.Header.Set("Idempotency-Key", key)

	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	fmt.Println("Final request status:", resp.StatusCode)
}

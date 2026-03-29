package main

import (
	"assignment_1/internal/handlers"
	"assignment_1/internal/middleware"
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/tasks", handlers.TasksHandler)

	var h http.Handler = mux
	h = middleware.Logging(h)
	h = middleware.APIKeyAuth(h)
	http.HandleFunc("/tasks", handlers.TasksHandler)
	fmt.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", h)
	fmt.Println("ListenAndServe error:", err)
}

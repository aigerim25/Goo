package main

import (
	"database/sql"
	"log"
	"net/http"
	"practice5/internal/handler"
	"practice5/internal/repository"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "postgres://postgres:aiko2502@localhost:5432/godb?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	repo := repository.NewUserRepository(db)
	h := handler.NewUserHandler(repo)

	http.HandleFunc("/users", h.GetPaginatedUsersHandler)
	http.HandleFunc("/common-friends", h.GetCommonFriendsHandler)
	log.Println("Server started on port 8080")
	http.ListenAndServe(":8080", nil)
}

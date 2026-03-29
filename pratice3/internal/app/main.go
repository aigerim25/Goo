package app

import (
	"context"
	"fmt"
	"practice3/internal/repository"
	"practice3/internal/repository/_postgres"
	_ "practice3/internal/repository/_postgres/users"
	"practice3/pkg/modules"
	"time"
)

func Run() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	dbConfig := initPostgreConfig()
	_postgre := _postgres.NewPGXDialect(ctx, dbConfig)
	repository := repository.NewRepositories(_postgre)
	newID, err := repository.CreateUser(modules.User{
		Name:  "Aigerim",
		Age:   18,
		Email: "aigerim@test.com",
		Phone: "8700455000",
	})
	fmt.Println("created id:", newID, "err:", err)
}
func initPostgreConfig() *modules.PostgreConfig {
	return &modules.PostgreConfig{
		Host:        "localhost",
		Port:        "5432",
		Username:    "postgres",
		Password:    "aiko2502",
		DBName:      "godb",
		SSLMode:     "disable",
		ExecTimeout: 5 * time.Second,
	}
}

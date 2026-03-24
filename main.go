package main

import (
	"log"
	"net/http"

	"idp-jobs-go/api"
	"idp-jobs-go/repository"
	"idp-jobs-go/service"
)

func main() {
	repo := repository.NewInMemoryJobRepository()
	svc := service.NewJobService(repo)
	handler := api.NewHandler(svc)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler.Router()))
}
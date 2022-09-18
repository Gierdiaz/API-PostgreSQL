package main

import (
	"app/configs"
	"app/handlers"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Server
func main() {
	err := configs.Load()

	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()

	r.Post("/", handlers.Create)
	r.Put("/{id}", handlers.Update)
	r.Delete("/{id}", handlers.Delete)
	r.Get("/", handlers.List)
	r.Get("/{id}", handlers.Get)

	fmt.Println("Escutando na porta 9000")

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)

}
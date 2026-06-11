package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {

	r := chi.NewRouter()

	fmt.Println("Servidor corriendo en puerto 8080")

	http.ListenAndServe(":8080", r)
}
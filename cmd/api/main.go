package main

import (
	"fmt"
	"net/http"

	"proyecto-detallesPersonalizados-api/internal/handlers"

	"github.com/go-chi/chi/v5"

)

// RUTAS CHI
func main() {

	r := chi.NewRouter()

	//commit:CONFIGURAR RUTAS CHI Y SERVIDOR HTTP PARA PEDIDOS
	r.Route("/api/v1/pedidos", func(r chi.Router) {
		r.Post("/", handlers.CrearPedido)
		r.Get("/", handlers.ObtenerPedidos)
		r.Get("/{id}", handlers.ObtenerPedidoPorID)
		r.Put("/{id}", handlers.ActualizarPedido)
		r.Delete("/{id}", handlers.EliminarPedido)

	})

	
	fmt.Println("Servidor corriendo en puerto 8080")

	http.ListenAndServe(":8080", r)
}

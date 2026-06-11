package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"proyecto-detalles-api/internal/handlers"
)

//RUTAS CHI
func main() {

	r := chi.NewRouter()


	//commit:CONFIGURAR RUTAS CHI PARA PEDIDOS
	r.Route("/api/v1/pedidos", func(r chi.Router) {
	r.Post("/", handlers.CrearPedido)
	r.Get("/", handlers.ObtenerPedidos)
	r.Get("/{id}", handlers.ObtenerPedidoPorID)
	r.Put("/{id}", handlers.ActualizarPedido)
	r.Delete("/{id}", handlers.EliminarPedido)

	})

	//commit:CONFIGURAR RUTAS CHI PARA REPARTIDORES
	r.Route("/api/v1/repartidores", func(r chi.Router) {

	r.Post("/", handlers.CrearRepartidor)
	r.Get("/", handlers.ObtenerRepartidores)

	r.Get("/{id}", handlers.ObtenerRepartidorPorID)
	r.Put("/{id}", handlers.ActualizarRepartidor)
	r.Delete("/{id}", handlers.EliminarRepartidor)
})


	fmt.Println("Servidor corriendo en puerto 8080")

	http.ListenAndServe(":8080", r)
}
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

	r.Route("/api/v1/personalizaciones", func(r chi.Router) {
		r.Post("/", handlers.CrearPersonalizacion)
		r.Get("/", handlers.ObtenerPersonalizaciones)
		r.Get("/{id}", handlers.ObtenerPersonalizacionPorID)
		r.Put("/{id}", handlers.ActualizarPersonalizacion)
		r.Delete("/{id}", handlers.EliminarPersonalizacion)
})

r.Route("/api/v1/productos-personalizados", func(r chi.Router) {
	r.Post("/", handlers.CrearProductoPersonalizado)
	r.Get("/", handlers.ObtenerProductosPersonalizados)
	r.Get("/{id}", handlers.ObtenerProductoPersonalizadoPorID)
	r.Put("/{id}", handlers.ActualizarProductoPersonalizado)
	r.Delete("/{id}", handlers.EliminarProductoPersonalizado)

})
	fmt.Println("Servidor corriendo en puerto 8080")

	http.ListenAndServe(":8080", r)
}

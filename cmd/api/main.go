package main

import (
	"fmt"
	"net/http"

	"proyecto-detallesPersonalizados-api/internal/handlers"

	"github.com/go-chi/chi/v5"

	 "github.com/go-chi/chi/v5/middleware"

    middlewareAPI "proyecto-detallesPersonalizados-api/internal/middleware"

)

// RUTAS CHI
func main() {

	r := chi.NewRouter()


r.Use(middleware.Logger)
r.Use(middlewareAPI.CORS)



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

	r.Route("/api/v1/solicitudes-urgentes", func(r chi.Router) {
		r.Post("/", handlers.CrearSolicitudUrgente)
		r.Get("/", handlers.ObtenerSolicitudesUrgentes)
		r.Get("/{id}", handlers.ObtenerSolicitudUrgentePorID)
		r.Put("/{id}", handlers.ActualizarSolicitudUrgente)
		r.Delete("/{id}", handlers.EliminarSolicitudUrgente)
	})

	r.Route("/api/v1/agendas-produccion", func(r chi.Router) {
		r.Post("/", handlers.CrearAgendaProduccion)
		r.Get("/", handlers.ObtenerAgendasProduccion)
		r.Get("/{id}", handlers.ObtenerAgendaProduccionPorID)
		r.Put("/{id}", handlers.ActualizarAgendaProduccion)
		r.Delete("/{id}", handlers.EliminarAgendaProduccion)
	})

	r.Route("/api/v1/slots-produccion", func(r chi.Router) {
		r.Post("/", handlers.CrearSlotProduccion)
		r.Get("/", handlers.ObtenerSlotsProduccion)
		r.Get("/{id}", handlers.ObtenerSlotProduccionPorID)
		r.Put("/{id}", handlers.ActualizarSlotProduccion)
		r.Delete("/{id}", handlers.EliminarSlotProduccion)
	})

	r.Route("/api/v1/clientes", func(r chi.Router) {
	r.Post("/", handlers.CrearCliente)
	r.Get("/", handlers.ObtenerClientes)
	r.Get("/{id}", handlers.ObtenerClientePorID)
	r.Put("/{id}", handlers.ActualizarCliente)
	r.Delete("/{id}", handlers.EliminarCliente)

})

	r.Route("/api/v1/seguimientos", func(r chi.Router) {
		r.Post("/", handlers.CrearSeguimientoPedido)
		r.Get("/", handlers.ObtenerSeguimientosPedido)
		r.Get("/{id}", handlers.ObtenerSeguimientoPedidoPorID)
		r.Put("/{id}", handlers.ActualizarSeguimientoPedido)
		r.Delete("/{id}", handlers.EliminarSeguimientoPedido)	
	})

	r.Route("/api/v1/reclamos", func(r chi.Router) {
		r.Post("/", handlers.CrearReclamo)
		r.Get("/", handlers.ObtenerReclamos)
		r.Get("/{id}", handlers.ObtenerReclamoPorID)
		r.Put("/{id}", handlers.ActualizarReclamo)
		r.Delete("/{id}", handlers.EliminarReclamo)
	})

	fmt.Println("Servidor corriendo en puerto 8080")

	http.ListenAndServe(":8080", r)
}

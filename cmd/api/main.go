package main

import (
	"fmt"
	"net/http"

	"proyecto-detallesPersonalizados-api/internal/handlers"
	"proyecto-detallesPersonalizados-api/internal/service"
	"proyecto-detallesPersonalizados-api/internal/storage"

	"github.com/go-chi/chi/v5"

	 "github.com/go-chi/chi/v5/middleware"

    middlewareAPI "proyecto-detallesPersonalizados-api/internal/middleware"

)

// RUTAS CHI
func main() {

	r := chi.NewRouter()


r.Use(middleware.Logger)
r.Use(middlewareAPI.CORS)

// INYECCIÓN DE DEPENDENCIAS
	// ==========================================

	Almacen := storage.NuevoAlmacenSQLite(db)

	pedidoService := service.NewPedidoService(Almacen)
	personalizacionService := service.NewPersonalizacionService(Almacen)
	productoPersonalizadoService := service.NewProductoPersonalizadoService(Almacen)
	solicitudUrgenteService := service.NewSolicitudUrgenteService(Almacen)
	agendaProduccionService := service.NewAgendaProduccionService(Almacen)
	slotProduccionService := service.NewSlotProduccionService(Almacen)
	clienteService := service.NewClienteService(Almacen)
	seguimientoService := service.NewSeguimientoPedidoService(Almacen)
	reclamoService := service.NewReclamoService(Almacen)
	
	

	server := handlers.NewServer(
		pedidoService,
		personalizacionService,
		productoPersonalizadoService,
		solicitudUrgenteService,
		agendaProduccionService,
		slotProduccionService,
		clienteService,
		seguimientoService,
		reclamoService,
	)



	


	//commit:CONFIGURAR RUTAS CHI Y SERVIDOR HTTP PARA PEDIDOS
	r.Route("/api/v1/pedidos", func(r chi.Router) {
		r.Post("/", server.CrearPedido)
		r.Get("/", server.ObtenerPedidos)
		r.Get("/{id}", server.ObtenerPedidoPorID)
		r.Put("/{id}", server.ActualizarPedido)
		r.Delete("/{id}", server.EliminarPedido)
	})


	r.Route("/api/v1/personalizaciones", func(r chi.Router) {
		r.Post("/", server.CrearPersonalizacion)
		r.Get("/", server.ObtenerPersonalizaciones)
		r.Get("/{id}", server.ObtenerPersonalizacionPorID)
		r.Put("/{id}", server.ActualizarPersonalizacion)
		r.Delete("/{id}", server.EliminarPersonalizacion)
	})

	r.Route("/api/v1/productos-personalizados", func(r chi.Router) {
		r.Post("/", server.CrearProductoPersonalizado)
		r.Get("/", server.ObtenerProductosPersonalizados)
		r.Get("/{id}", server.ObtenerProductoPersonalizadoPorID)
		r.Put("/{id}", server.ActualizarProductoPersonalizado)
		r.Delete("/{id}", server.EliminarProductoPersonalizado)

	})

	r.Route("/api/v1/solicitudes-urgentes", func(r chi.Router) {
		r.Post("/", server.CrearSolicitudUrgente)
		r.Get("/", server.ObtenerSolicitudesUrgentes)
		r.Get("/{id}", server.ObtenerSolicitudUrgentePorID)
		r.Put("/{id}", server.ActualizarSolicitudUrgente)
		r.Delete("/{id}", server.EliminarSolicitudUrgente)
	})

	r.Route("/api/v1/agendas-produccion", func(r chi.Router) {
		r.Post("/", server.CrearAgendaProduccion)
		r.Get("/", server.ObtenerAgendasProduccion)
		r.Get("/{id}", server.ObtenerAgendaProduccionPorID)
		r.Put("/{id}", server.ActualizarAgendaProduccion)
		r.Delete("/{id}", server.EliminarAgendaProduccion)
	})

	r.Route("/api/v1/slots-produccion", func(r chi.Router) {
		r.Post("/", server.CrearSlotProduccion)
		r.Get("/", server.ObtenerSlotsProduccion)
		r.Get("/{id}", server.ObtenerSlotProduccionPorID)
		r.Put("/{id}", server.ActualizarSlotProduccion)
		r.Delete("/{id}", server.EliminarSlotProduccion)
	})

	r.Route("/api/v1/clientes", func(r chi.Router) {
	r.Post("/", server.CrearCliente)
	r.Get("/", server.ObtenerClientes)
	r.Get("/{id}", server.ObtenerClientePorID)
	r.Put("/{id}", server.ActualizarCliente)
	r.Delete("/{id}", server.EliminarCliente)

})

	r.Route("/api/v1/reclamos", func(r chi.Router) {
		r.Post("/", server.CrearReclamo)
		r.Get("/", server.ObtenerReclamos)
		r.Get("/{id}", server.ObtenerReclamoPorID)
		r.Put("/{id}", server.ActualizarReclamo)
		r.Delete("/{id}", server.EliminarReclamo)
	})

	r.Route("/api/v1/seguimientos", func(r chi.Router) {
		r.Post("/", server.CrearSeguimientoPedido)
		r.Get("/", server.ObtenerSeguimientosPedido)
		r.Get("/{id}", server.ObtenerSeguimientoPedidoPorID)
		r.Put("/{id}", server.ActualizarSeguimientoPedido)
		r.Delete("/{id}", server.EliminarSeguimientoPedido)	
	})


	fmt.Println("Servidor corriendo en puerto 8080")

	http.ListenAndServe(":8080", r)
}

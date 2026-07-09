package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"proyecto-detallesPersonalizados-api/internal/config"
	"proyecto-detallesPersonalizados-api/internal/handlers"
	middlewareAPI "proyecto-detallesPersonalizados-api/internal/middleware"
	"proyecto-detallesPersonalizados-api/internal/service"
	"proyecto-detallesPersonalizados-api/internal/storage"
)

// ===================== MAIN =====================
func main() {
	cfg := config.Cargar()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middlewareAPI.CORS)

	db, almacen, err := storage.InicializarBaseDatos(cfg.DBDriver, cfg.RutaDB, cfg.DBDSN)
	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	pedidoService := service.NewPedidoService(almacen)
	personalizacionService := service.NewPersonalizacionService(almacen)
	productoPersonalizadoService := service.NewProductoPersonalizadoService(almacen)
	solicitudUrgenteService := service.NewSolicitudUrgenteService(almacen)
	agendaProduccionService := service.NewAgendaProduccionService(almacen)
	slotProduccionService := service.NewSlotProduccionService(almacen)
	clienteService := service.NewClienteService(almacen)
	seguimientoService := service.NewSeguimientoPedidoService(almacen)
	reclamoService := service.NewReclamoService(almacen)
	usuarioRepo := storage.NewUsuarioGORM(db)
	authService := service.NewAuthService(
		usuarioRepo,
		service.WithJWTSecreto(cfg.JWTSecreto),
		service.WithJWTDuracion(cfg.JWTDuracion),
	)

	server := handlers.NewServer(handlers.Deps{
		Pedidos:                 pedidoService,
		Personalizaciones:       personalizacionService,
		ProductosPersonalizados: productoPersonalizadoService,
		SolicitudesUrgentes:     solicitudUrgenteService,
		AgendasProduccion:       agendaProduccionService,
		SlotsProduccion:         slotProduccionService,
		Clientes:                clienteService,
		Seguimientos:            seguimientoService,
		Reclamos:                reclamoService,
		Auth:                    authService,
	})

	authMiddleware := middlewareAPI.Auth(authService)
	admin := middlewareAPI.RequiereRol(service.RolAdmin)
	clienteOAdmin := middlewareAPI.RequiereRol(service.RolCliente, service.RolAdmin)

	r.Route("/api/v1/auth", func(r chi.Router) {
		r.Post("/register", server.Registrar)
		r.Post("/login", server.Login)
	})

	r.Group(func(r chi.Router) {
		r.Use(authMiddleware)

		r.Route("/api/v1/pedidos", func(r chi.Router) {
			r.With(clienteOAdmin).Post("/", server.CrearPedido)
			r.With(admin).Get("/", server.ObtenerPedidos)
			r.With(clienteOAdmin).Get("/{id}", server.ObtenerPedidoPorID)
			r.With(admin).Put("/{id}", server.ActualizarPedido)
			r.With(admin).Delete("/{id}", server.EliminarPedido)
		})

		r.Route("/api/v1/personalizaciones", func(r chi.Router) {
			r.With(clienteOAdmin).Post("/", server.CrearPersonalizacion)
			r.With(admin).Get("/", server.ObtenerPersonalizaciones)
			r.With(clienteOAdmin).Get("/{id}", server.ObtenerPersonalizacionPorID)
			r.With(admin).Put("/{id}", server.ActualizarPersonalizacion)
			r.With(admin).Delete("/{id}", server.EliminarPersonalizacion)
		})

		r.Route("/api/v1/productos-personalizados", func(r chi.Router) {
			r.With(admin).Post("/", server.CrearProductoPersonalizado)
			r.With(clienteOAdmin).Get("/", server.ObtenerProductosPersonalizados)
			r.With(clienteOAdmin).Get("/{id}", server.ObtenerProductoPersonalizadoPorID)
			r.With(admin).Put("/{id}", server.ActualizarProductoPersonalizado)
			r.With(admin).Delete("/{id}", server.EliminarProductoPersonalizado)
		})

		r.Route("/api/v1/solicitudes-urgentes", func(r chi.Router) {
			r.With(clienteOAdmin).Post("/", server.CrearSolicitudUrgente)
			r.With(admin).Get("/", server.ObtenerSolicitudesUrgentes)
			r.With(clienteOAdmin).Get("/{id}", server.ObtenerSolicitudUrgentePorID)
			r.With(admin).Put("/{id}", server.ActualizarSolicitudUrgente)
			r.With(admin).Delete("/{id}", server.EliminarSolicitudUrgente)
		})

		r.Route("/api/v1/agendas-produccion", func(r chi.Router) {
			r.With(admin).Post("/", server.CrearAgendaProduccion)
			r.With(admin).Get("/", server.ObtenerAgendasProduccion)
			r.With(admin).Get("/{id}", server.ObtenerAgendaProduccionPorID)
			r.With(admin).Put("/{id}", server.ActualizarAgendaProduccion)
			r.With(admin).Delete("/{id}", server.EliminarAgendaProduccion)
		})

		r.Route("/api/v1/slots-produccion", func(r chi.Router) {
			r.With(admin).Post("/", server.CrearSlotProduccion)
			r.With(admin).Get("/", server.ObtenerSlotsProduccion)
			r.With(admin).Get("/{id}", server.ObtenerSlotProduccionPorID)
			r.With(admin).Put("/{id}", server.ActualizarSlotProduccion)
			r.With(admin).Delete("/{id}", server.EliminarSlotProduccion)
		})

		r.Route("/api/v1/clientes", func(r chi.Router) {
			r.With(clienteOAdmin).Post("/", server.CrearCliente)
			r.With(admin).Get("/", server.ObtenerClientes)
			r.With(clienteOAdmin).Get("/{id}", server.ObtenerClientePorID)
			r.With(admin).Put("/{id}", server.ActualizarCliente)
			r.With(admin).Delete("/{id}", server.EliminarCliente)
		})

		r.Route("/api/v1/reclamos", func(r chi.Router) {
			r.With(clienteOAdmin).Post("/", server.CrearReclamo)
			r.With(admin).Get("/", server.ObtenerReclamos)
			r.With(clienteOAdmin).Get("/{id}", server.ObtenerReclamoPorID)
			r.With(admin).Put("/{id}", server.ActualizarReclamo)
			r.With(admin).Delete("/{id}", server.EliminarReclamo)
		})

		r.Route("/api/v1/seguimientos", func(r chi.Router) {
			r.With(admin).Post("/", server.CrearSeguimientoPedido)
			r.With(clienteOAdmin).Get("/", server.ObtenerSeguimientosPedido)
			r.With(clienteOAdmin).Get("/{id}", server.ObtenerSeguimientoPedidoPorID)
			r.With(admin).Put("/{id}", server.ActualizarSeguimientoPedido)
			r.With(admin).Delete("/{id}", server.EliminarSeguimientoPedido)
		})
	})

	httpServer := &http.Server{
		Addr:         direccionHTTP(cfg.Puerto),
		Handler:      r,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		log.Printf("Servidor corriendo en %s", httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("error en servidor HTTP: %v", err)
		}
	}()

	<-ctx.Done()
	stop()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
	defer cancel()

	log.Println("Apagando servidor ordenadamente...")
	if err := httpServer.Shutdown(shutdownCtx); err != nil {
		log.Printf("apagado ordenado fallido: %v", err)
		if err := httpServer.Close(); err != nil {
			log.Printf("error cerrando servidor HTTP: %v", err)
		}
	}

	log.Println("Servidor apagado")
}

func direccionHTTP(puerto string) string {
	puerto = strings.TrimSpace(puerto)
	if strings.HasPrefix(puerto, ":") {
		return puerto
	}
	return ":" + puerto
}

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

	r.Route("/api/v1/auth", func(r chi.Router) {
		r.Post("/register", server.Registrar)
		r.Post("/login", server.Login)
	})

	r.Group(func(r chi.Router) {
		r.Use(authMiddleware)

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

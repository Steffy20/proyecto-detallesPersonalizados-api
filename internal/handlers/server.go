package handlers

import (
	"proyecto-detallesPersonalizados-api/internal/models"
	"proyecto-detallesPersonalizados-api/internal/service"
)

type PedidoHandlerService interface {
	Listar() []models.Pedido
	Obtener(id int) (models.Pedido, error)
	Crear(models.Pedido) (models.Pedido, error)
	Actualizar(id int, datos models.Pedido) (models.Pedido, error)
	Borrar(id int) error
}

type PersonalizacionHandlerService interface {
	Listar() []models.Personalizacion
	Obtener(id int) (models.Personalizacion, error)
	Crear(models.Personalizacion) (models.Personalizacion, error)
	Actualizar(id int, datos models.Personalizacion) (models.Personalizacion, error)
	Borrar(id int) error
}

type ProductoPersonalizadoHandlerService interface {
	Listar() []models.ProductoPersonalizado
	Obtener(id int) (models.ProductoPersonalizado, error)
	Crear(models.ProductoPersonalizado) (models.ProductoPersonalizado, error)
	Actualizar(id int, datos models.ProductoPersonalizado) (models.ProductoPersonalizado, error)
	Borrar(id int) error
}

type SolicitudUrgenteHandlerService interface {
	Listar() []models.SolicitudUrgente
	Obtener(id int) (models.SolicitudUrgente, error)
	Crear(models.SolicitudUrgente) (models.SolicitudUrgente, error)
	Actualizar(id int, datos models.SolicitudUrgente) (models.SolicitudUrgente, error)
	Borrar(id int) error
}

type AgendaProduccionHandlerService interface {
	Listar() []models.AgendaProduccion
	Obtener(id int) (models.AgendaProduccion, error)
	Crear(models.AgendaProduccion) (models.AgendaProduccion, error)
	Actualizar(id int, datos models.AgendaProduccion) (models.AgendaProduccion, error)
	Borrar(id int) error
}

type SlotProduccionHandlerService interface {
	Listar() []models.SlotProduccion
	Obtener(id int) (models.SlotProduccion, error)
	Crear(models.SlotProduccion) (models.SlotProduccion, error)
	Actualizar(id int, datos models.SlotProduccion) (models.SlotProduccion, error)
	Borrar(id int) error
}

type ClienteHandlerService interface {
	Listar() []models.Cliente
	Obtener(id int) (models.Cliente, error)
	Crear(models.Cliente) (models.Cliente, error)
	Actualizar(id int, datos models.Cliente) (models.Cliente, error)
	Borrar(id int) error
}

type SeguimientoPedidoHandlerService interface {
	Listar() []models.SeguimientoPedido
	Obtener(id int) (models.SeguimientoPedido, error)
	Crear(models.SeguimientoPedido) (models.SeguimientoPedido, error)
	Actualizar(id int, datos models.SeguimientoPedido) (models.SeguimientoPedido, error)
	Borrar(id int) error
}

type ReclamoHandlerService interface {
	Listar() []models.Reclamo
	Obtener(id int) (models.Reclamo, error)
	Crear(models.Reclamo) (models.Reclamo, error)
	Actualizar(id int, datos models.Reclamo) (models.Reclamo, error)
	Borrar(id int) error
}

type Server struct {
	Pedidos                 PedidoHandlerService
	Personalizaciones       PersonalizacionHandlerService
	ProductosPersonalizados ProductoPersonalizadoHandlerService

	SolicitudesUrgentes SolicitudUrgenteHandlerService
	AgendasProduccion   AgendaProduccionHandlerService
	SlotsProduccion     SlotProduccionHandlerService

	Clientes     ClienteHandlerService
	Seguimientos SeguimientoPedidoHandlerService
	Reclamos     ReclamoHandlerService
	Auth         *service.AuthService
}

func NewServer(
	pedidos PedidoHandlerService,
	personalizaciones PersonalizacionHandlerService,
	productosPersonalizados ProductoPersonalizadoHandlerService,

	solicitudesUrgentes SolicitudUrgenteHandlerService,
	agendasProduccion AgendaProduccionHandlerService,
	slotsProduccion SlotProduccionHandlerService,

	clientes ClienteHandlerService,
	seguimientos SeguimientoPedidoHandlerService,
	reclamos ReclamoHandlerService,
	auth *service.AuthService,
) *Server {
	return &Server{
		Pedidos:                 pedidos,
		Personalizaciones:       personalizaciones,
		ProductosPersonalizados: productosPersonalizados,

		SolicitudesUrgentes: solicitudesUrgentes,
		AgendasProduccion:   agendasProduccion,
		SlotsProduccion:     slotsProduccion,

		Clientes:     clientes,
		Seguimientos: seguimientos,
		Reclamos:     reclamos,
		Auth:         auth,
	}
}
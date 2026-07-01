package handlers

import "proyecto-detallesPersonalizados-api/internal/service"

type Server struct {
	Pedidos                 *service.PedidoService
	Personalizaciones       *service.PersonalizacionService
	ProductosPersonalizados *service.ProductoPersonalizadoService

	SolicitudesUrgentes *service.SolicitudUrgenteService
	AgendasProduccion   *service.AgendaProduccionService
	SlotsProduccion     *service.SlotProduccionService

	Clientes            *service.ClienteService
	Seguimientos  *service.SeguimientoPedidoService
	Reclamos            *service.ReclamoService
}

func NewServer(
	pedidos *service.PedidoService,
	personalizaciones *service.PersonalizacionService,
	productosPersonalizados *service.ProductoPersonalizadoService,

	solicitudesUrgentes *service.SolicitudUrgenteService,
	agendasProduccion *service.AgendaProduccionService,
	slotsProduccion *service.SlotProduccionService,

	clientes *service.ClienteService,
	seguimientos *service.SeguimientoPedidoService,
	reclamos *service.ReclamoService,
) *Server {

	return &Server{
		Pedidos:                 pedidos,
		Personalizaciones:       personalizaciones,
		ProductosPersonalizados: productosPersonalizados,

		SolicitudesUrgentes: solicitudesUrgentes,
		AgendasProduccion:   agendasProduccion,
		SlotsProduccion:     slotsProduccion,

		Clientes:           clientes,
		Seguimientos: seguimientos,
		Reclamos:           reclamos,
	}
}

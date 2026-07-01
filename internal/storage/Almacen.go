package storage

import "proyecto-detallesPersonalizados-api/internal/models"

type Almacen interface {

	// ===================== PEDIDOS =====================

	ListarPedidos() []models.Pedido
	BuscarPedidoPorID(id int) (models.Pedido, bool)
	CrearPedido(p models.Pedido) models.Pedido
	ActualizarPedido(id int, datos models.Pedido) (models.Pedido, bool)
	BorrarPedido(id int) bool

	// ===================== PERSONALIZACIONES =====================

	ListarPersonalizaciones() []models.Personalizacion
	BuscarPersonalizacionPorID(id int) (models.Personalizacion, bool)
	CrearPersonalizacion(p models.Personalizacion) models.Personalizacion
	ActualizarPersonalizacion(id int, datos models.Personalizacion) (models.Personalizacion, bool)
	BorrarPersonalizacion(id int) bool

	// ===================== PRODUCTOS PERSONALIZADOS =====================

	ListarProductosPersonalizados() []models.ProductoPersonalizado
	BuscarProductoPersonalizadoPorID(id int) (models.ProductoPersonalizado, bool)
	CrearProductoPersonalizado(p models.ProductoPersonalizado) models.ProductoPersonalizado
	ActualizarProductoPersonalizado(id int, datos models.ProductoPersonalizado) (models.ProductoPersonalizado, bool)
	BorrarProductoPersonalizado(id int) bool

	// ===================== SOLICITUDES URGENTES =====================

	ListarSolicitudesUrgentes() []models.SolicitudUrgente
	BuscarSolicitudUrgentePorID(id int) (models.SolicitudUrgente, bool)
	CrearSolicitudUrgente(s models.SolicitudUrgente) models.SolicitudUrgente
	ActualizarSolicitudUrgente(id int, datos models.SolicitudUrgente) (models.SolicitudUrgente, bool)
	BorrarSolicitudUrgente(id int) bool

	// ===================== AGENDAS DE PRODUCCIÓN =====================

	ListarAgendasProduccion() []models.AgendaProduccion
	BuscarAgendaProduccionPorID(id int) (models.AgendaProduccion, bool)
	CrearAgendaProduccion(a models.AgendaProduccion) models.AgendaProduccion
	ActualizarAgendaProduccion(id int, datos models.AgendaProduccion) (models.AgendaProduccion, bool)
	BorrarAgendaProduccion(id int) bool

	// ===================== SLOTS DE PRODUCCIÓN =====================

	ListarSlotsProduccion() []models.SlotProduccion
	BuscarSlotProduccionPorID(id int) (models.SlotProduccion, bool)
	CrearSlotProduccion(s models.SlotProduccion) models.SlotProduccion
	ActualizarSlotProduccion(id int, datos models.SlotProduccion) (models.SlotProduccion, bool)
	BorrarSlotProduccion(id int) bool

	// ===================== CLIENTES =====================

	ListarClientes() []models.Cliente
	BuscarClientePorID(id int) (models.Cliente, bool)
	CrearCliente(c models.Cliente) models.Cliente
	ActualizarCliente(id int, datos models.Cliente) (models.Cliente, bool)
	BorrarCliente(id int) bool

	// ===================== SEGUIMIENTOS =====================

	ListarSeguimientosPedido() []models.SeguimientoPedido
	BuscarSeguimientoPedidoPorID(id int) (models.SeguimientoPedido, bool)
	CrearSeguimientoPedido(s models.SeguimientoPedido) models.SeguimientoPedido
	ActualizarSeguimientoPedido(id int, datos models.SeguimientoPedido) (models.SeguimientoPedido, bool)
	BorrarSeguimientoPedido(id int) bool

	// ===================== RECLAMOS =====================

	ListarReclamos() []models.Reclamo
	BuscarReclamoPorID(id int) (models.Reclamo, bool)
	CrearReclamo(r models.Reclamo) models.Reclamo
	ActualizarReclamo(id int, datos models.Reclamo) (models.Reclamo, bool)
	BorrarReclamo(id int) bool
}


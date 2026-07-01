package storage

import (
	"context"
	"database/sql"

	"proyecto-detallesPersonalizados-api/internal/models"
	"proyecto-detallesPersonalizados-api/internal/storage/sqlcdb"
)

// AlmacenSQLC implementa la interfaz Almacen usando SQLC.
type AlmacenSQLC struct {
	q *sqlcdb.Queries
}

// Constructor
func NuevoAlmacenSQLC(db *sql.DB) *AlmacenSQLC {
	return &AlmacenSQLC{
		q: sqlcdb.New(db),
	}
}

// =========================================================
// MAPEO SQLC -> DOMINIO
// =========================================================

func aPedidoDominio(p sqlcdb.Pedido) models.Pedido {
	return models.Pedido{
		ID:       int(p.ID),
		Mensaje:  p.Mensaje,
		Estado:   p.Estado,
	}
}

// =========================================================
// PEDIDOS
// =========================================================

func (a *AlmacenSQLC) ListarPedidos() []models.Pedido {

	filas, err := a.q.ListarPedidos(context.Background())

	if err != nil {
		return nil
	}

	out := make([]models.Pedido, 0, len(filas))

	for _, fila := range filas {
		out = append(out, aPedidoDominio(fila))
	}

	return out
}

func (a *AlmacenSQLC) BuscarPedidoPorID(id int) (models.Pedido, bool) {

	fila, err := a.q.BuscarPedidoPorID(
		context.Background(),
		int64(id),
	)

	if err != nil {
		return models.Pedido{}, false
	}

	return aPedidoDominio(fila), true
}

func (a *AlmacenSQLC) CrearPedido(p models.Pedido) models.Pedido {

	fila, err := a.q.CrearPedido(
		context.Background(),
		sqlcdb.CrearPedidoParams{
			Mensaje: p.Mensaje,
			Estado:  p.Estado,
		},
	)

	if err != nil {
		return models.Pedido{}
	}

	return aPedidoDominio(fila)
}

func (a *AlmacenSQLC) ActualizarPedido(id int, datos models.Pedido) (models.Pedido, bool) {

	fila, err := a.q.ActualizarPedido(
		context.Background(),
		sqlcdb.ActualizarPedidoParams{
			Mensaje: datos.Mensaje,
			Estado:  datos.Estado,
			ID:      int64(id),
		},
	)

	if err != nil {
		return models.Pedido{}, false
	}

	return aPedidoDominio(fila), true
}

func (a *AlmacenSQLC) BorrarPedido(id int) bool {

	filas, err := a.q.BorrarPedido(
		context.Background(),
		int64(id),
	)

	if err != nil {
		return false
	}

	return filas > 0
}


// =========================================================
// MAPEO SQLC -> DOMINIO
// =========================================================

func aPersonalizacionDominio(p sqlcdb.Personalizacion) models.Personalizacion {
	return models.Personalizacion{
		ID:       int(p.ID),
		PedidoID: int(p.PedidoID),
		Mensaje:  p.Mensaje,
		Color:    p.Color,
	}
}

// =========================================================
// PERSONALIZACIONES
// =========================================================

func (a *AlmacenSQLC) ListarPersonalizaciones() []models.Personalizacion {

	filas, err := a.q.ListarPersonalizaciones(context.Background())

	if err != nil {
		return nil
	}

	out := make([]models.Personalizacion, 0, len(filas))

	for _, fila := range filas {
		out = append(out, aPersonalizacionDominio(fila))
	}

	return out
}

func (a *AlmacenSQLC) BuscarPersonalizacionPorID(id int) (models.Personalizacion, bool) {

	fila, err := a.q.BuscarPersonalizacionPorID(
		context.Background(),
		int64(id),
	)

	if err != nil {
		return models.Personalizacion{}, false
	}

	return aPersonalizacionDominio(fila), true
}

func (a *AlmacenSQLC) CrearPersonalizacion(p models.Personalizacion) models.Personalizacion {

	fila, err := a.q.CrearPersonalizacion(
		context.Background(),
		sqlcdb.CrearPersonalizacionParams{
			PedidoID: int64(p.PedidoID),
			Mensaje:  p.Mensaje,
			Color:    p.Color,
		},
	)

	if err != nil {
		return models.Personalizacion{}
	}

	return aPersonalizacionDominio(fila)
}

func (a *AlmacenSQLC) ActualizarPersonalizacion(id int, datos models.Personalizacion) (models.Personalizacion, bool) {

	fila, err := a.q.ActualizarPersonalizacion(
		context.Background(),
		sqlcdb.ActualizarPersonalizacionParams{
			PedidoID: int64(datos.PedidoID),
			Mensaje:  datos.Mensaje,
			Color:    datos.Color,
			ID:       int64(id),
		},
	)

	if err != nil {
		return models.Personalizacion{}, false
	}

	return aPersonalizacionDominio(fila), true
}

func (a *AlmacenSQLC) BorrarPersonalizacion(id int) bool {

	filas, err := a.q.BorrarPersonalizacion(
		context.Background(),
		int64(id),
	)

	if err != nil {
		return false
	}

	return filas > 0
}

// =========================================================
// MAPEO SQLC -> DOMINIO
// =========================================================

func aProductoPersonalizadoDominio(p sqlcdb.ProductoPersonalizado) models.ProductoPersonalizado {
	return models.ProductoPersonalizado{
		ID:       int(p.ID),
		PedidoID: int(p.PedidoID),
		Nombre:   p.Nombre,
		Cantidad: int(p.Cantidad),
		Precio:   p.Precio,
	}
}

// =========================================================
// PRODUCTOS PERSONALIZADOS
// =========================================================

func (a *AlmacenSQLC) ListarProductosPersonalizados() []models.ProductoPersonalizado {

	filas, err := a.q.ListarProductosPersonalizados(context.Background())

	if err != nil {
		return nil
	}

	out := make([]models.ProductoPersonalizado, 0, len(filas))

	for _, fila := range filas {
		out = append(out, aProductoPersonalizadoDominio(fila))
	}

	return out
}

func (a *AlmacenSQLC) BuscarProductoPersonalizadoPorID(id int) (models.ProductoPersonalizado, bool) {

	fila, err := a.q.BuscarProductoPersonalizadoPorID(
		context.Background(),
		int64(id),
	)

	if err != nil {
		return models.ProductoPersonalizado{}, false
	}

	return aProductoPersonalizadoDominio(fila), true
}

func (a *AlmacenSQLC) CrearProductoPersonalizado(p models.ProductoPersonalizado) models.ProductoPersonalizado {

	fila, err := a.q.CrearProductoPersonalizado(
		context.Background(),
		sqlcdb.CrearProductoPersonalizadoParams{
			PedidoID: int64(p.PedidoID),
			Nombre:   p.Nombre,
			Cantidad: int64(p.Cantidad),
			Precio:   p.Precio,
		},
	)

	if err != nil {
		return models.ProductoPersonalizado{}
	}

	return aProductoPersonalizadoDominio(fila)
}

func (a *AlmacenSQLC) ActualizarProductoPersonalizado(id int, datos models.ProductoPersonalizado) (models.ProductoPersonalizado, bool) {

	fila, err := a.q.ActualizarProductoPersonalizado(
		context.Background(),
		sqlcdb.ActualizarProductoPersonalizadoParams{
			PedidoID: int64(datos.PedidoID),
			Nombre:   datos.Nombre,
			Cantidad: int64(datos.Cantidad),
			Precio:   datos.Precio,
			ID:       int64(id),
		},
	)

	if err != nil {
		return models.ProductoPersonalizado{}, false
	}

	return aProductoPersonalizadoDominio(fila), true
}

func (a *AlmacenSQLC) BorrarProductoPersonalizado(id int) bool {

	filas, err := a.q.BorrarProductoPersonalizado(
		context.Background(),
		int64(id),
	)

	if err != nil {
		return false
	}

	return filas > 0
}


// =========================================================
// MAPEO SQLC -> DOMINIO
// =========================================================

func aSolicitudUrgenteDominio(s sqlcdb.Solicitudurgente) models.SolicitudUrgente {
	return models.SolicitudUrgente{
		ID:             int(s.ID),
		Cliente:        s.Cliente,
		Descripcion:    s.Descripcion,
		FechaRequerida: s.FechaRequerida,
		Estado:         s.Estado,
	}
}

// =========================================================
// SOLICITUDES URGENTES
// =========================================================

func (a *AlmacenSQLC) ListarSolicitudesUrgentes() []models.SolicitudUrgente {

	filas, err := a.q.ListarSolicitudesUrgentes(context.Background())

	if err != nil {
		return nil
	}

	out := make([]models.SolicitudUrgente, 0, len(filas))

	for _, fila := range filas {
		out = append(out, aSolicitudUrgenteDominio(fila))
	}

	return out
}

func (a *AlmacenSQLC) BuscarSolicitudUrgentePorID(id int) (models.SolicitudUrgente, bool) {

	fila, err := a.q.BuscarSolicitudUrgentePorID(
		context.Background(),
		int64(id),
	)

	if err != nil {
		return models.SolicitudUrgente{}, false
	}

	return aSolicitudUrgenteDominio(fila), true
}

func (a *AlmacenSQLC) CrearSolicitudUrgente(s models.SolicitudUrgente) models.SolicitudUrgente {

	fila, err := a.q.CrearSolicitudUrgente(
		context.Background(),
		sqlcdb.CrearSolicitudUrgenteParams{
			Cliente:        s.Cliente,
			Descripcion:    s.Descripcion,
			FechaRequerida: s.FechaRequerida,
			Estado:         s.Estado,
		},
	)

	if err != nil {
		return models.SolicitudUrgente{}
	}

	return aSolicitudUrgenteDominio(fila)
}

func (a *AlmacenSQLC) ActualizarSolicitudUrgente(id int, datos models.SolicitudUrgente) (models.SolicitudUrgente, bool) {

	fila, err := a.q.ActualizarSolicitudUrgente(
		context.Background(),
		sqlcdb.ActualizarSolicitudUrgenteParams{
			Cliente:        datos.Cliente,
			Descripcion:    datos.Descripcion,
			FechaRequerida: datos.FechaRequerida,
			Estado:         datos.Estado,
			ID:             int64(id),
		},
	)

	if err != nil {
		return models.SolicitudUrgente{}, false
	}

	return aSolicitudUrgenteDominio(fila), true
}

func (a *AlmacenSQLC) BorrarSolicitudUrgente(id int) bool {

	filas, err := a.q.BorrarSolicitudUrgente(
		context.Background(),
		int64(id),
	)

	if err != nil {
		return false
	}

	return filas > 0
}

// =========================================================
// MAPEO SQLC -> DOMINIO
// =========================================================

func aAgendaProduccionDominio(aq sqlcdb.Agendaproduccion) models.AgendaProduccion {
	return models.AgendaProduccion{
		ID:          int(aq.ID),
		Fecha:       aq.Fecha,
		Responsable: aq.Responsable,
		Estado:      aq.Estado,
	}
}

// =========================================================
// AGENDAS DE PRODUCCIÓN
// =========================================================

func (a *AlmacenSQLC) ListarAgendasProduccion() []models.AgendaProduccion {

	filas, err := a.q.ListarAgendasProduccion(context.Background())

	if err != nil {
		return nil
	}

	out := make([]models.AgendaProduccion, 0, len(filas))

	for _, fila := range filas {
		out = append(out, aAgendaProduccionDominio(fila))
	}

	return out
}

func (a *AlmacenSQLC) BuscarAgendaProduccionPorID(id int) (models.AgendaProduccion, bool) {

	fila, err := a.q.BuscarAgendaProduccionPorID(
		context.Background(),
		int64(id),
	)

	if err != nil {
		return models.AgendaProduccion{}, false
	}

	return aAgendaProduccionDominio(fila), true
}

func (a *AlmacenSQLC) CrearAgendaProduccion(agenda models.AgendaProduccion) models.AgendaProduccion {

	fila, err := a.q.CrearAgendaProduccion(
		context.Background(),
		sqlcdb.CrearAgendaProduccionParams{
			Fecha:       agenda.Fecha,
			Responsable: agenda.Responsable,
			Estado:      agenda.Estado,
		},
	)

	if err != nil {
		return models.AgendaProduccion{}
	}

	return aAgendaProduccionDominio(fila)
}

func (a *AlmacenSQLC) ActualizarAgendaProduccion(id int, datos models.AgendaProduccion) (models.AgendaProduccion, bool) {

	fila, err := a.q.ActualizarAgendaProduccion(
		context.Background(),
		sqlcdb.ActualizarAgendaProduccionParams{
			Fecha:       datos.Fecha,
			Responsable: datos.Responsable,
			Estado:      datos.Estado,
			ID:          int64(id),
		},
	)

	if err != nil {
		return models.AgendaProduccion{}, false
	}

	return aAgendaProduccionDominio(fila), true
}

func (a *AlmacenSQLC) BorrarAgendaProduccion(id int) bool {

	filas, err := a.q.BorrarAgendaProduccion(
		context.Background(),
		int64(id),
	)

	if err != nil {
		return false
	}

	return filas > 0
}


// =========================================================
// MAPEO SQLC -> DOMINIO
// =========================================================

func aSlotProduccionDominio(s sqlcdb.Slotproduccion) models.SlotProduccion {
	return models.SlotProduccion{
		ID:               int(s.ID),
		AgendaID:         int(s.AgendaID),
		CapacidadMaxima:  int(s.CapacidadMaxima),
		PedidosAsignados: int(s.PedidosAsignados),
	}
}

// =========================================================
// SLOTS DE PRODUCCIÓN
// =========================================================

func (a *AlmacenSQLC) ListarSlotsProduccion() []models.SlotProduccion {

	filas, err := a.q.ListarSlotsProduccion(context.Background())

	if err != nil {
		return nil
	}

	out := make([]models.SlotProduccion, 0, len(filas))

	for _, fila := range filas {
		out = append(out, aSlotProduccionDominio(fila))
	}

	return out
}

func (a *AlmacenSQLC) BuscarSlotProduccionPorID(id int) (models.SlotProduccion, bool) {

	fila, err := a.q.BuscarSlotProduccionPorID(
		context.Background(),
		int64(id),
	)

	if err != nil {
		return models.SlotProduccion{}, false
	}

	return aSlotProduccionDominio(fila), true
}

func (a *AlmacenSQLC) CrearSlotProduccion(slot models.SlotProduccion) models.SlotProduccion {

	fila, err := a.q.CrearSlotProduccion(
		context.Background(),
		sqlcdb.CrearSlotProduccionParams{
			AgendaID:         int64(slot.AgendaID),
			CapacidadMaxima:  int64(slot.CapacidadMaxima),
			PedidosAsignados: int64(slot.PedidosAsignados),
		},
	)

	if err != nil {
		return models.SlotProduccion{}
	}

	return aSlotProduccionDominio(fila)
}

func (a *AlmacenSQLC) ActualizarSlotProduccion(id int, datos models.SlotProduccion) (models.SlotProduccion, bool) {

	fila, err := a.q.ActualizarSlotProduccion(
		context.Background(),
		sqlcdb.ActualizarSlotProduccionParams{
			AgendaID:         int64(datos.AgendaID),
			CapacidadMaxima:  int64(datos.CapacidadMaxima),
			PedidosAsignados: int64(datos.PedidosAsignados),
			ID:               int64(id),
		},
	)

	if err != nil {
		return models.SlotProduccion{}, false
	}

	return aSlotProduccionDominio(fila), true
}

func (a *AlmacenSQLC) BorrarSlotProduccion(id int) bool {

	filas, err := a.q.BorrarSlotProduccion(
		context.Background(),
		int64(id),
	)

	if err != nil {
		return false
	}

	return filas > 0
}

// =========================================================
// MAPEO SQLC -> DOMINIO
// =========================================================

func aClienteDominio(c sqlcdb.Cliente) models.Cliente {
	return models.Cliente{
		ID:       int(c.ID),
		Nombre:   c.Nombre,
		Telefono: c.Telefono,
	}
}

// =========================================================
// CLIENTES
// =========================================================

func (a *AlmacenSQLC) ListarClientes() []models.Cliente {

	filas, err := a.q.ListarClientes(context.Background())

	if err != nil {
		return nil
	}

	out := make([]models.Cliente, 0, len(filas))

	for _, fila := range filas {
		out = append(out, aClienteDominio(fila))
	}

	return out
}

func (a *AlmacenSQLC) BuscarClientePorID(id int) (models.Cliente, bool) {

	fila, err := a.q.BuscarClientePorID(
		context.Background(),
		int64(id),
	)

	if err != nil {
		return models.Cliente{}, false
	}

	return aClienteDominio(fila), true
}

func (a *AlmacenSQLC) CrearCliente(c models.Cliente) models.Cliente {

	fila, err := a.q.CrearCliente(
		context.Background(),
		sqlcdb.CrearClienteParams{
			Nombre:   c.Nombre,
			Telefono: c.Telefono,
		},
	)

	if err != nil {
		return models.Cliente{}
	}

	return aClienteDominio(fila)
}

func (a *AlmacenSQLC) ActualizarCliente(id int, datos models.Cliente) (models.Cliente, bool) {

	fila, err := a.q.ActualizarCliente(
		context.Background(),
		sqlcdb.ActualizarClienteParams{
			Nombre:   datos.Nombre,
			Telefono: datos.Telefono,
			ID:       int64(id),
		},
	)

	if err != nil {
		return models.Cliente{}, false
	}

	return aClienteDominio(fila), true
}

func (a *AlmacenSQLC) BorrarCliente(id int) bool {

	filas, err := a.q.BorrarCliente(
		context.Background(),
		int64(id),
	)

	if err != nil {
		return false
	}

	return filas > 0
}


// =========================================================
// MAPEO SQLC -> DOMINIO
// =========================================================

func aReclamoDominio(r sqlcdb.Reclamo) models.Reclamo {
	return models.Reclamo{
		ID:          int(r.ID),
		ClienteID:   int(r.ClienteID),
		PedidoID:    int(r.PedidoID),
		Descripcion: r.Descripcion,
		Estado:      r.Estado,
	}
}

// =========================================================
// RECLAMOS
// =========================================================

func (a *AlmacenSQLC) ListarReclamos() []models.Reclamo {

	filas, err := a.q.ListarReclamos(context.Background())

	if err != nil {
		return nil
	}

	out := make([]models.Reclamo, 0, len(filas))

	for _, fila := range filas {
		out = append(out, aReclamoDominio(fila))
	}

	return out
}

func (a *AlmacenSQLC) BuscarReclamoPorID(id int) (models.Reclamo, bool) {

	fila, err := a.q.BuscarReclamoPorID(
		context.Background(),
		int64(id),
	)

	if err != nil {
		return models.Reclamo{}, false
	}

	return aReclamoDominio(fila), true
}

func (a *AlmacenSQLC) CrearReclamo(r models.Reclamo) models.Reclamo {

	fila, err := a.q.CrearReclamo(
		context.Background(),
		sqlcdb.CrearReclamoParams{
			ClienteID:   int64(r.ClienteID),
			PedidoID:    int64(r.PedidoID),
			Descripcion: r.Descripcion,
			Estado:      r.Estado,
		},
	)

	if err != nil {
		return models.Reclamo{}
	}

	return aReclamoDominio(fila)
}

func (a *AlmacenSQLC) ActualizarReclamo(id int, datos models.Reclamo) (models.Reclamo, bool) {

	fila, err := a.q.ActualizarReclamo(
		context.Background(),
		sqlcdb.ActualizarReclamoParams{
			ClienteID:   int64(datos.ClienteID),
			PedidoID:    int64(datos.PedidoID),
			Descripcion: datos.Descripcion,
			Estado:      datos.Estado,
			ID:          int64(id),
		},
	)

	if err != nil {
		return models.Reclamo{}, false
	}

	return aReclamoDominio(fila), true
}

func (a *AlmacenSQLC) BorrarReclamo(id int) bool {

	filas, err := a.q.BorrarReclamo(
		context.Background(),
		int64(id),
	)

	if err != nil {
		return false
	}

	return filas > 0
}


// =========================================================
// MAPEO SQLC -> DOMINIO
// =========================================================

func aSeguimientoPedidoDominio(s sqlcdb.Seguimientopedido) models.SeguimientoPedido {
	return models.SeguimientoPedido{
		ID:           int(s.ID),
		PedidoID:     int(s.PedidoID),
		Estado:       s.Estado,
		FechaEstado:  s.FechaEstado,
	}
}

// =========================================================
// SEGUIMIENTOS DE PEDIDO
// =========================================================

func (a *AlmacenSQLC) ListarSeguimientosPedido() []models.SeguimientoPedido {

	filas, err := a.q.ListarSeguimientosPedido(context.Background())

	if err != nil {
		return nil
	}

	out := make([]models.SeguimientoPedido, 0, len(filas))

	for _, fila := range filas {
		out = append(out, aSeguimientoPedidoDominio(fila))
	}

	return out
}

func (a *AlmacenSQLC) BuscarSeguimientoPedidoPorID(id int) (models.SeguimientoPedido, bool) {

	fila, err := a.q.BuscarSeguimientoPedidoPorID(
		context.Background(),
		int64(id),
	)

	if err != nil {
		return models.SeguimientoPedido{}, false
	}

	return aSeguimientoPedidoDominio(fila), true
}

func (a *AlmacenSQLC) CrearSeguimientoPedido(s models.SeguimientoPedido) models.SeguimientoPedido {

	fila, err := a.q.CrearSeguimientoPedido(
		context.Background(),
		sqlcdb.CrearSeguimientoPedidoParams{
			PedidoID:     int64(s.PedidoID),
			Estado:       s.Estado,
			FechaEstado:  s.FechaEstado,
		},
	)

	if err != nil {
		return models.SeguimientoPedido{}
	}

	return aSeguimientoPedidoDominio(fila)
}

func (a *AlmacenSQLC) ActualizarSeguimientoPedido(id int, datos models.SeguimientoPedido) (models.SeguimientoPedido, bool) {

	fila, err := a.q.ActualizarSeguimientoPedido(
		context.Background(),
		sqlcdb.ActualizarSeguimientoPedidoParams{
			PedidoID:     int64(datos.PedidoID),
			Estado:       datos.Estado,
			FechaEstado:  datos.FechaEstado,
			ID:           int64(id),
		},
	)

	if err != nil {
		return models.SeguimientoPedido{}, false
	}

	return aSeguimientoPedidoDominio(fila), true
}

func (a *AlmacenSQLC) BorrarSeguimientoPedido(id int) bool {

	filas, err := a.q.BorrarSeguimientoPedido(
		context.Background(),
		int64(id),
	)

	if err != nil {
		return false
	}

	return filas > 0
}


// Verificación en tiempo de compilación
var _ Almacen = (*AlmacenSQLC)(nil)

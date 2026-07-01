package storage

import (
	"gorm.io/gorm"

	"proyecto-detallesPersonalizados-api/internal/models"
)

type AlmacenSQLite struct {
	db *gorm.DB
}

func NuevoAlmacenSQLite(db *gorm.DB) *AlmacenSQLite {
	return &AlmacenSQLite{db: db}
}





// =========================================================
// PEDIDOS
// =========================================================

func (a *AlmacenSQLite) ListarPedidos() []models.Pedido {

	var pedidos []models.Pedido

	a.db.Find(&pedidos)

	return pedidos
}

func (a *AlmacenSQLite) BuscarPedidoPorID(id int) (models.Pedido, bool) {

	var pedido models.Pedido

	if err := a.db.First(&pedido, id).Error; err != nil {
		return models.Pedido{}, false
	}

	return pedido, true
}

func (a *AlmacenSQLite) CrearPedido(p models.Pedido) models.Pedido {

	a.db.Create(&p)

	return p
}

func (a *AlmacenSQLite) ActualizarPedido(id int, datos models.Pedido) (models.Pedido, bool) {

	var existente models.Pedido

	if err := a.db.First(&existente, id).Error; err != nil {
		return models.Pedido{}, false
	}

	datos.ID = id

	a.db.Save(&datos)

	return datos, true
}

func (a *AlmacenSQLite) BorrarPedido(id int) bool {

	res := a.db.Delete(&models.Pedido{}, id)

	return res.RowsAffected > 0
}

// =========================================================
// PERSONALIZACIONES
// =========================================================

func (a *AlmacenSQLite) ListarPersonalizaciones() []models.Personalizacion {

	var personalizaciones []models.Personalizacion

	a.db.Find(&personalizaciones)

	return personalizaciones
}

func (a *AlmacenSQLite) BuscarPersonalizacionPorID(id int) (models.Personalizacion, bool) {

	var personalizacion models.Personalizacion

	if err := a.db.First(&personalizacion, id).Error; err != nil {
		return models.Personalizacion{}, false
	}

	return personalizacion, true
}

func (a *AlmacenSQLite) CrearPersonalizacion(p models.Personalizacion) models.Personalizacion {

	a.db.Create(&p)

	return p
}

func (a *AlmacenSQLite) ActualizarPersonalizacion(id int, datos models.Personalizacion) (models.Personalizacion, bool) {

	var existente models.Personalizacion

	if err := a.db.First(&existente, id).Error; err != nil {
		return models.Personalizacion{}, false
	}

	datos.ID = id

	a.db.Save(&datos)

	return datos, true
}

func (a *AlmacenSQLite) BorrarPersonalizacion(id int) bool {

	res := a.db.Delete(&models.Personalizacion{}, id)

	return res.RowsAffected > 0
}


// =========================================================
// PRODUCTOS PERSONALIZADOS
// =========================================================

func (a *AlmacenSQLite) ListarProductosPersonalizados() []models.ProductoPersonalizado {

	var productos []models.ProductoPersonalizado

	a.db.Find(&productos)

	return productos
}

func (a *AlmacenSQLite) BuscarProductoPersonalizadoPorID(id int) (models.ProductoPersonalizado, bool) {

	var producto models.ProductoPersonalizado

	if err := a.db.First(&producto, id).Error; err != nil {
		return models.ProductoPersonalizado{}, false
	}

	return producto, true
}

func (a *AlmacenSQLite) CrearProductoPersonalizado(p models.ProductoPersonalizado) models.ProductoPersonalizado {

	a.db.Create(&p)

	return p
}

func (a *AlmacenSQLite) ActualizarProductoPersonalizado(id int, datos models.ProductoPersonalizado) (models.ProductoPersonalizado, bool) {

	var existente models.ProductoPersonalizado

	if err := a.db.First(&existente, id).Error; err != nil {
		return models.ProductoPersonalizado{}, false
	}

	datos.ID = id

	a.db.Save(&datos)

	return datos, true
}

func (a *AlmacenSQLite) BorrarProductoPersonalizado(id int) bool {

	res := a.db.Delete(&models.ProductoPersonalizado{}, id)

	return res.RowsAffected > 0
}


// =========================================================
// SOLICITUDES URGENTES
// =========================================================

func (a *AlmacenSQLite) ListarSolicitudesUrgentes() []models.SolicitudUrgente {

	var solicitudes []models.SolicitudUrgente

	a.db.Find(&solicitudes)

	return solicitudes
}

func (a *AlmacenSQLite) BuscarSolicitudUrgentePorID(id int) (models.SolicitudUrgente, bool) {

	var solicitud models.SolicitudUrgente

	if err := a.db.First(&solicitud, id).Error; err != nil {
		return models.SolicitudUrgente{}, false
	}

	return solicitud, true
}

func (a *AlmacenSQLite) CrearSolicitudUrgente(s models.SolicitudUrgente) models.SolicitudUrgente {

	a.db.Create(&s)

	return s
}

func (a *AlmacenSQLite) ActualizarSolicitudUrgente(id int, datos models.SolicitudUrgente) (models.SolicitudUrgente, bool) {

	var existente models.SolicitudUrgente

	if err := a.db.First(&existente, id).Error; err != nil {
		return models.SolicitudUrgente{}, false
	}

	datos.ID = id

	a.db.Save(&datos)

	return datos, true
}

func (a *AlmacenSQLite) BorrarSolicitudUrgente(id int) bool {

	res := a.db.Delete(&models.SolicitudUrgente{}, id)

	return res.RowsAffected > 0
}

// =========================================================
// AGENDAS DE PRODUCCIÓN
// =========================================================

func (a *AlmacenSQLite) ListarAgendasProduccion() []models.AgendaProduccion {

	var agendas []models.AgendaProduccion

	a.db.Find(&agendas)

	return agendas
}

func (a *AlmacenSQLite) BuscarAgendaProduccionPorID(id int) (models.AgendaProduccion, bool) {

	var agenda models.AgendaProduccion

	if err := a.db.First(&agenda, id).Error; err != nil {
		return models.AgendaProduccion{}, false
	}

	return agenda, true
}

func (a *AlmacenSQLite) CrearAgendaProduccion(ag models.AgendaProduccion) models.AgendaProduccion {

	a.db.Create(&ag)

	return ag
}

func (a *AlmacenSQLite) ActualizarAgendaProduccion(id int, datos models.AgendaProduccion) (models.AgendaProduccion, bool) {

	var existente models.AgendaProduccion

	if err := a.db.First(&existente, id).Error; err != nil {
		return models.AgendaProduccion{}, false
	}

	datos.ID = id

	a.db.Save(&datos)

	return datos, true
}

func (a *AlmacenSQLite) BorrarAgendaProduccion(id int) bool {

	res := a.db.Delete(&models.AgendaProduccion{}, id)

	return res.RowsAffected > 0
}

// =========================================================
// SLOTS DE PRODUCCIÓN
// =========================================================

func (a *AlmacenSQLite) ListarSlotsProduccion() []models.SlotProduccion {

	var slots []models.SlotProduccion

	a.db.Find(&slots)

	return slots
}

func (a *AlmacenSQLite) BuscarSlotProduccionPorID(id int) (models.SlotProduccion, bool) {

	var slot models.SlotProduccion

	if err := a.db.First(&slot, id).Error; err != nil {
		return models.SlotProduccion{}, false
	}

	return slot, true
}

func (a *AlmacenSQLite) CrearSlotProduccion(s models.SlotProduccion) models.SlotProduccion {

	a.db.Create(&s)

	return s
}

func (a *AlmacenSQLite) ActualizarSlotProduccion(id int, datos models.SlotProduccion) (models.SlotProduccion, bool) {

	var existente models.SlotProduccion

	if err := a.db.First(&existente, id).Error; err != nil {
		return models.SlotProduccion{}, false
	}

	datos.ID = id

	a.db.Save(&datos)

	return datos, true
}

func (a *AlmacenSQLite) BorrarSlotProduccion(id int) bool {

	res := a.db.Delete(&models.SlotProduccion{}, id)

	return res.RowsAffected > 0
}



// =========================================================
// CLIENTES
// =========================================================

func (a *AlmacenSQLite) ListarClientes() []models.Cliente {

	var clientes []models.Cliente

	a.db.Find(&clientes)

	return clientes
}

func (a *AlmacenSQLite) BuscarClientePorID(id int) (models.Cliente, bool) {

	var cliente models.Cliente

	if err := a.db.First(&cliente, id).Error; err != nil {
		return models.Cliente{}, false
	}

	return cliente, true
}

func (a *AlmacenSQLite) CrearCliente(c models.Cliente) models.Cliente {

	a.db.Create(&c)

	return c
}

func (a *AlmacenSQLite) ActualizarCliente(id int, datos models.Cliente) (models.Cliente, bool) {

	var existente models.Cliente

	if err := a.db.First(&existente, id).Error; err != nil {
		return models.Cliente{}, false
	}

	datos.ID = id

	a.db.Save(&datos)

	return datos, true
}

func (a *AlmacenSQLite) BorrarCliente(id int) bool {

	res := a.db.Delete(&models.Cliente{}, id)

	return res.RowsAffected > 0
}

// =========================================================
// RECLAMOS
// =========================================================

func (a *AlmacenSQLite) ListarReclamos() []models.Reclamo {

	var reclamos []models.Reclamo

	a.db.Find(&reclamos)

	return reclamos
}

func (a *AlmacenSQLite) BuscarReclamoPorID(id int) (models.Reclamo, bool) {

	var reclamo models.Reclamo

	if err := a.db.First(&reclamo, id).Error; err != nil {
		return models.Reclamo{}, false
	}

	return reclamo, true
}

func (a *AlmacenSQLite) CrearReclamo(r models.Reclamo) models.Reclamo {

	a.db.Create(&r)

	return r
}

func (a *AlmacenSQLite) ActualizarReclamo(id int, datos models.Reclamo) (models.Reclamo, bool) {

	var existente models.Reclamo

	if err := a.db.First(&existente, id).Error; err != nil {
		return models.Reclamo{}, false
	}

	datos.ID = id

	a.db.Save(&datos)

	return datos, true
}

func (a *AlmacenSQLite) BorrarReclamo(id int) bool {

	res := a.db.Delete(&models.Reclamo{}, id)

	return res.RowsAffected > 0
}

// =========================================================
// SEGUIMIENTOS DE PEDIDOS
// =========================================================

func (a *AlmacenSQLite) ListarSeguimientosPedido() []models.SeguimientoPedido {

	var seguimientos []models.SeguimientoPedido

	a.db.Find(&seguimientos)

	return seguimientos
}

func (a *AlmacenSQLite) BuscarSeguimientoPedidoPorID(id int) (models.SeguimientoPedido, bool) {

	var seguimiento models.SeguimientoPedido

	if err := a.db.First(&seguimiento, id).Error; err != nil {
		return models.SeguimientoPedido{}, false
	}

	return seguimiento, true
}

func (a *AlmacenSQLite) CrearSeguimientoPedido(s models.SeguimientoPedido) models.SeguimientoPedido {

	a.db.Create(&s)

	return s
}

func (a *AlmacenSQLite) ActualizarSeguimientoPedido(id int, datos models.SeguimientoPedido) (models.SeguimientoPedido, bool) {

	var existente models.SeguimientoPedido

	if err := a.db.First(&existente, id).Error; err != nil {
		return models.SeguimientoPedido{}, false
	}

	datos.ID = id

	a.db.Save(&datos)

	return datos, true
}

func (a *AlmacenSQLite) BorrarSeguimientoPedido(id int) bool {

	res := a.db.Delete(&models.SeguimientoPedido{}, id)

	return res.RowsAffected > 0
}

var _ Almacen = (*AlmacenSQLite)(nil)


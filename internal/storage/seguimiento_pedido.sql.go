package storage

import "proyecto-detallesPersonalizados-api/internal/models"


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


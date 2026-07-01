package storage

import "proyecto-detallesPersonalizados-api/internal/models"

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
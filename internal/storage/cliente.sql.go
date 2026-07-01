package storage

import "proyecto-detallesPersonalizados-api/internal/models"


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

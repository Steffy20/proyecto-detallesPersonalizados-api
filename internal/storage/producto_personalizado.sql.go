package storage

import "proyecto-detallesPersonalizados-api/internal/models"


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


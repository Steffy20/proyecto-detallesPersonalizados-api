package storage

import "proyecto-detallesPersonalizados-api/internal/models"


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

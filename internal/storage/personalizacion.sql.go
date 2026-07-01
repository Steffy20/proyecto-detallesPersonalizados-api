package storage

import "proyecto-detallesPersonalizados-api/internal/models"

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
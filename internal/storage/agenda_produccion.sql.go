package storage

import "proyecto-detallesPersonalizados-api/internal/models"


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

package storage

import "proyecto-detallesPersonalizados-api/internal/models"

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

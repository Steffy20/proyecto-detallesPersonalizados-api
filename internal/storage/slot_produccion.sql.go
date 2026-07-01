package storage

import "proyecto-detallesPersonalizados-api/internal/models"


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



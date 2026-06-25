package handlers

import (
	"encoding/json"
	"net/http"
	"proyecto-detallesPersonalizados-api/internal/models"
)

func CrearSlotProduccion(w http.ResponseWriter, r *http.Request) {
	var slot models.SlotProduccion
	err := json.NewDecoder(r.Body).Decode(&slot)
	if err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}
}

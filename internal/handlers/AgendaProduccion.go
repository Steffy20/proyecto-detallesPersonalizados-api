package handlers

import (
	"encoding/json"
	"net/http"
	"proyecto-detallesPersonalizados-api/internal/models"
)

func CrearAgendaProduccion(w http.ResponseWriter, r *http.Request) {
	var agenda models.AgendaProduccion
	err := json.NewDecoder(r.Body).Decode(&agenda)
	if err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}
}

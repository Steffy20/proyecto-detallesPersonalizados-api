package handlers

import (
	"encoding/json"
	"net/http"
	"proyecto-detallesPersonalizados-api/internal/models"
	"proyecto-detallesPersonalizados-api/internal/storage"
)

func CrearAgendaProduccion(w http.ResponseWriter, r *http.Request) {
	var agenda models.AgendaProduccion
	err := json.NewDecoder(r.Body).Decode(&agenda)
	if err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}
	// VALIDACIONES
	if agenda.Fecha == "" {
		http.Error(w, "La fecha es obligatoria", http.StatusBadRequest)
		return
	}
	if agenda.Responsable == "" {
		http.Error(w, "El responsable es obligatorio", http.StatusBadRequest)
		return
	}
	if agenda.Estado == "" {
		http.Error(w, "El estado es obligatorio", http.StatusBadRequest)
		return
	}
	// guardar en memoria
	agenda.ID = storage.AgendaProduccionID
	storage.AgendaProduccionID++
	storage.AgendasProduccion = append(
		storage.AgendasProduccion,
		agenda,
	)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(agenda)
}

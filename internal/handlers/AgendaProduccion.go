package handlers

import (
	"encoding/json"
	"net/http"
	"proyecto-detallesPersonalizados-api/internal/models"
	"proyecto-detallesPersonalizados-api/internal/storage"
	"strconv"

	"github.com/go-chi/chi/v5"
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
func ObtenerAgendasProduccion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(storage.AgendasProduccion)
}
func ObtenerAgendaProduccionPorID(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	for _, agenda := range storage.AgendasProduccion {
		if agenda.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(agenda)
			return
		}
	}
	http.Error(w, "Agenda de producción no encontrada", http.StatusNotFound)
}
func ActualizarAgendaProduccion(w http.ResponseWriter, r *http.Request) { 
idParam := chi.URLParam(r, "id") 
id, err := strconv.Atoi(idParam) 
if err != nil { 
http.Error(w, "ID inválido", http.StatusBadRequest) 
return 
} 
var agendaActualizada models.AgendaProduccion 
err = json.NewDecoder(r.Body).Decode(&agendaActualizada) 
if err != nil { 
http.Error(w, "Datos inválidos", http.StatusBadRequest) 
return 
} 
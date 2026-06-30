package handlers

import (
	"encoding/json"
	"net/http"
	"proyecto-detallesPersonalizados-api/internal/models"
	"proyecto-detallesPersonalizados-api/internal/storage"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func CrearSlotProduccion(w http.ResponseWriter, r *http.Request) {
	var slot models.SlotProduccion
	err := json.NewDecoder(r.Body).Decode(&slot)
	if err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}
	// Validar que la agenda exista
	agendaExiste := false

	for _, agenda := range storage.AgendasProduccion {
		if agenda.ID == slot.AgendaID {
			agendaExiste = true
			break
		}
	}
	if !agendaExiste {
		http.Error(w, "La agenda asociada no existe",
			http.StatusBadRequest)
		return
	}
	// Guardar en memoria
	slot.ID = storage.SlotProduccionID
	storage.SlotProduccionID++
	storage.SlotsProduccion = append(
		storage.SlotsProduccion,
		slot,
	)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(slot)
}
func ObtenerSlotsProduccion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(storage.SlotsProduccion)
}
func ObtenerSlotProduccionPorID(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	for _, slot := range storage.SlotsProduccion {
		if slot.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(slot)
			return
		}
	}
	http.Error(w, "Slot de producción no encontrado", http.StatusNotFound)
}
func ActualizarSlotProduccion(w http.ResponseWriter, r *http.Request) { 
idParam := chi.URLParam(r, "id") 
id, err := strconv.Atoi(idParam) 
if err != nil { 
http.Error(w, "ID inválido", http.StatusBadRequest) 
return 
} 
var slotActualizado models.SlotProduccion 
err = json.NewDecoder(r.Body).Decode(&slotActualizado) 
if err != nil { 
http.Error(w, "Datos inválidos", http.StatusBadRequest) 
return 
} 
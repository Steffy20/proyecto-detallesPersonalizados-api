package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"proyecto-detallesPersonalizados-api/internal/models"
	"proyecto-detallesPersonalizados-api/internal/storage"
	"proyecto-detallesPersonalizados-api/internal/service"
)
	var agendaProduccionService = service.NewAgendaProduccionService()

func CrearAgendaProduccion(w http.ResponseWriter, r *http.Request) {

	var agenda models.AgendaProduccion

	err := json.NewDecoder(r.Body).Decode(&agenda)

	if err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	err = agendaProduccionService.ValidarAgendaProduccion(
		&agenda,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

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

	// VALIDACIONES
	if agendaActualizada.Fecha == "" {
		http.Error(w, "La fecha es obligatoria", http.StatusBadRequest)
		return
	}

	if agendaActualizada.Responsable == "" {
		http.Error(w, "El responsable es obligatorio", http.StatusBadRequest)
		return
	}

	if agendaActualizada.Estado == "" {
		http.Error(w, "El estado es obligatorio", http.StatusBadRequest)
		return
	}

	for i, agenda := range storage.AgendasProduccion {

		if agenda.ID == id {

			agendaActualizada.ID = id

			storage.AgendasProduccion[i] = agendaActualizada

			w.Header().Set("Content-Type", "application/json")

			json.NewEncoder(w).Encode(agendaActualizada)

			return
		}
	}

	http.Error(w, "Agenda de producción no encontrada", http.StatusNotFound)
}


func EliminarAgendaProduccion(w http.ResponseWriter, r *http.Request) {

	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	for i, agenda := range storage.AgendasProduccion {

		if agenda.ID == id {

			storage.AgendasProduccion = append(
				storage.AgendasProduccion[:i],
				storage.AgendasProduccion[i+1:]...,
			)

			w.WriteHeader(http.StatusOK)

			w.Write([]byte("Agenda de producción eliminada"))

			return
		}
	}

	http.Error(w, "Agenda de producción no encontrada", http.StatusNotFound)
}

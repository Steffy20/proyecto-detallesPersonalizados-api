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
	var solicitudUrgenteService = service.NewSolicitudUrgenteService()

func CrearSolicitudUrgente(w http.ResponseWriter, r *http.Request) {

	var solicitud models.SolicitudUrgente

	err := json.NewDecoder(r.Body).Decode(&solicitud)

	if err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	err = solicitudUrgenteService.ValidarSolicitudUrgente(
		&solicitud,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	solicitud.ID = storage.SolicitudUrgenteID
	storage.SolicitudUrgenteID++

	storage.SolicitudesUrgentes = append(
		storage.SolicitudesUrgentes,
		solicitud,
	)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(solicitud)
}

func ObtenerSolicitudesUrgentes(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(storage.SolicitudesUrgentes)
}


func ObtenerSolicitudUrgentePorID(w http.ResponseWriter, r *http.Request) {

	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	for _, solicitud := range storage.SolicitudesUrgentes {

		if solicitud.ID == id {

			w.Header().Set("Content-Type", "application/json")

			json.NewEncoder(w).Encode(solicitud)

			return
		}
	}

	http.Error(w, "Solicitud urgente no encontrada", http.StatusNotFound)
}


func ActualizarSolicitudUrgente(w http.ResponseWriter, r *http.Request) {

	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var solicitudActualizada models.SolicitudUrgente

	err = json.NewDecoder(r.Body).Decode(&solicitudActualizada)

	if err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	// VALIDACIONES
	if solicitudActualizada.Cliente == "" {
		http.Error(w, "Cliente obligatorio", http.StatusBadRequest)
		return
	}

	if solicitudActualizada.Descripcion == "" {
		http.Error(w, "Descripción obligatoria", http.StatusBadRequest)
		return
	}

	if solicitudActualizada.FechaRequerida == "" {
		http.Error(w, "Fecha requerida obligatoria", http.StatusBadRequest)
		return
	}

	for i, solicitud := range storage.SolicitudesUrgentes {

		if solicitud.ID == id {

			solicitudActualizada.ID = id

			storage.SolicitudesUrgentes[i] = solicitudActualizada

			w.Header().Set("Content-Type", "application/json")

			json.NewEncoder(w).Encode(solicitudActualizada)

			return
		}
	}

	http.Error(w, "Solicitud urgente no encontrada", http.StatusNotFound)
}


func EliminarSolicitudUrgente(w http.ResponseWriter, r *http.Request) {

	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	for i, solicitud := range storage.SolicitudesUrgentes {

		if solicitud.ID == id {

			storage.SolicitudesUrgentes = append(
				storage.SolicitudesUrgentes[:i],
				storage.SolicitudesUrgentes[i+1:]...,
			)

			w.WriteHeader(http.StatusOK)

			w.Write([]byte("Solicitud urgente eliminada"))

			return
		}
	}

	http.Error(w, "Solicitud urgente no encontrada", http.StatusNotFound)
}

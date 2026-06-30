package handlers

import (
	"encoding/json"
	"net/http"
	"proyecto-detallesPersonalizados-api/internal/models"
	"proyecto-detallesPersonalizados-api/internal/storage"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func CrearSolicitudUrgente(w http.ResponseWriter, r *http.Request) {
	var solicitud models.SolicitudUrgente
	err := json.NewDecoder(r.Body).Decode(&solicitud)
	if err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}
	// VALIDACIONES
	if solicitud.Cliente == "" {
		http.Error(w, "Cliente obligatorio", http.StatusBadRequest)
		return
	}
	if solicitud.Descripcion == "" {
		http.Error(w, "Descripción obligatoria", http.StatusBadRequest)
		return
	}
	if solicitud.FechaRequerida == "" {
		http.Error(w, "Fecha requerida obligatoria", http.StatusBadRequest)
		return
	}
	if solicitud.Estado == "" {
		solicitud.Estado = "Pendiente"
	}

	// guardar en memoria
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

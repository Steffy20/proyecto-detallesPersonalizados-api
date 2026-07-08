package handlers

import (
	"encoding/json"
	"net/http"

	"proyecto-detallesPersonalizados-api/internal/models"
)

// ===================== CREAR =====================

func (s *Server) CrearSolicitudUrgente(w http.ResponseWriter, r *http.Request) {
	var solicitud models.SolicitudUrgente

	if err := json.NewDecoder(r.Body).Decode(&solicitud); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	creada, err := s.SolicitudesUrgentes.Crear(solicitud)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(creada)
}

// ===================== LISTAR =====================

func (s *Server) ObtenerSolicitudesUrgentes(w http.ResponseWriter, r *http.Request) {
	solicitudes := s.SolicitudesUrgentes.Listar()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(solicitudes)
}

// ===================== OBTENER =====================

func (s *Server) ObtenerSolicitudUrgentePorID(w http.ResponseWriter, r *http.Request) {
	id, err := idDeURL(r)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	solicitud, err := s.SolicitudesUrgentes.Obtener(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(solicitud)
}

// ===================== ACTUALIZAR =====================

func (s *Server) ActualizarSolicitudUrgente(w http.ResponseWriter, r *http.Request) {
	id, err := idDeURL(r)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var solicitud models.SolicitudUrgente
	if err := json.NewDecoder(r.Body).Decode(&solicitud); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	actualizada, err := s.SolicitudesUrgentes.Actualizar(id, solicitud)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(actualizada)
}

// ===================== ELIMINAR =====================

func (s *Server) EliminarSolicitudUrgente(w http.ResponseWriter, r *http.Request) {
	id, err := idDeURL(r)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	if err := s.SolicitudesUrgentes.Borrar(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

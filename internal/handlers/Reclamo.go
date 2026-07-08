package handlers

import (
	"encoding/json"
	"net/http"

	"proyecto-detallesPersonalizados-api/internal/models"
)

// ===================== CREAR =====================

func (s *Server) CrearReclamo(w http.ResponseWriter, r *http.Request) {

	var reclamo models.Reclamo

	if err := json.NewDecoder(r.Body).Decode(&reclamo); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	creado, err := s.Reclamos.Crear(reclamo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(creado)
}

// ===================== LISTAR =====================

func (s *Server) ObtenerReclamos(w http.ResponseWriter, r *http.Request) {

	reclamos := s.Reclamos.Listar()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reclamos)
}

// ===================== OBTENER =====================

func (s *Server) ObtenerReclamoPorID(w http.ResponseWriter, r *http.Request) {

	id, err := idDeURL(r)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	reclamo, err := s.Reclamos.Obtener(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reclamo)
}

// ===================== ACTUALIZAR =====================

func (s *Server) ActualizarReclamo(w http.ResponseWriter, r *http.Request) {

	id, err := idDeURL(r)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var reclamo models.Reclamo

	if err := json.NewDecoder(r.Body).Decode(&reclamo); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	actualizado, err := s.Reclamos.Actualizar(id, reclamo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(actualizado)
}

// ===================== ELIMINAR =====================

func (s *Server) EliminarReclamo(w http.ResponseWriter, r *http.Request) {

	id, err := idDeURL(r)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	if err := s.Reclamos.Borrar(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

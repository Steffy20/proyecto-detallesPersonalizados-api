package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"proyecto-detallesPersonalizados-api/internal/models"
)

// ===================== CREAR =====================

func (s *Server) CrearPersonalizacion(w http.ResponseWriter, r *http.Request) {

	var personalizacion models.Personalizacion

	if err := json.NewDecoder(r.Body).Decode(&personalizacion); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	creada, err := s.Personalizaciones.Crear(personalizacion)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(creada)
}

// ===================== LISTAR =====================

func (s *Server) ObtenerPersonalizaciones(w http.ResponseWriter, r *http.Request) {

	personalizaciones := s.Personalizaciones.Listar()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(personalizaciones)
}

// ===================== OBTENER =====================

func (s *Server) ObtenerPersonalizacionPorID(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	personalizacion, err := s.Personalizaciones.Obtener(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(personalizacion)
}

// ===================== ACTUALIZAR =====================

func (s *Server) ActualizarPersonalizacion(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var personalizacion models.Personalizacion

	if err := json.NewDecoder(r.Body).Decode(&personalizacion); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	actualizada, err := s.Personalizaciones.Actualizar(id, personalizacion)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(actualizada)
}

// ===================== ELIMINAR =====================

func (s *Server) EliminarPersonalizacion(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	if err := s.Personalizaciones.Borrar(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

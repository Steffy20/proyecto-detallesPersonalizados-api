package handlers

import (
	"encoding/json"
	"net/http"

	"proyecto-detallesPersonalizados-api/internal/models"
)

// ===================== CREAR =====================

func (s *Server) CrearSeguimientoPedido(w http.ResponseWriter, r *http.Request) {

	var seguimiento models.SeguimientoPedido

	if err := json.NewDecoder(r.Body).Decode(&seguimiento); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	creado, err := s.Seguimientos.Crear(seguimiento)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(creado)
}

// ===================== LISTAR =====================

func (s *Server) ObtenerSeguimientosPedido(w http.ResponseWriter, r *http.Request) {

	seguimientos := s.Seguimientos.Listar()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(seguimientos)
}

// ===================== OBTENER =====================

func (s *Server) ObtenerSeguimientoPedidoPorID(w http.ResponseWriter, r *http.Request) {

	id, err := idDeURL(r)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	seguimiento, err := s.Seguimientos.Obtener(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(seguimiento)
}

// ===================== ACTUALIZAR =====================

func (s *Server) ActualizarSeguimientoPedido(w http.ResponseWriter, r *http.Request) {

	id, err := idDeURL(r)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var seguimiento models.SeguimientoPedido

	if err := json.NewDecoder(r.Body).Decode(&seguimiento); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	actualizado, err := s.Seguimientos.Actualizar(id, seguimiento)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(actualizado)
}

// ===================== ELIMINAR =====================

func (s *Server) EliminarSeguimientoPedido(w http.ResponseWriter, r *http.Request) {

	id, err := idDeURL(r)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	if err := s.Seguimientos.Borrar(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

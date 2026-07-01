package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"proyecto-detallesPersonalizados-api/internal/models"
)


// ===================== CREAR =====================

func (s *Server) CrearProductoPersonalizado(w http.ResponseWriter, r *http.Request) {

	var producto models.ProductoPersonalizado

	if err := json.NewDecoder(r.Body).Decode(&producto); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	creado, err := s.ProductosPersonalizados.Crear(producto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(creado)
}

// ===================== LISTAR =====================

func (s *Server) ObtenerProductosPersonalizados(w http.ResponseWriter, r *http.Request) {

	productos := s.ProductosPersonalizados.Listar()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(productos)
}

// ===================== OBTENER =====================

func (s *Server) ObtenerProductoPersonalizadoPorID(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	producto, err := s.ProductosPersonalizados.Obtener(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(producto)
}

// ===================== ACTUALIZAR =====================

func (s *Server) ActualizarProductoPersonalizado(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var producto models.ProductoPersonalizado

	if err := json.NewDecoder(r.Body).Decode(&producto); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	actualizado, err := s.ProductosPersonalizados.Actualizar(id, producto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(actualizado)
}

// ===================== ELIMINAR =====================

func (s *Server) EliminarProductoPersonalizado(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	if err := s.ProductosPersonalizados.Borrar(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"proyecto-detallesPersonalizados-api/internal/models"
)// ===================== CREAR =====================

func (s *Server) CrearSlotProduccion(w http.ResponseWriter, r *http.Request) {

	var slot models.SlotProduccion

	if err := json.NewDecoder(r.Body).Decode(&slot); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	creado, err := s.SlotsProduccion.Crear(slot)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(creado)
}

// ===================== LISTAR =====================

func (s *Server) ObtenerSlotsProduccion(w http.ResponseWriter, r *http.Request) {

	slots := s.SlotsProduccion.Listar()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(slots)
}

// ===================== OBTENER =====================

func (s *Server) ObtenerSlotProduccionPorID(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	slot, err := s.SlotsProduccion.Obtener(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(slot)
}

// ===================== ACTUALIZAR =====================

func (s *Server) ActualizarSlotProduccion(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var slot models.SlotProduccion

	if err := json.NewDecoder(r.Body).Decode(&slot); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	actualizado, err := s.SlotsProduccion.Actualizar(id, slot)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(actualizado)
}

// ===================== ELIMINAR =====================

func (s *Server) EliminarSlotProduccion(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	if err := s.SlotsProduccion.Borrar(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

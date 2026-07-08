package handlers

import (
	"encoding/json"
	"net/http"

	"proyecto-detallesPersonalizados-api/internal/models"
)

// ===================== CREAR =====================

func (s *Server) CrearPedido(w http.ResponseWriter, r *http.Request) {

	var pedido models.Pedido

	if err := json.NewDecoder(r.Body).Decode(&pedido); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	creado, err := s.Pedidos.Crear(pedido)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(creado)
}

// ===================== LISTAR =====================

func (s *Server) ObtenerPedidos(w http.ResponseWriter, r *http.Request) {

	pedidos := s.Pedidos.Listar()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pedidos)
}

// ===================== OBTENER =====================

func (s *Server) ObtenerPedidoPorID(w http.ResponseWriter, r *http.Request) {

	id, err := idDeURL(r)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	pedido, err := s.Pedidos.Obtener(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pedido)
}

// ===================== ACTUALIZAR =====================

func (s *Server) ActualizarPedido(w http.ResponseWriter, r *http.Request) {

	id, err := idDeURL(r)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var pedido models.Pedido

	if err := json.NewDecoder(r.Body).Decode(&pedido); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	actualizado, err := s.Pedidos.Actualizar(id, pedido)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(actualizado)
}

// ===================== ELIMINAR =====================

func (s *Server) EliminarPedido(w http.ResponseWriter, r *http.Request) {

	id, err := idDeURL(r)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	if err := s.Pedidos.Borrar(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

package handlers

import (
	"encoding/json"
	"net/http"

	"proyecto-detallesPersonalizados-api/internal/models"
)

// ===================== CREAR =====================

func (s *Server) CrearCliente(w http.ResponseWriter, r *http.Request) {

	var cliente models.Cliente

	if err := json.NewDecoder(r.Body).Decode(&cliente); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	creado, err := s.Clientes.Crear(cliente)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(creado)
}

// ===================== LISTAR =====================

func (s *Server) ObtenerClientes(w http.ResponseWriter, r *http.Request) {

	clientes := s.Clientes.Listar()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(clientes)
}

// ===================== OBTENER =====================

func (s *Server) ObtenerClientePorID(w http.ResponseWriter, r *http.Request) {

	id, err := idDeURL(r)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	cliente, err := s.Clientes.Obtener(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cliente)
}

// ===================== ACTUALIZAR =====================

func (s *Server) ActualizarCliente(w http.ResponseWriter, r *http.Request) {

	id, err := idDeURL(r)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var cliente models.Cliente

	if err := json.NewDecoder(r.Body).Decode(&cliente); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	actualizado, err := s.Clientes.Actualizar(id, cliente)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(actualizado)
}

// ===================== ELIMINAR =====================

func (s *Server) EliminarCliente(w http.ResponseWriter, r *http.Request) {

	id, err := idDeURL(r)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	if err := s.Clientes.Borrar(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

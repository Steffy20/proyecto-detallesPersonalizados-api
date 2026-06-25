package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"proyecto-detallesPersonalizados-api/internal/models"
	"proyecto-detallesPersonalizados-api/internal/storage"
)

func CrearCliente(w http.ResponseWriter, r *http.Request) {
	var cliente models.Cliente
	err := json.NewDecoder(r.Body).Decode(&cliente)
	if err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}
	if cliente.Nombre == "" {
		http.Error(w, "Nombre obligatorio", http.StatusBadRequest)
		return
	}
	if cliente.Telefono == "" {
		http.Error(w, "Teléfono obligatorio", http.StatusBadRequest)
		return
	}
	cliente.ID = storage.ClienteID
	storage.ClienteID++
	storage.Clientes = append(storage.Clientes, cliente)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(cliente)
}
func ObtenerClientes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(storage.Clientes)
}

//obtener cliente por Id
func ObtenerClientePorID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	for _, cliente := range storage.Clientes {
		if cliente.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(cliente)
			return
		}
	}
	http.Error(w, "Cliente no encontrado", http.StatusNotFound)
}

//actualizar cliente 
func ActualizarCliente(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	var clienteActualizado models.Cliente
	err = json.NewDecoder(r.Body).Decode(&clienteActualizado)
	if err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}
	if clienteActualizado.Nombre == "" {
		http.Error(w, "Nombre obligatorio", http.StatusBadRequest)
		return
	}
	if clienteActualizado.Telefono == "" {
		http.Error(w, "Teléfono obligatorio", http.StatusBadRequest)
		return
	}
	for i, cliente := range storage.Clientes {
		if cliente.ID == id {
			clienteActualizado.ID = id
			storage.Clientes[i] = clienteActualizado
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(clienteActualizado)
			return
		}
	}
	http.Error(w, "Cliente no encontrado", http.StatusNotFound)
}
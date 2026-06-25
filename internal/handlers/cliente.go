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
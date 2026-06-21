package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"proyecto-detallesPersonalizados-api/internal/models"
	"proyecto-detallesPersonalizados-api/internal/storage"
)

func AsignarPersonalizacionAProducto(w http.ResponseWriter, r *http.Request) {
	var pp models.ProductoPersonalizacion

	if err := json.NewDecoder(r.Body).Decode(&pp); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	err := storage.InsertarProductoPersonalizacion(pp)
	if err != nil {
		http.Error(w, "Error al asignar", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(pp)
}

func ObtenerProductoPersonalizacionPorID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	data, err := storage.ObtenerProductoPersonalizacionPorID(id)
	if err != nil {
		http.Error(w, "No encontrado", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(data)
}


package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"proyecto-detallesPersonalizados-api/internal/models"
	"proyecto-detallesPersonalizados-api/internal/storage"
)
func CrearProductoPersonalizado(w http.ResponseWriter, r *http.Request) {

	var p models.ProductoPersonalizado

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	// VALIDACIONES
	if p.PedidoID <= 0 {
		http.Error(w, "PedidoID obligatorio", http.StatusBadRequest)
		return
	}

	if p.Nombre == "" {
		http.Error(w, "Nombre obligatorio", http.StatusBadRequest)
		return
	}

	if p.Cantidad <= 0 {
		http.Error(w, "Cantidad inválida", http.StatusBadRequest)
		return
	}

	if p.Precio <= 0 {
		http.Error(w, "Precio inválido", http.StatusBadRequest)
		return
	}

	// guardar en memoria
	p.ID = storage.ProductoPersonalizadoID
	storage.ProductoPersonalizadoID++

	storage.ProductosPersonalizados = append(storage.ProductosPersonalizados, p)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)
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

func ListarProductoPersonalizaciones(w http.ResponseWriter, r *http.Request) {
	lista, err := storage.ListarProductoPersonalizaciones()
	if err != nil {
		http.Error(w, "Error al listar", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(lista)
}

func ActualizarProductoPersonalizacion(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	var pp models.ProductoPersonalizacion
	if err := json.NewDecoder(r.Body).Decode(&pp); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	err := storage.ActualizarProductoPersonalizacion(id, pp)
	if err != nil {
		http.Error(w, "Error al actualizar", http.StatusInternalServerError)
		return
	}

	pp.ID = id
	json.NewEncoder(w).Encode(pp)
}

func EliminarProductoPersonalizacion(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	err := storage.EliminarProductoPersonalizacion(id)
	if err != nil {
		http.Error(w, "Error al eliminar", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

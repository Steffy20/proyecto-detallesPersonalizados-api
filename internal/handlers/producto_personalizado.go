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

func ObtenerProductosPersonalizados(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(storage.ProductosPersonalizados)
}

func ObtenerProductoPersonalizadoPorID(w http.ResponseWriter, r *http.Request) {

	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	for _, p := range storage.ProductosPersonalizados {
		if p.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(p)
			return
		}
	}

	http.Error(w, "No encontrado", http.StatusNotFound)
}

func ActualizarProductoPersonalizado(w http.ResponseWriter, r *http.Request) {

	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var p models.ProductoPersonalizado

	err = json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	for i, item := range storage.ProductosPersonalizados {

		if item.ID == id {

			p.ID = id
			storage.ProductosPersonalizados[i] = p

			json.NewEncoder(w).Encode(p)
			return
		}
	}

	http.Error(w, "No encontrado", http.StatusNotFound)
}
func EliminarProductoPersonalizado(w http.ResponseWriter, r *http.Request) {

	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	for i, p := range storage.ProductosPersonalizados {

		if p.ID == id {

			storage.ProductosPersonalizados = append(
				storage.ProductosPersonalizados[:i],
				storage.ProductosPersonalizados[i+1:]...,
			)

			w.Write([]byte("Eliminado correctamente"))
			return
		}
	}

	http.Error(w, "No encontrado", http.StatusNotFound)
}

package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"proyecto-detalles-api/internal/models"
	"proyecto-detalles-api/internal/storage"

	"github.com/go-chi/chi/v5"
)

// COMMIT: IMPLEMENTAR HANDLERS PARA ENTREGAS
func CrearEntrega(w http.ResponseWriter, r *http.Request) {

	var entrega models.Entrega

	err := json.NewDecoder(r.Body).Decode(&entrega)

	if err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	entrega.ID = storage.EntregaID
	storage.EntregaID++

	storage.Entregas = append(storage.Entregas, entrega)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(entrega)
}

// COMMIT: OBTENER TODAS LAS ENTREGAS
func ObtenerEntregas(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(storage.Entregas)
}

// COMMIT: OBTENER ENTREGA POR ID
func ObtenerEntregaPorID(w http.ResponseWriter, r *http.Request) {

	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	for _, entrega := range storage.Entregas {

		if entrega.ID == id {

			w.Header().Set("Content-Type", "application/json")

			json.NewEncoder(w).Encode(entrega)

			return
		}
	}

	http.Error(w, "Entrega no encontrada", http.StatusNotFound)
}

// COMMIT: ACTUALIZAR ENTREGA
func ActualizarEntrega(w http.ResponseWriter, r *http.Request) {

	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var entregaActualizada models.Entrega

	err = json.NewDecoder(r.Body).Decode(&entregaActualizada)

	if err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	for i, entrega := range storage.Entregas {

		if entrega.ID == id {

			entregaActualizada.ID = id

			storage.Entregas[i] = entregaActualizada

			w.Header().Set("Content-Type", "application/json")

			json.NewEncoder(w).Encode(entregaActualizada)

			return
		}
	}

	http.Error(w, "Entrega no encontrada", http.StatusNotFound)
}

package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"proyecto-detallesPersonalizados-api/internal/models"
	"proyecto-detallesPersonalizados-api/internal/storage"

	"github.com/go-chi/chi/v5"
)

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

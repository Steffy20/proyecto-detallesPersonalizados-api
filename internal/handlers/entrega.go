package handlers

import (
	"encoding/json"
	"net/http"

	"proyecto-detalles-api/internal/models"
	"proyecto-detalles-api/internal/storage"
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

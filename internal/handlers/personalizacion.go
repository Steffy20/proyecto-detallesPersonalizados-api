package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"proyecto-detallesPersonalizados-api/internal/models"
	"proyecto-detallesPersonalizados-api/internal/storage"
)

func CrearPersonalizacion(w http.ResponseWriter, r *http.Request) {
	var p models.Personalizacion

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	id, err := storage.InsertarPersonalizacion(p)
	if err != nil {
		http.Error(w, "Error al crear personalización", http.StatusInternalServerError)
		return
	}

	p.ID = id

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)
}
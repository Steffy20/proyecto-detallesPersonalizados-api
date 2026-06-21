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

func ObtenerPersonalizacionPorID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	p, err := storage.ObtenerPersonalizacionPorID(id)
	if err != nil {
		http.Error(w, "No encontrado", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(p)
}

func ListarPersonalizaciones(w http.ResponseWriter, r *http.Request) {
	lista, err := storage.ListarPersonalizaciones()
	if err != nil {
		http.Error(w, "Error al listar", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(lista)
}


func ActualizarPersonalizacion(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	var p models.Personalizacion
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	err := storage.ActualizarPersonalizacion(id, p)
	if err != nil {
		http.Error(w, "Error al actualizar", http.StatusInternalServerError)
		return
	}

	p.ID = id
	json.NewEncoder(w).Encode(p)
}

func EliminarPersonalizacion(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	err := storage.EliminarPersonalizacion(id)
	if err != nil {
		http.Error(w, "Error al eliminar", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

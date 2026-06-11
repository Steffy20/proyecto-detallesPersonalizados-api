package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"proyecto-detalles-api/internal/models"
	"proyecto-detalles-api/internal/storage"
)


//Crear función CrearRepartidor
func CrearRepartidor(w http.ResponseWriter, r *http.Request) {
    var repartidor models.Repartidor
    json.NewDecoder(r.Body).Decode(&repartidor)
}

//gAgregar validacion de datos al crear repartidor
if err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}
repartidor.ID = storage.RepartidorID
	storage.RepartidorID++

	storage.Repartidores = append(storage.Repartidores, repartidor)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(repartidor)
}

func ObtenerRepartidores(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(storage.Repartidores)
}
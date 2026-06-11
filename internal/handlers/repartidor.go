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
//crear obtener repartidores
func ObtenerRepartidores(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(storage.Repartidores)
}

//Crear ObtenerRepartidorPorID
func ObtenerRepartidorPorID(w http.ResponseWriter, r *http.Request) {

	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	for _, repartidor := range storage.Repartidores {

		if repartidor.ID == id {

			w.Header().Set("Content-Type", "application/json")

			json.NewEncoder(w).Encode(repartidor)

			return
		}
	}

	http.Error(w, "Repartidor no encontrado", http.StatusNotFound)
}
//actualizar repartidor
func ActualizarRepartidor(w http.ResponseWriter, r *http.Request) {

	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var repartidorActualizado models.Repartidor

	err = json.NewDecoder(r.Body).Decode(&repartidorActualizado)

	if err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	for i, repartidor := range storage.Repartidores {

		if repartidor.ID == id {

			repartidorActualizado.ID = id

			storage.Repartidores[i] = repartidorActualizado

			w.Header().Set("Content-Type", "application/json")

			json.NewEncoder(w).Encode(repartidorActualizado)

			return
		}
	}

	http.Error(w, "Repartidor no encontrado", http.StatusNotFound)
}
//crear eliminar repartidor
func EliminarRepartidor(w http.ResponseWriter, r *http.Request) {

	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	for i, repartidor := range storage.Repartidores {

		if repartidor.ID == id {

			storage.Repartidores = append(storage.Repartidores[:i], storage.Repartidores[i+1:]...)

			w.WriteHeader(http.StatusOK)

			w.Write([]byte("Repartidor eliminado"))

			return
		}
	}

	http.Error(w, "Repartidor no encontrado", http.StatusNotFound)
}

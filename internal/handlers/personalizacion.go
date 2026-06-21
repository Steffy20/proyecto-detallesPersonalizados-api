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
	var personalizacion models.Personalizacion

	err := json.NewDecoder(r.Body).Decode(&personalizacion)

	if err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

// VALIDACIONES
	if personalizacion.PedidoID == 0 {
		http.Error(w, "PedidoID es obligatorio", http.StatusBadRequest)
		return
	}
if personalizacion.Mensaje == "" {
		http.Error(w, "El mensaje es obligatorio", http.StatusBadRequest)
		return
	}

	if personalizacion.Color == "" {
		http.Error(w, "El color es obligatorio", http.StatusBadRequest)
		return
	}

// Validar que el PedidoID exista
	pedidoExiste := false

	for _, pedido := range storage.Pedidos {
		if pedido.ID == personalizacion.PedidoID {
			pedidoExiste = true
			break
		}
	}

	if !pedidoExiste {
		http.Error(w, "El pedido asociado no existe", http.StatusBadRequest)
		return
	}

    personalizacion.ID = storage.PersonalizacionID
	storage.PersonalizacionID++

    storage.Personalizaciones = append(
	    storage.Personalizaciones, 
	    personalizacion,
)

w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusCreated)

json.NewEncoder(w).Encode(personalizacion)
}


func ObtenerPersonalizaciones(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(storage.Personalizaciones)
}

func ObtenerPersonalizacionPorID(w http.ResponseWriter, r *http.Request) {

	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	for _, p := range storage.Personalizaciones {
		if p.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(p)
			return
		}
	}

	http.Error(w, "No encontrado", http.StatusNotFound)
}

func ActualizarPersonalizacion(w http.ResponseWriter, r *http.Request) {

	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var p models.Personalizacion

	err = json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	for i, item := range storage.Personalizaciones {

		if item.ID == id {

			p.ID = id
			storage.Personalizaciones[i] = p

			json.NewEncoder(w).Encode(p)
			return
		}
	}

	http.Error(w, "No encontrado", http.StatusNotFound)
}

func EliminarPersonalizacion(w http.ResponseWriter, r *http.Request) {

	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	for i, p := range storage.Personalizaciones {

		if p.ID == id {

			storage.Personalizaciones = append(
				storage.Personalizaciones[:i],
				storage.Personalizaciones[i+1:]...,
			)

			w.Write([]byte("Eliminado correctamente"))
			return
		}
	}

	http.Error(w, "No encontrado", http.StatusNotFound)
}

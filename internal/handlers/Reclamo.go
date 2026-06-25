package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"proyecto-detallesPersonalizados-api/internal/models"
	"proyecto-detallesPersonalizados-api/internal/storage"
	"proyecto-detallesPersonalizados-api/internal/service"
)
	var reclamoService = service.NewReclamoService()

func CrearReclamo(w http.ResponseWriter, r *http.Request) {

	var reclamo models.Reclamo

	err := json.NewDecoder(r.Body).Decode(&reclamo)

	if err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	err = reclamoService.ValidarReclamo(&reclamo)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validar cliente
	clienteExiste := false

	for _, cliente := range storage.Clientes {
		if cliente.ID == reclamo.ClienteID {
			clienteExiste = true
			break
		}
	}

	if !clienteExiste {
		http.Error(w, "El cliente asociado no existe", http.StatusBadRequest)
		return
	}

	// Validar pedido
	pedidoExiste := false

	for _, pedido := range storage.Pedidos {
		if pedido.ID == reclamo.PedidoID {
			pedidoExiste = true
			break
		}
	}

	if !pedidoExiste {
		http.Error(w, "El pedido asociado no existe", http.StatusBadRequest)
		return
	}

	reclamo.ID = storage.ReclamoID
	storage.ReclamoID++

	storage.Reclamos = append(storage.Reclamos, reclamo)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(reclamo)
}

func ObtenerReclamos(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(storage.Reclamos)
}


func ObtenerReclamoPorID(w http.ResponseWriter, r *http.Request) {

	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	for _, reclamo := range storage.Reclamos {

		if reclamo.ID == id {

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(reclamo)

			return
		}
	}

	http.Error(w, "Reclamo no encontrado", http.StatusNotFound)
}


func ActualizarReclamo(w http.ResponseWriter, r *http.Request) {

	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var reclamoActualizado models.Reclamo

	err = json.NewDecoder(r.Body).Decode(&reclamoActualizado)

	if err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	if reclamoActualizado.ClienteID <= 0 {
		http.Error(w, "ClienteID obligatorio", http.StatusBadRequest)
		return
	}

	if reclamoActualizado.PedidoID <= 0 {
		http.Error(w, "PedidoID obligatorio", http.StatusBadRequest)
		return
	}

	if reclamoActualizado.Descripcion == "" {
		http.Error(w, "Descripción obligatoria", http.StatusBadRequest)
		return
	}

	if reclamoActualizado.Estado == "" {
		http.Error(w, "Estado obligatorio", http.StatusBadRequest)
		return
	}

	for i, reclamo := range storage.Reclamos {

		if reclamo.ID == id {

			reclamoActualizado.ID = id

			storage.Reclamos[i] = reclamoActualizado

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(reclamoActualizado)

			return
		}
	}

	http.Error(w, "Reclamo no encontrado", http.StatusNotFound)
}


func EliminarReclamo(w http.ResponseWriter, r *http.Request) {

	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	for i, reclamo := range storage.Reclamos {

		if reclamo.ID == id {

			storage.Reclamos = append(
				storage.Reclamos[:i],
				storage.Reclamos[i+1:]...,
			)

			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Reclamo eliminado"))

			return
		}
	}

	http.Error(w, "Reclamo no encontrado", http.StatusNotFound)
}
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

var seguimientoPedidoService = service.NewSeguimientoPedidoService()

func CrearSeguimientoPedido(w http.ResponseWriter, r *http.Request) {

	var seguimiento models.SeguimientoPedido

	err := json.NewDecoder(r.Body).Decode(&seguimiento)

	if err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	err = seguimientoPedidoService.ValidarSeguimientoPedido(
		&seguimiento,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validar que exista el pedido
	pedidoExiste := false

	for _, pedido := range storage.Pedidos {
		if pedido.ID == seguimiento.PedidoID {
			pedidoExiste = true
			break
		}
	}

	if !pedidoExiste {
		http.Error(w, "El pedido asociado no existe", http.StatusBadRequest)
		return
	}

	seguimiento.ID = storage.SeguimientoPedidoID
	storage.SeguimientoPedidoID++

	storage.SeguimientosPedido = append(
		storage.SeguimientosPedido,
		seguimiento,
	)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(seguimiento)
}


func ObtenerSeguimientosPedido(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(storage.SeguimientosPedido)
}


func ObtenerSeguimientoPedidoPorID(w http.ResponseWriter, r *http.Request) {

	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	for _, seguimiento := range storage.SeguimientosPedido {

		if seguimiento.ID == id {

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(seguimiento)

			return
		}
	}

	http.Error(w, "Seguimiento no encontrado", http.StatusNotFound)
}


func ActualizarSeguimientoPedido(w http.ResponseWriter, r *http.Request) {

	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var seguimientoActualizado models.SeguimientoPedido

	err = json.NewDecoder(r.Body).Decode(&seguimientoActualizado)

	if err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	if seguimientoActualizado.PedidoID <= 0 {
		http.Error(w, "PedidoID obligatorio", http.StatusBadRequest)
		return
	}

	if seguimientoActualizado.Estado == "" {
		http.Error(w, "Estado obligatorio", http.StatusBadRequest)
		return
	}

	if seguimientoActualizado.FechaEstado == "" {
		http.Error(w, "FechaEstado obligatoria", http.StatusBadRequest)
		return
	}

	pedidoExiste := false

	for _, pedido := range storage.Pedidos {
		if pedido.ID == seguimientoActualizado.PedidoID {
			pedidoExiste = true
			break
		}
	}

	if !pedidoExiste {
		http.Error(w, "El pedido asociado no existe", http.StatusBadRequest)
		return
	}

	for i, seguimiento := range storage.SeguimientosPedido {

		if seguimiento.ID == id {

			seguimientoActualizado.ID = id

			storage.SeguimientosPedido[i] = seguimientoActualizado

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(seguimientoActualizado)

			return
		}
	}

	http.Error(w, "Seguimiento no encontrado", http.StatusNotFound)
}


func EliminarSeguimientoPedido(w http.ResponseWriter, r *http.Request) {

	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	for i, seguimiento := range storage.SeguimientosPedido {

		if seguimiento.ID == id {

			storage.SeguimientosPedido = append(
				storage.SeguimientosPedido[:i],
				storage.SeguimientosPedido[i+1:]...,
			)

			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Seguimiento eliminado"))

			return
		}
	}

	http.Error(w, "Seguimiento no encontrado", http.StatusNotFound)
}

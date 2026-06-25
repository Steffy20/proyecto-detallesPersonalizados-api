package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"proyecto-detallesPersonalizados-api/internal/models"
	"proyecto-detallesPersonalizados-api/internal/storage"
)

func CrearSeguimientoPedido(w http.ResponseWriter, r *http.Request) {
	var seguimiento models.SeguimientoPedido
	err := json.NewDecoder(r.Body).Decode(&seguimiento)
	if err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}
	if seguimiento.PedidoID <= 0 {
		http.Error(w, "PedidoID obligatorio", http.StatusBadRequest)
		return
	}
	if seguimiento.Estado == "" {
		http.Error(w, "Estado obligatorio", http.StatusBadRequest)
		return
	}
	if seguimiento.FechaEstado == "" {
		http.Error(w, "FechaEstado obligatoria", http.StatusBadRequest)
		return
	}
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
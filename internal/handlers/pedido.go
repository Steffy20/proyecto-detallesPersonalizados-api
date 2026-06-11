package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"proyecto-detallesPersonalizados-api/internal/models"
	"proyecto-detallesPersonalizados-api/internal/storage"
)

//commit: implementar endpoint crear pedido
func CrearPedido(w http.ResponseWriter, r *http.Request) {

	var pedido models.Pedido

	err := json.NewDecoder(r.Body).Decode(&pedido)

	if err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

//agregar validaciones para pedidos
	if pedido.Cliente == "" || pedido.Producto == "" {
	http.Error(w, "Cliente y producto son obligatorios", http.StatusBadRequest)
	return
}

//commit: almacenar pedidos en memoria
pedido.ID = storage.PedidoID
storage.PedidoID++

storage.Pedidos = append(storage.Pedidos, pedido)


	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(pedido)
}

// GET todos los pedidos
//commit: implementar listado de pedidos
func ObtenerPedidos(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(storage.Pedidos)
}

//GET POR ID
//commit: implementar busqueda de pedido por id

func ObtenerPedidoPorID(w http.ResponseWriter, r *http.Request) {

	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	for _, pedido := range storage.Pedidos {

		if pedido.ID == id {

			w.Header().Set("Content-Type", "application/json")

			json.NewEncoder(w).Encode(pedido)

			return
		}
	}

	http.Error(w, "Pedido no encontrado", http.StatusNotFound)
}

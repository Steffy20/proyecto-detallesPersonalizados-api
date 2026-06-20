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
	if pedido.Mensaje == "" {
	http.Error(w, "El mensaje personalizado es obligatorio", http.StatusBadRequest)
	return
}

if pedido.Estado == "" {
	pedido.Estado = "Pendiente"
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

//UPDATE
//commit:implementar actualizacion de pedidos


func ActualizarPedido(w http.ResponseWriter, r *http.Request) {

	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var pedidoActualizado models.Pedido

	err = json.NewDecoder(r.Body).Decode(&pedidoActualizado)

	if err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	for i, pedido := range storage.Pedidos {

		if pedido.ID == id {

			pedidoActualizado.ID = id

			storage.Pedidos[i] = pedidoActualizado

			w.Header().Set("Content-Type", "application/json")

			json.NewEncoder(w).Encode(pedidoActualizado)

			return
		}
	}

	http.Error(w, "Pedido no encontrado", http.StatusNotFound)
}

//DELETE
//commit: implementar eliminacion de pedidos


func EliminarPedido(w http.ResponseWriter, r *http.Request) {

	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	for i, pedido := range storage.Pedidos {

		if pedido.ID == id {

			storage.Pedidos = append(storage.Pedidos[:i], storage.Pedidos[i+1:]...)

			w.WriteHeader(http.StatusOK)

			w.Write([]byte("Pedido eliminado"))

			return
		}
	}

	http.Error(w, "Pedido no encontrado", http.StatusNotFound)
}

package handlers

import (
	"encoding/json"
	"net/http"

	"proyecto-detallesPersonalizados-api/internal/models"
	"proyecto-detallesPersonalizados-api/internal/storage"
)

func CrearPedido(w http.ResponseWriter, r *http.Request) {

	var pedido models.Pedido

	err := json.NewDecoder(r.Body).Decode(&pedido)

	if err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	if pedido.Cliente == "" || pedido.Producto == "" {
	http.Error(w, "Cliente y producto son obligatorios", http.StatusBadRequest)
	return
}
pedido.ID = storage.PedidoID
storage.PedidoID++

storage.Pedidos = append(storage.Pedidos, pedido)


	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(pedido)
}

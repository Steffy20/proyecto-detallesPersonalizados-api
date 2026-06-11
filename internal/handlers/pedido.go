package handlers

import (
	"encoding/json"
	"net/http"

	"proyecto-detallesPersonalizados-api/internal/models"
)

func CrearPedido(w http.ResponseWriter, r *http.Request) {

	var pedido models.Pedido

	err := json.NewDecoder(r.Body).Decode(&pedido)

	if pedido.Cliente == "" || pedido.Producto == "" {
	http.Error(w, "Cliente y producto son obligatorios", http.StatusBadRequest)
	return
}

	if err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(pedido)
}

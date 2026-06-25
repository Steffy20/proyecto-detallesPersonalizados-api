package models

type Reclamo struct {
	ID          int    `json:"id"`
	ClienteID   int    `json:"cliente_id"`
	PedidoID    int    `json:"pedido_id"`
	Descripcion string `json:"descripcion"`
	Estado      string `json:"estado"`
}

package models

type SeguimientoPedido struct {
	ID          int    `json:"id"`
	PedidoID    int    `json:"pedido_id"`
	Estado      string `json:"estado"`
	FechaEstado string `json:"fecha_estado"`
}
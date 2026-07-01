package models

type SeguimientoPedido struct {
	ID           int    `gorm:"primaryKey" json:"id"`
	PedidoID     int    `json:"pedido_id"`
	Estado       string `json:"estado"`
	FechaEstado  string `json:"fecha_estado"`
}

package models

type Reclamo struct {
	ID          int    `gorm:"primaryKey" json:"id"`
	ClienteID   int    `json:"cliente_id"`
	PedidoID    int    `json:"pedido_id"`
	Descripcion string `json:"descripcion"`
	Estado      string `json:"estado"`
}

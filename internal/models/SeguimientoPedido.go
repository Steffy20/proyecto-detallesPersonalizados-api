package models

type SeguimientoPedido struct {
	ID          int     `gorm:"primaryKey" json:"id"`
	PedidoID    int     `json:"pedido_id"`
	Pedido      *Pedido `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"pedido,omitempty"`
	Estado      string  `json:"estado"`
	FechaEstado string  `json:"fecha_estado"`
}

package models

type Reclamo struct {
	ID          int      `gorm:"primaryKey" json:"id"`
	ClienteID   int      `json:"cliente_id"`
	Cliente     *Cliente `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"cliente,omitempty"`
	PedidoID    int      `json:"pedido_id"`
	Pedido      *Pedido  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"pedido,omitempty"`
	Descripcion string   `json:"descripcion"`
	Estado      string   `json:"estado"`
}

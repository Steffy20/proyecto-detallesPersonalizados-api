package models

type Personalizacion struct {
	ID            int     `gorm:"primaryKey" json:"id"`
	PedidoID      int     `json:"pedido_id"`
	Pedido        *Pedido `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"pedido,omitempty"`
	Mensaje       string  `json:"mensaje"`
	Color         string  `json:"color"`
	Observaciones string  `json:"observaciones"`
}

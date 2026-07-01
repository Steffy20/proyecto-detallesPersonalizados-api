package models

type Personalizacion struct {
	ID            int    `gorm:"primaryKey" json:"id"`
	PedidoID      int    `json:"pedido_id"`
	Mensaje       string `json:"mensaje"`
	Color         string `json:"color"`
	Observaciones string `json:"observaciones"`
}
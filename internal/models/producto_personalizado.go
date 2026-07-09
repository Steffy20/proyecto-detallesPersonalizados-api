package models

type ProductoPersonalizado struct {
	ID        int     `gorm:"primaryKey" json:"id"`
	PedidoID  int     `json:"pedido_id"`
	Pedido    *Pedido `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"pedido,omitempty"`
	Nombre    string  `json:"nombre"`
	Categoria string  `json:"categoria"`
	Cantidad  int     `json:"cantidad"`
	Precio    float64 `json:"precio"`
}

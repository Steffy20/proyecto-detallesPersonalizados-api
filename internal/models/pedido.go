package models

// commit: agregar atributos al modelo pedido
type Pedido struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	Cliente  string `json:"cliente"`
	Producto string `json:"producto"`
	Mensaje  string `json:"mensaje"`
	Estado   string `json:"estado"`
}



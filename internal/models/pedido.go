package models

// commit: agregar atributos al modelo pedido
type Pedido struct {
	ID                      int                     `gorm:"primaryKey" json:"id"`
	Cliente                 string                  `json:"cliente"`
	Producto                string                  `json:"producto"`
	Mensaje                 string                  `json:"mensaje"`
	Estado                  string                  `json:"estado"`
	Personalizaciones       []Personalizacion       `gorm:"foreignKey:PedidoID" json:"personalizaciones,omitempty"`
	ProductosPersonalizados []ProductoPersonalizado `gorm:"foreignKey:PedidoID" json:"productos_personalizados,omitempty"`
	Reclamos                []Reclamo               `gorm:"foreignKey:PedidoID" json:"reclamos,omitempty"`
	SeguimientosPedido      []SeguimientoPedido     `gorm:"foreignKey:PedidoID" json:"seguimientos_pedido,omitempty"`
}

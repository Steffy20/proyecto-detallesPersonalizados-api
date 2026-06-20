package models

type ProductoPersonalizado struct {
	ID        int     `json:"id"`
	PedidoID  int     `json:"pedido_id"`
	Nombre    string  `json:"nombre"`
	Categoria string  `json:"categoria"`
	Cantidad  int     `json:"cantidad"`
	Precio    float64 `json:"precio"`
}

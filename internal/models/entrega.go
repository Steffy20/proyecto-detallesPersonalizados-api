package models

type Entrega struct {
	ID           int    `json:"id"`
	Direccion    string `json:"direccion"`
	Estado       string `json:"estado"`
	FechaEntrega string `json:"fecha_entrega"`
}

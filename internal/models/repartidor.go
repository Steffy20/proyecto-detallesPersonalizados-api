package models

type Repartidor struct {
	ID         int    `json:"id"`
	Nombre     string `json:"nombre"`
	Telefono   string `json:"telefono"`
	Disponible bool   `json:"disponible"`
	Zona       string `json:"zona"`
}
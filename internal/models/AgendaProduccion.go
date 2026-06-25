package models

type AgendaProduccion struct {
	ID          int    `json:"id"`
	Fecha       string `json:"fecha"`
	Responsable string `json:"responsable"`
	Estado      string `json:"estado"`
}
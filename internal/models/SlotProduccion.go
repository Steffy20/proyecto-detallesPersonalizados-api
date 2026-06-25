package models

type SlotProduccion struct {
	ID                int `json:"id"`
	AgendaID          int `json:"agenda_id"`
	CapacidadMaxima   int `json:"capacidad_maxima"`
	PedidosAsignados  int `json:"pedidos_asignados"`
}

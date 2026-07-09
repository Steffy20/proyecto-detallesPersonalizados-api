package models

type SlotProduccion struct {
	ID               int               `gorm:"primaryKey" json:"id"`
	AgendaID         int               `json:"agenda_id"`
	Agenda           *AgendaProduccion `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"agenda,omitempty"`
	CapacidadMaxima  int               `json:"capacidad_maxima"`
	PedidosAsignados int               `json:"pedidos_asignados"`
}

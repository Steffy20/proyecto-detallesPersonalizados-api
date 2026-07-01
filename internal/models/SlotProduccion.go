package models

type SlotProduccion struct {
<<<<<<< HEAD
	ID                int `gorm:"primaryKey" json:"id"`
	AgendaID          int `json:"agenda_id"`
	CapacidadMaxima   int `json:"capacidad_maxima"`
	PedidosAsignados  int `json:"pedidos_asignados"`
}
=======
	ID               int `json:"id"`
	AgendaID         int `json:"agenda_id"`
	CapacidadMaxima  int `json:"capacidad_maxima"`
	PedidosAsignados int `json:"pedidos_asignados"`
}
>>>>>>> 53c8cd66769e645b9a25e9c8a1464333d6df57d0

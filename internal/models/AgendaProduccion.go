package models

type AgendaProduccion struct {
<<<<<<< HEAD
	ID           int    `gorm:"primaryKey" json:"id"`
	Fecha        string `json:"fecha"`
	Responsable  string `json:"responsable"`
	Estado       string `json:"estado"`
}
=======
	ID          int    `json:"id"`
	Fecha       string `json:"fecha"`
	Responsable string `json:"responsable"`
	Estado      string `json:"estado"`
}
>>>>>>> 53c8cd66769e645b9a25e9c8a1464333d6df57d0

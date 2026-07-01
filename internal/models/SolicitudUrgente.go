package models

type SolicitudUrgente struct {
<<<<<<< HEAD
	ID             int    `gorm:"primaryKey" json:"id"`
=======
	ID             int    `json:"id"`
>>>>>>> 53c8cd66769e645b9a25e9c8a1464333d6df57d0
	Cliente        string `json:"cliente"`
	Descripcion    string `json:"descripcion"`
	FechaRequerida string `json:"fecha_requerida"`
	Estado         string `json:"estado"`
}
<<<<<<< HEAD
	
=======
>>>>>>> 53c8cd66769e645b9a25e9c8a1464333d6df57d0

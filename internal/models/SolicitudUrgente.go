package models

type SolicitudUrgente struct {
	ID             int    `gorm:"primaryKey" json:"id"`
	Cliente        string `json:"cliente"`
	Descripcion    string `json:"descripcion"`
	FechaRequerida string `json:"fecha_requerida"`
	Estado         string `json:"estado"`
}
	
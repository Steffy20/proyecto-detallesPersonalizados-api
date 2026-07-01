package models

import "time"

type Usuario struct {
	ID           int       `json:"id" gorm:"primaryKey"`
	Nombre       string    `json:"nombre,omitempty"`
	Email        string    `json:"email,omitempty" gorm:"unique;not null"`
	PasswordHash string    `json:"-" gorm:"column:password_hash"`
	Rol          string    `json:"rol,omitempty"`
	CreadoEn     time.Time `json:"creado_en,omitempty" gorm:"column:creado_en"`
}

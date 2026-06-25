package models

type Cliente struct {
    ID       int    `json:"id"`
    Nombre   string `json:"nombre"`
    Telefono string `json:"telefono"`
    Correo   string `json:"correo"`
}
package service

import (
	"errors"

	"proyecto-detallesPersonalizados-api/internal/models"
)

type ClienteService struct{}

func NewClienteService() *ClienteService {
	return &ClienteService{}
}

func (s *ClienteService) ValidarCliente(
	cliente *models.Cliente,
) error {

	if cliente.Nombre == "" {
		return errors.New("Nombre obligatorio")
	}

	if cliente.Telefono == "" {
		return errors.New("Teléfono obligatorio")
	}

	return nil
}
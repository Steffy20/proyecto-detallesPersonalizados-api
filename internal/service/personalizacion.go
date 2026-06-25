package service

import (
	"errors"

	"proyecto-detallesPersonalizados-api/internal/models"
)

type PersonalizacionService struct{}

func NewPersonalizacionService() *PersonalizacionService {
	return &PersonalizacionService{}
}

func (s *PersonalizacionService) ValidarPersonalizacion(
	personalizacion *models.Personalizacion,
) error {

	if personalizacion.PedidoID == 0 {
		return errors.New("PedidoID es obligatorio")
	}

	if personalizacion.Mensaje == "" {
		return errors.New("El mensaje es obligatorio")
	}

	if personalizacion.Color == "" {
		return errors.New("El color es obligatorio")
	}

	return nil
}

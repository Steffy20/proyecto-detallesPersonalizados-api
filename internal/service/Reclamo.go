package service

import (
	"errors"

	"proyecto-detallesPersonalizados-api/internal/models"
)

type ReclamoService struct{}

func NewReclamoService() *ReclamoService {
	return &ReclamoService{}
}

func (s *ReclamoService) ValidarReclamo(
	reclamo *models.Reclamo,
) error {

	if reclamo.ClienteID <= 0 {
		return errors.New("ClienteID obligatorio")
	}

	if reclamo.PedidoID <= 0 {
		return errors.New("PedidoID obligatorio")
	}

	if reclamo.Descripcion == "" {
		return errors.New("Descripción obligatoria")
	}

	if reclamo.Estado == "" {
		return errors.New("Estado obligatorio")
	}

	return nil
}

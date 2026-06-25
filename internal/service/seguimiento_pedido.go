package service

import (
	"errors"

	"proyecto-detallesPersonalizados-api/internal/models"
)

type SeguimientoPedidoService struct{}

func NewSeguimientoPedidoService() *SeguimientoPedidoService {
	return &SeguimientoPedidoService{}
}

func (s *SeguimientoPedidoService) ValidarSeguimientoPedido(
	seguimiento *models.SeguimientoPedido,
) error {

	if seguimiento.PedidoID <= 0 {
		return errors.New("PedidoID obligatorio")
	}

	if seguimiento.Estado == "" {
		return errors.New("Estado obligatorio")
	}

	if seguimiento.FechaEstado == "" {
		return errors.New("FechaEstado obligatoria")
	}

	return nil
}

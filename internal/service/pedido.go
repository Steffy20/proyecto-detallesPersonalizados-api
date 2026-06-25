package service

import (
	"errors"

	"proyecto-detallesPersonalizados-api/internal/models"
)

type PedidoService struct{}

func NewPedidoService() *PedidoService {
	return &PedidoService{}
}

func (s *PedidoService) ValidarPedido(pedido *models.Pedido) error {

	if pedido.Mensaje == "" {
		return errors.New("el mensaje personalizado es obligatorio")
	}

	// Estado por defecto
	if pedido.Estado == "" {
		pedido.Estado = "Pendiente"
	}

	return nil
}

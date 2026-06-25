package service

import (
	"errors"

	"proyecto-detallesPersonalizados-api/internal/models"
)

type SlotProduccionService struct{}

func NewSlotProduccionService() *SlotProduccionService {
	return &SlotProduccionService{}
}

func (s *SlotProduccionService) ValidarSlotProduccion(
	slot *models.SlotProduccion,
) error {

	if slot.AgendaID <= 0 {
		return errors.New("AgendaID obligatorio")
	}

	if slot.CapacidadMaxima <= 0 {
		return errors.New("Capacidad máxima inválida")
	}

	if slot.PedidosAsignados < 0 {
		return errors.New("Pedidos asignados inválidos")
	}

	return nil
}
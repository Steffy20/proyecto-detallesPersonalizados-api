package service

import (
	"errors"

	"proyecto-detallesPersonalizados-api/internal/models"
)

type AgendaProduccionService struct{}

func NewAgendaProduccionService() *AgendaProduccionService {
	return &AgendaProduccionService{}
}

func (s *AgendaProduccionService) ValidarAgendaProduccion(
	agenda *models.AgendaProduccion,
) error {

	if agenda.Fecha == "" {
		return errors.New("La fecha es obligatoria")
	}

	if agenda.Responsable == "" {
		return errors.New("El responsable es obligatorio")
	}

	if agenda.Estado == "" {
		return errors.New("El estado es obligatorio")
	}

	return nil
}

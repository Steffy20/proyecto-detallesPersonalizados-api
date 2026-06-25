package service

import (
	"errors"

	"proyecto-detallesPersonalizados-api/internal/models"
)

type SolicitudUrgenteService struct{}

func NewSolicitudUrgenteService() *SolicitudUrgenteService {
	return &SolicitudUrgenteService{}
}

func (s *SolicitudUrgenteService) ValidarSolicitudUrgente(
	solicitud *models.SolicitudUrgente,
) error {

	if solicitud.Cliente == "" {
		return errors.New("Cliente obligatorio")
	}

	if solicitud.Descripcion == "" {
		return errors.New("Descripción obligatoria")
	}

	if solicitud.FechaRequerida == "" {
		return errors.New("Fecha requerida obligatoria")
	}

	// Estado por defecto
	if solicitud.Estado == "" {
		solicitud.Estado = "Pendiente"
	}

	return nil
}
package service

import (
	"errors"

	"proyecto-detallesPersonalizados-api/internal/models"
	
)

type SolicitudUrgenteRepository interface {
	ListarSolicitudesUrgentes() []models.SolicitudUrgente
	BuscarSolicitudUrgentePorID(id int) (models.SolicitudUrgente, bool)
	CrearSolicitudUrgente(models.SolicitudUrgente) models.SolicitudUrgente
	ActualizarSolicitudUrgente(id int, datos models.SolicitudUrgente) (models.SolicitudUrgente, bool)
	BorrarSolicitudUrgente(id int) bool
}

type SolicitudUrgenteService struct {
	Almacen SolicitudUrgenteRepository
}

func NewSolicitudUrgenteService(almacen SolicitudUrgenteRepository) *SolicitudUrgenteService {
	return &SolicitudUrgenteService{
		Almacen: almacen,
	}
}

// ================= VALIDACIONES =================

func (s *SolicitudUrgenteService) ValidarSolicitudUrgente(sol *models.SolicitudUrgente) error {

	if sol.Cliente == "" {
		return errors.New("el cliente es obligatorio")
	}

	if sol.Descripcion == "" {
		return errors.New("la descripción es obligatoria")
	}

	if sol.FechaRequerida == "" {
		return errors.New("la fecha requerida es obligatoria")
	}

	if sol.Estado == "" {
		sol.Estado = "Pendiente"
	}

	return nil
}

// ================= CRUD =================

func (s *SolicitudUrgenteService) Listar() []models.SolicitudUrgente {
	return s.Almacen.ListarSolicitudesUrgentes()
}

func (s *SolicitudUrgenteService) Obtener(id int) (models.SolicitudUrgente, error) {

	solicitud, ok := s.Almacen.BuscarSolicitudUrgentePorID(id)

	if !ok {
		return models.SolicitudUrgente{}, errors.New("solicitud urgente no encontrada")
	}

	return solicitud, nil
}

func (s *SolicitudUrgenteService) Crear(sol models.SolicitudUrgente) (models.SolicitudUrgente, error) {

	if err := s.ValidarSolicitudUrgente(&sol); err != nil {
		return models.SolicitudUrgente{}, err
	}

	return s.Almacen.CrearSolicitudUrgente(sol), nil
}

func (s *SolicitudUrgenteService) Actualizar(id int, datos models.SolicitudUrgente) (models.SolicitudUrgente, error) {

	if err := s.ValidarSolicitudUrgente(&datos); err != nil {
		return models.SolicitudUrgente{}, err
	}

	solicitud, ok := s.Almacen.ActualizarSolicitudUrgente(id, datos)

	if !ok {
		return models.SolicitudUrgente{}, errors.New("solicitud urgente no encontrada")
	}

	return solicitud, nil
}

func (s *SolicitudUrgenteService) Borrar(id int) error {

	if !s.Almacen.BorrarSolicitudUrgente(id) {
		return errors.New("solicitud urgente no encontrada")
	}

	return nil
}
package service

import (
	"errors"

	"proyecto-detallesPersonalizados-api/internal/models"
	"proyecto-detallesPersonalizados-api/internal/storage"
)

type AgendaProduccionService struct {
	Almacen storage.Almacen
}


func NewAgendaProduccionService(almacen storage.Almacen) *AgendaProduccionService {
	return &AgendaProduccionService{
		Almacen: almacen,
	}
}

// ================= VALIDACIONES =================

func (s *AgendaProduccionService) ValidarAgendaProduccion(a *models.AgendaProduccion) error {

	if a.Fecha == "" {
		return errors.New("la fecha es obligatoria")
	}

	if a.Responsable == "" {
		return errors.New("el responsable es obligatorio")
	}

	if a.Estado == "" {
		a.Estado = "Pendiente"
	}

	return nil
}

// ================= CRUD =================

func (s *AgendaProduccionService) Listar() []models.AgendaProduccion {
	return s.Almacen.ListarAgendasProduccion()
}

func (s *AgendaProduccionService) Obtener(id int) (models.AgendaProduccion, error) {

	agenda, ok := s.Almacen.BuscarAgendaProduccionPorID(id)

	if !ok {
		return models.AgendaProduccion{}, errors.New("agenda de producción no encontrada")
	}

	return agenda, nil
}

func (s *AgendaProduccionService) Crear(a models.AgendaProduccion) (models.AgendaProduccion, error) {

	if err := s.ValidarAgendaProduccion(&a); err != nil {
		return models.AgendaProduccion{}, err
	}

	return s.Almacen.CrearAgendaProduccion(a), nil
}

func (s *AgendaProduccionService) Actualizar(id int, datos models.AgendaProduccion) (models.AgendaProduccion, error) {

	if err := s.ValidarAgendaProduccion(&datos); err != nil {
		return models.AgendaProduccion{}, err
	}

	agenda, ok := s.Almacen.ActualizarAgendaProduccion(id, datos)

	if !ok {
		return models.AgendaProduccion{}, errors.New("agenda de producción no encontrada")
	}

	return agenda, nil
}

func (s *AgendaProduccionService) Borrar(id int) error {

	if !s.Almacen.BorrarAgendaProduccion(id) {
		return errors.New("agenda de producción no encontrada")
	}

	return nil
}

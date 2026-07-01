package service

import (
	"errors"

	"proyecto-detallesPersonalizados-api/internal/models"
	
)

type ReclamoRepository interface {
	ListarReclamos() []models.Reclamo
	BuscarReclamoPorID(id int) (models.Reclamo, bool)
	CrearReclamo(models.Reclamo) models.Reclamo
	ActualizarReclamo(id int, datos models.Reclamo) (models.Reclamo, bool)
	BorrarReclamo(id int) bool
}

type ReclamoService struct {
	Almacen ReclamoRepository
}

func NewReclamoService(almacen ReclamoRepository) *ReclamoService {
	return &ReclamoService{
		Almacen: almacen,
	}
}

// ================= VALIDACIONES =================

func (s *ReclamoService) ValidarReclamo(r *models.Reclamo) error {

	if r.ClienteID <= 0 {
		return errors.New("el cliente es obligatorio")
	}

	if r.PedidoID <= 0 {
		return errors.New("el pedido es obligatorio")
	}

	if r.Descripcion == "" {
		return errors.New("la descripción es obligatoria")
	}

	if r.Estado == "" {
		r.Estado = "Pendiente"
	}

	return nil
}

// ================= CRUD =================

func (s *ReclamoService) Listar() []models.Reclamo {
	return s.Almacen.ListarReclamos()
}

func (s *ReclamoService) Obtener(id int) (models.Reclamo, error) {

	reclamo, ok := s.Almacen.BuscarReclamoPorID(id)

	if !ok {
		return models.Reclamo{}, errors.New("reclamo no encontrado")
	}

	return reclamo, nil
}

func (s *ReclamoService) Crear(r models.Reclamo) (models.Reclamo, error) {

	if err := s.ValidarReclamo(&r); err != nil {
		return models.Reclamo{}, err
	}

	return s.Almacen.CrearReclamo(r), nil
}

func (s *ReclamoService) Actualizar(id int, datos models.Reclamo) (models.Reclamo, error) {

	if err := s.ValidarReclamo(&datos); err != nil {
		return models.Reclamo{}, err
	}

	reclamo, ok := s.Almacen.ActualizarReclamo(id, datos)

	if !ok {
		return models.Reclamo{}, errors.New("reclamo no encontrado")
	}

	return reclamo, nil
}

func (s *ReclamoService) Borrar(id int) error {

	if !s.Almacen.BorrarReclamo(id) {
		return errors.New("reclamo no encontrado")
	}

	return nil
}

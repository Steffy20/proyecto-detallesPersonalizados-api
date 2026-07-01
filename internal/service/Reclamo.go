package service

import (
	"errors"

	"proyecto-detallesPersonalizados-api/internal/models"
	"proyecto-detallesPersonalizados-api/internal/storage"
)

type ReclamoService struct {
	Almacen storage.Almacen
}

func NewReclamoService(almacen storage.Almacen) *ReclamoService {
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

package service

import (
	"errors"

	"proyecto-detallesPersonalizados-api/internal/models"
	
)
type SeguimientoPedidoRepository interface {
	ListarSeguimientosPedido() []models.SeguimientoPedido
	BuscarSeguimientoPedidoPorID(id int) (models.SeguimientoPedido, bool)
	CrearSeguimientoPedido(models.SeguimientoPedido) models.SeguimientoPedido
	ActualizarSeguimientoPedido(id int, datos models.SeguimientoPedido) (models.SeguimientoPedido, bool)
	BorrarSeguimientoPedido(id int) bool
}

type SeguimientoPedidoService struct {
	Almacen SeguimientoPedidoRepository
}

func NewSeguimientoPedidoService(almacen SeguimientoPedidoRepository) *SeguimientoPedidoService {
	return &SeguimientoPedidoService{
		Almacen: almacen,
	}
}

// ================= VALIDACIONES =================

func (s *SeguimientoPedidoService) ValidarSeguimiento(seg *models.SeguimientoPedido) error {

	if seg.PedidoID <= 0 {
		return errors.New("el pedido es obligatorio")
	}

	if seg.Estado == "" {
		return errors.New("el estado es obligatorio")
	}

	if seg.FechaEstado == "" {
		return errors.New("la fecha del estado es obligatoria")
	}

	return nil
}

// ================= CRUD =================

func (s *SeguimientoPedidoService) Listar() []models.SeguimientoPedido {
	return s.Almacen.ListarSeguimientosPedido()
}

func (s *SeguimientoPedidoService) Obtener(id int) (models.SeguimientoPedido, error) {

	seguimiento, ok := s.Almacen.BuscarSeguimientoPedidoPorID(id)

	if !ok {
		return models.SeguimientoPedido{}, errors.New("seguimiento no encontrado")
	}

	return seguimiento, nil
}

func (s *SeguimientoPedidoService) Crear(seg models.SeguimientoPedido) (models.SeguimientoPedido, error) {

	if err := s.ValidarSeguimiento(&seg); err != nil {
		return models.SeguimientoPedido{}, err
	}

	return s.Almacen.CrearSeguimientoPedido(seg), nil
}

func (s *SeguimientoPedidoService) Actualizar(id int, datos models.SeguimientoPedido) (models.SeguimientoPedido, error) {

	if err := s.ValidarSeguimiento(&datos); err != nil {
		return models.SeguimientoPedido{}, err
	}

	seguimiento, ok := s.Almacen.ActualizarSeguimientoPedido(id, datos)

	if !ok {
		return models.SeguimientoPedido{}, errors.New("seguimiento no encontrado")
	}

	return seguimiento, nil
}

func (s *SeguimientoPedidoService) Borrar(id int) error {

	if !s.Almacen.BorrarSeguimientoPedido(id) {
		return errors.New("seguimiento no encontrado")
	}

	return nil
}

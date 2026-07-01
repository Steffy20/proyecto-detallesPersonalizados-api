package service

import (
	"errors"

	"proyecto-detallesPersonalizados-api/internal/models"
	
)

//INTERFACE DE PEDIDO PARA LOS MOCKS
type PedidoRepository interface {
	ListarPedidos() []models.Pedido
	BuscarPedidoPorID(id int) (models.Pedido, bool)
	CrearPedido(models.Pedido) models.Pedido
	ActualizarPedido(id int, datos models.Pedido) (models.Pedido, bool)
	BorrarPedido(id int) bool
}


type PedidoService struct {
	Almacen PedidoRepository
}

func NewPedidoService(Almacen PedidoRepository) *PedidoService {
	return &PedidoService{
		Almacen: Almacen,
	}
}


func (s *PedidoService) ValidarPedido(pedido *models.Pedido) error {

	if pedido.Mensaje == "" {
		return errors.New("el mensaje personalizado es obligatorio")
	}

	if pedido.Estado == "" {
		pedido.Estado = "Pendiente"
	}


	return nil
}

// ================= CRUD =================

func (s *PedidoService) Listar() []models.Pedido {
	return s.Almacen.ListarPedidos()
}

func (s *PedidoService) Obtener(id int) (models.Pedido, error) {

	pedido, ok := s.Almacen.BuscarPedidoPorID(id)

	if !ok {
		return models.Pedido{}, errors.New("pedido no encontrado")
	}

	return pedido, nil
}

func (s *PedidoService) Crear(p models.Pedido) (models.Pedido, error) {

	if err := s.ValidarPedido(&p); err != nil {
		return models.Pedido{}, err
	}

	return s.Almacen.CrearPedido(p), nil
}

func (s *PedidoService) Actualizar(id int, datos models.Pedido) (models.Pedido, error) {

	if err := s.ValidarPedido(&datos); err != nil {
		return models.Pedido{}, err
	}

	pedido, ok := s.Almacen.ActualizarPedido(id, datos)

	if !ok {
		return models.Pedido{}, errors.New("pedido no encontrado")
	}

	return pedido, nil
}

func (s *PedidoService) Borrar(id int) error {

	if !s.Almacen.BorrarPedido(id) {
		return errors.New("pedido no encontrado")
	}

	return nil
}


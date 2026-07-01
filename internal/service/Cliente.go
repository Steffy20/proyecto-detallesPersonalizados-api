package service

import (
	"errors"

	"proyecto-detallesPersonalizados-api/internal/models"
	"proyecto-detallesPersonalizados-api/internal/storage"
)

type ClienteService struct {
	Almacen storage.Almacen
}

func NewClienteService(almacen storage.Almacen) *ClienteService {
	return &ClienteService{
		Almacen: almacen,
	}
}

// ================= VALIDACIONES =================

func (s *ClienteService) ValidarCliente(c *models.Cliente) error {

	if c.Nombre == "" {
		return errors.New("el nombre es obligatorio")
	}

	if c.Telefono == "" {
		return errors.New("el teléfono es obligatorio")
	}

	return nil
}

// ================= CRUD =================

func (s *ClienteService) Listar() []models.Cliente {
	return s.Almacen.ListarClientes()
}

func (s *ClienteService) Obtener(id int) (models.Cliente, error) {

	cliente, ok := s.Almacen.BuscarClientePorID(id)

	if !ok {
		return models.Cliente{}, errors.New("cliente no encontrado")
	}

	return cliente, nil
}

func (s *ClienteService) Crear(c models.Cliente) (models.Cliente, error) {

	if err := s.ValidarCliente(&c); err != nil {
		return models.Cliente{}, err
	}

	return s.Almacen.CrearCliente(c), nil
}

func (s *ClienteService) Actualizar(id int, datos models.Cliente) (models.Cliente, error) {

	if err := s.ValidarCliente(&datos); err != nil {
		return models.Cliente{}, err
	}

	cliente, ok := s.Almacen.ActualizarCliente(id, datos)

	if !ok {
		return models.Cliente{}, errors.New("cliente no encontrado")
	}

	return cliente, nil
}

func (s *ClienteService) Borrar(id int) error {

	if !s.Almacen.BorrarCliente(id) {
		return errors.New("cliente no encontrado")
	}

	return nil
}
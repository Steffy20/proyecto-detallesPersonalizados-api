 package service

import (
	"errors"

	"proyecto-detallesPersonalizados-api/internal/models"
	"proyecto-detallesPersonalizados-api/internal/storage"
)


type ProductoPersonalizadoService struct{
	Almacen storage.Almacen
}

func NewProductoPersonalizadoService(almacen storage.Almacen) *ProductoPersonalizadoService {
	return &ProductoPersonalizadoService{
		Almacen: almacen,
	}
}

// ================= VALIDACIONES =================

func (s *ProductoPersonalizadoService) ValidarProductoPersonalizado(p *models.ProductoPersonalizado) error {

	if p.PedidoID <= 0 {
		return errors.New("el pedido es obligatorio")
	}

	if p.Nombre == "" {
		return errors.New("el nombre del producto es obligatorio")
	}

	if p.Cantidad <= 0 {
		return errors.New("la cantidad debe ser mayor que cero")
	}

	if p.Precio <= 0 {
		return errors.New("el precio debe ser mayor que cero")
	}

	return nil
}

// ================= CRUD =================

func (s *ProductoPersonalizadoService) Listar() []models.ProductoPersonalizado {
	return s.Almacen.ListarProductosPersonalizados()
}

func (s *ProductoPersonalizadoService) Obtener(id int) (models.ProductoPersonalizado, error) {

	producto, ok := s.Almacen.BuscarProductoPersonalizadoPorID(id)

	if !ok {
		return models.ProductoPersonalizado{}, errors.New("producto personalizado no encontrado")
	}

	return producto, nil
}

func (s *ProductoPersonalizadoService) Crear(p models.ProductoPersonalizado) (models.ProductoPersonalizado, error) {

	if err := s.ValidarProductoPersonalizado(&p); err != nil {
		return models.ProductoPersonalizado{}, err
	}

	return s.Almacen.CrearProductoPersonalizado(p), nil
}

func (s *ProductoPersonalizadoService) Actualizar(id int, datos models.ProductoPersonalizado) (models.ProductoPersonalizado, error) {

	if err := s.ValidarProductoPersonalizado(&datos); err != nil {
		return models.ProductoPersonalizado{}, err
	}

	producto, ok := s.Almacen.ActualizarProductoPersonalizado(id, datos)

	if !ok {
		return models.ProductoPersonalizado{}, errors.New("producto personalizado no encontrado")
	}

	return producto, nil
}

func (s *ProductoPersonalizadoService) Borrar(id int) error {

	if !s.Almacen.BorrarProductoPersonalizado(id) {
		return errors.New("producto personalizado no encontrado")
	}

	return nil
}
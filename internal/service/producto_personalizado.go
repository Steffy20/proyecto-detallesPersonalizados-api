 package service

import (
	"errors"

	"proyecto-detallesPersonalizados-api/internal/models"
)

type ProductoPersonalizadoService struct{}

func NewProductoPersonalizadoService() *ProductoPersonalizadoService {
	return &ProductoPersonalizadoService{}
}

func (s *ProductoPersonalizadoService) ValidarProductoPersonalizado(
	p *models.ProductoPersonalizado,
) error {

	if p.PedidoID <= 0 {
		return errors.New("PedidoID obligatorio")
	}

	if p.Nombre == "" {
		return errors.New("Nombre obligatorio")
	}

	if p.Cantidad <= 0 {
		return errors.New("Cantidad inválida")
	}

	if p.Precio <= 0 {
		return errors.New("Precio inválido")
	}

	return nil
}
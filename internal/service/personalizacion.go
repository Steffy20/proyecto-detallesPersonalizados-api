package service

import (
	"errors"

	"proyecto-detallesPersonalizados-api/internal/models"
	"proyecto-detallesPersonalizados-api/internal/storage"
)



type PersonalizacionService struct {
	Almacen storage.Almacen
}


func NewPersonalizacionService(almacen storage.Almacen) *PersonalizacionService {
	return &PersonalizacionService{
		Almacen: almacen,
	}
}


func (s *PersonalizacionService) ValidarPersonalizacion(p *models.Personalizacion) error {

	if p.PedidoID <= 0 {
		return errors.New("el pedido es obligatorio")
	}

	if p.Mensaje == "" {
		return errors.New("el mensaje es obligatorio")
	}

	if p.Color == "" {
		return errors.New("el color es obligatorio")
	}

	return nil
}

// ================= CRUD =================

func (s *PersonalizacionService) Listar() []models.Personalizacion {
	return s.Almacen.ListarPersonalizaciones()
}

func (s *PersonalizacionService) Obtener(id int) (models.Personalizacion, error) {

	personalizacion, ok := s.Almacen.BuscarPersonalizacionPorID(id)

	if !ok {
		return models.Personalizacion{}, errors.New("personalización no encontrada")
	}

	return personalizacion, nil
}

func (s *PersonalizacionService) Crear(p models.Personalizacion) (models.Personalizacion, error) {

	if err := s.ValidarPersonalizacion(&p); err != nil {
		return models.Personalizacion{}, err
	}

	return s.Almacen.CrearPersonalizacion(p), nil
}

func (s *PersonalizacionService) Actualizar(id int, datos models.Personalizacion) (models.Personalizacion, error) {

	if err := s.ValidarPersonalizacion(&datos); err != nil {
		return models.Personalizacion{}, err
	}

	personalizacion, ok := s.Almacen.ActualizarPersonalizacion(id, datos)

	if !ok {
		return models.Personalizacion{}, errors.New("personalización no encontrada")
	}

	return personalizacion, nil
}

func (s *PersonalizacionService) Borrar(id int) error {

	if !s.Almacen.BorrarPersonalizacion(id) {
		return errors.New("personalización no encontrada")
	}

	return nil
}

package service

import (
	"errors"

	"proyecto-detallesPersonalizados-api/internal/models"
	"proyecto-detallesPersonalizados-api/internal/storage"
)

type SlotProduccionService struct{
	Almacen storage.Almacen
}
func NewSlotProduccionService(almacen storage.Almacen) *SlotProduccionService {
	return &SlotProduccionService{
		Almacen: almacen,
	}
}

// ================= VALIDACIONES =================

func (s *SlotProduccionService) ValidarSlotProduccion(slot *models.SlotProduccion) error {

	if slot.AgendaID <= 0 {
		return errors.New("la agenda es obligatoria")
	}

	if slot.CapacidadMaxima <= 0 {
		return errors.New("la capacidad máxima debe ser mayor que cero")
	}

	if slot.PedidosAsignados < 0 {
		return errors.New("los pedidos asignados no pueden ser negativos")
	}

	if slot.PedidosAsignados > slot.CapacidadMaxima {
		return errors.New("los pedidos asignados no pueden superar la capacidad máxima")
	}

	return nil
}

// ================= CRUD =================

func (s *SlotProduccionService) Listar() []models.SlotProduccion {
	return s.Almacen.ListarSlotsProduccion()
}

func (s *SlotProduccionService) Obtener(id int) (models.SlotProduccion, error) {

	slot, ok := s.Almacen.BuscarSlotProduccionPorID(id)

	if !ok {
		return models.SlotProduccion{}, errors.New("slot de producción no encontrado")
	}

	return slot, nil
}

func (s *SlotProduccionService) Crear(slot models.SlotProduccion) (models.SlotProduccion, error) {

	if err := s.ValidarSlotProduccion(&slot); err != nil {
		return models.SlotProduccion{}, err
	}

	return s.Almacen.CrearSlotProduccion(slot), nil
}

func (s *SlotProduccionService) Actualizar(id int, datos models.SlotProduccion) (models.SlotProduccion, error) {

	if err := s.ValidarSlotProduccion(&datos); err != nil {
		return models.SlotProduccion{}, err
	}

	slot, ok := s.Almacen.ActualizarSlotProduccion(id, datos)

	if !ok {
		return models.SlotProduccion{}, errors.New("slot de producción no encontrado")
	}

	return slot, nil
}

func (s *SlotProduccionService) Borrar(id int) error {

	if !s.Almacen.BorrarSlotProduccion(id) {
		return errors.New("slot de producción no encontrado")
	}

	return nil
}
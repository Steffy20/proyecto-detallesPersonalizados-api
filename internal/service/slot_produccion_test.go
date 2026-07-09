package service

import (
	"testing"

	"proyecto-detallesPersonalizados-api/internal/models"

	"github.com/stretchr/testify/assert"
)

type mockSlotProduccionRepository struct {
	crearLlamado bool
}

func (m *mockSlotProduccionRepository) ListarSlotsProduccion() []models.SlotProduccion {
	return nil
}

func (m *mockSlotProduccionRepository) BuscarSlotProduccionPorID(id int) (models.SlotProduccion, bool) {
	if id == 1 {
		return models.SlotProduccion{
			ID:               1,
			AgendaID:         1,
			CapacidadMaxima:  10,
			PedidosAsignados: 2,
		}, true
	}

	return models.SlotProduccion{}, false
}
func (m *mockSlotProduccionRepository) CrearSlotProduccion(s models.SlotProduccion) models.SlotProduccion {
	m.crearLlamado = true
	return s
}

func (m *mockSlotProduccionRepository) ActualizarSlotProduccion(id int, datos models.SlotProduccion) (models.SlotProduccion, bool) {
	return models.SlotProduccion{}, false
}

func (m *mockSlotProduccionRepository) BorrarSlotProduccion(id int) bool {
	return id == 1
}

// ===================== TESTS =====================

func TestCrearSlotProduccionCapacidadInvalida(t *testing.T) {
	mock := &mockSlotProduccionRepository{}
	service := NewSlotProduccionService(mock)

	slot := models.SlotProduccion{
		AgendaID:         1,
		CapacidadMaxima:  0,
		PedidosAsignados: 0,
	}

	_, err := service.Crear(slot)

	assert.Error(t, err)
	assert.False(t, mock.crearLlamado)
}

func TestCrearSlotProduccionValido(t *testing.T) {
	mock := &mockSlotProduccionRepository{}
	service := NewSlotProduccionService(mock)

	slot := models.SlotProduccion{
		AgendaID:         1,
		CapacidadMaxima:  10,
		PedidosAsignados: 2,
	}

	_, err := service.Crear(slot)

	assert.NoError(t, err)
	assert.True(t, mock.crearLlamado)
}

func TestObtenerSlotProduccionExistente(t *testing.T) {
	mock := &mockSlotProduccionRepository{}
	service := NewSlotProduccionService(mock)

	slot, err := service.Obtener(1)

	assert.NoError(t, err)
	assert.Equal(t, 1, slot.ID)
}

func TestObtenerSlotProduccionNoEncontrado(t *testing.T) {
	mock := &mockSlotProduccionRepository{}
	service := NewSlotProduccionService(mock)

	_, err := service.Obtener(99)

	assert.Error(t, err)
}

func TestBorrarSlotProduccionExistente(t *testing.T) {
	mock := &mockSlotProduccionRepository{}
	service := NewSlotProduccionService(mock)

	err := service.Borrar(1)

	assert.NoError(t, err)
}

func TestBorrarSlotProduccionNoEncontrado(t *testing.T) {
	mock := &mockSlotProduccionRepository{}
	service := NewSlotProduccionService(mock)

	err := service.Borrar(99)

	assert.Error(t, err)
}
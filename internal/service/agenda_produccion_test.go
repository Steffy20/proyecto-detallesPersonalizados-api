package service

import (
	"testing"

	"proyecto-detallesPersonalizados-api/internal/models"

	"github.com/stretchr/testify/assert"
)

type mockAgendaProduccionRepository struct {
	crearLlamado bool
}

func (m *mockAgendaProduccionRepository) ListarAgendasProduccion() []models.AgendaProduccion {
	return nil
}
func (m *mockAgendaProduccionRepository) BuscarAgendaProduccionPorID(id int) (models.AgendaProduccion, bool) {
	if id == 1 {
		return models.AgendaProduccion{
			ID:          1,
			Fecha:       "2026-07-09",
			Responsable: "Carlos",
			Estado:      "Pendiente",
		}, true
	}

	return models.AgendaProduccion{}, false
}

func (m *mockAgendaProduccionRepository) CrearAgendaProduccion(a models.AgendaProduccion) models.AgendaProduccion {
	m.crearLlamado = true
	return a
}

func (m *mockAgendaProduccionRepository) ActualizarAgendaProduccion(id int, datos models.AgendaProduccion) (models.AgendaProduccion, bool) {
	return models.AgendaProduccion{}, false
}

func (m *mockAgendaProduccionRepository) BorrarAgendaProduccion(id int) bool {
	return id == 1
}

// ===================== TESTS =====================

func TestCrearAgendaProduccionFechaVacia(t *testing.T) {
	mock := &mockAgendaProduccionRepository{}
	service := NewAgendaProduccionService(mock)

	agenda := models.AgendaProduccion{
		Fecha:       "",
		Responsable: "Carlos",
		Estado:      "Pendiente",
	}

	_, err := service.Crear(agenda)

	assert.Error(t, err)
	assert.False(t, mock.crearLlamado)
}

func TestCrearAgendaProduccionValida(t *testing.T) {
	mock := &mockAgendaProduccionRepository{}
	service := NewAgendaProduccionService(mock)

	agenda := models.AgendaProduccion{
		Fecha:       "2026-07-09",
		Responsable: "Carlos",
		Estado:      "Pendiente",
	}

	_, err := service.Crear(agenda)

	assert.NoError(t, err)
	assert.True(t, mock.crearLlamado)
}
func TestObtenerAgendaProduccionExistente(t *testing.T) {
	mock := &mockAgendaProduccionRepository{}
	service := NewAgendaProduccionService(mock)

	agenda, err := service.Obtener(1)

	assert.NoError(t, err)
	assert.Equal(t, 1, agenda.ID)
}

func TestObtenerAgendaProduccionNoEncontrada(t *testing.T) {
	mock := &mockAgendaProduccionRepository{}
	service := NewAgendaProduccionService(mock)

	_, err := service.Obtener(99)

	assert.Error(t, err)
}

func TestBorrarAgendaProduccionExistente(t *testing.T) {
	mock := &mockAgendaProduccionRepository{}
	service := NewAgendaProduccionService(mock)

	err := service.Borrar(1)

	assert.NoError(t, err)
}

func TestBorrarAgendaProduccionNoEncontrada(t *testing.T) {
	mock := &mockAgendaProduccionRepository{}
	service := NewAgendaProduccionService(mock)

	err := service.Borrar(99)

	assert.Error(t, err)
}
package service

import (
	"testing"

	"proyecto-detallesPersonalizados-api/internal/models"

	"github.com/stretchr/testify/assert"
)

type mockSolicitudUrgenteRepository struct {
	crearLlamado bool
}

func (m *mockSolicitudUrgenteRepository) ListarSolicitudesUrgentes() []models.SolicitudUrgente {
	return nil
}

func (m *mockSolicitudUrgenteRepository) BuscarSolicitudUrgentePorID(id int) (models.SolicitudUrgente, bool) {
	if id == 1 {
		return models.SolicitudUrgente{
			ID:             1,
			Cliente:        "Juan",
			Descripcion:    "Pedido urgente",
			FechaRequerida: "2026-07-09",
			Estado:         "Pendiente",
		}, true
	}

	return models.SolicitudUrgente{}, false
}

func (m *mockSolicitudUrgenteRepository) CrearSolicitudUrgente(s models.SolicitudUrgente) models.SolicitudUrgente {
	m.crearLlamado = true
	return s
}

func (m *mockSolicitudUrgenteRepository) ActualizarSolicitudUrgente(id int, datos models.SolicitudUrgente) (models.SolicitudUrgente, bool) {
	return models.SolicitudUrgente{}, false
}

func (m *mockSolicitudUrgenteRepository) BorrarSolicitudUrgente(id int) bool {
	return id == 1
}

// ===================== TESTS =====================

func TestCrearSolicitudUrgenteClienteVacio(t *testing.T) {
	mock := &mockSolicitudUrgenteRepository{}
	service := NewSolicitudUrgenteService(mock)

	solicitud := models.SolicitudUrgente{
		Cliente:        "",
		Descripcion:    "Pedido urgente para hoy",
		FechaRequerida: "2026-07-09",
	}

	_, err := service.Crear(solicitud)

	assert.Error(t, err)
	assert.False(t, mock.crearLlamado)
}

func TestCrearSolicitudUrgenteValida(t *testing.T) {
	mock := &mockSolicitudUrgenteRepository{}
	service := NewSolicitudUrgenteService(mock)

	solicitud := models.SolicitudUrgente{
		Cliente:        "Juan",
		Descripcion:    "Pedido urgente para hoy",
		FechaRequerida: "2026-07-09",
	}

	_, err := service.Crear(solicitud)

	assert.NoError(t, err)
	assert.True(t, mock.crearLlamado)
}

func TestObtenerSolicitudUrgenteExistente(t *testing.T) {
	mock := &mockSolicitudUrgenteRepository{}
	service := NewSolicitudUrgenteService(mock)

	solicitud, err := service.Obtener(1)

	assert.NoError(t, err)
	assert.Equal(t, 1, solicitud.ID)
}

func TestObtenerSolicitudUrgenteNoEncontrada(t *testing.T) {
	mock := &mockSolicitudUrgenteRepository{}
	service := NewSolicitudUrgenteService(mock)

	_, err := service.Obtener(99)

	assert.Error(t, err)
}

func TestBorrarSolicitudUrgenteExistente(t *testing.T) {
	mock := &mockSolicitudUrgenteRepository{}
	service := NewSolicitudUrgenteService(mock)

	err := service.Borrar(1)

	assert.NoError(t, err)
}

func TestBorrarSolicitudUrgenteNoEncontrada(t *testing.T) {
	mock := &mockSolicitudUrgenteRepository{}
	service := NewSolicitudUrgenteService(mock)

	err := service.Borrar(99)

	assert.Error(t, err)
}
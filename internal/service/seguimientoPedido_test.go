package service

import (
	"testing"

	"proyecto-detallesPersonalizados-api/internal/models"

	"github.com/stretchr/testify/assert"
)

type mockSeguimientoPedidoRepository struct {
	crearLlamado bool
}

func (m *mockSeguimientoPedidoRepository) ListarSeguimientosPedido() []models.SeguimientoPedido {
	return nil
}

func (m *mockSeguimientoPedidoRepository) BuscarSeguimientoPedidoPorID(id int) (models.SeguimientoPedido, bool) {
	if id == 1 {
		return models.SeguimientoPedido{
			ID:          1,
			PedidoID:    1,
			Estado:      "En producción",
			FechaEstado: "2026-07-09",
		}, true
	}

	return models.SeguimientoPedido{}, false
}

func (m *mockSeguimientoPedidoRepository) CrearSeguimientoPedido(s models.SeguimientoPedido) models.SeguimientoPedido {
	m.crearLlamado = true
	return s
}

func (m *mockSeguimientoPedidoRepository) ActualizarSeguimientoPedido(id int, datos models.SeguimientoPedido) (models.SeguimientoPedido, bool) {
	return models.SeguimientoPedido{}, false
}

func (m *mockSeguimientoPedidoRepository) BorrarSeguimientoPedido(id int) bool {
	return id == 1
}

// ===================== TESTS =====================

func TestCrearSeguimientoPedidoEstadoVacio(t *testing.T) {
	mock := &mockSeguimientoPedidoRepository{}
	service := NewSeguimientoPedidoService(mock)

	seguimiento := models.SeguimientoPedido{
		PedidoID:    1,
		Estado:      "",
		FechaEstado: "2026-07-09",
	}

	_, err := service.Crear(seguimiento)

	assert.Error(t, err)
	assert.False(t, mock.crearLlamado)
}

func TestCrearSeguimientoPedidoValido(t *testing.T) {
	mock := &mockSeguimientoPedidoRepository{}
	service := NewSeguimientoPedidoService(mock)

	seguimiento := models.SeguimientoPedido{
		PedidoID:    1,
		Estado:      "En producción",
		FechaEstado: "2026-07-09",
	}

	_, err := service.Crear(seguimiento)

	assert.NoError(t, err)
	assert.True(t, mock.crearLlamado)
}

func TestObtenerSeguimientoPedidoExistente(t *testing.T) {
	mock := &mockSeguimientoPedidoRepository{}
	service := NewSeguimientoPedidoService(mock)

	seguimiento, err := service.Obtener(1)

	assert.NoError(t, err)
	assert.Equal(t, 1, seguimiento.ID)
}

func TestObtenerSeguimientoPedidoNoEncontrado(t *testing.T) {
	mock := &mockSeguimientoPedidoRepository{}
	service := NewSeguimientoPedidoService(mock)

	_, err := service.Obtener(99)

	assert.Error(t, err)
}

func TestBorrarSeguimientoPedidoExistente(t *testing.T) {
	mock := &mockSeguimientoPedidoRepository{}
	service := NewSeguimientoPedidoService(mock)

	err := service.Borrar(1)

	assert.NoError(t, err)
}

func TestBorrarSeguimientoPedidoNoEncontrado(t *testing.T) {
	mock := &mockSeguimientoPedidoRepository{}
	service := NewSeguimientoPedidoService(mock)

	err := service.Borrar(99)

	assert.Error(t, err)
}


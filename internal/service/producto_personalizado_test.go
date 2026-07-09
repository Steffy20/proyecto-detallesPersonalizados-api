package service

import (
	"testing"

	"proyecto-detallesPersonalizados-api/internal/models"

	"github.com/stretchr/testify/assert"
)

type mockProductoRepository struct {
	crearLlamado bool
}

func (m *mockProductoRepository) ListarProductosPersonalizados() []models.ProductoPersonalizado {
	return nil
}

func (m *mockProductoRepository) BuscarProductoPersonalizadoPorID(id int) (models.ProductoPersonalizado, bool) {
	if id == 1 {
		return models.ProductoPersonalizado{
			ID:       1,
			PedidoID: 1,
			Nombre:   "Taza",
			Cantidad: 2,
			Precio:   10,
		}, true
	}

	return models.ProductoPersonalizado{}, false
}

func (m *mockProductoRepository) CrearProductoPersonalizado(p models.ProductoPersonalizado) models.ProductoPersonalizado {
	m.crearLlamado = true
	return p
}

func (m *mockProductoRepository) ActualizarProductoPersonalizado(id int, datos models.ProductoPersonalizado) (models.ProductoPersonalizado, bool) {
	return models.ProductoPersonalizado{}, false
}

func (m *mockProductoRepository) BorrarProductoPersonalizado(id int) bool {
	return id == 1
}

// ===================== TESTS =====================

func TestCrearProductoCantidadInvalida(t *testing.T) {

	mock := &mockProductoRepository{}

	service := NewProductoPersonalizadoService(mock)

	p := models.ProductoPersonalizado{
		PedidoID: 1,
		Nombre:   "Taza",
		Cantidad: 0,
		Precio:   10,
	}

	_, err := service.Crear(p)

	assert.Error(t, err)
	assert.False(t, mock.crearLlamado)
}

func TestCrearProductoValido(t *testing.T) {

	mock := &mockProductoRepository{}

	service := NewProductoPersonalizadoService(mock)

	p := models.ProductoPersonalizado{
		PedidoID: 1,
		Nombre:   "Taza",
		Cantidad: 2,
		Precio:   10,
	}

	_, err := service.Crear(p)

	assert.NoError(t, err)
	assert.True(t, mock.crearLlamado)
}

func TestObtenerProductoPersonalizadoExistente(t *testing.T) {
	mock := &mockProductoRepository{}
	service := NewProductoPersonalizadoService(mock)

	producto, err := service.Obtener(1)

	assert.NoError(t, err)
	assert.Equal(t, 1, producto.ID)
}

func TestObtenerProductoPersonalizadoNoEncontrado(t *testing.T) {
	mock := &mockProductoRepository{}
	service := NewProductoPersonalizadoService(mock)

	_, err := service.Obtener(99)

	assert.Error(t, err)
}

func TestBorrarProductoPersonalizadoExistente(t *testing.T) {
	mock := &mockProductoRepository{}
	service := NewProductoPersonalizadoService(mock)

	err := service.Borrar(1)

	assert.NoError(t, err)
}

func TestBorrarProductoPersonalizadoNoEncontrado(t *testing.T) {
	mock := &mockProductoRepository{}
	service := NewProductoPersonalizadoService(mock)

	err := service.Borrar(99)

	assert.Error(t, err)
}
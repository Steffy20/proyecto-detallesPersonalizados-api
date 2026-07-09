package service

import (
	"testing"

	"proyecto-detallesPersonalizados-api/internal/models"

	"github.com/stretchr/testify/assert"
)


// ===================== Creacion de mock =====================


type mockPedidoRepository struct {
	crearLlamado bool
}

func (m *mockPedidoRepository) ListarPedidos() []models.Pedido {
	return nil
}
func (m *mockPedidoRepository) BuscarPedidoPorID(id int) (models.Pedido, bool) {
	if id == 1 {
		return models.Pedido{
			ID:       1,
			Cliente:  "Juan",
			Producto: "Taza",
			Mensaje:  "Hola",
			Estado:   "Pendiente",
		}, true
	}

	return models.Pedido{}, false
}

func (m *mockPedidoRepository) CrearPedido(p models.Pedido) models.Pedido {
	m.crearLlamado = true
	return p
}

func (m *mockPedidoRepository) ActualizarPedido(id int, datos models.Pedido) (models.Pedido, bool) {
	return models.Pedido{}, false
}
func (m *mockPedidoRepository) BorrarPedido(id int) bool {
	return id == 1
}

// ===================== TESTS =====================

func TestCrearPedidoMensajeVacio(t *testing.T) {

	mock := &mockPedidoRepository{}

	service := NewPedidoService(mock)

	pedido := models.Pedido{
		Cliente:  "Juan",
		Producto: "Taza",
		Mensaje:  "",
		Estado:   "Pendiente",
	}

	_, err := service.Crear(pedido)

	assert.Error(t, err)
	assert.False(t, mock.crearLlamado)
}

func TestCrearPedidoValido(t *testing.T) {

	mock := &mockPedidoRepository{}

	service := NewPedidoService(mock)

	pedido := models.Pedido{
		Cliente:  "Juan",
		Producto: "Taza",
		Mensaje:  "Feliz cumpleaños",
	}

	_, err := service.Crear(pedido)

	assert.NoError(t, err)
	assert.True(t, mock.crearLlamado)
}

func TestObtenerPedidoExistente(t *testing.T) {
	mock := &mockPedidoRepository{}
	service := NewPedidoService(mock)

	pedido, err := service.Obtener(1)

	assert.NoError(t, err)
	assert.Equal(t, 1, pedido.ID)
}

func TestObtenerPedidoNoEncontrado(t *testing.T) {
	mock := &mockPedidoRepository{}
	service := NewPedidoService(mock)

	_, err := service.Obtener(99)

	assert.Error(t, err)
}

func TestBorrarPedidoExistente(t *testing.T) {
	mock := &mockPedidoRepository{}
	service := NewPedidoService(mock)

	err := service.Borrar(1)

	assert.NoError(t, err)
}

func TestBorrarPedidoNoEncontrado(t *testing.T) {
	mock := &mockPedidoRepository{}
	service := NewPedidoService(mock)

	err := service.Borrar(99)

	assert.Error(t, err)
}
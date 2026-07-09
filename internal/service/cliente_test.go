package service

import (
	"testing"

	"proyecto-detallesPersonalizados-api/internal/models"

	"github.com/stretchr/testify/assert"
)

type mockClienteRepository struct {
	crearLlamado bool
}

func (m *mockClienteRepository) ListarClientes() []models.Cliente {
	return nil
}

func (m *mockClienteRepository) BuscarClientePorID(id int) (models.Cliente, bool) {
	if id == 1 {
		return models.Cliente{
			ID:       1,
			Nombre:   "Juan",
			Telefono: "0999999999",
			Correo:   "cliente@gmail.com",
		}, true
	}

	return models.Cliente{}, false
}

func (m *mockClienteRepository) CrearCliente(c models.Cliente) models.Cliente {
	m.crearLlamado = true
	return c
}

func (m *mockClienteRepository) ActualizarCliente(id int, datos models.Cliente) (models.Cliente, bool) {
	return models.Cliente{}, false
}

func (m *mockClienteRepository) BorrarCliente(id int) bool {
	return id == 1
}

// ===================== TESTS =====================

func TestCrearClienteNombreVacio(t *testing.T) {
	mock := &mockClienteRepository{}
	service := NewClienteService(mock)

	cliente := models.Cliente{
		Nombre:   "",
		Telefono: "0999999999",
		Correo:   "cliente@gmail.com",
	}

	_, err := service.Crear(cliente)

	assert.Error(t, err)
	assert.False(t, mock.crearLlamado)
}

func TestCrearClienteValido(t *testing.T) {
	mock := &mockClienteRepository{}
	service := NewClienteService(mock)

	cliente := models.Cliente{
		Nombre:   "Juan",
		Telefono: "0999999999",
		Correo:   "cliente@gmail.com",
	}

	_, err := service.Crear(cliente)

	assert.NoError(t, err)
	assert.True(t, mock.crearLlamado)
}

func TestObtenerClienteExistente(t *testing.T) {
	mock := &mockClienteRepository{}
	service := NewClienteService(mock)

	cliente, err := service.Obtener(1)

	assert.NoError(t, err)
	assert.Equal(t, 1, cliente.ID)
}

func TestObtenerClienteNoEncontrado(t *testing.T) {
	mock := &mockClienteRepository{}
	service := NewClienteService(mock)

	_, err := service.Obtener(99)

	assert.Error(t, err)
}

func TestBorrarClienteExistente(t *testing.T) {
	mock := &mockClienteRepository{}
	service := NewClienteService(mock)

	err := service.Borrar(1)

	assert.NoError(t, err)
}

func TestBorrarClienteNoEncontrado(t *testing.T) {
	mock := &mockClienteRepository{}
	service := NewClienteService(mock)

	err := service.Borrar(99)

	assert.Error(t, err)
}
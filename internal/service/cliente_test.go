package service

import (
	"testing"

	"proyecto-detallesPersonalizados-api/internal/models"
)

type mockClienteRepository struct {
	crearLlamado bool
}

func (m *mockClienteRepository) ListarClientes() []models.Cliente {
	return nil
}

func (m *mockClienteRepository) BuscarClientePorID(id int) (models.Cliente, bool) {
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
	return false
}

func TestCrearClienteNombreVacio(t *testing.T) {
	mock := &mockClienteRepository{}

	service := NewClienteService(mock)

	cliente := models.Cliente{
		Nombre:   "",
		Telefono: "0999999999",
		Correo:   "cliente@gmail.com",
	}

	_, err := service.Crear(cliente)

	if err == nil {
		t.Fatal("se esperaba un error por nombre vacío")
	}

	if mock.crearLlamado {
		t.Fatal("el repositorio no debía ser llamado")
	}
}
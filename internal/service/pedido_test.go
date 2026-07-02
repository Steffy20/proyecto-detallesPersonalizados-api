package service

import (
	"testing"

	"proyecto-detallesPersonalizados-api/internal/models"
)

//================= CREACION DE MOCK =================

type mockPedidoRepository struct {
	crearLlamado bool
}

func (m *mockPedidoRepository) ListarPedidos() []models.Pedido {
	return nil
}

func (m *mockPedidoRepository) BuscarPedidoPorID(id int) (models.Pedido, bool) {
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
	return false
}

// ================= TEST 1 =================

func TestCrearPedidoMensajeVacio(t *testing.T) {

	mock := &mockPedidoRepository{}

	service := NewPedidoService(mock)

	pedido := models.Pedido{
		Cliente:  "Juan",
		Producto: "Taza",
		Mensaje:  "",
	}

	_, err := service.Crear(pedido)

	if err == nil {
		t.Fatal("se esperaba un error")
	}

	if mock.crearLlamado {
		t.Fatal("el repositorio no debía ser llamado")
	}
}


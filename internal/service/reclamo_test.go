package service

import (
	"testing"

	"proyecto-detallesPersonalizados-api/internal/models"

	"github.com/stretchr/testify/assert"
)

type mockReclamoRepository struct {
	crearLlamado bool
}

func (m *mockReclamoRepository) ListarReclamos() []models.Reclamo {
	return nil
}

func (m *mockReclamoRepository) BuscarReclamoPorID(id int) (models.Reclamo, bool) {
	if id == 1 {
		return models.Reclamo{
			ID:          1,
			ClienteID:   1,
			PedidoID:    1,
			Descripcion: "Producto llegó con error",
			Estado:      "Pendiente",
		}, true
	}

	return models.Reclamo{}, false
}

func (m *mockReclamoRepository) CrearReclamo(r models.Reclamo) models.Reclamo {
	m.crearLlamado = true
	return r
}

func (m *mockReclamoRepository) ActualizarReclamo(id int, datos models.Reclamo) (models.Reclamo, bool) {
	return models.Reclamo{}, false
}

func (m *mockReclamoRepository) BorrarReclamo(id int) bool {
	return id == 1
}

// ===================== TESTS =====================

func TestCrearReclamoDescripcionVacia(t *testing.T) {
	mock := &mockReclamoRepository{}
	service := NewReclamoService(mock)

	reclamo := models.Reclamo{
		ClienteID:   1,
		PedidoID:    1,
		Descripcion: "",
		Estado:      "Pendiente",
	}

	_, err := service.Crear(reclamo)

	assert.Error(t, err)
	assert.False(t, mock.crearLlamado)
}

func TestCrearReclamoValido(t *testing.T) {
	mock := &mockReclamoRepository{}
	service := NewReclamoService(mock)

	reclamo := models.Reclamo{
		ClienteID:   1,
		PedidoID:    1,
		Descripcion: "Producto llegó con error",
		Estado:      "Pendiente",
	}

	_, err := service.Crear(reclamo)

	assert.NoError(t, err)
	assert.True(t, mock.crearLlamado)
}

func TestObtenerReclamoExistente(t *testing.T) {
	mock := &mockReclamoRepository{}
	service := NewReclamoService(mock)

	reclamo, err := service.Obtener(1)

	assert.NoError(t, err)
	assert.Equal(t, 1, reclamo.ID)
}

func TestObtenerReclamoNoEncontrado(t *testing.T) {
	mock := &mockReclamoRepository{}
	service := NewReclamoService(mock)

	_, err := service.Obtener(99)

	assert.Error(t, err)
}

func TestBorrarReclamoExistente(t *testing.T) {
	mock := &mockReclamoRepository{}
	service := NewReclamoService(mock)

	err := service.Borrar(1)

	assert.NoError(t, err)
}

func TestBorrarReclamoNoEncontrado(t *testing.T) {
	mock := &mockReclamoRepository{}
	service := NewReclamoService(mock)

	err := service.Borrar(99)

	assert.Error(t, err)
}

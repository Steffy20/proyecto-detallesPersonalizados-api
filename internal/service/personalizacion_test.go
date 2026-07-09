package service

import (
	"testing"

	"proyecto-detallesPersonalizados-api/internal/models"

	"github.com/stretchr/testify/assert"
)

type mockPersonalizacionRepository struct {
	crearLlamado bool
}

func (m *mockPersonalizacionRepository) ListarPersonalizaciones() []models.Personalizacion {
	return nil
}

func (m *mockPersonalizacionRepository) BuscarPersonalizacionPorID(id int) (models.Personalizacion, bool) {
	if id == 1 {
		return models.Personalizacion{
			ID:       1,
			PedidoID: 1,
			Mensaje:  "Hola",
			Color:    "Rojo",
		}, true
	}

	return models.Personalizacion{}, false
}

func (m *mockPersonalizacionRepository) CrearPersonalizacion(p models.Personalizacion) models.Personalizacion {
	m.crearLlamado = true
	return p
}

func (m *mockPersonalizacionRepository) ActualizarPersonalizacion(id int, datos models.Personalizacion) (models.Personalizacion, bool) {
	return models.Personalizacion{}, false
}

func (m *mockPersonalizacionRepository) BorrarPersonalizacion(id int) bool {
	return id == 1
}
// ===================== TESTS =====================

func TestCrearPersonalizacionColorVacio(t *testing.T) {

	mock := &mockPersonalizacionRepository{}

	service := NewPersonalizacionService(mock)

	p := models.Personalizacion{
		PedidoID: 1,
		Mensaje:  "Hola",
		Color:    "",
	}

	_, err := service.Crear(p)

	assert.Error(t, err)
	assert.False(t, mock.crearLlamado)
}

func TestCrearPersonalizacionValida(t *testing.T) {

	mock := &mockPersonalizacionRepository{}

	service := NewPersonalizacionService(mock)

	p := models.Personalizacion{
		PedidoID: 1,
		Mensaje:  "Hola",
		Color:    "Rojo",
	}

	_, err := service.Crear(p)

	assert.NoError(t, err)
	assert.True(t, mock.crearLlamado)
}

func TestObtenerPersonalizacionExistente(t *testing.T) {
	mock := &mockPersonalizacionRepository{}
	service := NewPersonalizacionService(mock)

	personalizacion, err := service.Obtener(1)

	assert.NoError(t, err)
	assert.Equal(t, 1, personalizacion.ID)
}

func TestObtenerPersonalizacionNoEncontrada(t *testing.T) {
	mock := &mockPersonalizacionRepository{}
	service := NewPersonalizacionService(mock)

	_, err := service.Obtener(99)

	assert.Error(t, err)
}

func TestBorrarPersonalizacionExistente(t *testing.T) {
	mock := &mockPersonalizacionRepository{}
	service := NewPersonalizacionService(mock)

	err := service.Borrar(1)

	assert.NoError(t, err)
}

func TestBorrarPersonalizacionNoEncontrada(t *testing.T) {
	mock := &mockPersonalizacionRepository{}
	service := NewPersonalizacionService(mock)

	err := service.Borrar(99)

	assert.Error(t, err)
}

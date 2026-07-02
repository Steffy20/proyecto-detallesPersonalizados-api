package service

import (
	"testing"

	"proyecto-detallesPersonalizados-api/internal/models"
)

type mockSolicitudUrgenteRepository struct {
	crearLlamado bool
}

func (m *mockSolicitudUrgenteRepository) ListarSolicitudesUrgentes() []models.SolicitudUrgente {
	return nil
}

func (m *mockSolicitudUrgenteRepository) BuscarSolicitudUrgentePorID(id int) (models.SolicitudUrgente, bool) {
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
	return false
}

func TestCrearSolicitudUrgenteClienteVacio(t *testing.T) {
	mock := &mockSolicitudUrgenteRepository{}

	service := NewSolicitudUrgenteService(mock)

	solicitud := models.SolicitudUrgente{
		Cliente:        "",
		Descripcion:    "Pedido urgente para hoy",
		FechaRequerida: "2026-07-01",
	}

	_, err := service.Crear(solicitud)

	if err == nil {
		t.Fatal("se esperaba un error por cliente vacío")
	}

	if mock.crearLlamado {
		t.Fatal("el repositorio no debía ser llamado")
	}
}
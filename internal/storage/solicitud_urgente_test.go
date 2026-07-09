package storage

import (
	"testing"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"

	"proyecto-detallesPersonalizados-api/internal/models"
)

func TestRepositorioSolicitudUrgenteCrearYListar(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("error al abrir sqlite en memoria: %v", err)
	}

	if err := db.AutoMigrate(&models.SolicitudUrgente{}); err != nil {
		t.Fatalf("error en AutoMigrate: %v", err)
	}

	almacen := NuevoAlmacenSQLite(db)

	solicitud := models.SolicitudUrgente{
		Cliente:        "María",
		Descripcion:    "Pedido urgente para cumpleaños",
		FechaRequerida: "2026-07-01",
		Estado:         "Pendiente",
	}

	creada := almacen.CrearSolicitudUrgente(solicitud)

	if creada.ID == 0 {
		t.Fatal("la solicitud urgente no fue creada")
	}

	solicitudes := almacen.ListarSolicitudesUrgentes()

	if len(solicitudes) != 1 {
		t.Fatalf("se esperaba 1 solicitud urgente y se obtuvieron %d", len(solicitudes))
	}

	if solicitudes[0].Cliente != "María" {
		t.Fatal("la solicitud urgente recuperada no coincide")
	}
}

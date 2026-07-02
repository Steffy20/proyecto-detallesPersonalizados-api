package storage

import (
	"testing"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"

	"proyecto-detallesPersonalizados-api/internal/models"
)

func TestRepositorioClienteCrearYListar(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("error al abrir sqlite en memoria: %v", err)
	}

	if err := db.AutoMigrate(&models.Cliente{}); err != nil {
		t.Fatalf("error en AutoMigrate: %v", err)
	}

	almacen := NuevoAlmacenSQLite(db)

	cliente := models.Cliente{
		Nombre:   "Tonny",
		Telefono: "0999999999",
		Correo:   "Tonny@gmail.com",
	}

	creado := almacen.CrearCliente(cliente)

	if creado.ID == 0 {
		t.Fatal("el cliente no fue creado")
	}

	clientes := almacen.ListarClientes()

	if len(clientes) != 1 {
		t.Fatalf("se esperaba 1 cliente y se obtuvieron %d", len(clientes))
	}

	if clientes[0].Nombre != "María" {
		t.Fatal("el cliente recuperado no coincide")
	}
}
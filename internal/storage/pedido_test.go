package storage

import (
	"testing"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"

	"proyecto-detallesPersonalizados-api/internal/models"
)

func TestRepositorioPedidoCrearYListar(t *testing.T) {

	// Base de datos temporal en memoria
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("error al abrir la base de datos: %v", err)
	}

	// Crear la tabla Pedido
	if err := db.AutoMigrate(&models.Pedido{}); err != nil {
		t.Fatalf("error en AutoMigrate: %v", err)
	}

	// Crear el repositorio real
	almacen := &AlmacenSQLite{
		db: db,
	}

	// Crear un pedido
	pedido := models.Pedido{
		Cliente:  "Juan",
		Producto: "Taza personalizada",
		Mensaje:  "Feliz cumpleaños",
		Estado:   "Pendiente",
	}

	creado := almacen.CrearPedido(pedido)

	if creado.ID == 0 {
		t.Fatal("el pedido no fue creado correctamente")
	}

	// Listar pedidos
	pedidos := almacen.ListarPedidos()

	if len(pedidos) != 1 {
		t.Fatalf("se esperaba 1 pedido, se obtuvo %d", len(pedidos))
	}

	if pedidos[0].Cliente != "Juan" {
		t.Fatal("el pedido recuperado no coincide")
	}
}

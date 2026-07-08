package storage

import (
	"fmt"
	"strings"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"proyecto-detallesPersonalizados-api/internal/models"
)

type AlmacenSQLite struct {
	db *gorm.DB
}

func NuevoAlmacenSQLite(db *gorm.DB) *AlmacenSQLite {
	return &AlmacenSQLite{db: db}
}

func InicializarBaseDatos(driver, rutaDB, dsn string) (*gorm.DB, *AlmacenSQLite, error) {
	var dialector gorm.Dialector

	switch strings.ToLower(strings.TrimSpace(driver)) {
	case "", "sqlite":
		dialector = sqlite.Open(rutaDB)
	case "postgres", "postgresql":
		if strings.TrimSpace(dsn) == "" {
			return nil, nil, fmt.Errorf("DB_DSN es obligatorio cuando DB_DRIVER=%q", driver)
		}
		dialector = postgres.Open(dsn)
	default:
		return nil, nil, fmt.Errorf("DB_DRIVER no soportado: %q", driver)
	}

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}

	if err := db.AutoMigrate(
		&models.Pedido{},
		&models.Personalizacion{},
		&models.ProductoPersonalizado{},
		&models.SolicitudUrgente{},
		&models.AgendaProduccion{},
		&models.SlotProduccion{},
		&models.Cliente{},
		&models.Reclamo{},
		&models.SeguimientoPedido{},
		&models.Usuario{},
	); err != nil {
		return nil, nil, err
	}

	return db, NuevoAlmacenSQLite(db), nil
}

func InicializarSQLite(rutaDB string) (*gorm.DB, *AlmacenSQLite, error) {
	return InicializarBaseDatos("sqlite", rutaDB, "")
}

var _ Almacen = (*AlmacenSQLite)(nil)

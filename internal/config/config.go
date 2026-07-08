package config

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	defaultPuerto          = "8080"
	defaultDBDriver        = "sqlite"
	defaultRutaDB          = "detallesPersonalizados.db"
	defaultJWTSecreto      = "detalles_personalizados_api"
	defaultJWTDuracion     = 24 * time.Hour
	defaultShutdownTimeout = 10 * time.Second
	defaultReadTimeout     = 5 * time.Second
	defaultWriteTimeout    = 10 * time.Second
	defaultIdleTimeout     = 60 * time.Second
)

type Config struct {
	Puerto          string
	DBDriver        string
	DBDSN           string
	RutaDB          string
	JWTSecreto      []byte
	JWTDuracion     time.Duration
	ShutdownTimeout time.Duration
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	IdleTimeout     time.Duration
}

func Cargar() Config {
	cargarDotEnv(".env")

	return Config{
		Puerto:          envString("PUERTO", defaultPuerto),
		DBDriver:        strings.ToLower(envString("DB_DRIVER", defaultDBDriver)),
		DBDSN:           envString("DB_DSN", ""),
		RutaDB:          envString("RUTA_DB", defaultRutaDB),
		JWTSecreto:      []byte(envString("JWT_SECRETO", defaultJWTSecreto)),
		JWTDuracion:     envDurationHours("JWT_DURACION_HORAS", defaultJWTDuracion),
		ShutdownTimeout: envDurationSeconds("SHUTDOWN_TIMEOUT_SEGUNDOS", defaultShutdownTimeout),
		ReadTimeout:     envDurationSeconds("HTTP_READ_TIMEOUT_SEGUNDOS", defaultReadTimeout),
		WriteTimeout:    envDurationSeconds("HTTP_WRITE_TIMEOUT_SEGUNDOS", defaultWriteTimeout),
		IdleTimeout:     envDurationSeconds("HTTP_IDLE_TIMEOUT_SEGUNDOS", defaultIdleTimeout),
	}
}

func cargarDotEnv(ruta string) {
	file, err := os.Open(ruta)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linea := strings.TrimSpace(scanner.Text())
		if linea == "" || strings.HasPrefix(linea, "#") {
			continue
		}

		clave, valor, ok := strings.Cut(linea, "=")
		if !ok {
			continue
		}

		clave = strings.TrimSpace(clave)
		valor = strings.Trim(strings.TrimSpace(valor), `"'`)

		if clave != "" && os.Getenv(clave) == "" {
			_ = os.Setenv(clave, valor)
		}
	}
}

func envString(clave, defecto string) string {
	if valor := strings.TrimSpace(os.Getenv(clave)); valor != "" {
		return valor
	}
	return defecto
}

func envDurationHours(clave string, defecto time.Duration) time.Duration {
	return envDuration(clave, defecto, time.Hour)
}

func envDurationSeconds(clave string, defecto time.Duration) time.Duration {
	return envDuration(clave, defecto, time.Second)
}

func envDuration(clave string, defecto, unidad time.Duration) time.Duration {
	valor := strings.TrimSpace(os.Getenv(clave))
	if valor == "" {
		return defecto
	}

	numero, err := strconv.Atoi(valor)
	if err != nil || numero <= 0 {
		return defecto
	}

	return time.Duration(numero) * unidad
}

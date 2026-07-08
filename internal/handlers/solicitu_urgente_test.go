package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"

	"proyecto-detallesPersonalizados-api/internal/models"
)

type fakeSolicitudUrgenteService struct {
	solicitudes []models.SolicitudUrgente
	nextID      int
}

func (f *fakeSolicitudUrgenteService) Listar() []models.SolicitudUrgente {
	return f.solicitudes
}

func (f *fakeSolicitudUrgenteService) Obtener(id int) (models.SolicitudUrgente, error) {
	for _, s := range f.solicitudes {
		if s.ID == id {
			return s, nil
		}
	}
	return models.SolicitudUrgente{}, errors.New("solicitud urgente no encontrada")
}

func (f *fakeSolicitudUrgenteService) Crear(s models.SolicitudUrgente) (models.SolicitudUrgente, error) {
	if s.Cliente == "" {
		return models.SolicitudUrgente{}, errors.New("el cliente es obligatorio")
	}

	f.nextID++
	s.ID = f.nextID

	if s.Estado == "" {
		s.Estado = "Pendiente"
	}

	f.solicitudes = append(f.solicitudes, s)

	return s, nil
}

func (f *fakeSolicitudUrgenteService) Actualizar(id int, datos models.SolicitudUrgente) (models.SolicitudUrgente, error) {
	for i, s := range f.solicitudes {
		if s.ID == id {
			datos.ID = id
			f.solicitudes[i] = datos
			return datos, nil
		}
	}
	return models.SolicitudUrgente{}, errors.New("solicitud urgente no encontrada")
}

func (f *fakeSolicitudUrgenteService) Borrar(id int) error {
	for i, s := range f.solicitudes {
		if s.ID == id {
			f.solicitudes = append(f.solicitudes[:i], f.solicitudes[i+1:]...)
			return nil
		}
	}
	return errors.New("solicitud urgente no encontrada")
}

func TestCrearSolicitudUrgenteHandler(t *testing.T) {
	fake := &fakeSolicitudUrgenteService{}

	server := &Server{
		SolicitudesUrgentes: fake,
	}

	body := models.SolicitudUrgente{
		Cliente:        "María",
		Descripcion:    "Pedido urgente para cumpleaños",
		FechaRequerida: "2026-07-01",
	}

	jsonBody, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/solicitudes-urgentes/", bytes.NewReader(jsonBody))
	rec := httptest.NewRecorder()

	server.CrearSolicitudUrgente(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("se esperaba status 201, se obtuvo %d", rec.Code)
	}

	if len(fake.solicitudes) != 1 {
		t.Fatalf("se esperaba 1 solicitud guardada, se obtuvo %d", len(fake.solicitudes))
	}
}

func TestRutaSolicitudUrgenteProtegidaSinTokenRetorna401(t *testing.T) {
	fake := &fakeSolicitudUrgenteService{}

	server := &Server{
		SolicitudesUrgentes: fake,
	}

	r := chi.NewRouter()

	authMiddlewareFake := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get("Authorization") == "" {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}

	r.Group(func(r chi.Router) {
		r.Use(authMiddlewareFake)

		r.Route("/api/v1/solicitudes-urgentes", func(r chi.Router) {
			r.Post("/", server.CrearSolicitudUrgente)
		})
	})

	body := models.SolicitudUrgente{
		Cliente:        "María",
		Descripcion:    "Pedido urgente para cumpleaños",
		FechaRequerida: "2026-07-01",
	}

	jsonBody, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/solicitudes-urgentes/", bytes.NewReader(jsonBody))
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("se esperaba status 401, se obtuvo %d", rec.Code)
	}
}
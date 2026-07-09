package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"

	"proyecto-detallesPersonalizados-api/internal/models"
)

// ===================== FAKE SERVICE =====================

type fakeAgendaProduccionService struct {
	agendas []models.AgendaProduccion
	nextID  int
}

func (f *fakeAgendaProduccionService) Listar() []models.AgendaProduccion {
	return f.agendas
}

func (f *fakeAgendaProduccionService) Obtener(id int) (models.AgendaProduccion, error) {
	for _, a := range f.agendas {
		if a.ID == id {
			return a, nil
		}
	}
	return models.AgendaProduccion{}, errors.New("agenda de producción no encontrada")
}

func (f *fakeAgendaProduccionService) Crear(a models.AgendaProduccion) (models.AgendaProduccion, error) {
	if a.Fecha == "" {
		return models.AgendaProduccion{}, errors.New("la fecha es obligatoria")
	}
	if a.Responsable == "" {
		return models.AgendaProduccion{}, errors.New("el responsable es obligatorio")
	}
	if a.Estado == "" {
		a.Estado = "Pendiente"
	}
	f.nextID++
	a.ID = f.nextID
	f.agendas = append(f.agendas, a)
	return a, nil
}

func (f *fakeAgendaProduccionService) Actualizar(id int, datos models.AgendaProduccion) (models.AgendaProduccion, error) {
	for i, a := range f.agendas {
		if a.ID == id {
			datos.ID = id
			f.agendas[i] = datos
			return datos, nil
		}
	}
	return models.AgendaProduccion{}, errors.New("agenda de producción no encontrada")
}

func (f *fakeAgendaProduccionService) Borrar(id int) error {
	for i, a := range f.agendas {
		if a.ID == id {
			f.agendas = append(f.agendas[:i], f.agendas[i+1:]...)
			return nil
		}
	}
	return errors.New("agenda de producción no encontrada")
}

// ===================== HELPERS =====================

func nuevoServerAgenda(fake *fakeAgendaProduccionService) *Server {
	return &Server{AgendasProduccion: fake}
}

func routerAgenda(server *Server) *chi.Mux {
	r := chi.NewRouter()
	r.Route("/api/v1/agendas-produccion", func(r chi.Router) {
		r.Post("/", server.CrearAgendaProduccion)
		r.Get("/", server.ObtenerAgendasProduccion)
		r.Get("/{id}", server.ObtenerAgendaProduccionPorID)
		r.Put("/{id}", server.ActualizarAgendaProduccion)
		r.Delete("/{id}", server.EliminarAgendaProduccion)
	})
	return r
}

// ===================== TESTS =====================

func TestCrearAgendaProduccion_OK(t *testing.T) {
	fake := &fakeAgendaProduccionService{}
	server := nuevoServerAgenda(fake)

	body := models.AgendaProduccion{
		Fecha:       "2026-07-15",
		Responsable: "Equipo A",
		Estado:      "programado",
	}
	jsonBody, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/agendas-produccion/", bytes.NewReader(jsonBody))
	rec := httptest.NewRecorder()

	server.CrearAgendaProduccion(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("esperaba 201, obtuvo %d", rec.Code)
	}
	if len(fake.agendas) != 1 {
		t.Fatalf("esperaba 1 agenda guardada, obtuvo %d", len(fake.agendas))
	}
	if fake.agendas[0].Responsable != "Equipo A" {
		t.Errorf("responsable incorrecto: %s", fake.agendas[0].Responsable)
	}
}

func TestCrearAgendaProduccion_SinFecha_Retorna400(t *testing.T) {
	fake := &fakeAgendaProduccionService{}
	server := nuevoServerAgenda(fake)

	body := models.AgendaProduccion{
		Responsable: "Equipo B",
	}
	jsonBody, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/agendas-produccion/", bytes.NewReader(jsonBody))
	rec := httptest.NewRecorder()

	server.CrearAgendaProduccion(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("esperaba 400 por fecha vacía, obtuvo %d", rec.Code)
	}
}

func TestListarAgendasProduccion_OK(t *testing.T) {
	fake := &fakeAgendaProduccionService{
		agendas: []models.AgendaProduccion{
			{ID: 1, Fecha: "2026-07-10", Responsable: "Equipo A", Estado: "programado"},
			{ID: 2, Fecha: "2026-07-11", Responsable: "Equipo B", Estado: "completado"},
		},
	}
	server := nuevoServerAgenda(fake)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/agendas-produccion/", nil)
	rec := httptest.NewRecorder()

	server.ObtenerAgendasProduccion(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("esperaba 200, obtuvo %d", rec.Code)
	}

	var resultado []models.AgendaProduccion
	if err := json.NewDecoder(rec.Body).Decode(&resultado); err != nil {
		t.Fatal("error decodificando respuesta:", err)
	}
	if len(resultado) != 2 {
		t.Fatalf("esperaba 2 agendas, obtuvo %d", len(resultado))
	}
}

func TestObtenerAgendaProduccionPorID_OK(t *testing.T) {
	fake := &fakeAgendaProduccionService{
		agendas: []models.AgendaProduccion{
			{ID: 1, Fecha: "2026-07-10", Responsable: "Equipo A", Estado: "programado"},
		},
	}
	server := nuevoServerAgenda(fake)
	r := routerAgenda(server)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/agendas-produccion/1", nil)
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("esperaba 200, obtuvo %d", rec.Code)
	}

	var resultado models.AgendaProduccion
	if err := json.NewDecoder(rec.Body).Decode(&resultado); err != nil {
		t.Fatal("error decodificando respuesta:", err)
	}
	if resultado.ID != 1 {
		t.Errorf("ID incorrecto: %d", resultado.ID)
	}
}

func TestObtenerAgendaProduccionPorID_NoExiste_Retorna404(t *testing.T) {
	fake := &fakeAgendaProduccionService{}
	server := nuevoServerAgenda(fake)
	r := routerAgenda(server)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/agendas-produccion/99", nil)
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("esperaba 404, obtuvo %d", rec.Code)
	}
}

func TestActualizarAgendaProduccion_OK(t *testing.T) {
	fake := &fakeAgendaProduccionService{
		agendas: []models.AgendaProduccion{
			{ID: 1, Fecha: "2026-07-10", Responsable: "Equipo A", Estado: "programado"},
		},
	}
	server := nuevoServerAgenda(fake)
	r := routerAgenda(server)

	actualizado := models.AgendaProduccion{
		Fecha:       "2026-07-20",
		Responsable: "Equipo C",
		Estado:      "en_proceso",
	}
	jsonBody, _ := json.Marshal(actualizado)

	req := httptest.NewRequest(http.MethodPut, "/api/v1/agendas-produccion/1", bytes.NewReader(jsonBody))
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("esperaba 200, obtuvo %d", rec.Code)
	}
	if fake.agendas[0].Responsable != "Equipo C" {
		t.Errorf("responsable no actualizado: %s", fake.agendas[0].Responsable)
	}
}

func TestActualizarAgendaProduccion_NoExiste_Retorna404(t *testing.T) {
	fake := &fakeAgendaProduccionService{}
	server := nuevoServerAgenda(fake)
	r := routerAgenda(server)

	body := models.AgendaProduccion{
		Fecha:       "2026-07-20",
		Responsable: "Equipo X",
		Estado:      "programado",
	}
	jsonBody, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPut, "/api/v1/agendas-produccion/99", bytes.NewReader(jsonBody))
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("esperaba 404, obtuvo %d", rec.Code)
	}
}

func TestEliminarAgendaProduccion_OK(t *testing.T) {
	fake := &fakeAgendaProduccionService{
		agendas: []models.AgendaProduccion{
			{ID: 1, Fecha: "2026-07-10", Responsable: "Equipo A", Estado: "programado"},
		},
	}
	server := nuevoServerAgenda(fake)
	r := routerAgenda(server)

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/agendas-produccion/1", nil)
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusNoContent {
		t.Fatalf("esperaba 204, obtuvo %d", rec.Code)
	}
	if len(fake.agendas) != 0 {
		t.Fatalf("esperaba 0 agendas tras eliminar, obtuvo %d", len(fake.agendas))
	}
}

func TestEliminarAgendaProduccion_NoExiste_Retorna404(t *testing.T) {
	fake := &fakeAgendaProduccionService{}
	server := nuevoServerAgenda(fake)
	r := routerAgenda(server)

	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/api/v1/agendas-produccion/%d", 99), nil)
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("esperaba 404, obtuvo %d", rec.Code)
	}
}

func TestRutaAgendaProduccionProtegidaSinTokenRetorna401(t *testing.T) {
	fake := &fakeAgendaProduccionService{}
	server := nuevoServerAgenda(fake)

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
		r.Route("/api/v1/agendas-produccion", func(r chi.Router) {
			r.Post("/", server.CrearAgendaProduccion)
		})
	})

	body := models.AgendaProduccion{
		Fecha:       "2026-07-15",
		Responsable: "Equipo A",
	}
	jsonBody, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/agendas-produccion/", bytes.NewReader(jsonBody))
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("esperaba 401, obtuvo %d", rec.Code)
	}
}

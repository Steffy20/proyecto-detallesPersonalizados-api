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

type fakeClienteService struct {
	clientes []models.Cliente
	nextID   int
}

func (f *fakeClienteService) Listar() []models.Cliente {
	return f.clientes
}

func (f *fakeClienteService) Obtener(id int) (models.Cliente, error) {
	for _, c := range f.clientes {
		if c.ID == id {
			return c, nil
		}
	}
	return models.Cliente{}, errors.New("cliente no encontrado")
}

func (f *fakeClienteService) Crear(c models.Cliente) (models.Cliente, error) {
	if c.Nombre == "" {
		return models.Cliente{}, errors.New("el nombre es obligatorio")
	}

	f.nextID++
	c.ID = f.nextID

	f.clientes = append(f.clientes, c)

	return c, nil
}

func (f *fakeClienteService) Actualizar(id int, datos models.Cliente) (models.Cliente, error) {
	for i, c := range f.clientes {
		if c.ID == id {
			datos.ID = id
			f.clientes[i] = datos
			return datos, nil
		}
	}

	return models.Cliente{}, errors.New("cliente no encontrado")
}

func (f *fakeClienteService) Borrar(id int) error {
	for i, c := range f.clientes {
		if c.ID == id {
			f.clientes = append(f.clientes[:i], f.clientes[i+1:]...)
			return nil
		}
	}

	return errors.New("cliente no encontrado")
}

func TestCrearClienteHandler(t *testing.T) {

	fake := &fakeClienteService{}

	server := &Server{
		Clientes: fake,
	}

	body := models.Cliente{
		Nombre:   "Juan Pérez",
		Telefono: "0999999999",
		Correo:   "juan@gmail.com",
	}

	jsonBody, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/clientes/", bytes.NewReader(jsonBody))
	rec := httptest.NewRecorder()

	server.CrearCliente(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("se esperaba status 201, se obtuvo %d", rec.Code)
	}

	if len(fake.clientes) != 1 {
		t.Fatalf("se esperaba 1 cliente guardado, se obtuvo %d", len(fake.clientes))
	}
}

func TestRutaClienteProtegidaSinTokenRetorna401(t *testing.T) {

	fake := &fakeClienteService{}

	server := &Server{
		Clientes: fake,
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

		r.Route("/api/v1/clientes", func(r chi.Router) {
			r.Post("/", server.CrearCliente)
		})
	})

	body := models.Cliente{
		Nombre:   "Juan Pérez",
		Telefono: "0999999999",
		Correo:   "juan@gmail.com",
	}

	jsonBody, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/clientes/", bytes.NewReader(jsonBody))
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("se esperaba status 401, se obtuvo %d", rec.Code)
	}
}
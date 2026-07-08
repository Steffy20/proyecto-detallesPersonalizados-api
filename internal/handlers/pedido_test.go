package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"

	"proyecto-detallesPersonalizados-api/internal/middleware"
	"proyecto-detallesPersonalizados-api/internal/models"
)

type fakePedidoService struct {
	pedidos []models.Pedido
	nextID  int
}

func (f *fakePedidoService) Listar() []models.Pedido {
	return f.pedidos
}

func (f *fakePedidoService) Obtener(id int) (models.Pedido, error) {
	for _, p := range f.pedidos {
		if p.ID == id {
			return p, nil
		}
	}
	return models.Pedido{}, errors.New("pedido no encontrado")
}

func (f *fakePedidoService) Crear(p models.Pedido) (models.Pedido, error) {
	if p.Mensaje == "" {
		return models.Pedido{}, errors.New("el mensaje personalizado es obligatorio")
	}

	f.nextID++
	p.ID = f.nextID

	if p.Estado == "" {
		p.Estado = "Pendiente"
	}

	f.pedidos = append(f.pedidos, p)

	return p, nil
}

func (f *fakePedidoService) Actualizar(id int, datos models.Pedido) (models.Pedido, error) {
	for i, p := range f.pedidos {
		if p.ID == id {
			datos.ID = id
			f.pedidos[i] = datos
			return datos, nil
		}
	}
	return models.Pedido{}, errors.New("pedido no encontrado")
}

func (f *fakePedidoService) Borrar(id int) error {
	for i, p := range f.pedidos {
		if p.ID == id {
			f.pedidos = append(f.pedidos[:i], f.pedidos[i+1:]...)
			return nil
		}
	}
	return errors.New("pedido no encontrado")
}

// PRIMER TEST DE HANDLERS

func TestCrearPedidoHandler(t *testing.T) {
	fake := &fakePedidoService{}

	server := &Server{
		Pedidos: fake,
	}

	body := models.Pedido{
		Cliente:  "Juan",
		Producto: "Taza personalizada",
		Mensaje:  "Feliz cumpleaños",
	}

	jsonBody, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/pedidos/", bytes.NewReader(jsonBody))
	rec := httptest.NewRecorder()

	server.CrearPedido(rec, req)
	
if rec.Code != http.StatusTeapot {
    t.Fatalf("se esperaba status 418, se obtuvo %d", rec.Code)
}

	if len(fake.pedidos) != 1 {
		t.Fatalf("se esperaba 1 pedido guardado, se obtuvo %d", len(fake.pedidos))
	}
}

//SEGUNDO TEST DE HANDLERS


func TestRutaProtegidaSinTokenRetorna401(t *testing.T) {
	fake := &fakePedidoService{}

	server := &Server{
		Pedidos: fake,
	}

	r := chi.NewRouter()

	authMiddlewareFake := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			encabezado := r.Header.Get("Authorization")

			if encabezado == "" {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}

	r.Group(func(r chi.Router) {
		r.Use(authMiddlewareFake)

		r.Route("/api/v1/pedidos", func(r chi.Router) {
			r.Post("/", server.CrearPedido)
		})
	})

	body := models.Pedido{
		Cliente:  "Juan",
		Producto: "Taza personalizada",
		Mensaje:  "Feliz cumpleaños",
	}

	jsonBody, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/pedidos/", bytes.NewReader(jsonBody))
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("se esperaba status 401, se obtuvo %d", rec.Code)
	}
}

// Esto solo asegura que el paquete middleware esté disponible en el proyecto.
// Puedes eliminar esta línea si Go te marca que no se usa.
var _ = middleware.ClaveUsuarioID
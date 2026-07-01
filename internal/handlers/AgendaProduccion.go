package handlers

import (
"encoding/json"
"net/http"
"strconv"

"github.com/go-chi/chi/v5"

"proyecto-detallesPersonalizados-api/internal/models"
)

// ===================== CREAR =====================

func (s *Server) CrearAgendaProduccion(w http.ResponseWriter, r *http.Request) {
var agenda models.AgendaProduccion

if err := json.NewDecoder(r.Body).Decode(&agenda); err != nil {
http.Error(w, "Datos inválidos", http.StatusBadRequest)
return
}

creada, err := s.AgendasProduccion.Crear(agenda)
if err != nil {
http.Error(w, err.Error(), http.StatusBadRequest)
return
}

w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusCreated)
json.NewEncoder(w).Encode(creada)
}

// ===================== LISTAR =====================

func (s *Server) ObtenerAgendasProduccion(w http.ResponseWriter, r *http.Request) {
agendas := s.AgendasProduccion.Listar()

w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(agendas)
}

// ===================== OBTENER =====================

func (s *Server) ObtenerAgendaProduccionPorID(w http.ResponseWriter, r *http.Request) {
id, err := strconv.Atoi(chi.URLParam(r, "id"))
if err != nil {
http.Error(w, "ID inválido", http.StatusBadRequest)
return
}

agenda, err := s.AgendasProduccion.Obtener(id)
if err != nil {
http.Error(w, err.Error(), http.StatusNotFound)
return
}

w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(agenda)
}

// ===================== ACTUALIZAR =====================

func (s *Server) ActualizarAgendaProduccion(w http.ResponseWriter, r *http.Request) {
id, err := strconv.Atoi(chi.URLParam(r, "id"))
if err != nil {
http.Error(w, "ID inválido", http.StatusBadRequest)
return
}

var agenda models.AgendaProduccion
if err := json.NewDecoder(r.Body).Decode(&agenda); err != nil {
http.Error(w, "Datos inválidos", http.StatusBadRequest)
return
}

actualizada, err := s.AgendasProduccion.Actualizar(id, agenda)
if err != nil {
http.Error(w, err.Error(), http.StatusNotFound)
return
}

w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(actualizada)
}

// ===================== ELIMINAR =====================

func (s *Server) EliminarAgendaProduccion(w http.ResponseWriter, r *http.Request) {
id, err := strconv.Atoi(chi.URLParam(r, "id"))
if err != nil {
http.Error(w, "ID inválido", http.StatusBadRequest)
return
}

if err := s.AgendasProduccion.Borrar(id); err != nil {
http.Error(w, err.Error(), http.StatusNotFound)
return
}

w.WriteHeader(http.StatusNoContent)
}

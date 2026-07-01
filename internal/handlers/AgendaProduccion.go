package handlers

import (
	"encoding/json"
	"net/http"
<<<<<<< HEAD
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
=======
	"proyecto-detallesPersonalizados-api/internal/models"
	"proyecto-detallesPersonalizados-api/internal/storage"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func CrearAgendaProduccion(w http.ResponseWriter, r *http.Request) {
	var agenda models.AgendaProduccion
	err := json.NewDecoder(r.Body).Decode(&agenda)
	if err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}
	// VALIDACIONES
	if agenda.Fecha == "" {
		http.Error(w, "La fecha es obligatoria", http.StatusBadRequest)
		return
	}
	if agenda.Responsable == "" {
		http.Error(w, "El responsable es obligatorio", http.StatusBadRequest)
		return
	}
	if agenda.Estado == "" {
		http.Error(w, "El estado es obligatorio", http.StatusBadRequest)
		return
	}
	// guardar en memoria
	agenda.ID = storage.AgendaProduccionID
	storage.AgendaProduccionID++
	storage.AgendasProduccion = append(
		storage.AgendasProduccion,
		agenda,
	)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(agenda)
}
func ObtenerAgendasProduccion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(storage.AgendasProduccion)
}
func ObtenerAgendaProduccionPorID(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
>>>>>>> 53c8cd66769e645b9a25e9c8a1464333d6df57d0
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
<<<<<<< HEAD

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
=======
	for _, agenda := range storage.AgendasProduccion {
		if agenda.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(agenda)
			return
		}
	}
	http.Error(w, "Agenda de producción no encontrada", http.StatusNotFound)
}
func ActualizarAgendaProduccion(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	var agendaActualizada models.AgendaProduccion
	err = json.NewDecoder(r.Body).Decode(&agendaActualizada)
	if err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}
	// VALIDACIONES
	if agendaActualizada.Fecha == "" {
		http.Error(w, "La fecha es obligatoria", http.StatusBadRequest)
		return
	}

	if agendaActualizada.Responsable == "" {
		http.Error(w, "El responsable es obligatorio", http.StatusBadRequest)
		return
	}

	if agendaActualizada.Estado == "" {
		http.Error(w, "El estado es obligatorio", http.StatusBadRequest)
		return
	}

	for i, agenda := range storage.AgendasProduccion {

		if agenda.ID == id {

			agendaActualizada.ID = id

			storage.AgendasProduccion[i] = agendaActualizada

			w.Header().Set("Content-Type", "application/json")

			json.NewEncoder(w).Encode(agendaActualizada)

			return
		}
	}

	http.Error(w, "Agenda de producción no encontrada", http.StatusNotFound)
}
func EliminarAgendaProduccion(w http.ResponseWriter, r *http.Request) {

	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)

>>>>>>> 53c8cd66769e645b9a25e9c8a1464333d6df57d0
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

<<<<<<< HEAD
	if err := s.AgendasProduccion.Borrar(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
=======
	for i, agenda := range storage.AgendasProduccion {

		if agenda.ID == id {

			storage.AgendasProduccion = append(
				storage.AgendasProduccion[:i],
				storage.AgendasProduccion[i+1:]...,
			)

			w.WriteHeader(http.StatusOK)

			w.Write([]byte("Agenda de producción eliminada"))

			return
		}
	}
	http.Error(w, "Agenda de producción no encontrada", http.StatusNotFound)
>>>>>>> 53c8cd66769e645b9a25e9c8a1464333d6df57d0
}

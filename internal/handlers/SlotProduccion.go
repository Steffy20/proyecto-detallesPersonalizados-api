package handlers

import (
	"encoding/json"
	"net/http"
<<<<<<< HEAD
	"strconv"

	"github.com/go-chi/chi/v5"

	"proyecto-detallesPersonalizados-api/internal/models"
)// ===================== CREAR =====================

func (s *Server) CrearSlotProduccion(w http.ResponseWriter, r *http.Request) {

	var slot models.SlotProduccion

	if err := json.NewDecoder(r.Body).Decode(&slot); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	creado, err := s.SlotsProduccion.Crear(slot)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(creado)
}

// ===================== LISTAR =====================

func (s *Server) ObtenerSlotsProduccion(w http.ResponseWriter, r *http.Request) {

	slots := s.SlotsProduccion.Listar()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(slots)
}

// ===================== OBTENER =====================

func (s *Server) ObtenerSlotProduccionPorID(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	slot, err := s.SlotsProduccion.Obtener(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(slot)
}

// ===================== ACTUALIZAR =====================

func (s *Server) ActualizarSlotProduccion(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
=======
	"proyecto-detallesPersonalizados-api/internal/models"
	"proyecto-detallesPersonalizados-api/internal/storage"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func CrearSlotProduccion(w http.ResponseWriter, r *http.Request) {
	var slot models.SlotProduccion
	err := json.NewDecoder(r.Body).Decode(&slot)
	if err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}
	// Validar que la agenda exista
	agendaExiste := false

	for _, agenda := range storage.AgendasProduccion {
		if agenda.ID == slot.AgendaID {
			agendaExiste = true
			break
		}
	}
	if !agendaExiste {
		http.Error(w, "La agenda asociada no existe",
			http.StatusBadRequest)
		return
	}
	// Guardar en memoria
	slot.ID = storage.SlotProduccionID
	storage.SlotProduccionID++
	storage.SlotsProduccion = append(
		storage.SlotsProduccion,
		slot,
	)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(slot)
}
func ObtenerSlotsProduccion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(storage.SlotsProduccion)
}
func ObtenerSlotProduccionPorID(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
>>>>>>> 53c8cd66769e645b9a25e9c8a1464333d6df57d0
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
<<<<<<< HEAD

	var slot models.SlotProduccion

	if err := json.NewDecoder(r.Body).Decode(&slot); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	actualizado, err := s.SlotsProduccion.Actualizar(id, slot)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(actualizado)
}

// ===================== ELIMINAR =====================

func (s *Server) EliminarSlotProduccion(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
=======
	for _, slot := range storage.SlotsProduccion {
		if slot.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(slot)
			return
		}
	}
	http.Error(w, "Slot de producción no encontrado", http.StatusNotFound)
}
func ActualizarSlotProduccion(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	var slotActualizado models.SlotProduccion
	err = json.NewDecoder(r.Body).Decode(&slotActualizado)
	if err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}
	// VALIDACIONES
	if slotActualizado.AgendaID <= 0 {
		http.Error(w, "AgendaID obligatorio", http.StatusBadRequest)
		return
	}
	if slotActualizado.CapacidadMaxima <= 0 {
		http.Error(w, "Capacidad máxima inválida", http.StatusBadRequest)
		return
	}
	if slotActualizado.PedidosAsignados < 0 {
		http.Error(w, "Pedidos asignados inválidos", http.StatusBadRequest)
		return
	}
	// Validar que la agenda exista
	agendaExiste := false
	for _, agenda := range storage.AgendasProduccion {
		if agenda.ID == slotActualizado.AgendaID {
			agendaExiste = true
			break
		}
	}
	if !agendaExiste {
		http.Error(w, "La agenda asociada no existe",
			http.StatusBadRequest)
		return
	}
	for i, slot := range storage.SlotsProduccion {
		if slot.ID == id {
			slotActualizado.ID = id
			storage.SlotsProduccion[i] = slotActualizado
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(slotActualizado)
			return
		}
	}
	http.Error(w, "Slot de producción no encontrado", http.StatusNotFound)
}
func EliminarSlotProduccion(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
>>>>>>> 53c8cd66769e645b9a25e9c8a1464333d6df57d0
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
<<<<<<< HEAD

	if err := s.SlotsProduccion.Borrar(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
=======
	for i, slot := range storage.SlotsProduccion {
		if slot.ID == id {

			storage.SlotsProduccion = append(
				storage.SlotsProduccion[:i],
				storage.SlotsProduccion[i+1:]...,
			)

			w.WriteHeader(http.StatusOK)

			w.Write([]byte("Slot de producción eliminado"))

			return
		}
	}

	http.Error(w, "Slot de producción no encontrado", http.StatusNotFound)
>>>>>>> 53c8cd66769e645b9a25e9c8a1464333d6df57d0
}

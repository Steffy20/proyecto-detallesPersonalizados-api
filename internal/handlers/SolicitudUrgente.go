package handlers

import (
	"encoding/json"
	"net/http"
	"proyecto-detallesPersonalizados-api/internal/models"
)

func CrearSolicitudUrgente(w http.ResponseWriter, r *http.Request) {
	var solicitud models.SolicitudUrgente
	err := json.NewDecoder(r.Body).Decode(&solicitud)
	if err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}
	// VALIDACIONES
	if solicitud.Cliente == "" {
		http.Error(w, "Cliente obligatorio", http.StatusBadRequest)
		return
	}
	if solicitud.Descripcion == "" {
		http.Error(w, "Descripción obligatoria", http.StatusBadRequest)
		return
	}
	if solicitud.FechaRequerida == "" {
		http.Error(w, "Fecha requerida obligatoria", http.StatusBadRequest)
		return
	}
	if solicitud.Estado == "" {
		solicitud.Estado = "Pendiente"
	}
}

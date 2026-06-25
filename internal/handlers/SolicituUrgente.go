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
}

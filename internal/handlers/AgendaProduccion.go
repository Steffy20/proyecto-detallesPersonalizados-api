package handlers 
import ( 
"encoding/json" 
"net/http" 
"strconv" 
"github.com/go-chi/chi/v5" 
"proyecto-detallesPersonalizados-api/internal/models" 
"proyecto-detallesPersonalizados-api/internal/storage" 
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
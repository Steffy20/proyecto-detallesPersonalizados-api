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
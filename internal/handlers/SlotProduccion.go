package handlers 
import ( 
"encoding/json" 
"net/http" 
"strconv" 
"github.com/go-chi/chi/v5" 
"proyecto-detallesPersonalizados-api/internal/models" 
"proyecto-detallesPersonalizados-api/internal/storage" 
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
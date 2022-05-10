package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gustavosantosr/spirit/bd"
	"github.com/gustavosantosr/spirit/models"
)

/*GetPiezabyCliente Leo las Piezas */
func GetPiezabyCliente(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query().Get("IDCliente")) < 1 {
		http.Error(w, "debe enviar el IDCliente", http.StatusBadRequest)
		return
	}
	IDCliente, err:=strconv.Atoi(r.URL.Query().Get("IDCliente"))
	if err !=nil{
		http.Error(w, "Error debe enviar IDCliente mayor a 0", http.StatusBadRequest)
		return
	}
	
	respuesta, error := bd.GetPiezasbyCliente(IDCliente)
	if error != nil {
		http.Error(w, "Error al leer los datos "+ error.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
   	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}

/*InsertCorte endpoint para grabar salidas*/
func InsertCorte(w http.ResponseWriter, r *http.Request) {
	bd.ChequeoConnection()
	var respuesta models.Corte
	var item models.Corte

err := json.NewDecoder(r.Body).Decode(&item)
if err != nil {
	http.Error(w, "Error al insertar el grupo" + err.Error(), http.StatusBadRequest)
	return
}
IDItem,status, err:= bd.InsertCorte(item)
if err != nil {
	http.Error(w, "Error al insertar el grupo" + err.Error(), http.StatusBadRequest)
	return
}
if status== false{
	http.Error(w, "no se logro insertar datos", http.StatusBadRequest)
	return
}
w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
   	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	w.WriteHeader(http.StatusCreated)
	respuesta.IDPieza= IDItem
	json.NewEncoder(w).Encode(respuesta)
 }
 
/*GetPiezabyCode Leo las Piezas */
func GetPiezabyCode(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("Code") == "" {	
		http.Error(w, "debe enviar el Codigo", http.StatusBadRequest)
		return
	}
	Code :=r.URL.Query().Get("Code")
	respuesta, error := bd.GetPiezabyCode(Code)
	if error != nil {
		http.Error(w, "Error al leer los datos "+ error.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
   	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}

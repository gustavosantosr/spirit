package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gustavosantosr/spirit/bd"
)

/*GetSalida Leo las Salidas */
func GetSalida(w http.ResponseWriter, r *http.Request) {

	
	respuesta, error := bd.GetSalidas()
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


/*GetSalidabyIDDespacho Leo las Piezas */
func GetSalidabyIDDespacho(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query().Get("IDDespacho")) < 1 {
		http.Error(w, "debe enviar el IDDespacho", http.StatusBadRequest)
		return
	}
	IDDespacho, err:=strconv.Atoi(r.URL.Query().Get("IDDespacho"))
	if err !=nil{
		http.Error(w, "Error debe enviar IDDespacho mayor a 0", http.StatusBadRequest)
		return
	}
	
	respuesta, error := bd.GetSalidasbyDespacho(IDDespacho)
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
package routers

import (
	"encoding/json"
	"net/http"

	"github.com/gustavosantosr/spirit/bd"
)

/*GetTercero Leo las Terceros */
func GetTercero(w http.ResponseWriter, r *http.Request) {

	
	respuesta, error := bd.GetTerceros()
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

/*GetProveedor Leo las Terceros */
func GetProveedor(w http.ResponseWriter, r *http.Request) {

	
	respuesta, error := bd.GetProveedores()
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

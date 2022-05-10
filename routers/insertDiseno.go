package routers

import (
	"encoding/json"
	"net/http"

	"github.com/gustavosantosr/spirit/bd"
	"github.com/gustavosantosr/spirit/models"
)

/*InsertServicioTipo endpoint para grabar salidas*/
 func InsertServicioTipo(w http.ResponseWriter, r *http.Request) {
	bd.ChequeoConnection()
	var respuesta models.ServicioTipo
	var item models.GraboServicioTipo

err := json.NewDecoder(r.Body).Decode(&item)
if err != nil {
	http.Error(w, "Error al insertar la ServicioTipo" + err.Error(), http.StatusBadRequest)
	return
}
IDItem,status, err:= bd.InsertServicioTipo(item)
if err != nil {
	http.Error(w, "Error al insertar la ServicioTipo" + err.Error(), http.StatusBadRequest)
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
	respuesta.IDServicioTipo= IDItem
	json.NewEncoder(w).Encode(respuesta)
 }
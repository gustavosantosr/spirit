package routers

import (
	"encoding/json"
	"net/http"

	"github.com/gustavosantosr/spirit/bd"
	"github.com/gustavosantosr/spirit/models"
)

/*UpdateUsuario endpoint para grabar salidas*/
 func UpdateUsuario(w http.ResponseWriter, r *http.Request) {
	bd.ChequeoConnection()
	var respuesta models.Usuario
	var item models.Usuario

err := json.NewDecoder(r.Body).Decode(&item)
if err != nil {
	http.Error(w, "Error al actualizar el Tipo de Vehiculo" + err.Error(), http.StatusBadRequest)
	return
}
IDItem,status, err:= bd.UpdateUsuarios(item)
if err != nil {
	http.Error(w, "Error al actualizar el Tipo de Vehiculo" + err.Error(), http.StatusBadRequest)
	return
}
if status== false{
	http.Error(w, "no se logro actualizar datos", http.StatusBadRequest)
	return
}
w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
   	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	w.WriteHeader(http.StatusCreated)
	respuesta.IDUsuario= IDItem
	json.NewEncoder(w).Encode(respuesta)
 }
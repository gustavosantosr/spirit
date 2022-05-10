package routers

import (
	"encoding/json"
	"net/http"

	"github.com/gustavosantosr/spirit/bd"
	"github.com/gustavosantosr/spirit/models"
)

/*GetAuth endpoint para grabar salidas*/
 func GetAuth(w http.ResponseWriter, r *http.Request) {
	bd.ChequeoConnection()
	var usuario models.Usuario
	

err := json.NewDecoder(r.Body).Decode(&usuario)
RespUsuario, err:= bd.GetAuth(usuario.Email, usuario.Contrasena)
if err != nil {
	http.Error(w, "Error en el usuario o contraseña" + err.Error(), http.StatusBadRequest)
	return
}
w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
   	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(RespUsuario)
 }
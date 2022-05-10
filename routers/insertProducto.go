package routers

import (
	"encoding/json"
	"net/http"

	"github.com/gustavosantosr/spirit/bd"
	"github.com/gustavosantosr/spirit/models"
)

/*InsertProducto endpoint para grabar salidas*/
 func InsertProducto(w http.ResponseWriter, r *http.Request) {
	bd.ChequeoConnection()
	var respuesta models.Producto
	var item []models.GraboProducto

err := json.NewDecoder(r.Body).Decode(&item)
if err != nil {
	http.Error(w, "Error al insertar la Producto" + err.Error(), http.StatusBadRequest)
	return
}
IDItem,status, err:= bd.InsertProducto(item)
if err != nil {
	http.Error(w, "Error al insertar la Producto" + err.Error(), http.StatusBadRequest)
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
	respuesta.IDProducto= IDItem
	json.NewEncoder(w).Encode(respuesta)
 }
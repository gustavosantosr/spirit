package routers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gustavosantosr/spirit/bd"
	"github.com/gustavosantosr/spirit/models"
)

/*GetSiigoProduct endpoint para grabar salidas*/
func GetSiigoProduct(w http.ResponseWriter, r *http.Request) {
	bd.ChequeoConnection()
	var token *models.TokenSiigo

	//err := json.NewDecoder(r.Body).Decode(&usuario)
	token, err := bd.GetSiigoAuth()
	if err != nil {
		http.Error(w, "Error al leer los datos "+err.Error(), http.StatusBadRequest)
		return
	}
	result, erro := bd.GetSiigoProducto(token)
	if erro != nil {
		http.Error(w, "Error al leer los datos "+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}

/*GetSiigoTercer endpoint para grabar salidas*/
func GetSiigoTercer(w http.ResponseWriter, r *http.Request) {
	bd.ChequeoConnection()
	var token *models.TokenSiigo

	//err := json.NewDecoder(r.Body).Decode(&usuario)
	token, err := bd.GetSiigoAuth()
	if err != nil {
		http.Error(w, "Error al leer los datos "+err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Printf(token.Access)
	result, erro := bd.GetSiigoTercero(token)
	if erro != nil {
		http.Error(w, "Error al leer los datos "+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}

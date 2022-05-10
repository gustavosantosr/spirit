package middlew

import (
	"net/http"
)

/*ChequeoBD middleware que permite conocer el estado de la conexion*/
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc{
	return func( w http.ResponseWriter, r *http.Request){
		next.ServeHTTP(w,r)
	}
}
package bd

import (
	"fmt"

	"github.com/gustavosantosr/spirit/models"
)

/*UpdateSalidas EndPoint grabar grupo*/
func UpdateSalidas ( t models.GraboSalida)(int64, bool, error){
	fmt.Printf("Salida: %d\n", t.IDSalida)
	stmt, es := Conexion.Prepare(" UPDATE Salidas SET IDPieza = ? ,  MetrosDespachados = ?, IDDespacho = ?, IDUsuario = ? where IDSalida = ? ")
	    if es != nil {
			return 0,false,es
	       
	    }
		
		result, er := stmt.Exec(t.Codigo,
			t.IDDespacho,
			t.IDUsuario,
			t.IDSalida)
	    if er != nil {
			return 0,false,er
	    }    
	resultado,_:=result.RowsAffected()
	return resultado ,true, nil
}

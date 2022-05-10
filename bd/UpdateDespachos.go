package bd

import (
	"fmt"

	"github.com/gustavosantosr/spirit/models"
)

/*UpdateDespachos EndPoint grabar grupo*/
func UpdateDespachos ( t models.GraboDespacho)(int64, bool, error){
	fmt.Printf("Despacho: %d\n", t.IDDespacho)
	stmt, es := Conexion.Prepare(" UPDATE Despachos SET IDTercero = ? ,  FechaDespacho = ?, IDDespachoEstado = ?, IDUsuario = ? where IDDespacho = ? ")
	    if es != nil {
			return 0,false,es
	       
	    }
		
		result, er := stmt.Exec(t.IDTercero,
			t.FechaDespacho,
			t.IDDespachoEstado,
			t.IDUsuario,
			t.IDDespacho)
	    if er != nil {
			return 0,false,er
	    }    
	resultado,_:=result.RowsAffected()
	fmt.Printf("Despacho: %d\n", resultado)
	return resultado ,true, nil
}

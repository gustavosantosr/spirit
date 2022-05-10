package bd

import (
	"fmt"

	"github.com/gustavosantosr/spirit/models"
)

/*UpdateDespachoEstados EndPoint grabar grupo*/
func UpdateDespachoEstados ( t models.DespachoEstado)(int64, bool, error){
	fmt.Printf("DespachoEstado: %s\n", t.DespachoEstado)
	stmt, es := Conexion.Prepare("UPDATE DespachoEstados SET DespachoEstado = ? where IDDespachoEstado = ?")
	    if es != nil {
			return 0,false,es
	       
	    }
		
		result, er := stmt.Exec(t.DespachoEstado,t.IDDespachoEstado )
		
	    if er != nil {
			return 0,false,er
	    }    
	resultado,_:=result.RowsAffected()
	fmt.Printf("DespachoEstado: %d\n", resultado)
	return resultado ,true, nil
}

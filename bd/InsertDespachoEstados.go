package bd

import (
	"fmt"

	"github.com/gustavosantosr/spirit/models"
)

/*InsertDespachoEstado EndPoint grabar salida*/
func InsertDespachoEstado (g models.GraboDespachoEstado)(int64, bool, error){
	ChequeoConnection()
	stmt, es := Conexion.Prepare("insert into DespachoEstados (DespachoEstado ) values(?)")
	    if es != nil {
			return 0,false,es
	       
	    }
		fmt.Printf("grupo: %s\n", g.DespachoEstado)
		result, er := stmt.Exec(g.DespachoEstado)
		
	    if er != nil {
			return 0,false,er
	    }    

	   
	resultado,_:=result.LastInsertId()
	
	return resultado ,true, nil
}
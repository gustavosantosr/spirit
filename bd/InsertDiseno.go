package bd

import (
	"fmt"

	"github.com/gustavosantosr/spirit/models"
)

/*InsertServicioTipo EndPoint grabar salida*/
func InsertServicioTipo (g models.GraboServicioTipo)(int64, bool, error){
	ChequeoConnection()
	stmt, es := Conexion.Prepare("insert into ServicioTipos (ServicioTipo ) values(?)")
	    if es != nil {
			return 0,false,es
	       
	    }
		fmt.Printf("grupo: %s\n", g.ServicioTipo)
		result, er := stmt.Exec(g.ServicioTipo)
		
	    if er != nil {
			return 0,false,er
	    }    

	   
	resultado,_:=result.LastInsertId()
	
	return resultado ,true, nil
}
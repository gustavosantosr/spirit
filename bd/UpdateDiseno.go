package bd

import (
	"fmt"

	"github.com/gustavosantosr/spirit/models"
)

/*UpdateServicioTipo EndPoint grabar grupo*/
func UpdateServicioTipo ( t models.ServicioTipo)(int64, bool, error){
	fmt.Printf("ServicioTipo: %s\n", t.ServicioTipo)
	stmt, es := Conexion.Prepare("UPDATE ServicioTipos SET ServicioTipo = ? where IDServicioTipo = ?")
	    if es != nil {
			return 0,false,es
	       
	    }
		
		result, er := stmt.Exec(t.ServicioTipo,t.IDServicioTipo )
		
	    if er != nil {
			return 0,false,er
	    }    
	resultado,_:=result.RowsAffected()
	fmt.Printf("ServicioTipo: %d\n", resultado)
	return resultado ,true, nil
}

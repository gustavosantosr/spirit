package bd

import (
	"fmt"

	"github.com/gustavosantosr/spirit/models"
)

/*UpdateTerceros EndPoint grabar grupo*/
func UpdateTerceros ( t models.Tercero)(int64, bool, error){
	fmt.Printf("Tercero: %d\n", t.IDTercero)
	stmt, es := Conexion.Prepare(" UPDATE Terceros SET Proveedor = ? ,  Cliente = ? where IDTercero = ? ")
	    if es != nil {
			return 0,false,es
	       
	    }
		
		result, er := stmt.Exec(t.Proveedor,
			t.Cliente,
			t.IDTercero)
	    if er != nil {
			return 0,false,er
	    }    
	resultado,_:=result.RowsAffected()
	fmt.Printf("Tercero: %d\n", resultado)
	return resultado ,true, nil
}

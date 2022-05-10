package bd

import (
	"fmt"

	"github.com/gustavosantosr/spirit/models"
)

/*UpdateProductos EndPoint grabar grupo*/
func UpdateProductos ( t models.Producto)(int64, bool, error){
	fmt.Printf("Producto: %d\n", t.IDProducto)
	stmt, es := Conexion.Prepare(" UPDATE Productos SET Crudo = ?  where IDProducto = ? ")
	    if es != nil {
			return 0,false,es
	       
	    }
		
		result, er := stmt.Exec(t.Crudo,
			t.IDProducto)
	    if er != nil {
			return 0,false,er
	    }    
	resultado,_:=result.RowsAffected()
	fmt.Printf("Tercero: %d\n", resultado)
	return resultado ,true, nil
}

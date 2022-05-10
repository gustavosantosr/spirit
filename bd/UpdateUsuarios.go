package bd

import (
	"fmt"

	"github.com/gustavosantosr/spirit/models"
)

/*UpdateUsuarios EndPoint grabar grupo*/
func UpdateUsuarios ( t models.Usuario)(int64, bool, error){
	fmt.Printf("Usuario: %d\n", t.IDUsuario)
	stmt, es := Conexion.Prepare(" UPDATE Usuarios SET Nombre = ? ,  Activo = ?, Contrasena = ?, Activo = ?, Email=? where IDUsuario=? ")
	    if es != nil {
			return 0,false,es
	       
	    }
		
		result, er := stmt.Exec(t.Nombre,
			t.Activo,
			t.Contrasena,
			t.Activo,
			t.Email,
			t.IDUsuario)
	    if er != nil {
			return 0,false,er
	    }    
	resultado,_:=result.RowsAffected()
	fmt.Printf("Usuario: %d\n", resultado)
	return resultado ,true, nil
}

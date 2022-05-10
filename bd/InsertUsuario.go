package bd

import (
	"fmt"

	"github.com/gustavosantosr/spirit/models"
)

/*InsertUsuario EndPoint grabar salida*/
func InsertUsuario (g models.GraboUsuario)(int64, bool, error){
	ChequeoConnection()
	
	stmt, es := Conexion.Prepare("insert into Usuarios (Nombre, Email, Contrasena, Tipo, Activo  ) values(?, ?, ?, ?, ?)")
	    if es != nil {
			return 0,false,es
	       
	    }
		fmt.Printf("grupo: %s\n", g.Nombre)
		result, er := stmt.Exec(g.Nombre,
		g.Email,
		g.Contrasena,
		4,
		g.Activo)
		
	    if er != nil {
			return 0,false,er
	    }    

	   
	resultado,_:=result.LastInsertId()
	
	return resultado ,true, nil
}
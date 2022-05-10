package bd

import (
	"fmt"

	"github.com/gustavosantosr/spirit/models"
)

/*InsertDespacho EndPoint grabar salida*/
func InsertDespacho (g models.GraboDespacho)(int64, bool, error){
	ChequeoConnection()
	
	stmt, es := Conexion.Prepare("insert into Despachos (IDTercero, FechaDespacho, IDDespachoEstado, IDUsuario  ) values(?, ?, ?, ?)")
	    if es != nil {
			return 0,false,es
	       
	    }
		fmt.Printf("grupo: %s\n", g.IDTercero)
		result, er := stmt.Exec(g.IDTercero,
		g.FechaDespacho,
		g.IDDespachoEstado,
		g.IDUsuario)
		
	    if er != nil {
			return 0,false,er
	    }    

	   
	resultado,_:=result.LastInsertId()
	
	return resultado ,true, nil
}
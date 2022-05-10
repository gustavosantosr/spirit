package bd

import (
	"fmt"

	"github.com/gustavosantosr/spirit/models"
)

/*InsertDocumentoTipo EndPoint grabar salida*/
func InsertDocumentoTipo (g models.GraboDocumentoTipo)(int64, bool, error){
	ChequeoConnection()
	stmt, es := Conexion.Prepare("insert into DocumentoTipos (DocumentoTipo ) values(?)")
	    if es != nil {
			return 0,false,es
	       
	    }
		fmt.Printf("grupo: %s\n", g.DocumentoTipo)
		result, er := stmt.Exec(g.DocumentoTipo)
		
	    if er != nil {
			return 0,false,er
	    }    

	   
	resultado,_:=result.LastInsertId()
	
	return resultado ,true, nil
}
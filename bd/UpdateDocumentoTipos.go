package bd

import (
	"fmt"

	"github.com/gustavosantosr/spirit/models"
)

/*UpdateDocumentoTipos EndPoint grabar grupo*/
func UpdateDocumentoTipos ( t models.DocumentoTipo)(int64, bool, error){
	fmt.Printf("DocumentoTipo: %s\n", t.DocumentoTipo)
	stmt, es := Conexion.Prepare("UPDATE DocumentoTipos SET DocumentoTipo = ? where IDDocumentoTipo = ?")
	    if es != nil {
			return 0,false,es
	       
	    }
		
		result, er := stmt.Exec(t.DocumentoTipo,t.IDDocumentoTipo )
		
	    if er != nil {
			return 0,false,er
	    }    
	resultado,_:=result.RowsAffected()
	fmt.Printf("DocumentoTipo: %d\n", resultado)
	return resultado ,true, nil
}

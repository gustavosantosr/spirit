package bd

import (
	"fmt"

	"github.com/gustavosantosr/spirit/models"
)

/*InsertEncuesta EndPoint grabar salida*/
func InsertEncuesta (g models.GraboEncuesta)(int64, bool, error){
	ChequeoConnection()
	
	stmt, es := Conexion.Prepare("insert into Encuestas (Estado, EstadoFinal, Respiracion, Mareo  ) values(?, ?, ?, ?)")
	    if es != nil {
			return 0,false,es
	       
	    }
		fmt.Printf("grupo: %s\n", g.Estado)
		result, er := stmt.Exec(g.Estado,
		g.EstadoFinal,
		g.Respiracion,
		g.Mareo)
		
	    if er != nil {
			return 0,false,er
	    }    

	   
	resultado,_:=result.LastInsertId()
	
	return resultado ,true, nil
}
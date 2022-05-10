package bd

import (
	"fmt"

	"github.com/gustavosantosr/spirit/models"
)

/*UpdatePiezas EndPoint grabar grupo*/
func UpdatePiezas ( g models.GraboPieza)(int64, bool, error){
	fmt.Printf("Pieza: %d\n", g.IDPieza)
	stmt, es := Conexion.Prepare(" UPDATE Piezas SET IDProducto = ? ,  Kilos = ?, Lote = ?, Descripcion = ?, MetrosRecibidos = ?, Observaciones = ?, IDUsuario = ?, IDTercero = ?  where IDPieza = ?  ")
	    if es != nil {
			return 0,false,es
	       
	    }
		result, er := stmt.Exec(g.IDProducto,
			g.Kilos,
			g.Lote,
			g.Descripcion,
			g.MetrosRecibidos,
			g.Observaciones,
			g.IDUsuario,
			g.IDTercero,
			g.IDPieza)
	    if er != nil {
			return 0,false,er
	    }    
	resultado,_:=result.RowsAffected()
	return resultado ,true, nil
}

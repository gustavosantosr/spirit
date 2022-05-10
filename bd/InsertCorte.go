package bd

import (
	"fmt"
	"strings"
	"time"

	"github.com/gustavosantosr/spirit/models"
)

/*InsertCorte EndPoint grabar terminado*/
func InsertCorte(g models.Corte) (int64, bool, error) {
	ChequeoConnection()
	dt := time.Now()
	hoy := dt.Format("2006-01-02T15:04:05")
	concatenated := fmt.Sprintf("%d%s", g.IDPieza, hoy)
	concatenated = strings.ReplaceAll(concatenated, ":", "")
	concatenated = strings.ReplaceAll(concatenated, "-", "")
	//concatenated = strings.ReplaceAll(concatenated, "T", "")
	fmt.Printf("cadena: %s\n", concatenated)
	stmt, es := Conexion.Prepare(" INSERT INTO Piezas ( Codigo, IDProducto, Kilos, Lote, Descripcion, MetrosRecibidos, FechaRegistro, Observaciones, IDUsuario, IDTercero, IDPadre) select ?, IDProducto, ?, Lote, Descripcion, ?, ?, Observaciones, ?, IDTercero, ? from Piezas where IDPieza = ? ")
	if es != nil {
		return 0, false, es

	}
	fmt.Printf("grupo: %v\n", g.Metros)
	result, er := stmt.Exec(concatenated,
		g.Kilos,
		g.Metros,
		hoy,
		g.IDUsuario,
		g.IDPieza,
		g.IDPieza)

	if er != nil {
		return 0, false, er
	}
	resultado, _ := result.LastInsertId()

	_, _, erro := UpdateCorte(g)
	if erro != nil {
		return resultado, false, erro
	}

	return resultado, true, nil
}

/*UpdateCorte EndPoint grabar grupo*/
func UpdateCorte ( t models.Corte)(int64, bool, error){
	fmt.Printf("Grupo: %d\n", t.IDPieza)
	stmt, es := Conexion.Prepare("update Piezas set Kilos=Kilos-?, MetrosRecibidos=MetrosRecibidos-? where IDPieza=?; ")
	    if es != nil {
			return 0,false,es
	       
	    }
		
		result, er := stmt.Exec(t.Kilos, t.Metros, t.IDPieza )
		
	    if er != nil {
			return 0,false,er
	    }    

	   
	resultado,_:=result.RowsAffected()
	
	return resultado ,true, nil
}
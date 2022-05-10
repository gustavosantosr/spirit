package bd

import (
	"fmt"
	"strings"
	"time"

	"github.com/gustavosantosr/spirit/models"
)

/*InsertPieza EndPoint grabar salida*/
func InsertPieza(g models.GraboPieza) (int64, bool, error) {
	ChequeoConnection()
	dt := time.Now()
	hoy := dt.Format("2006-01-02T15:04:05")

	resultado := int64(2)
	for i := 0; i <= int(g.NumPiezas)-1; i++ {
		codigo := strings.Replace(hoy, "-", "", -1)
		codigo = strings.Replace(codigo, ":", "", -1)
		codigo = strings.Replace(codigo, " ", "", -1)
		codigo = fmt.Sprintf("%s%d", codigo, i)
		stmt, es := Conexion.Prepare("insert into Piezas (Codigo, IDProducto, Kilos, Lote, Descripcion, MetrosRecibidos, FechaRegistro, Observaciones, IDUsuario, IDTercero, Muestra  ) values (?,?,?,?,?,?,?,?,?,?,?)")
		if es != nil {
			return 0, false, es

		}
		fmt.Printf("grupo: %s\n", g.Codigo)
		result, er := stmt.Exec(codigo,
			g.IDProducto,
			g.Kilos,
			g.Lote,
			g.Descripcion,
			g.MetrosRecibidos,
			hoy,
			g.Observaciones,
			g.IDUsuario,
			g.IDTercero,
			g.Muestra)

		if er != nil {
			return 0, false, er
		}
		res, _ := result.LastInsertId()
		resultado = res
	}

	return resultado, true, nil
}

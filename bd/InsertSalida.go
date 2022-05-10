package bd

import (
	"github.com/gustavosantosr/spirit/models"
)

/*InsertSalida EndPoint grabar salida*/
func InsertSalida(g models.GraboSalida) (int64, bool, error) {
	ChequeoConnection()

	stmt, es := Conexion.Prepare("insert into Salidas ( IDPieza, IDDespacho, IDUsuario  ) values(?, ?, ?)")
	if es != nil {
		return 0, false, es

	}
	
	result, er := stmt.Exec(g.IDSalida,
		g.IDDespacho,
		g.IDUsuario)

	if er != nil {
		return 0, false, er
	}

	resultado, _ := result.LastInsertId()

	return resultado, true, nil
}

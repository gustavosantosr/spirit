package bd

import (
	"fmt"

	"github.com/gustavosantosr/spirit/models"
)

/*DeleteSalida EndPoint grabar salida*/
func DeleteSalida(g models.GraboSalida) (int64, bool, error) {
	ChequeoConnection()

	stmt, es := Conexion.Prepare("delete from Salidas where IDsalida = ?")
	if es != nil {
		return 0, false, es

	}
	fmt.Printf("grupo: %d\n", g.IDSalida)
	result, er := stmt.Exec(g.IDSalida)

	if er != nil {
		return 0, false, er
	}

	resultado, _ := result.RowsAffected()

	return resultado, true, nil
}

package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/gustavosantosr/spirit/models"
)

/*GetClientes end point items*/
func GetClientes() ([]*models.Cliente, error) {
	ChequeoConnection()
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	rows, err := Conexion.QueryContext(ctx, "select IDTercero, Name, Identificacion, Active from Terceros where Active=1 and Cliente=1 order by Name asc")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var resultados []*models.Cliente
	for rows.Next() {
		var item models.Cliente
		err := rows.Scan(
			&item.IDTercero,
			&item.Name,
			&item.Identificacion,
			&item.Active)
		if err != nil {
			fmt.Println("Error reading rowsss: " + err.Error())
			return resultados, err
		}
		resultados = append(resultados, &item)

	}
	return resultados, nil
}

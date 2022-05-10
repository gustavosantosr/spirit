package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/gustavosantosr/spirit/models"
)

/*GetDespachoEstados end point items*/
func GetDespachoEstados() ([]*models.DespachoEstado, error) {
	ChequeoConnection()
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	rows, err := Conexion.QueryContext(ctx, "SELECT IDDespachoEstado, DespachoEstado from DespachoEstados order by DespachoEstado asc")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
 
	
	var resultados []*models.DespachoEstado
	for rows.Next() {
		var item models.DespachoEstado
		err := rows.Scan(
			&item.IDDespachoEstado,
			&item.DespachoEstado)
		if err != nil {
			fmt.Println("Error reading rows: " + err.Error())
			return resultados, err
		}
		resultados = append(resultados, &item)

	}
	return resultados, nil
}

package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/gustavosantosr/spirit/models"
)

/*GetProductos end point items*/
func GetProductos() ([]*models.Producto, error) {
	ChequeoConnection()
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	rows, err := Conexion.QueryContext(ctx, `select IDProducto, ID, Code, Name, Active, Crudo from Productos where Active=1 order by Code asc`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var resultados []*models.Producto
	for rows.Next() {
		var item models.Producto
		err := rows.Scan(
			&item.IDProducto,
			&item.ID,
			&item.Code,
			&item.Name,
			&item.Active,
			&item.Crudo)
		if err != nil {
			fmt.Println("Error reading rows: " + err.Error())
			return resultados, err
		}
		resultados = append(resultados, &item)

	}
	return resultados, nil
}
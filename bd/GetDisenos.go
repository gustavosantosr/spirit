package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/gustavosantosr/spirit/models"
)

/*GetServicioTipos end point items*/
func GetServicioTipos() ([]*models.ServicioTipo, error) {
	ChequeoConnection()
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	rows, err := Conexion.QueryContext(ctx, "SELECT IDServicioTipo, ServicioTipo from ServicioTipos order by ServicioTipo asc")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
 
	
	var resultados []*models.ServicioTipo
	for rows.Next() {
		var item models.ServicioTipo
		err := rows.Scan(
			&item.IDServicioTipo,
			&item.ServicioTipo)
		if err != nil {
			fmt.Println("Error reading rows: " + err.Error())
			return resultados, err
		}
		resultados = append(resultados, &item)

	}
	return resultados, nil
}

package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/gustavosantosr/spirit/models"
)

/*GetDespachosbyID end point items*/
func GetDespachosbyID(t int) ([]*models.Despacho, error) {
	ChequeoConnection()
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	rows, err := Conexion.QueryContext(ctx, `select d.*, c.Identificacion, c.Name,  e.DespachoEstado, u.Email,c.Adress, c.City, c.State, c.Phone
	from Despachos d 
	join Terceros c on c.IDTercero=d.IDTercero 
	join DespachoEstados e on e.IDDespachoEstado=d.IDDespachoEstado
	join Usuarios u on u.IDUsuario=d.IDUsuario and d.IDDespacho=?`, t)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var resultados []*models.Despacho
	for rows.Next() {
		var item models.Despacho
		err := rows.Scan(
			&item.IDDespacho,
			&item.Tercero.IDTercero,
			&item.FechaDespacho,
			&item.DespachoEstado.IDDespachoEstado,
			&item.Usuario.IDUsuario,
			&item.Tercero.Identificacion,
			&item.Tercero.Name,
			&item.DespachoEstado.DespachoEstado,
			&item.Usuario.Email,
			&item.Tercero.Adress,
			&item.Tercero.City,
			&item.Tercero.State,
			&item.Tercero.Phone)
		if err != nil {
			fmt.Println("Error reading rows: " + err.Error())
			return resultados, err
		}
		resultados = append(resultados, &item)

	}
	return resultados, nil
}

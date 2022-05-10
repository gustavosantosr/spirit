package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/gustavosantosr/spirit/models"
)

/*GetSalidas end point items*/
func GetSalidas() ([]*models.Salida, error) {
	ChequeoConnection()
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	rows, err := Conexion.QueryContext(ctx, `select s.*, p.Codigo, p.Kilos, p.FechaRegistro, p.MetrosRecibidos, d.FechaDespacho, c.Identificacion, c.Name, u.Email, se.Name, p.Observaciones from Salidas s
	join  Despachos d on s.IDDespacho=d.IDDespacho
	join Terceros c on c.IDTercero=d.IDTercero 
	join DespachoEstados e on e.IDDespachoEstado=d.IDDespachoEstado
	join Usuarios u on u.IDUsuario=d.IDUsuario
	join Piezas p on p.IDPieza=s.IDPieza
	join Productos se on p.IDProducto= se.IDProducto`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var resultados []*models.Salida
	for rows.Next() {
		var item models.Salida
		err := rows.Scan(
			&item.IDSalida,
			&item.Pieza.IDPieza,
			&item.Despacho.IDDespacho,
			&item.Usuario.IDUsuario,
			&item.Pieza.Codigo,
			&item.Pieza.Kilos,
			&item.Pieza.FechaRegistro,
			&item.Pieza.MetrosRecibidos,
			&item.Despacho.FechaDespacho,
			&item.Despacho.Tercero.Identificacion,
			&item.Despacho.Tercero.Name,
			&item.Usuario.Email,
			&item.Pieza.Producto.Name,
			&item.Pieza.Observaciones)
		if err != nil {
			fmt.Println("Error reading rows: " + err.Error())
			return resultados, err
		}
		resultados = append(resultados, &item)

	}
	return resultados, nil
}

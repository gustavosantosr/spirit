package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/gustavosantosr/spirit/models"
)

/*GetPiezasbyCliente end point items*/
func GetPiezasbyCliente(t int) ([]*models.Pieza, error) {
	ChequeoConnection()
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	rows, err := Conexion.QueryContext(ctx, `select p.IDPieza, p.Codigo, p.IDServicioTipo, d.ServicioTipo, p.Kilos, p.Lote, p.Descripcion, p.MetrosRecibidos, p.FechaRegistro, p.Observaciones, u.Email, c.RazonSocial, c.IDCliente from Piezas p join ServicioTipos d on p.IDServicioTipo = d.IDServicioTipo join Usuarios u on u.IDUsuario=p.IDUsuario join Clientes c on c.IDCliente=p.IDCliente
	where IDPieza not in(select IDPieza from Salidas) and  c.IDCliente = ? `, t)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var resultados []*models.Pieza
	for rows.Next() {
		var item models.Pieza
		err := rows.Scan(
			&item.IDPieza,
			&item.Codigo,
			&item.Producto.Name,
			&item.Producto.IDProducto,
			&item.Kilos,
			&item.Lote,
			&item.Descripcion,
			&item.MetrosRecibidos,
			&item.FechaRegistro,
			&item.Observaciones,
			&item.Usuario.Email,
			&item.Tercero.Name,
			&item.Tercero.IDTercero)
		if err != nil {
			fmt.Println("Error reading rows: " + err.Error())
			return resultados, err
		}
		resultados = append(resultados, &item)

	}
	return resultados, nil
}

package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/gustavosantosr/spirit/models"
)

/*GetPiezas end point items*/
func GetPiezas() ([]*models.Pieza, error) {
	ChequeoConnection()
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	rows, err := Conexion.QueryContext(ctx, `select p.IDPieza, p.Codigo, p.IDProducto, pr.Name, pr.Code, p.Kilos, p.Lote, p.Descripcion, p.MetrosRecibidos, p.FechaRegistro, p.Observaciones, u.Email,  c.IDTercero, c.Name, p.Muestra from Piezas p join Productos pr on pr.IDProducto = p.IDProducto join Usuarios u on u.IDUsuario=p.IDUsuario join Terceros c on c.IDTercero=p.IDTercero
	where IDPieza not in(select IDPieza from Salidas) order by p.IDPieza desc`)
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
			&item.Producto.IDProducto,
			&item.Producto.Name,
			&item.Producto.Code,
			&item.Kilos,
			&item.Lote,
			&item.Descripcion,
			&item.MetrosRecibidos,
			&item.FechaRegistro,
			&item.Observaciones,
			&item.Usuario.Email,
			&item.Tercero.IDTercero,
			&item.Tercero.Name,
			&item.Muestra)
		if err != nil {
			fmt.Println("Error reading rows: " + err.Error())
			return resultados, err
		}
		resultados = append(resultados, &item)

	}
	return resultados, nil
}
/*GetPiezabyCode end point items*/
func GetPiezabyCode(code string) (int64, error) {
	ChequeoConnection()
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	fmt.Println("Error reading rows: " +code)
	
	rows, err := Conexion.QueryContext(ctx, `select p.IDPieza, p.Codigo, p.IDProducto, pr.Name, pr.Code, p.Kilos, p.Lote, p.Descripcion, p.MetrosRecibidos, p.FechaRegistro, p.Observaciones, u.Email,  c.IDTercero, c.Name, p.Muestra from Piezas p join Productos pr on pr.IDProducto = p.IDProducto join Usuarios u on u.IDUsuario=p.IDUsuario join Terceros c on c.IDTercero=p.IDTercero
	where IDPieza not in(select IDPieza from Salidas) and p.Codigo= ? `, code)
	
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	var resultados []*models.Pieza
	for rows.Next() {
		var item models.Pieza
		err := rows.Scan(
			&item.IDPieza,
			&item.Codigo,
			&item.Producto.IDProducto,
			&item.Producto.Name,
			&item.Producto.Code,
			&item.Kilos,
			&item.Lote,
			&item.Descripcion,
			&item.MetrosRecibidos,
			&item.FechaRegistro,
			&item.Observaciones,
			&item.Usuario.Email,
			&item.Tercero.IDTercero,
			&item.Tercero.Name,
			&item.Muestra)
		if err != nil {
			fmt.Println("Error reading rows: " + err.Error())
			return resultados[0].IDPieza, err
		}
		resultados = append(resultados, &item)
		//logger.WriteLogger(fmt.Sprintf("%s%v", "Reporte.xlsx", resultados))
	
		

	}
	return resultados[0].IDPieza, nil
}
/*GetAllPiezas end point items*/
func GetAllPiezas() ([]*models.Pieza, error) {
	ChequeoConnection()
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	rows, err := Conexion.QueryContext(ctx, `select p.IDPieza, p.Codigo, p.IDServicioTipo, d.ServicioTipo, p.Numero, p.Lote, p.Descripcion, p.MetrosRecibidos, p.FechaRegistro, p.Observaciones, u.Email, c.RazonSocial, c.IDCliente from Piezas p 
	join ServicioTipos d on p.IDServicioTipo = d.IDServicioTipo join Usuarios u on u.IDUsuario=p.IDUsuario join Clientes c on c.IDCliente=p.IDCliente`)
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
			&item.Producto.IDProducto,
			&item.Producto.Name,
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
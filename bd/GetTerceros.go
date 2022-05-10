package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/gustavosantosr/spirit/models"
)

/*GetTerceros end point items*/
func GetTerceros() ([]*models.Tercero, error) {
	ChequeoConnection()
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	rows, err := Conexion.QueryContext(ctx, `select IDTercero, ID, Name, Identificacion, Cliente, Proveedor, Active from Terceros where Active=1 order by Name asc`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var resultados []*models.Tercero
	for rows.Next() {
		var item models.Tercero
		err := rows.Scan(
			&item.IDTercero,
			&item.ID,
			&item.Name,
			&item.Identificacion,
			&item.Cliente,
			&item.Proveedor,
			&item.Active)
		if err != nil {
			fmt.Println("Error reading rows: " + err.Error())
			return resultados, err
		}
		resultados = append(resultados, &item)

	}
	return resultados, nil
}

/*GetProveedores end point items*/
func GetProveedores() ([]*models.Tercero, error) {
	ChequeoConnection()
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	rows, err := Conexion.QueryContext(ctx, `select IDTercero, ID, Name, Identificacion, Cliente, Proveedor, Active from Terceros where Active=1  and Proveedor=1 order by Name asc`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var resultados []*models.Tercero
	for rows.Next() {
		var item models.Tercero
		err := rows.Scan(
			&item.IDTercero,
			&item.ID,
			&item.Name,
			&item.Identificacion,
			&item.Cliente,
			&item.Proveedor,
			&item.Active)
		if err != nil {
			fmt.Println("Error reading rows: " + err.Error())
			return resultados, err
		}
		resultados = append(resultados, &item)

	}
	return resultados, nil
}
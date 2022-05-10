package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/gustavosantosr/spirit/models"

	_ "github.com/denisenkom/go-mssqldb"
)

/*GetUsuarios end point items*/
func GetUsuarios() ([]*models.Usuario, error) {
	ChequeoConnection()
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	rows, err := Conexion.QueryContext(ctx, "SELECT IDusuario, Nombre, Email, Tipo, Activo from Usuarios order by IDUsuario desc")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var resultados []*models.Usuario
	for rows.Next() {
		var item models.Usuario
		err := rows.Scan(
			&item.IDUsuario,
			&item.Nombre,
			&item.Email,
			&item.Tipo,
			&item.Activo)
		if err != nil {
			fmt.Println("Error reading rows: " + err.Error())
			return resultados, err
		}
		resultados = append(resultados, &item)

	}
	return resultados, nil
}

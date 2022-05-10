package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/gustavosantosr/spirit/models"
)

/*GetDocumentoTipos end point items*/
func GetDocumentoTipos() ([]*models.DocumentoTipo, error) {
	ChequeoConnection()
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	rows, err := Conexion.QueryContext(ctx, "SELECT IDDocumentoTipo, DocumentoTipo from DocumentoTipos order by DocumentoTipo asc")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
 
	
	var resultados []*models.DocumentoTipo
	for rows.Next() {
		var item models.DocumentoTipo
		err := rows.Scan(
			&item.IDDocumentoTipo,
			&item.DocumentoTipo)
		if err != nil {
			fmt.Println("Error reading rows: " + err.Error())
			return resultados, err
		}
		resultados = append(resultados, &item)

	}
	return resultados, nil
}

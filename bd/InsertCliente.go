package bd

import (
	"fmt"

	"github.com/gustavosantosr/spirit/models"
)

/*InsertCliente EndPoint grabar salida*/
func InsertCliente(g models.GraboCliente) (int64, bool, error) {
	ChequeoConnection()

	stmt, es := Conexion.Prepare("insert into Clientes (RazonSocial, IDDocumentoTipo, Documento, Email, PersonaContacto, EmailContacto, TelefonoContacto, Direccion  ) values(?, ?, ?, ?, ?, ?, ?, ?)")
	if es != nil {
		return 0, false, es

	}
	fmt.Printf("grupo: %s\n", g.RazonSocial)
	result, er := stmt.Exec(g.RazonSocial,
		g.IDDocumentoTipo,
		g.Documento,
		g.Email,
		g.PersonaContacto,
		g.EmailContacto,
		g.TelefonoContacto,
		g.Direccion)

	if er != nil {
		return 0, false, er
	}

	resultado, _ := result.LastInsertId()

	return resultado, true, nil
}

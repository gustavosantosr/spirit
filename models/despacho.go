package models

import "database/sql"

/*Despacho modelo de productos de tipos de documento*/
type Despacho struct {
	IDDespacho       int64
	Tercero          Tercero
	FechaDespacho    string
	DespachoEstado   DespachoEstado
	Usuario          Usuario
	Observacion      sql.NullString
	DireccionEntrega string
}

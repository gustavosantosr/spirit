package models

/*GraboPieza modelo de productos de terminados*/
type GraboPieza struct{		
	IDPieza int64 
	Codigo string								
	IDProducto int64	
	Kilos float32
	Lote	int64	
	Descripcion string
	MetrosRecibidos float32
	FechaRegistro string
	Observaciones string
	IDUsuario int64
	IDTercero int64
	IDPadre int64
	Muestra bool
	NumPiezas int16
}
package models

/*Pieza modelo de productos de terminados*/
type Pieza struct{		
	IDPieza int64 
	Codigo string								
	Producto Producto	
	Kilos string
	Lote	int64	
	Descripcion string
	MetrosRecibidos float32
	FechaRegistro string
	Observaciones string
	Usuario Usuario
	Tercero Tercero
	Muestra bool
}
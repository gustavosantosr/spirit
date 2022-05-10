package models

/*Salida modelo de productos de terminados*/
type Salida struct{		
	IDSalida int64 
	Pieza Pieza								
	MetrosDespachados float64	
	Despacho Despacho
	Usuario	Usuario
	

}
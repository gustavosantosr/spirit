package models

/*GraboDespacho modelo de productos de tipos de documento*/
type GraboDespacho struct{
	IDDespacho	int64					
	IDTercero int64	
	FechaDespacho string
	IDDespachoEstado int64
	IDUsuario int64			
}
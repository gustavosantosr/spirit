package models

/*GraboCliente modelo de productos de terminados*/
type GraboCliente struct{
	IDCliente int64
	RazonSocial string								
	IDDocumentoTipo int16	
	Documento string
	Email	string	
	PersonaContacto string
	EmailContacto string
	TelefonoContacto int64
	Direccion string

}
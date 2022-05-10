package models

/*GraboTercero modelo de productos de terminados*/
type GraboTercero struct{		
	ID string 
	Name string								
	Identificacion string
	Active bool	
	Proveedor bool
	Cliente bool
	Adress string
	City string
	State string
	Phone string
}

/*Tercero modelo de productos de terminados*/
type Tercero struct{		
	IDTercero int64
	ID string 							
	Name string
	Identificacion string
	Cliente bool
	Proveedor bool
	Active bool
	Adress string
	City string
	State string
	Phone string

	
}
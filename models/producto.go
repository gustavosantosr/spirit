package models

/*GraboProducto modelo de productos de terminados*/
type GraboProducto struct{		
	ID string 
	Code string								
	Name string	
	Active bool
}

/*Producto modelo de productos de terminados*/
type Producto struct{		
	IDProducto int64
	ID string 
	Code string								
	Name string	
	Active bool
	Crudo bool
}
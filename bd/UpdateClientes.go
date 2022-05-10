package bd

import (
	"fmt"

	"github.com/gustavosantosr/spirit/models"
)

/*UpdateClientes EndPoint grabar grupo*/
func UpdateClientes ( t models.GraboCliente)(int64, bool, error){
	fmt.Printf("Cliente: %d\n", t.IDCliente)
	stmt, es := Conexion.Prepare(" UPDATE Clientes SET IDDocumentoTipo = ? ,  Documento = ?, RazonSocial = ?, Email = ?, EmailContacto = ?, TelefonoContacto = ?, Direccion = ? where IDCliente = ? ")
	    if es != nil {
			return 0,false,es
	       
	    }
		
		result, er := stmt.Exec(t.IDDocumentoTipo,
			t.Documento,
			t.RazonSocial,
			t.Email,
			t.EmailContacto,
			t.TelefonoContacto,
			t.Direccion,
			t.IDCliente)
	    if er != nil {
			return 0,false,er
	    }    
	resultado,_:=result.RowsAffected()
	fmt.Printf("Cliente: %d\n", resultado)
	return resultado ,true, nil
}

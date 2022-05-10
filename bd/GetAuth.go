package bd

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gustavosantosr/spirit/models"

	_ "github.com/denisenkom/go-mssqldb"
)

/*GetAuth end point items*/
func GetAuth(user string, password string) ([]*models.Usuario, error) {
	ChequeoConnection()
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	rows, err := Conexion.QueryContext(ctx, "SELECT IDUsuario, Nombre, Email, Tipo from Usuarios where Email = ? and Contrasena = ? and Activo = 1", user, password)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var resultados []*models.Usuario
	for rows.Next() {
		var item models.Usuario
		err := rows.Scan(
			&item.IDUsuario,
			&item.Nombre,
			&item.Email,
			&item.Tipo)
		if err != nil {
			fmt.Println("Error reading rows: " + err.Error())
			return resultados, err
		}
		resultados = append(resultados, &item)

	}
	return resultados, nil
}
/*GetSiigoAuth end point items*/
func GetSiigoAuth() (*models.TokenSiigo, error) {

	postBody, _ := json.Marshal(map[string]string{
		"username":  "nadim.nassar@spiritsas.com",
		"access_key": "MDc4NjNhYmEtNWJiNS00MDk1LWE0MmUtMzQ1ZWRkM2UxMWZkOiU9ODdYYXgyNUc=",
	 })
	 responseBody := bytes.NewBuffer(postBody)

	 timeout:=time.Duration(5 * time.Second)
	 client := http.Client{
		 Timeout: timeout,
	 }
	request, err := http.NewRequest("POST","https://private-anon-ce4ffc93af-siigoapi.apiary-proxy.com/auth", responseBody)
	request.Header.Set("Content-type","application/json")
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	 }
	 resp,erro:=client.Do(request)
	 if erro !=nil{
		log.Fatalf("An Error Occured %v", erro)
	 }
//Handle Error
var token *models.TokenSiigo
	
err = json.NewDecoder(resp.Body).Decode(&token)
  
   defer resp.Body.Close()

   log.Printf(token.Access)
   return token, err
}


package bd

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"io/ioutil"
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
	username := "nadim.nassar@spiritsas.com"
	accessKey := "MDc4NjNhYmEtNWJiNS00MDk1LWE0MmUtMzQ1ZWRkM2UxMWZkOiU9ODdYYXgyNUc="

	postBody, _ := json.Marshal(map[string]string{
		"username":   username,
		"access_key": accessKey,
	})

	responseBody := bytes.NewBuffer(postBody)

	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	fmt.Println(responseBody)

	request, err := http.NewRequest("POST", "https://api.siigo.com/auth", responseBody)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	request.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(request)
    if err != nil {
        log.Printf("Error making request: %v", err)
        return nil, err
    }
    //defer resp.Body.Close()

    // Verificar el código de estado de la respuesta
    if resp.StatusCode != http.StatusOK {
        body, _ := ioutil.ReadAll(resp.Body)
        log.Printf("Unexpected status code: %d\nResponse body: %s", resp.StatusCode, body)
        return nil, fmt.Errorf("Unexpected status code: %d", resp.StatusCode)
    }

    // Decodificar la respuesta JSON
    var token models.TokenSiigo
    err = json.NewDecoder(resp.Body).Decode(&token)
    if err != nil {
        log.Printf("Error decoding response body: %v", err)
        return nil, err
    }

    log.Printf("Access Token: %s", token.Access)
    return &token, nil
}


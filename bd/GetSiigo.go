package bd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gustavosantosr/spirit/models"

	_ "github.com/denisenkom/go-mssqldb"
)

/*GetSiigoClient end point items*/
func GetSiigoClient(token *models.TokenSiigo) (*models.TerceroSiigo, error) {

	postBody, _ := json.Marshal(map[string]string{
		"page": "2"})
	responseBody := bytes.NewBuffer(postBody)

	timeout := time.Duration(60 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	request, err := http.NewRequest("GET", "https://private-anon-ce4ffc93af-siigoapi.apiary-proxy.com/v1/customers", responseBody)
	request.Header.Set("Content-type", "application/json")
	request.Header.Set("Authorization", token.Access)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	resp, erro := client.Do(request)
	if erro != nil {
		log.Fatalf("An Error Occured %v", erro)
	}
	//Handle Error
	var cliente *models.TerceroSiigo

	err = json.NewDecoder(resp.Body).Decode(&cliente)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()

	//log.Printf(token.Pagination)
	return cliente, err
}

/*GetSiigoProducto end point items*/
func GetSiigoProducto(token *models.TokenSiigo) (*models.MessageSiigo, error) {
	var producto *models.ProductoSiigo
	var message models.MessageSiigo
	var erro error
	for page := 1; page <= 20; page++ {

		pageSize := 100
		postBody, _ := json.Marshal(map[string]string{
			"page": "2"})
		responseBody := bytes.NewBuffer(postBody)

		timeout := time.Duration(15 * time.Second)
		client := http.Client{
			Timeout: timeout,
		}
		request, err := http.NewRequest("GET", fmt.Sprintf("%s%d%s%d", "https://api.siigo.com/v1/products?page=", page, "&page_size=", pageSize), responseBody)
		request.Header.Set("Content-type", "application/json")
		request.Header.Set("Authorization", token.Access)
		request.Header.Set("Partner-Id", "Inventarios")
		if err != nil {
			log.Fatalf("An Error Occured %v", err)
		}
		resp, erro := client.Do(request)
		if erro != nil {
			log.Fatalf("An Error Occured %v", erro)
		}
		//Handle Error

		err = json.NewDecoder(resp.Body).Decode(&producto)
		if err != nil {
			log.Fatalf("An Error Occured %v", err)
		}
		defer resp.Body.Close()
		var items []models.GraboProducto
		for t := 0; t <= len(producto.Results)-1; t++ {
			var item models.GraboProducto
			item.Code = producto.Results[t].Code
			item.Name = producto.Results[t].Name
			item.Active = producto.Results[t].Active
			item.ID = producto.Results[t].ID
			items = append(items, item)
		}
		InsertProducto(items)
		//log.Printf(token.Pagination)
		erro = err
	}
	message.Message = "Productos cargados correctamentes"
	return &message, erro
}

/*GetSiigoTercero end point items*/
func GetSiigoTercero(token *models.TokenSiigo) (*models.MessageSiigo, error) {
	var tercero *models.TerceroSiigo
	var message models.MessageSiigo
	var erro error
	for page := 1; page <= 10; page++ {

		pageSize := 100
		postBody, _ := json.Marshal(map[string]string{
			"page": "1"})
		responseBody := bytes.NewBuffer(postBody)

		timeout := time.Duration(15 * time.Second)
		client := http.Client{
			Timeout: timeout,
		}
		request, err := http.NewRequest("GET", fmt.Sprintf("%s%d%s%d", "https://api.siigo.com/v1/customers?page=", page, "&page_size=", pageSize), responseBody)
		request.Header.Set("Content-type", "application/json")
		request.Header.Set("Authorization", token.Access)
		request.Header.Set("Partner-Id", "Inventarios")
		if err != nil {
			log.Fatalf("An Error Occured %v", err)
		}
		resp, erro := client.Do(request)
		if erro != nil {
			log.Fatalf("An Error Occured %v", erro)
		}
		//Handle Error

		err = json.NewDecoder(resp.Body).Decode(&tercero)
		if err != nil {
			log.Fatalf("An Error Occured %v", err)
		}
		defer resp.Body.Close()
		var items []models.GraboTercero
		for t := 0; t <= len(tercero.Results)-1; t++ {
			var item models.GraboTercero
			item.ID = tercero.Results[t].ID
			if len(tercero.Results[t].Name) > 1 {
				item.Name = fmt.Sprintf("%s%s%s", tercero.Results[t].Name[0], " ", tercero.Results[t].Name[1])
			} else {
				item.Name = tercero.Results[t].Name[0]
			}

			item.Identificacion = tercero.Results[t].Identification
			item.Active = tercero.Results[t].Active
			item.Adress = tercero.Results[t].Adress.Address
			item.City = tercero.Results[t].Adress.City.CityName
			item.State = tercero.Results[t].Adress.City.StateName
			item.Phone = fmt.Sprintf("%s%s", tercero.Results[t].Phones[0].Indicative, tercero.Results[t].Phones[0].Number)
			fmt.Printf("Adress: %s\n", tercero.Results[t].Adress.Address)
			items = append(items, item)
		}
		InsertTercero(items)
		//log.Printf(token.Pagination)
		erro = err
	}
	message.Message = "Terceros cargados correctamentes"
	return &message, erro
}

package bd

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)
var (
	server   = "drmonkey.co"
	port     = 1433
	user     = "drmonkey_honesolutionsAdmin"
	password = "Admin360!"
	database = "drmonkey_honesolutions"
)
/*Conexion es el objeto de conexión a la BD */
var Conexion = ConectarBD()

/*ConectarBD es la función que me permite conectar la BD */
func ConectarBD() *sql.DB {
	// Connect to database
	connString := fmt.Sprintf("drmonkey_Adminspirit:Admin360!@tcp(drmonkey.co:3306)/drmonkey_spirit")
	conn, err := sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	fmt.Printf("Connectedo\n")
	// defer conn.Close()
	return conn
}


/*ChequeoConnection es el Ping a la BD */
func ChequeoConnection() int {
	err := Conexion.PingContext(context.TODO())
	if err != nil {
		return 0
	}
	return 1
}

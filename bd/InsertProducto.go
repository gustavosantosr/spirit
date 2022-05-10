package bd

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/gustavosantosr/spirit/models"
)

/*InsertProducto EndPoint grabar salida*/
func InsertProducto(g []models.GraboProducto) (int64, bool, error) {
	var resultado int64 = 0
	for t := 0; t <= len(g)-1; t++ {
	
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	rows, err := Conexion.QueryContext(ctx, "SELECT COUNT(*) as count from Productos where ID = ? ", g[t].ID)
	if err != nil {
		return resultado, false, err
	}
	defer rows.Close()
	

		if checkCount(rows) != 0 {
			stmt, es := Conexion.Prepare(" UPDATE Productos SET  Code = ?, Name = ?, Active = ?  where ID = ?  ")
			if es != nil {
				return 0, false, es

			}
			result, er := stmt.Exec(g[t].Code,
				g[t].Name,
				g[t].Active,
				g[t].ID)
			if er != nil {
				return 0, false, er
			}
			resultado, _ = result.RowsAffected()
		} else {
			stmt, es := Conexion.Prepare("insert into Productos (ID, Code, Name, Active  ) values (?,?,?,?)")
			if es != nil {
				return 0, false, es

			}
			fmt.Printf("grupo: %s\n", g[t].Code)
			result, er := stmt.Exec(g[t].ID,
				g[t].Code,
				g[t].Name,
				g[t].Active)

			if er != nil {
				return 0, false, er
			}

			resultado, _ = result.LastInsertId()
		}
	
	}
	return resultado, true, nil
}
/*InsertTercero EndPoint grabar salida*/
func InsertTercero(g []models.GraboTercero) (int64, bool, error) {
	var resultado int64 = 0
	for t := 0; t <= len(g)-1; t++ {
	
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	rows, err := Conexion.QueryContext(ctx, "SELECT COUNT(*) as count from Terceros where ID = ? ", g[t].ID)
	if err != nil {
		return resultado, false, err
	}
	defer rows.Close()
	

		if checkCount(rows) != 0 {
			stmt, es := Conexion.Prepare(" UPDATE Terceros SET   Name = ?, Identificacion = ?, Active = ?, Adress=?, City=?, State=?, Phone=? where ID = ?  ")
			if es != nil {
				return 0, false, es

			}
			result, er := stmt.Exec(g[t].Name,
				g[t].Identificacion,
				g[t].Active,
				g[t].Adress,
				g[t].City,
				g[t].State,
				g[t].Phone,
				g[t].ID)
			if er != nil {
				return 0, false, er
			}
			resultado, _ = result.RowsAffected()
		} else {
			stmt, es := Conexion.Prepare("insert into Terceros (ID, Name, Identificacion, Active, Adress, City, State, Phone  ) values (?,?,?,?,?,?,?,?)")
			if es != nil {
				return 0, false, es

			}
			fmt.Printf("grupo: %s\n", g[t].ID)
			result, er := stmt.Exec(g[t].ID,
				g[t].Name,
				g[t].Identificacion,
				g[t].Active,
				g[t].Adress,
				g[t].City,
				g[t].State,
				g[t].Phone)

			if er != nil {
				return 0, false, er
			}

			resultado, _ = result.LastInsertId()
		}
	
	}
	return resultado, true, nil
}


func checkCount(rows *sql.Rows) (count int) {
	for rows.Next() {
	   err:= rows.Scan(&count)
	   checkErr(err)
   }   
   return count
}

func checkErr(err error) {
   if err != nil {
	   panic(err)
   }
}
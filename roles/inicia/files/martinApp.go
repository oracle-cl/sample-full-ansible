package main

import (
  "net/http"
  "log"
  "os"
  "fmt"
  "database/sql"
  _ "github.com/godror/godror"
)
 var pool *sql.DB 

func pagina(w http.ResponseWriter, r *http.Request) {
  hostname, _ := os.Hostname()
  arg := os.Args[1]
  datos := datos()
  message := fmt.Sprintf(`<!DOCTYPE html> <html lang=en-us>
<head> <meta name=viewport content=width=device-width, initial-scale=1.0>
<title> Ejemplo: Consulta SQL </title>
<style>
.datos { float: left; margin: 5px; padding: 15px; max-width: 300px;
    height: 300px; background-color:#610B0B; color:#ffffff; }
</style> </head>
<body style=background-color:#610B00;color:#ffffff;>
<div style=background-color:#8A0808;color:#ffffff;padding:2px; align=left>
  <h2>  Hostname %s Ambiente %s </h2> </div>

<div class=datos> <h3>SQL data count: %s</h3> </div> </body> </html> ` , hostname, arg, datos )

  w.Write([]byte(message))
}


func datos( ) string {
	if pool == nil {
	   return "Sin Conexion"
        } 
	row := pool.QueryRow("SELECT count(*) FROM martin.datos")

	var s string
	err := row.Scan(&s)
	if err != nil {
		return "Sin Datos"
     		log.Print(err)
	}
	return s
}

func main() {
  databaseDSN := os.Args[2]
  db , err := sql.Open("godror", databaseDSN)
  if err != nil {
     log.Print("Sin conexion a BD")
     log.Print(err)
     pool = nil
  } else {
     log.Print("Conexion a BD exitosa")
     pool = db
  }
  log.Print("Iniciando Server en 80" )
  http.HandleFunc("/", pagina)
  if err := http.ListenAndServe(":80", nil); err != nil {
    panic(err)
  }
}

package main

import (
	"log"
	"net/http"
	"os"

	"database/sql"

	_ "github.com/lib/pq"
)

func main() {
	//Coneccion App
	fileServer := http.FileServer(http.Dir("App"))
	http.Handle("/App/", http.StripPrefix("/App/", fileServer))

	//home Handle
	http.HandleFunc("/", handler)

	//Puertos
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	//DB
	connStr := "user=ozpxpneydsssjv dbname=db1sk6s96gfi85 sslmode=require"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	dat, err := db.Query("select * from nombre")
	println(dat)

	defer db.Close()

	db.Exec("CREATE DATABASE " + "Encuhfe")
	if err != nil {
		log.Fatal(err)
	}

	db.Exec("USE " + "Enchufe")
	if err != nil {
		log.Fatal(err)
	}

	db.Exec("CREATE TABLE example ( id integer, data varchar(32) )")
	if err != nil {
		log.Fatal(err)
	}
	//Host pagina
	http.ListenAndServe(":"+port, nil)
}
func handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")

}

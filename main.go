package main

import (
	"database/sql"
	"net/http"
	"os"
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
		db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
		if err != nil {
			panic(err)
		}
		defer db.Close()


			_, err = db.Exec("CREATE DATABASE " + "Encuhfe")
			if err != nil {
				panic(err)
			}

			_, err = db.Exec("USE " + "Enchufe")
			if err != nil {
				panic(err)
			}

			_, err = db.Exec("CREATE TABLE example ( id integer, data varchar(32) )")
			if err != nil {
				panic(err)
			}
	
	//Host pagina
	http.ListenAndServe(":"+port, nil)
}
func handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")

}

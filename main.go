package main

import (
	"bytes"
	"net/http"
	"os"
)

func main() {
	fileServer := http.FileServer(http.Dir("App"))
	http.Handle("/App/", http.StripPrefix("/App/", fileServer))

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	http.ListenAndServe(":8000", nil)

}
/*
func home(w http.Response, r *http.Request) {
	var b bytes.Buffer // A Buffer needs no initialization.
	b.Write([]byte("Hello World"))

}
*/
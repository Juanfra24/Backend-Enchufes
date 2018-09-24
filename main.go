package main
import (
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

	http.ListenAndServe(":"+port, nil)
}
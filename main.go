package main

import (
	"log"
	"net/http"

	"./routes"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	port := "8000"
	// Handle routes
	http.Handle("/", routes.Handlers())

	// serve
	log.Printf("Server up on port '%s'", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

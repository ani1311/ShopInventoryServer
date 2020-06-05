package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"./routes"
)

func main() {
	http.HandleFunc("/item", routes.ItemEndpoint)
	http.HandleFunc("/shop", routes.ShopEndpoint)
	log.Fatal(http.ListenAndServe(":8080", nil))
}


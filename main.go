package main

import (
	"log"
	"net/http"

	"./routes"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	http.HandleFunc("/item", routes.ItemEndpoint)
	http.HandleFunc("/shop", routes.ShopEndpoint)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

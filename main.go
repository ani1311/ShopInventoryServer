package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	http.HandleFunc("/item", routes.itemEndpoint)
	http.HandleFunc("/shop", routes.shopEndpoint)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

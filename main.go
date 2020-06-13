package main

import (
	"log"
	"net/http"

	"./routes"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	port := "8080"
	// Handle routes
	http.Handle("/", routes.Handlers())

	// dbUtils.InsertItem(models.Item{"4005900205711", "cream", 41})

	// serve
	log.Printf("Server up on port '%s'", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

	// dbUtils.InsertItem(models.Item{"dada", "fafa", 31})
	// dbUtils.InsertShop(models.Shop{"shsh7", "413", "414", "asd"})
	// dbUtils.UpdateOrInsertShopItem("dada", "shsh", false)
	// dbUtils.UpdateOrInsertShopItem("shsh7", "dada", true)

}

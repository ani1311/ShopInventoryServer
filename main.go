package main

import (
	"log"
	"net/http"

	// "./routes"
	// "github.com/gorilla/mux"

	"./routes"
	"github.com/gorilla/mux"
)

func main() {
	port := "8080"
	r := mux.NewRouter().StrictSlash(true)
	r.Use(CommonMiddleware)

	//Todo :  figure our how to use file Server
	r.HandleFunc("/index.html", routes.ServeIndexHTML)
	r.HandleFunc("/index.js", routes.ServeIndexJS)
	r.HandleFunc("/shopsWithItem.html", routes.ServeShopWithItemHTML)
	r.HandleFunc("/shopsWithItem.js", routes.ServeShopWithItemJS)

	r.HandleFunc("/register", routes.Register).Methods("POST")
	r.HandleFunc("/login", routes.Login).Methods("POST")
	r.HandleFunc("/test", routes.Test)
	r.HandleFunc("/getItems", routes.GetItemsEndpoint)
	r.HandleFunc("/shopWithItem", routes.ShopWithItemEndpoint).Methods("GET")

	s := r.PathPrefix("/auth").Subrouter()
	s.Use(routes.JwtVerify)
	s.HandleFunc("/addShopItem", routes.AddShopItemEndpoint)

	log.Printf("Server up on port '%s'", port)
	log.Fatal(http.ListenAndServe(":"+port, r))

	// fs := http.FileServer(http.Dir("./template"))
	// http.Handle("/", fs)

	// log.Println("Listening...")
	// http.ListenAndServe(":5558", nil)
}

func CommonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		next.ServeHTTP(w, r)
	})
}

package routes

//func Handlers() *mux.Router {

//	r := mux.NewRouter().StrictSlash(true)
//	r.Use(CommonMiddleware)

//	//website functions
//	fmt.Println(os.Open("website"))
//	r.Handle("/", http.FileServer(http.Dir("./website/")))

//	r.HandleFunc("/register", Register).Methods("POST")
//	r.HandleFunc("/login", Login).Methods("POST")
//	r.HandleFunc("/test", Test)
//	r.HandleFunc("/getItems", GetItemsEndpoint)
//	r.HandleFunc("/shopWithItem", ShopWithItemEndpoint)

//	// Auth route
//	s := r.PathPrefix("/auth").Subrouter()
//	s.Use(JwtVerify)
//	s.HandleFunc("/addShopItem", AddShopItemEndpoint)

//	return r
//}

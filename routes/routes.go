package routes

func itemEndpoint(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		dataBytes, err := ioutil.ReadAll(r.Body)
		checkError(err)
		fmt.Println(string(dataBytes))
		var data ItemData
		err = json.Unmarshal(dataBytes, &data)
		checkError(err)
		db, err := sql.Open("mysql", "aniket:aniket1311@(localhost:3306)/shop_inventory")
		fmt.Println(data)
		for _, item := range data.Data {
			insertItem(db, item)
		}
	}
}

func shopEndpoint(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		dataBytes, err := ioutil.ReadAll(r.Body)
		checkError(err)
		var data ShopData
		err = json.Unmarshal(dataBytes, data)
		checkError(err)
		db, err := sql.Open("mysql", "aniket:aniket1311@(localhost:3306)/shop_inventory")
		for _, shop := range data.Data {
			insertShop(db, shop)
		}
	}
}
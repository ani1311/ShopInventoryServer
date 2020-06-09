package routes

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"../dbUtils"
	"../models"
	"../utils"
)

func ShopEndpoint(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		dataBytes, err := ioutil.ReadAll(r.Body)
		utils.CheckError(err)
		var data models.ShopData
		err = json.Unmarshal(dataBytes, &data)
		utils.CheckError(err)
		db, err := sql.Open("mysql", "aniket:aniket1311@(localhost:3306)/shop_inventory")
		for _, shop := range data.Data {
			dbUtils.InsertShop(db, shop)
		}
		w.Write([]byte("Success"))
		w.WriteHeader(200)
	} else if r.Method == http.MethodGet {
		db, err := sql.Open("mysql", "aniket:aniket1311@(localhost:3306)/shop_inventory")
		shopId, ok := r.URL.Query()["shopId"]
		var data models.ShopData
		if !ok || len(shopId[0]) < 1 {
			data = dbUtils.SelectAllShop(db)
		} else {
			shopIdInt, err := strconv.ParseInt(shopId[0], 10, 32)
			utils.CheckError(err)
			data = dbUtils.SelectShopWithId(db, int(shopIdInt))
		}
		dataBytes, err := json.Marshal(data)
		utils.CheckError(err)
		w.Write(dataBytes)
	}
}

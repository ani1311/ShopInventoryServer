package routes

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"../dbUtils"
	"../utils"
	"../models"
)

func ItemEndpoint(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		dataBytes, err := ioutil.ReadAll(r.Body)
		utils.CheckError(err)
		var data models.ItemData
		err = json.Unmarshal(dataBytes, &data)
		utils.CheckError(err)
		db, err := sql.Open("mysql", "aniket:aniket1311@(localhost:3306)/shop_inventory")
		for _, item := range data.Data {
			dbUtils.InsertItem(db, item)
		}
	}
}

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
	}
}

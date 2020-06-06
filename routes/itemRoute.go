package routes

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"../dbUtils"
	"../models"
	"../utils"
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
	} else if r.Method == http.MethodGet {
		db, err := sql.Open("mysql", "aniket:aniket1311@(localhost:3306)/shop_inventory")
		barcode, ok := r.URL.Query()["barcode"]
		var data models.ItemData
		if !ok || len(barcode[0]) < 1 {
			data = dbUtils.SelectAllItem(db)
		} else {
			data = dbUtils.SelectItemWithBarcode(db, barcode[0])
		}
		dataBytes, err := json.Marshal(data)
		utils.CheckError(err)
		w.Write(dataBytes)
	}
}

package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"../dbUtils"
	"../models"
	"../utils"
)

func ItemEndpoint(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		fmt.Println("got stuff")
		dataBytes, err := ioutil.ReadAll(r.Body)
		utils.CheckError(err)
		var data models.ItemData
		err = json.Unmarshal(dataBytes, &data)
		utils.CheckError(err)
		for _, item := range data.Data {
			dbUtils.InsertItem(item)
		}
	} else if r.Method == http.MethodGet {
		barcode, ok := r.URL.Query()["barcode"]
		var data models.ItemData
		if !ok || len(barcode[0]) < 1 {
			data = dbUtils.SelectAllItem()
		} else {
			data = dbUtils.SelectItemWithBarcode(barcode[0])
		}
		dataBytes, err := json.Marshal(data)
		utils.CheckError(err)
		w.Write(dataBytes)
	}
}

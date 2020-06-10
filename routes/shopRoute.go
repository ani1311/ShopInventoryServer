package routes

import (
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
		for _, shop := range data.Data {
			dbUtils.InsertShop(shop)
		}
		w.Write([]byte("Success"))
		w.WriteHeader(200)
	} else if r.Method == http.MethodGet {
		shopId, ok := r.URL.Query()["shopId"]
		var data models.ShopData
		if !ok || len(shopId[0]) < 1 {
			data = dbUtils.SelectAllShop()
		} else {
			shopIdInt, err := strconv.ParseInt(shopId[0], 10, 32)
			utils.CheckError(err)
			data = dbUtils.SelectShopWithId(int(shopIdInt))
		}
		dataBytes, err := json.Marshal(data)
		utils.CheckError(err)
		w.Write(dataBytes)
	}
}

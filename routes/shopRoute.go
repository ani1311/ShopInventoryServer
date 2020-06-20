package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../dbUtils"
	"../mapUtils"
	"../models"
	"../utils"
)

func ShopWithItemEndpoint(w http.ResponseWriter, r *http.Request) {
	var shopData models.ShopData
	threshold := float64(4000)

	fmt.Print("Find item in shop ItemEndpoint : ")
	barcode, ok := r.URL.Query()[utils.BarcodeString]
	if !ok || len(barcode[0]) < 1 {
		fmt.Println("No Barcode Found ")
		json.NewEncoder(w).Encode(utils.InvalidRequestResponse)
		return
	}
	shopData = dbUtils.SelectShopWithItems(barcode[0])

	longitude, ok := r.URL.Query()[utils.LongitudeString]
	if !ok || len(longitude[0]) < 1 {
		fmt.Println("No Longitude Found ")
		json.NewEncoder(w).Encode(utils.InvalidRequestResponse)
		return
	}

	latitude, ok := r.URL.Query()[utils.LatitudeString]
	if !ok || len(latitude[0]) < 1 {
		fmt.Println("No Latitude Found ")
		json.NewEncoder(w).Encode(utils.InvalidRequestResponse)
		return
	}
	res := models.ShopData{}
	res.Data = make([]models.Shop, 0)
	

	for _, shop := range shopData.Data {
		if mapUtils.Distance(string(latitude[0]), string(longitude[0]), string(shop.Latitude), string(shop.Longitude)) < threshold {
			res.Data = append(res.Data, shop)
		}
	}

	dataBytes, err := json.Marshal(res)
	utils.CheckError(err)
	w.Write(dataBytes)
}

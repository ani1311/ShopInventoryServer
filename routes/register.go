package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"../dbUtils"
	"../models"
	"../utils"
)

func Register(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Register attempt : ")
	shop := models.Shop{}
	
	bytes, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(bytes, &shop)
	if err != nil {
		fmt.Println("Invalid Values in body ")
		json.NewEncoder(w).Encode(utils.InvalidRequestResponse)
		return
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(shop.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Wrong Password ",err)
		json.NewEncoder(w).Encode(utils.PasswordEncryptionFailResponse)
		return
	}

	if shop.Longitude == "" || shop.Latitude == "" {
		fmt.Println("No location ingformation ")
		json.NewEncoder(w).Encode(utils.InvalidRequestResponse)
		return
	}

	shop.Password = string(pass)
	if !dbUtils.InsertShop(shop) {
		fmt.Println("DB error ")
		json.NewEncoder(w).Encode(utils.InvalidRequestResponse)
		return
	}
	json.NewEncoder(w).Encode(utils.SuccessfulResponse)

}

package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"../dbUtils"
	"../models"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	csvfile, err := os.Open("/home/aniket/Documents/MyFiles/shopInventory/ShopInventoryServer/testFiles/testShops.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	r := csv.NewReader(csvfile)

	record, err := r.Read()
	for {
		record, err = r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		shop := models.Shop{record[0], record[2], record[1], record[3]}
		pass, err := bcrypt.GenerateFromPassword([]byte(shop.Password), bcrypt.DefaultCost)
		if err != nil {
			fmt.Println("Wrong Password ", err)
			return
		}

		if shop.Longitude == "" || shop.Latitude == "" {
			fmt.Println("No location ingformation ")
			return
		}

		shop.Password = string(pass)
		if !dbUtils.InsertShop(shop) {
			fmt.Println("DB error ")
			return
		}

	}
}

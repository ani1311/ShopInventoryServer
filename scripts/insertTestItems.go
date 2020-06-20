package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"

	"../dbUtils"
	"../models"
)

func main() {
	csvfile, err := os.Open("/home/aniket/Documents/MyFiles/shopInventory/ShopInventoryServer/testFiles/testItems.csv")
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
		dbUtils.InsertItem(models.Item{record[1], record[0], 1000})

	}
}

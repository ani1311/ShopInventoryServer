package main

import (
	"fmt"
	"math/rand"

	"../dbUtils"
)

func main() {
	allItems := dbUtils.SelectAllItem()
	allShops := dbUtils.SelectAllShop()

	for i := 0; i < 20; i++ {
		ti := rand.Int() % len(allItems.Data)
		tj := rand.Int() % len(allShops.Data)
		fmt.Println(ti, tj)
		dbUtils.UpdateOrInsertShopItem(allShops.Data[tj].Name, allItems.Data[ti].Barcode, true)
	}

}

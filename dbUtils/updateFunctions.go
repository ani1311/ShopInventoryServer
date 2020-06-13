package dbUtils

import (
	"../utils"
)

func UpdateShopItem(shopName, barcode string, available bool) {
	db := getDB()
	stmt, err := db.Prepare("UPDATE shop_item SET available=? WHERE barcode=? AND shop_name=?")
	utils.CheckError(err)
	_, err = stmt.Query(available, barcode, shopName)
	utils.CheckError(err)
}

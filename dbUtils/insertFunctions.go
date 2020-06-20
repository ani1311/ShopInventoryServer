package dbUtils

import (
	"fmt"

	"../models"
	"../utils"
)

func InsertShop(shop models.Shop) bool {
	db := getDB()
	defer db.Close()
	stmt, err := db.Prepare("INSERT shop SET name=?,longitude=?,latitude=?,password=?")
	if utils.CheckError(err) {
		return false
	}
	res, err := stmt.Exec(shop.Name, shop.Longitude, shop.Latitude, shop.Password)
	if utils.CheckError(err) {
		return false
	}
	affect, err := res.RowsAffected()
	if utils.CheckError(err) {
		return false
	}
	fmt.Print(" Inserted ", affect, " rows ")
	return true
}

func InsertItem(item models.Item) bool {
	db := getDB()
	stmt, err := db.Prepare("INSERT item SET barcode=?,name=?,price=?")
	if utils.CheckError(err) {

		return false
	}
	res, err := stmt.Exec(item.Barcode, item.Name, item.Price)
	if utils.CheckError(err) {
		return false
	}
	affect, err := res.RowsAffected()
	if utils.CheckError(err) {
		return false
	}
	fmt.Print(" Inserted ", affect, " rows ")
	return true
}

func InsertShopItem(shopItem models.ShopItem) bool {
	db := getDB()
	stmt, err := db.Prepare("INSERT shop_item SET shop_name=?,barcode=?,available=?")
	if utils.CheckError(err) {
		return false
	}
	res, err := stmt.Exec(shopItem.Shopname, shopItem.Barcode, shopItem.Available)

	if utils.CheckError(err) {
		return false
	}
	affect, err := res.RowsAffected()
	if utils.CheckError(err) {
		return false
	}
	fmt.Print(" Inserted ", affect, " rows ")
	return true
}

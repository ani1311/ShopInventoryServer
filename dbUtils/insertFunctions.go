package dbUtils

import (
	"fmt"

	"../models"
	"../utils"
)

func InsertShop(shop models.Shop) {
	db := getDB()
	defer db.Close()
	stmt, err := db.Prepare("INSERT shop SET shopid=?,name=?,longitude=?,latitude=?")
	utils.CheckError(err)
	res, err := stmt.Exec(shop.ShopId, shop.Name, shop.Longitude, shop.Latitude)
	utils.CheckError(err)
	affect, err := res.RowsAffected()
	utils.CheckError(err)
	fmt.Println(affect)
}

func InsertShopClient(shopClient models.ShopClient) {
	db := getDB()
	stmt, err := db.Prepare("INSERT shop_client SET username=?,shopid=?,password=?")
	utils.CheckError(err)
	res, err := stmt.Exec(shopClient.Username, shopClient.ShopId, shopClient.Password)
	utils.CheckError(err)
	affect, err := res.RowsAffected()
	utils.CheckError(err)
	fmt.Println(affect)
}

func InsertItem(item models.Item) {
	db := getDB()
	stmt, err := db.Prepare("INSERT item SET barcode=?,name=?,price=?")
	utils.CheckError(err)
	res, err := stmt.Exec(item.Barcode, item.Name, item.Price)
	utils.CheckError(err)
	affect, err := res.RowsAffected()
	utils.CheckError(err)
	fmt.Println(affect)
}

func InsertShopItem(shopItem models.ShopItem) {
	db := getDB()
	stmt, err := db.Prepare("INSERT shop_item SET shopid=?,barcode=?,quantity=?")
	utils.CheckError(err)
	res, err := stmt.Exec(shopItem.ShopId, shopItem.Barcode, shopItem.Quantity)
	utils.CheckError(err)
	affect, err := res.RowsAffected()
	utils.CheckError(err)
	fmt.Println(affect)
}

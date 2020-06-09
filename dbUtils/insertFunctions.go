package dbUtils

import (
	"database/sql"
	"fmt"

	"../models"
	"../utils"
)

func InsertShop(db *sql.DB, shop models.Shop) {
	stmt, err := db.Prepare("INSERT shop SET shopid=?,name=?,longitude=?,latitude=?")
	utils.CheckError(err)
	res, err := stmt.Exec(shop.ShopId, shop.Name, shop.Longitude, shop.Latitude)
	utils.CheckError(err)
	affect, err := res.RowsAffected()
	utils.CheckError(err)
	fmt.Println(affect)
}

func InsertItem(db *sql.DB, item models.Item) {
	stmt, err := db.Prepare("INSERT item SET barcode=?,name=?,price=?")
	utils.CheckError(err)
	res, err := stmt.Exec(item.Barcode, item.Name, item.Price)
	utils.CheckError(err)
	affect, err := res.RowsAffected()
	utils.CheckError(err)
	fmt.Println(affect)
}

func InsertShopItem(db *sql.DB, shopItem models.ShopItem) {
	stmt, err := db.Prepare("INSERT shop_item SET shopid=?,barcode=?,quantity=?")
	utils.CheckError(err)
	res, err := stmt.Exec(shopItem.ShopId, shopItem.Barcode, shopItem.Quantity)
	utils.CheckError(err)
	affect, err := res.RowsAffected()
	utils.CheckError(err)
	fmt.Println(affect)
}

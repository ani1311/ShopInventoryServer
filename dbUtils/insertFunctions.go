package dbUtils

import (
	"database/sql"
	"fmt"
	"../models"
	"../utils"
)

func InsertShop(db *sql.DB, shop models.Shop) {
	stmt, err := db.Prepare("INSERT shop SET shopid=?,barcode=?,quantity=?,name=?")
	utils.CheckError(err)
	res, err := stmt.Exec(shop.ShopId, shop.Barcode, shop.Quantity, shop.Name)
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

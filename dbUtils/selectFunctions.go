package dbUtils

import (
	"database/sql"

	"../models"
	"../utils"
)

func SelectShopWithId(db *sql.DB, shopId int) models.ShopData {
	stmt, err := db.Prepare("SELECT * FROM shop WHERE shopid=?")
	utils.CheckError(err)
	res, err := stmt.Query(shopId)
	utils.CheckError(err)
	var shopData models.ShopData
	shopData.Data = make([]models.Shop, 0)
	for res.Next() {
		var shop models.Shop
		res.Scan(&shop.ShopId, &shop.Barcode, &shop.Quantity, &shop.Name)
		shopData.Data = append(shopData.Data, shop)
	}
	return shopData
}

func SelectAllShop(db *sql.DB) models.ShopData {
	stmt, err := db.Prepare("SELECT * FROM shop ")
	utils.CheckError(err)
	res, err := stmt.Query()
	utils.CheckError(err)
	var shopData models.ShopData
	shopData.Data = make([]models.Shop, 0)
	for res.Next() {
		var shop models.Shop
		res.Scan(&shop.ShopId, &shop.Barcode, &shop.Quantity, &shop.Name)
		shopData.Data = append(shopData.Data, shop)
	}
	return shopData
}

func SelectItemWithBarcode(db *sql.DB, barcode string) models.ItemData {
	stmt, err := db.Prepare("SELECT * FROM item WHERE barcode=?")
	utils.CheckError(err)
	res, err := stmt.Query(barcode)
	utils.CheckError(err)
	var itemData models.ItemData
	itemData.Data = make([]models.Item, 0)
	for res.Next() {
		var item models.Item
		res.Scan(&item.Barcode, &item.Name, &item.Price)
		itemData.Data = append(itemData.Data, item)
	}
	return itemData
}

func SelectAllItem(db *sql.DB) models.ItemData {
	stmt, err := db.Prepare("SELECT * FROM item")
	utils.CheckError(err)
	res, err := stmt.Query()
	utils.CheckError(err)
	var itemData models.ItemData
	itemData.Data = make([]models.Item, 0)
	for res.Next() {
		var item models.Item
		res.Scan(&item.Barcode, &item.Name, &item.Price)
		itemData.Data = append(itemData.Data, item)
	}
	return itemData
}

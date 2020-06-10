package dbUtils

import (
	"../models"
	"../utils"
)

func SelectShopWithId(shopId int) models.ShopData {
	db := getDB()
	stmt, err := db.Prepare("SELECT * FROM shop WHERE shopid=?")
	utils.CheckError(err)
	res, err := stmt.Query(shopId)
	utils.CheckError(err)
	var shopData models.ShopData
	shopData.Data = make([]models.Shop, 0)
	for res.Next() {
		var shop models.Shop
		res.Scan(&shop.ShopId, &shop.Name, &shop.Longitude, &shop.Latitude)
		shopData.Data = append(shopData.Data, shop)
	}
	return shopData
}

func SelectAllShop() models.ShopData {
	db := getDB()
	stmt, err := db.Prepare("SELECT * FROM shop ")
	utils.CheckError(err)
	res, err := stmt.Query()
	utils.CheckError(err)
	var shopData models.ShopData
	shopData.Data = make([]models.Shop, 0)
	for res.Next() {
		var shop models.Shop
		res.Scan(&shop.ShopId, &shop.Name, &shop.Longitude, &shop.Latitude)
		shopData.Data = append(shopData.Data, shop)
	}
	return shopData
}

func SelectItemWithBarcode(barcode string) models.ItemData {
	db := getDB()
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

func SelectAllItem() models.ItemData {
	db := getDB()
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

func SelectShopClient(username string) *models.ShopClient {
	db := getDB()
	defer db.Close()
	stmt, err := db.Prepare("SELECT * FROM item WHERE username=?")
	utils.CheckError(err)
	res, err := stmt.Query()
	utils.CheckError(err)
	if res.Next() {
		var shopClient models.ShopClient
		res.Scan(&shopClient.Username, &shopClient.ShopId, &shopClient.Password)
		if res.Next() {
			return nil
		}
		return &shopClient
	}
	return nil
}

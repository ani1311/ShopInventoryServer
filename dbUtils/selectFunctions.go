package dbUtils

import (
	"../models"
	"../utils"
)

func SelectShopWithName(shopName string) *models.Shop {
	db := getDB()
	stmt, err := db.Prepare("SELECT * FROM shop WHERE name=?")
	utils.CheckError(err)
	res, err := stmt.Query(shopName)
	utils.CheckError(err)

	if res.Next() {
		var shop models.Shop
		res.Scan(&shop.Name, &shop.Longitude, &shop.Latitude, &shop.Password)
		if res.Next() {
			return nil
		}
		return &shop
	}
	return nil
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
		res.Scan(&shop.Name, &shop.Longitude, &shop.Latitude, &shop.Password)
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

func SelectShopItem(barcode, shopName string) models.ShopItemData {
	db := getDB()
	stmt, err := db.Prepare("SELECT * FROM shop_item WHERE barcode=? AND shop_name=?")
	utils.CheckError(err)
	res, err := stmt.Query(barcode, shopName)
	utils.CheckError(err)
	var shopItemData models.ShopItemData
	shopItemData.Data = make([]models.ShopItem, 0)
	for res.Next() {
		var item models.ShopItem
		res.Scan(&item.Barcode, &item.Shopname, &item.Available)
		shopItemData.Data = append(shopItemData.Data, item)
	}
	return shopItemData
}

func SelectShopWithItems(barcode string) models.ShopData {
	db := getDB()
	stmt, err := db.Prepare("SELECT * FROM shop WHERE name IN (SELECT shop_name FROM shop_item where barcode=?)")
	utils.CheckError(err)
	res, err := stmt.Query(barcode)
	utils.CheckError(err)
	var shopData models.ShopData
	shopData.Data = make([]models.Shop, 0)
	for res.Next() {
		var shop models.Shop
		res.Scan(&shop.Name, &shop.Longitude, &shop.Latitude, &shop.Password)
		shopData.Data = append(shopData.Data, shop)
	}
	return shopData
}

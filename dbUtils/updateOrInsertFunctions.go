package dbUtils

import (
	"../models"
)

func UpdateOrInsertShopItem(shopName, barcode string, available bool) {
	results := SelectShopItem(barcode, shopName)
	if len(results.Data) == 0 {
		InsertShopItem(models.ShopItem{shopName, barcode, available})
	} else {
		UpdateShopItem(shopName, barcode, available)
	}
}

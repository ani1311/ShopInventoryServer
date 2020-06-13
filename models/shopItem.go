package models

type ShopItem struct {
	Shopname  string `json:"shopName"`
	Barcode   string `json:"barcode"`
	Available bool   `json:"available"`
}

type ShopItemData struct {
	Data []ShopItem `json:"data"`
}

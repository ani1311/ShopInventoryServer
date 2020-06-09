package models

type ShopItem struct {
	ShopId   int    `json:"shopId"`
	Barcode  string `json:"barcode"`
	Quantity int    `json:"quantity"`
}

type ShopItemData struct {
	Data []ShopItem `json:"data"`
}

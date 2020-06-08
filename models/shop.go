package models

type Shop struct {
	ShopId   int    `json:"shopId"`
	Barcode  string `json:"barcode"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

type ShopData struct {
	Data []Shop `json:"data"`
}

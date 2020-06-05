package models

type Shop struct {
	ShopId   int
	Barcode  string
	Name     string
	Quantity int
}

type ShopData struct {
	Data []Shop `json:"data"`
}
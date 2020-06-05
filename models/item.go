package models

type Item struct {
	Barcode string `json:"barcode"`
	Name    string `json:"name"`
	Price   int    `json:"price"`
}

type ItemData struct {
	Data []Item `json:"data"`
}

package models

type Shop struct {
	ShopId    int    `json:"shopId"`
	Name      string `json:"name"`
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
}

type ShopData struct {
	Data []Shop `json:"data"`
}

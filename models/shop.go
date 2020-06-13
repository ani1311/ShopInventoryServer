package models

type Shop struct {
	Name      string `json:"name"`
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
	Password string `json:"password"`
}

type ShopData struct {
	Data []Shop `json:"data"`
}

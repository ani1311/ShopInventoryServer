package models

type ShopClient struct {
	Username string `json:"username"`
	ShopId   int    `json:"shopId"`
	Password string `json:"password"`
}

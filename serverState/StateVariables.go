package serverState

import "fmt"

var SessionToShopMap map[string]string
var ShopToSessionMap map[string]string

func init() {
	SessionToShopMap = make(map[string]string)
	ShopToSessionMap = make(map[string]string)
}

func PrintServerToShopMap() {
	fmt.Println(SessionToShopMap)
}

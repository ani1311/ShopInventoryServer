package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../dbUtils"
	"../searchUtils"
	"../serverState"
	"../utils"
)

func AddShopItemEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Item ItemEndpoint : ")
	barcode, ok := r.URL.Query()[utils.BarcodeString]
	if !ok || len(barcode[0]) < 1 {
		fmt.Println("No Barcode Found ")
		json.NewEncoder(w).Encode(utils.InvalidRequestResponse)
		return
	}
	operation, ok := r.URL.Query()[utils.OperationString]
	if !ok || len(operation[0]) < 1 {
		fmt.Println("No Operation Found ")
		json.NewEncoder(w).Encode(utils.InvalidRequestResponse)
		return
	}
	sessionID, ok := r.URL.Query()[utils.SessionIDString]
	if !ok || len(sessionID[0]) < 1 {
		fmt.Println("No Session ID Found ")
		json.NewEncoder(w).Encode(utils.InvalidRequestResponse)
		return
	}

	fmt.Println(barcode, operation, sessionID, serverState.SessionToShopMap[sessionID[0]])

	if operation[0] == utils.AddOperationString {
		fmt.Println("Add ")
		dbUtils.UpdateOrInsertShopItem(serverState.SessionToShopMap[sessionID[0]], barcode[0], true)
		// if !dbUtils.UpdateOrInsertShopItem(serverState.SessionToShopMap[sessionID[0]], barcode[0], true) {
		// 	fmt.Println("DB error ")
		// 	json.NewEncoder(w).Encode(utils.InvalidRequestResponse)
		// 	return
		// }
	} else if operation[0] == utils.RemoveOperationString {
		fmt.Println("Remove ")
		dbUtils.UpdateOrInsertShopItem(serverState.SessionToShopMap[sessionID[0]], barcode[0], false)
		// if !dbUtils.UpdateOrInsertShopItem(serverState.SessionToShopMap[sessionID[0]], barcode[0], false){
		// 	fmt.Println("DB error ")
		// 	json.NewEncoder(w).Encode(utils.InvalidRequestResponse)
		// 	return
		// }
	} else {
		fmt.Println("Invalid Operation ")
		json.NewEncoder(w).Encode(utils.InvalidRequestResponse)
		return
	}
}

func GetItemsEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting items")
	name, ok := r.URL.Query()["name"]
	if !ok || len(name[0]) < 1 {
		fmt.Println("No Name Found ")
		json.NewEncoder(w).Encode(utils.InvalidRequestResponse)
		return
	}
	items := searchUtils.SearchItem(name[0])
	resp, _ := json.Marshal(items)
	w.Write(resp)
}

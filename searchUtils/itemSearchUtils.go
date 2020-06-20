package searchUtils

import (
	"strings"

	"../dbUtils"
	"../utils"

	"../models"
	"github.com/texttheater/golang-levenshtein/levenshtein"
)

const distThreshold = 2

func SearchItem(searchName string) models.ItemData {
	allItems := dbUtils.SelectAllItem()
	var res models.ItemData
	res.Data = make([]models.Item, 0)
	searchWords := strings.Fields(searchName)
	for _, item := range allItems.Data {
		itemWords := strings.Fields(item.Name)
		for i := 0; i < utils.Min(len(itemWords), len(searchWords)); i++ {
			if checkName(searchWords[i], itemWords[i]) {
				res.Data = append(res.Data, item)
			}
		}
	}
	return res
}

func checkName(searchName, itemName string) bool {
	cost := levenshtein.DistanceForStrings([]rune(searchName), []rune(itemName), levenshtein.DefaultOptionsWithSub)
	if cost < distThreshold {
		return true
	}
	return false
}

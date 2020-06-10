package dbUtils

import (
	"database/sql"

	"../utils"
)

func getDB() *sql.DB {
	db, err := sql.Open("mysql", "aniket:aniket1311@(localhost:3306)/shop_inventory")
	utils.CheckError(err)
	return db

}

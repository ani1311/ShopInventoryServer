package dbUtils

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func getDB() *sql.DB {
	db, err := sql.Open("mysql", "aniket:aniket1311@(localhost:3306)/shop_inventory")
	fmt.Println(err)
	// utils.CheckError(err)
	return db

}

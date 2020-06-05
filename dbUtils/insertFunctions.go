package dbUtils

func insertShop(db *sql.DB, shop Shop) {
	stmt, err := db.Prepare("INSERT shop SET shopid=?,barcode=?,quantity=?,name=?")
	checkError(err)
	res, err := stmt.Exec(shop.ShopId, shop.Barcode, shop.Quantity, shop.Name)
	checkError(err)
	affect, err := res.RowsAffected()
	checkError(err)
	fmt.Println(affect)
}

func insertItem(db *sql.DB, item Item) {
	stmt, err := db.Prepare("INSERT item SET barcode=?,name=?,price=?")
	checkError(err)
	res, err := stmt.Exec(item.Barcode, item.Name, item.Price)
	checkError(err)
	affect, err := res.RowsAffected()
	checkError(err)
	fmt.Println(affect)
}

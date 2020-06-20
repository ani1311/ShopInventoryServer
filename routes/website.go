package routes

import (
	"net/http"
	"text/template"

	"../utils"
)

func ServeIndexHTML(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	utils.CheckError(err)
	data := make([]int, 0)
	tmpl.Execute(w, data)
}

func ServeIndexJS(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.js")
	utils.CheckError(err)
	data := make([]int, 0)
	tmpl.Execute(w, data)
}
func ServeShopWithItemHTML(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/shopsWithItem.html")
	utils.CheckError(err)
	data := make([]int, 0)
	tmpl.Execute(w, data)
}

func ServeShopWithItemJS(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/shopsWithItem.js")
	utils.CheckError(err)
	data := make([]int, 0)
	tmpl.Execute(w, data)
}

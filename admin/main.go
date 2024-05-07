package main

import (
	"net/http"

	"github.com/tiqueteo/adminv2-mock-api/api/controllers"
	"github.com/tiqueteo/adminv2-mock-api/db"
)

func main() {

	database, _ := db.InitDB()
	r := controllers.Init(database)
	http.ListenAndServe(":8080", r)
}

package main

import (
	"net/http"

	"github.com/tiqueteo/adminv2-mock-api/api/controllers"
)

func main() {
	r := controllers.Init()
	http.ListenAndServe(":8080", r)
}

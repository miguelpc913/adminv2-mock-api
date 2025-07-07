package main

import (
	"log"
	"net/http"
	"os"
	"runtime/debug"

	"github.com/tiqueteo/adminv2-mock-api/api/controllers"
)

func main() {

	defer func() {
		if err := recover(); err != nil {
			log.Printf("Panic occurred: %v\nStack trace:\n%s", err, debug.Stack())
			os.Exit(1)
		}
	}()
	port := "8080"
	r := controllers.Init()
	http.ListenAndServe(":"+port, r)
	log.Printf("Escuchando en el puerto %s", port)
}

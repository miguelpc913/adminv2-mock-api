package main

import (
	"log"
	"net/http"
	"os"
	"runtime/debug"

	"github.com/joho/godotenv"
	"github.com/tiqueteo/adminv2-mock-api/api/controllers"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("No %s file found or error loading it\n", ".env")
	}
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Panic occurred: %v\nStack trace:\n%s", err, debug.Stack())
			os.Exit(1)
		}
	}()

	r := controllers.Init()
	http.ListenAndServe(":8080", r)
}

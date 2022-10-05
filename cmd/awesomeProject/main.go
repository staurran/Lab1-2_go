package main

import (
	"awesomeProject/internal/pkg/app"
	"log"
	"os"
)

func main() {
	log.Println("Application start")
	a, err := app.New()
	if err != nil {
		log.Println("Application failed")
		os.Exit(2)
	}
	a.StartServer()
	log.Println("Application terminate")
}

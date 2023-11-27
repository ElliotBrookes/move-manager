package main

import (
	"ElliotBrookes/move-manager/internal/configuration"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	//coordMap := make(map[string]int)

	// Load env vars & tokens
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	app := configuration.New()
	app.Start()

}

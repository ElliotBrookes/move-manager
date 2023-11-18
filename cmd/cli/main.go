package main

import (
	"ElliotBrookes/move-manager/internal/adapters/retrieval/level_db"
	"ElliotBrookes/move-manager/internal/ingress"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	//coordMap := make(map[string]int)

	// Load env vars & tokens
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	t := time.Now()
	// Init/ Load up db
	db, err := level_db.OpenLevelDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Println("Time to Init/Load dataset", time.Since(t))

	// Collect all the args, discard the program name call
	err = ingress.HandleCliInvoc(os.Args)
	if err != nil {
		log.Fatal(err.Error())
	}

}

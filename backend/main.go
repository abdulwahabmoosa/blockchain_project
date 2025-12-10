package main

import (
	"backend/api"
	"backend/db"
	"log"
)

func main() {
	db, err := db.NewDatabase()

	if err != nil {
		panic(err)
	}

	log.Printf("db connection successful\n")

	handler := api.NewRequestHandler(db)
	handler.Start()
}

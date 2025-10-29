package main

import (
	"backend/api"
	"backend/db"
	"fmt"

	_ "ariga.io/atlas-provider-gorm/gormschema"
)

func main() {
	db, err := db.NewDatabase()

	if err != nil {
		panic(err)
	}

	fmt.Printf("db connection successful")

	handler := api.NewRequestHandler(db)
	handler.Start()
}

package main

import (
	"backend/api"
	"backend/blockchain"
	"backend/db"
	"log"
)

func main() {
	// 1. Initialize DB
	log.Printf("🔄 Initializing database...")
	database, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("❌ DB Connection failed: %v", err)
	}
	log.Printf("✅ DB connection successful")

	// 2. Initialize Blockchain Service
	log.Printf("🔄 Initializing blockchain service...")
	chainService, err := blockchain.NewChainServiceEnv()
	if err != nil {
		log.Printf("⚠️ Blockchain Connection failed: %v", err)
		log.Printf("⚠️ Continuing without blockchain functionality")
		chainService = nil
	} else {
		log.Printf("✅ Blockchain connection successful (Chain ID: %s)", chainService.ChainID.String())
	}

	// 3. Start API Server
	// Pass both DB and ChainService to the handler
	handler := api.NewRequestHandler(database, chainService)
	handler.Start()
}

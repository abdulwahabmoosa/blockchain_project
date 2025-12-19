package main

import (
	"backend/api"
	"backend/blockchain"
	"backend/blockchain/worker"
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

	// 3. Start blockchain event listeners (if blockchain is available)
	if chainService != nil {
		log.Printf("🔄 Starting blockchain event listeners...")
		worker.StartListeners(chainService, database)
		log.Printf("✅ Blockchain listeners started (may show warnings if RPC doesn't support subscriptions)")
	}

	// 4. Start API Server
	// Pass both DB and ChainService to the handler
	handler := api.NewRequestHandler(database, chainService)
	handler.Start()
}


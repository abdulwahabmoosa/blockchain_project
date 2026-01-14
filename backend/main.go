// main package - entry point for backend application
package main

import (
	"backend/api"
	"backend/blockchain"
	"backend/blockchain/worker"
	"backend/db"
	"log"
)

// main function - application startup sequence
func main() {
	// setup database first, required for everything
	// database init includes migrations and admin user creation
	log.Printf("Setting up database connection...")
	database, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Printf("Database connected successfully")

	// try blockchain connection, optional - system works without it
	// blockchain service handles smart contract interactions
	log.Printf("Attempting to connect to blockchain...")
	chainService, err := blockchain.NewChainServiceEnv()
	if err != nil {
		log.Printf("Warning: Could not connect to blockchain: %v", err)
		log.Printf("Warning: System will continue without blockchain features")
		chainService = nil
	} else {
		log.Printf("Blockchain connected (Chain ID: %s)", chainService.ChainID.String())
	}

	// start event listeners if blockchain available
	// listeners monitor contract events and update database
	if chainService != nil {
		log.Printf("Starting blockchain event monitoring...")
		worker.StartListeners(chainService, database)
		log.Printf("Event listeners active (might see warnings if RPC doesn't support event subscriptions)")
	}

	// finally start the API server
	// api handler needs both database and blockchain service
	log.Printf("Starting API server...")
	handler := api.NewRequestHandler(database, chainService)
	handler.Start()
}


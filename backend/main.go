package main

import (
	"backend/api"
	"backend/blockchain"
	"backend/db"
	"backend/db/models"
	"log"
	"time"
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

	// 3. Start Revenue Distribution Scheduler (if blockchain is available)
	if chainService != nil {
		log.Printf("🔄 Starting revenue distribution scheduler...")
		go startRevenueScheduler(database, chainService)
		log.Printf("✅ Revenue scheduler started (runs every 24 hours)")
	}

	// 4. Start API Server
	// Pass both DB and ChainService to the handler
	handler := api.NewRequestHandler(database, chainService)
	handler.Start()
}

// startRevenueScheduler runs every 24 hours to automatically distribute revenue to all active properties
func startRevenueScheduler(database *db.Database, chainService *blockchain.ChainService) {
	ticker := time.NewTicker(24 * time.Hour) // Run every 24 hours
	defer ticker.Stop()

	// Run immediately on startup for testing
	go distributeRevenueToAllProperties(database, chainService)

	for range ticker.C {
		log.Printf("🔄 Starting scheduled revenue distribution...")
		distributeRevenueToAllProperties(database, chainService)
	}
}

// distributeRevenueToAllProperties creates snapshots and distributes revenue to all active properties
func distributeRevenueToAllProperties(database *db.Database, chainService *blockchain.ChainService) {
	// Get all active properties
	properties, err := database.GetAllProperties()
	if err != nil {
		log.Printf("❌ Failed to get properties for revenue distribution: %v", err)
		return
	}

	activeProperties := make([]models.Property, 0)
	for _, prop := range properties {
		if prop.Status == "Active" {
			activeProperties = append(activeProperties, prop)
		}
	}

	if len(activeProperties) == 0 {
		log.Printf("ℹ️ No active properties found for revenue distribution")
		return
	}

	log.Printf("💰 Starting revenue distribution for %d active properties", len(activeProperties))

	// For testing: distribute 0.001 ETH worth of test USDC to each property
	// In production, this could be based on actual rental income
	testRevenueAmount := "1000000" // 1 million wei (0.001 ETH in wei for 18-decimal tokens)

	// Use a mock stablecoin address for testing (USDC on Sepolia)
	// In production, this should be a real stablecoin contract
	testStablecoinAddress := "0x1c7D4B196Cb0C7B01d743Fbc6116a902379C7238" // Example USDC address

	distributedCount := 0
	for _, property := range activeProperties {
		log.Printf("📊 Creating snapshot and distributing revenue for property: %s", property.OnchainTokenAddress)

		// Create snapshot and distribute revenue
		_, err := chainService.DistributeRevenue(
			property.OnchainTokenAddress,
			testStablecoinAddress,
			testRevenueAmount,
		)

		if err != nil {
			log.Printf("❌ Failed to distribute revenue to property %s: %v", property.ID, err)
			continue
		}

		log.Printf("✅ Successfully distributed revenue to property %s", property.ID)
		distributedCount++
	}

	log.Printf("🎉 Revenue distribution completed: %d/%d properties received revenue", distributedCount, len(activeProperties))
}

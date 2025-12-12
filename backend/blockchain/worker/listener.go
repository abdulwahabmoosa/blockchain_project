package worker

import (
	"backend/blockchain"
	"backend/blockchain/property_factory"
	"backend/blockchain/revenue_distribution"
	"backend/db"
	"backend/db/models"
	"log"
	"math/big"
	"time"

	"github.com/google/uuid"
)

// StartListeners spins up background routines to watch contract events
func StartListeners(chain *blockchain.ChainService, database *db.Database) {
	go listenForProperties(chain, database)
	go listenForRevenue(chain, database)
}

func listenForProperties(chain *blockchain.ChainService, database *db.Database) {
	sink := make(chan *property_factory.PropertyFactoryPropertyRegistered)

	// Subscribe to "PropertyRegistered" events
	// nil arguments mean "listen to all" (no filters)
	sub, err := chain.PropertyFactory.WatchPropertyRegistered(nil, sink, nil)
	if err != nil {
		log.Fatalf("Failed to subscribe to property events: %v", err)
	}

	log.Println("🎧 Listening for New Properties...")

	for {
		select {
		case err := <-sub.Err():
			log.Printf("Property Subscription error: %v", err)
			// Reconnect logic would go here in production
			time.Sleep(5 * time.Second)
		case event := <-sink:
			log.Printf("📢 Event: Property Registered at %s", event.PropertyAsset.Hex())

			// Convert Valuation BigInt to float64
			valFloat, _ := new(big.Float).SetInt(event.Valuation).Float64()

			newProp := models.Property{
				ID:                  uuid.New(),
				OnchainAssetAddress: event.PropertyAsset.Hex(),
				OnchainTokenAddress: event.PropertyToken.Hex(),
				OwnerWallet:         event.Owner.Hex(),
				MetadataHash:        event.PropertyDataHash,
				Valuation:           valFloat,
				Status:              models.StatusActive,
				CreatedAt:           time.Now(),
			}

			if err := database.CreateProperty(newProp); err != nil {
				log.Printf("❌ DB Error saving property: %v", err)
			} else {
				log.Printf("✅ Property saved to DB: %s", newProp.ID)
			}
		}
	}
}

func listenForRevenue(chain *blockchain.ChainService, database *db.Database) {
	sink := make(chan *revenue_distribution.RevenueDistributionRevenueDeposited)

	// Subscribe to "RevenueDeposited" events
	sub, err := chain.RevenueDistribution.WatchRevenueDeposited(nil, sink, nil, nil)
	if err != nil {
		log.Fatalf("Failed to subscribe to revenue events: %v", err)
	}

	log.Println("🎧 Listening for Revenue Deposits...")

	for {
		select {
		case err := <-sub.Err():
			log.Printf("Revenue Subscription error: %v", err)
			time.Sleep(5 * time.Second)
		case event := <-sink:
			log.Printf("📢 Event: Revenue Deposited for Token %s", event.Token.Hex())

			// 1. Find the Property ID based on the Token Address from the event
			prop, err := database.GetPropertyByTokenAddress(event.Token.Hex())
			if err != nil {
				log.Printf("❌ Error finding property for token %s: %v", event.Token.Hex(), err)
				continue
			}

			// 2. Create the Revenue Record
			newDist := models.RevenueDistribution{
				ID:               uuid.New(),
				PropertyID:       prop.ID,
				SnapshotID:       int32(event.SnapshotId.Int64()),
				StablecoinTxHash: event.Raw.TxHash.Hex(),
				TotalAmount:      event.Amount.String(), // Store exact BigInt as string
				CreatedAt:        time.Now(),
			}

			if err := database.CreateRevenueDistribution(newDist); err != nil {
				log.Printf("❌ DB Error saving revenue: %v", err)
			} else {
				log.Printf("✅ Revenue Distribution saved. Snapshot ID: %d", newDist.SnapshotID)
			}
		}
	}
}

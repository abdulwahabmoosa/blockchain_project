package worker

import (
	"backend/blockchain"
	"backend/blockchain/approval_service"
	"backend/blockchain/property_asset"
	"backend/blockchain/property_factory"
	"backend/blockchain/revenue_distribution"
	"backend/db"
	"backend/db/models"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
)

func StartListeners(chain *blockchain.ChainService, database *db.Database) {
	log.Printf("🔄 Starting blockchain event listeners...")

	if chain.PropertyFactory != nil {
		go listenForProperties(chain, database)
	} else {
		log.Printf("⚠️ Skipping property listener - contract not available")
	}

	if chain.RevenueDistribution != nil {
		go listenForRevenue(chain, database)
		go listenForRevenueClaims(chain, database)
	} else {
		log.Printf("⚠️ Skipping revenue listeners - contract not available")
	}

	if chain.Approval != nil {
		go listenForApprovals(chain, database)
	} else {
		log.Printf("⚠️ Skipping approval listener - contract not available")
	}

	log.Printf("✅ Event listeners started (only for available contracts)")
}

func listenForProperties(chain *blockchain.ChainService, database *db.Database) {
	sink := make(chan *property_factory.PropertyFactoryPropertyRegistered)
	sub, err := chain.PropertyFactory.WatchPropertyRegistered(nil, sink, nil)
	if err != nil {
		log.Printf("⚠️ Failed to subscribe to property events: %v", err)
		log.Printf("ℹ️ Property events will not be monitored (RPC may not support subscriptions)")
		return
	}
	log.Println("🎧 Listening for New Properties...")

	for {
		select {
		case err := <-sub.Err():
			log.Printf("Property Subscription error: %v", err)
			time.Sleep(5 * time.Second)
		case event := <-sink:
			log.Printf("📢 Event: Property Registered at %s", event.PropertyAsset.Hex())
			valFloat, _ := new(big.Float).SetInt(event.Valuation).Float64()

			// Fetch property name from PropertyAsset contract
			propertyName := ""
			if chain.Client != nil {
				propertyAssetAddr := common.HexToAddress(event.PropertyAsset.Hex())
				propertyAsset, err := property_asset.NewPropertyAsset(propertyAssetAddr, chain.Client)
				if err == nil {
					name, err := propertyAsset.Name(nil)
					if err == nil {
						propertyName = name
						log.Printf("✅ Fetched property name: %s", propertyName)
					} else {
						log.Printf("⚠️ Failed to fetch property name: %v", err)
					}
				} else {
					log.Printf("⚠️ Failed to connect to PropertyAsset: %v", err)
				}
			}

			newProp := models.Property{
				ID:                  uuid.New(),
				Name:                propertyName,
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
				log.Printf("✅ Property saved: %s", newProp.ID)
			}
		}
	}
}

func listenForRevenue(chain *blockchain.ChainService, database *db.Database) {
	sink := make(chan *revenue_distribution.RevenueDistributionRevenueDeposited)
	sub, err := chain.RevenueDistribution.WatchRevenueDeposited(nil, sink, nil, nil)
	if err != nil {
		log.Printf("⚠️ Failed to subscribe to revenue events: %v", err)
		log.Printf("ℹ️ Revenue events will not be monitored (RPC may not support subscriptions)")
		return
	}
	log.Println("🎧 Listening for Revenue Deposits...")

	for {
		select {
		case err := <-sub.Err():
			log.Printf("Revenue Subscription error: %v", err)
			time.Sleep(5 * time.Second)
		case event := <-sink:
			log.Printf("📢 Event: Revenue Deposited for Token %s", event.Token.Hex())
			prop, err := database.GetPropertyByTokenAddress(event.Token.Hex())
			if err != nil {
				log.Printf("❌ Property not found for token %s", event.Token.Hex())
				continue
			}

			newDist := models.RevenueDistribution{
				ID:               uuid.New(),
				PropertyID:       prop.ID,
				SnapshotID:       int32(event.SnapshotId.Int64()),
				StablecoinTxHash: event.Raw.TxHash.Hex(),
				TotalAmount:      event.Amount.String(),
				CreatedAt:        time.Now(),
			}

			if err := database.CreateRevenueDistribution(newDist); err != nil {
				log.Printf("❌ DB Error saving revenue: %v", err)
			} else {
				log.Printf("✅ Revenue Distribution saved.")
			}
		}
	}
}

func listenForApprovals(chain *blockchain.ChainService, database *db.Database) {
	sink := make(chan *approval_service.ApprovalServiceApproved)
	sub, err := chain.Approval.WatchApproved(nil, sink, nil)
	if err != nil {
		log.Printf("⚠️ Failed to subscribe to approval events: %v", err)
		log.Printf("ℹ️ Approval events will not be monitored (RPC may not support subscriptions)")
		return
	}
	log.Println("🎧 Listening for User Approvals...")

	for {
		select {
		case err := <-sub.Err():
			log.Printf("Approval Subscription error: %v", err)
			time.Sleep(5 * time.Second)
		case event := <-sink:
			userWallet := event.User.Hex()
			log.Printf("📢 Event: User Approved %s", userWallet)

			if err := database.UpdateUserApproval(userWallet, models.ApprovalApproved); err != nil {
				log.Printf("❌ DB Error updating user approval: %v", err)
			}
		}
	}
}

func listenForRevenueClaims(chain *blockchain.ChainService, database *db.Database) {
	sink := make(chan *revenue_distribution.RevenueDistributionRevenueClaimed)
	sub, err := chain.RevenueDistribution.WatchRevenueClaimed(nil, sink, nil, nil)
	if err != nil {
		log.Printf("⚠️ Failed to subscribe to claim events: %v", err)
		log.Printf("ℹ️ Claim events will not be monitored (RPC may not support subscriptions)")
		return
	}
	log.Println("🎧 Listening for Revenue Claims...")

	for {
		select {
		case err := <-sub.Err():
			log.Printf("Claim Subscription error: %v", err)
			time.Sleep(5 * time.Second)
		case event := <-sink:
			// Note: Finding the exact Distribution parent in DB is tricky purely from event data
			// if we don't index distributions by ID on chain easily.
			// For now, we log the raw claim. In a production app, we would query the Distribution ID.

			log.Printf("📢 Event: Revenue Claimed by %s Amount: %s", event.Claimant.Hex(), event.Amount.String())

			newClaim := models.RevenueClaim{
				ID:            uuid.New(),
				WalletAddress: event.Claimant.Hex(),
				Amount:        event.Amount.String(),
				TxHash:        event.Raw.TxHash.Hex(),
				ClaimedAt:     time.Now(),
				// RevenueDistributionID would be linked here if we query distributionId from DB
			}

			if err := database.CreateRevenueClaim(newClaim); err != nil {
				log.Printf("❌ DB Error saving claim: %v", err)
			}
		}
	}
}


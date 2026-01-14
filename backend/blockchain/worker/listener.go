package worker

import (
	"backend/blockchain"
	"backend/blockchain/approval_service"
	"backend/blockchain/property_factory"
	"backend/blockchain/revenue_distribution"
	"backend/db"
	"backend/db/models"
	"log"
	"time"

	"github.com/google/uuid"
)

func StartListeners(chain *blockchain.ChainService, database *db.Database) {
	log.Printf("Info: Starting blockchain event listeners...")

	if chain.PropertyFactory != nil {
		go listenForProperties(chain, database)
	} else {
		log.Printf("Warning: Skipping property listener - contract not available")
	}

	if chain.RevenueDistribution != nil {
		go listenForRevenue(chain, database)
		go listenForRevenueClaims(chain, database)
	} else {
		log.Printf("Warning: Skipping revenue listeners - contract not available")
	}

	if chain.Approval != nil {
		go listenForApprovals(chain, database)
	} else {
		log.Printf("Warning: Skipping approval listener - contract not available")
	}

	log.Printf("Success: Event listeners started (only for available contracts)")
}

func listenForProperties(chain *blockchain.ChainService, database *db.Database) {
	sink := make(chan *property_factory.PropertyFactoryPropertyRegistered)
	sub, err := chain.PropertyFactory.WatchPropertyRegistered(nil, sink, nil)
	if err != nil {
		log.Printf("Warning: Failed to subscribe to property events: %v", err)
		log.Printf("Info: Property events will not be monitored (RPC may not support subscriptions)")
		return
	}
	log.Println("Info: Listening for New Properties...")

	for {
		select {
		case err := <-sub.Err():
			log.Printf("Property Subscription error: %v", err)
			time.Sleep(5 * time.Second)
		case event := <-sink:
			log.Printf("Info: Event: Property Registered at %s", event.PropertyAsset.Hex())
			
			// Check if property already exists by asset address OR token address to prevent duplicates
			assetAddr := event.PropertyAsset.Hex()
			tokenAddr := event.PropertyToken.Hex()
			
			// Check by asset address first
			existingProp, err := database.GetPropertyByAssetAddress(assetAddr)
			if err == nil {
				log.Printf("Warning: Property already exists with asset address %s (ID: %s), skipping duplicate creation", assetAddr, existingProp.ID)
				continue
			}
			
			// Also check by token address as fallback
			existingProp, err = database.GetPropertyByTokenAddress(tokenAddr)
			if err == nil {
				log.Printf("Warning: Property already exists with token address %s (ID: %s), skipping duplicate creation", tokenAddr, existingProp.ID)
				continue
			}

			// Property not found in database - this event listener only checks for duplicates
// Property creation is handled by the API endpoints (ApprovePropertyUploadRequest, CreateProperty)
// This listener serves as a backup duplicate prevention mechanism onlyy
			log.Printf("Info: Property event detected but property not found in database. Property creation should be handled by API endpoints.")
		}
	}
}

func listenForRevenue(chain *blockchain.ChainService, database *db.Database) {
	sink := make(chan *revenue_distribution.RevenueDistributionRevenueDeposited)
	sub, err := chain.RevenueDistribution.WatchRevenueDeposited(nil, sink, nil, nil)
	if err != nil {
		log.Printf("Warning: Failed to subscribe to revenue events: %v", err)
		log.Printf("Info: Revenue events will not be monitored (RPC may not support subscriptions)")
		return
	}
	log.Println("Info: Listening for Revenue Deposits...")

	for {
		select {
		case err := <-sub.Err():
			log.Printf("Revenue Subscription error: %v", err)
			time.Sleep(5 * time.Second)
		case event := <-sink:
			log.Printf("Info: Event: Revenue Deposited for Token %s", event.Token.Hex())
			prop, err := database.GetPropertyByTokenAddress(event.Token.Hex())
			if err != nil {
				log.Printf("Error: Property not found for token %s", event.Token.Hex())
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
				log.Printf("Error: DB Error saving revenue: %v", err)
			} else {
				log.Printf("Success: Revenue Distribution saved.")
			}
		}
	}
}

func listenForApprovals(chain *blockchain.ChainService, database *db.Database) {
	sink := make(chan *approval_service.ApprovalServiceApproved)
	sub, err := chain.Approval.WatchApproved(nil, sink, nil)
	if err != nil {
		log.Printf("Warning: Failed to subscribe to approval events: %v", err)
		log.Printf("Info: Approval events will not be monitored (RPC may not support subscriptions)")
		return
	}
	log.Println("Info: Listening for User Approvals...")

	for {
		select {
		case err := <-sub.Err():
			log.Printf("Approval Subscription error: %v", err)
			time.Sleep(5 * time.Second)
		case event := <-sink:
			userWallet := event.User.Hex()
			log.Printf("Info: Event: User Approved %s", userWallet)

			if err := database.UpdateUserApproval(userWallet, models.ApprovalApproved); err != nil {
				log.Printf("Error: DB Error updating user approval: %v", err)
			}
		}
	}
}

func listenForRevenueClaims(chain *blockchain.ChainService, database *db.Database) {
	sink := make(chan *revenue_distribution.RevenueDistributionRevenueClaimed)
	sub, err := chain.RevenueDistribution.WatchRevenueClaimed(nil, sink, nil, nil)
	if err != nil {
		log.Printf("Warning: Failed to subscribe to claim events: %v", err)
		log.Printf("Info: Claim events will not be monitored (RPC may not support subscriptions)")
		return
	}
	log.Println("Info: Listening for Revenue Claims...")

	for {
		select {
		case err := <-sub.Err():
			log.Printf("Claim Subscription error: %v", err)
			time.Sleep(5 * time.Second)
		case event := <-sink:
			// Note: Finding the exact Distribution parent in DB is tricky purely from event data
			// if we don't index distributions by ID on chain easily.
			// For now, we log the raw claim. In a production app, we would query the Distribution ID.

			log.Printf("Info: Event: Revenue Claimed by %s Amount: %s", event.Claimant.Hex(), event.Amount.String())

			newClaim := models.RevenueClaim{
				ID:            uuid.New(),
				WalletAddress: event.Claimant.Hex(),
				Amount:        event.Amount.String(),
				TxHash:        event.Raw.TxHash.Hex(),
				ClaimedAt:     time.Now(),
				// RevenueDistributionID would be linked here if we query distributionId from DB
			}

			if err := database.CreateRevenueClaim(newClaim); err != nil {
				log.Printf("Error: DB Error saving claim: %v", err)
			}
		}
	}
}

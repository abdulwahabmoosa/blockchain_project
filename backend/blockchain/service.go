package blockchain

import (
	"backend/blockchain/approval_service"
	"backend/blockchain/platform_registry"
	"backend/blockchain/property_asset"
	"backend/blockchain/property_factory"
	"backend/blockchain/property_token"
	"backend/blockchain/revenue_distribution"
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// ChainService - main struct for blockchain operations
// contains client connection and contract instances
type ChainService struct {
	Client     *ethclient.Client
	PrivateKey *ecdsa.PrivateKey
	ChainID    *big.Int

	// Contract instance
	Registry            *platform_registry.PlatformRegistry
	PropertyFactory     *property_factory.PropertyFactory
	Approval            *approval_service.ApprovalService
	RevenueDistribution *revenue_distribution.RevenueDistribution
}

// NewChainServiceEnv - create service from environment variables
// loads RPC URL, private key, registry address from env
func NewChainServiceEnv() (*ChainService, error) {
	rpcURL := os.Getenv("SEPOLIA_RPC")
	if rpcURL == "" {
		rpcURL = "http://127.0.0.1:8545" // Use HTTP for local development
	}

	privateKey := os.Getenv("PRIVATE_KEY_2")
	if privateKey == "" {
		log.Printf("Warning PRIVATE_KEY_2 not found, trying PRIVATE_KEY...")
		privateKey = os.Getenv("PRIVATE_KEY")
		if privateKey == "" {
			return nil, errors.New("PRIVATE_KEY_2 or PRIVATE_KEY environment variable is required")
		}
	}

	registryAddr := os.Getenv("REGISTRY")
	if registryAddr == "" {
		return nil, errors.New("REGISTRY environment variable is required")
	}

	return NewChainService(rpcURL, privateKey, registryAddr)
}

func NewChainService(rpcUrl, privateKeyHex, registryAddr string) (*ChainService, error) {
	log.Printf("Connecting to blockchain at: %s", rpcUrl)
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to blockchain: %v", err)
	}

	privateKeyHex = strings.TrimPrefix(privateKeyHex, "0x")
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %v", err)
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get network ID: %v", err)
	}

	log.Printf("Connected to network: %s (Chain ID: %s)", rpcUrl, chainID.String())
	log.Printf("Using registry contract: %s", registryAddr)

	registryAddress := common.HexToAddress(registryAddr)
	registry, err := platform_registry.NewPlatformRegistry(registryAddress, client)
	if err != nil {
		log.Printf("Warning: Failed to connect to registry contract at %s: %v", registryAddr, err)
		log.Printf("Warning: Blockchain service will work in limited mode - contracts not deployed yet")
		return &ChainService{
			Client:     client,
			PrivateKey: privateKey,
			ChainID:    chainID,
			// Contract instances will be nil, methods will need to handle this
		}, nil
	}

	factoryAddr, err := registry.GetPropertyFactory(nil)
	if err != nil {
		log.Printf("Warning: Failed to get factory address from registry: %v", err)
		return &ChainService{
			Client:     client,
			PrivateKey: privateKey,
			ChainID:    chainID,
			Registry:   registry,
			// Other contracts will be nil
		}, nil
	}
	factory, err := property_factory.NewPropertyFactory(factoryAddr, client)
	if err != nil {
		log.Printf("Warning: Failed to connect to factory contract: %v", err)
		return &ChainService{
			Client:     client,
			PrivateKey: privateKey,
			ChainID:    chainID,
			Registry:   registry,
			// Factory and other contracts will be nil
		}, nil
	}

	approvalAddr, err := registry.GetApprovalService(nil)
	if err != nil {
		log.Printf("Warning: Failed to get approval service address from registry: %v", err)
		return &ChainService{
			Client:          client,
			PrivateKey:      privateKey,
			ChainID:         chainID,
			Registry:        registry,
			PropertyFactory: factory,
			// Approval and other contracts will be nil
		}, nil
	}
	approval, err := approval_service.NewApprovalService(approvalAddr, client)
	if err != nil {
		log.Printf("Warning: Failed to connect to approval service contract: %v", err)
		return &ChainService{
			Client:          client,
			PrivateKey:      privateKey,
			ChainID:         chainID,
			Registry:        registry,
			PropertyFactory: factory,
			// Approval and other contracts will be nil
		}, nil
	}

	revenueAddr, err := registry.GetRevenueDistribution(nil)
	if err != nil {
		log.Printf("Warning: Failed to get revenue distribution address from registry: %v", err)
		return &ChainService{
			Client:          client,
			PrivateKey:      privateKey,
			ChainID:         chainID,
			Registry:        registry,
			PropertyFactory: factory,
			Approval:        approval,
			// Revenue contract will be nil
		}, nil
	}
	revenue, err := revenue_distribution.NewRevenueDistribution(revenueAddr, client)
	if err != nil {
		log.Printf("Warning: Failed to connect to revenue distribution contract: %v", err)
		return &ChainService{
			Client:          client,
			PrivateKey:      privateKey,
			ChainID:         chainID,
			Registry:        registry,
			PropertyFactory: factory,
			Approval:        approval,
			// Revenue contract will be nil
		}, nil
	}

	return &ChainService{
		Client:              client,
		PrivateKey:          privateKey,
		ChainID:             chainID,
		Registry:            registry,
		PropertyFactory:     factory,
		Approval:            approval,
		RevenueDistribution: revenue,
	}, nil
}

func (s *ChainService) GetTransactor() (*bind.TransactOpts, error) {
	auth, err := bind.NewKeyedTransactorWithChainID(s.PrivateKey, s.ChainID)
	if err != nil {
		return nil, err
	}
	log.Printf("Backend using signer address: %s", auth.From.Hex())
	return auth, nil
}

// WaitForTx waits for a transaction to be mined and returns the receipt
func (s *ChainService) WaitForTx(txHash common.Hash) (*types.Receipt, error) {
	// Wait for the transaction to be mined
	for {
		receipt, err := s.Client.TransactionReceipt(context.Background(), txHash)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				// Transaction not yet mined, wait and try again
				time.Sleep(2 * time.Second)
				continue
			}
			return nil, err
		}
		return receipt, nil
	}
}

// PropertyCreationResult holds the result of creating a property on-chain
type PropertyCreationResult struct {
	AssetAddress string
	TokenAddress string
	TxHash       string
	PropertyName string
}

// CreateProperty - deploy new property contracts on blockchain
// creates PropertyAsset (NFT) and PropertyToken (ERC20), links them
func (s *ChainService) CreateProperty(ownerStr, name, symbol, dataHash string, valuation, supply int64) (*PropertyCreationResult, error) {
	if s.PropertyFactory == nil {
		log.Printf("Warning: Property factory contract not available - blockchain service in limited mode")
		return nil, fmt.Errorf("property factory contract not deployed - deploy contracts to enable property creation")
	}

	auth, err := s.GetTransactor()
	if err != nil {
		return nil, err
	}

	// Set higher gas limit for contract deployments
	auth.GasLimit = 8000000 // 8 million gas

	owner := common.HexToAddress(ownerStr)

	// Convert valuation from ETH to wei (multiply by 10^18)
	valBig := big.NewInt(valuation)
	weiMultiplier := big.NewInt(1000000000000000000) // 10^18
	valBig.Mul(valBig, weiMultiplier)

	// Convert token supply to wei (multiply by 10^18 for 18 decimal tokens)
	supplyBig := big.NewInt(supply)
	supplyBig.Mul(supplyBig, weiMultiplier)

	log.Printf("Submitting property creation transaction to blockchain...")

	// Submit transaction
	tx, err := s.PropertyFactory.CreateProperty(auth, owner, name, symbol, dataHash, valBig, supplyBig, name+" Token", "TKN")
	if err != nil {
		return nil, fmt.Errorf("failed to submit transaction: %v", err)
	}

	log.Printf("Transaction submitted: %s", tx.Hash().Hex())
	log.Printf("Waiting for transaction to be mined...")

	// Wait for transaction to be mined
	receipt, err := s.WaitForTx(tx.Hash())
	if err != nil {
		return nil, fmt.Errorf("failed to wait for transaction: %v", err)
	}

	// Check if transaction succeeded
	if receipt.Status != types.ReceiptStatusSuccessful {
		return nil, fmt.Errorf("transaction failed on-chain with status: %d", receipt.Status)
	}

	log.Printf("Transaction mined successfully in block: %d", receipt.BlockNumber)

	// Parse PropertyRegistered event from logs
	for _, vLog := range receipt.Logs {
		event, err := s.PropertyFactory.ParsePropertyRegistered(*vLog)
		if err != nil {
			continue // Not the event we're looking for
		}

		log.Printf("PropertyRegistered event found:")
		log.Printf("   Owner: %s", event.Owner.Hex())
		log.Printf("   Asset Address: %s", event.PropertyAsset.Hex())
		log.Printf("   Token Address: %s", event.PropertyToken.Hex())
		log.Printf("   Metadata Hash: %s", event.PropertyDataHash)

		return &PropertyCreationResult{
			AssetAddress: event.PropertyAsset.Hex(),
			TokenAddress: event.PropertyToken.Hex(),
			TxHash:       tx.Hash().Hex(),
			PropertyName: name,
		}, nil
	}

	return nil, fmt.Errorf("PropertyRegistered event not found in transaction logs")
}

// DistributeRevenue deposits funds into the Revenue contract
func (s *ChainService) DistributeRevenue(tokenAddrStr, stablecoinAddrStr string, amount int64) (*types.Transaction, error) {
	if s.RevenueDistribution == nil {
		log.Printf("Warning: Revenue distribution contract not available - blockchain service in limited mode")
		return nil, fmt.Errorf("revenue distribution contract not deployed - deploy contracts to enable revenue distribution")
	}

	auth, err := s.GetTransactor()
	if err != nil {
		return nil, err
	}

	tokenAddr := common.HexToAddress(tokenAddrStr)
	stablecoinAddr := common.HexToAddress(stablecoinAddrStr)
	amountBig := big.NewInt(amount)

	// Get RevenueDistribution contract address
	revenueDistributionAddr, err := s.Registry.GetRevenueDistribution(nil)
	if err != nil {
		log.Printf("Warning: Failed to get RevenueDistribution address from registry: %v", err)
		// Continue anyway - might work if role is already granted
	} else {
		// Check if RevenueDistribution has SNAPSHOT_ROLE on the PropertyToken
		// If not, grant it automatically
		err = s.ensureSnapshotRole(tokenAddr, revenueDistributionAddr, auth)
		if err != nil {
			log.Printf("Warning: Failed to ensure SNAPSHOT_ROLE (will try distribution anyway): %v", err)
			// Continue anyway - the error might be that role is already granted
		}
	}

	return s.RevenueDistribution.DepositRevenue(auth, tokenAddr, stablecoinAddr, amountBig)
}

// ensureSnapshotRole ensures that RevenueDistribution has SNAPSHOT_ROLE on the PropertyToken
func (s *ChainService) ensureSnapshotRole(tokenAddr, revenueDistributionAddr common.Address, auth *bind.TransactOpts) error {
	// Create PropertyToken instance
	token, err := property_token.NewPropertyToken(tokenAddr, s.Client)
	if err != nil {
		return fmt.Errorf("failed to create PropertyToken instance: %v", err)
	}

	// Get SNAPSHOT_ROLE bytes32
	snapshotRole, err := token.SNAPSHOTROLE(nil)
	if err != nil {
		return fmt.Errorf("failed to get SNAPSHOT_ROLE: %v", err)
	}

	// Check if RevenueDistribution already has the role
	hasRole, err := token.HasRole(nil, snapshotRole, revenueDistributionAddr)
	if err != nil {
		return fmt.Errorf("failed to check role: %v", err)
	}

	if hasRole {
		log.Printf("RevenueDistribution already has SNAPSHOT_ROLE on token %s", tokenAddr.Hex())
		return nil
	}

	log.Printf("Info: RevenueDistribution does not have SNAPSHOT_ROLE on token %s, attempting to grant...", tokenAddr.Hex())

	// PropertyFactory is the admin of PropertyTokens it creats
	// We need to use PropertyFactory's grantSnapshotRoleToRevenue function
	// Since it's not in bindings, we'll use PropertyFactory's GrantRole on the PropertyToken
	// But first, let's try calling PropertyFactory's grantSnapshotRoleToRevenue via raw transaction
	
	// Try using PropertyFactory to grant the role
	// PropertyFactory.grantSnapshotRoleToRevenue(tokenAddress) calls token.grantRole internally
	// Since PropertyFactory is admin, it can grant roles on tokens it created
	
	// For now, try to grant directly on the token
	// If the backend wallet is admin of PropertyFactory, and PropertyFactory is admin of token,
	// we might need to use PropertyFactory's function. But since bindings don't have it,
	// we'll try direct grant first (might work if backend wallet has admin role on token)
	
	// Convert snapshotRole from [32]byte to [32]byte for grantRole
	var roleBytes [32]byte
	copy(roleBytes[:], snapshotRole[:])
	
	// Try to grant the role directly on the PropertyToken first
	// This will work if the caller (backend wallet) is the admin of the PropertyToken
	tx, err := token.GrantRole(auth, roleBytes, revenueDistributionAddr)
	if err != nil {
		// Direct grant failed - PropertyFactory is likely the admin
		// Use PropertyFactory's grantSnapshotRoleToRevenue function via raw transaction
		log.Printf("Warning: Direct grant failed (PropertyFactory is admin). Using PropertyFactory.grantSnapshotRoleToRevenue...")
		
		if s.PropertyFactory == nil {
			return fmt.Errorf("PropertyFactory not available to grant role")
		}
		
		// Use PropertyFactory's raw Transact method to call grantSnapshotRoleToRevenue
		// Function signature: grantSnapshotRoleToRevenue(address tokenAddress)
		// PropertyFactoryRaw allows calling methods by name even if not in bindings
		rawFactory := &property_factory.PropertyFactoryRaw{Contract: s.PropertyFactory}
		tx, err = rawFactory.Transact(auth, "grantSnapshotRoleToRevenue", tokenAddr)
		if err != nil {
			log.Printf("Failed to call PropertyFactory.grantSnapshotRoleToRevenue: %v", err)
			return fmt.Errorf("failed to grant SNAPSHOT_ROLE via PropertyFactory: %v", err)
		}
		
		log.Printf("PropertyFactory.grantSnapshotRoleToRevenue transaction sent, waiting for confirmation...")
		receipt, err := s.WaitForTx(tx.Hash())
		if err != nil {
			return fmt.Errorf("failed to wait for grant role transaction: %v", err)
		}
		
		if receipt.Status == 0 {
			return fmt.Errorf("grant role transaction failed")
		}
		
		log.Printf("SNAPSHOT_ROLE successfully granted to RevenueDistribution via PropertyFactory")
		return nil
	}
	
	log.Printf("SNAPSHOT_ROLE grant transaction sent, waiting for confirmation...")
	receipt, err := s.WaitForTx(tx.Hash())
	if err != nil {
		return fmt.Errorf("failed to wait for grant role transaction: %v", err)
	}
	
	if receipt.Status == 0 {
		return fmt.Errorf("grant role transaction failed")
	}
	
	log.Printf("SNAPSHOT_ROLE successfully granted to RevenueDistribution on token %s", tokenAddr.Hex())
	return nil
}

// ApproveUser allows a specific wallet address to participate in the platform
func (s *ChainService) ApproveUser(userAddressStr string) (*types.Transaction, error) {
	log.Printf("Starting approval for address: %s", userAddressStr)

	if s.Approval == nil {
		log.Printf("Warning: Approval contract not available - blockchain service in limited mode")
		return nil, fmt.Errorf("approval contract not deployed - deploy contracts to enable blockchain functionality")
	}

	auth, err := s.GetTransactor()
	if err != nil {
		log.Printf("Failed to get transactor: %v", err)
		return nil, fmt.Errorf("failed to get transactor: %v", err)
	}

	userAddr := common.HexToAddress(userAddressStr)
	log.Printf("Debug: Calling smart contract Approve() for address: %s", userAddressStr)
	log.Printf("Backend signer address: %s", auth.From.Hex())

	tx, err := s.Approval.Approve(auth, userAddr)
	if err != nil {
		log.Printf("Smart contract Approve() failed for %s: %v", userAddressStr, err)
		log.Printf("Debug: Auth signer address: %s", auth.From.Hex())
		return nil, fmt.Errorf("smart contract approve failed: %v", err)
	}

	log.Printf("Smart contract approve() transaction sent, hash: %s", tx.Hash().Hex())
	log.Printf("Waiting for transaction confirmation...")

	// Wait for transaction to be mined
	receipt, err := s.WaitForTx(tx.Hash())
	if err != nil {
		log.Printf("Warning: Transaction confirmation failed: %v", err)
		return tx, fmt.Errorf("transaction confirmation failed: %v", err)
	}

	log.Printf("Transaction mined in block: %d", receipt.BlockNumber)
	return tx, nil
}

// IsApproved checks if a user address is approved
func (s *ChainService) IsApproved(userAddressStr string) (bool, error) {
	if s.Approval == nil {
		log.Printf("Warning: Approval contract not available - returning false for approval check")
		return false, nil // Return false if contract not available
	}

	userAddr := common.HexToAddress(userAddressStr)
	approved, err := s.Approval.Check(nil, userAddr)
	if err != nil {
		return false, fmt.Errorf("failed to check approval status: %v", err)
	}
	return approved, nil
}

// ApproveProperty sets a property status to Active on-chain
func (s *ChainService) ApproveProperty(propertyAssetAddrStr string) (*types.Transaction, error) {
	if s.Client == nil {
		return nil, fmt.Errorf("blockchain client not available")
	}

	propertyAssetAddr := common.HexToAddress(propertyAssetAddrStr)
	propertyAsset, err := property_asset.NewPropertyAsset(propertyAssetAddr, s.Client)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to property asset contract: %v", err)
	}

	auth, err := s.GetTransactor()
	if err != nil {
		return nil, err
	}

	// Status.Active = 0
	return propertyAsset.SetStatus(auth, 0)
}

// RejectProperty sets a property status to Closed on-chain
func (s *ChainService) RejectProperty(propertyAssetAddrStr string) (*types.Transaction, error) {
	if s.Client == nil {
		return nil, fmt.Errorf("blockchain client not available")
	}

	propertyAssetAddr := common.HexToAddress(propertyAssetAddrStr)
	propertyAsset, err := property_asset.NewPropertyAsset(propertyAssetAddr, s.Client)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to property asset contract: %v", err)
	}

	auth, err := s.GetTransactor()
	if err != nil {
		return nil, err
	}

	// Status.Closed = 3
	return propertyAsset.SetStatus(auth, 3)
}

// GetTokenBalance gets the token balance for a wallet address
func (s *ChainService) GetTokenBalance(tokenAddrStr, walletAddrStr string) (*big.Int, error) {
	if s.Client == nil {
		return nil, fmt.Errorf("blockchain client not available")
	}

	tokenAddr := common.HexToAddress(tokenAddrStr)
	walletAddr := common.HexToAddress(walletAddrStr)

	token, err := property_token.NewPropertyToken(tokenAddr, s.Client)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to property token contract: %v", err)
	}

	balance, err := token.BalanceOf(nil, walletAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to get token balance: %v", err)
	}

	return balance, nil
}

// TransferTokens transfers tokens from the backend wallet to another address
// Note: This requires the backend wallet to have tokens and approval
func (s *ChainService) TransferTokens(tokenAddrStr, toAddrStr string, amount *big.Int) (*types.Transaction, error) {
	if s.Client == nil {
		return nil, fmt.Errorf("blockchain client not available")
	}

	tokenAddr := common.HexToAddress(tokenAddrStr)
	toAddr := common.HexToAddress(toAddrStr)

	token, err := property_token.NewPropertyToken(tokenAddr, s.Client)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to property token contract: %v", err)
	}

	auth, err := s.GetTransactor()
	if err != nil {
		return nil, err
	}

	return token.Transfer(auth, toAddr, amount)
}

// GetTotalSupply gets the total token supply for a property token contract
func (s *ChainService) GetTotalSupply(tokenAddrStr string) (*big.Int, error) {
	if s.Client == nil {
		return nil, fmt.Errorf("blockchain client not available")
	}

	tokenAddr := common.HexToAddress(tokenAddrStr)

	token, err := property_token.NewPropertyToken(tokenAddr, s.Client)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to property token contract: %v", err)
	}

	totalSupply, err := token.TotalSupply(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get total supply: %v", err)
	}

	return totalSupply, nil
}

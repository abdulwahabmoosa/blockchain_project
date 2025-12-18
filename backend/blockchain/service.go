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

type ChainService struct {
	Client     *ethclient.Client
	PrivateKey *ecdsa.PrivateKey
	ChainID    *big.Int

	// Contract Instances
	Registry            *platform_registry.PlatformRegistry
	PropertyFactory     *property_factory.PropertyFactory
	Approval            *approval_service.ApprovalService
	RevenueDistribution *revenue_distribution.RevenueDistribution
}

func NewChainServiceEnv() (*ChainService, error) {
	rpcURL := os.Getenv("SEPOLIA_RPC")
	if rpcURL == "" {
		rpcURL = "http://127.0.0.1:8545" // Use HTTP for local development
	}

	privateKey := os.Getenv("PRIVATE_KEY_2")
	if privateKey == "" {
		log.Printf("⚠️ PRIVATE_KEY_2 not found, trying PRIVATE_KEY...")
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
	log.Printf("🔗 Connecting to blockchain at: %s", rpcUrl)
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

	log.Printf("✅ Connected to network: %s (Chain ID: %s)", rpcUrl, chainID.String())
	log.Printf("📋 Using registry contract: %s", registryAddr)

	registryAddress := common.HexToAddress(registryAddr)
	registry, err := platform_registry.NewPlatformRegistry(registryAddress, client)
	if err != nil {
		log.Printf("⚠️ Failed to connect to registry contract at %s: %v", registryAddr, err)
		log.Printf("⚠️ Blockchain service will work in limited mode - contracts not deployed yet")
		return &ChainService{
			Client:     client,
			PrivateKey: privateKey,
			ChainID:    chainID,
			// Contract instances will be nil, methods will need to handle this
		}, nil
	}

	factoryAddr, err := registry.GetPropertyFactory(nil)
	if err != nil {
		log.Printf("⚠️ Failed to get factory address from registry: %v", err)
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
		log.Printf("⚠️ Failed to connect to factory contract: %v", err)
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
		log.Printf("⚠️ Failed to get approval service address from registry: %v", err)
		return &ChainService{
			Client:     client,
			PrivateKey: privateKey,
			ChainID:    chainID,
			Registry:   registry,
			PropertyFactory: factory,
			// Approval and other contracts will be nil
		}, nil
	}
	approval, err := approval_service.NewApprovalService(approvalAddr, client)
	if err != nil {
		log.Printf("⚠️ Failed to connect to approval service contract: %v", err)
		return &ChainService{
			Client:     client,
			PrivateKey: privateKey,
			ChainID:    chainID,
			Registry:   registry,
			PropertyFactory: factory,
			// Approval and other contracts will be nil
		}, nil
	}

	revenueAddr, err := registry.GetRevenueDistribution(nil)
	if err != nil {
		log.Printf("⚠️ Failed to get revenue distribution address from registry: %v", err)
		return &ChainService{
			Client:     client,
			PrivateKey: privateKey,
			ChainID:    chainID,
			Registry:   registry,
			PropertyFactory: factory,
			Approval:   approval,
			// Revenue contract will be nil
		}, nil
	}
	revenue, err := revenue_distribution.NewRevenueDistribution(revenueAddr, client)
	if err != nil {
		log.Printf("⚠️ Failed to connect to revenue distribution contract: %v", err)
		return &ChainService{
			Client:     client,
			PrivateKey: privateKey,
			ChainID:    chainID,
			Registry:   registry,
			PropertyFactory: factory,
			Approval:   approval,
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
	log.Printf("🔑 Backend using signer address: %s", auth.From.Hex())
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

// CreateProperty deploys a new property via the Factory
func (s *ChainService) CreateProperty(ownerStr, name, symbol, dataHash string, valuation, supply int64) (*types.Transaction, error) {
	if s.PropertyFactory == nil {
		log.Printf("⚠️ Property factory contract not available - blockchain service in limited mode")
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

	// tokenName and tokenSymbol are derived for simplicity here
	return s.PropertyFactory.CreateProperty(auth, owner, name, symbol, dataHash, valBig, supplyBig, name+" Token", "TKN")
}

// DistributeRevenue deposits funds into the Revenue contract
func (s *ChainService) DistributeRevenue(tokenAddrStr, stablecoinAddrStr string, amount int64) (*types.Transaction, error) {
	if s.RevenueDistribution == nil {
		log.Printf("⚠️ Revenue distribution contract not available - blockchain service in limited mode")
		return nil, fmt.Errorf("revenue distribution contract not deployed - deploy contracts to enable revenue distribution")
	}

	auth, err := s.GetTransactor()
	if err != nil {
		return nil, err
	}

	tokenAddr := common.HexToAddress(tokenAddrStr)
	stablecoinAddr := common.HexToAddress(stablecoinAddrStr)
	amountBig := big.NewInt(amount)

	return s.RevenueDistribution.DepositRevenue(auth, tokenAddr, stablecoinAddr, amountBig)
}

// ApproveUser allows a specific wallet address to participate in the platform
func (s *ChainService) ApproveUser(userAddressStr string) (*types.Transaction, error) {
	log.Printf("🔄 Starting approval for address: %s", userAddressStr)

	if s.Approval == nil {
		log.Printf("⚠️ Approval contract not available - blockchain service in limited mode")
		return nil, fmt.Errorf("approval contract not deployed - deploy contracts to enable blockchain functionality")
	}

	auth, err := s.GetTransactor()
	if err != nil {
		log.Printf("❌ Failed to get transactor: %v", err)
		return nil, fmt.Errorf("failed to get transactor: %v", err)
	}

	userAddr := common.HexToAddress(userAddressStr)
	log.Printf("🔗 Calling smart contract Approve() for address: %s", userAddressStr)
	log.Printf("🔑 Backend signer address: %s", auth.From.Hex())

	tx, err := s.Approval.Approve(auth, userAddr)
	if err != nil {
		log.Printf("❌ Smart contract Approve() failed for %s: %v", userAddressStr, err)
		log.Printf("🔍 Auth signer address: %s", auth.From.Hex())
		return nil, fmt.Errorf("smart contract approve failed: %v", err)
	}

	log.Printf("✅ Smart contract approve() transaction sent, hash: %s", tx.Hash().Hex())
	log.Printf("⏳ Waiting for transaction confirmation...")

	// Wait for transaction to be mined
	receipt, err := s.WaitForTx(tx.Hash())
	if err != nil {
		log.Printf("⚠️ Transaction confirmation failed: %v", err)
		return tx, fmt.Errorf("transaction confirmation failed: %v", err)
	}

	log.Printf("✅ Transaction mined in block: %d", receipt.BlockNumber)
	return tx, nil
}

// IsApproved checks if a user address is approved
func (s *ChainService) IsApproved(userAddressStr string) (bool, error) {
	if s.Approval == nil {
		log.Printf("⚠️ Approval contract not available - returning false for approval check")
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

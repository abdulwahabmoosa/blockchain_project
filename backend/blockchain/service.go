package blockchain

import (
	"backend/blockchain/approval_service"
	"backend/blockchain/platform_registry"
	"backend/blockchain/property_factory"
	"backend/blockchain/revenue_distribution"
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"
	"os"
	"strings"

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
		rpcURL = "ws://127.0.0.1:8545" // Must use WS for event subscriptions
	}

	privateKey := os.Getenv("PRIVATE_KEY")
	if privateKey == "" {
		return nil, errors.New("PRIVATE_KEY environment variable is required")
	}

	registryAddr := os.Getenv("REGISTRY")
	if registryAddr == "" {
		return nil, errors.New("REGISTRY environment variable is required")
	}

	return NewChainService(rpcURL, privateKey, registryAddr)
}

func NewChainService(rpcUrl, privateKeyHex, registryAddr string) (*ChainService, error) {
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		return nil, err
	}

	privateKeyHex = strings.TrimPrefix(privateKeyHex, "0x")
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, err
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return nil, err
	}

	registryAddress := common.HexToAddress(registryAddr)
	registry, err := platform_registry.NewPlatformRegistry(registryAddress, client)
	if err != nil {
		return nil, err
	}

	factoryAddr, err := registry.GetPropertyFactory(nil)
	if err != nil {
		return nil, err
	}
	factory, err := property_factory.NewPropertyFactory(factoryAddr, client)
	if err != nil {
		return nil, err
	}

	approvalAddr, err := registry.GetApprovalService(nil)
	if err != nil {
		return nil, err
	}
	approval, err := approval_service.NewApprovalService(approvalAddr, client)
	if err != nil {
		return nil, err
	}

	revenueAddr, err := registry.GetRevenueDistribution(nil)
	if err != nil {
		return nil, err
	}
	revenue, err := revenue_distribution.NewRevenueDistribution(revenueAddr, client)
	if err != nil {
		return nil, err
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
	return auth, nil
}

// CreateProperty deploys a new property via the Factory
func (s *ChainService) CreateProperty(ownerStr, name, symbol, dataHash string, valuation, supply int64) (*types.Transaction, error) {
	auth, err := s.GetTransactor()
	if err != nil {
		return nil, err
	}

	owner := common.HexToAddress(ownerStr)
	valBig := big.NewInt(valuation)
	supplyBig := big.NewInt(supply)

	// tokenName and tokenSymbol are derived for simplicity here
	return s.PropertyFactory.CreateProperty(auth, owner, name, symbol, dataHash, valBig, supplyBig, name+" Token", "TKN")
}

// DistributeRevenue deposits funds into the Revenue contract
func (s *ChainService) DistributeRevenue(tokenAddrStr, stablecoinAddrStr string, amount int64) (*types.Transaction, error) {
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
	auth, err := s.GetTransactor()
	if err != nil {
		return nil, err
	}

	userAddr := common.HexToAddress(userAddressStr)
	return s.Approval.Approve(auth, userAddr)
}

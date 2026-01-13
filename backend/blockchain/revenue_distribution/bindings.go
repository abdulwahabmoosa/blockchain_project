// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package revenue_distribution

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// RevenueDistributionMetaData contains all meta data concerning the RevenueDistribution contract.
var RevenueDistributionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"distributionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"claimant\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RevenueClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"distributionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"snapshotId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"stablecoin\",\"type\":\"address\"}],\"name\":\"RevenueDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DISTRIBUTOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"distributionId\",\"type\":\"uint256\"}],\"name\":\"claimRevenue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"claimed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stablecoin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositRevenue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"distributions\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"snapshotId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stablecoin\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"withdrawn\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"distributionsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// RevenueDistributionABI is the input ABI used to generate the binding from.
// Deprecated: Use RevenueDistributionMetaData.ABI instead.
var RevenueDistributionABI = RevenueDistributionMetaData.ABI

// RevenueDistribution is an auto generated Go binding around an Ethereum contract.
type RevenueDistribution struct {
	RevenueDistributionCaller     // Read-only binding to the contract
	RevenueDistributionTransactor // Write-only binding to the contract
	RevenueDistributionFilterer   // Log filterer for contract events
}

// RevenueDistributionCaller is an auto generated read-only Go binding around an Ethereum contract.
type RevenueDistributionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RevenueDistributionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RevenueDistributionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RevenueDistributionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RevenueDistributionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RevenueDistributionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RevenueDistributionSession struct {
	Contract     *RevenueDistribution // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// RevenueDistributionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RevenueDistributionCallerSession struct {
	Contract *RevenueDistributionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// RevenueDistributionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RevenueDistributionTransactorSession struct {
	Contract     *RevenueDistributionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// RevenueDistributionRaw is an auto generated low-level Go binding around an Ethereum contract.
type RevenueDistributionRaw struct {
	Contract *RevenueDistribution // Generic contract binding to access the raw methods on
}

// RevenueDistributionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RevenueDistributionCallerRaw struct {
	Contract *RevenueDistributionCaller // Generic read-only contract binding to access the raw methods on
}

// RevenueDistributionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RevenueDistributionTransactorRaw struct {
	Contract *RevenueDistributionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRevenueDistribution creates a new instance of RevenueDistribution, bound to a specific deployed contract.
func NewRevenueDistribution(address common.Address, backend bind.ContractBackend) (*RevenueDistribution, error) {
	contract, err := bindRevenueDistribution(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RevenueDistribution{RevenueDistributionCaller: RevenueDistributionCaller{contract: contract}, RevenueDistributionTransactor: RevenueDistributionTransactor{contract: contract}, RevenueDistributionFilterer: RevenueDistributionFilterer{contract: contract}}, nil
}

// NewRevenueDistributionCaller creates a new read-only instance of RevenueDistribution, bound to a specific deployed contract.
func NewRevenueDistributionCaller(address common.Address, caller bind.ContractCaller) (*RevenueDistributionCaller, error) {
	contract, err := bindRevenueDistribution(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RevenueDistributionCaller{contract: contract}, nil
}

// NewRevenueDistributionTransactor creates a new write-only instance of RevenueDistribution, bound to a specific deployed contract.
func NewRevenueDistributionTransactor(address common.Address, transactor bind.ContractTransactor) (*RevenueDistributionTransactor, error) {
	contract, err := bindRevenueDistribution(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RevenueDistributionTransactor{contract: contract}, nil
}

// NewRevenueDistributionFilterer creates a new log filterer instance of RevenueDistribution, bound to a specific deployed contract.
func NewRevenueDistributionFilterer(address common.Address, filterer bind.ContractFilterer) (*RevenueDistributionFilterer, error) {
	contract, err := bindRevenueDistribution(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RevenueDistributionFilterer{contract: contract}, nil
}

// bindRevenueDistribution binds a generic wrapper to an already deployed contract.
func bindRevenueDistribution(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RevenueDistributionMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RevenueDistribution *RevenueDistributionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RevenueDistribution.Contract.RevenueDistributionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RevenueDistribution *RevenueDistributionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RevenueDistribution.Contract.RevenueDistributionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RevenueDistribution *RevenueDistributionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RevenueDistribution.Contract.RevenueDistributionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RevenueDistribution *RevenueDistributionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RevenueDistribution.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RevenueDistribution *RevenueDistributionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RevenueDistribution.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RevenueDistribution *RevenueDistributionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RevenueDistribution.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_RevenueDistribution *RevenueDistributionCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RevenueDistribution.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_RevenueDistribution *RevenueDistributionSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _RevenueDistribution.Contract.DEFAULTADMINROLE(&_RevenueDistribution.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_RevenueDistribution *RevenueDistributionCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _RevenueDistribution.Contract.DEFAULTADMINROLE(&_RevenueDistribution.CallOpts)
}

// DISTRIBUTORROLE is a free data retrieval call binding the contract method 0xf0bd87cc.
//
// Solidity: function DISTRIBUTOR_ROLE() view returns(bytes32)
func (_RevenueDistribution *RevenueDistributionCaller) DISTRIBUTORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RevenueDistribution.contract.Call(opts, &out, "DISTRIBUTOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DISTRIBUTORROLE is a free data retrieval call binding the contract method 0xf0bd87cc.
//
// Solidity: function DISTRIBUTOR_ROLE() view returns(bytes32)
func (_RevenueDistribution *RevenueDistributionSession) DISTRIBUTORROLE() ([32]byte, error) {
	return _RevenueDistribution.Contract.DISTRIBUTORROLE(&_RevenueDistribution.CallOpts)
}

// DISTRIBUTORROLE is a free data retrieval call binding the contract method 0xf0bd87cc.
//
// Solidity: function DISTRIBUTOR_ROLE() view returns(bytes32)
func (_RevenueDistribution *RevenueDistributionCallerSession) DISTRIBUTORROLE() ([32]byte, error) {
	return _RevenueDistribution.Contract.DISTRIBUTORROLE(&_RevenueDistribution.CallOpts)
}

// Claimed is a free data retrieval call binding the contract method 0x120aa877.
//
// Solidity: function claimed(uint256 , address ) view returns(bool)
func (_RevenueDistribution *RevenueDistributionCaller) Claimed(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _RevenueDistribution.contract.Call(opts, &out, "claimed", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Claimed is a free data retrieval call binding the contract method 0x120aa877.
//
// Solidity: function claimed(uint256 , address ) view returns(bool)
func (_RevenueDistribution *RevenueDistributionSession) Claimed(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _RevenueDistribution.Contract.Claimed(&_RevenueDistribution.CallOpts, arg0, arg1)
}

// Claimed is a free data retrieval call binding the contract method 0x120aa877.
//
// Solidity: function claimed(uint256 , address ) view returns(bool)
func (_RevenueDistribution *RevenueDistributionCallerSession) Claimed(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _RevenueDistribution.Contract.Claimed(&_RevenueDistribution.CallOpts, arg0, arg1)
}

// Distributions is a free data retrieval call binding the contract method 0x4487d3df.
//
// Solidity: function distributions(uint256 ) view returns(address token, uint256 snapshotId, uint256 totalAmount, address stablecoin, bool withdrawn)
func (_RevenueDistribution *RevenueDistributionCaller) Distributions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Token       common.Address
	SnapshotId  *big.Int
	TotalAmount *big.Int
	Stablecoin  common.Address
	Withdrawn   bool
}, error) {
	var out []interface{}
	err := _RevenueDistribution.contract.Call(opts, &out, "distributions", arg0)

	outstruct := new(struct {
		Token       common.Address
		SnapshotId  *big.Int
		TotalAmount *big.Int
		Stablecoin  common.Address
		Withdrawn   bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Token = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.SnapshotId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TotalAmount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Stablecoin = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Withdrawn = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// Distributions is a free data retrieval call binding the contract method 0x4487d3df.
//
// Solidity: function distributions(uint256 ) view returns(address token, uint256 snapshotId, uint256 totalAmount, address stablecoin, bool withdrawn)
func (_RevenueDistribution *RevenueDistributionSession) Distributions(arg0 *big.Int) (struct {
	Token       common.Address
	SnapshotId  *big.Int
	TotalAmount *big.Int
	Stablecoin  common.Address
	Withdrawn   bool
}, error) {
	return _RevenueDistribution.Contract.Distributions(&_RevenueDistribution.CallOpts, arg0)
}

// Distributions is a free data retrieval call binding the contract method 0x4487d3df.
//
// Solidity: function distributions(uint256 ) view returns(address token, uint256 snapshotId, uint256 totalAmount, address stablecoin, bool withdrawn)
func (_RevenueDistribution *RevenueDistributionCallerSession) Distributions(arg0 *big.Int) (struct {
	Token       common.Address
	SnapshotId  *big.Int
	TotalAmount *big.Int
	Stablecoin  common.Address
	Withdrawn   bool
}, error) {
	return _RevenueDistribution.Contract.Distributions(&_RevenueDistribution.CallOpts, arg0)
}

// DistributionsCount is a free data retrieval call binding the contract method 0xdaafd1b6.
//
// Solidity: function distributionsCount() view returns(uint256)
func (_RevenueDistribution *RevenueDistributionCaller) DistributionsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RevenueDistribution.contract.Call(opts, &out, "distributionsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DistributionsCount is a free data retrieval call binding the contract method 0xdaafd1b6.
//
// Solidity: function distributionsCount() view returns(uint256)
func (_RevenueDistribution *RevenueDistributionSession) DistributionsCount() (*big.Int, error) {
	return _RevenueDistribution.Contract.DistributionsCount(&_RevenueDistribution.CallOpts)
}

// DistributionsCount is a free data retrieval call binding the contract method 0xdaafd1b6.
//
// Solidity: function distributionsCount() view returns(uint256)
func (_RevenueDistribution *RevenueDistributionCallerSession) DistributionsCount() (*big.Int, error) {
	return _RevenueDistribution.Contract.DistributionsCount(&_RevenueDistribution.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_RevenueDistribution *RevenueDistributionCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _RevenueDistribution.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_RevenueDistribution *RevenueDistributionSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _RevenueDistribution.Contract.GetRoleAdmin(&_RevenueDistribution.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_RevenueDistribution *RevenueDistributionCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _RevenueDistribution.Contract.GetRoleAdmin(&_RevenueDistribution.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_RevenueDistribution *RevenueDistributionCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _RevenueDistribution.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_RevenueDistribution *RevenueDistributionSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _RevenueDistribution.Contract.HasRole(&_RevenueDistribution.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_RevenueDistribution *RevenueDistributionCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _RevenueDistribution.Contract.HasRole(&_RevenueDistribution.CallOpts, role, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RevenueDistribution *RevenueDistributionCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _RevenueDistribution.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RevenueDistribution *RevenueDistributionSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _RevenueDistribution.Contract.SupportsInterface(&_RevenueDistribution.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RevenueDistribution *RevenueDistributionCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _RevenueDistribution.Contract.SupportsInterface(&_RevenueDistribution.CallOpts, interfaceId)
}

// ClaimRevenue is a paid mutator transaction binding the contract method 0x1cb536ed.
//
// Solidity: function claimRevenue(uint256 distributionId) returns()
func (_RevenueDistribution *RevenueDistributionTransactor) ClaimRevenue(opts *bind.TransactOpts, distributionId *big.Int) (*types.Transaction, error) {
	return _RevenueDistribution.contract.Transact(opts, "claimRevenue", distributionId)
}

// ClaimRevenue is a paid mutator transaction binding the contract method 0x1cb536ed.
//
// Solidity: function claimRevenue(uint256 distributionId) returns()
func (_RevenueDistribution *RevenueDistributionSession) ClaimRevenue(distributionId *big.Int) (*types.Transaction, error) {
	return _RevenueDistribution.Contract.ClaimRevenue(&_RevenueDistribution.TransactOpts, distributionId)
}

// ClaimRevenue is a paid mutator transaction binding the contract method 0x1cb536ed.
//
// Solidity: function claimRevenue(uint256 distributionId) returns()
func (_RevenueDistribution *RevenueDistributionTransactorSession) ClaimRevenue(distributionId *big.Int) (*types.Transaction, error) {
	return _RevenueDistribution.Contract.ClaimRevenue(&_RevenueDistribution.TransactOpts, distributionId)
}

// DepositRevenue is a paid mutator transaction binding the contract method 0x96909992.
//
// Solidity: function depositRevenue(address tokenAddress, address stablecoin, uint256 amount) returns()
func (_RevenueDistribution *RevenueDistributionTransactor) DepositRevenue(opts *bind.TransactOpts, tokenAddress common.Address, stablecoin common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RevenueDistribution.contract.Transact(opts, "depositRevenue", tokenAddress, stablecoin, amount)
}

// DepositRevenue is a paid mutator transaction binding the contract method 0x96909992.
//
// Solidity: function depositRevenue(address tokenAddress, address stablecoin, uint256 amount) returns()
func (_RevenueDistribution *RevenueDistributionSession) DepositRevenue(tokenAddress common.Address, stablecoin common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RevenueDistribution.Contract.DepositRevenue(&_RevenueDistribution.TransactOpts, tokenAddress, stablecoin, amount)
}

// DepositRevenue is a paid mutator transaction binding the contract method 0x96909992.
//
// Solidity: function depositRevenue(address tokenAddress, address stablecoin, uint256 amount) returns()
func (_RevenueDistribution *RevenueDistributionTransactorSession) DepositRevenue(tokenAddress common.Address, stablecoin common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RevenueDistribution.Contract.DepositRevenue(&_RevenueDistribution.TransactOpts, tokenAddress, stablecoin, amount)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_RevenueDistribution *RevenueDistributionTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RevenueDistribution.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_RevenueDistribution *RevenueDistributionSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RevenueDistribution.Contract.GrantRole(&_RevenueDistribution.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_RevenueDistribution *RevenueDistributionTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RevenueDistribution.Contract.GrantRole(&_RevenueDistribution.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_RevenueDistribution *RevenueDistributionTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RevenueDistribution.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_RevenueDistribution *RevenueDistributionSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RevenueDistribution.Contract.RenounceRole(&_RevenueDistribution.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_RevenueDistribution *RevenueDistributionTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RevenueDistribution.Contract.RenounceRole(&_RevenueDistribution.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_RevenueDistribution *RevenueDistributionTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RevenueDistribution.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_RevenueDistribution *RevenueDistributionSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RevenueDistribution.Contract.RevokeRole(&_RevenueDistribution.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_RevenueDistribution *RevenueDistributionTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RevenueDistribution.Contract.RevokeRole(&_RevenueDistribution.TransactOpts, role, account)
}

// RevenueDistributionRevenueClaimedIterator is returned from FilterRevenueClaimed and is used to iterate over the raw logs and unpacked data for RevenueClaimed events raised by the RevenueDistribution contract.
type RevenueDistributionRevenueClaimedIterator struct {
	Event *RevenueDistributionRevenueClaimed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RevenueDistributionRevenueClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RevenueDistributionRevenueClaimed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RevenueDistributionRevenueClaimed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RevenueDistributionRevenueClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RevenueDistributionRevenueClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RevenueDistributionRevenueClaimed represents a RevenueClaimed event raised by the RevenueDistribution contract.
type RevenueDistributionRevenueClaimed struct {
	DistributionId *big.Int
	Claimant       common.Address
	Amount         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRevenueClaimed is a free log retrieval operation binding the contract event 0xb9e8470097faa00e83252475f2ee4b69007b0bb2405268ebec248676998a21b3.
//
// Solidity: event RevenueClaimed(uint256 indexed distributionId, address indexed claimant, uint256 amount)
func (_RevenueDistribution *RevenueDistributionFilterer) FilterRevenueClaimed(opts *bind.FilterOpts, distributionId []*big.Int, claimant []common.Address) (*RevenueDistributionRevenueClaimedIterator, error) {

	var distributionIdRule []interface{}
	for _, distributionIdItem := range distributionId {
		distributionIdRule = append(distributionIdRule, distributionIdItem)
	}
	var claimantRule []interface{}
	for _, claimantItem := range claimant {
		claimantRule = append(claimantRule, claimantItem)
	}

	logs, sub, err := _RevenueDistribution.contract.FilterLogs(opts, "RevenueClaimed", distributionIdRule, claimantRule)
	if err != nil {
		return nil, err
	}
	return &RevenueDistributionRevenueClaimedIterator{contract: _RevenueDistribution.contract, event: "RevenueClaimed", logs: logs, sub: sub}, nil
}

// WatchRevenueClaimed is a free log subscription operation binding the contract event 0xb9e8470097faa00e83252475f2ee4b69007b0bb2405268ebec248676998a21b3.
//
// Solidity: event RevenueClaimed(uint256 indexed distributionId, address indexed claimant, uint256 amount)
func (_RevenueDistribution *RevenueDistributionFilterer) WatchRevenueClaimed(opts *bind.WatchOpts, sink chan<- *RevenueDistributionRevenueClaimed, distributionId []*big.Int, claimant []common.Address) (event.Subscription, error) {

	var distributionIdRule []interface{}
	for _, distributionIdItem := range distributionId {
		distributionIdRule = append(distributionIdRule, distributionIdItem)
	}
	var claimantRule []interface{}
	for _, claimantItem := range claimant {
		claimantRule = append(claimantRule, claimantItem)
	}

	logs, sub, err := _RevenueDistribution.contract.WatchLogs(opts, "RevenueClaimed", distributionIdRule, claimantRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RevenueDistributionRevenueClaimed)
				if err := _RevenueDistribution.contract.UnpackLog(event, "RevenueClaimed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRevenueClaimed is a log parse operation binding the contract event 0xb9e8470097faa00e83252475f2ee4b69007b0bb2405268ebec248676998a21b3.
//
// Solidity: event RevenueClaimed(uint256 indexed distributionId, address indexed claimant, uint256 amount)
func (_RevenueDistribution *RevenueDistributionFilterer) ParseRevenueClaimed(log types.Log) (*RevenueDistributionRevenueClaimed, error) {
	event := new(RevenueDistributionRevenueClaimed)
	if err := _RevenueDistribution.contract.UnpackLog(event, "RevenueClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RevenueDistributionRevenueDepositedIterator is returned from FilterRevenueDeposited and is used to iterate over the raw logs and unpacked data for RevenueDeposited events raised by the RevenueDistribution contract.
type RevenueDistributionRevenueDepositedIterator struct {
	Event *RevenueDistributionRevenueDeposited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RevenueDistributionRevenueDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RevenueDistributionRevenueDeposited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RevenueDistributionRevenueDeposited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RevenueDistributionRevenueDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RevenueDistributionRevenueDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RevenueDistributionRevenueDeposited represents a RevenueDeposited event raised by the RevenueDistribution contract.
type RevenueDistributionRevenueDeposited struct {
	DistributionId *big.Int
	Token          common.Address
	Amount         *big.Int
	SnapshotId     *big.Int
	Stablecoin     common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRevenueDeposited is a free log retrieval operation binding the contract event 0x53dd77f5e24dd324ca2bd42ac3b4f8531af533c339e22123fcec4a4bea2e6b86.
//
// Solidity: event RevenueDeposited(uint256 indexed distributionId, address indexed token, uint256 amount, uint256 snapshotId, address stablecoin)
func (_RevenueDistribution *RevenueDistributionFilterer) FilterRevenueDeposited(opts *bind.FilterOpts, distributionId []*big.Int, token []common.Address) (*RevenueDistributionRevenueDepositedIterator, error) {

	var distributionIdRule []interface{}
	for _, distributionIdItem := range distributionId {
		distributionIdRule = append(distributionIdRule, distributionIdItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _RevenueDistribution.contract.FilterLogs(opts, "RevenueDeposited", distributionIdRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &RevenueDistributionRevenueDepositedIterator{contract: _RevenueDistribution.contract, event: "RevenueDeposited", logs: logs, sub: sub}, nil
}

// WatchRevenueDeposited is a free log subscription operation binding the contract event 0x53dd77f5e24dd324ca2bd42ac3b4f8531af533c339e22123fcec4a4bea2e6b86.
//
// Solidity: event RevenueDeposited(uint256 indexed distributionId, address indexed token, uint256 amount, uint256 snapshotId, address stablecoin)
func (_RevenueDistribution *RevenueDistributionFilterer) WatchRevenueDeposited(opts *bind.WatchOpts, sink chan<- *RevenueDistributionRevenueDeposited, distributionId []*big.Int, token []common.Address) (event.Subscription, error) {

	var distributionIdRule []interface{}
	for _, distributionIdItem := range distributionId {
		distributionIdRule = append(distributionIdRule, distributionIdItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _RevenueDistribution.contract.WatchLogs(opts, "RevenueDeposited", distributionIdRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RevenueDistributionRevenueDeposited)
				if err := _RevenueDistribution.contract.UnpackLog(event, "RevenueDeposited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRevenueDeposited is a log parse operation binding the contract event 0x53dd77f5e24dd324ca2bd42ac3b4f8531af533c339e22123fcec4a4bea2e6b86.
//
// Solidity: event RevenueDeposited(uint256 indexed distributionId, address indexed token, uint256 amount, uint256 snapshotId, address stablecoin)
func (_RevenueDistribution *RevenueDistributionFilterer) ParseRevenueDeposited(log types.Log) (*RevenueDistributionRevenueDeposited, error) {
	event := new(RevenueDistributionRevenueDeposited)
	if err := _RevenueDistribution.contract.UnpackLog(event, "RevenueDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RevenueDistributionRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the RevenueDistribution contract.
type RevenueDistributionRoleAdminChangedIterator struct {
	Event *RevenueDistributionRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RevenueDistributionRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RevenueDistributionRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RevenueDistributionRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RevenueDistributionRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RevenueDistributionRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RevenueDistributionRoleAdminChanged represents a RoleAdminChanged event raised by the RevenueDistribution contract.
type RevenueDistributionRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_RevenueDistribution *RevenueDistributionFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*RevenueDistributionRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _RevenueDistribution.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &RevenueDistributionRoleAdminChangedIterator{contract: _RevenueDistribution.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_RevenueDistribution *RevenueDistributionFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *RevenueDistributionRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _RevenueDistribution.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RevenueDistributionRoleAdminChanged)
				if err := _RevenueDistribution.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_RevenueDistribution *RevenueDistributionFilterer) ParseRoleAdminChanged(log types.Log) (*RevenueDistributionRoleAdminChanged, error) {
	event := new(RevenueDistributionRoleAdminChanged)
	if err := _RevenueDistribution.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RevenueDistributionRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the RevenueDistribution contract.
type RevenueDistributionRoleGrantedIterator struct {
	Event *RevenueDistributionRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RevenueDistributionRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RevenueDistributionRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RevenueDistributionRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RevenueDistributionRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RevenueDistributionRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RevenueDistributionRoleGranted represents a RoleGranted event raised by the RevenueDistribution contract.
type RevenueDistributionRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_RevenueDistribution *RevenueDistributionFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*RevenueDistributionRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _RevenueDistribution.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &RevenueDistributionRoleGrantedIterator{contract: _RevenueDistribution.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_RevenueDistribution *RevenueDistributionFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *RevenueDistributionRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _RevenueDistribution.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RevenueDistributionRoleGranted)
				if err := _RevenueDistribution.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_RevenueDistribution *RevenueDistributionFilterer) ParseRoleGranted(log types.Log) (*RevenueDistributionRoleGranted, error) {
	event := new(RevenueDistributionRoleGranted)
	if err := _RevenueDistribution.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RevenueDistributionRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the RevenueDistribution contract.
type RevenueDistributionRoleRevokedIterator struct {
	Event *RevenueDistributionRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RevenueDistributionRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RevenueDistributionRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RevenueDistributionRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RevenueDistributionRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RevenueDistributionRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RevenueDistributionRoleRevoked represents a RoleRevoked event raised by the RevenueDistribution contract.
type RevenueDistributionRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_RevenueDistribution *RevenueDistributionFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*RevenueDistributionRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _RevenueDistribution.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &RevenueDistributionRoleRevokedIterator{contract: _RevenueDistribution.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_RevenueDistribution *RevenueDistributionFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *RevenueDistributionRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _RevenueDistribution.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RevenueDistributionRoleRevoked)
				if err := _RevenueDistribution.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_RevenueDistribution *RevenueDistributionFilterer) ParseRoleRevoked(log types.Log) (*RevenueDistributionRoleRevoked, error) {
	event := new(RevenueDistributionRoleRevoked)
	if err := _RevenueDistribution.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

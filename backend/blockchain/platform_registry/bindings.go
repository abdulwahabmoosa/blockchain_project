// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package platform_registry

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

// PlatformRegistryMetaData contains all meta data concerning the PlatformRegistry contract.
var PlatformRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approval\",\"type\":\"address\"}],\"name\":\"ApprovalServiceSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"name\":\"PropertyFactorySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"revenue\",\"type\":\"address\"}],\"name\":\"RevenueDistributionSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"approvalService\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getApprovalService\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPropertyFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRevenueDistribution\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"propertyFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revenueDistribution\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_approval\",\"type\":\"address\"}],\"name\":\"setApprovalService\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"}],\"name\":\"setPropertyFactory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_revenue\",\"type\":\"address\"}],\"name\":\"setRevenueDistribution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PlatformRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use PlatformRegistryMetaData.ABI instead.
var PlatformRegistryABI = PlatformRegistryMetaData.ABI

// PlatformRegistry is an auto generated Go binding around an Ethereum contract.
type PlatformRegistry struct {
	PlatformRegistryCaller     // Read-only binding to the contract
	PlatformRegistryTransactor // Write-only binding to the contract
	PlatformRegistryFilterer   // Log filterer for contract events
}

// PlatformRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type PlatformRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlatformRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PlatformRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlatformRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PlatformRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlatformRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PlatformRegistrySession struct {
	Contract     *PlatformRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PlatformRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PlatformRegistryCallerSession struct {
	Contract *PlatformRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// PlatformRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PlatformRegistryTransactorSession struct {
	Contract     *PlatformRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// PlatformRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type PlatformRegistryRaw struct {
	Contract *PlatformRegistry // Generic contract binding to access the raw methods on
}

// PlatformRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PlatformRegistryCallerRaw struct {
	Contract *PlatformRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// PlatformRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PlatformRegistryTransactorRaw struct {
	Contract *PlatformRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPlatformRegistry creates a new instance of PlatformRegistry, bound to a specific deployed contract.
func NewPlatformRegistry(address common.Address, backend bind.ContractBackend) (*PlatformRegistry, error) {
	contract, err := bindPlatformRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PlatformRegistry{PlatformRegistryCaller: PlatformRegistryCaller{contract: contract}, PlatformRegistryTransactor: PlatformRegistryTransactor{contract: contract}, PlatformRegistryFilterer: PlatformRegistryFilterer{contract: contract}}, nil
}

// NewPlatformRegistryCaller creates a new read-only instance of PlatformRegistry, bound to a specific deployed contract.
func NewPlatformRegistryCaller(address common.Address, caller bind.ContractCaller) (*PlatformRegistryCaller, error) {
	contract, err := bindPlatformRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PlatformRegistryCaller{contract: contract}, nil
}

// NewPlatformRegistryTransactor creates a new write-only instance of PlatformRegistry, bound to a specific deployed contract.
func NewPlatformRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*PlatformRegistryTransactor, error) {
	contract, err := bindPlatformRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PlatformRegistryTransactor{contract: contract}, nil
}

// NewPlatformRegistryFilterer creates a new log filterer instance of PlatformRegistry, bound to a specific deployed contract.
func NewPlatformRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*PlatformRegistryFilterer, error) {
	contract, err := bindPlatformRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PlatformRegistryFilterer{contract: contract}, nil
}

// bindPlatformRegistry binds a generic wrapper to an already deployed contract.
func bindPlatformRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PlatformRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PlatformRegistry *PlatformRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PlatformRegistry.Contract.PlatformRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PlatformRegistry *PlatformRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlatformRegistry.Contract.PlatformRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PlatformRegistry *PlatformRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PlatformRegistry.Contract.PlatformRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PlatformRegistry *PlatformRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PlatformRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PlatformRegistry *PlatformRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlatformRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PlatformRegistry *PlatformRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PlatformRegistry.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_PlatformRegistry *PlatformRegistryCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PlatformRegistry.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_PlatformRegistry *PlatformRegistrySession) ADMINROLE() ([32]byte, error) {
	return _PlatformRegistry.Contract.ADMINROLE(&_PlatformRegistry.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_PlatformRegistry *PlatformRegistryCallerSession) ADMINROLE() ([32]byte, error) {
	return _PlatformRegistry.Contract.ADMINROLE(&_PlatformRegistry.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PlatformRegistry *PlatformRegistryCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PlatformRegistry.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PlatformRegistry *PlatformRegistrySession) DEFAULTADMINROLE() ([32]byte, error) {
	return _PlatformRegistry.Contract.DEFAULTADMINROLE(&_PlatformRegistry.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PlatformRegistry *PlatformRegistryCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _PlatformRegistry.Contract.DEFAULTADMINROLE(&_PlatformRegistry.CallOpts)
}

// ApprovalService is a free data retrieval call binding the contract method 0xf23d0c2f.
//
// Solidity: function approvalService() view returns(address)
func (_PlatformRegistry *PlatformRegistryCaller) ApprovalService(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PlatformRegistry.contract.Call(opts, &out, "approvalService")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ApprovalService is a free data retrieval call binding the contract method 0xf23d0c2f.
//
// Solidity: function approvalService() view returns(address)
func (_PlatformRegistry *PlatformRegistrySession) ApprovalService() (common.Address, error) {
	return _PlatformRegistry.Contract.ApprovalService(&_PlatformRegistry.CallOpts)
}

// ApprovalService is a free data retrieval call binding the contract method 0xf23d0c2f.
//
// Solidity: function approvalService() view returns(address)
func (_PlatformRegistry *PlatformRegistryCallerSession) ApprovalService() (common.Address, error) {
	return _PlatformRegistry.Contract.ApprovalService(&_PlatformRegistry.CallOpts)
}

// GetApprovalService is a free data retrieval call binding the contract method 0x3a90dbca.
//
// Solidity: function getApprovalService() view returns(address)
func (_PlatformRegistry *PlatformRegistryCaller) GetApprovalService(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PlatformRegistry.contract.Call(opts, &out, "getApprovalService")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApprovalService is a free data retrieval call binding the contract method 0x3a90dbca.
//
// Solidity: function getApprovalService() view returns(address)
func (_PlatformRegistry *PlatformRegistrySession) GetApprovalService() (common.Address, error) {
	return _PlatformRegistry.Contract.GetApprovalService(&_PlatformRegistry.CallOpts)
}

// GetApprovalService is a free data retrieval call binding the contract method 0x3a90dbca.
//
// Solidity: function getApprovalService() view returns(address)
func (_PlatformRegistry *PlatformRegistryCallerSession) GetApprovalService() (common.Address, error) {
	return _PlatformRegistry.Contract.GetApprovalService(&_PlatformRegistry.CallOpts)
}

// GetPropertyFactory is a free data retrieval call binding the contract method 0xe95038ea.
//
// Solidity: function getPropertyFactory() view returns(address)
func (_PlatformRegistry *PlatformRegistryCaller) GetPropertyFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PlatformRegistry.contract.Call(opts, &out, "getPropertyFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPropertyFactory is a free data retrieval call binding the contract method 0xe95038ea.
//
// Solidity: function getPropertyFactory() view returns(address)
func (_PlatformRegistry *PlatformRegistrySession) GetPropertyFactory() (common.Address, error) {
	return _PlatformRegistry.Contract.GetPropertyFactory(&_PlatformRegistry.CallOpts)
}

// GetPropertyFactory is a free data retrieval call binding the contract method 0xe95038ea.
//
// Solidity: function getPropertyFactory() view returns(address)
func (_PlatformRegistry *PlatformRegistryCallerSession) GetPropertyFactory() (common.Address, error) {
	return _PlatformRegistry.Contract.GetPropertyFactory(&_PlatformRegistry.CallOpts)
}

// GetRevenueDistribution is a free data retrieval call binding the contract method 0xef8f3681.
//
// Solidity: function getRevenueDistribution() view returns(address)
func (_PlatformRegistry *PlatformRegistryCaller) GetRevenueDistribution(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PlatformRegistry.contract.Call(opts, &out, "getRevenueDistribution")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRevenueDistribution is a free data retrieval call binding the contract method 0xef8f3681.
//
// Solidity: function getRevenueDistribution() view returns(address)
func (_PlatformRegistry *PlatformRegistrySession) GetRevenueDistribution() (common.Address, error) {
	return _PlatformRegistry.Contract.GetRevenueDistribution(&_PlatformRegistry.CallOpts)
}

// GetRevenueDistribution is a free data retrieval call binding the contract method 0xef8f3681.
//
// Solidity: function getRevenueDistribution() view returns(address)
func (_PlatformRegistry *PlatformRegistryCallerSession) GetRevenueDistribution() (common.Address, error) {
	return _PlatformRegistry.Contract.GetRevenueDistribution(&_PlatformRegistry.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PlatformRegistry *PlatformRegistryCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _PlatformRegistry.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PlatformRegistry *PlatformRegistrySession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _PlatformRegistry.Contract.GetRoleAdmin(&_PlatformRegistry.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PlatformRegistry *PlatformRegistryCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _PlatformRegistry.Contract.GetRoleAdmin(&_PlatformRegistry.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PlatformRegistry *PlatformRegistryCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _PlatformRegistry.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PlatformRegistry *PlatformRegistrySession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _PlatformRegistry.Contract.HasRole(&_PlatformRegistry.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PlatformRegistry *PlatformRegistryCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _PlatformRegistry.Contract.HasRole(&_PlatformRegistry.CallOpts, role, account)
}

// PropertyFactory is a free data retrieval call binding the contract method 0xa3db54b0.
//
// Solidity: function propertyFactory() view returns(address)
func (_PlatformRegistry *PlatformRegistryCaller) PropertyFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PlatformRegistry.contract.Call(opts, &out, "propertyFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PropertyFactory is a free data retrieval call binding the contract method 0xa3db54b0.
//
// Solidity: function propertyFactory() view returns(address)
func (_PlatformRegistry *PlatformRegistrySession) PropertyFactory() (common.Address, error) {
	return _PlatformRegistry.Contract.PropertyFactory(&_PlatformRegistry.CallOpts)
}

// PropertyFactory is a free data retrieval call binding the contract method 0xa3db54b0.
//
// Solidity: function propertyFactory() view returns(address)
func (_PlatformRegistry *PlatformRegistryCallerSession) PropertyFactory() (common.Address, error) {
	return _PlatformRegistry.Contract.PropertyFactory(&_PlatformRegistry.CallOpts)
}

// RevenueDistribution is a free data retrieval call binding the contract method 0x07db702a.
//
// Solidity: function revenueDistribution() view returns(address)
func (_PlatformRegistry *PlatformRegistryCaller) RevenueDistribution(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PlatformRegistry.contract.Call(opts, &out, "revenueDistribution")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RevenueDistribution is a free data retrieval call binding the contract method 0x07db702a.
//
// Solidity: function revenueDistribution() view returns(address)
func (_PlatformRegistry *PlatformRegistrySession) RevenueDistribution() (common.Address, error) {
	return _PlatformRegistry.Contract.RevenueDistribution(&_PlatformRegistry.CallOpts)
}

// RevenueDistribution is a free data retrieval call binding the contract method 0x07db702a.
//
// Solidity: function revenueDistribution() view returns(address)
func (_PlatformRegistry *PlatformRegistryCallerSession) RevenueDistribution() (common.Address, error) {
	return _PlatformRegistry.Contract.RevenueDistribution(&_PlatformRegistry.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PlatformRegistry *PlatformRegistryCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _PlatformRegistry.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PlatformRegistry *PlatformRegistrySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _PlatformRegistry.Contract.SupportsInterface(&_PlatformRegistry.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PlatformRegistry *PlatformRegistryCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _PlatformRegistry.Contract.SupportsInterface(&_PlatformRegistry.CallOpts, interfaceId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PlatformRegistry *PlatformRegistryTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PlatformRegistry.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PlatformRegistry *PlatformRegistrySession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PlatformRegistry.Contract.GrantRole(&_PlatformRegistry.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PlatformRegistry *PlatformRegistryTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PlatformRegistry.Contract.GrantRole(&_PlatformRegistry.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_PlatformRegistry *PlatformRegistryTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PlatformRegistry.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_PlatformRegistry *PlatformRegistrySession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PlatformRegistry.Contract.RenounceRole(&_PlatformRegistry.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_PlatformRegistry *PlatformRegistryTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PlatformRegistry.Contract.RenounceRole(&_PlatformRegistry.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PlatformRegistry *PlatformRegistryTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PlatformRegistry.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PlatformRegistry *PlatformRegistrySession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PlatformRegistry.Contract.RevokeRole(&_PlatformRegistry.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PlatformRegistry *PlatformRegistryTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PlatformRegistry.Contract.RevokeRole(&_PlatformRegistry.TransactOpts, role, account)
}

// SetApprovalService is a paid mutator transaction binding the contract method 0xc9011a41.
//
// Solidity: function setApprovalService(address _approval) returns()
func (_PlatformRegistry *PlatformRegistryTransactor) SetApprovalService(opts *bind.TransactOpts, _approval common.Address) (*types.Transaction, error) {
	return _PlatformRegistry.contract.Transact(opts, "setApprovalService", _approval)
}

// SetApprovalService is a paid mutator transaction binding the contract method 0xc9011a41.
//
// Solidity: function setApprovalService(address _approval) returns()
func (_PlatformRegistry *PlatformRegistrySession) SetApprovalService(_approval common.Address) (*types.Transaction, error) {
	return _PlatformRegistry.Contract.SetApprovalService(&_PlatformRegistry.TransactOpts, _approval)
}

// SetApprovalService is a paid mutator transaction binding the contract method 0xc9011a41.
//
// Solidity: function setApprovalService(address _approval) returns()
func (_PlatformRegistry *PlatformRegistryTransactorSession) SetApprovalService(_approval common.Address) (*types.Transaction, error) {
	return _PlatformRegistry.Contract.SetApprovalService(&_PlatformRegistry.TransactOpts, _approval)
}

// SetPropertyFactory is a paid mutator transaction binding the contract method 0xe897e085.
//
// Solidity: function setPropertyFactory(address _factory) returns()
func (_PlatformRegistry *PlatformRegistryTransactor) SetPropertyFactory(opts *bind.TransactOpts, _factory common.Address) (*types.Transaction, error) {
	return _PlatformRegistry.contract.Transact(opts, "setPropertyFactory", _factory)
}

// SetPropertyFactory is a paid mutator transaction binding the contract method 0xe897e085.
//
// Solidity: function setPropertyFactory(address _factory) returns()
func (_PlatformRegistry *PlatformRegistrySession) SetPropertyFactory(_factory common.Address) (*types.Transaction, error) {
	return _PlatformRegistry.Contract.SetPropertyFactory(&_PlatformRegistry.TransactOpts, _factory)
}

// SetPropertyFactory is a paid mutator transaction binding the contract method 0xe897e085.
//
// Solidity: function setPropertyFactory(address _factory) returns()
func (_PlatformRegistry *PlatformRegistryTransactorSession) SetPropertyFactory(_factory common.Address) (*types.Transaction, error) {
	return _PlatformRegistry.Contract.SetPropertyFactory(&_PlatformRegistry.TransactOpts, _factory)
}

// SetRevenueDistribution is a paid mutator transaction binding the contract method 0x64bcbe47.
//
// Solidity: function setRevenueDistribution(address _revenue) returns()
func (_PlatformRegistry *PlatformRegistryTransactor) SetRevenueDistribution(opts *bind.TransactOpts, _revenue common.Address) (*types.Transaction, error) {
	return _PlatformRegistry.contract.Transact(opts, "setRevenueDistribution", _revenue)
}

// SetRevenueDistribution is a paid mutator transaction binding the contract method 0x64bcbe47.
//
// Solidity: function setRevenueDistribution(address _revenue) returns()
func (_PlatformRegistry *PlatformRegistrySession) SetRevenueDistribution(_revenue common.Address) (*types.Transaction, error) {
	return _PlatformRegistry.Contract.SetRevenueDistribution(&_PlatformRegistry.TransactOpts, _revenue)
}

// SetRevenueDistribution is a paid mutator transaction binding the contract method 0x64bcbe47.
//
// Solidity: function setRevenueDistribution(address _revenue) returns()
func (_PlatformRegistry *PlatformRegistryTransactorSession) SetRevenueDistribution(_revenue common.Address) (*types.Transaction, error) {
	return _PlatformRegistry.Contract.SetRevenueDistribution(&_PlatformRegistry.TransactOpts, _revenue)
}

// PlatformRegistryApprovalServiceSetIterator is returned from FilterApprovalServiceSet and is used to iterate over the raw logs and unpacked data for ApprovalServiceSet events raised by the PlatformRegistry contract.
type PlatformRegistryApprovalServiceSetIterator struct {
	Event *PlatformRegistryApprovalServiceSet // Event containing the contract specifics and raw log

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
func (it *PlatformRegistryApprovalServiceSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlatformRegistryApprovalServiceSet)
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
		it.Event = new(PlatformRegistryApprovalServiceSet)
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
func (it *PlatformRegistryApprovalServiceSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlatformRegistryApprovalServiceSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlatformRegistryApprovalServiceSet represents a ApprovalServiceSet event raised by the PlatformRegistry contract.
type PlatformRegistryApprovalServiceSet struct {
	Approval common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalServiceSet is a free log retrieval operation binding the contract event 0xfb48851223fdfe38226c4119336d9a1100fcc0e8c84ce888a08843adb7cb1a2e.
//
// Solidity: event ApprovalServiceSet(address indexed approval)
func (_PlatformRegistry *PlatformRegistryFilterer) FilterApprovalServiceSet(opts *bind.FilterOpts, approval []common.Address) (*PlatformRegistryApprovalServiceSetIterator, error) {

	var approvalRule []interface{}
	for _, approvalItem := range approval {
		approvalRule = append(approvalRule, approvalItem)
	}

	logs, sub, err := _PlatformRegistry.contract.FilterLogs(opts, "ApprovalServiceSet", approvalRule)
	if err != nil {
		return nil, err
	}
	return &PlatformRegistryApprovalServiceSetIterator{contract: _PlatformRegistry.contract, event: "ApprovalServiceSet", logs: logs, sub: sub}, nil
}

// WatchApprovalServiceSet is a free log subscription operation binding the contract event 0xfb48851223fdfe38226c4119336d9a1100fcc0e8c84ce888a08843adb7cb1a2e.
//
// Solidity: event ApprovalServiceSet(address indexed approval)
func (_PlatformRegistry *PlatformRegistryFilterer) WatchApprovalServiceSet(opts *bind.WatchOpts, sink chan<- *PlatformRegistryApprovalServiceSet, approval []common.Address) (event.Subscription, error) {

	var approvalRule []interface{}
	for _, approvalItem := range approval {
		approvalRule = append(approvalRule, approvalItem)
	}

	logs, sub, err := _PlatformRegistry.contract.WatchLogs(opts, "ApprovalServiceSet", approvalRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlatformRegistryApprovalServiceSet)
				if err := _PlatformRegistry.contract.UnpackLog(event, "ApprovalServiceSet", log); err != nil {
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

// ParseApprovalServiceSet is a log parse operation binding the contract event 0xfb48851223fdfe38226c4119336d9a1100fcc0e8c84ce888a08843adb7cb1a2e.
//
// Solidity: event ApprovalServiceSet(address indexed approval)
func (_PlatformRegistry *PlatformRegistryFilterer) ParseApprovalServiceSet(log types.Log) (*PlatformRegistryApprovalServiceSet, error) {
	event := new(PlatformRegistryApprovalServiceSet)
	if err := _PlatformRegistry.contract.UnpackLog(event, "ApprovalServiceSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlatformRegistryPropertyFactorySetIterator is returned from FilterPropertyFactorySet and is used to iterate over the raw logs and unpacked data for PropertyFactorySet events raised by the PlatformRegistry contract.
type PlatformRegistryPropertyFactorySetIterator struct {
	Event *PlatformRegistryPropertyFactorySet // Event containing the contract specifics and raw log

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
func (it *PlatformRegistryPropertyFactorySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlatformRegistryPropertyFactorySet)
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
		it.Event = new(PlatformRegistryPropertyFactorySet)
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
func (it *PlatformRegistryPropertyFactorySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlatformRegistryPropertyFactorySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlatformRegistryPropertyFactorySet represents a PropertyFactorySet event raised by the PlatformRegistry contract.
type PlatformRegistryPropertyFactorySet struct {
	Factory common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPropertyFactorySet is a free log retrieval operation binding the contract event 0x556e0c3ac1d78ff14830546a21584bcef8595b6a4d6aad59e8b19b1cdcd0bb48.
//
// Solidity: event PropertyFactorySet(address indexed factory)
func (_PlatformRegistry *PlatformRegistryFilterer) FilterPropertyFactorySet(opts *bind.FilterOpts, factory []common.Address) (*PlatformRegistryPropertyFactorySetIterator, error) {

	var factoryRule []interface{}
	for _, factoryItem := range factory {
		factoryRule = append(factoryRule, factoryItem)
	}

	logs, sub, err := _PlatformRegistry.contract.FilterLogs(opts, "PropertyFactorySet", factoryRule)
	if err != nil {
		return nil, err
	}
	return &PlatformRegistryPropertyFactorySetIterator{contract: _PlatformRegistry.contract, event: "PropertyFactorySet", logs: logs, sub: sub}, nil
}

// WatchPropertyFactorySet is a free log subscription operation binding the contract event 0x556e0c3ac1d78ff14830546a21584bcef8595b6a4d6aad59e8b19b1cdcd0bb48.
//
// Solidity: event PropertyFactorySet(address indexed factory)
func (_PlatformRegistry *PlatformRegistryFilterer) WatchPropertyFactorySet(opts *bind.WatchOpts, sink chan<- *PlatformRegistryPropertyFactorySet, factory []common.Address) (event.Subscription, error) {

	var factoryRule []interface{}
	for _, factoryItem := range factory {
		factoryRule = append(factoryRule, factoryItem)
	}

	logs, sub, err := _PlatformRegistry.contract.WatchLogs(opts, "PropertyFactorySet", factoryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlatformRegistryPropertyFactorySet)
				if err := _PlatformRegistry.contract.UnpackLog(event, "PropertyFactorySet", log); err != nil {
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

// ParsePropertyFactorySet is a log parse operation binding the contract event 0x556e0c3ac1d78ff14830546a21584bcef8595b6a4d6aad59e8b19b1cdcd0bb48.
//
// Solidity: event PropertyFactorySet(address indexed factory)
func (_PlatformRegistry *PlatformRegistryFilterer) ParsePropertyFactorySet(log types.Log) (*PlatformRegistryPropertyFactorySet, error) {
	event := new(PlatformRegistryPropertyFactorySet)
	if err := _PlatformRegistry.contract.UnpackLog(event, "PropertyFactorySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlatformRegistryRevenueDistributionSetIterator is returned from FilterRevenueDistributionSet and is used to iterate over the raw logs and unpacked data for RevenueDistributionSet events raised by the PlatformRegistry contract.
type PlatformRegistryRevenueDistributionSetIterator struct {
	Event *PlatformRegistryRevenueDistributionSet // Event containing the contract specifics and raw log

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
func (it *PlatformRegistryRevenueDistributionSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlatformRegistryRevenueDistributionSet)
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
		it.Event = new(PlatformRegistryRevenueDistributionSet)
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
func (it *PlatformRegistryRevenueDistributionSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlatformRegistryRevenueDistributionSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlatformRegistryRevenueDistributionSet represents a RevenueDistributionSet event raised by the PlatformRegistry contract.
type PlatformRegistryRevenueDistributionSet struct {
	Revenue common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRevenueDistributionSet is a free log retrieval operation binding the contract event 0x8aaf7ef6a08e69b0c69f0719a51003d08c5357f3bbe169e26fac36ee32278756.
//
// Solidity: event RevenueDistributionSet(address indexed revenue)
func (_PlatformRegistry *PlatformRegistryFilterer) FilterRevenueDistributionSet(opts *bind.FilterOpts, revenue []common.Address) (*PlatformRegistryRevenueDistributionSetIterator, error) {

	var revenueRule []interface{}
	for _, revenueItem := range revenue {
		revenueRule = append(revenueRule, revenueItem)
	}

	logs, sub, err := _PlatformRegistry.contract.FilterLogs(opts, "RevenueDistributionSet", revenueRule)
	if err != nil {
		return nil, err
	}
	return &PlatformRegistryRevenueDistributionSetIterator{contract: _PlatformRegistry.contract, event: "RevenueDistributionSet", logs: logs, sub: sub}, nil
}

// WatchRevenueDistributionSet is a free log subscription operation binding the contract event 0x8aaf7ef6a08e69b0c69f0719a51003d08c5357f3bbe169e26fac36ee32278756.
//
// Solidity: event RevenueDistributionSet(address indexed revenue)
func (_PlatformRegistry *PlatformRegistryFilterer) WatchRevenueDistributionSet(opts *bind.WatchOpts, sink chan<- *PlatformRegistryRevenueDistributionSet, revenue []common.Address) (event.Subscription, error) {

	var revenueRule []interface{}
	for _, revenueItem := range revenue {
		revenueRule = append(revenueRule, revenueItem)
	}

	logs, sub, err := _PlatformRegistry.contract.WatchLogs(opts, "RevenueDistributionSet", revenueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlatformRegistryRevenueDistributionSet)
				if err := _PlatformRegistry.contract.UnpackLog(event, "RevenueDistributionSet", log); err != nil {
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

// ParseRevenueDistributionSet is a log parse operation binding the contract event 0x8aaf7ef6a08e69b0c69f0719a51003d08c5357f3bbe169e26fac36ee32278756.
//
// Solidity: event RevenueDistributionSet(address indexed revenue)
func (_PlatformRegistry *PlatformRegistryFilterer) ParseRevenueDistributionSet(log types.Log) (*PlatformRegistryRevenueDistributionSet, error) {
	event := new(PlatformRegistryRevenueDistributionSet)
	if err := _PlatformRegistry.contract.UnpackLog(event, "RevenueDistributionSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlatformRegistryRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the PlatformRegistry contract.
type PlatformRegistryRoleAdminChangedIterator struct {
	Event *PlatformRegistryRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *PlatformRegistryRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlatformRegistryRoleAdminChanged)
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
		it.Event = new(PlatformRegistryRoleAdminChanged)
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
func (it *PlatformRegistryRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlatformRegistryRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlatformRegistryRoleAdminChanged represents a RoleAdminChanged event raised by the PlatformRegistry contract.
type PlatformRegistryRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_PlatformRegistry *PlatformRegistryFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*PlatformRegistryRoleAdminChangedIterator, error) {

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

	logs, sub, err := _PlatformRegistry.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &PlatformRegistryRoleAdminChangedIterator{contract: _PlatformRegistry.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_PlatformRegistry *PlatformRegistryFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *PlatformRegistryRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _PlatformRegistry.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlatformRegistryRoleAdminChanged)
				if err := _PlatformRegistry.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_PlatformRegistry *PlatformRegistryFilterer) ParseRoleAdminChanged(log types.Log) (*PlatformRegistryRoleAdminChanged, error) {
	event := new(PlatformRegistryRoleAdminChanged)
	if err := _PlatformRegistry.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlatformRegistryRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the PlatformRegistry contract.
type PlatformRegistryRoleGrantedIterator struct {
	Event *PlatformRegistryRoleGranted // Event containing the contract specifics and raw log

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
func (it *PlatformRegistryRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlatformRegistryRoleGranted)
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
		it.Event = new(PlatformRegistryRoleGranted)
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
func (it *PlatformRegistryRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlatformRegistryRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlatformRegistryRoleGranted represents a RoleGranted event raised by the PlatformRegistry contract.
type PlatformRegistryRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_PlatformRegistry *PlatformRegistryFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PlatformRegistryRoleGrantedIterator, error) {

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

	logs, sub, err := _PlatformRegistry.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PlatformRegistryRoleGrantedIterator{contract: _PlatformRegistry.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_PlatformRegistry *PlatformRegistryFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *PlatformRegistryRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PlatformRegistry.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlatformRegistryRoleGranted)
				if err := _PlatformRegistry.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_PlatformRegistry *PlatformRegistryFilterer) ParseRoleGranted(log types.Log) (*PlatformRegistryRoleGranted, error) {
	event := new(PlatformRegistryRoleGranted)
	if err := _PlatformRegistry.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlatformRegistryRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the PlatformRegistry contract.
type PlatformRegistryRoleRevokedIterator struct {
	Event *PlatformRegistryRoleRevoked // Event containing the contract specifics and raw log

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
func (it *PlatformRegistryRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlatformRegistryRoleRevoked)
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
		it.Event = new(PlatformRegistryRoleRevoked)
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
func (it *PlatformRegistryRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlatformRegistryRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlatformRegistryRoleRevoked represents a RoleRevoked event raised by the PlatformRegistry contract.
type PlatformRegistryRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_PlatformRegistry *PlatformRegistryFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PlatformRegistryRoleRevokedIterator, error) {

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

	logs, sub, err := _PlatformRegistry.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PlatformRegistryRoleRevokedIterator{contract: _PlatformRegistry.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_PlatformRegistry *PlatformRegistryFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *PlatformRegistryRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PlatformRegistry.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlatformRegistryRoleRevoked)
				if err := _PlatformRegistry.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_PlatformRegistry *PlatformRegistryFilterer) ParseRoleRevoked(log types.Log) (*PlatformRegistryRoleRevoked, error) {
	event := new(PlatformRegistryRoleRevoked)
	if err := _PlatformRegistry.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

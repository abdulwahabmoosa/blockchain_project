// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package property_asset

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

// PropertyAssetMetaData contains all meta data concerning the PropertyAsset contract.
var PropertyAssetMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_assetId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_dataHash\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_valuation\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumPropertyAsset.Status\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"PropertyStatusChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PropertyTokenLinked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MANAGER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"assetId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStatus\",\"outputs\":[{\"internalType\":\"enumPropertyAsset.Status\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"propertyDataHash\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"propertyToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"setPropertyToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumPropertyAsset.Status\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"setStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"status\",\"outputs\":[{\"internalType\":\"enumPropertyAsset.Status\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"valuation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PropertyAssetABI is the input ABI used to generate the binding from.
// Deprecated: Use PropertyAssetMetaData.ABI instead.
var PropertyAssetABI = PropertyAssetMetaData.ABI

// PropertyAsset is an auto generated Go binding around an Ethereum contract.
type PropertyAsset struct {
	PropertyAssetCaller     // Read-only binding to the contract
	PropertyAssetTransactor // Write-only binding to the contract
	PropertyAssetFilterer   // Log filterer for contract events
}

// PropertyAssetCaller is an auto generated read-only Go binding around an Ethereum contract.
type PropertyAssetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PropertyAssetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PropertyAssetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PropertyAssetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PropertyAssetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PropertyAssetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PropertyAssetSession struct {
	Contract     *PropertyAsset    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PropertyAssetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PropertyAssetCallerSession struct {
	Contract *PropertyAssetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// PropertyAssetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PropertyAssetTransactorSession struct {
	Contract     *PropertyAssetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// PropertyAssetRaw is an auto generated low-level Go binding around an Ethereum contract.
type PropertyAssetRaw struct {
	Contract *PropertyAsset // Generic contract binding to access the raw methods on
}

// PropertyAssetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PropertyAssetCallerRaw struct {
	Contract *PropertyAssetCaller // Generic read-only contract binding to access the raw methods on
}

// PropertyAssetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PropertyAssetTransactorRaw struct {
	Contract *PropertyAssetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPropertyAsset creates a new instance of PropertyAsset, bound to a specific deployed contract.
func NewPropertyAsset(address common.Address, backend bind.ContractBackend) (*PropertyAsset, error) {
	contract, err := bindPropertyAsset(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PropertyAsset{PropertyAssetCaller: PropertyAssetCaller{contract: contract}, PropertyAssetTransactor: PropertyAssetTransactor{contract: contract}, PropertyAssetFilterer: PropertyAssetFilterer{contract: contract}}, nil
}

// NewPropertyAssetCaller creates a new read-only instance of PropertyAsset, bound to a specific deployed contract.
func NewPropertyAssetCaller(address common.Address, caller bind.ContractCaller) (*PropertyAssetCaller, error) {
	contract, err := bindPropertyAsset(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PropertyAssetCaller{contract: contract}, nil
}

// NewPropertyAssetTransactor creates a new write-only instance of PropertyAsset, bound to a specific deployed contract.
func NewPropertyAssetTransactor(address common.Address, transactor bind.ContractTransactor) (*PropertyAssetTransactor, error) {
	contract, err := bindPropertyAsset(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PropertyAssetTransactor{contract: contract}, nil
}

// NewPropertyAssetFilterer creates a new log filterer instance of PropertyAsset, bound to a specific deployed contract.
func NewPropertyAssetFilterer(address common.Address, filterer bind.ContractFilterer) (*PropertyAssetFilterer, error) {
	contract, err := bindPropertyAsset(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PropertyAssetFilterer{contract: contract}, nil
}

// bindPropertyAsset binds a generic wrapper to an already deployed contract.
func bindPropertyAsset(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PropertyAssetMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PropertyAsset *PropertyAssetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PropertyAsset.Contract.PropertyAssetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PropertyAsset *PropertyAssetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PropertyAsset.Contract.PropertyAssetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PropertyAsset *PropertyAssetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PropertyAsset.Contract.PropertyAssetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PropertyAsset *PropertyAssetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PropertyAsset.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PropertyAsset *PropertyAssetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PropertyAsset.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PropertyAsset *PropertyAssetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PropertyAsset.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PropertyAsset *PropertyAssetCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PropertyAsset.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PropertyAsset *PropertyAssetSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _PropertyAsset.Contract.DEFAULTADMINROLE(&_PropertyAsset.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PropertyAsset *PropertyAssetCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _PropertyAsset.Contract.DEFAULTADMINROLE(&_PropertyAsset.CallOpts)
}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_PropertyAsset *PropertyAssetCaller) MANAGERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PropertyAsset.contract.Call(opts, &out, "MANAGER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_PropertyAsset *PropertyAssetSession) MANAGERROLE() ([32]byte, error) {
	return _PropertyAsset.Contract.MANAGERROLE(&_PropertyAsset.CallOpts)
}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_PropertyAsset *PropertyAssetCallerSession) MANAGERROLE() ([32]byte, error) {
	return _PropertyAsset.Contract.MANAGERROLE(&_PropertyAsset.CallOpts)
}

// AssetId is a free data retrieval call binding the contract method 0x44de240a.
//
// Solidity: function assetId() view returns(uint256)
func (_PropertyAsset *PropertyAssetCaller) AssetId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PropertyAsset.contract.Call(opts, &out, "assetId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AssetId is a free data retrieval call binding the contract method 0x44de240a.
//
// Solidity: function assetId() view returns(uint256)
func (_PropertyAsset *PropertyAssetSession) AssetId() (*big.Int, error) {
	return _PropertyAsset.Contract.AssetId(&_PropertyAsset.CallOpts)
}

// AssetId is a free data retrieval call binding the contract method 0x44de240a.
//
// Solidity: function assetId() view returns(uint256)
func (_PropertyAsset *PropertyAssetCallerSession) AssetId() (*big.Int, error) {
	return _PropertyAsset.Contract.AssetId(&_PropertyAsset.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_PropertyAsset *PropertyAssetCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PropertyAsset.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_PropertyAsset *PropertyAssetSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _PropertyAsset.Contract.BalanceOf(&_PropertyAsset.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_PropertyAsset *PropertyAssetCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _PropertyAsset.Contract.BalanceOf(&_PropertyAsset.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_PropertyAsset *PropertyAssetCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PropertyAsset.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_PropertyAsset *PropertyAssetSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _PropertyAsset.Contract.GetApproved(&_PropertyAsset.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_PropertyAsset *PropertyAssetCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _PropertyAsset.Contract.GetApproved(&_PropertyAsset.CallOpts, tokenId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PropertyAsset *PropertyAssetCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _PropertyAsset.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PropertyAsset *PropertyAssetSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _PropertyAsset.Contract.GetRoleAdmin(&_PropertyAsset.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PropertyAsset *PropertyAssetCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _PropertyAsset.Contract.GetRoleAdmin(&_PropertyAsset.CallOpts, role)
}

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() view returns(uint8)
func (_PropertyAsset *PropertyAssetCaller) GetStatus(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _PropertyAsset.contract.Call(opts, &out, "getStatus")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() view returns(uint8)
func (_PropertyAsset *PropertyAssetSession) GetStatus() (uint8, error) {
	return _PropertyAsset.Contract.GetStatus(&_PropertyAsset.CallOpts)
}

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() view returns(uint8)
func (_PropertyAsset *PropertyAssetCallerSession) GetStatus() (uint8, error) {
	return _PropertyAsset.Contract.GetStatus(&_PropertyAsset.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PropertyAsset *PropertyAssetCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _PropertyAsset.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PropertyAsset *PropertyAssetSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _PropertyAsset.Contract.HasRole(&_PropertyAsset.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PropertyAsset *PropertyAssetCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _PropertyAsset.Contract.HasRole(&_PropertyAsset.CallOpts, role, account)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_PropertyAsset *PropertyAssetCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _PropertyAsset.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_PropertyAsset *PropertyAssetSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _PropertyAsset.Contract.IsApprovedForAll(&_PropertyAsset.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_PropertyAsset *PropertyAssetCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _PropertyAsset.Contract.IsApprovedForAll(&_PropertyAsset.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PropertyAsset *PropertyAssetCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PropertyAsset.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PropertyAsset *PropertyAssetSession) Name() (string, error) {
	return _PropertyAsset.Contract.Name(&_PropertyAsset.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PropertyAsset *PropertyAssetCallerSession) Name() (string, error) {
	return _PropertyAsset.Contract.Name(&_PropertyAsset.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_PropertyAsset *PropertyAssetCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PropertyAsset.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_PropertyAsset *PropertyAssetSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _PropertyAsset.Contract.OwnerOf(&_PropertyAsset.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_PropertyAsset *PropertyAssetCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _PropertyAsset.Contract.OwnerOf(&_PropertyAsset.CallOpts, tokenId)
}

// PropertyDataHash is a free data retrieval call binding the contract method 0xf5a1f56e.
//
// Solidity: function propertyDataHash() view returns(string)
func (_PropertyAsset *PropertyAssetCaller) PropertyDataHash(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PropertyAsset.contract.Call(opts, &out, "propertyDataHash")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// PropertyDataHash is a free data retrieval call binding the contract method 0xf5a1f56e.
//
// Solidity: function propertyDataHash() view returns(string)
func (_PropertyAsset *PropertyAssetSession) PropertyDataHash() (string, error) {
	return _PropertyAsset.Contract.PropertyDataHash(&_PropertyAsset.CallOpts)
}

// PropertyDataHash is a free data retrieval call binding the contract method 0xf5a1f56e.
//
// Solidity: function propertyDataHash() view returns(string)
func (_PropertyAsset *PropertyAssetCallerSession) PropertyDataHash() (string, error) {
	return _PropertyAsset.Contract.PropertyDataHash(&_PropertyAsset.CallOpts)
}

// PropertyToken is a free data retrieval call binding the contract method 0x43ae2695.
//
// Solidity: function propertyToken() view returns(address)
func (_PropertyAsset *PropertyAssetCaller) PropertyToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PropertyAsset.contract.Call(opts, &out, "propertyToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PropertyToken is a free data retrieval call binding the contract method 0x43ae2695.
//
// Solidity: function propertyToken() view returns(address)
func (_PropertyAsset *PropertyAssetSession) PropertyToken() (common.Address, error) {
	return _PropertyAsset.Contract.PropertyToken(&_PropertyAsset.CallOpts)
}

// PropertyToken is a free data retrieval call binding the contract method 0x43ae2695.
//
// Solidity: function propertyToken() view returns(address)
func (_PropertyAsset *PropertyAssetCallerSession) PropertyToken() (common.Address, error) {
	return _PropertyAsset.Contract.PropertyToken(&_PropertyAsset.CallOpts)
}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() view returns(uint8)
func (_PropertyAsset *PropertyAssetCaller) Status(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _PropertyAsset.contract.Call(opts, &out, "status")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() view returns(uint8)
func (_PropertyAsset *PropertyAssetSession) Status() (uint8, error) {
	return _PropertyAsset.Contract.Status(&_PropertyAsset.CallOpts)
}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() view returns(uint8)
func (_PropertyAsset *PropertyAssetCallerSession) Status() (uint8, error) {
	return _PropertyAsset.Contract.Status(&_PropertyAsset.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PropertyAsset *PropertyAssetCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _PropertyAsset.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PropertyAsset *PropertyAssetSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _PropertyAsset.Contract.SupportsInterface(&_PropertyAsset.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PropertyAsset *PropertyAssetCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _PropertyAsset.Contract.SupportsInterface(&_PropertyAsset.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PropertyAsset *PropertyAssetCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PropertyAsset.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PropertyAsset *PropertyAssetSession) Symbol() (string, error) {
	return _PropertyAsset.Contract.Symbol(&_PropertyAsset.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PropertyAsset *PropertyAssetCallerSession) Symbol() (string, error) {
	return _PropertyAsset.Contract.Symbol(&_PropertyAsset.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_PropertyAsset *PropertyAssetCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _PropertyAsset.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_PropertyAsset *PropertyAssetSession) TokenURI(tokenId *big.Int) (string, error) {
	return _PropertyAsset.Contract.TokenURI(&_PropertyAsset.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_PropertyAsset *PropertyAssetCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _PropertyAsset.Contract.TokenURI(&_PropertyAsset.CallOpts, tokenId)
}

// Valuation is a free data retrieval call binding the contract method 0x21a7cfe4.
//
// Solidity: function valuation() view returns(uint256)
func (_PropertyAsset *PropertyAssetCaller) Valuation(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PropertyAsset.contract.Call(opts, &out, "valuation")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Valuation is a free data retrieval call binding the contract method 0x21a7cfe4.
//
// Solidity: function valuation() view returns(uint256)
func (_PropertyAsset *PropertyAssetSession) Valuation() (*big.Int, error) {
	return _PropertyAsset.Contract.Valuation(&_PropertyAsset.CallOpts)
}

// Valuation is a free data retrieval call binding the contract method 0x21a7cfe4.
//
// Solidity: function valuation() view returns(uint256)
func (_PropertyAsset *PropertyAssetCallerSession) Valuation() (*big.Int, error) {
	return _PropertyAsset.Contract.Valuation(&_PropertyAsset.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_PropertyAsset *PropertyAssetTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _PropertyAsset.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_PropertyAsset *PropertyAssetSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _PropertyAsset.Contract.Approve(&_PropertyAsset.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_PropertyAsset *PropertyAssetTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _PropertyAsset.Contract.Approve(&_PropertyAsset.TransactOpts, to, tokenId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PropertyAsset *PropertyAssetTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PropertyAsset.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PropertyAsset *PropertyAssetSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PropertyAsset.Contract.GrantRole(&_PropertyAsset.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PropertyAsset *PropertyAssetTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PropertyAsset.Contract.GrantRole(&_PropertyAsset.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_PropertyAsset *PropertyAssetTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PropertyAsset.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_PropertyAsset *PropertyAssetSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PropertyAsset.Contract.RenounceRole(&_PropertyAsset.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_PropertyAsset *PropertyAssetTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PropertyAsset.Contract.RenounceRole(&_PropertyAsset.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PropertyAsset *PropertyAssetTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PropertyAsset.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PropertyAsset *PropertyAssetSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PropertyAsset.Contract.RevokeRole(&_PropertyAsset.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PropertyAsset *PropertyAssetTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PropertyAsset.Contract.RevokeRole(&_PropertyAsset.TransactOpts, role, account)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_PropertyAsset *PropertyAssetTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _PropertyAsset.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_PropertyAsset *PropertyAssetSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _PropertyAsset.Contract.SafeTransferFrom(&_PropertyAsset.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_PropertyAsset *PropertyAssetTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _PropertyAsset.Contract.SafeTransferFrom(&_PropertyAsset.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_PropertyAsset *PropertyAssetTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _PropertyAsset.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_PropertyAsset *PropertyAssetSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _PropertyAsset.Contract.SafeTransferFrom0(&_PropertyAsset.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_PropertyAsset *PropertyAssetTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _PropertyAsset.Contract.SafeTransferFrom0(&_PropertyAsset.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_PropertyAsset *PropertyAssetTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _PropertyAsset.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_PropertyAsset *PropertyAssetSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _PropertyAsset.Contract.SetApprovalForAll(&_PropertyAsset.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_PropertyAsset *PropertyAssetTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _PropertyAsset.Contract.SetApprovalForAll(&_PropertyAsset.TransactOpts, operator, approved)
}

// SetPropertyToken is a paid mutator transaction binding the contract method 0x63065557.
//
// Solidity: function setPropertyToken(address token) returns()
func (_PropertyAsset *PropertyAssetTransactor) SetPropertyToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _PropertyAsset.contract.Transact(opts, "setPropertyToken", token)
}

// SetPropertyToken is a paid mutator transaction binding the contract method 0x63065557.
//
// Solidity: function setPropertyToken(address token) returns()
func (_PropertyAsset *PropertyAssetSession) SetPropertyToken(token common.Address) (*types.Transaction, error) {
	return _PropertyAsset.Contract.SetPropertyToken(&_PropertyAsset.TransactOpts, token)
}

// SetPropertyToken is a paid mutator transaction binding the contract method 0x63065557.
//
// Solidity: function setPropertyToken(address token) returns()
func (_PropertyAsset *PropertyAssetTransactorSession) SetPropertyToken(token common.Address) (*types.Transaction, error) {
	return _PropertyAsset.Contract.SetPropertyToken(&_PropertyAsset.TransactOpts, token)
}

// SetStatus is a paid mutator transaction binding the contract method 0x2e49d78b.
//
// Solidity: function setStatus(uint8 newStatus) returns()
func (_PropertyAsset *PropertyAssetTransactor) SetStatus(opts *bind.TransactOpts, newStatus uint8) (*types.Transaction, error) {
	return _PropertyAsset.contract.Transact(opts, "setStatus", newStatus)
}

// SetStatus is a paid mutator transaction binding the contract method 0x2e49d78b.
//
// Solidity: function setStatus(uint8 newStatus) returns()
func (_PropertyAsset *PropertyAssetSession) SetStatus(newStatus uint8) (*types.Transaction, error) {
	return _PropertyAsset.Contract.SetStatus(&_PropertyAsset.TransactOpts, newStatus)
}

// SetStatus is a paid mutator transaction binding the contract method 0x2e49d78b.
//
// Solidity: function setStatus(uint8 newStatus) returns()
func (_PropertyAsset *PropertyAssetTransactorSession) SetStatus(newStatus uint8) (*types.Transaction, error) {
	return _PropertyAsset.Contract.SetStatus(&_PropertyAsset.TransactOpts, newStatus)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_PropertyAsset *PropertyAssetTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _PropertyAsset.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_PropertyAsset *PropertyAssetSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _PropertyAsset.Contract.TransferFrom(&_PropertyAsset.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_PropertyAsset *PropertyAssetTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _PropertyAsset.Contract.TransferFrom(&_PropertyAsset.TransactOpts, from, to, tokenId)
}

// PropertyAssetApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the PropertyAsset contract.
type PropertyAssetApprovalIterator struct {
	Event *PropertyAssetApproval // Event containing the contract specifics and raw log

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
func (it *PropertyAssetApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PropertyAssetApproval)
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
		it.Event = new(PropertyAssetApproval)
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
func (it *PropertyAssetApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PropertyAssetApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PropertyAssetApproval represents a Approval event raised by the PropertyAsset contract.
type PropertyAssetApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_PropertyAsset *PropertyAssetFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*PropertyAssetApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _PropertyAsset.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &PropertyAssetApprovalIterator{contract: _PropertyAsset.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_PropertyAsset *PropertyAssetFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PropertyAssetApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _PropertyAsset.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PropertyAssetApproval)
				if err := _PropertyAsset.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_PropertyAsset *PropertyAssetFilterer) ParseApproval(log types.Log) (*PropertyAssetApproval, error) {
	event := new(PropertyAssetApproval)
	if err := _PropertyAsset.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PropertyAssetApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the PropertyAsset contract.
type PropertyAssetApprovalForAllIterator struct {
	Event *PropertyAssetApprovalForAll // Event containing the contract specifics and raw log

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
func (it *PropertyAssetApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PropertyAssetApprovalForAll)
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
		it.Event = new(PropertyAssetApprovalForAll)
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
func (it *PropertyAssetApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PropertyAssetApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PropertyAssetApprovalForAll represents a ApprovalForAll event raised by the PropertyAsset contract.
type PropertyAssetApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_PropertyAsset *PropertyAssetFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*PropertyAssetApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _PropertyAsset.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &PropertyAssetApprovalForAllIterator{contract: _PropertyAsset.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_PropertyAsset *PropertyAssetFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *PropertyAssetApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _PropertyAsset.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PropertyAssetApprovalForAll)
				if err := _PropertyAsset.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_PropertyAsset *PropertyAssetFilterer) ParseApprovalForAll(log types.Log) (*PropertyAssetApprovalForAll, error) {
	event := new(PropertyAssetApprovalForAll)
	if err := _PropertyAsset.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PropertyAssetPropertyStatusChangedIterator is returned from FilterPropertyStatusChanged and is used to iterate over the raw logs and unpacked data for PropertyStatusChanged events raised by the PropertyAsset contract.
type PropertyAssetPropertyStatusChangedIterator struct {
	Event *PropertyAssetPropertyStatusChanged // Event containing the contract specifics and raw log

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
func (it *PropertyAssetPropertyStatusChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PropertyAssetPropertyStatusChanged)
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
		it.Event = new(PropertyAssetPropertyStatusChanged)
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
func (it *PropertyAssetPropertyStatusChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PropertyAssetPropertyStatusChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PropertyAssetPropertyStatusChanged represents a PropertyStatusChanged event raised by the PropertyAsset contract.
type PropertyAssetPropertyStatusChanged struct {
	NewStatus uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPropertyStatusChanged is a free log retrieval operation binding the contract event 0x43edd70e8ffe87f291d1c4defd1eb43ac7261af725fb3a8ad1b69abea32623e4.
//
// Solidity: event PropertyStatusChanged(uint8 newStatus)
func (_PropertyAsset *PropertyAssetFilterer) FilterPropertyStatusChanged(opts *bind.FilterOpts) (*PropertyAssetPropertyStatusChangedIterator, error) {

	logs, sub, err := _PropertyAsset.contract.FilterLogs(opts, "PropertyStatusChanged")
	if err != nil {
		return nil, err
	}
	return &PropertyAssetPropertyStatusChangedIterator{contract: _PropertyAsset.contract, event: "PropertyStatusChanged", logs: logs, sub: sub}, nil
}

// WatchPropertyStatusChanged is a free log subscription operation binding the contract event 0x43edd70e8ffe87f291d1c4defd1eb43ac7261af725fb3a8ad1b69abea32623e4.
//
// Solidity: event PropertyStatusChanged(uint8 newStatus)
func (_PropertyAsset *PropertyAssetFilterer) WatchPropertyStatusChanged(opts *bind.WatchOpts, sink chan<- *PropertyAssetPropertyStatusChanged) (event.Subscription, error) {

	logs, sub, err := _PropertyAsset.contract.WatchLogs(opts, "PropertyStatusChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PropertyAssetPropertyStatusChanged)
				if err := _PropertyAsset.contract.UnpackLog(event, "PropertyStatusChanged", log); err != nil {
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

// ParsePropertyStatusChanged is a log parse operation binding the contract event 0x43edd70e8ffe87f291d1c4defd1eb43ac7261af725fb3a8ad1b69abea32623e4.
//
// Solidity: event PropertyStatusChanged(uint8 newStatus)
func (_PropertyAsset *PropertyAssetFilterer) ParsePropertyStatusChanged(log types.Log) (*PropertyAssetPropertyStatusChanged, error) {
	event := new(PropertyAssetPropertyStatusChanged)
	if err := _PropertyAsset.contract.UnpackLog(event, "PropertyStatusChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PropertyAssetPropertyTokenLinkedIterator is returned from FilterPropertyTokenLinked and is used to iterate over the raw logs and unpacked data for PropertyTokenLinked events raised by the PropertyAsset contract.
type PropertyAssetPropertyTokenLinkedIterator struct {
	Event *PropertyAssetPropertyTokenLinked // Event containing the contract specifics and raw log

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
func (it *PropertyAssetPropertyTokenLinkedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PropertyAssetPropertyTokenLinked)
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
		it.Event = new(PropertyAssetPropertyTokenLinked)
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
func (it *PropertyAssetPropertyTokenLinkedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PropertyAssetPropertyTokenLinkedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PropertyAssetPropertyTokenLinked represents a PropertyTokenLinked event raised by the PropertyAsset contract.
type PropertyAssetPropertyTokenLinked struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPropertyTokenLinked is a free log retrieval operation binding the contract event 0x117a641ebd71f76e3dac70b5aeac84fd5d125c150048e77dddaef3f3cdc50738.
//
// Solidity: event PropertyTokenLinked(address token)
func (_PropertyAsset *PropertyAssetFilterer) FilterPropertyTokenLinked(opts *bind.FilterOpts) (*PropertyAssetPropertyTokenLinkedIterator, error) {

	logs, sub, err := _PropertyAsset.contract.FilterLogs(opts, "PropertyTokenLinked")
	if err != nil {
		return nil, err
	}
	return &PropertyAssetPropertyTokenLinkedIterator{contract: _PropertyAsset.contract, event: "PropertyTokenLinked", logs: logs, sub: sub}, nil
}

// WatchPropertyTokenLinked is a free log subscription operation binding the contract event 0x117a641ebd71f76e3dac70b5aeac84fd5d125c150048e77dddaef3f3cdc50738.
//
// Solidity: event PropertyTokenLinked(address token)
func (_PropertyAsset *PropertyAssetFilterer) WatchPropertyTokenLinked(opts *bind.WatchOpts, sink chan<- *PropertyAssetPropertyTokenLinked) (event.Subscription, error) {

	logs, sub, err := _PropertyAsset.contract.WatchLogs(opts, "PropertyTokenLinked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PropertyAssetPropertyTokenLinked)
				if err := _PropertyAsset.contract.UnpackLog(event, "PropertyTokenLinked", log); err != nil {
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

// ParsePropertyTokenLinked is a log parse operation binding the contract event 0x117a641ebd71f76e3dac70b5aeac84fd5d125c150048e77dddaef3f3cdc50738.
//
// Solidity: event PropertyTokenLinked(address token)
func (_PropertyAsset *PropertyAssetFilterer) ParsePropertyTokenLinked(log types.Log) (*PropertyAssetPropertyTokenLinked, error) {
	event := new(PropertyAssetPropertyTokenLinked)
	if err := _PropertyAsset.contract.UnpackLog(event, "PropertyTokenLinked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PropertyAssetRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the PropertyAsset contract.
type PropertyAssetRoleAdminChangedIterator struct {
	Event *PropertyAssetRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *PropertyAssetRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PropertyAssetRoleAdminChanged)
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
		it.Event = new(PropertyAssetRoleAdminChanged)
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
func (it *PropertyAssetRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PropertyAssetRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PropertyAssetRoleAdminChanged represents a RoleAdminChanged event raised by the PropertyAsset contract.
type PropertyAssetRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_PropertyAsset *PropertyAssetFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*PropertyAssetRoleAdminChangedIterator, error) {

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

	logs, sub, err := _PropertyAsset.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &PropertyAssetRoleAdminChangedIterator{contract: _PropertyAsset.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_PropertyAsset *PropertyAssetFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *PropertyAssetRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _PropertyAsset.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PropertyAssetRoleAdminChanged)
				if err := _PropertyAsset.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_PropertyAsset *PropertyAssetFilterer) ParseRoleAdminChanged(log types.Log) (*PropertyAssetRoleAdminChanged, error) {
	event := new(PropertyAssetRoleAdminChanged)
	if err := _PropertyAsset.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PropertyAssetRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the PropertyAsset contract.
type PropertyAssetRoleGrantedIterator struct {
	Event *PropertyAssetRoleGranted // Event containing the contract specifics and raw log

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
func (it *PropertyAssetRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PropertyAssetRoleGranted)
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
		it.Event = new(PropertyAssetRoleGranted)
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
func (it *PropertyAssetRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PropertyAssetRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PropertyAssetRoleGranted represents a RoleGranted event raised by the PropertyAsset contract.
type PropertyAssetRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_PropertyAsset *PropertyAssetFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PropertyAssetRoleGrantedIterator, error) {

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

	logs, sub, err := _PropertyAsset.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PropertyAssetRoleGrantedIterator{contract: _PropertyAsset.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_PropertyAsset *PropertyAssetFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *PropertyAssetRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PropertyAsset.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PropertyAssetRoleGranted)
				if err := _PropertyAsset.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_PropertyAsset *PropertyAssetFilterer) ParseRoleGranted(log types.Log) (*PropertyAssetRoleGranted, error) {
	event := new(PropertyAssetRoleGranted)
	if err := _PropertyAsset.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PropertyAssetRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the PropertyAsset contract.
type PropertyAssetRoleRevokedIterator struct {
	Event *PropertyAssetRoleRevoked // Event containing the contract specifics and raw log

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
func (it *PropertyAssetRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PropertyAssetRoleRevoked)
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
		it.Event = new(PropertyAssetRoleRevoked)
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
func (it *PropertyAssetRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PropertyAssetRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PropertyAssetRoleRevoked represents a RoleRevoked event raised by the PropertyAsset contract.
type PropertyAssetRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_PropertyAsset *PropertyAssetFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PropertyAssetRoleRevokedIterator, error) {

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

	logs, sub, err := _PropertyAsset.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PropertyAssetRoleRevokedIterator{contract: _PropertyAsset.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_PropertyAsset *PropertyAssetFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *PropertyAssetRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PropertyAsset.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PropertyAssetRoleRevoked)
				if err := _PropertyAsset.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_PropertyAsset *PropertyAssetFilterer) ParseRoleRevoked(log types.Log) (*PropertyAssetRoleRevoked, error) {
	event := new(PropertyAssetRoleRevoked)
	if err := _PropertyAsset.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PropertyAssetTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the PropertyAsset contract.
type PropertyAssetTransferIterator struct {
	Event *PropertyAssetTransfer // Event containing the contract specifics and raw log

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
func (it *PropertyAssetTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PropertyAssetTransfer)
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
		it.Event = new(PropertyAssetTransfer)
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
func (it *PropertyAssetTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PropertyAssetTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PropertyAssetTransfer represents a Transfer event raised by the PropertyAsset contract.
type PropertyAssetTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_PropertyAsset *PropertyAssetFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*PropertyAssetTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _PropertyAsset.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &PropertyAssetTransferIterator{contract: _PropertyAsset.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_PropertyAsset *PropertyAssetFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PropertyAssetTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _PropertyAsset.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PropertyAssetTransfer)
				if err := _PropertyAsset.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_PropertyAsset *PropertyAssetFilterer) ParseTransfer(log types.Log) (*PropertyAssetTransfer, error) {
	event := new(PropertyAssetTransfer)
	if err := _PropertyAsset.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

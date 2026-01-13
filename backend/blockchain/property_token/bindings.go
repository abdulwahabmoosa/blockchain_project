// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package property_token

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

// PropertyTokenMetaData contains all meta data concerning the PropertyToken contract.
var PropertyTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"admin_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_propertyAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_approvalService\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"approval\",\"type\":\"address\"}],\"name\":\"ApprovalServiceLinked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"PropertyAssetLinked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"Snapshot\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SNAPSHOT_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"approvalService\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"snapshotId\",\"type\":\"uint256\"}],\"name\":\"balanceOfAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"propertyAsset\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"snapshotNow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"snapshotId\",\"type\":\"uint256\"}],\"name\":\"totalSupplyAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PropertyTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use PropertyTokenMetaData.ABI instead.
var PropertyTokenABI = PropertyTokenMetaData.ABI

// PropertyToken is an auto generated Go binding around an Ethereum contract.
type PropertyToken struct {
	PropertyTokenCaller     // Read-only binding to the contract
	PropertyTokenTransactor // Write-only binding to the contract
	PropertyTokenFilterer   // Log filterer for contract events
}

// PropertyTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type PropertyTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PropertyTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PropertyTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PropertyTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PropertyTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PropertyTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PropertyTokenSession struct {
	Contract     *PropertyToken    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PropertyTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PropertyTokenCallerSession struct {
	Contract *PropertyTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// PropertyTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PropertyTokenTransactorSession struct {
	Contract     *PropertyTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// PropertyTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type PropertyTokenRaw struct {
	Contract *PropertyToken // Generic contract binding to access the raw methods on
}

// PropertyTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PropertyTokenCallerRaw struct {
	Contract *PropertyTokenCaller // Generic read-only contract binding to access the raw methods on
}

// PropertyTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PropertyTokenTransactorRaw struct {
	Contract *PropertyTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPropertyToken creates a new instance of PropertyToken, bound to a specific deployed contract.
func NewPropertyToken(address common.Address, backend bind.ContractBackend) (*PropertyToken, error) {
	contract, err := bindPropertyToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PropertyToken{PropertyTokenCaller: PropertyTokenCaller{contract: contract}, PropertyTokenTransactor: PropertyTokenTransactor{contract: contract}, PropertyTokenFilterer: PropertyTokenFilterer{contract: contract}}, nil
}

// NewPropertyTokenCaller creates a new read-only instance of PropertyToken, bound to a specific deployed contract.
func NewPropertyTokenCaller(address common.Address, caller bind.ContractCaller) (*PropertyTokenCaller, error) {
	contract, err := bindPropertyToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PropertyTokenCaller{contract: contract}, nil
}

// NewPropertyTokenTransactor creates a new write-only instance of PropertyToken, bound to a specific deployed contract.
func NewPropertyTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*PropertyTokenTransactor, error) {
	contract, err := bindPropertyToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PropertyTokenTransactor{contract: contract}, nil
}

// NewPropertyTokenFilterer creates a new log filterer instance of PropertyToken, bound to a specific deployed contract.
func NewPropertyTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*PropertyTokenFilterer, error) {
	contract, err := bindPropertyToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PropertyTokenFilterer{contract: contract}, nil
}

// bindPropertyToken binds a generic wrapper to an already deployed contract.
func bindPropertyToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PropertyTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PropertyToken *PropertyTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PropertyToken.Contract.PropertyTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PropertyToken *PropertyTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PropertyToken.Contract.PropertyTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PropertyToken *PropertyTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PropertyToken.Contract.PropertyTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PropertyToken *PropertyTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PropertyToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PropertyToken *PropertyTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PropertyToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PropertyToken *PropertyTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PropertyToken.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PropertyToken *PropertyTokenCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PropertyToken.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PropertyToken *PropertyTokenSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _PropertyToken.Contract.DEFAULTADMINROLE(&_PropertyToken.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PropertyToken *PropertyTokenCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _PropertyToken.Contract.DEFAULTADMINROLE(&_PropertyToken.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_PropertyToken *PropertyTokenCaller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PropertyToken.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_PropertyToken *PropertyTokenSession) MINTERROLE() ([32]byte, error) {
	return _PropertyToken.Contract.MINTERROLE(&_PropertyToken.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_PropertyToken *PropertyTokenCallerSession) MINTERROLE() ([32]byte, error) {
	return _PropertyToken.Contract.MINTERROLE(&_PropertyToken.CallOpts)
}

// SNAPSHOTROLE is a free data retrieval call binding the contract method 0x7028e2cd.
//
// Solidity: function SNAPSHOT_ROLE() view returns(bytes32)
func (_PropertyToken *PropertyTokenCaller) SNAPSHOTROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PropertyToken.contract.Call(opts, &out, "SNAPSHOT_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SNAPSHOTROLE is a free data retrieval call binding the contract method 0x7028e2cd.
//
// Solidity: function SNAPSHOT_ROLE() view returns(bytes32)
func (_PropertyToken *PropertyTokenSession) SNAPSHOTROLE() ([32]byte, error) {
	return _PropertyToken.Contract.SNAPSHOTROLE(&_PropertyToken.CallOpts)
}

// SNAPSHOTROLE is a free data retrieval call binding the contract method 0x7028e2cd.
//
// Solidity: function SNAPSHOT_ROLE() view returns(bytes32)
func (_PropertyToken *PropertyTokenCallerSession) SNAPSHOTROLE() ([32]byte, error) {
	return _PropertyToken.Contract.SNAPSHOTROLE(&_PropertyToken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_PropertyToken *PropertyTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PropertyToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_PropertyToken *PropertyTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _PropertyToken.Contract.Allowance(&_PropertyToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_PropertyToken *PropertyTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _PropertyToken.Contract.Allowance(&_PropertyToken.CallOpts, owner, spender)
}

// ApprovalService is a free data retrieval call binding the contract method 0xf23d0c2f.
//
// Solidity: function approvalService() view returns(address)
func (_PropertyToken *PropertyTokenCaller) ApprovalService(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PropertyToken.contract.Call(opts, &out, "approvalService")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ApprovalService is a free data retrieval call binding the contract method 0xf23d0c2f.
//
// Solidity: function approvalService() view returns(address)
func (_PropertyToken *PropertyTokenSession) ApprovalService() (common.Address, error) {
	return _PropertyToken.Contract.ApprovalService(&_PropertyToken.CallOpts)
}

// ApprovalService is a free data retrieval call binding the contract method 0xf23d0c2f.
//
// Solidity: function approvalService() view returns(address)
func (_PropertyToken *PropertyTokenCallerSession) ApprovalService() (common.Address, error) {
	return _PropertyToken.Contract.ApprovalService(&_PropertyToken.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_PropertyToken *PropertyTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PropertyToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_PropertyToken *PropertyTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _PropertyToken.Contract.BalanceOf(&_PropertyToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_PropertyToken *PropertyTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _PropertyToken.Contract.BalanceOf(&_PropertyToken.CallOpts, account)
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address account, uint256 snapshotId) view returns(uint256)
func (_PropertyToken *PropertyTokenCaller) BalanceOfAt(opts *bind.CallOpts, account common.Address, snapshotId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PropertyToken.contract.Call(opts, &out, "balanceOfAt", account, snapshotId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address account, uint256 snapshotId) view returns(uint256)
func (_PropertyToken *PropertyTokenSession) BalanceOfAt(account common.Address, snapshotId *big.Int) (*big.Int, error) {
	return _PropertyToken.Contract.BalanceOfAt(&_PropertyToken.CallOpts, account, snapshotId)
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address account, uint256 snapshotId) view returns(uint256)
func (_PropertyToken *PropertyTokenCallerSession) BalanceOfAt(account common.Address, snapshotId *big.Int) (*big.Int, error) {
	return _PropertyToken.Contract.BalanceOfAt(&_PropertyToken.CallOpts, account, snapshotId)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_PropertyToken *PropertyTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _PropertyToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_PropertyToken *PropertyTokenSession) Decimals() (uint8, error) {
	return _PropertyToken.Contract.Decimals(&_PropertyToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_PropertyToken *PropertyTokenCallerSession) Decimals() (uint8, error) {
	return _PropertyToken.Contract.Decimals(&_PropertyToken.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PropertyToken *PropertyTokenCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _PropertyToken.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PropertyToken *PropertyTokenSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _PropertyToken.Contract.GetRoleAdmin(&_PropertyToken.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PropertyToken *PropertyTokenCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _PropertyToken.Contract.GetRoleAdmin(&_PropertyToken.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PropertyToken *PropertyTokenCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _PropertyToken.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PropertyToken *PropertyTokenSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _PropertyToken.Contract.HasRole(&_PropertyToken.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PropertyToken *PropertyTokenCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _PropertyToken.Contract.HasRole(&_PropertyToken.CallOpts, role, account)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PropertyToken *PropertyTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PropertyToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PropertyToken *PropertyTokenSession) Name() (string, error) {
	return _PropertyToken.Contract.Name(&_PropertyToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PropertyToken *PropertyTokenCallerSession) Name() (string, error) {
	return _PropertyToken.Contract.Name(&_PropertyToken.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PropertyToken *PropertyTokenCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PropertyToken.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PropertyToken *PropertyTokenSession) Paused() (bool, error) {
	return _PropertyToken.Contract.Paused(&_PropertyToken.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PropertyToken *PropertyTokenCallerSession) Paused() (bool, error) {
	return _PropertyToken.Contract.Paused(&_PropertyToken.CallOpts)
}

// PropertyAsset is a free data retrieval call binding the contract method 0xf3b564d1.
//
// Solidity: function propertyAsset() view returns(address)
func (_PropertyToken *PropertyTokenCaller) PropertyAsset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PropertyToken.contract.Call(opts, &out, "propertyAsset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PropertyAsset is a free data retrieval call binding the contract method 0xf3b564d1.
//
// Solidity: function propertyAsset() view returns(address)
func (_PropertyToken *PropertyTokenSession) PropertyAsset() (common.Address, error) {
	return _PropertyToken.Contract.PropertyAsset(&_PropertyToken.CallOpts)
}

// PropertyAsset is a free data retrieval call binding the contract method 0xf3b564d1.
//
// Solidity: function propertyAsset() view returns(address)
func (_PropertyToken *PropertyTokenCallerSession) PropertyAsset() (common.Address, error) {
	return _PropertyToken.Contract.PropertyAsset(&_PropertyToken.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PropertyToken *PropertyTokenCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _PropertyToken.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PropertyToken *PropertyTokenSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _PropertyToken.Contract.SupportsInterface(&_PropertyToken.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PropertyToken *PropertyTokenCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _PropertyToken.Contract.SupportsInterface(&_PropertyToken.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PropertyToken *PropertyTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PropertyToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PropertyToken *PropertyTokenSession) Symbol() (string, error) {
	return _PropertyToken.Contract.Symbol(&_PropertyToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PropertyToken *PropertyTokenCallerSession) Symbol() (string, error) {
	return _PropertyToken.Contract.Symbol(&_PropertyToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PropertyToken *PropertyTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PropertyToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PropertyToken *PropertyTokenSession) TotalSupply() (*big.Int, error) {
	return _PropertyToken.Contract.TotalSupply(&_PropertyToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PropertyToken *PropertyTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _PropertyToken.Contract.TotalSupply(&_PropertyToken.CallOpts)
}

// TotalSupplyAt is a free data retrieval call binding the contract method 0x981b24d0.
//
// Solidity: function totalSupplyAt(uint256 snapshotId) view returns(uint256)
func (_PropertyToken *PropertyTokenCaller) TotalSupplyAt(opts *bind.CallOpts, snapshotId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PropertyToken.contract.Call(opts, &out, "totalSupplyAt", snapshotId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupplyAt is a free data retrieval call binding the contract method 0x981b24d0.
//
// Solidity: function totalSupplyAt(uint256 snapshotId) view returns(uint256)
func (_PropertyToken *PropertyTokenSession) TotalSupplyAt(snapshotId *big.Int) (*big.Int, error) {
	return _PropertyToken.Contract.TotalSupplyAt(&_PropertyToken.CallOpts, snapshotId)
}

// TotalSupplyAt is a free data retrieval call binding the contract method 0x981b24d0.
//
// Solidity: function totalSupplyAt(uint256 snapshotId) view returns(uint256)
func (_PropertyToken *PropertyTokenCallerSession) TotalSupplyAt(snapshotId *big.Int) (*big.Int, error) {
	return _PropertyToken.Contract.TotalSupplyAt(&_PropertyToken.CallOpts, snapshotId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_PropertyToken *PropertyTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PropertyToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_PropertyToken *PropertyTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PropertyToken.Contract.Approve(&_PropertyToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_PropertyToken *PropertyTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PropertyToken.Contract.Approve(&_PropertyToken.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_PropertyToken *PropertyTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _PropertyToken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_PropertyToken *PropertyTokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _PropertyToken.Contract.DecreaseAllowance(&_PropertyToken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_PropertyToken *PropertyTokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _PropertyToken.Contract.DecreaseAllowance(&_PropertyToken.TransactOpts, spender, subtractedValue)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PropertyToken *PropertyTokenTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PropertyToken.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PropertyToken *PropertyTokenSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PropertyToken.Contract.GrantRole(&_PropertyToken.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PropertyToken *PropertyTokenTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PropertyToken.Contract.GrantRole(&_PropertyToken.TransactOpts, role, account)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_PropertyToken *PropertyTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _PropertyToken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_PropertyToken *PropertyTokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _PropertyToken.Contract.IncreaseAllowance(&_PropertyToken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_PropertyToken *PropertyTokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _PropertyToken.Contract.IncreaseAllowance(&_PropertyToken.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_PropertyToken *PropertyTokenTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PropertyToken.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_PropertyToken *PropertyTokenSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PropertyToken.Contract.Mint(&_PropertyToken.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_PropertyToken *PropertyTokenTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PropertyToken.Contract.Mint(&_PropertyToken.TransactOpts, to, amount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PropertyToken *PropertyTokenTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PropertyToken.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PropertyToken *PropertyTokenSession) Pause() (*types.Transaction, error) {
	return _PropertyToken.Contract.Pause(&_PropertyToken.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PropertyToken *PropertyTokenTransactorSession) Pause() (*types.Transaction, error) {
	return _PropertyToken.Contract.Pause(&_PropertyToken.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_PropertyToken *PropertyTokenTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PropertyToken.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_PropertyToken *PropertyTokenSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PropertyToken.Contract.RenounceRole(&_PropertyToken.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_PropertyToken *PropertyTokenTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PropertyToken.Contract.RenounceRole(&_PropertyToken.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PropertyToken *PropertyTokenTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PropertyToken.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PropertyToken *PropertyTokenSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PropertyToken.Contract.RevokeRole(&_PropertyToken.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PropertyToken *PropertyTokenTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PropertyToken.Contract.RevokeRole(&_PropertyToken.TransactOpts, role, account)
}

// SnapshotNow is a paid mutator transaction binding the contract method 0xbe723e6e.
//
// Solidity: function snapshotNow() returns(uint256)
func (_PropertyToken *PropertyTokenTransactor) SnapshotNow(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PropertyToken.contract.Transact(opts, "snapshotNow")
}

// SnapshotNow is a paid mutator transaction binding the contract method 0xbe723e6e.
//
// Solidity: function snapshotNow() returns(uint256)
func (_PropertyToken *PropertyTokenSession) SnapshotNow() (*types.Transaction, error) {
	return _PropertyToken.Contract.SnapshotNow(&_PropertyToken.TransactOpts)
}

// SnapshotNow is a paid mutator transaction binding the contract method 0xbe723e6e.
//
// Solidity: function snapshotNow() returns(uint256)
func (_PropertyToken *PropertyTokenTransactorSession) SnapshotNow() (*types.Transaction, error) {
	return _PropertyToken.Contract.SnapshotNow(&_PropertyToken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_PropertyToken *PropertyTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PropertyToken.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_PropertyToken *PropertyTokenSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PropertyToken.Contract.Transfer(&_PropertyToken.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_PropertyToken *PropertyTokenTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PropertyToken.Contract.Transfer(&_PropertyToken.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_PropertyToken *PropertyTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PropertyToken.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_PropertyToken *PropertyTokenSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PropertyToken.Contract.TransferFrom(&_PropertyToken.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_PropertyToken *PropertyTokenTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PropertyToken.Contract.TransferFrom(&_PropertyToken.TransactOpts, from, to, amount)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PropertyToken *PropertyTokenTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PropertyToken.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PropertyToken *PropertyTokenSession) Unpause() (*types.Transaction, error) {
	return _PropertyToken.Contract.Unpause(&_PropertyToken.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PropertyToken *PropertyTokenTransactorSession) Unpause() (*types.Transaction, error) {
	return _PropertyToken.Contract.Unpause(&_PropertyToken.TransactOpts)
}

// PropertyTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the PropertyToken contract.
type PropertyTokenApprovalIterator struct {
	Event *PropertyTokenApproval // Event containing the contract specifics and raw log

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
func (it *PropertyTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PropertyTokenApproval)
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
		it.Event = new(PropertyTokenApproval)
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
func (it *PropertyTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PropertyTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PropertyTokenApproval represents a Approval event raised by the PropertyToken contract.
type PropertyTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_PropertyToken *PropertyTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*PropertyTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _PropertyToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &PropertyTokenApprovalIterator{contract: _PropertyToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_PropertyToken *PropertyTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PropertyTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _PropertyToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PropertyTokenApproval)
				if err := _PropertyToken.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_PropertyToken *PropertyTokenFilterer) ParseApproval(log types.Log) (*PropertyTokenApproval, error) {
	event := new(PropertyTokenApproval)
	if err := _PropertyToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PropertyTokenApprovalServiceLinkedIterator is returned from FilterApprovalServiceLinked and is used to iterate over the raw logs and unpacked data for ApprovalServiceLinked events raised by the PropertyToken contract.
type PropertyTokenApprovalServiceLinkedIterator struct {
	Event *PropertyTokenApprovalServiceLinked // Event containing the contract specifics and raw log

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
func (it *PropertyTokenApprovalServiceLinkedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PropertyTokenApprovalServiceLinked)
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
		it.Event = new(PropertyTokenApprovalServiceLinked)
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
func (it *PropertyTokenApprovalServiceLinkedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PropertyTokenApprovalServiceLinkedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PropertyTokenApprovalServiceLinked represents a ApprovalServiceLinked event raised by the PropertyToken contract.
type PropertyTokenApprovalServiceLinked struct {
	Approval common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalServiceLinked is a free log retrieval operation binding the contract event 0x60763e93bb64eec076e8ecf7071b5f891238624b3f022f845d03e12cac31f5e5.
//
// Solidity: event ApprovalServiceLinked(address approval)
func (_PropertyToken *PropertyTokenFilterer) FilterApprovalServiceLinked(opts *bind.FilterOpts) (*PropertyTokenApprovalServiceLinkedIterator, error) {

	logs, sub, err := _PropertyToken.contract.FilterLogs(opts, "ApprovalServiceLinked")
	if err != nil {
		return nil, err
	}
	return &PropertyTokenApprovalServiceLinkedIterator{contract: _PropertyToken.contract, event: "ApprovalServiceLinked", logs: logs, sub: sub}, nil
}

// WatchApprovalServiceLinked is a free log subscription operation binding the contract event 0x60763e93bb64eec076e8ecf7071b5f891238624b3f022f845d03e12cac31f5e5.
//
// Solidity: event ApprovalServiceLinked(address approval)
func (_PropertyToken *PropertyTokenFilterer) WatchApprovalServiceLinked(opts *bind.WatchOpts, sink chan<- *PropertyTokenApprovalServiceLinked) (event.Subscription, error) {

	logs, sub, err := _PropertyToken.contract.WatchLogs(opts, "ApprovalServiceLinked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PropertyTokenApprovalServiceLinked)
				if err := _PropertyToken.contract.UnpackLog(event, "ApprovalServiceLinked", log); err != nil {
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

// ParseApprovalServiceLinked is a log parse operation binding the contract event 0x60763e93bb64eec076e8ecf7071b5f891238624b3f022f845d03e12cac31f5e5.
//
// Solidity: event ApprovalServiceLinked(address approval)
func (_PropertyToken *PropertyTokenFilterer) ParseApprovalServiceLinked(log types.Log) (*PropertyTokenApprovalServiceLinked, error) {
	event := new(PropertyTokenApprovalServiceLinked)
	if err := _PropertyToken.contract.UnpackLog(event, "ApprovalServiceLinked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PropertyTokenPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the PropertyToken contract.
type PropertyTokenPausedIterator struct {
	Event *PropertyTokenPaused // Event containing the contract specifics and raw log

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
func (it *PropertyTokenPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PropertyTokenPaused)
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
		it.Event = new(PropertyTokenPaused)
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
func (it *PropertyTokenPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PropertyTokenPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PropertyTokenPaused represents a Paused event raised by the PropertyToken contract.
type PropertyTokenPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PropertyToken *PropertyTokenFilterer) FilterPaused(opts *bind.FilterOpts) (*PropertyTokenPausedIterator, error) {

	logs, sub, err := _PropertyToken.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PropertyTokenPausedIterator{contract: _PropertyToken.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PropertyToken *PropertyTokenFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PropertyTokenPaused) (event.Subscription, error) {

	logs, sub, err := _PropertyToken.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PropertyTokenPaused)
				if err := _PropertyToken.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PropertyToken *PropertyTokenFilterer) ParsePaused(log types.Log) (*PropertyTokenPaused, error) {
	event := new(PropertyTokenPaused)
	if err := _PropertyToken.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PropertyTokenPropertyAssetLinkedIterator is returned from FilterPropertyAssetLinked and is used to iterate over the raw logs and unpacked data for PropertyAssetLinked events raised by the PropertyToken contract.
type PropertyTokenPropertyAssetLinkedIterator struct {
	Event *PropertyTokenPropertyAssetLinked // Event containing the contract specifics and raw log

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
func (it *PropertyTokenPropertyAssetLinkedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PropertyTokenPropertyAssetLinked)
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
		it.Event = new(PropertyTokenPropertyAssetLinked)
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
func (it *PropertyTokenPropertyAssetLinkedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PropertyTokenPropertyAssetLinkedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PropertyTokenPropertyAssetLinked represents a PropertyAssetLinked event raised by the PropertyToken contract.
type PropertyTokenPropertyAssetLinked struct {
	Asset common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPropertyAssetLinked is a free log retrieval operation binding the contract event 0xbb1afc35b9af22510ceda1fcb658bb22be1c113e5eb596a3bc75ab0e075238e0.
//
// Solidity: event PropertyAssetLinked(address asset)
func (_PropertyToken *PropertyTokenFilterer) FilterPropertyAssetLinked(opts *bind.FilterOpts) (*PropertyTokenPropertyAssetLinkedIterator, error) {

	logs, sub, err := _PropertyToken.contract.FilterLogs(opts, "PropertyAssetLinked")
	if err != nil {
		return nil, err
	}
	return &PropertyTokenPropertyAssetLinkedIterator{contract: _PropertyToken.contract, event: "PropertyAssetLinked", logs: logs, sub: sub}, nil
}

// WatchPropertyAssetLinked is a free log subscription operation binding the contract event 0xbb1afc35b9af22510ceda1fcb658bb22be1c113e5eb596a3bc75ab0e075238e0.
//
// Solidity: event PropertyAssetLinked(address asset)
func (_PropertyToken *PropertyTokenFilterer) WatchPropertyAssetLinked(opts *bind.WatchOpts, sink chan<- *PropertyTokenPropertyAssetLinked) (event.Subscription, error) {

	logs, sub, err := _PropertyToken.contract.WatchLogs(opts, "PropertyAssetLinked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PropertyTokenPropertyAssetLinked)
				if err := _PropertyToken.contract.UnpackLog(event, "PropertyAssetLinked", log); err != nil {
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

// ParsePropertyAssetLinked is a log parse operation binding the contract event 0xbb1afc35b9af22510ceda1fcb658bb22be1c113e5eb596a3bc75ab0e075238e0.
//
// Solidity: event PropertyAssetLinked(address asset)
func (_PropertyToken *PropertyTokenFilterer) ParsePropertyAssetLinked(log types.Log) (*PropertyTokenPropertyAssetLinked, error) {
	event := new(PropertyTokenPropertyAssetLinked)
	if err := _PropertyToken.contract.UnpackLog(event, "PropertyAssetLinked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PropertyTokenRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the PropertyToken contract.
type PropertyTokenRoleAdminChangedIterator struct {
	Event *PropertyTokenRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *PropertyTokenRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PropertyTokenRoleAdminChanged)
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
		it.Event = new(PropertyTokenRoleAdminChanged)
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
func (it *PropertyTokenRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PropertyTokenRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PropertyTokenRoleAdminChanged represents a RoleAdminChanged event raised by the PropertyToken contract.
type PropertyTokenRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_PropertyToken *PropertyTokenFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*PropertyTokenRoleAdminChangedIterator, error) {

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

	logs, sub, err := _PropertyToken.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &PropertyTokenRoleAdminChangedIterator{contract: _PropertyToken.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_PropertyToken *PropertyTokenFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *PropertyTokenRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _PropertyToken.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PropertyTokenRoleAdminChanged)
				if err := _PropertyToken.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_PropertyToken *PropertyTokenFilterer) ParseRoleAdminChanged(log types.Log) (*PropertyTokenRoleAdminChanged, error) {
	event := new(PropertyTokenRoleAdminChanged)
	if err := _PropertyToken.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PropertyTokenRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the PropertyToken contract.
type PropertyTokenRoleGrantedIterator struct {
	Event *PropertyTokenRoleGranted // Event containing the contract specifics and raw log

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
func (it *PropertyTokenRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PropertyTokenRoleGranted)
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
		it.Event = new(PropertyTokenRoleGranted)
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
func (it *PropertyTokenRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PropertyTokenRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PropertyTokenRoleGranted represents a RoleGranted event raised by the PropertyToken contract.
type PropertyTokenRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_PropertyToken *PropertyTokenFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PropertyTokenRoleGrantedIterator, error) {

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

	logs, sub, err := _PropertyToken.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PropertyTokenRoleGrantedIterator{contract: _PropertyToken.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_PropertyToken *PropertyTokenFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *PropertyTokenRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PropertyToken.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PropertyTokenRoleGranted)
				if err := _PropertyToken.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_PropertyToken *PropertyTokenFilterer) ParseRoleGranted(log types.Log) (*PropertyTokenRoleGranted, error) {
	event := new(PropertyTokenRoleGranted)
	if err := _PropertyToken.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PropertyTokenRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the PropertyToken contract.
type PropertyTokenRoleRevokedIterator struct {
	Event *PropertyTokenRoleRevoked // Event containing the contract specifics and raw log

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
func (it *PropertyTokenRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PropertyTokenRoleRevoked)
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
		it.Event = new(PropertyTokenRoleRevoked)
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
func (it *PropertyTokenRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PropertyTokenRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PropertyTokenRoleRevoked represents a RoleRevoked event raised by the PropertyToken contract.
type PropertyTokenRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_PropertyToken *PropertyTokenFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PropertyTokenRoleRevokedIterator, error) {

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

	logs, sub, err := _PropertyToken.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PropertyTokenRoleRevokedIterator{contract: _PropertyToken.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_PropertyToken *PropertyTokenFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *PropertyTokenRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PropertyToken.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PropertyTokenRoleRevoked)
				if err := _PropertyToken.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_PropertyToken *PropertyTokenFilterer) ParseRoleRevoked(log types.Log) (*PropertyTokenRoleRevoked, error) {
	event := new(PropertyTokenRoleRevoked)
	if err := _PropertyToken.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PropertyTokenSnapshotIterator is returned from FilterSnapshot and is used to iterate over the raw logs and unpacked data for Snapshot events raised by the PropertyToken contract.
type PropertyTokenSnapshotIterator struct {
	Event *PropertyTokenSnapshot // Event containing the contract specifics and raw log

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
func (it *PropertyTokenSnapshotIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PropertyTokenSnapshot)
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
		it.Event = new(PropertyTokenSnapshot)
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
func (it *PropertyTokenSnapshotIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PropertyTokenSnapshotIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PropertyTokenSnapshot represents a Snapshot event raised by the PropertyToken contract.
type PropertyTokenSnapshot struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSnapshot is a free log retrieval operation binding the contract event 0x8030e83b04d87bef53480e26263266d6ca66863aa8506aca6f2559d18aa1cb67.
//
// Solidity: event Snapshot(uint256 id)
func (_PropertyToken *PropertyTokenFilterer) FilterSnapshot(opts *bind.FilterOpts) (*PropertyTokenSnapshotIterator, error) {

	logs, sub, err := _PropertyToken.contract.FilterLogs(opts, "Snapshot")
	if err != nil {
		return nil, err
	}
	return &PropertyTokenSnapshotIterator{contract: _PropertyToken.contract, event: "Snapshot", logs: logs, sub: sub}, nil
}

// WatchSnapshot is a free log subscription operation binding the contract event 0x8030e83b04d87bef53480e26263266d6ca66863aa8506aca6f2559d18aa1cb67.
//
// Solidity: event Snapshot(uint256 id)
func (_PropertyToken *PropertyTokenFilterer) WatchSnapshot(opts *bind.WatchOpts, sink chan<- *PropertyTokenSnapshot) (event.Subscription, error) {

	logs, sub, err := _PropertyToken.contract.WatchLogs(opts, "Snapshot")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PropertyTokenSnapshot)
				if err := _PropertyToken.contract.UnpackLog(event, "Snapshot", log); err != nil {
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

// ParseSnapshot is a log parse operation binding the contract event 0x8030e83b04d87bef53480e26263266d6ca66863aa8506aca6f2559d18aa1cb67.
//
// Solidity: event Snapshot(uint256 id)
func (_PropertyToken *PropertyTokenFilterer) ParseSnapshot(log types.Log) (*PropertyTokenSnapshot, error) {
	event := new(PropertyTokenSnapshot)
	if err := _PropertyToken.contract.UnpackLog(event, "Snapshot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PropertyTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the PropertyToken contract.
type PropertyTokenTransferIterator struct {
	Event *PropertyTokenTransfer // Event containing the contract specifics and raw log

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
func (it *PropertyTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PropertyTokenTransfer)
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
		it.Event = new(PropertyTokenTransfer)
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
func (it *PropertyTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PropertyTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PropertyTokenTransfer represents a Transfer event raised by the PropertyToken contract.
type PropertyTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_PropertyToken *PropertyTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*PropertyTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PropertyToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PropertyTokenTransferIterator{contract: _PropertyToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_PropertyToken *PropertyTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PropertyTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PropertyToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PropertyTokenTransfer)
				if err := _PropertyToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_PropertyToken *PropertyTokenFilterer) ParseTransfer(log types.Log) (*PropertyTokenTransfer, error) {
	event := new(PropertyTokenTransfer)
	if err := _PropertyToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PropertyTokenUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the PropertyToken contract.
type PropertyTokenUnpausedIterator struct {
	Event *PropertyTokenUnpaused // Event containing the contract specifics and raw log

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
func (it *PropertyTokenUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PropertyTokenUnpaused)
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
		it.Event = new(PropertyTokenUnpaused)
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
func (it *PropertyTokenUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PropertyTokenUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PropertyTokenUnpaused represents a Unpaused event raised by the PropertyToken contract.
type PropertyTokenUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PropertyToken *PropertyTokenFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PropertyTokenUnpausedIterator, error) {

	logs, sub, err := _PropertyToken.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PropertyTokenUnpausedIterator{contract: _PropertyToken.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PropertyToken *PropertyTokenFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PropertyTokenUnpaused) (event.Subscription, error) {

	logs, sub, err := _PropertyToken.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PropertyTokenUnpaused)
				if err := _PropertyToken.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PropertyToken *PropertyTokenFilterer) ParseUnpaused(log types.Log) (*PropertyTokenUnpaused, error) {
	event := new(PropertyTokenUnpaused)
	if err := _PropertyToken.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

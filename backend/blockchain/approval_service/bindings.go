// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package approval_service

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

// ApprovalServiceMetaData contains all meta data concerning the ApprovalService contract.
var ApprovalServiceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"Approved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"Revoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"check\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isApproved\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"revoke\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ApprovalServiceABI is the input ABI used to generate the binding from.
// Deprecated: Use ApprovalServiceMetaData.ABI instead.
var ApprovalServiceABI = ApprovalServiceMetaData.ABI

// ApprovalService is an auto generated Go binding around an Ethereum contract.
type ApprovalService struct {
	ApprovalServiceCaller     // Read-only binding to the contract
	ApprovalServiceTransactor // Write-only binding to the contract
	ApprovalServiceFilterer   // Log filterer for contract events
}

// ApprovalServiceCaller is an auto generated read-only Go binding around an Ethereum contract.
type ApprovalServiceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApprovalServiceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ApprovalServiceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApprovalServiceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ApprovalServiceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApprovalServiceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ApprovalServiceSession struct {
	Contract     *ApprovalService  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApprovalServiceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ApprovalServiceCallerSession struct {
	Contract *ApprovalServiceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ApprovalServiceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ApprovalServiceTransactorSession struct {
	Contract     *ApprovalServiceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ApprovalServiceRaw is an auto generated low-level Go binding around an Ethereum contract.
type ApprovalServiceRaw struct {
	Contract *ApprovalService // Generic contract binding to access the raw methods on
}

// ApprovalServiceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ApprovalServiceCallerRaw struct {
	Contract *ApprovalServiceCaller // Generic read-only contract binding to access the raw methods on
}

// ApprovalServiceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ApprovalServiceTransactorRaw struct {
	Contract *ApprovalServiceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewApprovalService creates a new instance of ApprovalService, bound to a specific deployed contract.
func NewApprovalService(address common.Address, backend bind.ContractBackend) (*ApprovalService, error) {
	contract, err := bindApprovalService(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ApprovalService{ApprovalServiceCaller: ApprovalServiceCaller{contract: contract}, ApprovalServiceTransactor: ApprovalServiceTransactor{contract: contract}, ApprovalServiceFilterer: ApprovalServiceFilterer{contract: contract}}, nil
}

// NewApprovalServiceCaller creates a new read-only instance of ApprovalService, bound to a specific deployed contract.
func NewApprovalServiceCaller(address common.Address, caller bind.ContractCaller) (*ApprovalServiceCaller, error) {
	contract, err := bindApprovalService(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ApprovalServiceCaller{contract: contract}, nil
}

// NewApprovalServiceTransactor creates a new write-only instance of ApprovalService, bound to a specific deployed contract.
func NewApprovalServiceTransactor(address common.Address, transactor bind.ContractTransactor) (*ApprovalServiceTransactor, error) {
	contract, err := bindApprovalService(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ApprovalServiceTransactor{contract: contract}, nil
}

// NewApprovalServiceFilterer creates a new log filterer instance of ApprovalService, bound to a specific deployed contract.
func NewApprovalServiceFilterer(address common.Address, filterer bind.ContractFilterer) (*ApprovalServiceFilterer, error) {
	contract, err := bindApprovalService(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ApprovalServiceFilterer{contract: contract}, nil
}

// bindApprovalService binds a generic wrapper to an already deployed contract.
func bindApprovalService(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ApprovalServiceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ApprovalService *ApprovalServiceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ApprovalService.Contract.ApprovalServiceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ApprovalService *ApprovalServiceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ApprovalService.Contract.ApprovalServiceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ApprovalService *ApprovalServiceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ApprovalService.Contract.ApprovalServiceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ApprovalService *ApprovalServiceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ApprovalService.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ApprovalService *ApprovalServiceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ApprovalService.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ApprovalService *ApprovalServiceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ApprovalService.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_ApprovalService *ApprovalServiceCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ApprovalService.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_ApprovalService *ApprovalServiceSession) ADMINROLE() ([32]byte, error) {
	return _ApprovalService.Contract.ADMINROLE(&_ApprovalService.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_ApprovalService *ApprovalServiceCallerSession) ADMINROLE() ([32]byte, error) {
	return _ApprovalService.Contract.ADMINROLE(&_ApprovalService.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ApprovalService *ApprovalServiceCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ApprovalService.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ApprovalService *ApprovalServiceSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ApprovalService.Contract.DEFAULTADMINROLE(&_ApprovalService.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ApprovalService *ApprovalServiceCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ApprovalService.Contract.DEFAULTADMINROLE(&_ApprovalService.CallOpts)
}

// Check is a free data retrieval call binding the contract method 0xc23697a8.
//
// Solidity: function check(address user) view returns(bool)
func (_ApprovalService *ApprovalServiceCaller) Check(opts *bind.CallOpts, user common.Address) (bool, error) {
	var out []interface{}
	err := _ApprovalService.contract.Call(opts, &out, "check", user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Check is a free data retrieval call binding the contract method 0xc23697a8.
//
// Solidity: function check(address user) view returns(bool)
func (_ApprovalService *ApprovalServiceSession) Check(user common.Address) (bool, error) {
	return _ApprovalService.Contract.Check(&_ApprovalService.CallOpts, user)
}

// Check is a free data retrieval call binding the contract method 0xc23697a8.
//
// Solidity: function check(address user) view returns(bool)
func (_ApprovalService *ApprovalServiceCallerSession) Check(user common.Address) (bool, error) {
	return _ApprovalService.Contract.Check(&_ApprovalService.CallOpts, user)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ApprovalService *ApprovalServiceCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ApprovalService.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ApprovalService *ApprovalServiceSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ApprovalService.Contract.GetRoleAdmin(&_ApprovalService.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ApprovalService *ApprovalServiceCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ApprovalService.Contract.GetRoleAdmin(&_ApprovalService.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ApprovalService *ApprovalServiceCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ApprovalService.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ApprovalService *ApprovalServiceSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ApprovalService.Contract.HasRole(&_ApprovalService.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ApprovalService *ApprovalServiceCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ApprovalService.Contract.HasRole(&_ApprovalService.CallOpts, role, account)
}

// IsApproved is a free data retrieval call binding the contract method 0x673448dd.
//
// Solidity: function isApproved(address ) view returns(bool)
func (_ApprovalService *ApprovalServiceCaller) IsApproved(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ApprovalService.contract.Call(opts, &out, "isApproved", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApproved is a free data retrieval call binding the contract method 0x673448dd.
//
// Solidity: function isApproved(address ) view returns(bool)
func (_ApprovalService *ApprovalServiceSession) IsApproved(arg0 common.Address) (bool, error) {
	return _ApprovalService.Contract.IsApproved(&_ApprovalService.CallOpts, arg0)
}

// IsApproved is a free data retrieval call binding the contract method 0x673448dd.
//
// Solidity: function isApproved(address ) view returns(bool)
func (_ApprovalService *ApprovalServiceCallerSession) IsApproved(arg0 common.Address) (bool, error) {
	return _ApprovalService.Contract.IsApproved(&_ApprovalService.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ApprovalService *ApprovalServiceCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ApprovalService.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ApprovalService *ApprovalServiceSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ApprovalService.Contract.SupportsInterface(&_ApprovalService.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ApprovalService *ApprovalServiceCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ApprovalService.Contract.SupportsInterface(&_ApprovalService.CallOpts, interfaceId)
}

// Approve is a paid mutator transaction binding the contract method 0xdaea85c5.
//
// Solidity: function approve(address user) returns()
func (_ApprovalService *ApprovalServiceTransactor) Approve(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _ApprovalService.contract.Transact(opts, "approve", user)
}

// Approve is a paid mutator transaction binding the contract method 0xdaea85c5.
//
// Solidity: function approve(address user) returns()
func (_ApprovalService *ApprovalServiceSession) Approve(user common.Address) (*types.Transaction, error) {
	return _ApprovalService.Contract.Approve(&_ApprovalService.TransactOpts, user)
}

// Approve is a paid mutator transaction binding the contract method 0xdaea85c5.
//
// Solidity: function approve(address user) returns()
func (_ApprovalService *ApprovalServiceTransactorSession) Approve(user common.Address) (*types.Transaction, error) {
	return _ApprovalService.Contract.Approve(&_ApprovalService.TransactOpts, user)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ApprovalService *ApprovalServiceTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ApprovalService.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ApprovalService *ApprovalServiceSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ApprovalService.Contract.GrantRole(&_ApprovalService.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ApprovalService *ApprovalServiceTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ApprovalService.Contract.GrantRole(&_ApprovalService.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ApprovalService *ApprovalServiceTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ApprovalService.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ApprovalService *ApprovalServiceSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ApprovalService.Contract.RenounceRole(&_ApprovalService.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ApprovalService *ApprovalServiceTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ApprovalService.Contract.RenounceRole(&_ApprovalService.TransactOpts, role, account)
}

// Revoke is a paid mutator transaction binding the contract method 0x74a8f103.
//
// Solidity: function revoke(address user) returns()
func (_ApprovalService *ApprovalServiceTransactor) Revoke(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _ApprovalService.contract.Transact(opts, "revoke", user)
}

// Revoke is a paid mutator transaction binding the contract method 0x74a8f103.
//
// Solidity: function revoke(address user) returns()
func (_ApprovalService *ApprovalServiceSession) Revoke(user common.Address) (*types.Transaction, error) {
	return _ApprovalService.Contract.Revoke(&_ApprovalService.TransactOpts, user)
}

// Revoke is a paid mutator transaction binding the contract method 0x74a8f103.
//
// Solidity: function revoke(address user) returns()
func (_ApprovalService *ApprovalServiceTransactorSession) Revoke(user common.Address) (*types.Transaction, error) {
	return _ApprovalService.Contract.Revoke(&_ApprovalService.TransactOpts, user)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ApprovalService *ApprovalServiceTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ApprovalService.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ApprovalService *ApprovalServiceSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ApprovalService.Contract.RevokeRole(&_ApprovalService.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ApprovalService *ApprovalServiceTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ApprovalService.Contract.RevokeRole(&_ApprovalService.TransactOpts, role, account)
}

// ApprovalServiceApprovedIterator is returned from FilterApproved and is used to iterate over the raw logs and unpacked data for Approved events raised by the ApprovalService contract.
type ApprovalServiceApprovedIterator struct {
	Event *ApprovalServiceApproved // Event containing the contract specifics and raw log

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
func (it *ApprovalServiceApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApprovalServiceApproved)
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
		it.Event = new(ApprovalServiceApproved)
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
func (it *ApprovalServiceApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApprovalServiceApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApprovalServiceApproved represents a Approved event raised by the ApprovalService contract.
type ApprovalServiceApproved struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterApproved is a free log retrieval operation binding the contract event 0x5d91bd0cecc45fef102af61de92c5462fadc884a5ce9d21c15e8a85198f2349e.
//
// Solidity: event Approved(address indexed user)
func (_ApprovalService *ApprovalServiceFilterer) FilterApproved(opts *bind.FilterOpts, user []common.Address) (*ApprovalServiceApprovedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ApprovalService.contract.FilterLogs(opts, "Approved", userRule)
	if err != nil {
		return nil, err
	}
	return &ApprovalServiceApprovedIterator{contract: _ApprovalService.contract, event: "Approved", logs: logs, sub: sub}, nil
}

// WatchApproved is a free log subscription operation binding the contract event 0x5d91bd0cecc45fef102af61de92c5462fadc884a5ce9d21c15e8a85198f2349e.
//
// Solidity: event Approved(address indexed user)
func (_ApprovalService *ApprovalServiceFilterer) WatchApproved(opts *bind.WatchOpts, sink chan<- *ApprovalServiceApproved, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ApprovalService.contract.WatchLogs(opts, "Approved", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApprovalServiceApproved)
				if err := _ApprovalService.contract.UnpackLog(event, "Approved", log); err != nil {
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

// ParseApproved is a log parse operation binding the contract event 0x5d91bd0cecc45fef102af61de92c5462fadc884a5ce9d21c15e8a85198f2349e.
//
// Solidity: event Approved(address indexed user)
func (_ApprovalService *ApprovalServiceFilterer) ParseApproved(log types.Log) (*ApprovalServiceApproved, error) {
	event := new(ApprovalServiceApproved)
	if err := _ApprovalService.contract.UnpackLog(event, "Approved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApprovalServiceRevokedIterator is returned from FilterRevoked and is used to iterate over the raw logs and unpacked data for Revoked events raised by the ApprovalService contract.
type ApprovalServiceRevokedIterator struct {
	Event *ApprovalServiceRevoked // Event containing the contract specifics and raw log

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
func (it *ApprovalServiceRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApprovalServiceRevoked)
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
		it.Event = new(ApprovalServiceRevoked)
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
func (it *ApprovalServiceRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApprovalServiceRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApprovalServiceRevoked represents a Revoked event raised by the ApprovalService contract.
type ApprovalServiceRevoked struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRevoked is a free log retrieval operation binding the contract event 0xb6fa8b8bd5eab60f292eca876e3ef90722275b785309d84b1de113ce0b8c4e74.
//
// Solidity: event Revoked(address indexed user)
func (_ApprovalService *ApprovalServiceFilterer) FilterRevoked(opts *bind.FilterOpts, user []common.Address) (*ApprovalServiceRevokedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ApprovalService.contract.FilterLogs(opts, "Revoked", userRule)
	if err != nil {
		return nil, err
	}
	return &ApprovalServiceRevokedIterator{contract: _ApprovalService.contract, event: "Revoked", logs: logs, sub: sub}, nil
}

// WatchRevoked is a free log subscription operation binding the contract event 0xb6fa8b8bd5eab60f292eca876e3ef90722275b785309d84b1de113ce0b8c4e74.
//
// Solidity: event Revoked(address indexed user)
func (_ApprovalService *ApprovalServiceFilterer) WatchRevoked(opts *bind.WatchOpts, sink chan<- *ApprovalServiceRevoked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ApprovalService.contract.WatchLogs(opts, "Revoked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApprovalServiceRevoked)
				if err := _ApprovalService.contract.UnpackLog(event, "Revoked", log); err != nil {
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

// ParseRevoked is a log parse operation binding the contract event 0xb6fa8b8bd5eab60f292eca876e3ef90722275b785309d84b1de113ce0b8c4e74.
//
// Solidity: event Revoked(address indexed user)
func (_ApprovalService *ApprovalServiceFilterer) ParseRevoked(log types.Log) (*ApprovalServiceRevoked, error) {
	event := new(ApprovalServiceRevoked)
	if err := _ApprovalService.contract.UnpackLog(event, "Revoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApprovalServiceRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the ApprovalService contract.
type ApprovalServiceRoleAdminChangedIterator struct {
	Event *ApprovalServiceRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ApprovalServiceRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApprovalServiceRoleAdminChanged)
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
		it.Event = new(ApprovalServiceRoleAdminChanged)
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
func (it *ApprovalServiceRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApprovalServiceRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApprovalServiceRoleAdminChanged represents a RoleAdminChanged event raised by the ApprovalService contract.
type ApprovalServiceRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ApprovalService *ApprovalServiceFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ApprovalServiceRoleAdminChangedIterator, error) {

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

	logs, sub, err := _ApprovalService.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ApprovalServiceRoleAdminChangedIterator{contract: _ApprovalService.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ApprovalService *ApprovalServiceFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ApprovalServiceRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _ApprovalService.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApprovalServiceRoleAdminChanged)
				if err := _ApprovalService.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_ApprovalService *ApprovalServiceFilterer) ParseRoleAdminChanged(log types.Log) (*ApprovalServiceRoleAdminChanged, error) {
	event := new(ApprovalServiceRoleAdminChanged)
	if err := _ApprovalService.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApprovalServiceRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ApprovalService contract.
type ApprovalServiceRoleGrantedIterator struct {
	Event *ApprovalServiceRoleGranted // Event containing the contract specifics and raw log

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
func (it *ApprovalServiceRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApprovalServiceRoleGranted)
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
		it.Event = new(ApprovalServiceRoleGranted)
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
func (it *ApprovalServiceRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApprovalServiceRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApprovalServiceRoleGranted represents a RoleGranted event raised by the ApprovalService contract.
type ApprovalServiceRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ApprovalService *ApprovalServiceFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ApprovalServiceRoleGrantedIterator, error) {

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

	logs, sub, err := _ApprovalService.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ApprovalServiceRoleGrantedIterator{contract: _ApprovalService.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ApprovalService *ApprovalServiceFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ApprovalServiceRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ApprovalService.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApprovalServiceRoleGranted)
				if err := _ApprovalService.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_ApprovalService *ApprovalServiceFilterer) ParseRoleGranted(log types.Log) (*ApprovalServiceRoleGranted, error) {
	event := new(ApprovalServiceRoleGranted)
	if err := _ApprovalService.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApprovalServiceRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ApprovalService contract.
type ApprovalServiceRoleRevokedIterator struct {
	Event *ApprovalServiceRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ApprovalServiceRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApprovalServiceRoleRevoked)
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
		it.Event = new(ApprovalServiceRoleRevoked)
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
func (it *ApprovalServiceRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApprovalServiceRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApprovalServiceRoleRevoked represents a RoleRevoked event raised by the ApprovalService contract.
type ApprovalServiceRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ApprovalService *ApprovalServiceFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ApprovalServiceRoleRevokedIterator, error) {

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

	logs, sub, err := _ApprovalService.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ApprovalServiceRoleRevokedIterator{contract: _ApprovalService.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ApprovalService *ApprovalServiceFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ApprovalServiceRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ApprovalService.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApprovalServiceRoleRevoked)
				if err := _ApprovalService.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_ApprovalService *ApprovalServiceFilterer) ParseRoleRevoked(log types.Log) (*ApprovalServiceRoleRevoked, error) {
	event := new(ApprovalServiceRoleRevoked)
	if err := _ApprovalService.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

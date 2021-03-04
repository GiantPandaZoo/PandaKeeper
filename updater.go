// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// AggregateUpdaterABI is the input ABI used to generate the binding from.
const AggregateUpdaterABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"_owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIOptionPool\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"addPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNextUpdateTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pools\",\"outputs\":[{\"internalType\":\"contractIOptionPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIOptionPool\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"removePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// AggregateUpdater is an auto generated Go binding around an Ethereum contract.
type AggregateUpdater struct {
	AggregateUpdaterCaller     // Read-only binding to the contract
	AggregateUpdaterTransactor // Write-only binding to the contract
	AggregateUpdaterFilterer   // Log filterer for contract events
}

// AggregateUpdaterCaller is an auto generated read-only Go binding around an Ethereum contract.
type AggregateUpdaterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggregateUpdaterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AggregateUpdaterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggregateUpdaterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AggregateUpdaterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggregateUpdaterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AggregateUpdaterSession struct {
	Contract     *AggregateUpdater // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AggregateUpdaterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AggregateUpdaterCallerSession struct {
	Contract *AggregateUpdaterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// AggregateUpdaterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AggregateUpdaterTransactorSession struct {
	Contract     *AggregateUpdaterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// AggregateUpdaterRaw is an auto generated low-level Go binding around an Ethereum contract.
type AggregateUpdaterRaw struct {
	Contract *AggregateUpdater // Generic contract binding to access the raw methods on
}

// AggregateUpdaterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AggregateUpdaterCallerRaw struct {
	Contract *AggregateUpdaterCaller // Generic read-only contract binding to access the raw methods on
}

// AggregateUpdaterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AggregateUpdaterTransactorRaw struct {
	Contract *AggregateUpdaterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAggregateUpdater creates a new instance of AggregateUpdater, bound to a specific deployed contract.
func NewAggregateUpdater(address common.Address, backend bind.ContractBackend) (*AggregateUpdater, error) {
	contract, err := bindAggregateUpdater(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AggregateUpdater{AggregateUpdaterCaller: AggregateUpdaterCaller{contract: contract}, AggregateUpdaterTransactor: AggregateUpdaterTransactor{contract: contract}, AggregateUpdaterFilterer: AggregateUpdaterFilterer{contract: contract}}, nil
}

// NewAggregateUpdaterCaller creates a new read-only instance of AggregateUpdater, bound to a specific deployed contract.
func NewAggregateUpdaterCaller(address common.Address, caller bind.ContractCaller) (*AggregateUpdaterCaller, error) {
	contract, err := bindAggregateUpdater(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AggregateUpdaterCaller{contract: contract}, nil
}

// NewAggregateUpdaterTransactor creates a new write-only instance of AggregateUpdater, bound to a specific deployed contract.
func NewAggregateUpdaterTransactor(address common.Address, transactor bind.ContractTransactor) (*AggregateUpdaterTransactor, error) {
	contract, err := bindAggregateUpdater(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AggregateUpdaterTransactor{contract: contract}, nil
}

// NewAggregateUpdaterFilterer creates a new log filterer instance of AggregateUpdater, bound to a specific deployed contract.
func NewAggregateUpdaterFilterer(address common.Address, filterer bind.ContractFilterer) (*AggregateUpdaterFilterer, error) {
	contract, err := bindAggregateUpdater(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AggregateUpdaterFilterer{contract: contract}, nil
}

// bindAggregateUpdater binds a generic wrapper to an already deployed contract.
func bindAggregateUpdater(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AggregateUpdaterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AggregateUpdater *AggregateUpdaterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AggregateUpdater.Contract.AggregateUpdaterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AggregateUpdater *AggregateUpdaterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AggregateUpdater.Contract.AggregateUpdaterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AggregateUpdater *AggregateUpdaterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AggregateUpdater.Contract.AggregateUpdaterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AggregateUpdater *AggregateUpdaterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AggregateUpdater.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AggregateUpdater *AggregateUpdaterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AggregateUpdater.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AggregateUpdater *AggregateUpdaterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AggregateUpdater.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(address)
func (_AggregateUpdater *AggregateUpdaterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AggregateUpdater.contract.Call(opts, &out, "_owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(address)
func (_AggregateUpdater *AggregateUpdaterSession) Owner() (common.Address, error) {
	return _AggregateUpdater.Contract.Owner(&_AggregateUpdater.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(address)
func (_AggregateUpdater *AggregateUpdaterCallerSession) Owner() (common.Address, error) {
	return _AggregateUpdater.Contract.Owner(&_AggregateUpdater.CallOpts)
}

// GetNextUpdateTime is a free data retrieval call binding the contract method 0x7696a015.
//
// Solidity: function getNextUpdateTime() view returns(uint256)
func (_AggregateUpdater *AggregateUpdaterCaller) GetNextUpdateTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AggregateUpdater.contract.Call(opts, &out, "getNextUpdateTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextUpdateTime is a free data retrieval call binding the contract method 0x7696a015.
//
// Solidity: function getNextUpdateTime() view returns(uint256)
func (_AggregateUpdater *AggregateUpdaterSession) GetNextUpdateTime() (*big.Int, error) {
	return _AggregateUpdater.Contract.GetNextUpdateTime(&_AggregateUpdater.CallOpts)
}

// GetNextUpdateTime is a free data retrieval call binding the contract method 0x7696a015.
//
// Solidity: function getNextUpdateTime() view returns(uint256)
func (_AggregateUpdater *AggregateUpdaterCallerSession) GetNextUpdateTime() (*big.Int, error) {
	return _AggregateUpdater.Contract.GetNextUpdateTime(&_AggregateUpdater.CallOpts)
}

// Pools is a free data retrieval call binding the contract method 0xac4afa38.
//
// Solidity: function pools(uint256 ) view returns(address)
func (_AggregateUpdater *AggregateUpdaterCaller) Pools(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AggregateUpdater.contract.Call(opts, &out, "pools", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pools is a free data retrieval call binding the contract method 0xac4afa38.
//
// Solidity: function pools(uint256 ) view returns(address)
func (_AggregateUpdater *AggregateUpdaterSession) Pools(arg0 *big.Int) (common.Address, error) {
	return _AggregateUpdater.Contract.Pools(&_AggregateUpdater.CallOpts, arg0)
}

// Pools is a free data retrieval call binding the contract method 0xac4afa38.
//
// Solidity: function pools(uint256 ) view returns(address)
func (_AggregateUpdater *AggregateUpdaterCallerSession) Pools(arg0 *big.Int) (common.Address, error) {
	return _AggregateUpdater.Contract.Pools(&_AggregateUpdater.CallOpts, arg0)
}

// AddPool is a paid mutator transaction binding the contract method 0xd914cd4b.
//
// Solidity: function addPool(address pool) returns()
func (_AggregateUpdater *AggregateUpdaterTransactor) AddPool(opts *bind.TransactOpts, pool common.Address) (*types.Transaction, error) {
	return _AggregateUpdater.contract.Transact(opts, "addPool", pool)
}

// AddPool is a paid mutator transaction binding the contract method 0xd914cd4b.
//
// Solidity: function addPool(address pool) returns()
func (_AggregateUpdater *AggregateUpdaterSession) AddPool(pool common.Address) (*types.Transaction, error) {
	return _AggregateUpdater.Contract.AddPool(&_AggregateUpdater.TransactOpts, pool)
}

// AddPool is a paid mutator transaction binding the contract method 0xd914cd4b.
//
// Solidity: function addPool(address pool) returns()
func (_AggregateUpdater *AggregateUpdaterTransactorSession) AddPool(pool common.Address) (*types.Transaction, error) {
	return _AggregateUpdater.Contract.AddPool(&_AggregateUpdater.TransactOpts, pool)
}

// RemovePool is a paid mutator transaction binding the contract method 0x3b7d0946.
//
// Solidity: function removePool(address pool) returns()
func (_AggregateUpdater *AggregateUpdaterTransactor) RemovePool(opts *bind.TransactOpts, pool common.Address) (*types.Transaction, error) {
	return _AggregateUpdater.contract.Transact(opts, "removePool", pool)
}

// RemovePool is a paid mutator transaction binding the contract method 0x3b7d0946.
//
// Solidity: function removePool(address pool) returns()
func (_AggregateUpdater *AggregateUpdaterSession) RemovePool(pool common.Address) (*types.Transaction, error) {
	return _AggregateUpdater.Contract.RemovePool(&_AggregateUpdater.TransactOpts, pool)
}

// RemovePool is a paid mutator transaction binding the contract method 0x3b7d0946.
//
// Solidity: function removePool(address pool) returns()
func (_AggregateUpdater *AggregateUpdaterTransactorSession) RemovePool(pool common.Address) (*types.Transaction, error) {
	return _AggregateUpdater.Contract.RemovePool(&_AggregateUpdater.TransactOpts, pool)
}

// Update is a paid mutator transaction binding the contract method 0xa2e62045.
//
// Solidity: function update() returns()
func (_AggregateUpdater *AggregateUpdaterTransactor) Update(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AggregateUpdater.contract.Transact(opts, "update")
}

// Update is a paid mutator transaction binding the contract method 0xa2e62045.
//
// Solidity: function update() returns()
func (_AggregateUpdater *AggregateUpdaterSession) Update() (*types.Transaction, error) {
	return _AggregateUpdater.Contract.Update(&_AggregateUpdater.TransactOpts)
}

// Update is a paid mutator transaction binding the contract method 0xa2e62045.
//
// Solidity: function update() returns()
func (_AggregateUpdater *AggregateUpdaterTransactorSession) Update() (*types.Transaction, error) {
	return _AggregateUpdater.Contract.Update(&_AggregateUpdater.TransactOpts)
}

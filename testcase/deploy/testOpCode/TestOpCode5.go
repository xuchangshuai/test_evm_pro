// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package OpCpde

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

// TestOpCode5ABI is the input ABI used to generate the binding from.
const TestOpCode5ABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"callByFun\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TestOpCode5 is an auto generated Go binding around an Ethereum contract.
type TestOpCode5 struct {
	TestOpCode5Caller     // Read-only binding to the contract
	TestOpCode5Transactor // Write-only binding to the contract
	TestOpCode5Filterer   // Log filterer for contract events
}

// TestOpCode5Caller is an auto generated read-only Go binding around an Ethereum contract.
type TestOpCode5Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestOpCode5Transactor is an auto generated write-only Go binding around an Ethereum contract.
type TestOpCode5Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestOpCode5Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestOpCode5Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestOpCode5Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestOpCode5Session struct {
	Contract     *TestOpCode5      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestOpCode5CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestOpCode5CallerSession struct {
	Contract *TestOpCode5Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// TestOpCode5TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestOpCode5TransactorSession struct {
	Contract     *TestOpCode5Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TestOpCode5Raw is an auto generated low-level Go binding around an Ethereum contract.
type TestOpCode5Raw struct {
	Contract *TestOpCode5 // Generic contract binding to access the raw methods on
}

// TestOpCode5CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestOpCode5CallerRaw struct {
	Contract *TestOpCode5Caller // Generic read-only contract binding to access the raw methods on
}

// TestOpCode5TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestOpCode5TransactorRaw struct {
	Contract *TestOpCode5Transactor // Generic write-only contract binding to access the raw methods on
}

// NewTestOpCode5 creates a new instance of TestOpCode5, bound to a specific deployed contract.
func NewTestOpCode5(address common.Address, backend bind.ContractBackend) (*TestOpCode5, error) {
	contract, err := bindTestOpCode5(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestOpCode5{TestOpCode5Caller: TestOpCode5Caller{contract: contract}, TestOpCode5Transactor: TestOpCode5Transactor{contract: contract}, TestOpCode5Filterer: TestOpCode5Filterer{contract: contract}}, nil
}

// NewTestOpCode5Caller creates a new read-only instance of TestOpCode5, bound to a specific deployed contract.
func NewTestOpCode5Caller(address common.Address, caller bind.ContractCaller) (*TestOpCode5Caller, error) {
	contract, err := bindTestOpCode5(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestOpCode5Caller{contract: contract}, nil
}

// NewTestOpCode5Transactor creates a new write-only instance of TestOpCode5, bound to a specific deployed contract.
func NewTestOpCode5Transactor(address common.Address, transactor bind.ContractTransactor) (*TestOpCode5Transactor, error) {
	contract, err := bindTestOpCode5(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestOpCode5Transactor{contract: contract}, nil
}

// NewTestOpCode5Filterer creates a new log filterer instance of TestOpCode5, bound to a specific deployed contract.
func NewTestOpCode5Filterer(address common.Address, filterer bind.ContractFilterer) (*TestOpCode5Filterer, error) {
	contract, err := bindTestOpCode5(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestOpCode5Filterer{contract: contract}, nil
}

// bindTestOpCode5 binds a generic wrapper to an already deployed contract.
func bindTestOpCode5(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestOpCode5ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestOpCode5 *TestOpCode5Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestOpCode5.Contract.TestOpCode5Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestOpCode5 *TestOpCode5Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestOpCode5.Contract.TestOpCode5Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestOpCode5 *TestOpCode5Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestOpCode5.Contract.TestOpCode5Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestOpCode5 *TestOpCode5CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestOpCode5.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestOpCode5 *TestOpCode5TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestOpCode5.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestOpCode5 *TestOpCode5TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestOpCode5.Contract.contract.Transact(opts, method, params...)
}

// CallByFun is a paid mutator transaction binding the contract method 0xd11f17ad.
//
// Solidity: function callByFun(address addr) returns(bool, bytes)
func (_TestOpCode5 *TestOpCode5Transactor) CallByFun(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _TestOpCode5.contract.Transact(opts, "callByFun", addr)
}

// CallByFun is a paid mutator transaction binding the contract method 0xd11f17ad.
//
// Solidity: function callByFun(address addr) returns(bool, bytes)
func (_TestOpCode5 *TestOpCode5Session) CallByFun(addr common.Address) (*types.Transaction, error) {
	return _TestOpCode5.Contract.CallByFun(&_TestOpCode5.TransactOpts, addr)
}

// CallByFun is a paid mutator transaction binding the contract method 0xd11f17ad.
//
// Solidity: function callByFun(address addr) returns(bool, bytes)
func (_TestOpCode5 *TestOpCode5TransactorSession) CallByFun(addr common.Address) (*types.Transaction, error) {
	return _TestOpCode5.Contract.CallByFun(&_TestOpCode5.TransactOpts, addr)
}

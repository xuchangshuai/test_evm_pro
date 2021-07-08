// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package TestOpCpde

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

// TestOpCode3ABI is the input ABI used to generate the binding from.
const TestOpCode3ABI = "[{\"inputs\":[],\"name\":\"dup\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"jump\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"push\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sloadAndSstore\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swap\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// TestOpCode3 is an auto generated Go binding around an Ethereum contract.
type TestOpCode3 struct {
	TestOpCode3Caller     // Read-only binding to the contract
	TestOpCode3Transactor // Write-only binding to the contract
	TestOpCode3Filterer   // Log filterer for contract events
}

// TestOpCode3Caller is an auto generated read-only Go binding around an Ethereum contract.
type TestOpCode3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestOpCode3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type TestOpCode3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestOpCode3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestOpCode3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestOpCode3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestOpCode3Session struct {
	Contract     *TestOpCode3      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestOpCode3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestOpCode3CallerSession struct {
	Contract *TestOpCode3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// TestOpCode3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestOpCode3TransactorSession struct {
	Contract     *TestOpCode3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TestOpCode3Raw is an auto generated low-level Go binding around an Ethereum contract.
type TestOpCode3Raw struct {
	Contract *TestOpCode3 // Generic contract binding to access the raw methods on
}

// TestOpCode3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestOpCode3CallerRaw struct {
	Contract *TestOpCode3Caller // Generic read-only contract binding to access the raw methods on
}

// TestOpCode3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestOpCode3TransactorRaw struct {
	Contract *TestOpCode3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewTestOpCode3 creates a new instance of TestOpCode3, bound to a specific deployed contract.
func NewTestOpCode3(address common.Address, backend bind.ContractBackend) (*TestOpCode3, error) {
	contract, err := bindTestOpCode3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestOpCode3{TestOpCode3Caller: TestOpCode3Caller{contract: contract}, TestOpCode3Transactor: TestOpCode3Transactor{contract: contract}, TestOpCode3Filterer: TestOpCode3Filterer{contract: contract}}, nil
}

// NewTestOpCode3Caller creates a new read-only instance of TestOpCode3, bound to a specific deployed contract.
func NewTestOpCode3Caller(address common.Address, caller bind.ContractCaller) (*TestOpCode3Caller, error) {
	contract, err := bindTestOpCode3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestOpCode3Caller{contract: contract}, nil
}

// NewTestOpCode3Transactor creates a new write-only instance of TestOpCode3, bound to a specific deployed contract.
func NewTestOpCode3Transactor(address common.Address, transactor bind.ContractTransactor) (*TestOpCode3Transactor, error) {
	contract, err := bindTestOpCode3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestOpCode3Transactor{contract: contract}, nil
}

// NewTestOpCode3Filterer creates a new log filterer instance of TestOpCode3, bound to a specific deployed contract.
func NewTestOpCode3Filterer(address common.Address, filterer bind.ContractFilterer) (*TestOpCode3Filterer, error) {
	contract, err := bindTestOpCode3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestOpCode3Filterer{contract: contract}, nil
}

// bindTestOpCode3 binds a generic wrapper to an already deployed contract.
func bindTestOpCode3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestOpCode3ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestOpCode3 *TestOpCode3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestOpCode3.Contract.TestOpCode3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestOpCode3 *TestOpCode3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestOpCode3.Contract.TestOpCode3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestOpCode3 *TestOpCode3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestOpCode3.Contract.TestOpCode3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestOpCode3 *TestOpCode3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestOpCode3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestOpCode3 *TestOpCode3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestOpCode3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestOpCode3 *TestOpCode3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestOpCode3.Contract.contract.Transact(opts, method, params...)
}

// Dup is a free data retrieval call binding the contract method 0x83243ff9.
//
// Solidity: function dup() pure returns()
func (_TestOpCode3 *TestOpCode3Caller) Dup(opts *bind.CallOpts) error {
	var out []interface{}
	err := _TestOpCode3.contract.Call(opts, &out, "dup")

	if err != nil {
		return err
	}

	return err

}

// Dup is a free data retrieval call binding the contract method 0x83243ff9.
//
// Solidity: function dup() pure returns()
func (_TestOpCode3 *TestOpCode3Session) Dup() error {
	return _TestOpCode3.Contract.Dup(&_TestOpCode3.CallOpts)
}

// Dup is a free data retrieval call binding the contract method 0x83243ff9.
//
// Solidity: function dup() pure returns()
func (_TestOpCode3 *TestOpCode3CallerSession) Dup() error {
	return _TestOpCode3.Contract.Dup(&_TestOpCode3.CallOpts)
}

// Jump is a free data retrieval call binding the contract method 0x8f6c696b.
//
// Solidity: function jump() pure returns()
func (_TestOpCode3 *TestOpCode3Caller) Jump(opts *bind.CallOpts) error {
	var out []interface{}
	err := _TestOpCode3.contract.Call(opts, &out, "jump")

	if err != nil {
		return err
	}

	return err

}

// Jump is a free data retrieval call binding the contract method 0x8f6c696b.
//
// Solidity: function jump() pure returns()
func (_TestOpCode3 *TestOpCode3Session) Jump() error {
	return _TestOpCode3.Contract.Jump(&_TestOpCode3.CallOpts)
}

// Jump is a free data retrieval call binding the contract method 0x8f6c696b.
//
// Solidity: function jump() pure returns()
func (_TestOpCode3 *TestOpCode3CallerSession) Jump() error {
	return _TestOpCode3.Contract.Jump(&_TestOpCode3.CallOpts)
}

// Push is a free data retrieval call binding the contract method 0x8035f0ce.
//
// Solidity: function push() pure returns()
func (_TestOpCode3 *TestOpCode3Caller) Push(opts *bind.CallOpts) error {
	var out []interface{}
	err := _TestOpCode3.contract.Call(opts, &out, "push")

	if err != nil {
		return err
	}

	return err

}

// Push is a free data retrieval call binding the contract method 0x8035f0ce.
//
// Solidity: function push() pure returns()
func (_TestOpCode3 *TestOpCode3Session) Push() error {
	return _TestOpCode3.Contract.Push(&_TestOpCode3.CallOpts)
}

// Push is a free data retrieval call binding the contract method 0x8035f0ce.
//
// Solidity: function push() pure returns()
func (_TestOpCode3 *TestOpCode3CallerSession) Push() error {
	return _TestOpCode3.Contract.Push(&_TestOpCode3.CallOpts)
}

// SloadAndSstore is a free data retrieval call binding the contract method 0xce02ded6.
//
// Solidity: function sloadAndSstore() view returns()
func (_TestOpCode3 *TestOpCode3Caller) SloadAndSstore(opts *bind.CallOpts) error {
	var out []interface{}
	err := _TestOpCode3.contract.Call(opts, &out, "sloadAndSstore")

	if err != nil {
		return err
	}

	return err

}

// SloadAndSstore is a free data retrieval call binding the contract method 0xce02ded6.
//
// Solidity: function sloadAndSstore() view returns()
func (_TestOpCode3 *TestOpCode3Session) SloadAndSstore() error {
	return _TestOpCode3.Contract.SloadAndSstore(&_TestOpCode3.CallOpts)
}

// SloadAndSstore is a free data retrieval call binding the contract method 0xce02ded6.
//
// Solidity: function sloadAndSstore() view returns()
func (_TestOpCode3 *TestOpCode3CallerSession) SloadAndSstore() error {
	return _TestOpCode3.Contract.SloadAndSstore(&_TestOpCode3.CallOpts)
}

// Swap is a free data retrieval call binding the contract method 0x8119c065.
//
// Solidity: function swap() pure returns()
func (_TestOpCode3 *TestOpCode3Caller) Swap(opts *bind.CallOpts) error {
	var out []interface{}
	err := _TestOpCode3.contract.Call(opts, &out, "swap")

	if err != nil {
		return err
	}

	return err

}

// Swap is a free data retrieval call binding the contract method 0x8119c065.
//
// Solidity: function swap() pure returns()
func (_TestOpCode3 *TestOpCode3Session) Swap() error {
	return _TestOpCode3.Contract.Swap(&_TestOpCode3.CallOpts)
}

// Swap is a free data retrieval call binding the contract method 0x8119c065.
//
// Solidity: function swap() pure returns()
func (_TestOpCode3 *TestOpCode3CallerSession) Swap() error {
	return _TestOpCode3.Contract.Swap(&_TestOpCode3.CallOpts)
}

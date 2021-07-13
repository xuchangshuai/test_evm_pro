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

// TestAddModAndMulModABI is the input ABI used to generate the binding from.
const TestAddModAndMulModABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"c\",\"type\":\"uint256\"}],\"name\":\"addmod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"c\",\"type\":\"uint256\"}],\"name\":\"mulmod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// TestAddModAndMulMod is an auto generated Go binding around an Ethereum contract.
type TestAddModAndMulMod struct {
	TestAddModAndMulModCaller     // Read-only binding to the contract
	TestAddModAndMulModTransactor // Write-only binding to the contract
	TestAddModAndMulModFilterer   // Log filterer for contract events
}

// TestAddModAndMulModCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestAddModAndMulModCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestAddModAndMulModTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestAddModAndMulModTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestAddModAndMulModFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestAddModAndMulModFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestAddModAndMulModSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestAddModAndMulModSession struct {
	Contract     *TestAddModAndMulMod // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TestAddModAndMulModCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestAddModAndMulModCallerSession struct {
	Contract *TestAddModAndMulModCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// TestAddModAndMulModTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestAddModAndMulModTransactorSession struct {
	Contract     *TestAddModAndMulModTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// TestAddModAndMulModRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestAddModAndMulModRaw struct {
	Contract *TestAddModAndMulMod // Generic contract binding to access the raw methods on
}

// TestAddModAndMulModCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestAddModAndMulModCallerRaw struct {
	Contract *TestAddModAndMulModCaller // Generic read-only contract binding to access the raw methods on
}

// TestAddModAndMulModTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestAddModAndMulModTransactorRaw struct {
	Contract *TestAddModAndMulModTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestAddModAndMulMod creates a new instance of TestAddModAndMulMod, bound to a specific deployed contract.
func NewTestAddModAndMulMod(address common.Address, backend bind.ContractBackend) (*TestAddModAndMulMod, error) {
	contract, err := bindTestAddModAndMulMod(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestAddModAndMulMod{TestAddModAndMulModCaller: TestAddModAndMulModCaller{contract: contract}, TestAddModAndMulModTransactor: TestAddModAndMulModTransactor{contract: contract}, TestAddModAndMulModFilterer: TestAddModAndMulModFilterer{contract: contract}}, nil
}

// NewTestAddModAndMulModCaller creates a new read-only instance of TestAddModAndMulMod, bound to a specific deployed contract.
func NewTestAddModAndMulModCaller(address common.Address, caller bind.ContractCaller) (*TestAddModAndMulModCaller, error) {
	contract, err := bindTestAddModAndMulMod(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestAddModAndMulModCaller{contract: contract}, nil
}

// NewTestAddModAndMulModTransactor creates a new write-only instance of TestAddModAndMulMod, bound to a specific deployed contract.
func NewTestAddModAndMulModTransactor(address common.Address, transactor bind.ContractTransactor) (*TestAddModAndMulModTransactor, error) {
	contract, err := bindTestAddModAndMulMod(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestAddModAndMulModTransactor{contract: contract}, nil
}

// NewTestAddModAndMulModFilterer creates a new log filterer instance of TestAddModAndMulMod, bound to a specific deployed contract.
func NewTestAddModAndMulModFilterer(address common.Address, filterer bind.ContractFilterer) (*TestAddModAndMulModFilterer, error) {
	contract, err := bindTestAddModAndMulMod(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestAddModAndMulModFilterer{contract: contract}, nil
}

// bindTestAddModAndMulMod binds a generic wrapper to an already deployed contract.
func bindTestAddModAndMulMod(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestAddModAndMulModABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestAddModAndMulMod *TestAddModAndMulModRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestAddModAndMulMod.Contract.TestAddModAndMulModCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestAddModAndMulMod *TestAddModAndMulModRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestAddModAndMulMod.Contract.TestAddModAndMulModTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestAddModAndMulMod *TestAddModAndMulModRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestAddModAndMulMod.Contract.TestAddModAndMulModTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestAddModAndMulMod *TestAddModAndMulModCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestAddModAndMulMod.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestAddModAndMulMod *TestAddModAndMulModTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestAddModAndMulMod.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestAddModAndMulMod *TestAddModAndMulModTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestAddModAndMulMod.Contract.contract.Transact(opts, method, params...)
}

// Addmod is a free data retrieval call binding the contract method 0x9796df37.
//
// Solidity: function addmod(uint256 a, uint256 b, uint256 c) pure returns(uint256)
func (_TestAddModAndMulMod *TestAddModAndMulModCaller) Addmod(opts *bind.CallOpts, a *big.Int, b *big.Int, c *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TestAddModAndMulMod.contract.Call(opts, &out, "addmod", a, b, c)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Addmod is a free data retrieval call binding the contract method 0x9796df37.
//
// Solidity: function addmod(uint256 a, uint256 b, uint256 c) pure returns(uint256)
func (_TestAddModAndMulMod *TestAddModAndMulModSession) Addmod(a *big.Int, b *big.Int, c *big.Int) (*big.Int, error) {
	return _TestAddModAndMulMod.Contract.Addmod(&_TestAddModAndMulMod.CallOpts, a, b, c)
}

// Addmod is a free data retrieval call binding the contract method 0x9796df37.
//
// Solidity: function addmod(uint256 a, uint256 b, uint256 c) pure returns(uint256)
func (_TestAddModAndMulMod *TestAddModAndMulModCallerSession) Addmod(a *big.Int, b *big.Int, c *big.Int) (*big.Int, error) {
	return _TestAddModAndMulMod.Contract.Addmod(&_TestAddModAndMulMod.CallOpts, a, b, c)
}

// Mulmod is a free data retrieval call binding the contract method 0x4902ddae.
//
// Solidity: function mulmod(uint256 a, uint256 b, uint256 c) pure returns(uint256)
func (_TestAddModAndMulMod *TestAddModAndMulModCaller) Mulmod(opts *bind.CallOpts, a *big.Int, b *big.Int, c *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TestAddModAndMulMod.contract.Call(opts, &out, "mulmod", a, b, c)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Mulmod is a free data retrieval call binding the contract method 0x4902ddae.
//
// Solidity: function mulmod(uint256 a, uint256 b, uint256 c) pure returns(uint256)
func (_TestAddModAndMulMod *TestAddModAndMulModSession) Mulmod(a *big.Int, b *big.Int, c *big.Int) (*big.Int, error) {
	return _TestAddModAndMulMod.Contract.Mulmod(&_TestAddModAndMulMod.CallOpts, a, b, c)
}

// Mulmod is a free data retrieval call binding the contract method 0x4902ddae.
//
// Solidity: function mulmod(uint256 a, uint256 b, uint256 c) pure returns(uint256)
func (_TestAddModAndMulMod *TestAddModAndMulModCallerSession) Mulmod(a *big.Int, b *big.Int, c *big.Int) (*big.Int, error) {
	return _TestAddModAndMulMod.Contract.Mulmod(&_TestAddModAndMulMod.CallOpts, a, b, c)
}

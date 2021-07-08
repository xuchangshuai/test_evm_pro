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

// TestOpCode4ABI is the input ABI used to generate the binding from.
const TestOpCode4ABI = "[{\"inputs\":[],\"name\":\"age\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"increaseAge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TestOpCode4 is an auto generated Go binding around an Ethereum contract.
type TestOpCode4 struct {
	TestOpCode4Caller     // Read-only binding to the contract
	TestOpCode4Transactor // Write-only binding to the contract
	TestOpCode4Filterer   // Log filterer for contract events
}

// TestOpCode4Caller is an auto generated read-only Go binding around an Ethereum contract.
type TestOpCode4Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestOpCode4Transactor is an auto generated write-only Go binding around an Ethereum contract.
type TestOpCode4Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestOpCode4Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestOpCode4Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestOpCode4Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestOpCode4Session struct {
	Contract     *TestOpCode4      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestOpCode4CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestOpCode4CallerSession struct {
	Contract *TestOpCode4Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// TestOpCode4TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestOpCode4TransactorSession struct {
	Contract     *TestOpCode4Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TestOpCode4Raw is an auto generated low-level Go binding around an Ethereum contract.
type TestOpCode4Raw struct {
	Contract *TestOpCode4 // Generic contract binding to access the raw methods on
}

// TestOpCode4CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestOpCode4CallerRaw struct {
	Contract *TestOpCode4Caller // Generic read-only contract binding to access the raw methods on
}

// TestOpCode4TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestOpCode4TransactorRaw struct {
	Contract *TestOpCode4Transactor // Generic write-only contract binding to access the raw methods on
}

// NewTestOpCode4 creates a new instance of TestOpCode4, bound to a specific deployed contract.
func NewTestOpCode4(address common.Address, backend bind.ContractBackend) (*TestOpCode4, error) {
	contract, err := bindTestOpCode4(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestOpCode4{TestOpCode4Caller: TestOpCode4Caller{contract: contract}, TestOpCode4Transactor: TestOpCode4Transactor{contract: contract}, TestOpCode4Filterer: TestOpCode4Filterer{contract: contract}}, nil
}

// NewTestOpCode4Caller creates a new read-only instance of TestOpCode4, bound to a specific deployed contract.
func NewTestOpCode4Caller(address common.Address, caller bind.ContractCaller) (*TestOpCode4Caller, error) {
	contract, err := bindTestOpCode4(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestOpCode4Caller{contract: contract}, nil
}

// NewTestOpCode4Transactor creates a new write-only instance of TestOpCode4, bound to a specific deployed contract.
func NewTestOpCode4Transactor(address common.Address, transactor bind.ContractTransactor) (*TestOpCode4Transactor, error) {
	contract, err := bindTestOpCode4(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestOpCode4Transactor{contract: contract}, nil
}

// NewTestOpCode4Filterer creates a new log filterer instance of TestOpCode4, bound to a specific deployed contract.
func NewTestOpCode4Filterer(address common.Address, filterer bind.ContractFilterer) (*TestOpCode4Filterer, error) {
	contract, err := bindTestOpCode4(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestOpCode4Filterer{contract: contract}, nil
}

// bindTestOpCode4 binds a generic wrapper to an already deployed contract.
func bindTestOpCode4(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestOpCode4ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestOpCode4 *TestOpCode4Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestOpCode4.Contract.TestOpCode4Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestOpCode4 *TestOpCode4Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestOpCode4.Contract.TestOpCode4Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestOpCode4 *TestOpCode4Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestOpCode4.Contract.TestOpCode4Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestOpCode4 *TestOpCode4CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestOpCode4.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestOpCode4 *TestOpCode4TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestOpCode4.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestOpCode4 *TestOpCode4TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestOpCode4.Contract.contract.Transact(opts, method, params...)
}

// Age is a free data retrieval call binding the contract method 0x262a9dff.
//
// Solidity: function age() view returns(uint256)
func (_TestOpCode4 *TestOpCode4Caller) Age(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestOpCode4.contract.Call(opts, &out, "age")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Age is a free data retrieval call binding the contract method 0x262a9dff.
//
// Solidity: function age() view returns(uint256)
func (_TestOpCode4 *TestOpCode4Session) Age() (*big.Int, error) {
	return _TestOpCode4.Contract.Age(&_TestOpCode4.CallOpts)
}

// Age is a free data retrieval call binding the contract method 0x262a9dff.
//
// Solidity: function age() view returns(uint256)
func (_TestOpCode4 *TestOpCode4CallerSession) Age() (*big.Int, error) {
	return _TestOpCode4.Contract.Age(&_TestOpCode4.CallOpts)
}

// GetAge is a paid mutator transaction binding the contract method 0x967e6e65.
//
// Solidity: function getAge() returns(uint256)
func (_TestOpCode4 *TestOpCode4Transactor) GetAge(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestOpCode4.contract.Transact(opts, "getAge")
}

// GetAge is a paid mutator transaction binding the contract method 0x967e6e65.
//
// Solidity: function getAge() returns(uint256)
func (_TestOpCode4 *TestOpCode4Session) GetAge() (*types.Transaction, error) {
	return _TestOpCode4.Contract.GetAge(&_TestOpCode4.TransactOpts)
}

// GetAge is a paid mutator transaction binding the contract method 0x967e6e65.
//
// Solidity: function getAge() returns(uint256)
func (_TestOpCode4 *TestOpCode4TransactorSession) GetAge() (*types.Transaction, error) {
	return _TestOpCode4.Contract.GetAge(&_TestOpCode4.TransactOpts)
}

// IncreaseAge is a paid mutator transaction binding the contract method 0x42855eb3.
//
// Solidity: function increaseAge(uint256 num) returns(uint256)
func (_TestOpCode4 *TestOpCode4Transactor) IncreaseAge(opts *bind.TransactOpts, num *big.Int) (*types.Transaction, error) {
	return _TestOpCode4.contract.Transact(opts, "increaseAge", num)
}

// IncreaseAge is a paid mutator transaction binding the contract method 0x42855eb3.
//
// Solidity: function increaseAge(uint256 num) returns(uint256)
func (_TestOpCode4 *TestOpCode4Session) IncreaseAge(num *big.Int) (*types.Transaction, error) {
	return _TestOpCode4.Contract.IncreaseAge(&_TestOpCode4.TransactOpts, num)
}

// IncreaseAge is a paid mutator transaction binding the contract method 0x42855eb3.
//
// Solidity: function increaseAge(uint256 num) returns(uint256)
func (_TestOpCode4 *TestOpCode4TransactorSession) IncreaseAge(num *big.Int) (*types.Transaction, error) {
	return _TestOpCode4.Contract.IncreaseAge(&_TestOpCode4.TransactOpts, num)
}

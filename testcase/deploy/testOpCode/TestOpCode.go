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

// TestOpCodeABI is the input ABI used to generate the binding from.
const TestOpCodeABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"add\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"a\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"b\",\"type\":\"int256\"}],\"name\":\"and\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"div\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"a\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"b\",\"type\":\"int256\"}],\"name\":\"eq\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"exp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"gt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"a\",\"type\":\"int256\"}],\"name\":\"iszero\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"lt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"mod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"mul\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"a\",\"type\":\"int256\"}],\"name\":\"not\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"a\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"b\",\"type\":\"int256\"}],\"name\":\"or\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"a\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"b\",\"type\":\"int256\"}],\"name\":\"sdiv\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"a\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"b\",\"type\":\"int256\"}],\"name\":\"sgt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"a\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"b\",\"type\":\"string\"}],\"name\":\"sha\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"shl\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"shr\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"a\",\"type\":\"int256\"}],\"name\":\"shr\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"a\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"b\",\"type\":\"int256\"}],\"name\":\"slt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"a\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"b\",\"type\":\"int256\"}],\"name\":\"smod\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"sub\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"a\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"b\",\"type\":\"int256\"}],\"name\":\"xor\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// TestOpCode is an auto generated Go binding around an Ethereum contract.
type TestOpCode struct {
	TestOpCodeCaller     // Read-only binding to the contract
	TestOpCodeTransactor // Write-only binding to the contract
	TestOpCodeFilterer   // Log filterer for contract events
}

// TestOpCodeCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestOpCodeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestOpCodeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestOpCodeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestOpCodeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestOpCodeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestOpCodeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestOpCodeSession struct {
	Contract     *TestOpCode       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestOpCodeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestOpCodeCallerSession struct {
	Contract *TestOpCodeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// TestOpCodeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestOpCodeTransactorSession struct {
	Contract     *TestOpCodeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TestOpCodeRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestOpCodeRaw struct {
	Contract *TestOpCode // Generic contract binding to access the raw methods on
}

// TestOpCodeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestOpCodeCallerRaw struct {
	Contract *TestOpCodeCaller // Generic read-only contract binding to access the raw methods on
}

// TestOpCodeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestOpCodeTransactorRaw struct {
	Contract *TestOpCodeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestOpCode creates a new instance of TestOpCode, bound to a specific deployed contract.
func NewTestOpCode(address common.Address, backend bind.ContractBackend) (*TestOpCode, error) {
	contract, err := bindTestOpCode(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestOpCode{TestOpCodeCaller: TestOpCodeCaller{contract: contract}, TestOpCodeTransactor: TestOpCodeTransactor{contract: contract}, TestOpCodeFilterer: TestOpCodeFilterer{contract: contract}}, nil
}

// NewTestOpCodeCaller creates a new read-only instance of TestOpCode, bound to a specific deployed contract.
func NewTestOpCodeCaller(address common.Address, caller bind.ContractCaller) (*TestOpCodeCaller, error) {
	contract, err := bindTestOpCode(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestOpCodeCaller{contract: contract}, nil
}

// NewTestOpCodeTransactor creates a new write-only instance of TestOpCode, bound to a specific deployed contract.
func NewTestOpCodeTransactor(address common.Address, transactor bind.ContractTransactor) (*TestOpCodeTransactor, error) {
	contract, err := bindTestOpCode(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestOpCodeTransactor{contract: contract}, nil
}

// NewTestOpCodeFilterer creates a new log filterer instance of TestOpCode, bound to a specific deployed contract.
func NewTestOpCodeFilterer(address common.Address, filterer bind.ContractFilterer) (*TestOpCodeFilterer, error) {
	contract, err := bindTestOpCode(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestOpCodeFilterer{contract: contract}, nil
}

// bindTestOpCode binds a generic wrapper to an already deployed contract.
func bindTestOpCode(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestOpCodeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestOpCode *TestOpCodeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestOpCode.Contract.TestOpCodeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestOpCode *TestOpCodeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestOpCode.Contract.TestOpCodeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestOpCode *TestOpCodeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestOpCode.Contract.TestOpCodeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestOpCode *TestOpCodeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestOpCode.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestOpCode *TestOpCodeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestOpCode.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestOpCode *TestOpCodeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestOpCode.Contract.contract.Transact(opts, method, params...)
}

// Add is a free data retrieval call binding the contract method 0x771602f7.
//
// Solidity: function add(uint256 a, uint256 b) pure returns(uint256)
func (_TestOpCode *TestOpCodeCaller) Add(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TestOpCode.contract.Call(opts, &out, "add", a, b)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Add is a free data retrieval call binding the contract method 0x771602f7.
//
// Solidity: function add(uint256 a, uint256 b) pure returns(uint256)
func (_TestOpCode *TestOpCodeSession) Add(a *big.Int, b *big.Int) (*big.Int, error) {
	return _TestOpCode.Contract.Add(&_TestOpCode.CallOpts, a, b)
}

// Add is a free data retrieval call binding the contract method 0x771602f7.
//
// Solidity: function add(uint256 a, uint256 b) pure returns(uint256)
func (_TestOpCode *TestOpCodeCallerSession) Add(a *big.Int, b *big.Int) (*big.Int, error) {
	return _TestOpCode.Contract.Add(&_TestOpCode.CallOpts, a, b)
}

// And is a free data retrieval call binding the contract method 0x726e2bdf.
//
// Solidity: function and(int256 a, int256 b) pure returns(int256)
func (_TestOpCode *TestOpCodeCaller) And(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TestOpCode.contract.Call(opts, &out, "and", a, b)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// And is a free data retrieval call binding the contract method 0x726e2bdf.
//
// Solidity: function and(int256 a, int256 b) pure returns(int256)
func (_TestOpCode *TestOpCodeSession) And(a *big.Int, b *big.Int) (*big.Int, error) {
	return _TestOpCode.Contract.And(&_TestOpCode.CallOpts, a, b)
}

// And is a free data retrieval call binding the contract method 0x726e2bdf.
//
// Solidity: function and(int256 a, int256 b) pure returns(int256)
func (_TestOpCode *TestOpCodeCallerSession) And(a *big.Int, b *big.Int) (*big.Int, error) {
	return _TestOpCode.Contract.And(&_TestOpCode.CallOpts, a, b)
}

// Div is a free data retrieval call binding the contract method 0xa391c15b.
//
// Solidity: function div(uint256 a, uint256 b) pure returns(uint256)
func (_TestOpCode *TestOpCodeCaller) Div(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TestOpCode.contract.Call(opts, &out, "div", a, b)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Div is a free data retrieval call binding the contract method 0xa391c15b.
//
// Solidity: function div(uint256 a, uint256 b) pure returns(uint256)
func (_TestOpCode *TestOpCodeSession) Div(a *big.Int, b *big.Int) (*big.Int, error) {
	return _TestOpCode.Contract.Div(&_TestOpCode.CallOpts, a, b)
}

// Div is a free data retrieval call binding the contract method 0xa391c15b.
//
// Solidity: function div(uint256 a, uint256 b) pure returns(uint256)
func (_TestOpCode *TestOpCodeCallerSession) Div(a *big.Int, b *big.Int) (*big.Int, error) {
	return _TestOpCode.Contract.Div(&_TestOpCode.CallOpts, a, b)
}

// Eq is a free data retrieval call binding the contract method 0xaa2c45bf.
//
// Solidity: function eq(int256 a, int256 b) pure returns(bool)
func (_TestOpCode *TestOpCodeCaller) Eq(opts *bind.CallOpts, a *big.Int, b *big.Int) (bool, error) {
	var out []interface{}
	err := _TestOpCode.contract.Call(opts, &out, "eq", a, b)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Eq is a free data retrieval call binding the contract method 0xaa2c45bf.
//
// Solidity: function eq(int256 a, int256 b) pure returns(bool)
func (_TestOpCode *TestOpCodeSession) Eq(a *big.Int, b *big.Int) (bool, error) {
	return _TestOpCode.Contract.Eq(&_TestOpCode.CallOpts, a, b)
}

// Eq is a free data retrieval call binding the contract method 0xaa2c45bf.
//
// Solidity: function eq(int256 a, int256 b) pure returns(bool)
func (_TestOpCode *TestOpCodeCallerSession) Eq(a *big.Int, b *big.Int) (bool, error) {
	return _TestOpCode.Contract.Eq(&_TestOpCode.CallOpts, a, b)
}

// Exp is a free data retrieval call binding the contract method 0xf5f565f8.
//
// Solidity: function exp(uint256 a, uint256 b) pure returns(uint256)
func (_TestOpCode *TestOpCodeCaller) Exp(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TestOpCode.contract.Call(opts, &out, "exp", a, b)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Exp is a free data retrieval call binding the contract method 0xf5f565f8.
//
// Solidity: function exp(uint256 a, uint256 b) pure returns(uint256)
func (_TestOpCode *TestOpCodeSession) Exp(a *big.Int, b *big.Int) (*big.Int, error) {
	return _TestOpCode.Contract.Exp(&_TestOpCode.CallOpts, a, b)
}

// Exp is a free data retrieval call binding the contract method 0xf5f565f8.
//
// Solidity: function exp(uint256 a, uint256 b) pure returns(uint256)
func (_TestOpCode *TestOpCodeCallerSession) Exp(a *big.Int, b *big.Int) (*big.Int, error) {
	return _TestOpCode.Contract.Exp(&_TestOpCode.CallOpts, a, b)
}

// Gt is a free data retrieval call binding the contract method 0x21e5749b.
//
// Solidity: function gt(uint256 a, uint256 b) pure returns(bool)
func (_TestOpCode *TestOpCodeCaller) Gt(opts *bind.CallOpts, a *big.Int, b *big.Int) (bool, error) {
	var out []interface{}
	err := _TestOpCode.contract.Call(opts, &out, "gt", a, b)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Gt is a free data retrieval call binding the contract method 0x21e5749b.
//
// Solidity: function gt(uint256 a, uint256 b) pure returns(bool)
func (_TestOpCode *TestOpCodeSession) Gt(a *big.Int, b *big.Int) (bool, error) {
	return _TestOpCode.Contract.Gt(&_TestOpCode.CallOpts, a, b)
}

// Gt is a free data retrieval call binding the contract method 0x21e5749b.
//
// Solidity: function gt(uint256 a, uint256 b) pure returns(bool)
func (_TestOpCode *TestOpCodeCallerSession) Gt(a *big.Int, b *big.Int) (bool, error) {
	return _TestOpCode.Contract.Gt(&_TestOpCode.CallOpts, a, b)
}

// Iszero is a free data retrieval call binding the contract method 0x319d5414.
//
// Solidity: function iszero(int256 a) pure returns(bool)
func (_TestOpCode *TestOpCodeCaller) Iszero(opts *bind.CallOpts, a *big.Int) (bool, error) {
	var out []interface{}
	err := _TestOpCode.contract.Call(opts, &out, "iszero", a)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Iszero is a free data retrieval call binding the contract method 0x319d5414.
//
// Solidity: function iszero(int256 a) pure returns(bool)
func (_TestOpCode *TestOpCodeSession) Iszero(a *big.Int) (bool, error) {
	return _TestOpCode.Contract.Iszero(&_TestOpCode.CallOpts, a)
}

// Iszero is a free data retrieval call binding the contract method 0x319d5414.
//
// Solidity: function iszero(int256 a) pure returns(bool)
func (_TestOpCode *TestOpCodeCallerSession) Iszero(a *big.Int) (bool, error) {
	return _TestOpCode.Contract.Iszero(&_TestOpCode.CallOpts, a)
}

// Lt is a free data retrieval call binding the contract method 0x118fc88c.
//
// Solidity: function lt(uint256 a, uint256 b) pure returns(bool)
func (_TestOpCode *TestOpCodeCaller) Lt(opts *bind.CallOpts, a *big.Int, b *big.Int) (bool, error) {
	var out []interface{}
	err := _TestOpCode.contract.Call(opts, &out, "lt", a, b)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Lt is a free data retrieval call binding the contract method 0x118fc88c.
//
// Solidity: function lt(uint256 a, uint256 b) pure returns(bool)
func (_TestOpCode *TestOpCodeSession) Lt(a *big.Int, b *big.Int) (bool, error) {
	return _TestOpCode.Contract.Lt(&_TestOpCode.CallOpts, a, b)
}

// Lt is a free data retrieval call binding the contract method 0x118fc88c.
//
// Solidity: function lt(uint256 a, uint256 b) pure returns(bool)
func (_TestOpCode *TestOpCodeCallerSession) Lt(a *big.Int, b *big.Int) (bool, error) {
	return _TestOpCode.Contract.Lt(&_TestOpCode.CallOpts, a, b)
}

// Mod is a free data retrieval call binding the contract method 0xf43f523a.
//
// Solidity: function mod(uint256 a, uint256 b) pure returns(uint256)
func (_TestOpCode *TestOpCodeCaller) Mod(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TestOpCode.contract.Call(opts, &out, "mod", a, b)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Mod is a free data retrieval call binding the contract method 0xf43f523a.
//
// Solidity: function mod(uint256 a, uint256 b) pure returns(uint256)
func (_TestOpCode *TestOpCodeSession) Mod(a *big.Int, b *big.Int) (*big.Int, error) {
	return _TestOpCode.Contract.Mod(&_TestOpCode.CallOpts, a, b)
}

// Mod is a free data retrieval call binding the contract method 0xf43f523a.
//
// Solidity: function mod(uint256 a, uint256 b) pure returns(uint256)
func (_TestOpCode *TestOpCodeCallerSession) Mod(a *big.Int, b *big.Int) (*big.Int, error) {
	return _TestOpCode.Contract.Mod(&_TestOpCode.CallOpts, a, b)
}

// Mul is a free data retrieval call binding the contract method 0xc8a4ac9c.
//
// Solidity: function mul(uint256 a, uint256 b) pure returns(uint256)
func (_TestOpCode *TestOpCodeCaller) Mul(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TestOpCode.contract.Call(opts, &out, "mul", a, b)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Mul is a free data retrieval call binding the contract method 0xc8a4ac9c.
//
// Solidity: function mul(uint256 a, uint256 b) pure returns(uint256)
func (_TestOpCode *TestOpCodeSession) Mul(a *big.Int, b *big.Int) (*big.Int, error) {
	return _TestOpCode.Contract.Mul(&_TestOpCode.CallOpts, a, b)
}

// Mul is a free data retrieval call binding the contract method 0xc8a4ac9c.
//
// Solidity: function mul(uint256 a, uint256 b) pure returns(uint256)
func (_TestOpCode *TestOpCodeCallerSession) Mul(a *big.Int, b *big.Int) (*big.Int, error) {
	return _TestOpCode.Contract.Mul(&_TestOpCode.CallOpts, a, b)
}

// Not is a free data retrieval call binding the contract method 0xa4c82e02.
//
// Solidity: function not(int256 a) pure returns(int256)
func (_TestOpCode *TestOpCodeCaller) Not(opts *bind.CallOpts, a *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TestOpCode.contract.Call(opts, &out, "not", a)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Not is a free data retrieval call binding the contract method 0xa4c82e02.
//
// Solidity: function not(int256 a) pure returns(int256)
func (_TestOpCode *TestOpCodeSession) Not(a *big.Int) (*big.Int, error) {
	return _TestOpCode.Contract.Not(&_TestOpCode.CallOpts, a)
}

// Not is a free data retrieval call binding the contract method 0xa4c82e02.
//
// Solidity: function not(int256 a) pure returns(int256)
func (_TestOpCode *TestOpCodeCallerSession) Not(a *big.Int) (*big.Int, error) {
	return _TestOpCode.Contract.Not(&_TestOpCode.CallOpts, a)
}

// Or is a free data retrieval call binding the contract method 0xa22fcaa6.
//
// Solidity: function or(int256 a, int256 b) pure returns(int256)
func (_TestOpCode *TestOpCodeCaller) Or(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TestOpCode.contract.Call(opts, &out, "or", a, b)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Or is a free data retrieval call binding the contract method 0xa22fcaa6.
//
// Solidity: function or(int256 a, int256 b) pure returns(int256)
func (_TestOpCode *TestOpCodeSession) Or(a *big.Int, b *big.Int) (*big.Int, error) {
	return _TestOpCode.Contract.Or(&_TestOpCode.CallOpts, a, b)
}

// Or is a free data retrieval call binding the contract method 0xa22fcaa6.
//
// Solidity: function or(int256 a, int256 b) pure returns(int256)
func (_TestOpCode *TestOpCodeCallerSession) Or(a *big.Int, b *big.Int) (*big.Int, error) {
	return _TestOpCode.Contract.Or(&_TestOpCode.CallOpts, a, b)
}

// Sdiv is a free data retrieval call binding the contract method 0x397b3a49.
//
// Solidity: function sdiv(int256 a, int256 b) pure returns(int256)
func (_TestOpCode *TestOpCodeCaller) Sdiv(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TestOpCode.contract.Call(opts, &out, "sdiv", a, b)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Sdiv is a free data retrieval call binding the contract method 0x397b3a49.
//
// Solidity: function sdiv(int256 a, int256 b) pure returns(int256)
func (_TestOpCode *TestOpCodeSession) Sdiv(a *big.Int, b *big.Int) (*big.Int, error) {
	return _TestOpCode.Contract.Sdiv(&_TestOpCode.CallOpts, a, b)
}

// Sdiv is a free data retrieval call binding the contract method 0x397b3a49.
//
// Solidity: function sdiv(int256 a, int256 b) pure returns(int256)
func (_TestOpCode *TestOpCodeCallerSession) Sdiv(a *big.Int, b *big.Int) (*big.Int, error) {
	return _TestOpCode.Contract.Sdiv(&_TestOpCode.CallOpts, a, b)
}

// Sgt is a free data retrieval call binding the contract method 0x2912581c.
//
// Solidity: function sgt(int256 a, int256 b) pure returns(bool)
func (_TestOpCode *TestOpCodeCaller) Sgt(opts *bind.CallOpts, a *big.Int, b *big.Int) (bool, error) {
	var out []interface{}
	err := _TestOpCode.contract.Call(opts, &out, "sgt", a, b)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Sgt is a free data retrieval call binding the contract method 0x2912581c.
//
// Solidity: function sgt(int256 a, int256 b) pure returns(bool)
func (_TestOpCode *TestOpCodeSession) Sgt(a *big.Int, b *big.Int) (bool, error) {
	return _TestOpCode.Contract.Sgt(&_TestOpCode.CallOpts, a, b)
}

// Sgt is a free data retrieval call binding the contract method 0x2912581c.
//
// Solidity: function sgt(int256 a, int256 b) pure returns(bool)
func (_TestOpCode *TestOpCodeCallerSession) Sgt(a *big.Int, b *big.Int) (bool, error) {
	return _TestOpCode.Contract.Sgt(&_TestOpCode.CallOpts, a, b)
}

// Sha is a free data retrieval call binding the contract method 0xf82f396f.
//
// Solidity: function sha(string a, string b) pure returns(bool)
func (_TestOpCode *TestOpCodeCaller) Sha(opts *bind.CallOpts, a string, b string) (bool, error) {
	var out []interface{}
	err := _TestOpCode.contract.Call(opts, &out, "sha", a, b)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Sha is a free data retrieval call binding the contract method 0xf82f396f.
//
// Solidity: function sha(string a, string b) pure returns(bool)
func (_TestOpCode *TestOpCodeSession) Sha(a string, b string) (bool, error) {
	return _TestOpCode.Contract.Sha(&_TestOpCode.CallOpts, a, b)
}

// Sha is a free data retrieval call binding the contract method 0xf82f396f.
//
// Solidity: function sha(string a, string b) pure returns(bool)
func (_TestOpCode *TestOpCodeCallerSession) Sha(a string, b string) (bool, error) {
	return _TestOpCode.Contract.Sha(&_TestOpCode.CallOpts, a, b)
}

// Shl is a free data retrieval call binding the contract method 0xbe9cdc52.
//
// Solidity: function shl(uint256 a) pure returns(uint256)
func (_TestOpCode *TestOpCodeCaller) Shl(opts *bind.CallOpts, a *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TestOpCode.contract.Call(opts, &out, "shl", a)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Shl is a free data retrieval call binding the contract method 0xbe9cdc52.
//
// Solidity: function shl(uint256 a) pure returns(uint256)
func (_TestOpCode *TestOpCodeSession) Shl(a *big.Int) (*big.Int, error) {
	return _TestOpCode.Contract.Shl(&_TestOpCode.CallOpts, a)
}

// Shl is a free data retrieval call binding the contract method 0xbe9cdc52.
//
// Solidity: function shl(uint256 a) pure returns(uint256)
func (_TestOpCode *TestOpCodeCallerSession) Shl(a *big.Int) (*big.Int, error) {
	return _TestOpCode.Contract.Shl(&_TestOpCode.CallOpts, a)
}

// Shr is a free data retrieval call binding the contract method 0x391e4dc1.
//
// Solidity: function shr(uint256 a) pure returns(uint256)
func (_TestOpCode *TestOpCodeCaller) Shr(opts *bind.CallOpts, a *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TestOpCode.contract.Call(opts, &out, "shr", a)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Shr is a free data retrieval call binding the contract method 0x391e4dc1.
//
// Solidity: function shr(uint256 a) pure returns(uint256)
func (_TestOpCode *TestOpCodeSession) Shr(a *big.Int) (*big.Int, error) {
	return _TestOpCode.Contract.Shr(&_TestOpCode.CallOpts, a)
}

// Shr is a free data retrieval call binding the contract method 0x391e4dc1.
//
// Solidity: function shr(uint256 a) pure returns(uint256)
func (_TestOpCode *TestOpCodeCallerSession) Shr(a *big.Int) (*big.Int, error) {
	return _TestOpCode.Contract.Shr(&_TestOpCode.CallOpts, a)
}

// Shr0 is a free data retrieval call binding the contract method 0x402e3313.
//
// Solidity: function shr(int256 a) pure returns(int256)
func (_TestOpCode *TestOpCodeCaller) Shr0(opts *bind.CallOpts, a *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TestOpCode.contract.Call(opts, &out, "shr0", a)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Shr0 is a free data retrieval call binding the contract method 0x402e3313.
//
// Solidity: function shr(int256 a) pure returns(int256)
func (_TestOpCode *TestOpCodeSession) Shr0(a *big.Int) (*big.Int, error) {
	return _TestOpCode.Contract.Shr0(&_TestOpCode.CallOpts, a)
}

// Shr0 is a free data retrieval call binding the contract method 0x402e3313.
//
// Solidity: function shr(int256 a) pure returns(int256)
func (_TestOpCode *TestOpCodeCallerSession) Shr0(a *big.Int) (*big.Int, error) {
	return _TestOpCode.Contract.Shr0(&_TestOpCode.CallOpts, a)
}

// Slt is a free data retrieval call binding the contract method 0x42a08c38.
//
// Solidity: function slt(int256 a, int256 b) pure returns(bool)
func (_TestOpCode *TestOpCodeCaller) Slt(opts *bind.CallOpts, a *big.Int, b *big.Int) (bool, error) {
	var out []interface{}
	err := _TestOpCode.contract.Call(opts, &out, "slt", a, b)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Slt is a free data retrieval call binding the contract method 0x42a08c38.
//
// Solidity: function slt(int256 a, int256 b) pure returns(bool)
func (_TestOpCode *TestOpCodeSession) Slt(a *big.Int, b *big.Int) (bool, error) {
	return _TestOpCode.Contract.Slt(&_TestOpCode.CallOpts, a, b)
}

// Slt is a free data retrieval call binding the contract method 0x42a08c38.
//
// Solidity: function slt(int256 a, int256 b) pure returns(bool)
func (_TestOpCode *TestOpCodeCallerSession) Slt(a *big.Int, b *big.Int) (bool, error) {
	return _TestOpCode.Contract.Slt(&_TestOpCode.CallOpts, a, b)
}

// Smod is a free data retrieval call binding the contract method 0x75df4bb9.
//
// Solidity: function smod(int256 a, int256 b) pure returns(int256)
func (_TestOpCode *TestOpCodeCaller) Smod(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TestOpCode.contract.Call(opts, &out, "smod", a, b)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Smod is a free data retrieval call binding the contract method 0x75df4bb9.
//
// Solidity: function smod(int256 a, int256 b) pure returns(int256)
func (_TestOpCode *TestOpCodeSession) Smod(a *big.Int, b *big.Int) (*big.Int, error) {
	return _TestOpCode.Contract.Smod(&_TestOpCode.CallOpts, a, b)
}

// Smod is a free data retrieval call binding the contract method 0x75df4bb9.
//
// Solidity: function smod(int256 a, int256 b) pure returns(int256)
func (_TestOpCode *TestOpCodeCallerSession) Smod(a *big.Int, b *big.Int) (*big.Int, error) {
	return _TestOpCode.Contract.Smod(&_TestOpCode.CallOpts, a, b)
}

// Sub is a free data retrieval call binding the contract method 0xb67d77c5.
//
// Solidity: function sub(uint256 a, uint256 b) pure returns(uint256)
func (_TestOpCode *TestOpCodeCaller) Sub(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TestOpCode.contract.Call(opts, &out, "sub", a, b)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Sub is a free data retrieval call binding the contract method 0xb67d77c5.
//
// Solidity: function sub(uint256 a, uint256 b) pure returns(uint256)
func (_TestOpCode *TestOpCodeSession) Sub(a *big.Int, b *big.Int) (*big.Int, error) {
	return _TestOpCode.Contract.Sub(&_TestOpCode.CallOpts, a, b)
}

// Sub is a free data retrieval call binding the contract method 0xb67d77c5.
//
// Solidity: function sub(uint256 a, uint256 b) pure returns(uint256)
func (_TestOpCode *TestOpCodeCallerSession) Sub(a *big.Int, b *big.Int) (*big.Int, error) {
	return _TestOpCode.Contract.Sub(&_TestOpCode.CallOpts, a, b)
}

// Xor is a free data retrieval call binding the contract method 0xf1388b6f.
//
// Solidity: function xor(int256 a, int256 b) pure returns(int256)
func (_TestOpCode *TestOpCodeCaller) Xor(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TestOpCode.contract.Call(opts, &out, "xor", a, b)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Xor is a free data retrieval call binding the contract method 0xf1388b6f.
//
// Solidity: function xor(int256 a, int256 b) pure returns(int256)
func (_TestOpCode *TestOpCodeSession) Xor(a *big.Int, b *big.Int) (*big.Int, error) {
	return _TestOpCode.Contract.Xor(&_TestOpCode.CallOpts, a, b)
}

// Xor is a free data retrieval call binding the contract method 0xf1388b6f.
//
// Solidity: function xor(int256 a, int256 b) pure returns(int256)
func (_TestOpCode *TestOpCodeCallerSession) Xor(a *big.Int, b *big.Int) (*big.Int, error) {
	return _TestOpCode.Contract.Xor(&_TestOpCode.CallOpts, a, b)
}

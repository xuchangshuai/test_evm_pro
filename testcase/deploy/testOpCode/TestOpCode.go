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

// TestOpCodeABI is the input ABI used to generate the binding from.
const TestOpCodeABI = "[{\"inputs\":[],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"and\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"byteFun\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"div\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eq\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exp\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gt\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lt\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mod\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mul\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"not\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"or\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sdiv\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sgt\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slt\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"smod\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sub\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"xor\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

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

// Add is a free data retrieval call binding the contract method 0x4f2be91f.
//
// Solidity: function add() pure returns()
func (_TestOpCode *TestOpCodeCaller) Add(opts *bind.CallOpts) error {
	var out []interface{}
	err := _TestOpCode.contract.Call(opts, &out, "add")

	if err != nil {
		return err
	}

	return err

}

// Add is a free data retrieval call binding the contract method 0x4f2be91f.
//
// Solidity: function add() pure returns()
func (_TestOpCode *TestOpCodeSession) Add() error {
	return _TestOpCode.Contract.Add(&_TestOpCode.CallOpts)
}

// Add is a free data retrieval call binding the contract method 0x4f2be91f.
//
// Solidity: function add() pure returns()
func (_TestOpCode *TestOpCodeCallerSession) Add() error {
	return _TestOpCode.Contract.Add(&_TestOpCode.CallOpts)
}

// And is a free data retrieval call binding the contract method 0x9542f2e0.
//
// Solidity: function and() pure returns()
func (_TestOpCode *TestOpCodeCaller) And(opts *bind.CallOpts) error {
	var out []interface{}
	err := _TestOpCode.contract.Call(opts, &out, "and")

	if err != nil {
		return err
	}

	return err

}

// And is a free data retrieval call binding the contract method 0x9542f2e0.
//
// Solidity: function and() pure returns()
func (_TestOpCode *TestOpCodeSession) And() error {
	return _TestOpCode.Contract.And(&_TestOpCode.CallOpts)
}

// And is a free data retrieval call binding the contract method 0x9542f2e0.
//
// Solidity: function and() pure returns()
func (_TestOpCode *TestOpCodeCallerSession) And() error {
	return _TestOpCode.Contract.And(&_TestOpCode.CallOpts)
}

// ByteFun is a free data retrieval call binding the contract method 0x5c12da89.
//
// Solidity: function byteFun() pure returns()
func (_TestOpCode *TestOpCodeCaller) ByteFun(opts *bind.CallOpts) error {
	var out []interface{}
	err := _TestOpCode.contract.Call(opts, &out, "byteFun")

	if err != nil {
		return err
	}

	return err

}

// ByteFun is a free data retrieval call binding the contract method 0x5c12da89.
//
// Solidity: function byteFun() pure returns()
func (_TestOpCode *TestOpCodeSession) ByteFun() error {
	return _TestOpCode.Contract.ByteFun(&_TestOpCode.CallOpts)
}

// ByteFun is a free data retrieval call binding the contract method 0x5c12da89.
//
// Solidity: function byteFun() pure returns()
func (_TestOpCode *TestOpCodeCallerSession) ByteFun() error {
	return _TestOpCode.Contract.ByteFun(&_TestOpCode.CallOpts)
}

// Div is a free data retrieval call binding the contract method 0xf9fa48c3.
//
// Solidity: function div() pure returns()
func (_TestOpCode *TestOpCodeCaller) Div(opts *bind.CallOpts) error {
	var out []interface{}
	err := _TestOpCode.contract.Call(opts, &out, "div")

	if err != nil {
		return err
	}

	return err

}

// Div is a free data retrieval call binding the contract method 0xf9fa48c3.
//
// Solidity: function div() pure returns()
func (_TestOpCode *TestOpCodeSession) Div() error {
	return _TestOpCode.Contract.Div(&_TestOpCode.CallOpts)
}

// Div is a free data retrieval call binding the contract method 0xf9fa48c3.
//
// Solidity: function div() pure returns()
func (_TestOpCode *TestOpCodeCallerSession) Div() error {
	return _TestOpCode.Contract.Div(&_TestOpCode.CallOpts)
}

// Eq is a free data retrieval call binding the contract method 0xb5165b60.
//
// Solidity: function eq() pure returns()
func (_TestOpCode *TestOpCodeCaller) Eq(opts *bind.CallOpts) error {
	var out []interface{}
	err := _TestOpCode.contract.Call(opts, &out, "eq")

	if err != nil {
		return err
	}

	return err

}

// Eq is a free data retrieval call binding the contract method 0xb5165b60.
//
// Solidity: function eq() pure returns()
func (_TestOpCode *TestOpCodeSession) Eq() error {
	return _TestOpCode.Contract.Eq(&_TestOpCode.CallOpts)
}

// Eq is a free data retrieval call binding the contract method 0xb5165b60.
//
// Solidity: function eq() pure returns()
func (_TestOpCode *TestOpCodeCallerSession) Eq() error {
	return _TestOpCode.Contract.Eq(&_TestOpCode.CallOpts)
}

// Exp is a free data retrieval call binding the contract method 0xab60ffda.
//
// Solidity: function exp() pure returns()
func (_TestOpCode *TestOpCodeCaller) Exp(opts *bind.CallOpts) error {
	var out []interface{}
	err := _TestOpCode.contract.Call(opts, &out, "exp")

	if err != nil {
		return err
	}

	return err

}

// Exp is a free data retrieval call binding the contract method 0xab60ffda.
//
// Solidity: function exp() pure returns()
func (_TestOpCode *TestOpCodeSession) Exp() error {
	return _TestOpCode.Contract.Exp(&_TestOpCode.CallOpts)
}

// Exp is a free data retrieval call binding the contract method 0xab60ffda.
//
// Solidity: function exp() pure returns()
func (_TestOpCode *TestOpCodeCallerSession) Exp() error {
	return _TestOpCode.Contract.Exp(&_TestOpCode.CallOpts)
}

// Gt is a free data retrieval call binding the contract method 0xaf0ab80a.
//
// Solidity: function gt() pure returns()
func (_TestOpCode *TestOpCodeCaller) Gt(opts *bind.CallOpts) error {
	var out []interface{}
	err := _TestOpCode.contract.Call(opts, &out, "gt")

	if err != nil {
		return err
	}

	return err

}

// Gt is a free data retrieval call binding the contract method 0xaf0ab80a.
//
// Solidity: function gt() pure returns()
func (_TestOpCode *TestOpCodeSession) Gt() error {
	return _TestOpCode.Contract.Gt(&_TestOpCode.CallOpts)
}

// Gt is a free data retrieval call binding the contract method 0xaf0ab80a.
//
// Solidity: function gt() pure returns()
func (_TestOpCode *TestOpCodeCallerSession) Gt() error {
	return _TestOpCode.Contract.Gt(&_TestOpCode.CallOpts)
}

// Lt is a free data retrieval call binding the contract method 0xa22ca2a6.
//
// Solidity: function lt() pure returns()
func (_TestOpCode *TestOpCodeCaller) Lt(opts *bind.CallOpts) error {
	var out []interface{}
	err := _TestOpCode.contract.Call(opts, &out, "lt")

	if err != nil {
		return err
	}

	return err

}

// Lt is a free data retrieval call binding the contract method 0xa22ca2a6.
//
// Solidity: function lt() pure returns()
func (_TestOpCode *TestOpCodeSession) Lt() error {
	return _TestOpCode.Contract.Lt(&_TestOpCode.CallOpts)
}

// Lt is a free data retrieval call binding the contract method 0xa22ca2a6.
//
// Solidity: function lt() pure returns()
func (_TestOpCode *TestOpCodeCallerSession) Lt() error {
	return _TestOpCode.Contract.Lt(&_TestOpCode.CallOpts)
}

// Mod is a free data retrieval call binding the contract method 0x2986e054.
//
// Solidity: function mod() pure returns()
func (_TestOpCode *TestOpCodeCaller) Mod(opts *bind.CallOpts) error {
	var out []interface{}
	err := _TestOpCode.contract.Call(opts, &out, "mod")

	if err != nil {
		return err
	}

	return err

}

// Mod is a free data retrieval call binding the contract method 0x2986e054.
//
// Solidity: function mod() pure returns()
func (_TestOpCode *TestOpCodeSession) Mod() error {
	return _TestOpCode.Contract.Mod(&_TestOpCode.CallOpts)
}

// Mod is a free data retrieval call binding the contract method 0x2986e054.
//
// Solidity: function mod() pure returns()
func (_TestOpCode *TestOpCodeCallerSession) Mod() error {
	return _TestOpCode.Contract.Mod(&_TestOpCode.CallOpts)
}

// Mul is a free data retrieval call binding the contract method 0x58931c46.
//
// Solidity: function mul() pure returns()
func (_TestOpCode *TestOpCodeCaller) Mul(opts *bind.CallOpts) error {
	var out []interface{}
	err := _TestOpCode.contract.Call(opts, &out, "mul")

	if err != nil {
		return err
	}

	return err

}

// Mul is a free data retrieval call binding the contract method 0x58931c46.
//
// Solidity: function mul() pure returns()
func (_TestOpCode *TestOpCodeSession) Mul() error {
	return _TestOpCode.Contract.Mul(&_TestOpCode.CallOpts)
}

// Mul is a free data retrieval call binding the contract method 0x58931c46.
//
// Solidity: function mul() pure returns()
func (_TestOpCode *TestOpCodeCallerSession) Mul() error {
	return _TestOpCode.Contract.Mul(&_TestOpCode.CallOpts)
}

// Not is a free data retrieval call binding the contract method 0x2e13bd10.
//
// Solidity: function not() pure returns()
func (_TestOpCode *TestOpCodeCaller) Not(opts *bind.CallOpts) error {
	var out []interface{}
	err := _TestOpCode.contract.Call(opts, &out, "not")

	if err != nil {
		return err
	}

	return err

}

// Not is a free data retrieval call binding the contract method 0x2e13bd10.
//
// Solidity: function not() pure returns()
func (_TestOpCode *TestOpCodeSession) Not() error {
	return _TestOpCode.Contract.Not(&_TestOpCode.CallOpts)
}

// Not is a free data retrieval call binding the contract method 0x2e13bd10.
//
// Solidity: function not() pure returns()
func (_TestOpCode *TestOpCodeCallerSession) Not() error {
	return _TestOpCode.Contract.Not(&_TestOpCode.CallOpts)
}

// Or is a free data retrieval call binding the contract method 0x2d5e74d8.
//
// Solidity: function or() pure returns()
func (_TestOpCode *TestOpCodeCaller) Or(opts *bind.CallOpts) error {
	var out []interface{}
	err := _TestOpCode.contract.Call(opts, &out, "or")

	if err != nil {
		return err
	}

	return err

}

// Or is a free data retrieval call binding the contract method 0x2d5e74d8.
//
// Solidity: function or() pure returns()
func (_TestOpCode *TestOpCodeSession) Or() error {
	return _TestOpCode.Contract.Or(&_TestOpCode.CallOpts)
}

// Or is a free data retrieval call binding the contract method 0x2d5e74d8.
//
// Solidity: function or() pure returns()
func (_TestOpCode *TestOpCodeCallerSession) Or() error {
	return _TestOpCode.Contract.Or(&_TestOpCode.CallOpts)
}

// Sdiv is a free data retrieval call binding the contract method 0xd5cfd270.
//
// Solidity: function sdiv() pure returns()
func (_TestOpCode *TestOpCodeCaller) Sdiv(opts *bind.CallOpts) error {
	var out []interface{}
	err := _TestOpCode.contract.Call(opts, &out, "sdiv")

	if err != nil {
		return err
	}

	return err

}

// Sdiv is a free data retrieval call binding the contract method 0xd5cfd270.
//
// Solidity: function sdiv() pure returns()
func (_TestOpCode *TestOpCodeSession) Sdiv() error {
	return _TestOpCode.Contract.Sdiv(&_TestOpCode.CallOpts)
}

// Sdiv is a free data retrieval call binding the contract method 0xd5cfd270.
//
// Solidity: function sdiv() pure returns()
func (_TestOpCode *TestOpCodeCallerSession) Sdiv() error {
	return _TestOpCode.Contract.Sdiv(&_TestOpCode.CallOpts)
}

// Sgt is a free data retrieval call binding the contract method 0x357a0ba2.
//
// Solidity: function sgt() pure returns()
func (_TestOpCode *TestOpCodeCaller) Sgt(opts *bind.CallOpts) error {
	var out []interface{}
	err := _TestOpCode.contract.Call(opts, &out, "sgt")

	if err != nil {
		return err
	}

	return err

}

// Sgt is a free data retrieval call binding the contract method 0x357a0ba2.
//
// Solidity: function sgt() pure returns()
func (_TestOpCode *TestOpCodeSession) Sgt() error {
	return _TestOpCode.Contract.Sgt(&_TestOpCode.CallOpts)
}

// Sgt is a free data retrieval call binding the contract method 0x357a0ba2.
//
// Solidity: function sgt() pure returns()
func (_TestOpCode *TestOpCodeCallerSession) Sgt() error {
	return _TestOpCode.Contract.Sgt(&_TestOpCode.CallOpts)
}

// Slt is a free data retrieval call binding the contract method 0x1c408f25.
//
// Solidity: function slt() pure returns()
func (_TestOpCode *TestOpCodeCaller) Slt(opts *bind.CallOpts) error {
	var out []interface{}
	err := _TestOpCode.contract.Call(opts, &out, "slt")

	if err != nil {
		return err
	}

	return err

}

// Slt is a free data retrieval call binding the contract method 0x1c408f25.
//
// Solidity: function slt() pure returns()
func (_TestOpCode *TestOpCodeSession) Slt() error {
	return _TestOpCode.Contract.Slt(&_TestOpCode.CallOpts)
}

// Slt is a free data retrieval call binding the contract method 0x1c408f25.
//
// Solidity: function slt() pure returns()
func (_TestOpCode *TestOpCodeCallerSession) Slt() error {
	return _TestOpCode.Contract.Slt(&_TestOpCode.CallOpts)
}

// Smod is a free data retrieval call binding the contract method 0x18572920.
//
// Solidity: function smod() pure returns()
func (_TestOpCode *TestOpCodeCaller) Smod(opts *bind.CallOpts) error {
	var out []interface{}
	err := _TestOpCode.contract.Call(opts, &out, "smod")

	if err != nil {
		return err
	}

	return err

}

// Smod is a free data retrieval call binding the contract method 0x18572920.
//
// Solidity: function smod() pure returns()
func (_TestOpCode *TestOpCodeSession) Smod() error {
	return _TestOpCode.Contract.Smod(&_TestOpCode.CallOpts)
}

// Smod is a free data retrieval call binding the contract method 0x18572920.
//
// Solidity: function smod() pure returns()
func (_TestOpCode *TestOpCodeCallerSession) Smod() error {
	return _TestOpCode.Contract.Smod(&_TestOpCode.CallOpts)
}

// Sub is a free data retrieval call binding the contract method 0xc54124be.
//
// Solidity: function sub() pure returns()
func (_TestOpCode *TestOpCodeCaller) Sub(opts *bind.CallOpts) error {
	var out []interface{}
	err := _TestOpCode.contract.Call(opts, &out, "sub")

	if err != nil {
		return err
	}

	return err

}

// Sub is a free data retrieval call binding the contract method 0xc54124be.
//
// Solidity: function sub() pure returns()
func (_TestOpCode *TestOpCodeSession) Sub() error {
	return _TestOpCode.Contract.Sub(&_TestOpCode.CallOpts)
}

// Sub is a free data retrieval call binding the contract method 0xc54124be.
//
// Solidity: function sub() pure returns()
func (_TestOpCode *TestOpCodeCallerSession) Sub() error {
	return _TestOpCode.Contract.Sub(&_TestOpCode.CallOpts)
}

// Xor is a free data retrieval call binding the contract method 0x5ada0c68.
//
// Solidity: function xor() pure returns()
func (_TestOpCode *TestOpCodeCaller) Xor(opts *bind.CallOpts) error {
	var out []interface{}
	err := _TestOpCode.contract.Call(opts, &out, "xor")

	if err != nil {
		return err
	}

	return err

}

// Xor is a free data retrieval call binding the contract method 0x5ada0c68.
//
// Solidity: function xor() pure returns()
func (_TestOpCode *TestOpCodeSession) Xor() error {
	return _TestOpCode.Contract.Xor(&_TestOpCode.CallOpts)
}

// Xor is a free data retrieval call binding the contract method 0x5ada0c68.
//
// Solidity: function xor() pure returns()
func (_TestOpCode *TestOpCodeCallerSession) Xor() error {
	return _TestOpCode.Contract.Xor(&_TestOpCode.CallOpts)
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package outputCont

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

// OutputContABI is the input ABI used to generate the binding from.
const OutputContABI = "[{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// OutputCont is an auto generated Go binding around an Ethereum contract.
type OutputCont struct {
	OutputContCaller     // Read-only binding to the contract
	OutputContTransactor // Write-only binding to the contract
	OutputContFilterer   // Log filterer for contract events
}

// OutputContCaller is an auto generated read-only Go binding around an Ethereum contract.
type OutputContCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OutputContTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OutputContTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OutputContFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OutputContFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OutputContSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OutputContSession struct {
	Contract     *OutputCont       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OutputContCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OutputContCallerSession struct {
	Contract *OutputContCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// OutputContTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OutputContTransactorSession struct {
	Contract     *OutputContTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// OutputContRaw is an auto generated low-level Go binding around an Ethereum contract.
type OutputContRaw struct {
	Contract *OutputCont // Generic contract binding to access the raw methods on
}

// OutputContCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OutputContCallerRaw struct {
	Contract *OutputContCaller // Generic read-only contract binding to access the raw methods on
}

// OutputContTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OutputContTransactorRaw struct {
	Contract *OutputContTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOutputCont creates a new instance of OutputCont, bound to a specific deployed contract.
func NewOutputCont(address common.Address, backend bind.ContractBackend) (*OutputCont, error) {
	contract, err := bindOutputCont(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OutputCont{OutputContCaller: OutputContCaller{contract: contract}, OutputContTransactor: OutputContTransactor{contract: contract}, OutputContFilterer: OutputContFilterer{contract: contract}}, nil
}

// NewOutputContCaller creates a new read-only instance of OutputCont, bound to a specific deployed contract.
func NewOutputContCaller(address common.Address, caller bind.ContractCaller) (*OutputContCaller, error) {
	contract, err := bindOutputCont(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OutputContCaller{contract: contract}, nil
}

// NewOutputContTransactor creates a new write-only instance of OutputCont, bound to a specific deployed contract.
func NewOutputContTransactor(address common.Address, transactor bind.ContractTransactor) (*OutputContTransactor, error) {
	contract, err := bindOutputCont(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OutputContTransactor{contract: contract}, nil
}

// NewOutputContFilterer creates a new log filterer instance of OutputCont, bound to a specific deployed contract.
func NewOutputContFilterer(address common.Address, filterer bind.ContractFilterer) (*OutputContFilterer, error) {
	contract, err := bindOutputCont(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OutputContFilterer{contract: contract}, nil
}

// bindOutputCont binds a generic wrapper to an already deployed contract.
func bindOutputCont(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OutputContABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OutputCont *OutputContRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OutputCont.Contract.OutputContCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OutputCont *OutputContRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OutputCont.Contract.OutputContTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OutputCont *OutputContRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OutputCont.Contract.OutputContTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OutputCont *OutputContCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OutputCont.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OutputCont *OutputContTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OutputCont.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OutputCont *OutputContTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OutputCont.Contract.contract.Transact(opts, method, params...)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256 balance)
func (_OutputCont *OutputContCaller) GetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OutputCont.contract.Call(opts, &out, "getBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256 balance)
func (_OutputCont *OutputContSession) GetBalance() (*big.Int, error) {
	return _OutputCont.Contract.GetBalance(&_OutputCont.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256 balance)
func (_OutputCont *OutputContCallerSession) GetBalance() (*big.Int, error) {
	return _OutputCont.Contract.GetBalance(&_OutputCont.CallOpts)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_OutputCont *OutputContTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _OutputCont.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_OutputCont *OutputContSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _OutputCont.Contract.Fallback(&_OutputCont.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_OutputCont *OutputContTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _OutputCont.Contract.Fallback(&_OutputCont.TransactOpts, calldata)
}

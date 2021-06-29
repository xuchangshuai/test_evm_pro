// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package inputCont

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

// InputContABI is the input ABI used to generate the binding from.
const InputContABI = "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"outputAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"_outputAddress\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transferAccounts\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// InputCont is an auto generated Go binding around an Ethereum contract.
type InputCont struct {
	InputContCaller     // Read-only binding to the contract
	InputContTransactor // Write-only binding to the contract
	InputContFilterer   // Log filterer for contract events
}

// InputContCaller is an auto generated read-only Go binding around an Ethereum contract.
type InputContCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InputContTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InputContTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InputContFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InputContFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InputContSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InputContSession struct {
	Contract     *InputCont        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InputContCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InputContCallerSession struct {
	Contract *InputContCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// InputContTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InputContTransactorSession struct {
	Contract     *InputContTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// InputContRaw is an auto generated low-level Go binding around an Ethereum contract.
type InputContRaw struct {
	Contract *InputCont // Generic contract binding to access the raw methods on
}

// InputContCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InputContCallerRaw struct {
	Contract *InputContCaller // Generic read-only contract binding to access the raw methods on
}

// InputContTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InputContTransactorRaw struct {
	Contract *InputContTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInputCont creates a new instance of InputCont, bound to a specific deployed contract.
func NewInputCont(address common.Address, backend bind.ContractBackend) (*InputCont, error) {
	contract, err := bindInputCont(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InputCont{InputContCaller: InputContCaller{contract: contract}, InputContTransactor: InputContTransactor{contract: contract}, InputContFilterer: InputContFilterer{contract: contract}}, nil
}

// NewInputContCaller creates a new read-only instance of InputCont, bound to a specific deployed contract.
func NewInputContCaller(address common.Address, caller bind.ContractCaller) (*InputContCaller, error) {
	contract, err := bindInputCont(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InputContCaller{contract: contract}, nil
}

// NewInputContTransactor creates a new write-only instance of InputCont, bound to a specific deployed contract.
func NewInputContTransactor(address common.Address, transactor bind.ContractTransactor) (*InputContTransactor, error) {
	contract, err := bindInputCont(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InputContTransactor{contract: contract}, nil
}

// NewInputContFilterer creates a new log filterer instance of InputCont, bound to a specific deployed contract.
func NewInputContFilterer(address common.Address, filterer bind.ContractFilterer) (*InputContFilterer, error) {
	contract, err := bindInputCont(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InputContFilterer{contract: contract}, nil
}

// bindInputCont binds a generic wrapper to an already deployed contract.
func bindInputCont(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InputContABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InputCont *InputContRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InputCont.Contract.InputContCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InputCont *InputContRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InputCont.Contract.InputContTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InputCont *InputContRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InputCont.Contract.InputContTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InputCont *InputContCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InputCont.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InputCont *InputContTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InputCont.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InputCont *InputContTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InputCont.Contract.contract.Transact(opts, method, params...)
}

// OutputAddress is a free data retrieval call binding the contract method 0xca28cda4.
//
// Solidity: function _outputAddress() view returns(address)
func (_InputCont *InputContCaller) OutputAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InputCont.contract.Call(opts, &out, "_outputAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OutputAddress is a free data retrieval call binding the contract method 0xca28cda4.
//
// Solidity: function _outputAddress() view returns(address)
func (_InputCont *InputContSession) OutputAddress() (common.Address, error) {
	return _InputCont.Contract.OutputAddress(&_InputCont.CallOpts)
}

// OutputAddress is a free data retrieval call binding the contract method 0xca28cda4.
//
// Solidity: function _outputAddress() view returns(address)
func (_InputCont *InputContCallerSession) OutputAddress() (common.Address, error) {
	return _InputCont.Contract.OutputAddress(&_InputCont.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256 balance)
func (_InputCont *InputContCaller) GetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InputCont.contract.Call(opts, &out, "getBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256 balance)
func (_InputCont *InputContSession) GetBalance() (*big.Int, error) {
	return _InputCont.Contract.GetBalance(&_InputCont.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256 balance)
func (_InputCont *InputContCallerSession) GetBalance() (*big.Int, error) {
	return _InputCont.Contract.GetBalance(&_InputCont.CallOpts)
}

// TransferAccounts is a paid mutator transaction binding the contract method 0x2372cdde.
//
// Solidity: function transferAccounts() payable returns(bool success)
func (_InputCont *InputContTransactor) TransferAccounts(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InputCont.contract.Transact(opts, "transferAccounts")
}

// TransferAccounts is a paid mutator transaction binding the contract method 0x2372cdde.
//
// Solidity: function transferAccounts() payable returns(bool success)
func (_InputCont *InputContSession) TransferAccounts() (*types.Transaction, error) {
	return _InputCont.Contract.TransferAccounts(&_InputCont.TransactOpts)
}

// TransferAccounts is a paid mutator transaction binding the contract method 0x2372cdde.
//
// Solidity: function transferAccounts() payable returns(bool success)
func (_InputCont *InputContTransactorSession) TransferAccounts() (*types.Transaction, error) {
	return _InputCont.Contract.TransferAccounts(&_InputCont.TransactOpts)
}

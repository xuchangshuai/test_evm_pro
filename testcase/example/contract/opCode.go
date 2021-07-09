// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// OpcodeABI is the input ABI used to generate the binding from.
const OpcodeABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"add\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// Opcode is an auto generated Go binding around an Ethereum contract.
type Opcode struct {
	OpcodeCaller     // Read-only binding to the contract
	OpcodeTransactor // Write-only binding to the contract
	OpcodeFilterer   // Log filterer for contract events
}

// OpcodeCaller is an auto generated read-only Go binding around an Ethereum contract.
type OpcodeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpcodeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OpcodeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpcodeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OpcodeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpcodeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OpcodeSession struct {
	Contract     *Opcode           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OpcodeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OpcodeCallerSession struct {
	Contract *OpcodeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OpcodeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OpcodeTransactorSession struct {
	Contract     *OpcodeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OpcodeRaw is an auto generated low-level Go binding around an Ethereum contract.
type OpcodeRaw struct {
	Contract *Opcode // Generic contract binding to access the raw methods on
}

// OpcodeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OpcodeCallerRaw struct {
	Contract *OpcodeCaller // Generic read-only contract binding to access the raw methods on
}

// OpcodeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OpcodeTransactorRaw struct {
	Contract *OpcodeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOpcode creates a new instance of Opcode, bound to a specific deployed contract.
func NewOpcode(address common.Address, backend bind.ContractBackend) (*Opcode, error) {
	contract, err := bindOpcode(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Opcode{OpcodeCaller: OpcodeCaller{contract: contract}, OpcodeTransactor: OpcodeTransactor{contract: contract}, OpcodeFilterer: OpcodeFilterer{contract: contract}}, nil
}

// NewOpcodeCaller creates a new read-only instance of Opcode, bound to a specific deployed contract.
func NewOpcodeCaller(address common.Address, caller bind.ContractCaller) (*OpcodeCaller, error) {
	contract, err := bindOpcode(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OpcodeCaller{contract: contract}, nil
}

// NewOpcodeTransactor creates a new write-only instance of Opcode, bound to a specific deployed contract.
func NewOpcodeTransactor(address common.Address, transactor bind.ContractTransactor) (*OpcodeTransactor, error) {
	contract, err := bindOpcode(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OpcodeTransactor{contract: contract}, nil
}

// NewOpcodeFilterer creates a new log filterer instance of Opcode, bound to a specific deployed contract.
func NewOpcodeFilterer(address common.Address, filterer bind.ContractFilterer) (*OpcodeFilterer, error) {
	contract, err := bindOpcode(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OpcodeFilterer{contract: contract}, nil
}

// bindOpcode binds a generic wrapper to an already deployed contract.
func bindOpcode(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OpcodeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Opcode *OpcodeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Opcode.Contract.OpcodeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Opcode *OpcodeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Opcode.Contract.OpcodeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Opcode *OpcodeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Opcode.Contract.OpcodeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Opcode *OpcodeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Opcode.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Opcode *OpcodeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Opcode.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Opcode *OpcodeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Opcode.Contract.contract.Transact(opts, method, params...)
}

// Add is a paid mutator transaction binding the contract method 0x1003e2d2.
//
// Solidity: function add(uint256 i) payable returns(uint256)
func (_Opcode *OpcodeTransactor) Add(opts *bind.TransactOpts, i *big.Int) (*types.Transaction, error) {
	return _Opcode.contract.Transact(opts, "add", i)
}

// Add is a paid mutator transaction binding the contract method 0x1003e2d2.
//
// Solidity: function add(uint256 i) payable returns(uint256)
func (_Opcode *OpcodeSession) Add(i *big.Int) (*types.Transaction, error) {
	return _Opcode.Contract.Add(&_Opcode.TransactOpts, i)
}

// Add is a paid mutator transaction binding the contract method 0x1003e2d2.
//
// Solidity: function add(uint256 i) payable returns(uint256)
func (_Opcode *OpcodeTransactorSession) Add(i *big.Int) (*types.Transaction, error) {
	return _Opcode.Contract.Add(&_Opcode.TransactOpts, i)
}

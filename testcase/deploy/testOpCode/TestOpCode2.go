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

// TestOpCode2ABI is the input ABI used to generate the binding from.
const TestOpCode2ABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"getAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlockDifficulty\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"getBlockHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCallDataCopy\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCaller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCoinBase\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGasPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOrigin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mload\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mstore\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pop\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// TestOpCode2 is an auto generated Go binding around an Ethereum contract.
type TestOpCode2 struct {
	TestOpCode2Caller     // Read-only binding to the contract
	TestOpCode2Transactor // Write-only binding to the contract
	TestOpCode2Filterer   // Log filterer for contract events
}

// TestOpCode2Caller is an auto generated read-only Go binding around an Ethereum contract.
type TestOpCode2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestOpCode2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type TestOpCode2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestOpCode2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestOpCode2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestOpCode2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestOpCode2Session struct {
	Contract     *TestOpCode2      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestOpCode2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestOpCode2CallerSession struct {
	Contract *TestOpCode2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// TestOpCode2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestOpCode2TransactorSession struct {
	Contract     *TestOpCode2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TestOpCode2Raw is an auto generated low-level Go binding around an Ethereum contract.
type TestOpCode2Raw struct {
	Contract *TestOpCode2 // Generic contract binding to access the raw methods on
}

// TestOpCode2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestOpCode2CallerRaw struct {
	Contract *TestOpCode2Caller // Generic read-only contract binding to access the raw methods on
}

// TestOpCode2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestOpCode2TransactorRaw struct {
	Contract *TestOpCode2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewTestOpCode2 creates a new instance of TestOpCode2, bound to a specific deployed contract.
func NewTestOpCode2(address common.Address, backend bind.ContractBackend) (*TestOpCode2, error) {
	contract, err := bindTestOpCode2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestOpCode2{TestOpCode2Caller: TestOpCode2Caller{contract: contract}, TestOpCode2Transactor: TestOpCode2Transactor{contract: contract}, TestOpCode2Filterer: TestOpCode2Filterer{contract: contract}}, nil
}

// NewTestOpCode2Caller creates a new read-only instance of TestOpCode2, bound to a specific deployed contract.
func NewTestOpCode2Caller(address common.Address, caller bind.ContractCaller) (*TestOpCode2Caller, error) {
	contract, err := bindTestOpCode2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestOpCode2Caller{contract: contract}, nil
}

// NewTestOpCode2Transactor creates a new write-only instance of TestOpCode2, bound to a specific deployed contract.
func NewTestOpCode2Transactor(address common.Address, transactor bind.ContractTransactor) (*TestOpCode2Transactor, error) {
	contract, err := bindTestOpCode2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestOpCode2Transactor{contract: contract}, nil
}

// NewTestOpCode2Filterer creates a new log filterer instance of TestOpCode2, bound to a specific deployed contract.
func NewTestOpCode2Filterer(address common.Address, filterer bind.ContractFilterer) (*TestOpCode2Filterer, error) {
	contract, err := bindTestOpCode2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestOpCode2Filterer{contract: contract}, nil
}

// bindTestOpCode2 binds a generic wrapper to an already deployed contract.
func bindTestOpCode2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestOpCode2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestOpCode2 *TestOpCode2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestOpCode2.Contract.TestOpCode2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestOpCode2 *TestOpCode2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestOpCode2.Contract.TestOpCode2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestOpCode2 *TestOpCode2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestOpCode2.Contract.TestOpCode2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestOpCode2 *TestOpCode2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestOpCode2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestOpCode2 *TestOpCode2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestOpCode2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestOpCode2 *TestOpCode2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestOpCode2.Contract.contract.Transact(opts, method, params...)
}

// GetAddress is a free data retrieval call binding the contract method 0x38cc4831.
//
// Solidity: function getAddress() view returns(address)
func (_TestOpCode2 *TestOpCode2Caller) GetAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TestOpCode2.contract.Call(opts, &out, "getAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddress is a free data retrieval call binding the contract method 0x38cc4831.
//
// Solidity: function getAddress() view returns(address)
func (_TestOpCode2 *TestOpCode2Session) GetAddress() (common.Address, error) {
	return _TestOpCode2.Contract.GetAddress(&_TestOpCode2.CallOpts)
}

// GetAddress is a free data retrieval call binding the contract method 0x38cc4831.
//
// Solidity: function getAddress() view returns(address)
func (_TestOpCode2 *TestOpCode2CallerSession) GetAddress() (common.Address, error) {
	return _TestOpCode2.Contract.GetAddress(&_TestOpCode2.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address addr) view returns(uint256)
func (_TestOpCode2 *TestOpCode2Caller) GetBalance(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TestOpCode2.contract.Call(opts, &out, "getBalance", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address addr) view returns(uint256)
func (_TestOpCode2 *TestOpCode2Session) GetBalance(addr common.Address) (*big.Int, error) {
	return _TestOpCode2.Contract.GetBalance(&_TestOpCode2.CallOpts, addr)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address addr) view returns(uint256)
func (_TestOpCode2 *TestOpCode2CallerSession) GetBalance(addr common.Address) (*big.Int, error) {
	return _TestOpCode2.Contract.GetBalance(&_TestOpCode2.CallOpts, addr)
}

// GetBlockDifficulty is a free data retrieval call binding the contract method 0x12e05dd1.
//
// Solidity: function getBlockDifficulty() view returns(uint256)
func (_TestOpCode2 *TestOpCode2Caller) GetBlockDifficulty(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestOpCode2.contract.Call(opts, &out, "getBlockDifficulty")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlockDifficulty is a free data retrieval call binding the contract method 0x12e05dd1.
//
// Solidity: function getBlockDifficulty() view returns(uint256)
func (_TestOpCode2 *TestOpCode2Session) GetBlockDifficulty() (*big.Int, error) {
	return _TestOpCode2.Contract.GetBlockDifficulty(&_TestOpCode2.CallOpts)
}

// GetBlockDifficulty is a free data retrieval call binding the contract method 0x12e05dd1.
//
// Solidity: function getBlockDifficulty() view returns(uint256)
func (_TestOpCode2 *TestOpCode2CallerSession) GetBlockDifficulty() (*big.Int, error) {
	return _TestOpCode2.Contract.GetBlockDifficulty(&_TestOpCode2.CallOpts)
}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 num) view returns(bytes32)
func (_TestOpCode2 *TestOpCode2Caller) GetBlockHash(opts *bind.CallOpts, num *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _TestOpCode2.contract.Call(opts, &out, "getBlockHash", num)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 num) view returns(bytes32)
func (_TestOpCode2 *TestOpCode2Session) GetBlockHash(num *big.Int) ([32]byte, error) {
	return _TestOpCode2.Contract.GetBlockHash(&_TestOpCode2.CallOpts, num)
}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 num) view returns(bytes32)
func (_TestOpCode2 *TestOpCode2CallerSession) GetBlockHash(num *big.Int) ([32]byte, error) {
	return _TestOpCode2.Contract.GetBlockHash(&_TestOpCode2.CallOpts, num)
}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256)
func (_TestOpCode2 *TestOpCode2Caller) GetBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestOpCode2.contract.Call(opts, &out, "getBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256)
func (_TestOpCode2 *TestOpCode2Session) GetBlockNumber() (*big.Int, error) {
	return _TestOpCode2.Contract.GetBlockNumber(&_TestOpCode2.CallOpts)
}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256)
func (_TestOpCode2 *TestOpCode2CallerSession) GetBlockNumber() (*big.Int, error) {
	return _TestOpCode2.Contract.GetBlockNumber(&_TestOpCode2.CallOpts)
}

// GetCallDataCopy is a free data retrieval call binding the contract method 0x02e84ecd.
//
// Solidity: function getCallDataCopy() pure returns(bytes)
func (_TestOpCode2 *TestOpCode2Caller) GetCallDataCopy(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _TestOpCode2.contract.Call(opts, &out, "getCallDataCopy")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetCallDataCopy is a free data retrieval call binding the contract method 0x02e84ecd.
//
// Solidity: function getCallDataCopy() pure returns(bytes)
func (_TestOpCode2 *TestOpCode2Session) GetCallDataCopy() ([]byte, error) {
	return _TestOpCode2.Contract.GetCallDataCopy(&_TestOpCode2.CallOpts)
}

// GetCallDataCopy is a free data retrieval call binding the contract method 0x02e84ecd.
//
// Solidity: function getCallDataCopy() pure returns(bytes)
func (_TestOpCode2 *TestOpCode2CallerSession) GetCallDataCopy() ([]byte, error) {
	return _TestOpCode2.Contract.GetCallDataCopy(&_TestOpCode2.CallOpts)
}

// GetCaller is a free data retrieval call binding the contract method 0xab470f05.
//
// Solidity: function getCaller() view returns(address)
func (_TestOpCode2 *TestOpCode2Caller) GetCaller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TestOpCode2.contract.Call(opts, &out, "getCaller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCaller is a free data retrieval call binding the contract method 0xab470f05.
//
// Solidity: function getCaller() view returns(address)
func (_TestOpCode2 *TestOpCode2Session) GetCaller() (common.Address, error) {
	return _TestOpCode2.Contract.GetCaller(&_TestOpCode2.CallOpts)
}

// GetCaller is a free data retrieval call binding the contract method 0xab470f05.
//
// Solidity: function getCaller() view returns(address)
func (_TestOpCode2 *TestOpCode2CallerSession) GetCaller() (common.Address, error) {
	return _TestOpCode2.Contract.GetCaller(&_TestOpCode2.CallOpts)
}

// GetCoinBase is a free data retrieval call binding the contract method 0x971ce641.
//
// Solidity: function getCoinBase() view returns(address)
func (_TestOpCode2 *TestOpCode2Caller) GetCoinBase(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TestOpCode2.contract.Call(opts, &out, "getCoinBase")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCoinBase is a free data retrieval call binding the contract method 0x971ce641.
//
// Solidity: function getCoinBase() view returns(address)
func (_TestOpCode2 *TestOpCode2Session) GetCoinBase() (common.Address, error) {
	return _TestOpCode2.Contract.GetCoinBase(&_TestOpCode2.CallOpts)
}

// GetCoinBase is a free data retrieval call binding the contract method 0x971ce641.
//
// Solidity: function getCoinBase() view returns(address)
func (_TestOpCode2 *TestOpCode2CallerSession) GetCoinBase() (common.Address, error) {
	return _TestOpCode2.Contract.GetCoinBase(&_TestOpCode2.CallOpts)
}

// GetData is a free data retrieval call binding the contract method 0x3bc5de30.
//
// Solidity: function getData() pure returns(bytes)
func (_TestOpCode2 *TestOpCode2Caller) GetData(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _TestOpCode2.contract.Call(opts, &out, "getData")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetData is a free data retrieval call binding the contract method 0x3bc5de30.
//
// Solidity: function getData() pure returns(bytes)
func (_TestOpCode2 *TestOpCode2Session) GetData() ([]byte, error) {
	return _TestOpCode2.Contract.GetData(&_TestOpCode2.CallOpts)
}

// GetData is a free data retrieval call binding the contract method 0x3bc5de30.
//
// Solidity: function getData() pure returns(bytes)
func (_TestOpCode2 *TestOpCode2CallerSession) GetData() ([]byte, error) {
	return _TestOpCode2.Contract.GetData(&_TestOpCode2.CallOpts)
}

// GetGasPrice is a free data retrieval call binding the contract method 0x455259cb.
//
// Solidity: function getGasPrice() view returns(uint256)
func (_TestOpCode2 *TestOpCode2Caller) GetGasPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestOpCode2.contract.Call(opts, &out, "getGasPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGasPrice is a free data retrieval call binding the contract method 0x455259cb.
//
// Solidity: function getGasPrice() view returns(uint256)
func (_TestOpCode2 *TestOpCode2Session) GetGasPrice() (*big.Int, error) {
	return _TestOpCode2.Contract.GetGasPrice(&_TestOpCode2.CallOpts)
}

// GetGasPrice is a free data retrieval call binding the contract method 0x455259cb.
//
// Solidity: function getGasPrice() view returns(uint256)
func (_TestOpCode2 *TestOpCode2CallerSession) GetGasPrice() (*big.Int, error) {
	return _TestOpCode2.Contract.GetGasPrice(&_TestOpCode2.CallOpts)
}

// GetOrigin is a free data retrieval call binding the contract method 0xdf1f29ee.
//
// Solidity: function getOrigin() view returns(address)
func (_TestOpCode2 *TestOpCode2Caller) GetOrigin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TestOpCode2.contract.Call(opts, &out, "getOrigin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOrigin is a free data retrieval call binding the contract method 0xdf1f29ee.
//
// Solidity: function getOrigin() view returns(address)
func (_TestOpCode2 *TestOpCode2Session) GetOrigin() (common.Address, error) {
	return _TestOpCode2.Contract.GetOrigin(&_TestOpCode2.CallOpts)
}

// GetOrigin is a free data retrieval call binding the contract method 0xdf1f29ee.
//
// Solidity: function getOrigin() view returns(address)
func (_TestOpCode2 *TestOpCode2CallerSession) GetOrigin() (common.Address, error) {
	return _TestOpCode2.Contract.GetOrigin(&_TestOpCode2.CallOpts)
}

// GetTimestamp is a free data retrieval call binding the contract method 0x188ec356.
//
// Solidity: function getTimestamp() view returns(uint256)
func (_TestOpCode2 *TestOpCode2Caller) GetTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestOpCode2.contract.Call(opts, &out, "getTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTimestamp is a free data retrieval call binding the contract method 0x188ec356.
//
// Solidity: function getTimestamp() view returns(uint256)
func (_TestOpCode2 *TestOpCode2Session) GetTimestamp() (*big.Int, error) {
	return _TestOpCode2.Contract.GetTimestamp(&_TestOpCode2.CallOpts)
}

// GetTimestamp is a free data retrieval call binding the contract method 0x188ec356.
//
// Solidity: function getTimestamp() view returns(uint256)
func (_TestOpCode2 *TestOpCode2CallerSession) GetTimestamp() (*big.Int, error) {
	return _TestOpCode2.Contract.GetTimestamp(&_TestOpCode2.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_TestOpCode2 *TestOpCode2Caller) Minter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TestOpCode2.contract.Call(opts, &out, "minter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_TestOpCode2 *TestOpCode2Session) Minter() (common.Address, error) {
	return _TestOpCode2.Contract.Minter(&_TestOpCode2.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_TestOpCode2 *TestOpCode2CallerSession) Minter() (common.Address, error) {
	return _TestOpCode2.Contract.Minter(&_TestOpCode2.CallOpts)
}

// Mload is a free data retrieval call binding the contract method 0x36ff50b0.
//
// Solidity: function mload() pure returns(int256)
func (_TestOpCode2 *TestOpCode2Caller) Mload(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestOpCode2.contract.Call(opts, &out, "mload")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Mload is a free data retrieval call binding the contract method 0x36ff50b0.
//
// Solidity: function mload() pure returns(int256)
func (_TestOpCode2 *TestOpCode2Session) Mload() (*big.Int, error) {
	return _TestOpCode2.Contract.Mload(&_TestOpCode2.CallOpts)
}

// Mload is a free data retrieval call binding the contract method 0x36ff50b0.
//
// Solidity: function mload() pure returns(int256)
func (_TestOpCode2 *TestOpCode2CallerSession) Mload() (*big.Int, error) {
	return _TestOpCode2.Contract.Mload(&_TestOpCode2.CallOpts)
}

// Mstore is a free data retrieval call binding the contract method 0x0794157a.
//
// Solidity: function mstore() pure returns()
func (_TestOpCode2 *TestOpCode2Caller) Mstore(opts *bind.CallOpts) error {
	var out []interface{}
	err := _TestOpCode2.contract.Call(opts, &out, "mstore")

	if err != nil {
		return err
	}

	return err

}

// Mstore is a free data retrieval call binding the contract method 0x0794157a.
//
// Solidity: function mstore() pure returns()
func (_TestOpCode2 *TestOpCode2Session) Mstore() error {
	return _TestOpCode2.Contract.Mstore(&_TestOpCode2.CallOpts)
}

// Mstore is a free data retrieval call binding the contract method 0x0794157a.
//
// Solidity: function mstore() pure returns()
func (_TestOpCode2 *TestOpCode2CallerSession) Mstore() error {
	return _TestOpCode2.Contract.Mstore(&_TestOpCode2.CallOpts)
}

// Pop is a free data retrieval call binding the contract method 0xa4ece52c.
//
// Solidity: function pop() pure returns()
func (_TestOpCode2 *TestOpCode2Caller) Pop(opts *bind.CallOpts) error {
	var out []interface{}
	err := _TestOpCode2.contract.Call(opts, &out, "pop")

	if err != nil {
		return err
	}

	return err

}

// Pop is a free data retrieval call binding the contract method 0xa4ece52c.
//
// Solidity: function pop() pure returns()
func (_TestOpCode2 *TestOpCode2Session) Pop() error {
	return _TestOpCode2.Contract.Pop(&_TestOpCode2.CallOpts)
}

// Pop is a free data retrieval call binding the contract method 0xa4ece52c.
//
// Solidity: function pop() pure returns()
func (_TestOpCode2 *TestOpCode2CallerSession) Pop() error {
	return _TestOpCode2.Contract.Pop(&_TestOpCode2.CallOpts)
}

// GetValue is a paid mutator transaction binding the contract method 0x20965255.
//
// Solidity: function getValue() payable returns(uint256)
func (_TestOpCode2 *TestOpCode2Transactor) GetValue(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestOpCode2.contract.Transact(opts, "getValue")
}

// GetValue is a paid mutator transaction binding the contract method 0x20965255.
//
// Solidity: function getValue() payable returns(uint256)
func (_TestOpCode2 *TestOpCode2Session) GetValue() (*types.Transaction, error) {
	return _TestOpCode2.Contract.GetValue(&_TestOpCode2.TransactOpts)
}

// GetValue is a paid mutator transaction binding the contract method 0x20965255.
//
// Solidity: function getValue() payable returns(uint256)
func (_TestOpCode2 *TestOpCode2TransactorSession) GetValue() (*types.Transaction, error) {
	return _TestOpCode2.Contract.GetValue(&_TestOpCode2.TransactOpts)
}

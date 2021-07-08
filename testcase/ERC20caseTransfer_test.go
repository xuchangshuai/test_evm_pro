package testcase

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/smartystreets/goconvey/convey"
	"log"
	"math/big"
	"test_evm/config"
	"test_evm/testcase/deploy"
	ERC20Test "test_evm/testcase/deploy/testEPC20"
	"test_evm/utils"
	"testing"
)

func TestERC20_transfer(t *testing.T) {
	Convey("TestERC20_transfer_01 合约拥有者账户转账给config.getAddressList[0]转账5000代币", t, func() {
		ethClient := config.GetEthClient()
		chainId, err := ethClient.ChainID(context.Background())
		_checkErr(err)

		// 备注: 对应的MetaMask钱包第一个账户私钥
		fromPrivateKey := config.FromPrivate
		fromPrivate, _ := crypto.HexToECDSA(fromPrivateKey)
		fromAddress := crypto.PubkeyToAddress(fromPrivate.PublicKey)
		log.Printf("fromAddress: %s", fromAddress.String())
		erc20ContractAddress := getERC20ContractAddress()
		erc20Contract, err := ERC20Test.NewERC20Test(erc20ContractAddress, ethClient)
		_checkErr(err)
		callOpts := &bind.CallOpts{
			From: fromAddress,
		}
		erp20Name, err := erc20Contract.Name(callOpts)
		log.Printf("erp20Name: %s", erp20Name)
		erp20Symbol, err := erc20Contract.Symbol(callOpts)
		log.Printf("erp20Symbol: %s", erp20Symbol)

		nonce, err := ethClient.PendingNonceAt(context.Background(), fromAddress)
		_checkErr(err)
		opts, err := bind.NewKeyedTransactorWithChainID(fromPrivate, chainId)
		_checkErr(err)
		opts.Nonce = big.NewInt(int64(nonce))
		opts.GasPrice = big.NewInt(500)
		opts.GasLimit = 3000000
		_to1 := config.GetAddressList()[0]
		//代币转账前
		fromBalance, err := erc20Contract.BalanceOf(callOpts, fromAddress)
		_checkErr(err)
		toBalance, err := erc20Contract.BalanceOf(callOpts, _to1)
		_checkErr(err)

		amount := big.NewInt(5000)
		fromBalanceOng := utils.GetBalance(fromAddress.String())
		res, err := erc20Contract.Transfer(opts, _to1, amount)
		_checkErr(err)
		log.Printf("res.Hash(): %s", res.Hash().String())
		receipt, err := utils.GetTransferInfoByHash(res.Hash())
		_checkErr(err)
		log.Printf("receipt.Status: %d", receipt.Status)
		if receipt.Status != 1 {
			panic("Tx fail!")
		}
		useGas := big.NewInt(int64(receipt.CumulativeGasUsed))
		log.Printf("useGas: %d", useGas)
		fromBalanceOngAfter := utils.GetBalance(fromAddress.String())
		log.Printf("fromBalanceOng: %d", fromBalanceOng)
		log.Printf("fromBalanceOngAfter: %d", fromBalanceOngAfter)

		// 断言ong消耗
		So(fromBalanceOng.Cmp(fromBalanceOngAfter.Add(fromBalanceOngAfter, useGas)) == 0, ShouldBeTrue)
		toBalanceAfter, err := erc20Contract.BalanceOf(callOpts, _to1)
		_checkErr(err)
		//log.Printf("toBalanceAfter: %d", toBalanceAfter)
		fromBalanceAfter, err := erc20Contract.BalanceOf(callOpts, fromAddress)
		_checkErr(err)
		//log.Printf("fromBalanceAfter: %d", fromBalanceAfter)

		// 断言代币转账后from账户资金变动情况
		So(fromBalance.Cmp(fromBalanceAfter.Add(fromBalanceAfter, amount)) == 0, ShouldBeTrue)
		// 断言代币转账后to账户资金变动情况
		So(toBalanceAfter.Cmp(toBalance.Add(toBalance, amount)) == 0, ShouldBeTrue)
	})
	Convey("TestERC20_transfer_02 合约拥有者账户转账给config.getAddressList[0]转账0代币", t, func() {
		ethClient := config.GetEthClient()
		chainId, err := ethClient.ChainID(context.Background())
		_checkErr(err)

		// 备注: 对应的MetaMask钱包第一个账户私钥
		fromPrivateKey := config.FromPrivate
		fromPrivate, _ := crypto.HexToECDSA(fromPrivateKey)
		fromAddress := crypto.PubkeyToAddress(fromPrivate.PublicKey)
		log.Printf("fromAddress: %s", fromAddress.String())
		erc20ContractAddress := getERC20ContractAddress()
		erc20Contract, err := ERC20Test.NewERC20Test(erc20ContractAddress, ethClient)
		_checkErr(err)
		callOpts := &bind.CallOpts{
			From: fromAddress,
		}
		erp20Name, err := erc20Contract.Name(callOpts)
		log.Printf("erp20Name: %s", erp20Name)
		erp20Symbol, err := erc20Contract.Symbol(callOpts)
		log.Printf("erp20Symbol: %s", erp20Symbol)

		nonce, err := ethClient.PendingNonceAt(context.Background(), fromAddress)
		_checkErr(err)
		opts, err := bind.NewKeyedTransactorWithChainID(fromPrivate, chainId)
		_checkErr(err)
		opts.Nonce = big.NewInt(int64(nonce))
		opts.GasPrice = big.NewInt(500)
		opts.GasLimit = 3000000
		_to1 := config.GetAddressList()[1]
		//代币转账前
		fromBalance, err := erc20Contract.BalanceOf(callOpts, fromAddress)
		_checkErr(err)
		toBalance, err := erc20Contract.BalanceOf(callOpts, _to1)
		_checkErr(err)

		amount := big.NewInt(0)
		fromBalanceOng := utils.GetBalance(fromAddress.String())
		res, err := erc20Contract.Transfer(opts, _to1, amount)
		_checkErr(err)
		log.Printf("res.Hash(): %s", res.Hash().String())
		receipt, err := utils.GetTransferInfoByHash(res.Hash())
		_checkErr(err)
		log.Printf("receipt.Status: %d", receipt.Status)
		useGas := big.NewInt(int64(receipt.CumulativeGasUsed))
		fromBalanceOngAfter := utils.GetBalance(fromAddress.String())
		//断言消耗ong
		So(fromBalanceOng.Cmp(fromBalanceOngAfter.Add(fromBalanceOngAfter, useGas)) == 0, ShouldBeTrue)
		// 断言代币转账状态0失败
		So(receipt.Status == 0, ShouldBeTrue)
		toBalanceAfter, err := erc20Contract.BalanceOf(callOpts, _to1)
		_checkErr(err)
		//log.Printf("toBalanceAfter: %d", toBalanceAfter)
		fromBalanceAfter, err := erc20Contract.BalanceOf(callOpts, fromAddress)
		_checkErr(err)
		//log.Printf("fromBalanceAfter: %d", fromBalanceAfter)

		// 断言代币转账后from账户资金变动情况
		So(fromBalance.Cmp(fromBalanceAfter) == 0, ShouldBeTrue)
		// 断言代币转账后to账户资金变动情况
		So(toBalanceAfter.Cmp(toBalance) == 0, ShouldBeTrue)

	})
	Convey("TestERC20_transfer_03 config.getAddressList[0]转账给config.getAddressList[1]转账（数量小于自身拥有代币数量）", t, func() {
		ethClient := config.GetEthClient()
		chainId, err := ethClient.ChainID(context.Background())
		_checkErr(err)
		// 备注: 对应的MetaMask钱包第一个账户私钥
		fromPrivateKey := config.FromPrivate


		fromCoinPrivateKey := config.GetPrivateList()[0]
		fromCoinPrivate, _ := crypto.HexToECDSA(fromCoinPrivateKey)
		fromCoinAddress := crypto.PubkeyToAddress(fromCoinPrivate.PublicKey)
		log.Printf("fromCoinAddress: %s", fromCoinAddress.String())

		// 给list0账户转账  确保有足够的可消耗ong
		txHash,err := deploy.SendTransfer(fromPrivateKey,fromCoinAddress.String(),big.NewInt(10),uint64(200000))
		_checkErr(err)
		txReceipt,err :=utils.GetTransferInfoByHash(txHash)
		_checkErr(err)
		if txReceipt.Status != 1{
			panic("tx error!")
		}

		erc20ContractAddress := getERC20ContractAddress()
		erc20Contract, err := ERC20Test.NewERC20Test(erc20ContractAddress, ethClient)
		_checkErr(err)
		callOpts := &bind.CallOpts{
			From: fromCoinAddress,
		}
		erp20Name, err := erc20Contract.Name(callOpts)
		log.Printf("erp20Name: %s", erp20Name)
		erp20Symbol, err := erc20Contract.Symbol(callOpts)
		log.Printf("erp20Symbol: %s", erp20Symbol)

		nonce, err := ethClient.PendingNonceAt(context.Background(), fromCoinAddress)
		_checkErr(err)
		opts, err := bind.NewKeyedTransactorWithChainID(fromCoinPrivate, chainId)
		_checkErr(err)
		opts.Nonce = big.NewInt(int64(nonce))
		opts.GasPrice = big.NewInt(500)
		opts.GasLimit = 3000000
		_to1 := config.GetAddressList()[1]
		//代币转账前
		fromBalance, err := erc20Contract.BalanceOf(callOpts, fromCoinAddress)
		_checkErr(err)
		toBalance, err := erc20Contract.BalanceOf(callOpts, _to1)
		_checkErr(err)

		amount := big.NewInt(500)
		fromBalanceOng := utils.GetBalance(fromCoinAddress.String())
		res, err := erc20Contract.Transfer(opts, _to1, amount)
		_checkErr(err)
		log.Printf("res.Hash(): %s", res.Hash().String())
		receipt, err := utils.GetTransferInfoByHash(res.Hash())
		_checkErr(err)
		log.Printf("receipt.Status: %d", receipt.Status)
		if receipt.Status != 1 {
			panic("Tx fail!")
		}
		useGas := big.NewInt(int64(receipt.CumulativeGasUsed))
		log.Printf("useGas: %d", useGas)
		fromBalanceOngAfter := utils.GetBalance(fromCoinAddress.String())
		log.Printf("fromBalanceOng: %d", fromBalanceOng)
		log.Printf("fromBalanceOngAfter: %d", fromBalanceOngAfter)

		// 断言ong消耗
		So(fromBalanceOng.Cmp(fromBalanceOngAfter.Add(fromBalanceOngAfter, useGas)) == 0, ShouldBeTrue)
		toBalanceAfter, err := erc20Contract.BalanceOf(callOpts, _to1)
		_checkErr(err)
		//log.Printf("toBalanceAfter: %d", toBalanceAfter)
		fromBalanceAfter, err := erc20Contract.BalanceOf(callOpts, fromCoinAddress)
		_checkErr(err)
		//log.Printf("fromBalanceAfter: %d", fromBalanceAfter)

		log.Printf("fromCoinBalance: %d",fromBalance)
		log.Printf("fromCoinBalanceAfter: %d",fromBalanceAfter)
		log.Printf("toCoinBalance: %d",toBalance)
		log.Printf("toCoinBalanceAfter: %d",toBalanceAfter)

		// 断言代币转账后from账户资金变动情况
		So(fromBalance.Cmp(fromBalanceAfter.Add(fromBalanceAfter, amount)) == 0, ShouldBeTrue)
		// 断言代币转账后to账户资金变动情况
		So(toBalanceAfter.Cmp(toBalance.Add(toBalance, amount)) == 0, ShouldBeTrue)
	})
	Convey("TestERC20_transfer_04 config.getAddressList[0]转账给config.getAddressList[1]转账（数量大于自身拥有代币数量）", t, func() {
		ethClient := config.GetEthClient()
		chainId, err := ethClient.ChainID(context.Background())
		_checkErr(err)
		// 备注: 对应的MetaMask钱包第一个账户私钥
		fromPrivateKey := config.FromPrivate


		fromCoinPrivateKey := config.GetPrivateList()[0]
		fromCoinPrivate, _ := crypto.HexToECDSA(fromCoinPrivateKey)
		fromCoinAddress := crypto.PubkeyToAddress(fromCoinPrivate.PublicKey)
		log.Printf("fromCoinAddress: %s", fromCoinAddress.String())

		// 给list0账户转账  确保有足够的可消耗ong
		txHash,err := deploy.SendTransfer(fromPrivateKey,fromCoinAddress.String(),big.NewInt(10),uint64(200000))
		_checkErr(err)
		txReceipt,err :=utils.GetTransferInfoByHash(txHash)
		_checkErr(err)
		if txReceipt.Status != 1{
			panic("tx error!")
		}

		erc20ContractAddress := getERC20ContractAddress()
		erc20Contract, err := ERC20Test.NewERC20Test(erc20ContractAddress, ethClient)
		_checkErr(err)
		callOpts := &bind.CallOpts{
			From: fromCoinAddress,
		}
		erp20Name, err := erc20Contract.Name(callOpts)
		log.Printf("erp20Name: %s", erp20Name)
		erp20Symbol, err := erc20Contract.Symbol(callOpts)
		log.Printf("erp20Symbol: %s", erp20Symbol)

		nonce, err := ethClient.PendingNonceAt(context.Background(), fromCoinAddress)
		_checkErr(err)
		opts, err := bind.NewKeyedTransactorWithChainID(fromCoinPrivate, chainId)
		_checkErr(err)
		opts.Nonce = big.NewInt(int64(nonce))
		opts.GasPrice = big.NewInt(500)
		opts.GasLimit = 3000000
		_to1 := config.GetAddressList()[1]
		//代币转账前
		fromBalance, err := erc20Contract.BalanceOf(callOpts, fromCoinAddress)
		_checkErr(err)
		toBalance, err := erc20Contract.BalanceOf(callOpts, _to1)
		_checkErr(err)

		amount := big.NewInt(5000)
		fromBalanceOng := utils.GetBalance(fromCoinAddress.String())
		res, err := erc20Contract.Transfer(opts, _to1, amount)
		_checkErr(err)
		log.Printf("res.Hash(): %s", res.Hash().String())
		receipt, err := utils.GetTransferInfoByHash(res.Hash())
		_checkErr(err)
		log.Printf("receipt.Status: %d", receipt.Status)

		//断言交易失败
		So(receipt.Status == 0,ShouldBeTrue)
		useGas := big.NewInt(int64(receipt.CumulativeGasUsed))
		log.Printf("useGas: %d", useGas)
		fromBalanceOngAfter := utils.GetBalance(fromCoinAddress.String())
		log.Printf("fromBalanceOng: %d", fromBalanceOng)
		log.Printf("fromBalanceOngAfter: %d", fromBalanceOngAfter)

		// 断言ong消耗
		So(fromBalanceOng.Cmp(fromBalanceOngAfter.Add(fromBalanceOngAfter, useGas)) == 0, ShouldBeTrue)
		toBalanceAfter, err := erc20Contract.BalanceOf(callOpts, _to1)
		_checkErr(err)
		//log.Printf("toBalanceAfter: %d", toBalanceAfter)
		fromBalanceAfter, err := erc20Contract.BalanceOf(callOpts, fromCoinAddress)
		_checkErr(err)
		//log.Printf("fromBalanceAfter: %d", fromBalanceAfter)

		log.Printf("fromCoinBalance: %d",fromBalance)
		log.Printf("fromCoinBalanceAfter: %d",fromBalanceAfter)
		log.Printf("toCoinBalance: %d",toBalance)
		log.Printf("toCoinBalanceAfter: %d",toBalanceAfter)

		// 断言代币转账后from账户资金变动情况
		So(fromBalance.Cmp(fromBalanceAfter) == 0, ShouldBeTrue)
		// 断言代币转账后to账户资金变动情况
		So(toBalanceAfter.Cmp(toBalance) == 0, ShouldBeTrue)
	})
	Convey("TestERC20_transfer_05 config.getAddressList[0]转账给config.getAddressList[1]转账（数量等于自身拥有代币数量）", t, func() {
		ethClient := config.GetEthClient()
		chainId, err := ethClient.ChainID(context.Background())
		_checkErr(err)
		// 备注: 对应的MetaMask钱包第一个账户私钥
		fromPrivateKey := config.FromPrivate

		fromCoinPrivateKey := config.GetPrivateList()[0]
		fromCoinPrivate, _ := crypto.HexToECDSA(fromCoinPrivateKey)
		fromCoinAddress := crypto.PubkeyToAddress(fromCoinPrivate.PublicKey)
		log.Printf("fromCoinAddress: %s", fromCoinAddress.String())

		// 给list0账户转账  确保有足够的可消耗ong
		txHash,err := deploy.SendTransfer(fromPrivateKey,fromCoinAddress.String(),big.NewInt(10),uint64(200000))
		_checkErr(err)
		txReceipt,err :=utils.GetTransferInfoByHash(txHash)
		_checkErr(err)
		if txReceipt.Status != 1{
			panic("tx error!")
		}

		erc20ContractAddress := getERC20ContractAddress()
		erc20Contract, err := ERC20Test.NewERC20Test(erc20ContractAddress, ethClient)
		_checkErr(err)
		callOpts := &bind.CallOpts{
			From: fromCoinAddress,
		}
		erp20Name, err := erc20Contract.Name(callOpts)
		log.Printf("erp20Name: %s", erp20Name)
		erp20Symbol, err := erc20Contract.Symbol(callOpts)
		log.Printf("erp20Symbol: %s", erp20Symbol)

		nonce, err := ethClient.PendingNonceAt(context.Background(), fromCoinAddress)
		_checkErr(err)
		opts, err := bind.NewKeyedTransactorWithChainID(fromCoinPrivate, chainId)
		_checkErr(err)
		opts.Nonce = big.NewInt(int64(nonce))
		opts.GasPrice = big.NewInt(500)
		opts.GasLimit = 3000000
		_to1 := config.GetAddressList()[1]
		//代币转账前
		fromBalance, err := erc20Contract.BalanceOf(callOpts, fromCoinAddress)
		_checkErr(err)
		toBalance, err := erc20Contract.BalanceOf(callOpts, _to1)
		_checkErr(err)

		fromBalanceOng := utils.GetBalance(fromCoinAddress.String())
		res, err := erc20Contract.Transfer(opts, _to1, fromBalance)
		_checkErr(err)
		log.Printf("res.Hash(): %s", res.Hash().String())
		receipt, err := utils.GetTransferInfoByHash(res.Hash())
		_checkErr(err)
		log.Printf("receipt.Status: %d", receipt.Status)
		if receipt.Status != 1 {
			panic("Tx fail!")
		}
		useGas := big.NewInt(int64(receipt.CumulativeGasUsed))
		log.Printf("useGas: %d", useGas)
		fromBalanceOngAfter := utils.GetBalance(fromCoinAddress.String())
		log.Printf("fromBalanceOng: %d", fromBalanceOng)
		log.Printf("fromBalanceOngAfter: %d", fromBalanceOngAfter)

		// 断言ong消耗
		So(fromBalanceOng.Cmp(fromBalanceOngAfter.Add(fromBalanceOngAfter, useGas)) == 0, ShouldBeTrue)
		toBalanceAfter, err := erc20Contract.BalanceOf(callOpts, _to1)
		_checkErr(err)
		//log.Printf("toBalanceAfter: %d", toBalanceAfter)
		fromBalanceAfter, err := erc20Contract.BalanceOf(callOpts, fromCoinAddress)
		_checkErr(err)
		//log.Printf("fromBalanceAfter: %d", fromBalanceAfter)

		log.Printf("fromCoinBalance: %d",fromBalance)
		log.Printf("fromCoinBalanceAfter: %d",fromBalanceAfter)
		log.Printf("toCoinBalance: %d",toBalance)
		log.Printf("toCoinBalanceAfter: %d",toBalanceAfter)

		// 断言代币转账后from账户资金变动情况
		So(fromBalance.Cmp(fromBalanceAfter.Add(fromBalanceAfter, fromBalance)) == 0, ShouldBeTrue)
		// 断言代币转账后to账户资金变动情况
		So(toBalanceAfter.Cmp(toBalance.Add(toBalance, fromBalance)) == 0, ShouldBeTrue)
	})
}

func initERC20ContractAddress() common.Address {

	// 备注: 对应的MetaMask钱包第一个账户私钥
	fromPrivateKey := config.FromPrivate
	fromPrivate, _ := crypto.HexToECDSA(fromPrivateKey)
	fromAddress := crypto.PubkeyToAddress(fromPrivate.PublicKey)
	log.Printf("fromAddress: %s", fromAddress.String())
	erc20Abi, err := utils.ReadAll("deploy/testEPC20/ERC20_example_sol_TokenERC20.abi")
	_checkErr(err)
	erc20Bin, err := utils.ReadAll("deploy/testEPC20/ERC20_example_sol_TokenERC20.bin")
	_checkErr(err)
	erc20ContractHash := deploy.ContractDeploy(string(erc20Bin), string(erc20Abi), big.NewInt(10000000), "牛币", "NBB")
	log.Printf("erc20ContractHash: %s", erc20ContractHash.String())

	erc20ContractReceipt, err := utils.GetTransferInfoByHash(erc20ContractHash)
	_checkErr(err)
	if erc20ContractReceipt.Status != 1 {
		fmt.Println("erc20ContractReceipt.Status : ", erc20ContractReceipt.Status)
		panic("TxTransaction fail !!!")
	}
	//erc20ContractAddress合约地址
	log.Printf("erc20ContractAddress: %s", erc20ContractReceipt.ContractAddress.String())
	return erc20ContractReceipt.ContractAddress
}

//erc20合约地址
var ERC20Address = common.Address{}

func getERC20ContractAddress() common.Address {
	if ERC20Address.String() != "0x0000000000000000000000000000000000000000" {
		return ERC20Address
	}
	ERC20Address = initERC20ContractAddress()
	return ERC20Address
}

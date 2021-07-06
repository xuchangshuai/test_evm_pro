package testcase

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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

func TestERC20_approve(t *testing.T) {
	Convey("TestERC20_approve_01 使用合约拥有者授权给账户A可使用代币数量5000", t, func() {
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

		amount := big.NewInt(5000)
		fromBalanceOng := utils.GetBalance(fromAddress.String())
		log.Printf("fromBalanceOng: %d", fromBalanceOng)
		res, err := erc20Contract.Approve(opts, _to1, amount)
		_checkErr(err)
		receipt, err := utils.GetTransferInfoByHash(res.Hash())
		fmt.Println(receipt)
		_checkErr(err)
		if receipt.Status != 1 {
			panic("Tx fail!")
		}

		approveAmount, err := erc20Contract.Allowance(callOpts, fromAddress, _to1)
		_checkErr(err)
		log.Printf("approveAmount: %d", approveAmount)
		log.Printf("amount: %d", amount)

		So(approveAmount.Cmp(amount) == 0, ShouldBeTrue)
		useGas := big.NewInt(int64(receipt.CumulativeGasUsed))
		log.Printf("useGas: %d", useGas)
		fromBalanceOngAfter := utils.GetBalance(fromAddress.String())

		// 断言ong消耗
		So(fromBalanceOng.Cmp(fromBalanceOngAfter.Add(fromBalanceOngAfter, useGas)) == 0, ShouldBeTrue)
	})
	Convey("TestERC20_approve_02 使用合约拥有者授权给账户A可使用代币数量0", t, func() {
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

		amount := big.NewInt(0)
		fromBalanceOng := utils.GetBalance(fromAddress.String())
		log.Printf("fromBalanceOng: %d", fromBalanceOng)
		res, err := erc20Contract.Approve(opts, _to1, amount)
		_checkErr(err)
		receipt, err := utils.GetTransferInfoByHash(res.Hash())
		fmt.Println(receipt)
		_checkErr(err)
		if receipt.Status != 1 {
			panic("Tx fail!")
		}

		approveAmount, err := erc20Contract.Allowance(callOpts, fromAddress, _to1)
		_checkErr(err)
		log.Printf("approveAmount: %d", approveAmount)
		log.Printf("amount: %d", amount)

		So(approveAmount.Cmp(amount) == 0, ShouldBeTrue)
		useGas := big.NewInt(int64(receipt.CumulativeGasUsed))
		log.Printf("useGas: %d", useGas)
		fromBalanceOngAfter := utils.GetBalance(fromAddress.String())

		// 断言ong消耗
		So(fromBalanceOng.Cmp(fromBalanceOngAfter.Add(fromBalanceOngAfter, useGas)) == 0, ShouldBeTrue)
	})
	Convey("TestERC20_approve_03 使用合约拥有者授权给账户A可使用代币数量超过拥有数量", t, func() {
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

		amount, err := erc20Contract.BalanceOf(callOpts, fromAddress)
		_checkErr(err)
		amount.Add(amount, big.NewInt(1))
		fromBalanceOng := utils.GetBalance(fromAddress.String())
		log.Printf("fromBalanceOng: %d", fromBalanceOng)
		res, err := erc20Contract.Approve(opts, _to1, amount)
		_checkErr(err)
		receipt, err := utils.GetTransferInfoByHash(res.Hash())
		fmt.Println(receipt)
		_checkErr(err)
		if receipt.Status != 1 {
			panic("Tx fail!")
		}

		approveAmount, err := erc20Contract.Allowance(callOpts, fromAddress, _to1)
		_checkErr(err)
		log.Printf("approveAmount: %d", approveAmount)
		log.Printf("amount: %d", amount)

		//合约未作该限制 故断言成功
		So(approveAmount.Cmp(amount) == 0, ShouldBeTrue)
		useGas := big.NewInt(int64(receipt.CumulativeGasUsed))
		log.Printf("useGas: %d", useGas)
		fromBalanceOngAfter := utils.GetBalance(fromAddress.String())

		// 断言ong消耗
		So(fromBalanceOng.Cmp(fromBalanceOngAfter.Add(fromBalanceOngAfter, useGas)) == 0, ShouldBeTrue)
	})
	Convey("TestERC20_approve_04 使用账户A授权给账户B可使用代币数量5000", t, func() {
		ethClient := config.GetEthClient()
		chainId, err := ethClient.ChainID(context.Background())
		_checkErr(err)

		// 备注: 对应的MetaMask钱包第一个账户私钥
		fromPrivateKey := config.GetPrivateList()[0]
		fromPrivate, _ := crypto.HexToECDSA(fromPrivateKey)
		fromAddress := crypto.PubkeyToAddress(fromPrivate.PublicKey)
		checkFromAddressBalance(fromAddress.String())
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

		amount := big.NewInt(5000)
		fromBalanceOng := utils.GetBalance(fromAddress.String())
		log.Printf("fromBalanceOng: %d", fromBalanceOng)
		res, err := erc20Contract.Approve(opts, _to1, amount)
		_checkErr(err)
		receipt, err := utils.GetTransferInfoByHash(res.Hash())
		fmt.Println(receipt)
		_checkErr(err)
		if receipt.Status != 1 {
			panic("Tx fail!")
		}

		approveAmount, err := erc20Contract.Allowance(callOpts, fromAddress, _to1)
		_checkErr(err)
		log.Printf("approveAmount: %d", approveAmount)
		log.Printf("amount: %d", amount)

		//合约未作该限制 故断言成功
		So(approveAmount.Cmp(amount) == 0, ShouldBeTrue)
		useGas := big.NewInt(int64(receipt.CumulativeGasUsed))
		log.Printf("useGas: %d", useGas)
		fromBalanceOngAfter := utils.GetBalance(fromAddress.String())

		// 断言ong消耗
		So(fromBalanceOng.Cmp(fromBalanceOngAfter.Add(fromBalanceOngAfter, useGas)) == 0, ShouldBeTrue)
	})
	Convey("TestERC20_approve_05 使用账户A授权给账户B可使用代币数量0", t, func() {
		ethClient := config.GetEthClient()
		chainId, err := ethClient.ChainID(context.Background())
		_checkErr(err)

		// 备注: 对应的MetaMask钱包第一个账户私钥
		fromPrivateKey := config.GetPrivateList()[0]
		fromPrivate, _ := crypto.HexToECDSA(fromPrivateKey)
		fromAddress := crypto.PubkeyToAddress(fromPrivate.PublicKey)
		checkFromAddressBalance(fromAddress.String())
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

		amount := big.NewInt(0)
		fromBalanceOng := utils.GetBalance(fromAddress.String())
		log.Printf("fromBalanceOng: %d", fromBalanceOng)
		res, err := erc20Contract.Approve(opts, _to1, amount)
		_checkErr(err)
		receipt, err := utils.GetTransferInfoByHash(res.Hash())
		fmt.Println(receipt)
		_checkErr(err)
		if receipt.Status != 1 {
			panic("Tx fail!")
		}

		approveAmount, err := erc20Contract.Allowance(callOpts, fromAddress, _to1)
		_checkErr(err)
		log.Printf("approveAmount: %d", approveAmount)
		log.Printf("amount: %d", amount)

		//合约未作该限制 故断言成功
		So(approveAmount.Cmp(amount) == 0, ShouldBeTrue)
		useGas := big.NewInt(int64(receipt.CumulativeGasUsed))
		log.Printf("useGas: %d", useGas)
		fromBalanceOngAfter := utils.GetBalance(fromAddress.String())

		// 断言ong消耗
		So(fromBalanceOng.Cmp(fromBalanceOngAfter.Add(fromBalanceOngAfter, useGas)) == 0, ShouldBeTrue)
	})
	Convey("TestERC20_approve_06 使用账户A给账户B可使用代币数量超过拥有数量", t, func() {
		ethClient := config.GetEthClient()
		chainId, err := ethClient.ChainID(context.Background())
		_checkErr(err)

		// 备注: 对应的MetaMask钱包第一个账户私钥
		fromPrivateKey := config.GetPrivateList()[0]
		fromPrivate, _ := crypto.HexToECDSA(fromPrivateKey)
		fromAddress := crypto.PubkeyToAddress(fromPrivate.PublicKey)
		checkFromAddressBalance(fromAddress.String())
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

		amount, err := erc20Contract.BalanceOf(callOpts, fromAddress)
		_checkErr(err)
		amount.Add(amount, big.NewInt(1))
		fromBalanceOng := utils.GetBalance(fromAddress.String())
		log.Printf("fromBalanceOng: %d", fromBalanceOng)
		res, err := erc20Contract.Approve(opts, _to1, amount)
		_checkErr(err)
		receipt, err := utils.GetTransferInfoByHash(res.Hash())
		fmt.Println(receipt)
		_checkErr(err)
		if receipt.Status != 1 {
			panic("Tx fail!")
		}

		approveAmount, err := erc20Contract.Allowance(callOpts, fromAddress, _to1)
		_checkErr(err)
		log.Printf("approveAmount: %d", approveAmount)
		log.Printf("amount: %d", amount)

		//合约未作该限制 故断言成功
		So(approveAmount.Cmp(amount) == 0, ShouldBeTrue)
		useGas := big.NewInt(int64(receipt.CumulativeGasUsed))
		log.Printf("useGas: %d", useGas)
		fromBalanceOngAfter := utils.GetBalance(fromAddress.String())

		// 断言ong消耗
		So(fromBalanceOng.Cmp(fromBalanceOngAfter.Add(fromBalanceOngAfter, useGas)) == 0, ShouldBeTrue)
	})
}

func checkFromAddressBalance(fromAddress string) {
	if utils.GetBalance(fromAddress).Cmp(big.NewInt(int64(100000000000))) < 0 {
		txHash, err := deploy.SendTransfer(config.FromPrivate, fromAddress, big.NewInt(100000000000), uint64(200000))
		_checkErr(err)
		re, err := utils.GetTransferInfoByHash(txHash)
		_checkErr(err)
		if re.Status != 1 {
			panic("error ! formAddress balance low")
		}
	}
}

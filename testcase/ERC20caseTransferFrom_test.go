package testcase

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/smartystreets/goconvey/convey"
	"log"
	"math/big"
	"test_evm/config"
	ERC20Test "test_evm/testcase/deploy/testEPC20"
	"test_evm/utils"
	"testing"
)

func TestERC20_transferFrom(t *testing.T) {
	Convey("TestERC20_transfer_01 1.合约拥有者账户approve授权给账户A n 代币 2.账户A调用合约该方法转账给B账户<n代币", t, func() {
		ethClient := config.GetEthClient()
		chainId, err := ethClient.ChainID(context.Background())
		_checkErr(err)

		// 备注: 对应的MetaMask钱包第一个账户私钥
		fromPrivateKey := config.FromPrivate
		fromPrivate, _ := crypto.HexToECDSA(fromPrivateKey)
		fromAddress := crypto.PubkeyToAddress(fromPrivate.PublicKey)

		fromCoinPrivateKey := config.GetPrivateList()[2]
		fromCoinPrivate, _ := crypto.HexToECDSA(fromCoinPrivateKey)
		fromCoinAddress := crypto.PubkeyToAddress(fromCoinPrivate.PublicKey)

		checkFromAmount(fromCoinAddress)

		erc20ContractAddress := getERC20ContractAddress()
		erc20Contract, err := ERC20Test.NewERC20Test(erc20ContractAddress, ethClient)
		_checkErr(err)
		callOpts := &bind.CallOpts{
			From: fromCoinAddress,
		}
		_to1 := fromCoinAddress
		checkFromAddressBalance(_to1.String())
		amountCoin := big.NewInt(500)
		_to2 := config.GetAddressList()[3]

		optsCoin, err := bind.NewKeyedTransactorWithChainID(fromCoinPrivate, chainId)
		_checkErr(err)
		coinNonce, err := ethClient.PendingNonceAt(context.Background(), fromCoinAddress)
		_checkErr(err)
		optsCoin.Nonce = big.NewInt(int64(coinNonce))
		optsCoin.GasPrice = big.NewInt(500000000000)
		optsCoin.GasLimit = 3000000

		fromAddressBalance, err := erc20Contract.BalanceOf(callOpts, fromAddress)
		_checkErr(err)
		_to1Balance, err := erc20Contract.BalanceOf(callOpts, _to1)
		_checkErr(err)
		_to2Balance, err := erc20Contract.BalanceOf(callOpts, _to2)
		_checkErr(err)

		transferFromHash, err := erc20Contract.TransferFrom(optsCoin, fromAddress, _to2, amountCoin)
		_checkErr(err)
		txReceipt, err := utils.GetTransferInfoByHash(transferFromHash.Hash())
		_checkErr(err)
		if txReceipt.Status != 1 {
			panic("tx error!")
		}

		fromAddressBalanceAfter, err := erc20Contract.BalanceOf(callOpts, fromAddress)
		_checkErr(err)
		_to1BalanceAfter, err := erc20Contract.BalanceOf(callOpts, _to1)
		_checkErr(err)
		_to2BalanceAfter, err := erc20Contract.BalanceOf(callOpts, _to2)
		_checkErr(err)

		log.Printf("fromAddressBalance: %d", fromAddressBalance)
		log.Printf("fromAddressBalanceAfter: %d", fromAddressBalanceAfter)
		log.Printf("_to1Balance: %d", _to1Balance)
		log.Printf("_to1BalanceAfter: %d", _to1BalanceAfter)
		log.Printf("_to2Balance: %d", _to2Balance)
		log.Printf("_to2BalanceAfter: %d", _to2BalanceAfter)
		So(_to1Balance.Cmp(_to1BalanceAfter) == 0, ShouldBeTrue)
		So(fromAddressBalance.Cmp(fromAddressBalanceAfter.Add(fromAddressBalanceAfter, amountCoin)) == 0, ShouldBeTrue)
		So(_to2BalanceAfter.Cmp(_to2Balance.Add(_to2Balance, amountCoin)) == 0, ShouldBeTrue)

	})
	Convey("TestERC20_transfer_02 1.合约拥有者账户approve授权给账户A n 代币 2.账户A调用合约该方法转账给B账户=n代币", t, func() {
		ethClient := config.GetEthClient()
		chainId, err := ethClient.ChainID(context.Background())
		_checkErr(err)

		// 备注: 对应的MetaMask钱包第一个账户私钥
		fromPrivateKey := config.FromPrivate
		fromPrivate, _ := crypto.HexToECDSA(fromPrivateKey)
		fromAddress := crypto.PubkeyToAddress(fromPrivate.PublicKey)

		fromCoinPrivateKey := config.GetPrivateList()[2]
		fromCoinPrivate, _ := crypto.HexToECDSA(fromCoinPrivateKey)
		fromCoinAddress := crypto.PubkeyToAddress(fromCoinPrivate.PublicKey)

		checkFromAmount(fromCoinAddress)

		erc20ContractAddress := getERC20ContractAddress()
		erc20Contract, err := ERC20Test.NewERC20Test(erc20ContractAddress, ethClient)
		_checkErr(err)
		callOpts := &bind.CallOpts{
			From: fromCoinAddress,
		}
		_to2 := config.GetAddressList()[3]
		amountCoin, err := erc20Contract.Allowance(callOpts, fromAddress, fromCoinAddress)
		log.Printf("amountCoin: %d", amountCoin)
		_checkErr(err)
		optsCoin, err := bind.NewKeyedTransactorWithChainID(fromCoinPrivate, chainId)
		_checkErr(err)
		coinNonce, err := ethClient.PendingNonceAt(context.Background(), fromCoinAddress)
		_checkErr(err)
		optsCoin.Nonce = big.NewInt(int64(coinNonce))
		optsCoin.GasPrice = big.NewInt(500000000000)
		optsCoin.GasLimit = 3000000

		fromAddressBalance, err := erc20Contract.BalanceOf(callOpts, fromAddress)
		_checkErr(err)
		_to1Balance, err := erc20Contract.BalanceOf(callOpts, fromCoinAddress)
		_checkErr(err)
		_to2Balance, err := erc20Contract.BalanceOf(callOpts, _to2)
		_checkErr(err)

		transferFromHash, err := erc20Contract.TransferFrom(optsCoin, fromAddress, _to2, amountCoin)
		_checkErr(err)
		txReceipt, err := utils.GetTransferInfoByHash(transferFromHash.Hash())
		_checkErr(err)
		if txReceipt.Status != 1 {
			panic("tx error!")
		}

		fromAddressBalanceAfter, err := erc20Contract.BalanceOf(callOpts, fromAddress)
		_checkErr(err)
		_to1BalanceAfter, err := erc20Contract.BalanceOf(callOpts, fromCoinAddress)
		_checkErr(err)
		_to2BalanceAfter, err := erc20Contract.BalanceOf(callOpts, _to2)
		_checkErr(err)

		log.Printf("fromAddressBalance: %d", fromAddressBalance)
		log.Printf("fromAddressBalanceAfter: %d", fromAddressBalanceAfter)
		log.Printf("_to1Balance: %d", _to1Balance)
		log.Printf("_to1BalanceAfter: %d", _to1BalanceAfter)
		log.Printf("_to2Balance: %d", _to2Balance)
		log.Printf("_to2BalanceAfter: %d", _to2BalanceAfter)
		So(_to1Balance.Cmp(_to1BalanceAfter) == 0, ShouldBeTrue)
		So(fromAddressBalance.Cmp(fromAddressBalanceAfter.Add(fromAddressBalanceAfter, amountCoin)) == 0, ShouldBeTrue)
		So(_to2BalanceAfter.Cmp(_to2Balance.Add(_to2Balance, amountCoin)) == 0, ShouldBeTrue)

	})
	Convey("TestERC20_transfer_03 1.合约拥有者账户approve授权给账户A n 代币 2.账户A调用合约该方法转账给B账户>n代币", t, func() {
		ethClient := config.GetEthClient()
		chainId, err := ethClient.ChainID(context.Background())
		_checkErr(err)

		// 备注: 对应的MetaMask钱包第一个账户私钥
		fromPrivateKey := config.FromPrivate
		fromPrivate, _ := crypto.HexToECDSA(fromPrivateKey)
		fromAddress := crypto.PubkeyToAddress(fromPrivate.PublicKey)

		fromCoinPrivateKey := config.GetPrivateList()[2]
		fromCoinPrivate, _ := crypto.HexToECDSA(fromCoinPrivateKey)
		fromCoinAddress := crypto.PubkeyToAddress(fromCoinPrivate.PublicKey)

		checkFromAmount(fromCoinAddress)

		erc20ContractAddress := getERC20ContractAddress()
		erc20Contract, err := ERC20Test.NewERC20Test(erc20ContractAddress, ethClient)
		_checkErr(err)
		callOpts := &bind.CallOpts{

			From: fromCoinAddress,
		}
		_to2 := config.GetAddressList()[3]
		amountCoin, err := erc20Contract.Allowance(callOpts, fromAddress, fromCoinAddress)
		_checkErr(err)
		amountCoin = amountCoin.Add(amountCoin, big.NewInt(1))

		optsCoin, err := bind.NewKeyedTransactorWithChainID(fromCoinPrivate, chainId)
		_checkErr(err)
		coinNonce, err := ethClient.PendingNonceAt(context.Background(), fromCoinAddress)
		_checkErr(err)
		optsCoin.Nonce = big.NewInt(int64(coinNonce))
		optsCoin.GasPrice = big.NewInt(500000000000)
		optsCoin.GasLimit = 3000000

		fromAddressBalance, err := erc20Contract.BalanceOf(callOpts, fromAddress)
		_checkErr(err)
		_to1Balance, err := erc20Contract.BalanceOf(callOpts, fromCoinAddress)
		_checkErr(err)
		_to2Balance, err := erc20Contract.BalanceOf(callOpts, _to2)
		_checkErr(err)

		transferFromHash, err := erc20Contract.TransferFrom(optsCoin, fromAddress, _to2, amountCoin)
		_checkErr(err)
		txReceipt, err := utils.GetTransferInfoByHash(transferFromHash.Hash())
		_checkErr(err)
		So(txReceipt.Status == 0, ShouldBeTrue)
		fromAddressBalanceAfter, err := erc20Contract.BalanceOf(callOpts, fromAddress)
		_checkErr(err)
		_to1BalanceAfter, err := erc20Contract.BalanceOf(callOpts, fromCoinAddress)
		_checkErr(err)
		_to2BalanceAfter, err := erc20Contract.BalanceOf(callOpts, _to2)
		_checkErr(err)

		log.Printf("fromAddressBalance: %d", fromAddressBalance)
		log.Printf("fromAddressBalanceAfter: %d", fromAddressBalanceAfter)
		log.Printf("_to1Balance: %d", _to1Balance)
		log.Printf("_to1BalanceAfter: %d", _to1BalanceAfter)
		log.Printf("_to2Balance: %d", _to2Balance)
		log.Printf("_to2BalanceAfter: %d", _to2BalanceAfter)
		So(_to1Balance.Cmp(_to1BalanceAfter) == 0, ShouldBeTrue)
		So(fromAddressBalance.Cmp(fromAddressBalanceAfter) == 0, ShouldBeTrue)
		So(_to2BalanceAfter.Cmp(_to2Balance) == 0, ShouldBeTrue)

	})
	Convey("TestERC20_transfer_04 1.合约拥有者账户approve授权给账户A n 代币 2.账户A调用合约该方法转账给B账户0代币", t, func() {
		ethClient := config.GetEthClient()
		chainId, err := ethClient.ChainID(context.Background())
		_checkErr(err)

		// 备注: 对应的MetaMask钱包第一个账户私钥
		fromPrivateKey := config.FromPrivate
		fromPrivate, _ := crypto.HexToECDSA(fromPrivateKey)
		fromAddress := crypto.PubkeyToAddress(fromPrivate.PublicKey)

		fromCoinPrivateKey := config.GetPrivateList()[2]
		fromCoinPrivate, _ := crypto.HexToECDSA(fromCoinPrivateKey)
		fromCoinAddress := crypto.PubkeyToAddress(fromCoinPrivate.PublicKey)

		checkFromAmount(fromCoinAddress)

		erc20ContractAddress := getERC20ContractAddress()
		erc20Contract, err := ERC20Test.NewERC20Test(erc20ContractAddress, ethClient)
		_checkErr(err)
		callOpts := &bind.CallOpts{
			From: fromCoinAddress,
		}
		_to2 := config.GetAddressList()[3]
		amountCoin := big.NewInt(0)
		optsCoin, err := bind.NewKeyedTransactorWithChainID(fromCoinPrivate, chainId)
		_checkErr(err)
		coinNonce, err := ethClient.PendingNonceAt(context.Background(), fromCoinAddress)
		_checkErr(err)
		optsCoin.Nonce = big.NewInt(int64(coinNonce))
		optsCoin.GasPrice = big.NewInt(500000000000)
		optsCoin.GasLimit = 3000000

		fromAddressBalance, err := erc20Contract.BalanceOf(callOpts, fromAddress)
		_checkErr(err)
		_to1Balance, err := erc20Contract.BalanceOf(callOpts, fromCoinAddress)
		_checkErr(err)
		_to2Balance, err := erc20Contract.BalanceOf(callOpts, _to2)
		_checkErr(err)

		transferFromHash, err := erc20Contract.TransferFrom(optsCoin, fromAddress, _to2, amountCoin)
		_checkErr(err)
		txReceipt, err := utils.GetTransferInfoByHash(transferFromHash.Hash())
		_checkErr(err)
		So(txReceipt.Status == 0, ShouldBeTrue)
		fromAddressBalanceAfter, err := erc20Contract.BalanceOf(callOpts, fromAddress)
		_checkErr(err)
		_to1BalanceAfter, err := erc20Contract.BalanceOf(callOpts, fromCoinAddress)
		_checkErr(err)
		_to2BalanceAfter, err := erc20Contract.BalanceOf(callOpts, _to2)
		_checkErr(err)

		log.Printf("fromAddressBalance: %d", fromAddressBalance)
		log.Printf("fromAddressBalanceAfter: %d", fromAddressBalanceAfter)
		log.Printf("_to1Balance: %d", _to1Balance)
		log.Printf("_to1BalanceAfter: %d", _to1BalanceAfter)
		log.Printf("_to2Balance: %d", _to2Balance)
		log.Printf("_to2BalanceAfter: %d", _to2BalanceAfter)
		So(_to1Balance.Cmp(_to1BalanceAfter) == 0, ShouldBeTrue)
		So(fromAddressBalance.Cmp(fromAddressBalanceAfter) == 0, ShouldBeTrue)
		So(_to2BalanceAfter.Cmp(_to2Balance) == 0, ShouldBeTrue)

	})
}

func checkFromAmount(_to common.Address) {
	ethClient := config.GetEthClient()
	chainId, err := ethClient.ChainID(context.Background())
	_checkErr(err)

	// 备注: 对应的MetaMask钱包第一个账户私钥
	fromPrivateKey := config.FromPrivate
	fromPrivate, _ := crypto.HexToECDSA(fromPrivateKey)
	fromAddress := crypto.PubkeyToAddress(fromPrivate.PublicKey)

	erc20ContractAddress := getERC20ContractAddress()
	erc20Contract, err := ERC20Test.NewERC20Test(erc20ContractAddress, ethClient)
	_checkErr(err)
	callOpts := &bind.CallOpts{
		From: fromAddress,
	}

	nonce, err := ethClient.PendingNonceAt(context.Background(), fromAddress)
	_checkErr(err)
	opts, err := bind.NewKeyedTransactorWithChainID(fromPrivate, chainId)
	_checkErr(err)
	opts.Nonce = big.NewInt(int64(nonce))
	opts.GasPrice = big.NewInt(500000000000)
	opts.GasLimit = 3000000

	amount, err := erc20Contract.Allowance(callOpts, fromAddress, _to)
	log.Printf("checkFromAmount: %d", amount)
	_checkErr(err)
	if amount.Cmp(big.NewInt(5000)) < 0 {
		res, err := erc20Contract.Approve(opts, _to, big.NewInt(5000))
		_checkErr(err)
		receipt, err := utils.GetTransferInfoByHash(res.Hash())
		_checkErr(err)
		if receipt.Status != 1 {
			panic("Tx fail!")
		}
	}
	amount, err = erc20Contract.Allowance(callOpts, fromAddress, _to)
	_checkErr(err)
	log.Printf("checkFromAmount retrun: %d", amount)
}

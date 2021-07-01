package testcase

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	. "github.com/smartystreets/goconvey/convey"
	"math/big"
	"test_evm/testcase/deploy"
	"test_evm/utils"
	"testing"
	"time"
)

func TestEVMTransfer006(t *testing.T) {
	Convey("Test 前置发一笔正常交易", t, func() {
		url := "http://127.0.0.1:20339"
		toAddress := "0x5f9a233682e80e8bb1bfdf18cd21d82af2fd7edd"
		amount := big.NewInt(1000000000)
		//  备注
		fromPrivateKey := "9b4600df58ea9f11c35b82ff660c819255ed5f71355eeef78bc994ccada4d862"
		fromPrivate, _ := crypto.HexToECDSA(fromPrivateKey)
		fromAddress := crypto.PubkeyToAddress(fromPrivate.PublicKey)
		fmt.Println("fromAddress: ", fromAddress)

		//交易前账户余额
		fromBalance := utils.GetBalance(url, fromAddress.String())
		toBalance := utils.GetBalance(url, toAddress)
		managerBalance := utils.GetBalance(url, "0x0000000000000000000000000000000000000007")

		hash, err := deploy.SendTransfer(url, fromPrivateKey, toAddress, amount, uint64(200000))
		if err != nil {
			panic(err)
		}
		//fmt.Println(hash)
		receipt := make(map[string]interface{})
		for i := 0; i < 10; i++ { //设置10次请求，每次间隔3s 30s超时
			time.Sleep(3000000000)
			if len(receipt) != 0 {
				break
			} else {
				receipt = utils.GetTransferInfo(url, hash)
			}
		}
		if len(receipt) == 0 {
			panic("交易30s未落账!")
		}
		//fmt.Println(receipt)
		status := utils.Strval(receipt["status"])
		//fmt.Println(status)
		if status != "0x1" {
			fmt.Println("status : ", status)
			panic("TxTransaction fail !!!")
		}
		useGas := utils.HexToDec(utils.Strval(receipt["cumulativeGasUsed"]))
		fmt.Println("useGas", useGas)
		fromBalanceAfter := utils.GetBalance(url, fromAddress.String())
		toBalanceAfter := utils.GetBalance(url, toAddress)
		managerBalanceAfter := utils.GetBalance(url, "0x0000000000000000000000000000000000000007")
		fmt.Println("fromBalance", fromBalance)
		fmt.Println("toBalance", toBalance)
		fmt.Println("managerBalance", managerBalance)
		fmt.Println("fromBalanceAfter", fromBalanceAfter)
		fmt.Println("toBalanceAfter", toBalanceAfter)
		fmt.Println("managerBalanceAfter", managerBalanceAfter)
		//1 断言from账户资金是否 = fromBalance1 = fromBalance - gasUsed - amount
		fromBalance1 := fromBalanceAfter.Add(fromBalanceAfter, amount).Add(fromBalanceAfter, useGas)
		So(fromBalance.Cmp(fromBalance1) == 0, ShouldBeTrue)
		//2 断言to账户资金是否 = toBalance + amount
		So(toBalanceAfter.Cmp(toBalance.Add(toBalance, amount)) == 0, ShouldBeTrue)
		//3 断言治理合约地址是否正常获得消耗gas
		So(managerBalanceAfter.Cmp(managerBalance.Add(managerBalance, useGas)) == 0, ShouldBeTrue)
	})
	Convey("Test006 测试EIP155交易普通转账ong时 nonce值错误(nonce-1) ", t, func() {
		url := "http://127.0.0.1:20339"
		toAddress := "0x5f9a233682e80e8bb1bfdf18cd21d82af2fd7edd"
		amount := big.NewInt(500000000)
		gasLimit := uint64(200000)
		//  备注： 对应的MetaMask钱包第5个账户私钥
		fromPrivateKey := "9b4600df58ea9f11c35b82ff660c819255ed5f71355eeef78bc994ccada4d862"
		fromPrivate, _ := crypto.HexToECDSA(fromPrivateKey)
		fromAddress := crypto.PubkeyToAddress(fromPrivate.PublicKey)

		//交易前账户余额
		fromBalance := utils.GetBalance(url, fromAddress.String())
		managerBalance := utils.GetBalance(url, "0x0000000000000000000000000000000000000007")

		// 构建一个私链对象
		ethClient, err := ethclient.Dial(url)
		_checkErr(err)
		defer ethClient.Close()
		//查chainId
		chainId, err := ethClient.ChainID(context.Background())
		_checkErr(err)
		// 查nonce值
		nonce, err := ethClient.PendingNonceAt(context.Background(), fromAddress)
		_checkErr(err)
		_to := common.HexToAddress(toAddress)
		gasPrice := big.NewInt(500)
		rawTx := types.NewTransaction(nonce-1, _to, amount, gasLimit, gasPrice, nil)

		signer := types.NewEIP155Signer(chainId) //big.NewInt(1)//当前入参链id
		sigTransaction, err := types.SignTx(rawTx, signer, fromPrivate)
		_checkErr(err)
		err = ethClient.SendTransaction(context.Background(), sigTransaction)
		fmt.Println(err)
		// 断言nonce = nonce-1 构造的交易对象,发送交易抛异常
		So(err, ShouldNotBeNil)
		managerBalanceAfter := utils.GetBalance(url, "0x0000000000000000000000000000000000000007")
		fromBalanceAfter := utils.GetBalance(url, fromAddress.String())
		fmt.Println("fromBalance: ", fromBalance)
		fmt.Println("managerBalance: ", managerBalance)
		fmt.Println("managerBalanceAfter: ", managerBalanceAfter)
		fmt.Println("fromBalanceAfter: ", fromBalanceAfter)
		So(fromBalanceAfter.Cmp(fromBalance) == 0, ShouldBeTrue)
		So(managerBalanceAfter.Cmp(managerBalance) == 0, ShouldBeTrue)

	})
	Convey("Test007 测试EIP155交易普通转账ong时 nonce值错误(nonce+1) ", t, func() {
		url := "http://127.0.0.1:20339"
		toAddress := "0x5f9a233682e80e8bb1bfdf18cd21d82af2fd7edd"
		amount := big.NewInt(500000000)
		gasLimit := uint64(200000)
		//  备注： 对应的MetaMask钱包第5个账户私钥
		fromPrivateKey := "9b4600df58ea9f11c35b82ff660c819255ed5f71355eeef78bc994ccada4d862"
		fromPrivate, _ := crypto.HexToECDSA(fromPrivateKey)
		fromAddress := crypto.PubkeyToAddress(fromPrivate.PublicKey)

		//交易前账户余额
		fromBalance := utils.GetBalance(url, fromAddress.String())
		toBalance := utils.GetBalance(url, toAddress)
		managerBalance := utils.GetBalance(url, "0x0000000000000000000000000000000000000007")

		// 构建一个私链对象
		ethClient, err := ethclient.Dial(url)
		_checkErr(err)
		defer ethClient.Close()
		//查chainId
		chainId, err := ethClient.ChainID(context.Background())
		_checkErr(err)
		// 查nonce值
		nonce, err := ethClient.PendingNonceAt(context.Background(), fromAddress)
		_checkErr(err)
		_to := common.HexToAddress(toAddress)
		gasPrice := big.NewInt(500)
		rawTx := types.NewTransaction(nonce+1, _to, amount, gasLimit, gasPrice, nil)

		signer := types.NewEIP155Signer(chainId) //big.NewInt(1)//当前入参链id
		sigTransaction, err := types.SignTx(rawTx, signer, fromPrivate)
		_checkErr(err)
		err = ethClient.SendTransaction(context.Background(), sigTransaction)
		_checkErr(err)
		receipt := make(map[string]interface{})
		for i := 0; i < 10; i++ { //设置10次请求，每次间隔3s 30s超时
			time.Sleep(3000000000)
			if len(receipt) != 0 {
				break
			} else {
				receipt = utils.GetTransferInfo(url, sigTransaction.Hash())
			}
		}
		// nonce = nonce + 1 断言receipt查不到
		So(len(receipt) == 0, ShouldBeTrue)

		managerBalanceAfter := utils.GetBalance(url, "0x0000000000000000000000000000000000000007")
		fromBalanceAfter := utils.GetBalance(url, fromAddress.String())
		toBalanceAfter := utils.GetBalance(url, toAddress)
		fmt.Println("fromBalance: ", fromBalance)
		fmt.Println("fromBalanceAfter: ", fromBalanceAfter)
		fmt.Println("toBalance: ", toBalance)
		fmt.Println("toBalanceAfter: ", toBalanceAfter)
		fmt.Println("managerBalance: ", managerBalance)
		fmt.Println("managerBalanceAfter: ", managerBalanceAfter)

		So(fromBalanceAfter.Cmp(fromBalance) == 0, ShouldBeTrue)
		So(toBalanceAfter.Cmp(toBalance) == 0, ShouldBeTrue)
		So(managerBalanceAfter.Cmp(managerBalance) == 0, ShouldBeTrue)

		// 发送一笔正常交易  防止当前case影响测试
		rawTx1 := types.NewTransaction(nonce, _to, amount, gasLimit, gasPrice, nil)
		signer1 := types.NewEIP155Signer(chainId) //big.NewInt(1)//当前入参链id
		sigTransaction1, err := types.SignTx(rawTx1, signer1, fromPrivate)
		_checkErr(err)
		err = ethClient.SendTransaction(context.Background(), sigTransaction1)
		_checkErr(err)
		receipt2 := make(map[string]interface{})
		for i := 0; i < 10; i++ { //设置10次请求，每次间隔3s 30s超时
			time.Sleep(3000000000)
			if len(receipt2) != 0 {
				break
			} else {
				receipt2 = utils.GetTransferInfo(url, sigTransaction1.Hash())
			}
		}
		if len(receipt2) == 0 {
			panic("交易30s未落账!")
		}
		//fmt.Println(receipt)
		status := utils.Strval(receipt2["status"])
		//fmt.Println(status)
		if status != "0x1" {
			fmt.Println("status : ", status)
			panic("TxTransaction fail !!!")
		}
		receipt = utils.GetTransferInfo(url, sigTransaction.Hash())
		fmt.Println("receipt: ", receipt)
		useGas := utils.HexToDec(utils.Strval(receipt["cumulativeGasUsed"]))//原nonce+=1case发送的交易收据

		useGas1 := utils.HexToDec(utils.Strval(receipt2["cumulativeGasUsed"]))
		managerBalanceAfter = utils.GetBalance(url, "0x0000000000000000000000000000000000000007")
		fromBalanceAfter = utils.GetBalance(url, fromAddress.String())
		toBalanceAfter = utils.GetBalance(url, toAddress)

		fmt.Println("useGas1", useGas1)
		fmt.Println("fromBalance", fromBalance)
		fmt.Println("toBalance", toBalance)
		fmt.Println("managerBalance", managerBalance)
		fmt.Println("fromBalanceAfter", fromBalanceAfter)
		fmt.Println("toBalanceAfter", toBalanceAfter)
		fmt.Println("managerBalanceAfter", managerBalanceAfter)


		fromBalance1 := fromBalanceAfter.Add(fromBalanceAfter, amount).Add(fromBalanceAfter, useGas1).Add(fromBalanceAfter, amount).Add(fromBalanceAfter, useGas)
		So(fromBalance.Cmp(fromBalance1) == 0, ShouldBeTrue)
		//2 断言to账户资金是否 = toBalance + amount
		So(toBalanceAfter.Cmp(toBalance.Add(toBalance, amount).Add(toBalance, amount)) == 0, ShouldBeTrue)
		//3 断言治理合约地址是否正常获得消耗gas
		So(managerBalanceAfter.Cmp(managerBalance.Add(managerBalance, useGas).Add(managerBalance, useGas1)) == 0, ShouldBeTrue)

	})
}

package testcase

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/smartystreets/goconvey/convey"
	"log"
	"math/big"
	"test_evm/config"
	"test_evm/testcase/deploy"
	"test_evm/utils"
	"testing"
)

func TestEVMTransfer006(t *testing.T) {
	Convey("Test 前置发一笔正常交易", t, func() {

		toAddress := "0xf45505D1F482EBc8881dacA97B122B62771B9e1d"
		amount := big.NewInt(1000000000000000000)
		//  备注
		fromPrivateKey := config.FromPrivate
		fromPrivate, _ := crypto.HexToECDSA(fromPrivateKey)
		fromAddress := crypto.PubkeyToAddress(fromPrivate.PublicKey)
		//log.Printf("fromAddress: %d", fromAddress)
		log.Printf("fromAddress: %s", fromAddress.String())

		//交易前账户余额
		fromBalance := utils.GetBalance(fromAddress.String())
		toBalance := utils.GetBalance(toAddress)
		managerBalance := utils.GetBalance("0x0000000000000000000000000000000000000007")

		hash, err := deploy.SendTransfer(fromPrivateKey, toAddress, amount, uint64(200000))
		if err != nil {
			panic(err)
		}
		//log.Printf(hash)
		receipt, err := utils.GetTransferInfoByHash(hash)
		if err != nil {
			panic(err)
		}

		//log.Printf(status)
		if receipt.Status != 1 {
			log.Printf("receipt.Status :  %d", receipt.Status)
			panic("TxTransaction fail !!!")
		}
		useGas := big.NewInt(int64(receipt.CumulativeGasUsed*1000000000))
		log.Printf("useGas %d", useGas)
		fromBalanceAfter := utils.GetBalance(fromAddress.String())
		toBalanceAfter := utils.GetBalance(toAddress)
		managerBalanceAfter := utils.GetBalance("0x0000000000000000000000000000000000000007")
		log.Printf("fromBalance %d", fromBalance)
		log.Printf("toBalance %d", toBalance)
		log.Printf("managerBalance %d", managerBalance)
		log.Printf("fromBalanceAfter %d", fromBalanceAfter)
		log.Printf("toBalanceAfter %d", toBalanceAfter)
		log.Printf("managerBalanceAfter %d", managerBalanceAfter)
		//1 断言from账户资金是否 = fromBalance1 = fromBalance - gasUsed - amount
		fromBalance1 := fromBalanceAfter.Add(fromBalanceAfter, amount).Add(fromBalanceAfter, useGas)
		So(fromBalance.Cmp(fromBalance1) == 0, ShouldBeTrue)
		//2 断言to账户资金是否 = toBalance + amount
		So(toBalanceAfter.Cmp(toBalance.Add(toBalance, amount)) == 0, ShouldBeTrue)
		//3 断言治理合约地址是否正常获得消耗gas
		So(managerBalanceAfter.Cmp(managerBalance.Add(managerBalance, useGas)) == 0, ShouldBeTrue)
	})
	Convey("Test006 测试EIP155交易普通转账ong时 nonce值错误(nonce-1) ", t, func() {

		toAddress := "0xf45505D1F482EBc8881dacA97B122B62771B9e1d"
		amount := big.NewInt(500000000000000000)
		gasLimit := uint64(200000)
		//  备注： 对应的MetaMask钱包第5个账户私钥
		fromPrivateKey := config.FromPrivate
		fromPrivate, _ := crypto.HexToECDSA(fromPrivateKey)
		fromAddress := crypto.PubkeyToAddress(fromPrivate.PublicKey)

		//交易前账户余额
		fromBalance := utils.GetBalance(fromAddress.String())
		managerBalance := utils.GetBalance("0x0000000000000000000000000000000000000007")

		// 构建一个私链对象
		ethClient := config.GetEthClient()
		//查chainId
		chainId, err := ethClient.ChainID(context.Background())
		_checkErr(err)
		// 查nonce值
		nonce, err := ethClient.PendingNonceAt(context.Background(), fromAddress)
		_checkErr(err)
		_to := common.HexToAddress(toAddress)
		gasPrice := big.NewInt(500000000000)
		rawTx := types.NewTransaction(nonce-1, _to, amount, gasLimit, gasPrice, nil)

		signer := types.NewEIP155Signer(chainId) //big.NewInt(1)//当前入参链id
		sigTransaction, err := types.SignTx(rawTx, signer, fromPrivate)
		_checkErr(err)
		err = ethClient.SendTransaction(context.Background(), sigTransaction)
		log.Println(err)
		// 断言nonce = nonce-1 构造的交易对象,发送交易抛异常
		So(err, ShouldNotBeNil)
		managerBalanceAfter := utils.GetBalance("0x0000000000000000000000000000000000000007")
		fromBalanceAfter := utils.GetBalance(fromAddress.String())
		log.Printf("fromBalance:  %d", fromBalance)
		log.Printf("managerBalance:  %d", managerBalance)
		log.Printf("managerBalanceAfter:  %d", managerBalanceAfter)
		log.Printf("fromBalanceAfter:  %d", fromBalanceAfter)
		So(fromBalanceAfter.Cmp(fromBalance) == 0, ShouldBeTrue)
		So(managerBalanceAfter.Cmp(managerBalance) == 0, ShouldBeTrue)

	})
	Convey("Test007 测试EIP155交易普通转账ong时 nonce值错误(nonce+1) ", t, func() {

		toAddress := "0xf45505D1F482EBc8881dacA97B122B62771B9e1d"
		amount := big.NewInt(500000000000000000)
		gasLimit := uint64(200000)
		//  备注： 对应的MetaMask钱包第5个账户私钥
		fromPrivateKey := config.FromPrivate
		fromPrivate, _ := crypto.HexToECDSA(fromPrivateKey)
		fromAddress := crypto.PubkeyToAddress(fromPrivate.PublicKey)

		//交易前账户余额
		fromBalance := utils.GetBalance(fromAddress.String())
		toBalance := utils.GetBalance(toAddress)
		managerBalance := utils.GetBalance("0x0000000000000000000000000000000000000007")

		// 构建一个私链对象
		ethClient := config.GetEthClient()
		//查chainId
		chainId, err := ethClient.ChainID(context.Background())
		_checkErr(err)
		// 查nonce值
		nonce, err := ethClient.PendingNonceAt(context.Background(), fromAddress)
		_checkErr(err)
		_to := common.HexToAddress(toAddress)
		gasPrice := big.NewInt(500000000000)
		rawTx := types.NewTransaction(nonce+1, _to, amount, gasLimit, gasPrice, nil)

		signer := types.NewEIP155Signer(chainId) //big.NewInt(1)//当前入参链id
		sigTransaction, err := types.SignTx(rawTx, signer, fromPrivate)
		_checkErr(err)
		err = ethClient.SendTransaction(context.Background(), sigTransaction)
		_checkErr(err)
		receipt, err := utils.GetTransferInfoByHash(sigTransaction.Hash())
		So(err, ShouldBeError)

		managerBalanceAfter := utils.GetBalance("0x0000000000000000000000000000000000000007")
		fromBalanceAfter := utils.GetBalance(fromAddress.String())
		toBalanceAfter := utils.GetBalance(toAddress)
		log.Printf("fromBalance:  %d", fromBalance)
		log.Printf("fromBalanceAfter:  %d", fromBalanceAfter)
		log.Printf("toBalance:  %d", toBalance)
		log.Printf("toBalanceAfter:  %d", toBalanceAfter)
		log.Printf("managerBalance:  %d", managerBalance)
		log.Printf("managerBalanceAfter:  %d", managerBalanceAfter)

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
		receipt2, err := utils.GetTransferInfoByHash(sigTransaction1.Hash())
		if err != nil {
			panic(err)
		}
		//log.Printf(status)
		if receipt2.Status != 1 {
			log.Printf("receipt2.Status :  %d", receipt2.Status)
			panic("TxTransaction fail !!!")
		}
		receipt, err = utils.GetTransferInfoByHash(sigTransaction.Hash())
		if err != nil {
			panic(err)
		}
		useGas := big.NewInt(int64(receipt.CumulativeGasUsed*1000000000)) //原nonce+=1case发送的交易收据
		useGas1 := big.NewInt(int64(receipt2.CumulativeGasUsed*1000000000))
		managerBalanceAfter = utils.GetBalance("0x0000000000000000000000000000000000000007")
		fromBalanceAfter = utils.GetBalance(fromAddress.String())
		toBalanceAfter = utils.GetBalance(toAddress)

		log.Printf("useGas1 %d", useGas1)
		log.Printf("fromBalance %d", fromBalance)
		log.Printf("toBalance %d", toBalance)
		log.Printf("managerBalance %d", managerBalance)
		log.Printf("fromBalanceAfter %d", fromBalanceAfter)
		log.Printf("toBalanceAfter %d", toBalanceAfter)
		log.Printf("managerBalanceAfter %d", managerBalanceAfter)

		fromBalance1 := fromBalanceAfter.Add(fromBalanceAfter, amount).Add(fromBalanceAfter, useGas1).Add(fromBalanceAfter, amount).Add(fromBalanceAfter, useGas)
		So(fromBalance.Cmp(fromBalance1) == 0, ShouldBeTrue)
		//2 断言to账户资金是否 = toBalance + amount
		So(toBalanceAfter.Cmp(toBalance.Add(toBalance, amount).Add(toBalance, amount)) == 0, ShouldBeTrue)
		//3 断言治理合约地址是否正常获得消耗gas
		So(managerBalanceAfter.Cmp(managerBalance.Add(managerBalance, useGas).Add(managerBalance, useGas1)) == 0, ShouldBeTrue)

	})
}

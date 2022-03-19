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

func TestEVMTransfer_add(t *testing.T) {
	Convey("Test 测试ong,ong  balance < gas price * gas limit", t, func() {
		toAddress := "0xf45505D1F482EBc8881dacA97B122B62771B9e1d"
		//amount := big.NewInt(1000000000)
		//  备注： 对应的MetaMask钱包第一个账户私钥
		moneyPrivateKey := config.FromPrivate
		moneyPrivate, _ := crypto.HexToECDSA(moneyPrivateKey)
		moneyAddress := crypto.PubkeyToAddress(moneyPrivate.PublicKey)
		log.Printf("moneyAddress: %s", moneyAddress.String())
		//交易前账户余额
		moneyBalance := utils.GetBalance(moneyAddress.String())
		log.Printf("moneyBalance: %d", moneyBalance)

		fromPrivateKey := config.GetPrivateList()[3]
		fromPrivate, _ := crypto.HexToECDSA(fromPrivateKey)
		fromAddress := crypto.PubkeyToAddress(fromPrivate.PublicKey)
		log.Printf("fromAddress: %s", fromAddress.String())
		//交易前账户余额
		checkFromBalance(fromAddress.String(),big.NewInt(100000000000000000))
		fromBalance := utils.GetBalance(fromAddress.String())
		log.Printf("fromBalance: %d", fromBalance)
		toBalance := utils.GetBalance(toAddress)
		managerBalance := utils.GetBalance("0x0000000000000000000000000000000000000007")
		log.Printf("toBalance: %d", toBalance)
		log.Printf("managerBalance: %d", managerBalance)

		_,err :=deploy.SendTransfer(fromPrivateKey, toAddress, big.NewInt(0), uint64(2000000))
		So(err,ShouldNotBeNil)
		fromBalanceAfter := utils.GetBalance(fromAddress.String())
		toBalanceAfter := utils.GetBalance(toAddress)
		managerBalanceAfter := utils.GetBalance("0x0000000000000000000000000000000000000007")
		log.Printf("fromBalanceAfter: %d", fromBalanceAfter)
		log.Printf("toBalanceAfter: %d", toBalanceAfter)
		log.Printf("managerBalanceAfter: %d", managerBalanceAfter)
		So(fromBalance.Cmp(fromBalanceAfter)==0,ShouldBeTrue)
		So(toBalance.Cmp(toBalanceAfter)==0,ShouldBeTrue)
		So(managerBalance.Cmp(managerBalanceAfter)==0,ShouldBeTrue)
	})
}

func TestEVMTransfer_add2(t *testing.T) {
	Convey("Test 测试ong,ong  balance < gas price * gas limit", t, func() {

		toAddress := "0xf45505D1F482EBc8881dacA97B122B62771B9e1d"
		gasLimit := uint64(21000)//200000
		//  备注
		moneyPrivateKey := config.FromPrivate
		moneyPrivate, _ := crypto.HexToECDSA(moneyPrivateKey)
		moneyAddress := crypto.PubkeyToAddress(moneyPrivate.PublicKey)
		log.Printf("moneyAddress: %s", moneyAddress.String())
		//交易前账户余额
		moneyBalance := utils.GetBalance(moneyAddress.String())
		log.Printf("moneyBalance: %d", moneyBalance)

		fromPrivateKey := config.GetPrivateList()[4]
		fromPrivate, _ := crypto.HexToECDSA(fromPrivateKey)
		fromAddress := crypto.PubkeyToAddress(fromPrivate.PublicKey)
		//交易前账户余额
		checkFromBalance(fromAddress.String(),big.NewInt(15000000000000000))
		log.Printf("fromAddress: %s", fromAddress.String())

		//交易前账户余额
		fromBalance := utils.GetBalance(fromAddress.String())
		log.Printf("fromBalance : %d", fromBalance)
		toBalance := utils.GetBalance(toAddress)
		managerBalance := utils.GetBalance("0x0000000000000000000000000000000000000007")

		// 构建一个私链对象
		ethClient := config.GetEthClient()
		//查chainId
		chainId, err := ethClient.ChainID(context.Background())
		_checkErr(err)
		// 查nonce值
		nonce, err := ethClient.PendingNonceAt(context.Background(), fromAddress)
		log.Printf("nonce: %d", nonce)
		_checkErr(err)
		_to := common.HexToAddress(toAddress)
		gasPrice := big.NewInt(500000000000)
		amount := big.NewInt(0)
		rawTx1 := types.NewTransaction(nonce, _to, amount, gasLimit, gasPrice, nil)
		rawTx2 := types.NewTransaction(nonce+1, _to, amount, gasLimit, gasPrice, nil)
		hash1, _ := _singTransaction(chainId, fromPrivateKey, rawTx1, ethClient)
		hash2, _ := _singTransaction(chainId, fromPrivateKey, rawTx2, ethClient)

		log.Printf("hash1: %s", hash1)
		log.Printf("hash2: %s", hash2)

		receipt1, err := utils.GetTransferInfoByHash(hash1)
		if err != nil {
			panic(err)
		}
		receipt2, err := utils.GetTransferInfoByHash(hash2)
		if err != nil {
			panic(err)
		}

		if (receipt1.Status != 1 || receipt2.Status != 1 ) && (receipt1.Status == 0 || receipt2.Status == 0 ){
			log.Printf("receipt1.Status : %d", receipt1.Status)
			log.Printf("receipt2.Status : %d", receipt2.Status)
			//panic("TxTransaction fail !!!")
		}
		So(receipt2.Status, ShouldEqual, 0) // 断言第二笔交易状态是否为失败

		useGas1 := big.NewInt(int64(receipt1.CumulativeGasUsed*1000000000))
		log.Printf("useGas1 %d", useGas1)
		useGas2 := big.NewInt(int64(receipt2.CumulativeGasUsed*1000000000))
		log.Printf("useGas2 %d", useGas2)
		fromBalanceAfter := utils.GetBalance(fromAddress.String())
		toBalanceAfter := utils.GetBalance(toAddress)
		managerBalanceAfter := utils.GetBalance("0x0000000000000000000000000000000000000007")
		log.Printf("fromBalance  %d", fromBalance)
		log.Printf("toBalance  %d", toBalance)
		log.Printf("managerBalance  %d", managerBalance)
		log.Printf("fromBalanceAfter  %d", fromBalanceAfter)
		log.Printf("toBalanceAfter  %d", toBalanceAfter)
		log.Printf("managerBalanceAfter  %d", managerBalanceAfter)

		//1.请求一调用成功，交易正常落账
		//2.请求二EVM事务回滚，扣除消耗gas

		//1 断言from账户资金是否 = fromBalance1 = fromBalance - gasUsed1 - gasUsed2 - amount
		log.Printf("useGas1 %d", useGas1)
		log.Printf("useGas2 %d", useGas2)
		log.Printf("amount %d", amount)

		fromBalance1 := fromBalanceAfter.Add(fromBalanceAfter, amount).Add(fromBalanceAfter, useGas1).Add(fromBalanceAfter, useGas2)
		So(fromBalance.Cmp(fromBalance1) == 0, ShouldBeTrue)
		//2 断言to账户资金是否 = toBalance + amount
		So(toBalanceAfter.Cmp(toBalance.Add(toBalance, amount)) == 0, ShouldBeTrue)
		//3 断言治理合约地址是否正常获得消耗gas
		log.Printf("managerBalance_1  %d", managerBalance)
		//log.Printf("result %d", managerBalance.Add(managerBalance, useGas1).Add(managerBalance, useGas2))
		So(managerBalanceAfter.Cmp(managerBalance.Add(managerBalance, useGas1).Add(managerBalance, useGas2)) == 0, ShouldBeTrue)

		})
}


func checkFromBalance(fromAddress string,amount *big.Int) {
	fromBalance := utils.GetBalance(fromAddress)
	if fromBalance.Cmp(amount) < 0 {
		diff := big.NewInt(0).Sub(amount,fromBalance)
		txHash, err := deploy.SendTransfer(config.FromPrivate, fromAddress, diff, uint64(200000))
		_checkErr(err)
		re, err := utils.GetTransferInfoByHash(txHash)
		_checkErr(err)
		if re.Status != 1 {
			panic("error ! formAddress balance low")
		}
	}
}
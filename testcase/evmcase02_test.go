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
	"test_evm/utils"
	"testing"
	"time"
)

func TestEVMTransfer005(t *testing.T) {
	Convey("Test005 测试EIP155交易普通转账ong 并发(ong不足两次操作消耗和发送的数量)", t, func() {
		url := "http://127.0.0.1:20339"
		toAddress := "0x5f9a233682e80e8bb1bfdf18cd21d82af2fd7edd"
		amount := big.NewInt(500000000)
		gasLimit := uint64(200000)
		//  备注： 对应的MetaMask钱包第5个账户私钥
		fromPrivateKey := "4929d50ed435df55ad5f654a3fb8f67d75c50ae5aa95ed8a9ae62ffff48f90b9"
		fromPrivate, _ := crypto.HexToECDSA(fromPrivateKey)
		fromAddress := crypto.PubkeyToAddress(fromPrivate.PublicKey)

		//交易前账户余额
		fromBalance := utils.GetBalance(url, fromAddress.String())
		fmt.Println("fromBalance : ", fromBalance)
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
		//amount := big.NewInt(257000000000)
		rawTx1 := types.NewTransaction(nonce, _to, amount, gasLimit, gasPrice, nil)
		rawTx2 := types.NewTransaction(nonce+1, _to, amount, gasLimit, gasPrice, nil)
		hash1, _ := _singTransaction(chainId, fromPrivateKey, rawTx1, ethClient)
		hash2, _ := _singTransaction(chainId, fromPrivateKey, rawTx2, ethClient)

		fmt.Println("hash1: ", hash1)
		fmt.Println("hash2: ", hash2)

		receipt1 := make(map[string]interface{})
		receipt2 := make(map[string]interface{})
		for i := 0; i < 10; i++ { //设置10次请求，每次间隔3s 30s超时
			time.Sleep(3000000000)
			if len(receipt1) != 0 {
				break
			} else {
				receipt1 = utils.GetTransferInfo(url, hash1)
				receipt2 = utils.GetTransferInfo(url, hash2)
			}
		}
		fmt.Println("receipt1 : ", receipt1)
		fmt.Println("receipt2 : ", receipt2)

		if len(receipt1) == 0 || len(receipt2) == 0 {
			panic("交易30s未落账!")
		}
		status1 := utils.Strval(receipt1["status"])
		if status1 != "0x1" {
			fmt.Println("status1 : ", status1)
			panic("TxTransaction fail !!!")
		}
		fmt.Println("status: 1", status1)
		status2 := utils.Strval(receipt2["status"])
		So(status2, ShouldEqual, "0x0") // 断言第二笔交易状态是否为失败

		useGas1 := utils.HexToDec(utils.Strval(receipt1["cumulativeGasUsed"]))
		fmt.Println("useGas1 ", useGas1)
		useGas2 := utils.HexToDec(utils.Strval(receipt2["cumulativeGasUsed"]))
		fmt.Println("useGas2 ", useGas2)
		fromBalanceAfter := utils.GetBalance(url, fromAddress.String())
		toBalanceAfter := utils.GetBalance(url, toAddress)
		managerBalanceAfter := utils.GetBalance(url, "0x0000000000000000000000000000000000000007")
		fmt.Println("fromBalance ", fromBalance)
		fmt.Println("toBalance ", toBalance)
		fmt.Println("managerBalance ", managerBalance)
		fmt.Println("fromBalanceAfter ", fromBalanceAfter)
		fmt.Println("toBalanceAfter ", toBalanceAfter)
		fmt.Println("managerBalanceAfter ", managerBalanceAfter)

		//1.请求一调用成功，交易正常落账
		//2.请求二EVM事务回滚，扣除消耗gas

		//1 断言from账户资金是否 = fromBalance1 = fromBalance - gasUsed1 - gasUsed2 - amount
		fmt.Println("useGas1",useGas1)
		fmt.Println("useGas2",useGas2)
		fmt.Println("amount",amount)
		fromBalance1 := fromBalanceAfter.Add(fromBalanceAfter, amount).Add(fromBalanceAfter, useGas1).Add(fromBalanceAfter, useGas2)
		fmt.Println(fromBalance1)

		//2 断言to账户资金是否 = toBalance + amount
		So(toBalanceAfter.Cmp(toBalance.Add(toBalance, amount)) == 0, ShouldBeTrue)
		//3 断言治理合约地址是否正常获得消耗gas
		So(managerBalanceAfter.Cmp(managerBalance.Add(managerBalance, useGas1).Add(managerBalance, useGas2)) == 0, ShouldBeTrue)

		So(fromBalance.Cmp(fromBalance1) == 0, ShouldBeTrue)
	})

}

func _checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func _singTransaction(chainId *big.Int, fromPrivateKey string, rawTx *types.Transaction, ethClient *ethclient.Client) (common.Hash, error) {
	signer := types.NewEIP155Signer(chainId) //big.NewInt(1)//当前入参链id
	key, err := crypto.HexToECDSA(fromPrivateKey)
	if err != nil {
		return common.Hash{}, err
	}
	//对交易对象做签名 0交易对象  1签名类型types.NewEIP155Signer(chainId)  2签名账户私钥
	sigTransaction, err := types.SignTx(rawTx, signer, key)
	if err != nil {
		return common.Hash{}, err
	}
	go func() {
		err = ethClient.SendTransaction(context.Background(), sigTransaction)
	}()
	fmt.Println("tx hash(evm): ", sigTransaction.Hash())
	return sigTransaction.Hash(), nil
}

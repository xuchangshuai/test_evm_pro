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
	"test_evm/config"
	"test_evm/testcase/deploy"
	"test_evm/utils"
	"testing"
)

func TestEVMTransfer005(t *testing.T) {
	Convey("Test005 前置 fromAddress账户资金准备1ong", t, func() {

		fromPrivateKey := config.FromPrivate
		toAddress := "0x69ed51D295cD121d279EfF97a227Ad510F61E3c0"
		toBalance := utils.GetBalance(toAddress)
		amount := big.NewInt(0)
		if toBalance.Cmp(big.NewInt(1000000000)) == 0 {
			return
		}
		amount.Sub(big.NewInt(1000000000), toBalance)
		fmt.Println("amount: ", amount)
		hash, err := deploy.SendTransfer(fromPrivateKey, toAddress, amount, uint64(200000))
		if err != nil {
			panic(err)
		}
		//fmt.Println(hash)
		receipt, err := utils.GetTransferInfoByHash(hash)
		if err != nil {
			panic(err)
		}
		if receipt.Status != 1 {
			fmt.Println("receipt.Status: ", receipt.Status)
			panic("交易30s未落账!")
		}
		toBalance = utils.GetBalance(toAddress)
		//断言to账户资金是否为1
		fmt.Println("toBalance: ", toBalance)
		So(toBalance.Cmp(big.NewInt(1000000000)) == 0, ShouldBeTrue)

	})
	Convey("Test005 测试EIP155交易普通转账ong 并发(ong不足两次操作消耗和发送的数量)", t, func() {

		toAddress := "0xf45505D1F482EBc8881dacA97B122B62771B9e1d"
		amount := big.NewInt(500000000)
		gasLimit := uint64(200000)
		//  备注
		//0x69ed51D295cD121d279EfF97a227Ad510F61E3c0 //ARRxttxxivKVfzkfdC5T3FBNJTEDrvbT9X
		fromPrivateKey := "4929d50ed435df55ad5f654a3fb8f67d75c50ae5aa95ed8a9ae62ffff48f90b9"
		fromPrivate, _ := crypto.HexToECDSA(fromPrivateKey)
		fromAddress := crypto.PubkeyToAddress(fromPrivate.PublicKey)
		fmt.Println("fromAddress: ", fromAddress)

		//交易前账户余额
		fromBalance := utils.GetBalance(fromAddress.String())
		fmt.Println("fromBalance : ", fromBalance)
		toBalance := utils.GetBalance(toAddress)
		managerBalance := utils.GetBalance("0x0000000000000000000000000000000000000007")

		// 构建一个私链对象
		ethClient := config.GetEthClient()
		//查chainId
		chainId, err := ethClient.ChainID(context.Background())
		_checkErr(err)
		// 查nonce值
		nonce, err := ethClient.PendingNonceAt(context.Background(), fromAddress)
		fmt.Println("nonce: ", nonce)
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

		receipt1, err := utils.GetTransferInfoByHash(hash1)
		if err != nil {
			panic(err)
		}
		receipt2, err := utils.GetTransferInfoByHash(hash2)
		if err != nil {
			panic(err)
		}

		if receipt1.Status != 1 {
			fmt.Println("receipt1.Status : ", receipt1.Status)
			panic("TxTransaction fail !!!")
		}
		So(receipt2.Status, ShouldEqual, 0) // 断言第二笔交易状态是否为失败

		useGas1 := big.NewInt(int64(receipt1.CumulativeGasUsed))
		fmt.Println("useGas1 ", useGas1)
		useGas2 := big.NewInt(int64(receipt2.CumulativeGasUsed))
		fmt.Println("useGas2 ", useGas2)
		fromBalanceAfter := utils.GetBalance(fromAddress.String())
		toBalanceAfter := utils.GetBalance(toAddress)
		managerBalanceAfter := utils.GetBalance("0x0000000000000000000000000000000000000007")
		fmt.Println("fromBalance ", fromBalance)
		fmt.Println("toBalance ", toBalance)
		fmt.Println("managerBalance ", managerBalance)
		fmt.Println("fromBalanceAfter ", fromBalanceAfter)
		fmt.Println("toBalanceAfter ", toBalanceAfter)
		fmt.Println("managerBalanceAfter ", managerBalanceAfter)

		//1.请求一调用成功，交易正常落账
		//2.请求二EVM事务回滚，扣除消耗gas

		//1 断言from账户资金是否 = fromBalance1 = fromBalance - gasUsed1 - gasUsed2 - amount
		fmt.Println("useGas1", useGas1)
		fmt.Println("useGas2", useGas2)
		fmt.Println("amount", amount)
		fromBalance1 := fromBalanceAfter.Add(fromBalanceAfter, amount).Add(fromBalanceAfter, useGas1).Add(fromBalanceAfter, useGas2)
		So(fromBalance.Cmp(fromBalance1) == 0, ShouldBeTrue)
		//2 断言to账户资金是否 = toBalance + amount
		So(toBalanceAfter.Cmp(toBalance.Add(toBalance, amount)) == 0, ShouldBeTrue)
		//3 断言治理合约地址是否正常获得消耗gas
		So(managerBalanceAfter.Cmp(managerBalance.Add(managerBalance, useGas1).Add(managerBalance, useGas2)) == 0, ShouldBeTrue)

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

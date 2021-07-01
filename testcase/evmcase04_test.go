package testcase

import (
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/smartystreets/goconvey/convey"
	"math/big"
	"test_evm/testcase/deploy"
	"test_evm/utils"
	"testing"
	"time"
)

func TestEVMTransfer008(t *testing.T) {
	Convey("Test008 测试EIP155交易 - 部署合约 - 往合约转账 {断言账户资金变动 治理地址手续费情况}", t, func() {
		url := "http://127.0.0.1:20339"
		// 备注: 对应的MetaMask钱包第一个账户私钥
		fromPrivateKey := "80a68081edc0aed4ddf8fa9f6a2e7cf8d0a69c998d4ef646f6446cbf4cfe9145"
		fromPrivate, _ := crypto.HexToECDSA(fromPrivateKey)
		fromAddress := crypto.PubkeyToAddress(fromPrivate.PublicKey)
		fmt.Println("fromAddress: ", fromAddress)
		fromBalance := utils.GetBalance(url, fromAddress.String())
		managerBalance := utils.GetBalance(url, "0x0000000000000000000000000000000000000007")

		abi, err := utils.ReadAll("deploy/outputCont/OutputCont_sol_OutputContract.abi")
		_checkErr(err)
		bin, err := utils.ReadAll("deploy/outputCont/OutputCont_sol_OutputContract.bin")
		_checkErr(err)
		hash := deploy.ContractDeploy(url, string(bin), string(abi))
		fmt.Println("hash: ", hash)
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

		fmt.Println("receipt: ", receipt)
		deployContractUseGas := utils.HexToDec(utils.Strval(receipt["cumulativeGasUsed"]))
		fmt.Println("deployContractUseGas: ", deployContractUseGas)
		managerBalanceAfter := utils.GetBalance(url, "0x0000000000000000000000000000000000000007")
		fmt.Println( "managerBalanceAfter: ", managerBalanceAfter)
		//断言部署合约手续费
		So(managerBalanceAfter.Cmp(managerBalance.Add(managerBalance, deployContractUseGas)) == 0, ShouldBeTrue)
		status := utils.Strval(receipt["status"])
		//fmt.Println(status)
		if status != "0x1" {
			fmt.Println("status : ", status)
			panic("TxTransaction fail !!!")
		}

		//交易后账户余额
		fromBalanceAfter := utils.GetBalance(url, fromAddress.String())
		fromBalanceAfterRes := *fromBalanceAfter
		fmt.Println("fromBalanceFirstAfter: ", fromBalanceAfter)
		So(fromBalance.Cmp(fromBalanceAfterRes.Add(&fromBalanceAfterRes, deployContractUseGas)) == 0, ShouldBeTrue)
		contractAddress := utils.Strval(receipt["contractAddress"])
		fmt.Println("contractAddress: ",contractAddress)
		contractAddressBalance := utils.GetBalance(url, contractAddress)
		fmt.Println("contractAddressBalance: ", contractAddressBalance)

		amount := big.NewInt(1000000000)
		txHash,err := deploy.SendTransfer(url, fromPrivateKey, contractAddress, amount, uint64(200000))
		_checkErr(err)
		txReceipt := make(map[string]interface{})
		for i := 0; i < 10; i++ { //设置10次请求，每次间隔3s 30s超时
			time.Sleep(3000000000)
			if len(txReceipt) != 0 {
				break
			} else {
				txReceipt = utils.GetTransferInfo(url, txHash)
			}
		}
		if len(txReceipt) == 0 {
			panic("交易30s未落账!")
		}
		fmt.Println(txReceipt)
		txStatus := utils.Strval(txReceipt["status"])
		//fmt.Println(status)
		if txStatus != "0x1" {
			fmt.Println("txStatus : ", txStatus)
			panic("TxTransaction fail !!!")
		}
		useGas := utils.HexToDec(utils.Strval(txReceipt["cumulativeGasUsed"]))
		fromBalanceSecondAfter := utils.GetBalance(url, fromAddress.String())
		contractAddressBalanceAfter := utils.GetBalance(url, contractAddress)
		managerBalanceSecondAfter := utils.GetBalance(url, "0x0000000000000000000000000000000000000007")

		//1 断言合约账户账户资金是否 = contractAddressBalance + amount
		So(contractAddressBalanceAfter.Cmp(contractAddressBalance.Add(contractAddressBalance, amount)) == 0, ShouldBeTrue)
		//2 断言治理合约地址是否正常获得消耗gas
		So(managerBalanceSecondAfter.Cmp(managerBalanceAfter.Add(managerBalanceAfter, useGas)) == 0, ShouldBeTrue)
		//3 from账户资金= fromBalanceAfter -fromBalanceSecondAfter -amount-useGas
		So(fromBalanceAfter.Cmp(fromBalanceSecondAfter.Add(fromBalanceSecondAfter, amount).Add(fromBalanceSecondAfter, useGas)) == 0, ShouldBeTrue)

	})
}

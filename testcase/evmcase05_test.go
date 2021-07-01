package testcase

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	. "github.com/smartystreets/goconvey/convey"
	"math/big"
	"test_evm/testcase/deploy"
	"test_evm/testcase/deploy/inputCont"
	"test_evm/utils"
	"testing"
	"time"
)

func TestEVMTransfer009(t *testing.T) {
	Convey("Test008 测试EIP155交易 - 部署合约 - 往合约转账A(合约内部往合约B转账) {断言账户资金变动 治理地址手续费情况}", t, func() {
		url := "http://127.0.0.1:20339"
		ethClient, err := ethclient.Dial(url)
		_checkErr(err)
		defer ethClient.Close()
		chainId, err := ethClient.ChainID(context.Background())
		_checkErr(err)

		// 备注: 对应的MetaMask钱包第一个账户私钥
		fromPrivateKey := "80a68081edc0aed4ddf8fa9f6a2e7cf8d0a69c998d4ef646f6446cbf4cfe9145"
		fromPrivate, _ := crypto.HexToECDSA(fromPrivateKey)
		fromAddress := crypto.PubkeyToAddress(fromPrivate.PublicKey)
		fmt.Println("fromAddress: ", fromAddress)

		outputContractAbi, err := utils.ReadAll("deploy/outputCont/OutputCont_sol_OutputContract.abi")
		_checkErr(err)
		outputContractBin, err := utils.ReadAll("deploy/outputCont/OutputCont_sol_OutputContract.bin")
		_checkErr(err)
		outputContractHash := deploy.ContractDeploy(url, string(outputContractBin), string(outputContractAbi))
		fmt.Println("outputContractHash: ", outputContractHash)
		outputContractReceipt := make(map[string]interface{})
		for i := 0; i < 10; i++ { //设置10次请求，每次间隔3s 30s超时
			time.Sleep(3000000000)
			if len(outputContractReceipt) != 0 {
				break
			} else {
				outputContractReceipt = utils.GetTransferInfo(url, outputContractHash)
			}
		}
		if len(outputContractReceipt) == 0 {
			panic("交易30s未落账!")
		}
		outputContractStatus := utils.Strval(outputContractReceipt["status"])
		//fmt.Println(outputContractStatus)
		if outputContractStatus != "0x1" {
			fmt.Println("outputContractStatus : ", outputContractStatus)
			panic("TxTransaction fail !!!")
		}
		//outputContract合约地址
		outputContractAddress := utils.Strval(outputContractReceipt["contractAddress"])
		fmt.Println("outputContractAddress: ", outputContractAddress)

		inputContractAbi, err := utils.ReadAll("deploy/inputCont/InputCont_sol_InputContract.abi")
		_checkErr(err)
		inputContractBin, err := utils.ReadAll("deploy/inputCont/InputCont_sol_InputContract.bin")
		_checkErr(err)
		inputContractHash := deploy.ContractDeploy(url, string(inputContractBin), string(inputContractAbi), common.HexToAddress(outputContractAddress))
		fmt.Println("inputContractHash: ", inputContractHash)
		inputContractReceipt := make(map[string]interface{})
		for i := 0; i < 10; i++ { //设置10次请求，每次间隔3s 30s超时
			time.Sleep(3000000000)
			if len(inputContractReceipt) != 0 {
				break
			} else {
				inputContractReceipt = utils.GetTransferInfo(url, inputContractHash)
			}
		}
		if len(inputContractReceipt) == 0 {
			panic("交易30s未落账!")
		}
		inputContractStatus := utils.Strval(inputContractReceipt["status"])
		//fmt.Println(outputContractStatus)
		if inputContractStatus != "0x1" {
			fmt.Println("inputContractStatus : ", inputContractStatus)
			panic("TxTransaction fail !!!")
		}
		//outputContract合约地址
		inputContractAddress := utils.Strval(inputContractReceipt["contractAddress"])
		fmt.Println("inputContractAddress: ", inputContractAddress)

		fromBalance := utils.GetBalance(url, fromAddress.String())
		managerBalance := utils.GetBalance(url, "0x0000000000000000000000000000000000000007")
		outputContractBalance := utils.GetBalance(url, outputContractAddress)
		inputContractBalance := utils.GetBalance(url, inputContractAddress)

		nonce, err := ethClient.PendingNonceAt(context.Background(), fromAddress)
		_checkErr(err)
		amount := big.NewInt(10000000000)
		opts, err := bind.NewKeyedTransactorWithChainID(fromPrivate, chainId)
		_checkErr(err)
		opts.Value = amount
		opts.Nonce = big.NewInt(int64(nonce))
		opts.GasPrice = big.NewInt(500)
		opts.GasLimit = 3000000

		inputContract, err := inputCont.NewInputCont(common.HexToAddress(inputContractAddress), ethClient)
		_checkErr(err)
		tx, err := inputContract.TransferAccounts(opts)
		_checkErr(err)
		fmt.Println("tx.Hash: ", tx.Hash())
		receipt := make(map[string]interface{})
		for i := 0; i < 10; i++ { //设置最多10次请求，每次间隔3s 30s超时
			time.Sleep(3000000000)
			if len(receipt) != 0 {
				break
			} else {
				receipt = utils.GetTransferInfo(url, tx.Hash())
			}
		}
		if len(receipt) == 0 {
			panic("交易30s未落账!")
		}
		fmt.Println("receipt: ", receipt)
		useGas := utils.HexToDec(utils.Strval(receipt["cumulativeGasUsed"]))

		fromBalanceAfter := utils.GetBalance(url, fromAddress.String())
		managerBalanceAfter := utils.GetBalance(url, "0x0000000000000000000000000000000000000007")
		outputContractBalanceAfter := utils.GetBalance(url, outputContractAddress)
		inputContractBalanceAfter := utils.GetBalance(url, inputContractAddress)

		fmt.Println("useGas: ", useGas)
		fmt.Println("fromBalance: ", fromBalance)
		fmt.Println("fromBalanceAfter: ", fromBalanceAfter)
		fmt.Println("managerBalance: ", managerBalance)
		fmt.Println("managerBalanceAfter: ", managerBalanceAfter)
		fmt.Println("outputContractBalance: ", outputContractBalance)
		fmt.Println("outputContractBalanceAfter: ", outputContractBalanceAfter)
		fmt.Println("inputContractBalance: ", inputContractBalance)
		fmt.Println("inputContractBalanceAfter: ", inputContractBalanceAfter)

		// 断言from账户资金变动
		So(fromBalance.Cmp(fromBalanceAfter.Add(fromBalanceAfter, useGas).Add(fromBalanceAfter,amount)) == 0, ShouldBeTrue)
		// 断言治理合约账户资金变动 手续费情况
		So(managerBalanceAfter.Cmp(managerBalance.Add(managerBalance, useGas)) == 0, ShouldBeTrue)
		// 断言output合约账户资金情况
		outputAmount :=big.NewInt(1)
		outputAmount.Sub(amount,big.NewInt(1000000000))
		So(outputContractBalanceAfter.Cmp(outputContractBalance.Add(outputContractBalance, outputAmount)) == 0, ShouldBeTrue)
		// 断言input合约账户资金情况
		inputAmount :=big.NewInt(1000000000)
		So(inputContractBalanceAfter.Cmp(inputContractBalance.Add(inputContractBalance, inputAmount)) == 0, ShouldBeTrue)



	})
}

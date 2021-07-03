package testcase

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/smartystreets/goconvey/convey"
	"math/big"
	"test_evm/config"
	"test_evm/testcase/deploy"
	"test_evm/testcase/deploy/inputCont"
	"test_evm/utils"
	"testing"
)

func TestEVMTransfer009(t *testing.T) {
	Convey("Test008 测试EIP155交易 - 部署合约 - 往合约转账A(合约内部往合约B转账) {断言账户资金变动 治理地址手续费情况}", t, func() {
		ethClient:=config.GetEthClient()
		defer ethClient.Close()
		chainId, err := ethClient.ChainID(context.Background())
		_checkErr(err)

		// 备注: 对应的MetaMask钱包第一个账户私钥
		fromPrivateKey := config.FromPrivate
		fromPrivate, _ := crypto.HexToECDSA(fromPrivateKey)
		fromAddress := crypto.PubkeyToAddress(fromPrivate.PublicKey)
		fmt.Println("fromAddress: ", fromAddress)

		outputContractAbi, err := utils.ReadAll("deploy/outputCont/OutputCont_sol_OutputContract.abi")
		_checkErr(err)
		outputContractBin, err := utils.ReadAll("deploy/outputCont/OutputCont_sol_OutputContract.bin")
		_checkErr(err)
		outputContractHash := deploy.ContractDeploy(  string(outputContractBin), string(outputContractAbi))
		fmt.Println("outputContractHash: ", outputContractHash)
		outputContractReceipt, err := utils.GetTransferInfoByHash(outputContractHash)
		_checkErr(err)
		if outputContractReceipt.Status != 1 {
			fmt.Println("outputContractReceipt.Status : ", outputContractReceipt.Status)
			panic("TxTransaction fail !!!")
		}
		//outputContract合约地址
		outputContractAddress := outputContractReceipt.ContractAddress.String()
		fmt.Println("outputContractAddress: ", outputContractAddress)
		inputContractAbi, err := utils.ReadAll("deploy/inputCont/InputCont_sol_InputContract.abi")
		_checkErr(err)
		inputContractBin, err := utils.ReadAll("deploy/inputCont/InputCont_sol_InputContract.bin")
		_checkErr(err)
		inputContractHash := deploy.ContractDeploy(  string(inputContractBin), string(inputContractAbi), common.HexToAddress(outputContractAddress))
		fmt.Println("inputContractHash: ", inputContractHash)
		inputContractReceipt,err:=utils.GetTransferInfoByHash(inputContractHash)
		_checkErr(err)
		if inputContractReceipt.Status != 1 {
			fmt.Println("inputContractReceipt.Status : ", inputContractReceipt.Status)
			panic("TxTransaction fail !!!")
		}
		//outputContract合约地址
		inputContractAddress := inputContractReceipt.ContractAddress.String()
		fmt.Println("inputContractAddress: ", inputContractAddress)

		fromBalance := utils.GetBalance(  fromAddress.String())
		managerBalance := utils.GetBalance(  "0x0000000000000000000000000000000000000007")
		outputContractBalance := utils.GetBalance(  outputContractAddress)
		inputContractBalance := utils.GetBalance(  inputContractAddress)

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
		receipt,err := utils.GetTransferInfoByHash(tx.Hash())
		_checkErr(err)
		fmt.Println("receipt: ", receipt)
		useGas := big.NewInt(int64(receipt.CumulativeGasUsed))

		fromBalanceAfter := utils.GetBalance(  fromAddress.String())
		managerBalanceAfter := utils.GetBalance(  "0x0000000000000000000000000000000000000007")
		outputContractBalanceAfter := utils.GetBalance(  outputContractAddress)
		inputContractBalanceAfter := utils.GetBalance(  inputContractAddress)

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

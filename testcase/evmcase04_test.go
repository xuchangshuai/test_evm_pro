package testcase

import (
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/smartystreets/goconvey/convey"
	"math/big"
	"reflect"
	"test_evm/config"
	"test_evm/testcase/deploy"
	"test_evm/utils"
	"testing"
)

func TestEVMTransfer008(t *testing.T) {
	Convey("Test008 测试EIP155交易 - 部署合约 - 往合约转账 {断言账户资金变动 治理地址手续费情况}", t, func() {

		// 备注: 对应的MetaMask钱包第一个账户私钥
		fromPrivateKey := config.FromPrivate
		fromPrivate, _ := crypto.HexToECDSA(fromPrivateKey)
		fromAddress := crypto.PubkeyToAddress(fromPrivate.PublicKey)
		fmt.Println("fromAddress: ", fromAddress)
		fromBalance := utils.GetBalance(fromAddress.String())
		managerBalance := utils.GetBalance("0x0000000000000000000000000000000000000007")

		abi, err := utils.ReadAll("deploy/outputCont/OutputCont_sol_OutputContract.abi")
		_checkErr(err)
		bin, err := utils.ReadAll("deploy/outputCont/OutputCont_sol_OutputContract.bin")
		_checkErr(err)
		hash := deploy.ContractDeploy(string(bin), string(abi))
		fmt.Println("hash: ", hash)
		receipt, err := utils.GetTransferInfoByHash(hash)
		if err != nil {
			panic(err)
		}

		fmt.Println("receipt: ", receipt)
		deployContractUseGas := big.NewInt(int64(receipt.CumulativeGasUsed*1000000000))
		fmt.Println("deployContractUseGas: ", deployContractUseGas)
		managerBalanceAfter := utils.GetBalance("0x0000000000000000000000000000000000000007")
		fmt.Println("managerBalanceAfter: ", managerBalanceAfter)
		fmt.Println(reflect.TypeOf(managerBalanceAfter))
		//断言部署合约手续费
		So(managerBalanceAfter.Cmp(managerBalance.Add(managerBalance, deployContractUseGas)) == 0, ShouldBeTrue)
		if receipt.Status != 1 {
			fmt.Println("receipt.Status : ", receipt.Status)
			panic("TxTransaction fail !!!")
		}
		//交易后账户余额
		fromBalanceAfter := utils.GetBalance(fromAddress.String())
		fromBalanceAfterRes := utils.GetBalance(fromAddress.String())
		fmt.Println("fromBalanceFirstAfter: ", fromBalanceAfter)
		So(fromBalance.Cmp(fromBalanceAfterRes.Add(fromBalanceAfterRes, deployContractUseGas)) == 0, ShouldBeTrue)
		contractAddress := receipt.ContractAddress
		fmt.Println("contractAddress: ", contractAddress.String())
		contractAddressBalance := utils.GetBalance(contractAddress.String())
		fmt.Println("contractAddressBalance: ", contractAddressBalance)

		amount := big.NewInt(1000000000000000000)
		txHash, err := deploy.SendTransfer(fromPrivateKey, contractAddress.String(), amount, uint64(200000))
		_checkErr(err)
		txReceipt, err := utils.GetTransferInfoByHash(txHash)
		if err != nil {
			panic(err)
		}
		if txReceipt.Status != 1 {
			fmt.Println("txReceipt.Status : ", txReceipt.Status)
			panic("TxTransaction fail !!!")
		}
		useGas := big.NewInt(0).Mul(big.NewInt(int64(txReceipt.CumulativeGasUsed)),big.NewInt(1000000000))
		fmt.Println("useGas: ", useGas)
		fmt.Println(reflect.TypeOf(useGas))
		fromBalanceSecondAfter := utils.GetBalance(fromAddress.String())
		fmt.Println("fromBalanceSecondAfter: ", fromBalanceSecondAfter)
		contractAddressBalanceAfter := utils.GetBalance(contractAddress.String())
		fmt.Println("contractAddressBalanceAfter: ", contractAddressBalanceAfter)
		managerBalanceSecondAfter := utils.GetBalance("0x0000000000000000000000000000000000000007")
		fmt.Println("managerBalanceSecondAfter: ", managerBalanceSecondAfter)
		fmt.Println(reflect.TypeOf(managerBalanceSecondAfter))
		//1 断言合约账户账户资金是否 = contractAddressBalance + amount
		So(contractAddressBalanceAfter.Cmp(contractAddressBalance.Add(contractAddressBalance, amount)) == 0, ShouldBeTrue)
		//2 断言治理合约地址是否正常获得消耗gas
		So(managerBalanceSecondAfter.Cmp(managerBalanceAfter.Add(managerBalanceAfter, useGas)) == 0, ShouldBeTrue)
		//3 from账户资金= fromBalanceAfter -fromBalanceSecondAfter -amount-useGas
		fmt.Println("fromBalanceAfter: ", fromBalanceAfter)
		//fmt.Println("After: ", fromBalanceSecondAfter.Add(fromBalanceSecondAfter, amount).Add(fromBalanceSecondAfter, useGas))
		So(fromBalanceAfter.Cmp(fromBalanceSecondAfter.Add(fromBalanceSecondAfter, amount).Add(fromBalanceSecondAfter, useGas)) == 0, ShouldBeTrue)

	})
}

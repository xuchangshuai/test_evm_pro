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

func TestEVMTransfer001(t *testing.T) {
	Convey("Test001 测试EIP155交易普通转账ong ", t, func() {
		url := "http://127.0.0.1:20339"
		toAddress := "0xf45505D1F482EBc8881dacA97B122B62771B9e1d"
		amount := big.NewInt(1000000000)
		//  备注： 对应的MetaMask钱包第一个账户私钥
		fromPrivateKey := "80a68081edc0aed4ddf8fa9f6a2e7cf8d0a69c998d4ef646f6446cbf4cfe9145"
		fromPrivate, _ := crypto.HexToECDSA(fromPrivateKey)
		fromAddress := crypto.PubkeyToAddress(fromPrivate.PublicKey)
		fmt.Println("fromAddress: ", fromAddress)

		//交易前账户余额
		fromBalance := utils.GetBalance(url, fromAddress.String())
		toBalance := utils.GetBalance(url, toAddress)
		managerBalance := utils.GetBalance(url, "0x0000000000000000000000000000000000000007")

		hash,err := deploy.SendTransfer(url, fromPrivateKey, toAddress, amount, uint64(200000))
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
}

func TestEVMTransfer002(t *testing.T) {
	Convey("Test002 测试EIP155交易普通转账ong(amount:255) ", t, func() {
		url := "http://127.0.0.1:20339"
		toAddress := "0xf45505D1F482EBc8881dacA97B122B62771B9e1d"
		amount := big.NewInt(255000000000)
		//  备注： 对应的MetaMask钱包第一个账户私钥
		fromPrivateKey := "80a68081edc0aed4ddf8fa9f6a2e7cf8d0a69c998d4ef646f6446cbf4cfe9145"
		fromPrivate, _ := crypto.HexToECDSA(fromPrivateKey)
		fromAddress := crypto.PubkeyToAddress(fromPrivate.PublicKey)

		//交易前账户余额
		fromBalance := utils.GetBalance(url, fromAddress.String())
		toBalance := utils.GetBalance(url, toAddress)
		managerBalance := utils.GetBalance(url, "0x0000000000000000000000000000000000000007")

		hash,err := deploy.SendTransfer(url, fromPrivateKey, toAddress, amount, uint64(200000))
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
}

func TestEVMTransfer003(t *testing.T) {
	Convey("Test003 测试EIP155交易普通转账ong(amount:256) ", t, func() {
		url := "http://127.0.0.1:20339"
		toAddress := "0xf45505D1F482EBc8881dacA97B122B62771B9e1d"
		amount := big.NewInt(256000000000)
		//  备注： 对应的MetaMask钱包第一个账户私钥
		fromPrivateKey := "80a68081edc0aed4ddf8fa9f6a2e7cf8d0a69c998d4ef646f6446cbf4cfe9145"
		fromPrivate, _ := crypto.HexToECDSA(fromPrivateKey)
		fromAddress := crypto.PubkeyToAddress(fromPrivate.PublicKey)

		//交易前账户余额
		fromBalance := utils.GetBalance(url, fromAddress.String())
		toBalance := utils.GetBalance(url, toAddress)
		managerBalance := utils.GetBalance(url, "0x0000000000000000000000000000000000000007")

		hash,err := deploy.SendTransfer(url, fromPrivateKey, toAddress, amount, uint64(200000))
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
}

func TestEVMTransfer004(t *testing.T) {
	Convey("Test004 测试EIP155交易普通转账ong(amount:257) ", t, func() {
		url := "http://127.0.0.1:20339"
		toAddress := "0xf45505D1F482EBc8881dacA97B122B62771B9e1d"
		amount := big.NewInt(257000000000)
		//  备注： 对应的MetaMask钱包第一个账户私钥
		fromPrivateKey := "80a68081edc0aed4ddf8fa9f6a2e7cf8d0a69c998d4ef646f6446cbf4cfe9145"
		fromPrivate, _ := crypto.HexToECDSA(fromPrivateKey)
		fromAddress := crypto.PubkeyToAddress(fromPrivate.PublicKey)

		//交易前账户余额
		fromBalance := utils.GetBalance(url, fromAddress.String())
		toBalance := utils.GetBalance(url, toAddress)
		managerBalance := utils.GetBalance(url, "0x0000000000000000000000000000000000000007")

		hash,err := deploy.SendTransfer(url, fromPrivateKey, toAddress, amount, uint64(200000))
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
}

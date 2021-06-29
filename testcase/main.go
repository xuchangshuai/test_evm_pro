package main

import (
	"fmt"
	"math/big"
	"test_evm/testcase/deploy"
	"test_evm/utils"
	"time"
)

func main() {
	url := "http://127.0.0.1:20339"
	toAddress := "0x5f9a233682e80e8bb1bfdf18cd21d82af2fd7edd"
	//交易前账户余额
	fromBalance := utils.GetBalance(url, "0x42195C051eafc0E328a5e8AED9bB604a9569DCBc")
	toBalance := utils.GetBalance(url, toAddress)
	managerBalance := utils.GetBalance(url, "0x0000000000000000000000000000000000000007")
	hash := deploy.SendTransfer(url, toAddress, big.NewInt(1000000000), uint64(200000))
	fmt.Println(hash)
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
	fmt.Println(status)
	if status != "0x1" {
		fmt.Println("status : ", status)
		panic("TxTransaction fail !!!")
	}
	useGas := utils.HexToDec(utils.Strval(receipt["cumulativeGasUsed"]))
	fmt.Println("useGas",useGas)

	fromBalanceAfter := utils.GetBalance(url, "0x42195C051eafc0E328a5e8AED9bB604a9569DCBc")
	toBalanceAfter := utils.GetBalance(url, toAddress)
	managerBalanceAfter := utils.GetBalance(url, "0x0000000000000000000000000000000000000007")
	fmt.Println("fromBalance", fromBalance)
	fmt.Println("toBalance", toBalance)
	fmt.Println("managerBalance", managerBalance)
	fmt.Println("fromBalanceAfter", fromBalanceAfter)
	fmt.Println("toBalanceAfter", toBalanceAfter)
	fmt.Println("managerBalanceAfter", managerBalanceAfter)
	fmt.Println(managerBalanceAfter.Sub(managerBalanceAfter,managerBalance))
}

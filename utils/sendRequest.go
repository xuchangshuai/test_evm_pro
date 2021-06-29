package utils

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
)

func GetTransferInfo(url string, hash common.Hash) map[string]interface{} {
	//url = "http://127.0.0.1:20339"
	client, _ := rpc.Dial(url)
	defer client.Close()
	_receipt := make(map[string]interface{})
	err := client.Call(&_receipt, "eth_getTransactionReceipt", hash)
	if err != nil {
		panic(err)
	}
	fmt.Println("eth_getTransactionReceipt res :", _receipt)
	status := HexToDec(Strval(_receipt["status"]))
	fmt.Println("eth_getTransactionReceipt status :", status)
	return _receipt
}

func GetBalance(url string, address string) *big.Int {
	//url = "http://127.0.0.1:20339"
	client, _ := rpc.Dial(url)
	defer client.Close()
	var balance string
	err := client.Call(&balance, "eth_getBalance", address, "latest")
	if err != nil {
		panic(err)
	}
	res := HexToDec(balance) //16 to 10
	return res
}

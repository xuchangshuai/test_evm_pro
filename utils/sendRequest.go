package utils

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
	"test_evm/config"
	"time"
)

func GetBalance(address string) *big.Int {
	client, _ := rpc.Dial(config.URL)
	defer client.Close()
	var balance string
	err := client.Call(&balance, "eth_getBalance", address, "latest")
	if err != nil {
		panic(err)
	}
	res := HexToDec(balance) //16 to 10
	return res
}

func GetTransferInfoByHash(hash common.Hash) (*types.Receipt, error) {
	var resErr error
	var res *types.Receipt
	for i := 0; i < 10; i++ {
		time.Sleep(3000000000) //设置10次请求，每次间隔3s 30s不落块超时 抛异常
		receipt, err := config.GetEthClient().TransactionReceipt(context.Background(), hash)
		if err == nil {
			resErr = nil
			res = receipt
			break
		} else {
			resErr = err
		}
	}
	return res, resErr
}

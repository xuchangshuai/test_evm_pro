package example

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"test_evm/testcase/deploy/inputCont"

	//"github.com/ontio/ontology/test/testcase/deploy/inputCont"
	"math/big"
)

func inputContInteraction() {

	// 构建一个私链对象
	ethClient, err := ethclient.Dial("http://127.0.0.1:20339")
	cont,err := inputCont.NewInputCont(ethCommon.HexToAddress("0x3596746b0a4d2e6a6a41d3cf846e5958d459554c"), ethClient)
	if err != nil {
		panic(err)
	}
	fmt.Println("outputCont :", cont)
	testPrivateKeyStr := "80a68081edc0aed4ddf8fa9f6a2e7cf8d0a69c998d4ef646f6446cbf4cfe9145" //
	testPrivateKey, err := crypto.HexToECDSA(testPrivateKeyStr)
	if err != nil {
		panic(err)
	}
	fmt.Println("28ln  testPrivateKey:", testPrivateKey)
	addr := crypto.PubkeyToAddress(testPrivateKey.PublicKey)
	fmt.Println("31ln address: ", addr)

	// 查nonce值
	nonce, err := ethClient.PendingNonceAt(context.Background(), addr)
	if err != nil {
		panic(err)
	}
	//查chainId
	chainId, _ := ethClient.ChainID(context.Background())
	//构建一个TransactOpts对象
	opts, err := bind.NewKeyedTransactorWithChainID(testPrivateKey, chainId)
	if err != nil {
		panic(err)
	}
	opts.From = addr
	opts.Nonce = big.NewInt(int64(nonce))
	opts.GasPrice = big.NewInt(500)
	opts.GasLimit = 300000
	opts.Value=big.NewInt(101000000000)

	res, err := cont.TransferAccounts(opts)
	if err != nil {
		panic(err)
	}

	fmt.Println("res :",res)
	fmt.Println("res: (evm)", res.Hash())
	//ha := common.Uint256(res.Hash())
	//fmt.Println("res: ", ha.ToHexString())
	//fmt.Println(chainId)
	////0x42195C051eafc0E328a5e8AED9bB604a9569DCBc  AMoNdAWjcrFj5t1GyAx2vKiM6BkCmHaRS7
	////0x0000000000000000000000000000000000000007  AFmseVrdL9f9oyCzZefL9tG6UbviEH9ugK    治理地址

}

package example

func main()  {
	
}
//
//import (
//	"context"
//	"fmt"
//	"github.com/ethereum/go-ethereum/accounts/abi/bind"
//	ethCommon "github.com/ethereum/go-ethereum/common"
//	"github.com/ethereum/go-ethereum/crypto"
//	"github.com/ethereum/go-ethereum/ethclient"
//	coin "test_evm/test"
//
//	//"github.com/ontio/ontology/common"
//	"math/big"
//	"time"
//)
//
//func main() {
//
//	// 构建一个私链对象
//	ethClient, err := ethclient.Dial("http://127.0.0.1:20339")
//	token, err := coin.NewCoin(ethCommon.HexToAddress("0x4e777507ce95770df9f7fa1961df8e1ea7fdfad1"), ethClient)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("token :", token)
//	testPrivateKeyStr := "80a68081edc0aed4ddf8fa9f6a2e7cf8d0a69c998d4ef646f6446cbf4cfe9145" //
//	testPrivateKey, err := crypto.HexToECDSA(testPrivateKeyStr)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("28ln  testPrivateKey:", testPrivateKey)
//	addr := crypto.PubkeyToAddress(testPrivateKey.PublicKey)
//	fmt.Println("31ln address: ", addr)
//
//	// 查nonce值
//	nonce, err := ethClient.PendingNonceAt(context.Background(), addr)
//	if err != nil {
//		panic(err)
//	}
//	//查chainId
//	chainId, _ := ethClient.ChainID(context.Background())
//	//构建一个TransactOpts对象
//	opts, err := bind.NewKeyedTransactorWithChainID(testPrivateKey, chainId)
//	if err != nil {
//		panic(err)
//	}
//	opts.From = addr
//	opts.Nonce = big.NewInt(int64(nonce))
//	opts.GasPrice = big.NewInt(500)
//	opts.GasLimit = 300000
//
//	res, err := token.Mint(opts, ethCommon.HexToAddress("0x319b6EcbBA71b6197c19E3159E7A69de74458def"), big.NewInt(50000))
//	if err != nil {
//		panic(err)
//	}
//
//	fmt.Println("res: (evm)", res.Hash())
//	//fmt.Println("res: ", common.Uint256(res.Hash()).ToHexString())
//
//	opt := bind.CallOpts{
//		From: addr,
//	}
//	time.Sleep(5000000)
//	mintAddress, err := token.Balances(&opt, ethCommon.HexToAddress("0x319b6EcbBA71b6197c19E3159E7A69de74458def"))
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("mintAddress :", mintAddress)
//	//fmt.Println(chainId)
//	////0x42195C051eafc0E328a5e8AED9bB604a9569DCBc  AMoNdAWjcrFj5t1GyAx2vKiM6BkCmHaRS7
//	////0x0000000000000000000000000000000000000007  AFmseVrdL9f9oyCzZefL9tG6UbviEH9ugK    治理地址
//
//}

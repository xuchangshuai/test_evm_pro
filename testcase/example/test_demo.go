package example

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"test_evm/config"
)

func ma111in() {
	//前置准备
	ethClient, err := ethclient.Dial("http://127.0.0.1:20339") //实例化当前链对象
	if err != nil {
		fmt.Println("ethclient.Dial failed: ", err.Error())
		return
	}
	chainId, err := ethClient.ChainID(context.Background()) //拿当前的链id
	if err != nil {
		panic(err)
	}
	fmt.Println("chainId : ", chainId)
	nonce, err := ethClient.PendingNonceAt(context.Background(), common.HexToAddress("0x42195C051eafc0E328a5e8AED9bB604a9569DCBc"))
	if err != nil {
		panic(err)
	}
	fmt.Println("nonce: ", nonce)
	//私钥
	key, err := crypto.HexToECDSA(config.FromPrivate)
	if err != nil {
		fmt.Println("crypto.HexToECDSA failed: ", err.Error())
		return
	}
	fmt.Println("address: ", key.PublicKey)

	//一、ABI编码请求参数  合约交互的时候处理  如果没有data 传nil
	methodId := crypto.Keccak256([]byte("mint()"))[:4] // 取函数原型的keccak256哈希的前四个字节
	fmt.Println("methodId: ", methodId)
	fmt.Println("methodId: ", common.Bytes2Hex(methodId))

	paramValue := math.U256Bytes(new(big.Int).Set(big.NewInt(0))) // 函数入参 转换成16进制的[]byte
	fmt.Println("paramValue: ", paramValue)
	fmt.Println("paramValue: ", common.Bytes2Hex(paramValue))
	input := append(methodId, paramValue...)
	fmt.Println("input: ", input)
	fmt.Println("input: ", common.Bytes2Hex(input))

	from := crypto.PubkeyToAddress(key.PublicKey)
	contractAddr := common.HexToAddress("0xDf1dAD556528f2BCe8511E47D607EeC740A1E3D7")
	msg := ethereum.CallMsg{
		From:     from,
		To:       &contractAddr,
		Gas:      30000,   // if 0, the call executes with near-infinite gas
		GasPrice: big.NewInt(300000), // wei <-> gas exchange ratio
		//Value:    nil, // amount of wei sent along with the call
		Data:     input,
	}
	res, err := ethClient.CallContract(context.Background(), msg, nil)
	fmt.Println(err)
	fmt.Println("res:", res)
	fmt.Println("res:", hex.EncodeToString(res))

	//// 二、构造交易对象
	//value := big.NewInt(0)
	//gasLimit := uint64(3000000)
	//gasPrice := big.NewInt(500)
	//rawTx := types.NewTransaction(nonce, common.HexToAddress("0x320fa3ed0fd7b026cd4317f3e73d1c1b2da78f58"), value, gasLimit, gasPrice, input)
	//jsonRawTx, _ := rawTx.MarshalJSON()
	//fmt.Println("rawTx: ", string(jsonRawTx))

	//// 三、交易签名
	//signer := types.NewEIP155Signer(chainId) //big.NewInt(1)//当前入参链id
	//
	//sigTransaction, err := types.SignTx(rawTx, signer, key)
	//if err != nil {
	//	fmt.Println("types.SignTx failed: ", err.Error())
	//	return
	//}
	//jsonSigTx, _ := sigTransaction.MarshalJSON()
	//fmt.Println("sigTransaction: ", string(jsonSigTx))
	//
	//// 四、发送交易
	//err = ethClient.SendTransaction(context.Background(), sigTransaction)
	//if err != nil {
	//	fmt.Println("ethClient.SendTransaction failed: ", err.Error())
	//	return
	//}
	//fmt.Println("send transaction success,tx: ", sigTransaction.Hash().Hex())
}

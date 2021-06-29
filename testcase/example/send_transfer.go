package example

import (
	"context"
	"fmt"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

func main() {
	fmt.Println("1")
	// 构建一个私链对象
	ethClient, err := ethclient.Dial("http://127.0.0.1:20339")
	fmt.Println("2")
	checkErrFromSendTransfer(err)
	//查chainId
	chainId, err := ethClient.ChainID(context.Background())
	fmt.Println(err)
	fmt.Println("23ln chainId: ", chainId)
	//0x42195C051eafc0E328a5e8AED9bB604a9569DCBc  AMoNdAWjcrFj5t1GyAx2vKiM6BkCmHaRS7
	//0x0000000000000000000000000000000000000007  AFmseVrdL9f9oyCzZefL9tG6UbviEH9ugK    治理地址
	testPrivateKeyStr := "80a68081edc0aed4ddf8fa9f6a2e7cf8d0a69c998d4ef646f6446cbf4cfe9145" //
	testPrivateKey, err := crypto.HexToECDSA(testPrivateKeyStr)
	checkErrFromSendTransfer(err)
	fmt.Println("28ln  testPrivateKey:", testPrivateKey)
	addr := crypto.PubkeyToAddress(testPrivateKey.PublicKey)
	fmt.Println("31ln address: ", addr)
	// 查nonce值
	nonce, err := ethClient.PendingNonceAt(context.Background(), addr)
	checkErrFromSendTransfer(err)
	fmt.Println("35ln nonce :", nonce)
	//私链第二个账户:AYxAW7xXTNQVPE44wtgopcmrXt9vBJCJrV  0xbc6c5e45ed2eb101b26eff56da0a0e8827ef78a9
	//合约地址  0x8291d9e5e6014af5b84f751416a3368071f75bf8  ATgGDMRBMRwStAxeK4A4s1pStGuMqHrJoS
	to := ethCommon.HexToAddress("0x8291d9e5e6014af5b84f751416a3368071f75bf8")
	gasPrice := big.NewInt(500)
	rawTx := types.NewTransaction(nonce, to, big.NewInt(257000000000), uint64(300000), gasPrice, nil)
	signer := types.NewEIP155Signer(chainId) //big.NewInt(1)//当前入参链id
	key, err := crypto.HexToECDSA(testPrivateKeyStr)
	if err != nil {
		fmt.Println("crypto.HexToECDSA failed: ", err.Error())
		return
	}
	//对交易对象做签名 0交易对象  1签名类型types.NewEIP155Signer(chainId)  2签名账户私钥
	sigTransaction, err := types.SignTx(rawTx, signer, key)
	if err != nil {
		fmt.Println("types.SignTx failed: ", err.Error())
		return
	}
	//buf, err := json.MarshalIndent(sigTransaction, "", "  ")
	//checkErrFromSendTransfer(err)
	fmt.Println("52ln rawTx: ", sigTransaction)
	err = ethClient.SendTransaction(context.Background(), sigTransaction)
	checkErrFromSendTransfer(err)
	fmt.Println("tx hash(evm): ", sigTransaction.Hash())
	//fmt.Println("tx hash: ",common.Uint256(sigTransaction.Hash()).ToHexString())

}

func checkErrFromSendTransfer(err error) {
	if err != nil {
		panic(err)
	}
}

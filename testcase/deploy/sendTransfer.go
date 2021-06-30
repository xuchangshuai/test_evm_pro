package deploy

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

func SendTransfer(url string, fromPrivateKey string, to string, amount *big.Int, gasLimit uint64) common.Hash {

	// 构建一个私链对象
	ethClient, err := ethclient.Dial(url)
	_checkErr(err)
	defer ethClient.Close()
	//查chainId
	chainId, err := ethClient.ChainID(context.Background())
	_checkErr(err)
	//0x42195C051eafc0E328a5e8AED9bB604a9569DCBc  AMoNdAWjcrFj5t1GyAx2vKiM6BkCmHaRS7
	//testPrivateKeyStr := "80a68081edc0aed4ddf8fa9f6a2e7cf8d0a69c998d4ef646f6446cbf4cfe9145" //
	testPrivateKey, err := crypto.HexToECDSA(fromPrivateKey)
	_checkErr(err)
	addr := crypto.PubkeyToAddress(testPrivateKey.PublicKey)
	// 查nonce值
	nonce, err := ethClient.PendingNonceAt(context.Background(), addr)
	_checkErr(err)
	fmt.Println("nonce: ", nonce)
	_to := common.HexToAddress(to)
	gasPrice := big.NewInt(500)
	//amount := big.NewInt(257000000000)
	rawTx := types.NewTransaction(nonce, _to, amount, gasLimit, gasPrice, nil)
	signer := types.NewEIP155Signer(chainId) //big.NewInt(1)//当前入参链id
	key, err := crypto.HexToECDSA(fromPrivateKey)
	if err != nil {
		fmt.Println("crypto.HexToECDSA failed: ", err.Error())
		panic(err)
	}
	//对交易对象做签名 0交易对象  1签名类型types.NewEIP155Signer(chainId)  2签名账户私钥
	sigTransaction, err := types.SignTx(rawTx, signer, key)
	if err != nil {
		fmt.Println("types.SignTx failed: ", err.Error())
		panic(err)
	}
	//fmt.Println("52ln rawTx: ", sigTransaction)
	err = ethClient.SendTransaction(context.Background(), sigTransaction)
	_checkErr(err)
	fmt.Println("tx hash(evm): ", sigTransaction.Hash())
	return sigTransaction.Hash()
}

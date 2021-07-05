package deploy

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
	"test_evm/config"
)

func SendTransfer(fromPrivateKey string, to string, amount *big.Int, gasLimit uint64) (common.Hash, error) {

	// 构建一个私链对象
	ethClient := config.GetEthClient()
	//查chainId
	chainId, err := ethClient.ChainID(context.Background())
	if err != nil {
		return common.Hash{}, err
	}
	//0x42195C051eafc0E328a5e8AED9bB604a9569DCBc  AMoNdAWjcrFj5t1GyAx2vKiM6BkCmHaRS7
	//testPrivateKeyStr := config.FromPrivate //
	testPrivateKey, err := crypto.HexToECDSA(fromPrivateKey)
	if err != nil {
		return common.Hash{}, err
	}
	addr := crypto.PubkeyToAddress(testPrivateKey.PublicKey)
	// 查nonce值
	nonce, err := ethClient.PendingNonceAt(context.Background(), addr)
	if err != nil {
		return common.Hash{}, err
	}
	log.Printf("nonce: %d", nonce)
	_to := common.HexToAddress(to)
	gasPrice := big.NewInt(500)
	//amount := big.NewInt(257000000000)
	rawTx := types.NewTransaction(nonce, _to, amount, gasLimit, gasPrice, nil)
	signer := types.NewEIP155Signer(chainId) //big.NewInt(1)//当前入参链id
	key, err := crypto.HexToECDSA(fromPrivateKey)
	if err != nil {
		return common.Hash{}, err
	}
	//对交易对象做签名 0交易对象  1签名类型types.NewEIP155Signer(chainId)  2签名账户私钥
	sigTransaction, err := types.SignTx(rawTx, signer, key)
	if err != nil {
		return common.Hash{}, err
	}
	//fmt.Println("52ln rawTx: ", sigTransaction)
	err = ethClient.SendTransaction(context.Background(), sigTransaction)
	if err != nil {
		return common.Hash{}, err
	}
	log.Printf("tx hash(evm): %s", sigTransaction.Hash())
	return sigTransaction.Hash(), nil
}

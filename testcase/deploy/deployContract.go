package deploy

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"strings"
	"test_evm/config"
)

func ContractDeploy(bin string, abi string, params ...interface{}) ethCommon.Hash {

	fmt.Println("bin: ", bin)
	fmt.Println("abi: ", abi)

	// 构建一个私链对象
	ethClient := config.GetEthClient()
	//查chainId
	chainId, _ := ethClient.ChainID(context.Background())
	//fmt.Println("23ln chainId: ", chainId)
	//0x42195C051eafc0E328a5e8AED9bB604a9569DCBc  AMoNdAWjcrFj5t1GyAx2vKiM6BkCmHaRS7
	//0x0000000000000000000000000000000000000007  AFmseVrdL9f9oyCzZefL9tG6UbviEH9ugK    治理地址
	testPrivateKeyStr := config.FromPrivate //私钥 账户0x4219的私钥
	testPrivateKey, err := crypto.HexToECDSA(testPrivateKeyStr)
	_checkErr(err)
	//fmt.Println("28ln  testPrivateKey:", testPrivateKey)
	addr := crypto.PubkeyToAddress(testPrivateKey.PublicKey)
	//ontAddr := common.Address(addr)
	//fmt.Println("31ln address: ", addr, ontAddr.ToBase58())
	// 查nonce值
	nonce, err := ethClient.PendingNonceAt(context.Background(), addr)
	_checkErr(err)
	//fmt.Println("35ln nonce :", nonce)
	//to := ethCommon.HexToAddress("0xbc6c5e45ed2eb101b26eff56da0a0e8827ef78a9") //私链第二个账户:AYxAW7xXTNQVPE44wtgopcmrXt9vBJCJrV
	//contractAddress = ethCommon.HexToAddress("0xf218394768660e1ae1e12f9f881b73e0c6125a7e")
	// 合约文件
	opts, err := bind.NewKeyedTransactorWithChainID(testPrivateKey, chainId)
	_checkErr(err)
	opts.Nonce = big.NewInt(int64(nonce))
	opts.GasPrice = big.NewInt(500)
	opts.GasLimit = 3000000
	binDecode, _ := hex.DecodeString(bin)
	//fmt.Println("binDecode: ", binDecode)
	rawTx, err := _newDeployEvmContract(opts, binDecode, abi, params...)
	//fmt.Println("rawTx :", rawTx)
	_checkErr(err)
	err = ethClient.SendTransaction(context.Background(), rawTx)
	_checkErr(err)
	//fmt.Println("tx hash: ", rawTx.Hash())
	return rawTx.Hash()

}
func _newDeployEvmContract(opts *bind.TransactOpts, code []byte, jsonABI string, params ...interface{}) (*types.Transaction, error) {
	fmt.Println("params :", params)
	parsed, err := abi.JSON(strings.NewReader(jsonABI))
	_checkErr(err)
	input, err := parsed.Pack("", params...)
	//fmt.Println("input :", input)
	_checkErr(err)
	input = append(code, input...)
	deployTx := types.NewContractCreation(opts.Nonce.Uint64(), opts.Value, opts.GasLimit, opts.GasPrice, input)
	signedTx, err := opts.Signer(opts.From, deployTx)
	_checkErr(err)
	return signedTx, err
}

func _checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

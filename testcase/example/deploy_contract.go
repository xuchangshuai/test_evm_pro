package example

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ontio/ontology/common"
	"math/big"
	"strings"
)

func test1() {
	// 构建一个私链对象
	ethClient, err := ethclient.Dial("http://127.0.0.1:20339")
	defer ethClient.Close()
	checkErrFromDeployContract(err)
	//查chainId
	chainId, _ := ethClient.ChainID(context.Background())
	fmt.Println("23ln chainId: ", chainId)
	//0x42195C051eafc0E328a5e8AED9bB604a9569DCBc  AMoNdAWjcrFj5t1GyAx2vKiM6BkCmHaRS7
	//0x0000000000000000000000000000000000000007  AFmseVrdL9f9oyCzZefL9tG6UbviEH9ugK    治理地址
	testPrivateKeyStr := "80a68081edc0aed4ddf8fa9f6a2e7cf8d0a69c998d4ef646f6446cbf4cfe9145" //私钥 账户0x4219的私钥
	testPrivateKey, err := crypto.HexToECDSA(testPrivateKeyStr)
	checkErrFromDeployContract(err)
	fmt.Println("28ln  testPrivateKey:", testPrivateKey)
	addr := crypto.PubkeyToAddress(testPrivateKey.PublicKey)
	ontAddr := common.Address(addr)
	fmt.Println("31ln address: ", addr, ontAddr.ToBase58())
	// 查nonce值
	nonce, err := ethClient.PendingNonceAt(context.Background(), addr)
	checkErrFromDeployContract(err)
	fmt.Println("35ln nonce :", nonce)
	//to := ethCommon.HexToAddress("0xbc6c5e45ed2eb101b26eff56da0a0e8827ef78a9") //私链第二个账户:AYxAW7xXTNQVPE44wtgopcmrXt9vBJCJrV

	// 合约文件
	opts, err := bind.NewKeyedTransactorWithChainID(testPrivateKey, chainId)
	checkErrFromDeployContract(err)
	opts.Nonce = big.NewInt(int64(nonce))
	opts.GasPrice = big.NewInt(500)
	opts.GasLimit = 3000000
	// bin文件
	bin := "60806040523480156100115760006000fd5b505b33600060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b61005a565b6105e0806100696000396000f3fe60806040523480156100115760006000fd5b50600436106100515760003560e01c8063075461721461005757806327e235e31461007557806340c10f19146100a5578063d0679d34146100c157610051565b60006000fd5b61005f6100dd565b60405161006c91906103ec565b60405180910390f35b61008f600480360381019061008a919061035c565b610103565b60405161009c9190610440565b60405180910390f35b6100bf60048036038101906100ba919061038a565b61011e565b005b6100db60048036038101906100d6919061038a565b6101dd565b005b600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60016000506020528060005260406000206000915090505481565b600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561017b5760006000fd5b80600160005060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082828250546101cf9190610467565b9250508190909055505b5050565b80600160005060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005054101515156102325760006000fd5b80600160005060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828282505461028691906104be565b92505081909090555080600160005060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082828250546102e39190610467565b9250508190909055507f3990db2d31862302a685e8086b5755072a6e2b5b780af1ee81ece35ee3cd334533838360405161031f93929190610408565b60405180910390a15b5050566105a9565b60008135905061033f81610573565b5b92915050565b6000813590506103558161058e565b5b92915050565b6000602082840312156103725761037161056b565b5b600061038084828501610330565b9150505b92915050565b60006000604083850312156103a2576103a161056b565b5b60006103b085828601610330565b92505060206103c185828601610346565b9150505b9250929050565b6103d5816104f3565b825250505b565b6103e581610527565b825250505b565b600060208201905061040160008301846103cc565b5b92915050565b600060608201905061041d60008301866103cc565b61042a60208301856103cc565b61043760408301846103dc565b5b949350505050565b600060208201905061045560008301846103dc565b5b92915050565b600060405190505b90565b600061047282610527565b915061047d83610527565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156104b2576104b1610532565b5b82820190505b92915050565b60006104c982610527565b91506104d483610527565b9250828210156104e7576104e6610532565b5b82820390505b92915050565b60006104fe82610506565b90505b919050565b600073ffffffffffffffffffffffffffffffffffffffff821690505b919050565b60008190505b919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b565b60006000fd5b565b60006000fd5b565b61057c816104f3565b8114151561058a5760006000fd5b505b565b61059781610527565b811415156105a55760006000fd5b505b565bfea26469706673582212205f6bb7ac9215ba469d0cf773fd7d8014b26bee11723dc2e2c5fd830eadc67ed464736f6c63430008050033"
	// bi
	contractAbi:= "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Sent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receive1\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receive2\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
	binDecode, _ := hex.DecodeString(bin)
	//fmt.Println("binDecode: ", binDecode)

	rawTx, err := NewDeployEvmContract(opts, binDecode, contractAbi)
	//fmt.Println("rawTx :", rawTx)
	checkErrFromDeployContract(err)


	err = ethClient.SendTransaction(context.Background(), rawTx)
	checkErrFromDeployContract(err)
	fmt.Println("tx hash: ", rawTx.Hash())
	//320fa3ed0fd7b026cd4317f3e73d1c1b2da78f58 coin合约地址  ALLaAoKbMHTifiL46DhwbpSvqTXJ6H9vJW

}
func NewDeployEvmContract(opts *bind.TransactOpts, code []byte, jsonABI string, params ...interface{}) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(jsonABI))
	checkErrFromDeployContract(err)
	input, err := parsed.Pack("", params...)
	checkErrFromDeployContract(err)
	input = append(code, input...)
	deployTx := types.NewContractCreation(opts.Nonce.Uint64(), opts.Value, opts.GasLimit, opts.GasPrice, input)
	signedTx, err := opts.Signer(opts.From, deployTx)
	checkErrFromDeployContract(err)
	return signedTx, err
}

func checkErrFromDeployContract(err error)  {
	if err != nil {
		panic(err)
	}
}

//signer := types.NewEIP155Signer(chainId) //big.NewInt(1)//当前入参链id
//key, err := crypto.HexToECDSA(testPrivateKeyStr)
//if err != nil {
//	fmt.Println("crypto.HexToECDSA failed: ", err.Error())
//	return
//}
//对交易对象做签名 0交易对象  1签名类型types.NewEIP155Signer(chainId)  2签名账户私钥
//sigTransaction, err := types.SignTx(rawTx, signer, key)
//if err != nil {
//	fmt.Println("types.SignTx failed: ", err.Error())
//	return
//}
//buf, err := json.MarshalIndent(sigTransaction, "", "  ")
//checkErrFromDeployContract(err)
//fmt.Println("52ln rawTx: ", buf)
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
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strings"
)

// ContractDeploy url = "http://127.0.0.1:20339"
// bin文件
//ContractDeploy bin = "60806040523480156100115760006000fd5b5060405161041d38038061041d83398181016040528101906100339190610096565b5b80600060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b5061012f5661012e565b60008151905061008f81610113565b5b92915050565b6000602082840312156100ac576100ab61010b565b5b60006100ba84828501610080565b9150505b92915050565b600060405190505b90565b60006100da826100e2565b90505b919050565b600073ffffffffffffffffffffffffffffffffffffffff821690505b919050565b60006000fd5b565b60006000fd5b565b61011c816100cf565b8114151561012a5760006000fd5b505b565b5b6102df8061013e6000396000f3fe6080604052600436106100385760003560e01c806312065fe01461003e5780632372cdde1461006a578063ca28cda41461008857610038565b60006000fd5b34801561004b5760006000fd5b506100546100b4565b60405161006191906101da565b60405180910390f35b6100726100c1565b60405161007f91906101be565b60405180910390f35b3480156100955760006000fd5b5061009e610148565b6040516100ab91906101a2565b60405180910390f35b60004790506100be565b90565b6000600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc633b9aca003461010f91906101f6565b9081150290604051600060405180830381858888f1935050505015801561013b573d600060003e3d6000fd5b5060019050610145565b90565b600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681566102a8565b61017b8161022b565b825250505b565b61018b8161023e565b825250505b565b61019b8161026c565b825250505b565b60006020820190506101b76000830184610172565b5b92915050565b60006020820190506101d36000830184610182565b5b92915050565b60006020820190506101ef6000830184610192565b5b92915050565b60006102018261026c565b915061020c8361026c565b92508282101561021f5761021e610277565b5b82820390505b92915050565b60006102368261024b565b90505b919050565b600081151590505b919050565b600073ffffffffffffffffffffffffffffffffffffffff821690505b919050565b60008190505b919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b565bfea26469706673582212207de6707531b56dca7ce580692d53349fd9c66b2646c3ce31ec2e16cb66166ab464736f6c63430008050033"
// bi
//ContractDeploy abi = "[{\"inputs\":[{\"internalType\":\"address payable\",\"name\":\"outputAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"_outputAddress\",\"outputs\":[{\"internalType\":\"address payable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transferAccounts\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]"
func ContractDeploy(url string, bin string, abi string, params ...interface{}) ethCommon.Hash {

	fmt.Println("bin: ",bin)
	fmt.Println("abi: ",abi)

	// 构建一个私链对象
	ethClient, err := ethclient.Dial(url)
	defer ethClient.Close()
	_checkErr(err)
	//查chainId
	chainId, _ := ethClient.ChainID(context.Background())
	//fmt.Println("23ln chainId: ", chainId)
	//0x42195C051eafc0E328a5e8AED9bB604a9569DCBc  AMoNdAWjcrFj5t1GyAx2vKiM6BkCmHaRS7
	//0x0000000000000000000000000000000000000007  AFmseVrdL9f9oyCzZefL9tG6UbviEH9ugK    治理地址
	testPrivateKeyStr := "80a68081edc0aed4ddf8fa9f6a2e7cf8d0a69c998d4ef646f6446cbf4cfe9145" //私钥 账户0x4219的私钥
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

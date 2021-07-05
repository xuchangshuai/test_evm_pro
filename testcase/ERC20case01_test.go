package testcase

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/smartystreets/goconvey/convey"
	"log"
	"math/big"
	"test_evm/config"
	"test_evm/testcase/deploy"
	ERC20Test "test_evm/testcase/deploy/testEPC20"
	"test_evm/utils"
	"testing"
)

func TestERC20_01(t *testing.T) {
	Convey("TestERC20", t, func() {
		//ethClient:=config.GetEthClient()
		//chainId, err := ethClient.ChainID(context.Background())
		//_checkErr(err)

		// 备注: 对应的MetaMask钱包第一个账户私钥
		fromPrivateKey := config.FromPrivate
		fromPrivate, _ := crypto.HexToECDSA(fromPrivateKey)
		fromAddress := crypto.PubkeyToAddress(fromPrivate.PublicKey)
		log.Printf("fromAddress: %s", fromAddress.String())

		erc20Abi, err := utils.ReadAll("deploy/testEPC20/EPC20_example_sol_TokenERC20.abi")
		_checkErr(err)
		erc20Bin, err := utils.ReadAll("deploy/testEPC20/EPC20_example_sol_TokenERC20.bin")
		_checkErr(err)

		erc20ContractHash := deploy.ContractDeploy(string(erc20Bin), string(erc20Abi), big.NewInt(10000000), "牛币", "NBB")
		log.Printf("erc20ContractHash: %s", erc20ContractHash.String())

		erc20ContractReceipt, err := utils.GetTransferInfoByHash(erc20ContractHash)
		_checkErr(err)
		if erc20ContractReceipt.Status != 1 {
			fmt.Println("erc20ContractReceipt.Status : ", erc20ContractReceipt.Status)
			panic("TxTransaction fail !!!")
		}
		//erc20ContractAddress合约地址
		erc20ContractAddress := erc20ContractReceipt.ContractAddress
		log.Printf("erc20ContractAddress: %s", erc20ContractAddress.String())

		ethClient := config.GetEthClient()
		//chainId, err := ethClient.ChainID(context.Background())
		//_checkErr(err)
		//nonce, err := ethClient.PendingNonceAt(context.Background(), fromAddress)
		//_checkErr(err)

		epc20Contract, err := ERC20Test.NewERC20Test(erc20ContractAddress, ethClient)
		_checkErr(err)
		//opts, err := bind.NewKeyedTransactorWithChainID(fromPrivate, chainId)
		//_checkErr(err)
		//opts.Nonce = big.NewInt(int64(nonce))
		//opts.GasPrice = big.NewInt(500)
		//opts.GasLimit = 3000000
		erp20Name, err := epc20Contract.Name(&bind.CallOpts{
			From: fromAddress,
		})
		log.Printf("erp20Name: %s", erp20Name)
		erp20Symbol, err := epc20Contract.Symbol(&bind.CallOpts{
			From: fromAddress,
		})
		log.Printf("erp20Symbol: %s", erp20Symbol)

	})
}

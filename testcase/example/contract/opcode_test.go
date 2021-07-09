package contract

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/smartystreets/goconvey/convey"
	"log"
	"math/big"
	"test_evm/config"
	"test_evm/testcase/deploy"
	"test_evm/utils"
	"testing"
)

func TestOpCode(t *testing.T) {
	Convey("Test OpCode", t, func() {
		testContractAbi, err := utils.ReadAll("TestOpCodeExam_sol_simple.abi")
		checkErr(err)
		testContractBin, err := utils.ReadAll("TestOpCodeExam_sol_simple.bin")
		checkErr(err)
		hash := deploy.ContractDeploy(string(testContractBin), string(testContractAbi))
		receipt, err := utils.GetTransferInfoByHash(hash)
		checkErr(err)
		if receipt.Status != 1 {
			panic("tx error!")
		}
		testContractAddress := receipt.ContractAddress
		log.Printf("testContractAddress: %s", testContractAddress)

		opCodeContract, err := NewOpcode(testContractAddress, config.GetEthClient())
		checkErr(err)
		private, err := crypto.HexToECDSA(config.FromPrivate)
		checkErr(err)
		chainId, err := config.GetEthClient().ChainID(context.Background())
		opts, err := bind.NewKeyedTransactorWithChainID(private, chainId)
		checkErr(err)
		nonce, err := config.GetEthClient().PendingNonceAt(context.Background(), crypto.PubkeyToAddress(private.PublicKey))
		checkErr(err)
		opts.Value = big.NewInt(1)
		opts.Nonce = big.NewInt(int64(nonce))
		opts.GasPrice = big.NewInt(500)
		opts.GasLimit = 3000000
		tx, err := opCodeContract.Add(opts, big.NewInt(1))
		So(err, ShouldBeNil)
		res, err := utils.GetTransferInfoByHash(tx.Hash())
		checkErr(err)
		if res.Status != 1 {
			panic("tx fail!")
		}
		log.Printf("gasUsed %d", res.CumulativeGasUsed)
		a:=big.NewInt(int64(res.CumulativeGasUsed)).Sub(big.NewInt(int64(res.CumulativeGasUsed)),big.NewInt(500*21000))
		log.Printf("%d",a)
		So(big.NewInt(int64(res.CumulativeGasUsed)).Cmp(big.NewInt(500*21000)) > 0, ShouldBeTrue)

	})
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

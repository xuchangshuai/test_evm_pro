package testcase

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/smartystreets/goconvey/convey"
	"log"
	"math/big"
	"test_evm/config"
	"test_evm/testcase/deploy"
	TestOpCode "test_evm/testcase/deploy/testOpCode"
	"test_evm/utils"
	"testing"
)

func TestOpCode_(t *testing.T) {
	Convey("Test OpCode", t, func() {
		testContractAbi, err := utils.ReadAll("deploy/TestOpCode/TestOpCode_sol_Test.abi")
		checkErr(err)
		testContractBin, err := utils.ReadAll("deploy/TestOpCode/TestOpCode_sol_Test.bin")
		checkErr(err)
		hash := deploy.ContractDeploy(string(testContractBin), string(testContractAbi))
		receipt, err := utils.GetTransferInfoByHash(hash)
		checkErr(err)
		if receipt.Status != 1 {
			panic("tx error!")
		}
		testContractAddress := receipt.ContractAddress
		log.Printf("testContractAddress: %s", testContractAddress)
		testOpCodeContract, err := TestOpCode.NewTestOpCode(testContractAddress, config.GetEthClient())
		checkErr(err)
		callOpts := &bind.CallOpts{From: common.HexToAddress(config.GetPrivateList()[1])}
		err = testOpCodeContract.Add(callOpts)
		So(err, ShouldBeNil)
		err = testOpCodeContract.Sub(callOpts)
		So(err, ShouldBeNil)
		err = testOpCodeContract.Mul(callOpts)
		So(err, ShouldBeNil)
		err = testOpCodeContract.Div(callOpts)
		So(err, ShouldBeNil)
		err = testOpCodeContract.Sdiv(callOpts)
		So(err, ShouldBeNil)
		err = testOpCodeContract.Mod(callOpts)
		So(err, ShouldBeNil)
		err = testOpCodeContract.Smod(callOpts)
		So(err, ShouldBeNil)
		err = testOpCodeContract.Exp(callOpts)
		So(err, ShouldBeNil)
		err = testOpCodeContract.And(callOpts)
		So(err, ShouldBeNil)
		err = testOpCodeContract.Or(callOpts)
		So(err, ShouldBeNil)
		err = testOpCodeContract.Xor(callOpts)
		So(err, ShouldBeNil)
		err = testOpCodeContract.Not(callOpts)
		So(err, ShouldBeNil)
		err = testOpCodeContract.ByteFun(callOpts)
		So(err, ShouldBeNil)
		err = testOpCodeContract.Lt(callOpts)
		So(err, ShouldBeNil)
		err = testOpCodeContract.Gt(callOpts)
		So(err, ShouldBeNil)
		err = testOpCodeContract.Slt(callOpts)
		So(err, ShouldBeNil)
		err = testOpCodeContract.Sgt(callOpts)
		So(err, ShouldBeNil)
		err = testOpCodeContract.Eq(callOpts)
		So(err, ShouldBeNil)
	})
	Convey("Test OpCode2", t, func() {
		testContractAbi, err := utils.ReadAll("deploy/testOpCode/TestOpCode2_sol_TestCoin.abi")
		checkErr(err)
		testContractBin, err := utils.ReadAll("deploy/testOpCode/TestOpCode2_sol_TestCoin.bin")
		checkErr(err)
		hash := deploy.ContractDeploy(string(testContractBin), string(testContractAbi))
		receipt, err := utils.GetTransferInfoByHash(hash)
		checkErr(err)
		if receipt.Status != 1 {
			panic("tx error!")
		}
		testContractAddress := receipt.ContractAddress
		log.Printf("testContractAddress: %s", testContractAddress)
		testOpCodeContract, err := TestOpCode.NewTestOpCode2(testContractAddress, config.GetEthClient())
		checkErr(err)
		callOpts := &bind.CallOpts{From: common.HexToAddress(config.GetPrivateList()[1])}
		_, err = testOpCodeContract.GetAddress(callOpts)
		So(err, ShouldBeNil)
		_, err = testOpCodeContract.GetBalance(callOpts, config.GetAddressList()[1])
		So(err, ShouldBeNil)
		_, err = testOpCodeContract.GetOrigin(callOpts)
		So(err, ShouldBeNil)
		_, err = testOpCodeContract.GetCaller(callOpts)
		So(err, ShouldBeNil)

		private, err := crypto.HexToECDSA(config.FromPrivate)
		checkErr(err)
		chainId, err := config.GetEthClient().ChainID(context.Background())
		opts, err := bind.NewKeyedTransactorWithChainID(private, chainId)
		_checkErr(err)
		nonce, err := config.GetEthClient().PendingNonceAt(context.Background(), crypto.PubkeyToAddress(private.PublicKey))
		_checkErr(err)
		opts.Value = big.NewInt(1)
		opts.Nonce = big.NewInt(int64(nonce))
		opts.GasPrice = big.NewInt(500)
		opts.GasLimit = 3000000
		_, err = testOpCodeContract.GetValue(opts)
		So(err, ShouldBeNil)

		_, err = testOpCodeContract.GetData(callOpts)
		So(err, ShouldBeNil)
		_, err = testOpCodeContract.GetCallDataCopy(callOpts)
		So(err, ShouldBeNil)
		_, err = testOpCodeContract.GetGasPrice(callOpts)
		So(err, ShouldBeNil)
		_, err = testOpCodeContract.GetBlockHash(callOpts, big.NewInt(1))
		So(err, ShouldBeNil)
		_, err = testOpCodeContract.GetCoinBase(callOpts)
		So(err, ShouldBeNil)
		_, err = testOpCodeContract.GetTimestamp(callOpts)
		So(err, ShouldBeNil)
		_, err = testOpCodeContract.GetBlockNumber(callOpts)
		So(err, ShouldBeNil)
		_, err = testOpCodeContract.GetBlockDifficulty(callOpts)
		So(err, ShouldBeNil)
		err = testOpCodeContract.Pop(callOpts)
		So(err, ShouldBeNil)
		_, err = testOpCodeContract.Mload(callOpts)
		So(err, ShouldBeNil)
		err = testOpCodeContract.Mstore(callOpts)
		So(err, ShouldBeNil)
	})
	Convey("Test OpCode3", t, func() {
		testContractAbi, err := utils.ReadAll("deploy/testOpCode/TestOpCode3_sol_testStorege.abi")
		checkErr(err)
		testContractBin, err := utils.ReadAll("deploy/testOpCode/TestOpCode3_sol_testStorege.bin")
		checkErr(err)
		hash := deploy.ContractDeploy(string(testContractBin), string(testContractAbi))
		receipt, err := utils.GetTransferInfoByHash(hash)
		checkErr(err)
		if receipt.Status != 1 {
			panic("tx error!")
		}
		testContractAddress := receipt.ContractAddress
		log.Printf("testContractAddress: %s", testContractAddress)
		testOpCodeContract, err := TestOpCode.NewTestOpCode3(testContractAddress, config.GetEthClient())
		checkErr(err)
		callOpts := &bind.CallOpts{From: common.HexToAddress(config.GetPrivateList()[1])}
		err = testOpCodeContract.SloadAndSstore(callOpts)
		So(err, ShouldBeNil)
		err = testOpCodeContract.Jump(callOpts)
		So(err, ShouldBeNil)
		err = testOpCodeContract.Push(callOpts)
		So(err, ShouldBeNil)
		err = testOpCodeContract.Dup(callOpts)
		So(err, ShouldBeNil)
		err = testOpCodeContract.Swap(callOpts)
		So(err, ShouldBeNil)
	})
	Convey("Test OpCode5", t, func() {
		testContract4Abi, err := utils.ReadAll("deploy/testOpCode/TestOpCode4_sol_Person.abi")
		checkErr(err)
		testContract4Bin, err := utils.ReadAll("deploy/testOpCode/TestOpCode4_sol_Person.bin")
		checkErr(err)
		txhash1 := deploy.ContractDeploy(string(testContract4Bin), string(testContract4Abi))
		receipt1, err := utils.GetTransferInfoByHash(txhash1)
		checkErr(err)
		if receipt1.Status != 1 {
			panic("tx error!")
		}
		testContract4Address := receipt1.ContractAddress
		log.Printf("testContract4Address: %s",testContract4Address)

		testContractAbi, err := utils.ReadAll("deploy/testOpCode/TestOpCode5_sol_CallTest.abi")
		checkErr(err)
		testContractBin, err := utils.ReadAll("deploy/testOpCode/TestOpCode5_sol_CallTest.bin")
		checkErr(err)
		hash2 := deploy.ContractDeploy(string(testContractBin), string(testContractAbi))
		receipt, err := utils.GetTransferInfoByHash(hash2)
		checkErr(err)
		if receipt.Status != 1 {
			panic("tx error!")
		}
		testContractAddress := receipt.ContractAddress
		log.Printf("testContractAddress: %s", testContractAddress)
		testOpCodeContract, err := TestOpCode.NewTestOpCode5(testContractAddress, config.GetEthClient())
		checkErr(err)
		private, err := crypto.HexToECDSA(config.FromPrivate)
		checkErr(err)
		chainId, err := config.GetEthClient().ChainID(context.Background())
		opts, err := bind.NewKeyedTransactorWithChainID(private, chainId)
		checkErr(err)
		nonce, err := config.GetEthClient().PendingNonceAt(context.Background(), crypto.PubkeyToAddress(private.PublicKey))
		checkErr(err)
		opts.Value = big.NewInt(0)
		opts.Nonce = big.NewInt(int64(nonce))
		opts.GasPrice = big.NewInt(500)
		opts.GasLimit = 3000000
		tx, err := testOpCodeContract.CallByFun(opts, testContract4Address)
		So(err, ShouldBeNil)
		res, err := utils.GetTransferInfoByHash(tx.Hash())
		checkErr(err)
		if res.Status != 1 {
			panic("tx fail!")
		}
	})
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

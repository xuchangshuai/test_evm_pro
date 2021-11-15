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
		addRes, err := testOpCodeContract.Add(callOpts, big.NewInt(500), big.NewInt(80000))
		checkErr(err)
		log.Printf("addRes: %d", addRes)
		So(addRes.Cmp(big.NewInt(80500)) == 0, ShouldBeTrue)
		subRes, err := testOpCodeContract.Sub(callOpts, big.NewInt(500), big.NewInt(80000))
		So(err, ShouldNotBeNil)
		subRes, err = testOpCodeContract.Sub(callOpts, big.NewInt(500000), big.NewInt(80000))
		checkErr(err)
		log.Printf("subRes: %d", subRes)
		So(subRes.Cmp(big.NewInt(420000)) == 0, ShouldBeTrue)
		mulRes, err := testOpCodeContract.Mul(callOpts, big.NewInt(5), big.NewInt(80000))
		checkErr(err)
		So(mulRes.Cmp(big.NewInt(400000)) == 0, ShouldBeTrue)
		divRes, err := testOpCodeContract.Div(callOpts, big.NewInt(2), big.NewInt(8))
		log.Printf("2/8: %d", divRes)
		checkErr(err)
		So(divRes.Cmp(big.NewInt(0)) == 0, ShouldBeTrue)
		divRes1, err := testOpCodeContract.Div(callOpts, big.NewInt(80007), big.NewInt(8))
		checkErr(err)
		So(divRes1.Cmp(big.NewInt(10000)) == 0, ShouldBeTrue)
		sdivRes, err := testOpCodeContract.Sdiv(callOpts, big.NewInt(46465), big.NewInt(5955))
		checkErr(err)
		So(sdivRes.Cmp(big.NewInt(0).Div(big.NewInt(46465), big.NewInt(5955))) == 0, ShouldBeTrue)
		modRes, err := testOpCodeContract.Mod(callOpts, big.NewInt(887), big.NewInt(8))
		_checkErr(err)
		So(modRes.Cmp(big.NewInt(0).Mod(big.NewInt(887), big.NewInt(8))) == 0, ShouldBeTrue)
		smodRes, err := testOpCodeContract.Smod(callOpts, big.NewInt(95), big.NewInt(665))
		checkErr(err)
		So(smodRes.Cmp(big.NewInt(0).Mod(big.NewInt(95), big.NewInt(665))) == 0, ShouldBeTrue)
		expRes, err := testOpCodeContract.Exp(callOpts, big.NewInt(2), big.NewInt(8))
		checkErr(err)
		So(expRes.Cmp(big.NewInt(256)) == 0, ShouldBeTrue)

		ltRes, err := testOpCodeContract.Lt(callOpts, big.NewInt(888), big.NewInt(999))
		So(ltRes, ShouldBeTrue)
		ltRes, err = testOpCodeContract.Lt(callOpts, big.NewInt(55959), big.NewInt(999))
		So(ltRes, ShouldBeFalse)
		gtRes, err := testOpCodeContract.Gt(callOpts, big.NewInt(888), big.NewInt(999))
		So(gtRes, ShouldBeFalse)
		gtRes, err = testOpCodeContract.Gt(callOpts, big.NewInt(55959), big.NewInt(999))
		So(gtRes, ShouldBeTrue)

		sltRes, err := testOpCodeContract.Slt(callOpts, big.NewInt(-888), big.NewInt(-999))
		So(sltRes, ShouldBeFalse)
		sltRes, err = testOpCodeContract.Slt(callOpts, big.NewInt(-55959), big.NewInt(-999))
		So(sltRes, ShouldBeTrue)
		gltRes, err := testOpCodeContract.Sgt(callOpts, big.NewInt(-8889), big.NewInt(-999))
		So(gltRes, ShouldBeFalse)
		gltRes, err = testOpCodeContract.Sgt(callOpts, big.NewInt(-9), big.NewInt(-999))
		So(gltRes, ShouldBeTrue)

		eqRes, err := testOpCodeContract.Eq(callOpts, big.NewInt(0), big.NewInt(0))
		So(eqRes, ShouldBeTrue)
		eqRes, err = testOpCodeContract.Eq(callOpts, big.NewInt(4645), big.NewInt(9898515))
		So(eqRes, ShouldBeFalse)

		iszero, err := testOpCodeContract.Iszero(callOpts, big.NewInt(0))
		So(iszero, ShouldBeTrue)
		iszero, err = testOpCodeContract.Iszero(callOpts, big.NewInt(98989))
		So(iszero, ShouldBeFalse)

		andRes, err := testOpCodeContract.And(callOpts, big.NewInt(806), big.NewInt(656))
		checkErr(err)
		So(andRes.Cmp(big.NewInt(512)) == 0, ShouldBeTrue)

		orRes, err := testOpCodeContract.Or(callOpts, big.NewInt(23), big.NewInt(999))
		checkErr(err)
		So(orRes.Cmp(big.NewInt(1015)) == 0, ShouldBeTrue)
		xorRes, err := testOpCodeContract.Xor(callOpts, big.NewInt(55), big.NewInt(66))
		checkErr(err)
		So(xorRes.Cmp(big.NewInt(117)) == 0, ShouldBeTrue)

		notRes, err := testOpCodeContract.Not(callOpts, big.NewInt(55))
		checkErr(err)
		So(notRes.Cmp(big.NewInt(-56)) == 0, ShouldBeTrue)

		shlRes, err := testOpCodeContract.Shl(callOpts, big.NewInt(99))
		checkErr(err)
		So(shlRes.Cmp(big.NewInt(396)) == 0, ShouldBeTrue)

		shrRes, err := testOpCodeContract.Shr(callOpts, big.NewInt(98))
		checkErr(err)
		So(shrRes.Cmp(big.NewInt(24)) == 0, ShouldBeTrue)

		shaRes, err := testOpCodeContract.Sha(callOpts, "随便", "无所谓")
		So(shaRes, ShouldBeFalse)

		shaRes, err = testOpCodeContract.Sha(callOpts, "随便", "随便")
		So(shaRes, ShouldBeTrue)
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
		opts.GasPrice = big.NewInt(500000000000)
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
		log.Printf("testContract4Address: %s", testContract4Address)

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
		opts.GasPrice = big.NewInt(500000000000)
		opts.GasLimit = 3000000
		tx, err := testOpCodeContract.CallByFun(opts, testContract4Address)
		So(err, ShouldBeNil)
		res, err := utils.GetTransferInfoByHash(tx.Hash())
		checkErr(err)
		if res.Status != 1 {
			panic("tx fail!")
		}
	})
	//Convey("Test OpCode_addMod_mulMod", t, func() {
	//	testContractAbi, err := utils.ReadAll("deploy/TestOpCode/TestAddModAndMulMod_sol_TestAddModAndMulMod.abi")
	//	checkErr(err)
	//	testContractBin, err := utils.ReadAll("deploy/TestOpCode/TestAddModAndMulMod_sol_TestAddModAndMulMod.bin")
	//	checkErr(err)
	//	hash := deploy.ContractDeploy(string(testContractBin), string(testContractAbi))
	//	receipt, err := utils.GetTransferInfoByHash(hash)
	//	checkErr(err)
	//	if receipt.Status != 1 {
	//		panic("tx error!")
	//	}
	//	testContractAddress := receipt.ContractAddress
	//	log.Printf("testContractAddress: %s", testContractAddress)
	//	testOpCodeContract, err := TestOpCode.NewTestAddModAndMulMod(testContractAddress, con/
	//addModRes, err := testOpCodeContract.Addmod(callOpts, big.NewInt(500), big.NewInt(800),big.NewInt(800))
	//checkErr(err)
	//log.Printf("addModRes: %d", addModRes)
	//So(addModRes.Cmp(big.NewInt(80500)) == 0, ShouldBeTrue)
	//subRes, err := testOpCodeContract.Mulmod(callOpts, big.NewInt(500), big.NewInt(80000),big.NewInt(888))
	//So(err, ShouldNotBeNil)
	//log.Printf("subRes: %d", subRes)
	//So(subRes.Cmp(big.NewInt(80500)) == 0, ShouldBeTrue) })

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

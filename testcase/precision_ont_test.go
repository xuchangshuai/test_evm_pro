package testcase

import (
	"github.com/ontio/ontology/common"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	ontology_go_sdk "github.com/xuchangshuai/ontologyGoSdk"
	sdkcom "github.com/xuchangshuai/ontologyGoSdk/common"
	"log"
	"time"

	"math/big"
	"test_evm/config"
	"testing"
)


func Test_precision_ont(t *testing.T) {
	Convey("Test Ont BalanceV2&&transferV2(ont < 1)", t, func() {
		testOntSdk := config.GetSdk()
		testWallet, _ := testOntSdk.OpenWallet("./../wallet.dat")
		testDefAcc, err := testWallet.GetDefaultAccount([]byte("123456"))
		assert.Nil(t, err)

		toWallet, _ := testOntSdk.OpenWallet("./../wallet2.dat")
		toDefAcc, err := toWallet.GetDefaultAccount([]byte("123456"))
		assert.Nil(t, err)
		testGasLimit := uint64(20000)
		testGasPrice := uint64(2500)
		_, err = testOntSdk.Native.Ong.Transfer(testGasPrice, testGasLimit, nil, testDefAcc, toDefAcc.Address, 10000000000)
		assert.Nil(t, err)
		_, _ = testOntSdk.WaitForGenerateBlock(30 * time.Second)
		amount := big.NewInt(1)
		beforeBalance, err := testOntSdk.Native.Ont.BalanceOfV2(toDefAcc.Address)
		txHash, err := testOntSdk.Native.Ont.TransferV2(testGasPrice, testGasLimit, nil, testDefAcc, toDefAcc.Address, amount)
		assert.Nil(t, err)
		_, _ = testOntSdk.WaitForGenerateBlock(30 * time.Second)
		contractEvent, err := testOntSdk.GetSmartContractEvent(txHash.ToHexString())
		log.Println("contractEvent:", contractEvent)
		assert.Nil(t, err)
		for _, notify := range contractEvent.Notify {
			transfer, err := testOntSdk.ParseNativeTransferEventV2(notify)
			assert.Nil(t, err)
			t.Logf("transfer:%v", transfer)
		}
		balance, err := testOntSdk.Native.Ont.BalanceOfV2(toDefAcc.Address)
		assert.Nil(t, err)
		t.Logf("balance:%v", balance)
		So(balance.Cmp(big.NewInt(0).Add(beforeBalance,amount)) == 0, ShouldBeTrue)
	})
	Convey("Test Ont BalanceV2&&transferV2(ont > 1)", t, func() {
		testNetUrl := config.SDK_URL
		testOntSdk := ontology_go_sdk.NewOntologySdk()
		testOntSdk.NewRpcClient().SetAddress(testNetUrl)
		testWallet, _ := testOntSdk.OpenWallet("./../wallet.dat")
		testDefAcc, err := testWallet.GetDefaultAccount([]byte("123456"))
		assert.Nil(t, err)

		toWallet, _ := testOntSdk.OpenWallet("./../wallet2.dat")
		toDefAcc, err := toWallet.GetDefaultAccount([]byte("123456"))
		assert.Nil(t, err)
		testGasLimit := uint64(20000)
		testGasPrice := uint64(2500)
		amount := big.NewInt(1000000000)
		beforeBalance, err := testOntSdk.Native.Ont.BalanceOfV2(toDefAcc.Address)
		txHash, err := testOntSdk.Native.Ont.TransferV2(testGasPrice, testGasLimit, nil, testDefAcc, toDefAcc.Address, amount)
		assert.Nil(t, err)
		_, _ = testOntSdk.WaitForGenerateBlock(30 * time.Second)
		contractEvent, err := testOntSdk.GetSmartContractEvent(txHash.ToHexString())
		log.Println("contractEvent:", contractEvent)
		assert.Nil(t, err)
		for _, notify := range contractEvent.Notify {
			transfer, err := testOntSdk.ParseNativeTransferEventV2(notify)
			assert.Nil(t, err)
			t.Logf("transfer:%v", transfer)
		}
		balance, err := testOntSdk.Native.Ont.BalanceOfV2(toDefAcc.Address)
		assert.Nil(t, err)
		t.Logf("balance:%v", balance)
		So(balance.Cmp(big.NewInt(0).Add(beforeBalance,amount)) == 0, ShouldBeTrue)
	})
	Convey("Test Ont BalanceV2&&transferV2(ont > 1000w)", t, func() {
		testNetUrl := config.SDK_URL
		testOntSdk := ontology_go_sdk.NewOntologySdk()
		testOntSdk.NewRpcClient().SetAddress(testNetUrl)
		testWallet, _ := testOntSdk.OpenWallet("./../wallet.dat")
		testDefAcc, err := testWallet.GetDefaultAccount([]byte("123456"))
		assert.Nil(t, err)

		toWallet, _ := testOntSdk.OpenWallet("./../wallet2.dat")
		toDefAcc, err := toWallet.GetDefaultAccount([]byte("123456"))
		assert.Nil(t, err)
		testGasLimit := uint64(20000)
		testGasPrice := uint64(2500)
		amount := big.NewInt(100000000000000000)
		beforeBalance, err := testOntSdk.Native.Ont.BalanceOfV2(toDefAcc.Address)
		txHash, err := testOntSdk.Native.Ont.TransferV2(testGasPrice, testGasLimit, nil, testDefAcc, toDefAcc.Address, amount)
		assert.Nil(t, err)
		_, _ = testOntSdk.WaitForGenerateBlock(30 * time.Second)
		contractEvent, err := testOntSdk.GetSmartContractEvent(txHash.ToHexString())
		log.Println("contractEvent:", contractEvent)
		assert.Nil(t, err)
		for _, notify := range contractEvent.Notify {
			transfer, err := testOntSdk.ParseNativeTransferEventV2(notify)
			assert.Nil(t, err)
			t.Logf("transfer:%v", transfer)
		}
		balance, err := testOntSdk.Native.Ont.BalanceOfV2(toDefAcc.Address)
		assert.Nil(t, err)
		t.Logf("balance:%v", balance)
		So(balance.Cmp(big.NewInt(0).Add(beforeBalance,amount)) == 0, ShouldBeTrue)
	})
	Convey("Test Ont transferV2(amount = 1000000001)", t, func() {
		testNetUrl := config.SDK_URL
		testOntSdk := ontology_go_sdk.NewOntologySdk()
		testOntSdk.NewRpcClient().SetAddress(testNetUrl)
		testWallet, _ := testOntSdk.OpenWallet("./../wallet.dat")
		testDefAcc, err := testWallet.GetDefaultAccount([]byte("123456"))
		assert.Nil(t, err)

		toWallet, _ := testOntSdk.OpenWallet("./../wallet2.dat")
		toDefAcc, err := toWallet.GetDefaultAccount([]byte("123456"))
		assert.Nil(t, err)
		testGasLimit := uint64(20000)
		testGasPrice := uint64(2500)
		amount := big.NewInt(1000000001)
		beforeBalance, err := testOntSdk.Native.Ont.BalanceOfV2(toDefAcc.Address)
		txHash, err := testOntSdk.Native.Ont.TransferV2(testGasPrice, testGasLimit, nil, testDefAcc, toDefAcc.Address, amount)
		assert.Nil(t, err)
		_, _ = testOntSdk.WaitForGenerateBlock(30 * time.Second)
		contractEvent, err := testOntSdk.GetSmartContractEvent(txHash.ToHexString())
		log.Println("contractEvent:", contractEvent)
		assert.Nil(t, err)
		for _, notify := range contractEvent.Notify {
			transfer, err := testOntSdk.ParseNativeTransferEventV2(notify)
			assert.Nil(t, err)
			t.Logf("transfer:%v", transfer)
		}
		balance, err := testOntSdk.Native.Ont.BalanceOfV2(toDefAcc.Address)
		assert.Nil(t, err)
		t.Logf("balance:%v", balance)
		So(balance.Cmp(big.NewInt(0).Add(beforeBalance,amount)) == 0, ShouldBeTrue)
	})
	Convey("Test Ont transferV2(amount = 99999999)", t, func() {
		testNetUrl := config.SDK_URL
		testOntSdk := ontology_go_sdk.NewOntologySdk()
		testOntSdk.NewRpcClient().SetAddress(testNetUrl)
		testWallet, _ := testOntSdk.OpenWallet("./../wallet.dat")
		testDefAcc, err := testWallet.GetDefaultAccount([]byte("123456"))
		assert.Nil(t, err)

		toWallet, _ := testOntSdk.OpenWallet("./../wallet2.dat")
		toDefAcc, err := toWallet.GetDefaultAccount([]byte("123456"))
		assert.Nil(t, err)
		testGasLimit := uint64(20000)
		testGasPrice := uint64(2500)
		amount := big.NewInt(99999999)
		beforeBalance, err := testOntSdk.Native.Ont.BalanceOfV2(toDefAcc.Address)
		txHash, err := testOntSdk.Native.Ont.TransferV2(testGasPrice, testGasLimit, nil, testDefAcc, toDefAcc.Address, amount)
		assert.Nil(t, err)
		_, _ = testOntSdk.WaitForGenerateBlock(30 * time.Second)
		contractEvent, err := testOntSdk.GetSmartContractEvent(txHash.ToHexString())
		log.Println("contractEvent:", contractEvent)
		assert.Nil(t, err)
		for _, notify := range contractEvent.Notify {
			transfer, err := testOntSdk.ParseNativeTransferEventV2(notify)
			assert.Nil(t, err)
			t.Logf("transfer:%v", transfer)
		}
		balance, err := testOntSdk.Native.Ont.BalanceOfV2(toDefAcc.Address)
		assert.Nil(t, err)
		t.Logf("balance:%v", balance)
		So(balance.Cmp(big.NewInt(0).Add(beforeBalance,amount)) == 0, ShouldBeTrue)
	})
	Convey("Test Ont transferV2(amount = 1000000000)", t, func() {
		testNetUrl := config.SDK_URL
		testOntSdk := ontology_go_sdk.NewOntologySdk()
		testOntSdk.NewRpcClient().SetAddress(testNetUrl)
		testWallet, _ := testOntSdk.OpenWallet("./../wallet.dat")
		testDefAcc, err := testWallet.GetDefaultAccount([]byte("123456"))
		assert.Nil(t, err)

		toWallet, _ := testOntSdk.OpenWallet("./../wallet2.dat")
		toDefAcc, err := toWallet.GetDefaultAccount([]byte("123456"))
		assert.Nil(t, err)
		testGasLimit := uint64(20000)
		testGasPrice := uint64(2500)
		amount := big.NewInt(1000000000)
		beforeBalance, err := testOntSdk.Native.Ont.BalanceOfV2(toDefAcc.Address)
		txHash, err := testOntSdk.Native.Ont.TransferV2(testGasPrice, testGasLimit, nil, testDefAcc, toDefAcc.Address, amount)
		assert.Nil(t, err)
		_, _ = testOntSdk.WaitForGenerateBlock(30 * time.Second)
		contractEvent, err := testOntSdk.GetSmartContractEvent(txHash.ToHexString())
		log.Println("contractEvent:", contractEvent)
		assert.Nil(t, err)
		for _, notify := range contractEvent.Notify {
			transfer, err := testOntSdk.ParseNativeTransferEventV2(notify)
			assert.Nil(t, err)
			t.Logf("transfer:%v", transfer)
		}
		balance, err := testOntSdk.Native.Ont.BalanceOfV2(toDefAcc.Address)
		assert.Nil(t, err)
		t.Logf("balance:%v", balance)
		So(balance.Cmp(big.NewInt(0).Add(beforeBalance,amount)) == 0, ShouldBeTrue)
	})
	Convey("Test Ont transferV2(amount = 12345678123456789)", t, func() {
		testNetUrl := config.SDK_URL
		testOntSdk := ontology_go_sdk.NewOntologySdk()
		testOntSdk.NewRpcClient().SetAddress(testNetUrl)
		testWallet, _ := testOntSdk.OpenWallet("./../wallet.dat")
		testDefAcc, err := testWallet.GetDefaultAccount([]byte("123456"))
		assert.Nil(t, err)

		toWallet, _ := testOntSdk.OpenWallet("./../wallet2.dat")
		toDefAcc, err := toWallet.GetDefaultAccount([]byte("123456"))
		assert.Nil(t, err)
		testGasLimit := uint64(20000)
		testGasPrice := uint64(2500)
		amount := big.NewInt(12345678123456789)
		beforeBalance, err := testOntSdk.Native.Ont.BalanceOfV2(toDefAcc.Address)
		txHash, err := testOntSdk.Native.Ont.TransferV2(testGasPrice, testGasLimit, nil, testDefAcc, toDefAcc.Address, amount)
		assert.Nil(t, err)
		_, _ = testOntSdk.WaitForGenerateBlock(30 * time.Second)
		contractEvent, err := testOntSdk.GetSmartContractEvent(txHash.ToHexString())
		log.Println("contractEvent:", contractEvent)
		assert.Nil(t, err)
		for _, notify := range contractEvent.Notify {
			transfer, err := testOntSdk.ParseNativeTransferEventV2(notify)
			assert.Nil(t, err)
			t.Logf("transfer:%v", transfer)
		}
		balance, err := testOntSdk.Native.Ont.BalanceOfV2(toDefAcc.Address)
		assert.Nil(t, err)
		t.Logf("balance:%v", balance)
		So(balance.Cmp(big.NewInt(0).Add(beforeBalance,amount)) == 0, ShouldBeTrue)
	})
	Convey("Test Ont transferV2&&transfer", t, func() {
		testNetUrl := config.SDK_URL
		testOntSdk := ontology_go_sdk.NewOntologySdk()
		testOntSdk.NewRpcClient().SetAddress(testNetUrl)
		testWallet, _ := testOntSdk.OpenWallet("./../wallet.dat")
		testDefAcc, err := testWallet.GetDefaultAccount([]byte("123456"))
		assert.Nil(t, err)

		toWallet, _ := testOntSdk.OpenWallet("./../wallet2.dat")
		toDefAcc, err := toWallet.GetDefaultAccount([]byte("123456"))
		assert.Nil(t, err)
		testGasLimit := uint64(20000)
		testGasPrice := uint64(2500)
		amount := big.NewInt(1000)
		beforeBalance, err := testOntSdk.Native.Ont.BalanceOfV2(toDefAcc.Address)
		txHash, err := testOntSdk.Native.Ont.TransferV2(testGasPrice, testGasLimit, nil, testDefAcc, toDefAcc.Address, amount)
		assert.Nil(t, err)
		_, _ = testOntSdk.WaitForGenerateBlock(30 * time.Second)
		contractEvent, err := testOntSdk.GetSmartContractEvent(txHash.ToHexString())
		log.Println("contractEvent:", contractEvent)
		assert.Nil(t, err)
		for _, notify := range contractEvent.Notify {
			transfer, err := testOntSdk.ParseNativeTransferEventV2(notify)
			assert.Nil(t, err)
			t.Logf("transfer:%v", transfer)
		}
		balance, err := testOntSdk.Native.Ont.BalanceOfV2(toDefAcc.Address)
		assert.Nil(t, err)
		t.Logf("balance:%v", balance)
		So(balance.Cmp(big.NewInt(0).Add(beforeBalance,amount)) == 0, ShouldBeTrue)
		// transferV1
		amount2 := uint64(12)
		txHash, err = testOntSdk.Native.Ont.Transfer(testGasPrice, testGasLimit, nil, testDefAcc, toDefAcc.Address, amount2)
		assert.Nil(t, err)
		_, _ = testOntSdk.WaitForGenerateBlock(30 * time.Second)
		balance2, err := testOntSdk.Native.Ont.BalanceOfV2(toDefAcc.Address)
		assert.Nil(t, err)
		So(balance2.Cmp(big.NewInt(0).Add(balance,big.NewInt(0).Mul(big.NewInt(int64(amount2)),big.NewInt(1000000000)))) == 0, ShouldBeTrue)
	})
	Convey("Test Ont transferV2(amount = 0)", t, func() {
		testNetUrl := config.SDK_URL
		testOntSdk := ontology_go_sdk.NewOntologySdk()
		testOntSdk.NewRpcClient().SetAddress(testNetUrl)
		testWallet, _ := testOntSdk.OpenWallet("./../wallet.dat")
		testDefAcc, err := testWallet.GetDefaultAccount([]byte("123456"))
		assert.Nil(t, err)

		toWallet, _ := testOntSdk.OpenWallet("./../wallet2.dat")
		toDefAcc, err := toWallet.GetDefaultAccount([]byte("123456"))
		assert.Nil(t, err)
		testGasLimit := uint64(20000)
		testGasPrice := uint64(2500)
		amount := big.NewInt(0)
		beforeBalance, err := testOntSdk.Native.Ont.BalanceOfV2(toDefAcc.Address)
		txHash, err := testOntSdk.Native.Ont.TransferV2(testGasPrice, testGasLimit, nil, testDefAcc, toDefAcc.Address, amount)
		assert.Nil(t, err)
		_, _ = testOntSdk.WaitForGenerateBlock(30 * time.Second)
		contractEvent, err := testOntSdk.GetSmartContractEvent(txHash.ToHexString())
		log.Println("contractEvent:", contractEvent)
		assert.Nil(t, err)
		for _, notify := range contractEvent.Notify {
			transfer, err := testOntSdk.ParseNativeTransferEventV2(notify)
			assert.Nil(t, err)
			t.Logf("transfer:%v", transfer)
		}
		balance, err := testOntSdk.Native.Ont.BalanceOfV2(toDefAcc.Address)
		assert.Nil(t, err)
		t.Logf("balance:%v", balance)
		So(balance.Cmp(big.NewInt(0).Add(beforeBalance,amount)) == 0, ShouldBeTrue)
	})
	Convey("Test Ont transferV2(amount = -1)", t, func() {
		testNetUrl := config.SDK_URL
		testOntSdk := ontology_go_sdk.NewOntologySdk()
		testOntSdk.NewRpcClient().SetAddress(testNetUrl)
		testWallet, _ := testOntSdk.OpenWallet("./../wallet.dat")
		testDefAcc, err := testWallet.GetDefaultAccount([]byte("123456"))
		assert.Nil(t, err)

		toWallet, _ := testOntSdk.OpenWallet("./../wallet2.dat")
		toDefAcc, err := toWallet.GetDefaultAccount([]byte("123456"))
		assert.Nil(t, err)
		testGasLimit := uint64(20000)
		testGasPrice := uint64(2500)
		amount := big.NewInt(-1)
		_, err = testOntSdk.Native.Ont.TransferV2(testGasPrice, testGasLimit, nil, testDefAcc, toDefAcc.Address, amount)
		So(err,ShouldNotBeNil)
	})
	Convey("Test Ont transferV2(可用ont小于转账ont)", t, func() {
		testNetUrl := config.SDK_URL
		testOntSdk := ontology_go_sdk.NewOntologySdk()
		testOntSdk.NewRpcClient().SetAddress(testNetUrl)
		testWallet, _ := testOntSdk.OpenWallet("./../wallet.dat")
		testDefAcc, err := testWallet.GetDefaultAccount([]byte("123456"))
		assert.Nil(t, err)

		toWallet, _ := testOntSdk.OpenWallet("./../wallet2.dat")
		toDefAcc, err := toWallet.GetDefaultAccount([]byte("123456"))
		assert.Nil(t, err)
		testGasLimit := uint64(20000)
		testGasPrice := uint64(2500)
		amount, err := testOntSdk.Native.Ont.BalanceOfV2(testDefAcc.Address)
		assert.Nil(t, err)
		amount.Add(amount, big.NewInt(1))
		_, err = testOntSdk.Native.Ont.TransferV2(testGasPrice, testGasLimit, nil, testDefAcc, toDefAcc.Address, amount)
		So(err, ShouldNotBeNil)
	})
	//ApproveV2&&AllowanceV2&&TransferFromV2
	Convey("Test Ont ApproveV2&&AllowanceV2", t, func() {
		testNetUrl := config.SDK_URL
		testOntSdk := ontology_go_sdk.NewOntologySdk()
		testOntSdk.NewRpcClient().SetAddress(testNetUrl)
		testWallet, _ := testOntSdk.OpenWallet("./../wallet.dat")
		testDefAcc, err := testWallet.GetDefaultAccount([]byte("123456"))
		assert.Nil(t, err)

		toWallet, _ := testOntSdk.OpenWallet("./../wallet2.dat")
		toDefAcc, err := toWallet.GetDefaultAccount([]byte("123456"))
		assert.Nil(t, err)
		testGasPrice := uint64(2500)
		testGasLimit := uint64(20000)
		// case1: amount = 500
		amount := big.NewInt(500)
		// testDefAcc调用ApproveV2 -> toDefAcc amount 500
		_, err = testOntSdk.Native.Ont.ApproveV2(testGasPrice, testGasLimit, nil, testDefAcc, toDefAcc.Address, amount)
		assert.Nil(t, err)
		_, _ = testOntSdk.WaitForGenerateBlock(30 * time.Second)
		// toDefAcc调用AllowanceV2-> testDefAcc amount 500
		allowanceAmount, err := testOntSdk.Native.Ont.AllowanceV2(testDefAcc.Address, toDefAcc.Address)
		assert.Nil(t, err)
		So(allowanceAmount.Cmp(amount) == 0, ShouldBeTrue)
		// toDefAcc调用AllowanceV2-> testDefAcc amount 0
		allowanceAmount2, err := testOntSdk.Native.Ont.Allowance(testDefAcc.Address, toDefAcc.Address)
		assert.Nil(t, err)
		So(allowanceAmount2 == 0, ShouldBeTrue)
		// case2: amount = 50 * 10 ** 9
		amount = big.NewInt(50000000000)
		_, err = testOntSdk.Native.Ont.ApproveV2(testGasPrice, testGasLimit, nil, testDefAcc, toDefAcc.Address, amount)
		assert.Nil(t, err)
		_, _ = testOntSdk.WaitForGenerateBlock(40 * time.Second)
		allowanceAmount, err = testOntSdk.Native.Ont.AllowanceV2(testDefAcc.Address, toDefAcc.Address)
		assert.Nil(t, err)
		So(allowanceAmount.Cmp(amount) == 0, ShouldBeTrue)
		allowanceAmount2, err = testOntSdk.Native.Ont.Allowance(testDefAcc.Address, toDefAcc.Address)
		assert.Nil(t, err)
		So(allowanceAmount2 == uint64(50), ShouldBeTrue)
		// case3: amount = 0
		amount = big.NewInt(0)
		_, err = testOntSdk.Native.Ont.ApproveV2(testGasPrice, testGasLimit, nil, testDefAcc, toDefAcc.Address, amount)
		assert.Nil(t, err)
		_, _ = testOntSdk.WaitForGenerateBlock(30 * time.Second)
		allowanceAmount, err = testOntSdk.Native.Ont.AllowanceV2(testDefAcc.Address, toDefAcc.Address)
		assert.Nil(t, err)
		So(allowanceAmount.Cmp(big.NewInt(0)) == 0, ShouldBeTrue)
		allowanceAmount2, err = testOntSdk.Native.Ont.Allowance(testDefAcc.Address, toDefAcc.Address)
		t.Logf("allowance Amount 2: %v", allowanceAmount2)
		assert.Nil(t, err)
		So(allowanceAmount2 == 0, ShouldBeTrue)
		// case4: amount = 12345678_987654321
		amount = big.NewInt(12345678987654321)
		amount.Add(amount, big.NewInt(1))
		_, err = testOntSdk.Native.Ont.ApproveV2(testGasPrice, testGasLimit, nil, testDefAcc, toDefAcc.Address, amount)
		assert.Nil(t, err)
		_, _ = testOntSdk.WaitForGenerateBlock(30 * time.Second)
		allowanceAmount, err = testOntSdk.Native.Ont.AllowanceV2(testDefAcc.Address, toDefAcc.Address)
		assert.Nil(t, err)
		So(allowanceAmount.Cmp(amount) == 0, ShouldBeTrue)
		allowanceAmount2, err = testOntSdk.Native.Ont.Allowance(testDefAcc.Address, toDefAcc.Address)
		assert.Nil(t, err)
		So(allowanceAmount2 == 12345678, ShouldBeTrue)
	})
	Convey("Test Ont TransferFromV2", t, func() {
		testNetUrl := config.SDK_URL
		testOntSdk := ontology_go_sdk.NewOntologySdk()
		testOntSdk.NewRpcClient().SetAddress(testNetUrl)
		testWallet, _ := testOntSdk.OpenWallet("./../wallet.dat")
		testDefAcc, err := testWallet.GetDefaultAccount([]byte("123456"))
		assert.Nil(t, err)

		toWallet, _ := testOntSdk.OpenWallet("./../wallet2.dat")
		toDefAcc, err := toWallet.GetDefaultAccount([]byte("123456"))
		assert.Nil(t, err)
		testGasPrice := uint64(2500)
		testGasLimit := uint64(20000)
		// case1: 用户a approve给用户b 50000000000 ont 用户b调用TransferFromV2转账给c >50000000000
		cAddr, err := common.AddressFromBase58("AJ3scNwZF9pUoeQYhrXJs7DTEBxmNJzSWa")
		assert.Nil(t, err)
		amount := big.NewInt(50000000000)
		// testDefAcc调用ApproveV2 -> toDefAcc amount 50000000000
		_, err = testOntSdk.Native.Ont.ApproveV2(testGasPrice, testGasLimit, nil, testDefAcc, toDefAcc.Address, amount)
		assert.Nil(t, err)
		_, _ = testOntSdk.WaitForGenerateBlock(30 * time.Second)
		// 用户b调用TransferFromV2转账给c >50000000000
		allowanceAmount, err := testOntSdk.Native.Ont.AllowanceV2(testDefAcc.Address, toDefAcc.Address)
		assert.Nil(t, err)
		t.Logf("allowanceAmount :%v", allowanceAmount)
		So(allowanceAmount.Cmp(amount) == 0,ShouldBeTrue)
		// 给b账户转一个ont  手续费
		res,err := testOntSdk.Native.Ong.Transfer(testGasPrice, testGasLimit, nil, testDefAcc,toDefAcc.Address,1000000000)
		assert.Nil(t, err)
		t.Logf("res :%s", res.ToHexString())
		_, _ = testOntSdk.WaitForGenerateBlock(30 * time.Second)
		toDefAccBalance,err :=testOntSdk.Native.Ong.BalanceOfV2(toDefAcc.Address)
		assert.Nil(t, err)
		t.Logf("toDefAccBalance :%v", toDefAccBalance)
		So(toDefAccBalance.Cmp(big.NewInt(1000000000000000000))>=0,ShouldBeTrue)
		// err
		res,err = testOntSdk.Native.Ont.TransferFromV2(testGasPrice, testGasLimit, nil, toDefAcc, testDefAcc.Address,cAddr,big.NewInt(0).Add(allowanceAmount, big.NewInt(1)))
		So(err, ShouldNotBeNil)
		// case2: 用户b调用TransferFromV2转账给c = 0
		testDefAccBalance,err := testOntSdk.Native.Ont.BalanceOfV2(testDefAcc.Address)
		assert.Nil(t, err)
		cAddrBalance,err := testOntSdk.Native.Ont.BalanceOfV2(cAddr)
		assert.Nil(t, err)
		res,err = testOntSdk.Native.Ont.TransferFromV2(testGasPrice, testGasLimit, nil, toDefAcc, testDefAcc.Address,cAddr,big.NewInt(0))
		assert.Nil(t, err)
		_, _ = testOntSdk.WaitForGenerateBlock(30 * time.Second)
		testDefAccBalanceAfter,err := testOntSdk.Native.Ont.BalanceOfV2(testDefAcc.Address)
		assert.Nil(t, err)
		cAddrBalanceAfter,err := testOntSdk.Native.Ont.BalanceOfV2(cAddr)
		assert.Nil(t, err)
		// 断言资金无变动
		So(testDefAccBalance.Cmp(testDefAccBalanceAfter)==0,ShouldBeTrue)
		So(cAddrBalance.Cmp(cAddrBalanceAfter)==0,ShouldBeTrue)
		// case3: 用户b调用TransferFromV2转账给c < allowance
		testDefAccBalance,err = testOntSdk.Native.Ont.BalanceOfV2(testDefAcc.Address)
		assert.Nil(t, err)
		cAddrBalance,err = testOntSdk.Native.Ont.BalanceOfV2(cAddr)
		assert.Nil(t, err)
		amount2 := big.NewInt(0).Sub(amount, big.NewInt(10000000000))
		res,err = testOntSdk.Native.Ont.TransferFromV2(testGasPrice, testGasLimit, nil, toDefAcc, testDefAcc.Address,cAddr,amount2)
		assert.Nil(t, err)
		_, _ = testOntSdk.WaitForGenerateBlock(30 * time.Second)
		testDefAccBalanceAfter,err = testOntSdk.Native.Ont.BalanceOfV2(testDefAcc.Address)
		assert.Nil(t, err)
		cAddrBalanceAfter,err = testOntSdk.Native.Ont.BalanceOfV2(cAddr)
		assert.Nil(t, err)
		// 断言资金变动
		So(testDefAccBalance.Cmp(big.NewInt(0).Add(testDefAccBalanceAfter,amount2))==0,ShouldBeTrue)
		So(cAddrBalance.Cmp(big.NewInt(0).Sub(cAddrBalanceAfter, amount2))==0,ShouldBeTrue)
		allowanceAmountAfter, err := testOntSdk.Native.Ont.AllowanceV2(testDefAcc.Address, toDefAcc.Address)
		assert.Nil(t, err)
		t.Logf("allowanceAmountAfter :%v", allowanceAmountAfter)
		So(allowanceAmountAfter.Cmp(big.NewInt(0).Sub(allowanceAmount, amount2))==0,ShouldBeTrue)
		// case4: 用户b调用TransferFromV2转账给c = allowance
		testDefAccBalance2,err := testOntSdk.Native.Ont.BalanceOfV2(testDefAcc.Address)
		assert.Nil(t, err)
		cAddrBalance2,err := testOntSdk.Native.Ont.BalanceOfV2(cAddr)
		assert.Nil(t, err)
		res,err = testOntSdk.Native.Ont.TransferFromV2(testGasPrice, testGasLimit, nil, toDefAcc, testDefAcc.Address,cAddr,allowanceAmountAfter)
		assert.Nil(t, err)
		_, _ = testOntSdk.WaitForGenerateBlock(30 * time.Second)
		testDefAccBalanceAfter,err = testOntSdk.Native.Ont.BalanceOfV2(testDefAcc.Address)
		assert.Nil(t, err)
		cAddrBalanceAfter,err = testOntSdk.Native.Ont.BalanceOfV2(cAddr)
		assert.Nil(t, err)
		// 断言资金变动
		So(testDefAccBalance2.Cmp(big.NewInt(0).Add(testDefAccBalanceAfter,allowanceAmountAfter))==0,ShouldBeTrue)
		So(cAddrBalanceAfter.Cmp(big.NewInt(0).Add(cAddrBalance2, allowanceAmountAfter))==0,ShouldBeTrue)
		allowanceAmountAfter, err = testOntSdk.Native.Ont.AllowanceV2(testDefAcc.Address, toDefAcc.Address)
		assert.Nil(t, err)
		t.Logf("allowanceAmountAfter :%v", allowanceAmountAfter)
		So(allowanceAmountAfter.Cmp(big.NewInt(0))==0,ShouldBeTrue)
		// case5: 用户b调用TransferFromV2转账给c =A -> B APPROVE = 0
		res,err = testOntSdk.Native.Ont.TransferFromV2(testGasPrice, testGasLimit, nil, toDefAcc, testDefAcc.Address,cAddr,big.NewInt(1))
		So(err, ShouldNotBeNil)
		// case6:   balance = testOntSdk.Native.Ont.BalanceOfV2( a )；
		//	    	ApproveV2（balance+1）；
		//	    	用户b TransferFromV2() balance； 抛异常
		testDefAccBalance,err = testOntSdk.Native.Ont.BalanceOfV2(testDefAcc.Address)
		assert.Nil(t, err)
		amount = big.NewInt(0).Add(testDefAccBalance, big.NewInt(1))
		_, err = testOntSdk.Native.Ont.ApproveV2(testGasPrice, testGasLimit, nil, testDefAcc, toDefAcc.Address, amount)
		assert.Nil(t, err)
		_, _ = testOntSdk.WaitForGenerateBlock(30 * time.Second)
		allowanceAmount, err = testOntSdk.Native.Ont.AllowanceV2(testDefAcc.Address, toDefAcc.Address)
		assert.Nil(t, err)
		So(allowanceAmount.Cmp(amount) == 0, ShouldBeTrue)
		res,err = testOntSdk.Native.Ont.TransferFromV2(testGasPrice, testGasLimit, nil, toDefAcc, testDefAcc.Address,cAddr,amount)
		So(err, ShouldNotBeNil)
	})
	Convey("Test Ont MultipleTransferV2", t, func() {
		testNetUrl := config.SDK_URL
		testOntSdk := ontology_go_sdk.NewOntologySdk()
		testOntSdk.NewRpcClient().SetAddress(testNetUrl)
		testWallet, _ := testOntSdk.OpenWallet("./../wallet.dat")
		testDefAcc, err := testWallet.GetDefaultAccount([]byte("123456"))
		assert.Nil(t, err)

		toWallet, _ := testOntSdk.OpenWallet("./../wallet2.dat")
		toDefAcc, err := toWallet.GetDefaultAccount([]byte("123456"))
		assert.Nil(t, err)
		testGasPrice := uint64(2500)
		testGasLimit := uint64(20000)
		// case1. {from:a,to:b,amount:1}{from:a,to:c,amount:1}成功
		cAddr, err := common.AddressFromBase58("AJ3scNwZF9pUoeQYhrXJs7DTEBxmNJzSWa")
		assert.Nil(t, err)
		amount:=big.NewInt(1)
		tx1:= &sdkcom.TransferStateV2 {
			From: testDefAcc.Address,
			To: toDefAcc.Address,
			Value: amount,
		}
		tx2:= &sdkcom.TransferStateV2 {
			From: testDefAcc.Address,
			To: cAddr,
			Value: amount,
		}
		txs := []*sdkcom.TransferStateV2{tx1,tx2}

		testDefAccBalance,_ := testOntSdk.Native.Ont.BalanceOfV2(testDefAcc.Address)
		toAccountBalance,_ := testOntSdk.Native.Ont.BalanceOfV2(toDefAcc.Address)
		cAccountBalance,_ := testOntSdk.Native.Ont.BalanceOfV2(cAddr)
		_,err = testOntSdk.Native.Ont.MultiTransferV2(testGasPrice, testGasLimit,nil ,txs,testDefAcc)
		assert.Nil(t, err)
		_, _ = testOntSdk.WaitForGenerateBlock(30 * time.Second)
		testDefAccAfter,_ := testOntSdk.Native.Ont.BalanceOfV2(testDefAcc.Address)
		toAccountBalanceAfter,_ := testOntSdk.Native.Ont.BalanceOfV2(toDefAcc.Address)
		cAccountBalanceTsAfter,_ := testOntSdk.Native.Ont.BalanceOfV2(cAddr)
		So(testDefAccBalance.Cmp(big.NewInt(0).Add(testDefAccAfter,big.NewInt(0).Add(amount, amount))) == 0,ShouldBeTrue)
		So(toAccountBalanceAfter.Cmp(big.NewInt(0).Add(toAccountBalance,amount)) == 0, ShouldBeTrue)
		So(cAccountBalanceTsAfter.Cmp(big.NewInt(0).Add(cAccountBalance,amount)) == 0, ShouldBeTrue)
		// case2 . {from:a,to:b,amount:0}{from:a,to:c,amount:0}成功
		amount = big.NewInt(0)
		tx1 = &sdkcom.TransferStateV2 {
			From: testDefAcc.Address,
			To: toDefAcc.Address,
			Value: amount,
		}
		tx2 = &sdkcom.TransferStateV2 {
			From: testDefAcc.Address,
			To: cAddr,
			Value: amount,
		}
		txs = []*sdkcom.TransferStateV2{tx1,tx2}
		testDefAccBalance,_ = testOntSdk.Native.Ont.BalanceOfV2(testDefAcc.Address)
		toAccountBalance,_ = testOntSdk.Native.Ont.BalanceOfV2(toDefAcc.Address)
		cAccountBalance,_ = testOntSdk.Native.Ont.BalanceOfV2(cAddr)
		_,err = testOntSdk.Native.Ont.MultiTransferV2(testGasPrice, testGasLimit,nil,txs,testDefAcc)
		assert.Nil(t, err)
		_, _ = testOntSdk.WaitForGenerateBlock(30 * time.Second)
		testDefAccAfter,_ = testOntSdk.Native.Ont.BalanceOfV2(testDefAcc.Address)
		toAccountBalanceAfter,_ = testOntSdk.Native.Ont.BalanceOfV2(toDefAcc.Address)
		cAccountBalanceTsAfter,_ = testOntSdk.Native.Ont.BalanceOfV2(cAddr)
		So(testDefAccBalance.Cmp(big.NewInt(0).Add(testDefAccAfter,big.NewInt(0).Add(amount, amount))) == 0,ShouldBeTrue)
		So(toAccountBalanceAfter.Cmp(big.NewInt(0).Add(toAccountBalance,amount)) == 0, ShouldBeTrue)
		So(cAccountBalanceTsAfter.Cmp(big.NewInt(0).Add(cAccountBalance,amount)) == 0, ShouldBeTrue)
		// case 3.{from:a,to:b,amount:-1}{from:a,to:c,amount:-1}
		amount = big.NewInt(-1)
		tx1 = &sdkcom.TransferStateV2 {
			From: testDefAcc.Address,
			To: toDefAcc.Address,
			Value: amount,
		}
		tx2 = &sdkcom.TransferStateV2 {
			From: testDefAcc.Address,
			To: cAddr,
			Value: amount,
		}
		txs = []*sdkcom.TransferStateV2{tx1,tx2}
		_,err = testOntSdk.Native.Ont.MultiTransferV2(testGasPrice, testGasLimit,nil,txs,testDefAcc)
		So(err, ShouldNotBeNil)
		// case4 . {from:a,to:b,amount:1000000001}{from:a,to:c,amount:1000000001}成功
		amount = big.NewInt(1000000001)
		tx1 = &sdkcom.TransferStateV2 {
			From: testDefAcc.Address,
			To: toDefAcc.Address,
			Value: amount,
		}
		tx2 = &sdkcom.TransferStateV2 {
			From: testDefAcc.Address,
			To: cAddr,
			Value: amount,
		}
		txs = []*sdkcom.TransferStateV2{tx1,tx2}
		testDefAccBalance,_ = testOntSdk.Native.Ont.BalanceOfV2(testDefAcc.Address)
		toAccountBalance,_ = testOntSdk.Native.Ont.BalanceOfV2(toDefAcc.Address)
		cAccountBalance,_ = testOntSdk.Native.Ont.BalanceOfV2(cAddr)
		_,err = testOntSdk.Native.Ont.MultiTransferV2(testGasPrice, testGasLimit,nil,txs,testDefAcc)
		assert.Nil(t, err)
		_, _ = testOntSdk.WaitForGenerateBlock(30 * time.Second)
		testDefAccAfter,_ = testOntSdk.Native.Ont.BalanceOfV2(testDefAcc.Address)
		toAccountBalanceAfter,_ = testOntSdk.Native.Ont.BalanceOfV2(toDefAcc.Address)
		cAccountBalanceTsAfter,_ = testOntSdk.Native.Ont.BalanceOfV2(cAddr)
		So(testDefAccBalance.Cmp(big.NewInt(0).Add(testDefAccAfter,big.NewInt(0).Add(amount, amount))) == 0,ShouldBeTrue)
		So(toAccountBalanceAfter.Cmp(big.NewInt(0).Add(toAccountBalance,amount)) == 0, ShouldBeTrue)
		So(cAccountBalanceTsAfter.Cmp(big.NewInt(0).Add(cAccountBalance,amount)) == 0, ShouldBeTrue)
		// case5 . {from:a,to:b,amount:99999999}{from:a,to:c,amount:99999999}成功
		amount = big.NewInt(99999999)
		tx1 = &sdkcom.TransferStateV2 {
			From: testDefAcc.Address,
			To: toDefAcc.Address,
			Value: amount,
		}
		tx2 = &sdkcom.TransferStateV2 {
			From: testDefAcc.Address,
			To: cAddr,
			Value: amount,
		}
		txs = []*sdkcom.TransferStateV2{tx1,tx2}
		testDefAccBalance,_ = testOntSdk.Native.Ont.BalanceOfV2(testDefAcc.Address)
		toAccountBalance,_ = testOntSdk.Native.Ont.BalanceOfV2(toDefAcc.Address)
		cAccountBalance,_ = testOntSdk.Native.Ont.BalanceOfV2(cAddr)
		_,err = testOntSdk.Native.Ont.MultiTransferV2(testGasPrice, testGasLimit,nil,txs,testDefAcc)
		assert.Nil(t, err)
		_, _ = testOntSdk.WaitForGenerateBlock(30 * time.Second)
		testDefAccAfter,_ = testOntSdk.Native.Ont.BalanceOfV2(testDefAcc.Address)
		toAccountBalanceAfter,_ = testOntSdk.Native.Ont.BalanceOfV2(toDefAcc.Address)
		cAccountBalanceTsAfter,_ = testOntSdk.Native.Ont.BalanceOfV2(cAddr)
		So(testDefAccBalance.Cmp(big.NewInt(0).Add(testDefAccAfter,big.NewInt(0).Add(amount, amount))) == 0,ShouldBeTrue)
		So(toAccountBalanceAfter.Cmp(big.NewInt(0).Add(toAccountBalance,amount)) == 0, ShouldBeTrue)
		So(cAccountBalanceTsAfter.Cmp(big.NewInt(0).Add(cAccountBalance,amount)) == 0, ShouldBeTrue)
		// case6 . {from:a,to:b,amount:12345678123456789000000000}{from:a,to:c,amount:12345678123456789000000000}
		amount = big.NewInt(12345678123456789)
		tx1 = &sdkcom.TransferStateV2 {
			From: testDefAcc.Address,
			To: toDefAcc.Address,
			Value: amount,
		}
		tx2 = &sdkcom.TransferStateV2 {
			From: testDefAcc.Address,
			To: cAddr,
			Value: amount,
		}
		txs = []*sdkcom.TransferStateV2{tx1,tx2}
		testDefAccBalance,_ = testOntSdk.Native.Ont.BalanceOfV2(testDefAcc.Address)
		toAccountBalance,_ = testOntSdk.Native.Ont.BalanceOfV2(toDefAcc.Address)
		cAccountBalance,_ = testOntSdk.Native.Ont.BalanceOfV2(cAddr)
		_,err = testOntSdk.Native.Ont.MultiTransferV2(testGasPrice, testGasLimit,nil,txs,testDefAcc)
		assert.Nil(t, err)
		_, _ = testOntSdk.WaitForGenerateBlock(30 * time.Second)
		testDefAccAfter,_ = testOntSdk.Native.Ont.BalanceOfV2(testDefAcc.Address)
		toAccountBalanceAfter,_ = testOntSdk.Native.Ont.BalanceOfV2(toDefAcc.Address)
		cAccountBalanceTsAfter,_ = testOntSdk.Native.Ont.BalanceOfV2(cAddr)
		So(testDefAccBalance.Cmp(big.NewInt(0).Add(testDefAccAfter,big.NewInt(0).Add(amount, amount))) == 0,ShouldBeTrue)
		So(toAccountBalanceAfter.Cmp(big.NewInt(0).Add(toAccountBalance,amount)) == 0, ShouldBeTrue)
		So(cAccountBalanceTsAfter.Cmp(big.NewInt(0).Add(cAccountBalance,amount)) == 0, ShouldBeTrue)
		// case6 7 .amount = balance(a)/2 {from:a,to:b,amount:amount}{from:a,to:c,amount:amount}  交易失败
		testDefAccBalance,_ = testOntSdk.Native.Ont.BalanceOfV2(testDefAcc.Address)
		t.Logf("testDefAccBalance: %v", testDefAccBalance)
		toAccountBalance,_ = testOntSdk.Native.Ont.BalanceOfV2(toDefAcc.Address)
		t.Logf("toAccountBalance: %v", toAccountBalance)
		cAccountBalance,_ = testOntSdk.Native.Ont.BalanceOfV2(cAddr)
		t.Logf("cAccountBalance: %v", cAccountBalance)
		amount = big.NewInt(0).Div(testDefAccBalance, big.NewInt(2))
		t.Logf("amount: %v", amount)
		tx1 = &sdkcom.TransferStateV2 {
			From: testDefAcc.Address,
			To: toDefAcc.Address,
			Value: amount,
		}
		tx2 = &sdkcom.TransferStateV2 {
			From: testDefAcc.Address,
			To: cAddr,
			Value: big.NewInt(0).Add(amount, big.NewInt(1)),
		}
		txs = []*sdkcom.TransferStateV2{tx1,tx2}

		_,err = testOntSdk.Native.Ont.MultiTransferV2(testGasPrice, testGasLimit,nil,txs,testDefAcc)
		So(err, ShouldNotBeNil)
	})
}


package testcase

import (
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/smartystreets/goconvey/convey"
	"math/big"
	"test_evm/config"
	"test_evm/testcase/deploy"
	"test_evm/utils"
	"testing"
)

func TestEVMTransfer001(t *testing.T) {
	Convey("Test001 测试EIP155交易普通转账ong ", t, func() {
		toAddress := "0xf45505D1F482EBc8881dacA97B122B62771B9e1d"
		amount := big.NewInt(1000000000)
		//  备注： 对应的MetaMask钱包第一个账户私钥
		fromPrivateKey := config.FromPrivate
		fromPrivate, _ := crypto.HexToECDSA(fromPrivateKey)
		fromAddress := crypto.PubkeyToAddress(fromPrivate.PublicKey)
		fmt.Println("fromAddress: ", fromAddress)

		//交易前账户余额
		fromBalance := utils.GetBalance( fromAddress.String())
		toBalance := utils.GetBalance( toAddress)
		managerBalance := utils.GetBalance( "0x0000000000000000000000000000000000000007")

		hash, err := deploy.SendTransfer(fromPrivateKey, toAddress, amount, uint64(200000))
		if err != nil {
			panic(err)
		}
		//fmt.Println(hash)
		receipt, err := utils.GetTransferInfoByHash(hash)
		if err != nil {
			panic("交易未查到！")
		}
		//fmt.Println(receipt)
		if receipt.Status != 1 {
			fmt.Println("status : ", receipt.Status)
			panic("TxTransaction fail !!!")
		}
		useGas := receipt.CumulativeGasUsed
		fmt.Println("useGas", useGas)
		fromBalanceAfter := utils.GetBalance( fromAddress.String())
		toBalanceAfter := utils.GetBalance( toAddress)
		managerBalanceAfter := utils.GetBalance( "0x0000000000000000000000000000000000000007")
		fmt.Println("fromBalance", fromBalance)
		fmt.Println("toBalance", toBalance)
		fmt.Println("managerBalance", managerBalance)
		fmt.Println("fromBalanceAfter", fromBalanceAfter)
		fmt.Println("toBalanceAfter", toBalanceAfter)
		fmt.Println("managerBalanceAfter", managerBalanceAfter)
		//1 断言from账户资金是否 = fromBalance1 = fromBalance - gasUsed - amount
		fromBalance1 := fromBalanceAfter.Add(fromBalanceAfter, amount).Add(fromBalanceAfter, big.NewInt(int64(useGas)))
		So(fromBalance.Cmp(fromBalance1) == 0, ShouldBeTrue)
		//2 断言to账户资金是否 = toBalance + amount
		So(toBalanceAfter.Cmp(toBalance.Add(toBalance, amount)) == 0, ShouldBeTrue)
		//3 断言治理合约地址是否正常获得消耗gas
		So(managerBalanceAfter.Cmp(managerBalance.Add(managerBalance, big.NewInt(int64(useGas)))) == 0, ShouldBeTrue)
	})
}

func TestEVMTransfer002(t *testing.T) {
	Convey("Test002 测试EIP155交易普通转账ong(amount:255) ", t, func() {
		toAddress := "0xf45505D1F482EBc8881dacA97B122B62771B9e1d"
		amount := big.NewInt(255000000000)
		//  备注： 对应的MetaMask钱包第一个账户私钥
		fromPrivateKey := config.FromPrivate
		fromPrivate, _ := crypto.HexToECDSA(fromPrivateKey)
		fromAddress := crypto.PubkeyToAddress(fromPrivate.PublicKey)

		//交易前账户余额
		fromBalance := utils.GetBalance(  fromAddress.String())
		toBalance := utils.GetBalance(  toAddress)
		managerBalance := utils.GetBalance(  "0x0000000000000000000000000000000000000007")

		hash, err := deploy.SendTransfer(  fromPrivateKey, toAddress, amount, uint64(200000))
		if err != nil {
			panic(err)
		}
		//fmt.Println(hash)
		receipt, err := utils.GetTransferInfoByHash(hash)
		if err != nil {
			fmt.Println("交易未落账！")
			panic(err)
		}
		if receipt.Status != 1 {
			fmt.Println("status : ", receipt.Status)
			panic("TxTransaction fail !!!")
		}
		useGas := receipt.CumulativeGasUsed
		fmt.Println("useGas", useGas)
		fromBalanceAfter := utils.GetBalance(  fromAddress.String())
		toBalanceAfter := utils.GetBalance(  toAddress)
		managerBalanceAfter := utils.GetBalance(  "0x0000000000000000000000000000000000000007")
		fmt.Println("fromBalance", fromBalance)
		fmt.Println("toBalance", toBalance)
		fmt.Println("managerBalance", managerBalance)
		fmt.Println("fromBalanceAfter", fromBalanceAfter)
		fmt.Println("toBalanceAfter", toBalanceAfter)
		fmt.Println("managerBalanceAfter", managerBalanceAfter)
		//1 断言from账户资金是否 = fromBalance1 = fromBalance - gasUsed - amount
		fromBalance1 := fromBalanceAfter.Add(fromBalanceAfter, amount).Add(fromBalanceAfter, big.NewInt(int64(useGas)))
		So(fromBalance.Cmp(fromBalance1) == 0, ShouldBeTrue)
		//2 断言to账户资金是否 = toBalance + amount
		So(toBalanceAfter.Cmp(toBalance.Add(toBalance, amount)) == 0, ShouldBeTrue)
		//3 断言治理合约地址是否正常获得消耗gas
		So(managerBalanceAfter.Cmp(managerBalance.Add(managerBalance, big.NewInt(int64(useGas)))) == 0, ShouldBeTrue)
	})
}

func TestEVMTransfer003(t *testing.T) {
	Convey("Test003 测试EIP155交易普通转账ong(amount:256) ", t, func() {
		toAddress := "0xf45505D1F482EBc8881dacA97B122B62771B9e1d"
		amount := big.NewInt(256000000000)
		//  备注： 对应的MetaMask钱包第一个账户私钥
		fromPrivateKey := config.FromPrivate
		fromPrivate, _ := crypto.HexToECDSA(fromPrivateKey)
		fromAddress := crypto.PubkeyToAddress(fromPrivate.PublicKey)

		//交易前账户余额
		fromBalance := utils.GetBalance(  fromAddress.String())
		toBalance := utils.GetBalance(  toAddress)
		managerBalance := utils.GetBalance(  "0x0000000000000000000000000000000000000007")

		hash, err := deploy.SendTransfer(  fromPrivateKey, toAddress, amount, uint64(200000))
		if err != nil {
			panic(err)
		}
		receipt, err := utils.GetTransferInfoByHash(hash)
		if err != nil {
			panic(err)
		}
		//fmt.Println(receipt)
		if receipt.Status != 1 {
			fmt.Println("status : ", receipt.Status)
			panic("TxTransaction fail !!!")
		}
		useGas := receipt.CumulativeGasUsed
		fmt.Println("useGas", useGas)
		fromBalanceAfter := utils.GetBalance(  fromAddress.String())
		toBalanceAfter := utils.GetBalance(  toAddress)
		managerBalanceAfter := utils.GetBalance(  "0x0000000000000000000000000000000000000007")
		fmt.Println("fromBalance", fromBalance)
		fmt.Println("toBalance", toBalance)
		fmt.Println("managerBalance", managerBalance)
		fmt.Println("fromBalanceAfter", fromBalanceAfter)
		fmt.Println("toBalanceAfter", toBalanceAfter)
		fmt.Println("managerBalanceAfter", managerBalanceAfter)
		//1 断言from账户资金是否 = fromBalance1 = fromBalance - gasUsed - amount
		fromBalance1 := fromBalanceAfter.Add(fromBalanceAfter, amount).Add(fromBalanceAfter, big.NewInt(int64(useGas)))
		So(fromBalance.Cmp(fromBalance1) == 0, ShouldBeTrue)
		//2 断言to账户资金是否 = toBalance + amount
		So(toBalanceAfter.Cmp(toBalance.Add(toBalance, amount)) == 0, ShouldBeTrue)
		//3 断言治理合约地址是否正常获得消耗gas
		So(managerBalanceAfter.Cmp(managerBalance.Add(managerBalance, big.NewInt(int64(useGas)))) == 0, ShouldBeTrue)
	})
}

func TestEVMTransfer004(t *testing.T) {
	Convey("Test004 测试EIP155交易普通转账ong(amount:257) ", t, func() {
		toAddress := "0xf45505D1F482EBc8881dacA97B122B62771B9e1d"
		amount := big.NewInt(257000000000)
		//  备注： 对应的MetaMask钱包第一个账户私钥
		fromPrivateKey := config.FromPrivate
		fromPrivate, _ := crypto.HexToECDSA(fromPrivateKey)
		fromAddress := crypto.PubkeyToAddress(fromPrivate.PublicKey)

		//交易前账户余额
		fromBalance := utils.GetBalance(  fromAddress.String())
		toBalance := utils.GetBalance(  toAddress)
		managerBalance := utils.GetBalance(  "0x0000000000000000000000000000000000000007")

		hash, err := deploy.SendTransfer(  fromPrivateKey, toAddress, amount, uint64(200000))
		if err != nil {
			panic(err)
		}
		//fmt.Println(hash)
		receipt, err := utils.GetTransferInfoByHash(hash)
		if err != nil {
			panic(err)
		}
		//fmt.Println(receipt)
		if receipt.Status != 1 {
			fmt.Println("status : ", receipt.Status)
			panic("TxTransaction fail !!!")
		}
		useGas := receipt.CumulativeGasUsed
		fmt.Println("useGas", useGas)
		fromBalanceAfter := utils.GetBalance(  fromAddress.String())
		toBalanceAfter := utils.GetBalance(  toAddress)
		managerBalanceAfter := utils.GetBalance(  "0x0000000000000000000000000000000000000007")
		fmt.Println("fromBalance", fromBalance)
		fmt.Println("toBalance", toBalance)
		fmt.Println("managerBalance", managerBalance)
		fmt.Println("fromBalanceAfter", fromBalanceAfter)
		fmt.Println("toBalanceAfter", toBalanceAfter)
		fmt.Println("managerBalanceAfter", managerBalanceAfter)
		//1 断言from账户资金是否 = fromBalance1 = fromBalance - gasUsed - amount
		fromBalance1 := fromBalanceAfter.Add(fromBalanceAfter, amount).Add(fromBalanceAfter, big.NewInt(int64(useGas)))
		So(fromBalance.Cmp(fromBalance1) == 0, ShouldBeTrue)
		//2 断言to账户资金是否 = toBalance + amount
		So(toBalanceAfter.Cmp(toBalance.Add(toBalance, amount)) == 0, ShouldBeTrue)
		//3 断言治理合约地址是否正常获得消耗gas
		So(managerBalanceAfter.Cmp(managerBalance.Add(managerBalance, big.NewInt(int64(useGas)))) == 0, ShouldBeTrue)
	})
}

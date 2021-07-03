package example

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

func TestExample(t *testing.T) {
	Convey("TestTestExample", t, func() {
		// 得到SDK的实例
		//ontSdk := goSdk.NewOntologySdk()
		//ontSdk.NewRpcClient().SetAddress("http://localhost:20336")
		//// 生成native合约调用交易; 如果想调用NEO VM合约，可以使用NewNeoVMSInvokeTransaction方法
		//addr,_ := common.AddressFromHexString("Ae9mEb7VZQj1Pd4rsFZytzhnUNJ95JnbmT")
		//res,_ := ontSdk.Native.Ong.BalanceOf(addr)
		//fmt.Println(res)
		//ethcli, err := ethclient.Dial(config.URL)
		//if err != nil {
		//	panic(err)
		//}

		url := config.URL
		toAddress := "0x0000000000000000000000000000000000000007"
		amount := big.NewInt(1000000000)
		//  备注
		fromPrivateKey := config.FromPrivate
		fromPrivate, _ := crypto.HexToECDSA(fromPrivateKey)
		fromAddress := crypto.PubkeyToAddress(fromPrivate.PublicKey)
		fmt.Println("fromAddress: ", fromAddress)

		//交易前账户余额
		//fromBalance := utils.GetBalance(url, fromAddress.String())
		//toBalance := utils.GetBalance(url, toAddress)
		//managerBalance := utils.GetBalance(url, "0x0000000000000000000000000000000000000007")

		hash, err := deploy.SendTransfer(url, fromPrivateKey, toAddress, amount, uint64(200000))
		if err != nil {
			panic(err)
		}
		res, err := utils.GetTransferInfoByHash(hash)
		fmt.Println(err)
		fmt.Println(res.Status)

		// cversion 为合约的版本, method 是要调用的合约方法名, params 是该方法需要的参数
		// 例如：
		// NewNativeInvokeTransaction(0, 200000, byte(0),utils.ParamContractAddress,
		//      "getGlobalParam", []interface{}{global_params.ParamNameList{"gasPrice"}})

		//tx, err := rpcClient.NewNativeInvokeTransaction(gasPrice, gasLimit, cversion, contractAddress, method, params)
		//if err != nil {
		//	return common.UINT256_EMPTY, err
		//}
		//// 对交易签名，signer为交易的发送者
		//err = rpcClient.SignToTransaction(tx, signer)
		//if err != nil {
		//	return common.UINT256_EMPTY, err
		//}
		//
		//txbf := new(bytes.Buffer)
		//err = tx.Serialize(txbf);
		//hexCode = common.ToHexString(txbf.Bytes())
		//client, err := account.Open("./wallet.dat")
		//if err != nil {
		//	panic(err)
		//}
		//password := "123456"
		//account, err := client.GetAccountByIndex(1, []byte(password))
		//fmt.Println(err)
		//if err != nil {
		//	panic(err)
		//}
		//fmt.Println(account.Address)
		//fmt.Println(account.PrivateKey)
		//fmt.Println(account.PublicKey)
		//fmt.Println(common.PubKeyToHex(account.PublicKey))

		//url :=  config.URL
		//toAddress := "0xf45505D1F482EBc8881dacA97B122B62771B9e1d"
		//amount := big.NewInt(1000000000)
		//  备注： 对应的MetaMask钱包第一个账户私钥
		//fromPrivateKey := "b03fce4f5df24210e77de096bf33e2765f273cbe377f3ccc997dafdf90c9b01c"
		//fromPrivate, _ := crypto.HexToECDSA(fromPrivateKey)
		//fromAddress := crypto.PubkeyToAddress(fromPrivate.PublicKey)
		//fmt.Println(fromPrivate)
		//fmt.Println(fromAddress)
		//fmt.Println(fromPrivate.PublicKey)

		//ethClient, _ := ethclient.Dial(url)
		//defer ethClient.Close()

		//	//交易前账户余额
		//	fromBalance := utils.GetBalance(url, account.Address())
		//	toBalance := utils.GetBalance(url, toAddress)
		//	managerBalance := utils.GetBalance(url, "0x0000000000000000000000000000000000000007")
		//
		//	hash,err := deploy.SendTransfer(url, fromPrivateKey, toAddress, amount, uint64(200000))
		//	if err != nil {
		//		panic(err)
		//	}
		//	//fmt.Println(hash)
		//	receipt := make(map[string]interface{})
		//	for i := 0; i < 10; i++ { //设置10次请求，每次间隔3s 30s超时
		//		time.Sleep(3000000000)
		//		if len(receipt) != 0 {
		//			break
		//		} else {
		//			receipt = utils.GetTransferInfo(url, hash)
		//		}
		//	}
		//	if len(receipt) == 0 {
		//		panic("交易30s未落账!")
		//	}
		//	//fmt.Println(receipt)
		//})
	})
}

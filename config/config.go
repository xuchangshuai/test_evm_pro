package config

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	ontology_go_sdk "github.com/ontio/ontology-go-sdk"
)

const URL string = "http://127.0.0.1:20339"

const SDK_URL string = "http://127.0.0.1:20336"

const FromPrivate string = "80a68081edc0aed4ddf8fa9f6a2e7cf8d0a69c998d4ef646f6446cbf4cfe9145"

func GetEthClient() *ethclient.Client {
	client, err := ethclient.Dial(URL)
	if err != nil {
		panic(err)
	}
	return client
}

func GetAddressList() [5]common.Address { //账户列表
	return [5]common.Address{
		common.HexToAddress("0xf45505D1F482EBc8881dacA97B122B62771B9e1d"),
		common.HexToAddress("0xb71380C3bcfB8469692c5a18BE1974897764c46A"),
		common.HexToAddress("0x1cd1A95C6366e03527E892Ac49cd7C4cEFEF277D"),
		common.HexToAddress("0x68C60D4557e03dA58B0406B5A5058533c58D6c75"),
		common.HexToAddress("0x60271dd6F02d3661Ed9C8829A211da2c463c13f5"),
	}
}

func GetPrivateList() [5]string { // 账户对应的私钥
	return [5]string{
		"cb8f6ca317340db01cb8787218576153f7a9e996dd9e7953fc54fd74db301d57",
		"9f084729bd77604f8951a68514447c590f088930a7bbf27fece6d298f028b1ed",
		"754a993bb73798c4fa37415c78b3d8bdce20d1889a083fd9e209a0cd0d839eb7",
		"f9499d66522b8b5bc6c3833ac3bc7d64437ff72e7d2880b88559368f664dee2a",
		"4ed22b50a7dc27afedfeefb4127ad952dfeb47f838ffbc7745614a0a894a457a",
	}
}

func GetSdk() *ontology_go_sdk.OntologySdk{
	testOntSdk := ontology_go_sdk.NewOntologySdk()
	testOntSdk.NewRpcClient().SetAddress(SDK_URL)
	return testOntSdk
}
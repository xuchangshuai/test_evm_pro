package config

import "github.com/ethereum/go-ethereum/ethclient"

const URL string = "http://127.0.0.1:20339"

const FromPrivate string = "80a68081edc0aed4ddf8fa9f6a2e7cf8d0a69c998d4ef646f6446cbf4cfe9145"

func GetEthClient() *ethclient.Client {
	client, err := ethclient.Dial(URL)
	if err != nil {
		panic(err)
	}
	return client
}

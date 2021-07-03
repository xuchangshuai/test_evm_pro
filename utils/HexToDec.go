package utils

import (
	"math/big"
	"strings"
)

func HexToDec(hex string) *big.Int {
	if strings.HasPrefix(hex, "0x") {
		res, _ := new(big.Int).SetString(hex, 0)
		return res
	}
	return nil
}

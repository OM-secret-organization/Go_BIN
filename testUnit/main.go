package main

import (
	"fmt"
	"math/big"
	"reflect"

	"github.com/ethereum/go-ethereum/params"
)

func main() {
	n := new(big.Int)
	n, ok := n.SetString("99972712335400000000", 10)
	if !ok {
			fmt.Println("SetString: error")
			return
	}
	bigFloat := weiToEther(n)
	str := bigFloat.Text('f', -1)
	fmt.Println(bigFloat, reflect.TypeOf(bigFloat))
  fmt.Println(str, reflect.TypeOf(str))	
}


// func ChangeValueUnit(value interface{}, chainName string) (string, error) {
// 	switch chainName{
// 	case "BTC":
// 		// float64 -> string
// 		strVal := fmt.Sprintf("%f", value)
// 		return strVal, nil
// 	case "ETH":
// 		// big.Int -> string
// 		val := value.(big.Int)
// 		ethVal := weiToEtherStr(&val)
// 		ethVal.Text(f, int(precision))
// 		fmt.Println("BTC unit change")
// 		return ethValStr, nil
// 	case "XRP":
// 		fmt.Println("BTC unit change")
// 		return "1", nil
// 	}
// 	return "", errors.New("chain name error")
// }

func weiToEther(wei *big.Int) *big.Float {
	return new(big.Float).Quo(new(big.Float).SetInt(wei), big.NewFloat(params.Ether))
}
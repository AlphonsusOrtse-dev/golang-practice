package main

import (
	"fmt"
	"strconv"
)

func hexToDecimal(hexStr string) (int64, error) {

	// 	this is another short way of converting from hex 2 dec
	// return strconv.ParseInt(hexStr, 16, 64)

	// 	value, err := strconv.ParseInt(hexStr, 16, 64)

	// 	if err != nil {
	// 		return 0, err
	// 	}

	// return value, nil

	//converting from dec to binary....	fmt.Println(hexToDecimal("1001"))

	return strconv.ParseInt(hexStr, 2, 64)

}
func main() {
	fmt.Println(hexToDecimal("E1"))
}

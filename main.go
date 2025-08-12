package main

import (
	"fmt"
	"math/big"
	"strings"
)

func main() {
	fmt.Println("Enter hex number or 'stop' to exit:")

	var input string

	for {
		fmt.Scanln(&input)

		if strings.ToLower(input) == "stop" {
			break
		}

		i := new(big.Int)

		if _, ok := i.SetString(processHex(input), 16); !ok {
			fmt.Println("Invalid hex number")
			continue
		}

		fmt.Println(i)
	}
}

func processHex(hexStr string) string {
	return strings.TrimPrefix(hexStr, "0x")
}

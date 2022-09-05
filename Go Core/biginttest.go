package main

import "fmt"
import "math/big"

func main() {
	intHash := big.NewInt(3000)
	target := big.NewInt(2000)
	if intHash.Cmp(target) == -1 {
		fmt.Println("Check is -1", intHash.Cmp(target))
	} else {
		fmt.Println("Check is what?", intHash.Cmp(target))
	}
}

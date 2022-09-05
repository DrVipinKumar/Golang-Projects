package main

import (
	"fmt"
	"myblockchain/blockchain"
)

func main() {
	chain := blockchain.InitBlockchain()
	chain.AddBlock("First block after Genesis")
	chain.AddBlock("Second block after Genesis")
	chain.AddBlock("Third block after Genesis")
	for _, block := range chain.Blocks {
		fmt.Printf("Previous Hash=%x\n", block.PrevHash)
		fmt.Printf("Data=%s\n", block.Data)
		fmt.Printf("Hash=%x\n", block.Hash)
		pow := blockchain.NewProof(block)
		fmt.Println("Check Valid=>", pow.Validate())

	}

}

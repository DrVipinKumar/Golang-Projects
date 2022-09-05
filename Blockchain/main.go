package main

import "fmt"
import "myblockchain/blockchain"
import "strconv"

func main() {
	chain := blockchain.InitBlockChain()
	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("second block")
	chain.AddBlock("third block")
	for _, block := range chain.Blocks {
		fmt.Printf("Previous Hash:%x\n", block.PrevHash)
		fmt.Printf("Data:%s\n", block.Data)
		fmt.Printf("Hash:%x\n", block.Hash)
		pow := blockchain.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}

}

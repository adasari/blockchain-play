package main

import (
	"blockchain-play/chain"
	"fmt"
	"strconv"
)

func main() {

	genesis := chain.Genesis()
	bchain := chain.NewBlockchain(genesis)

	bchain.AddBlock("First Block after Genesis")
	bchain.AddBlock("Second Block after Genesis")
	bchain.AddBlock("Third Block after Genesis")

	for _, block := range bchain.Blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := chain.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}

package main

import (
	"fmt"
	"strconv"

	"github.com/YoungsoonLee/blockchain-go/block"
)

func main() {
	bc := block.NewBlockchain() // make genesis block

	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")

	for _, bl := range bc.Blocks {
		fmt.Printf("Prev. hash: %x\n", bl.PrevBlockHash)
		fmt.Printf("Data: %s\n", bl.Data)
		fmt.Printf("Hash: %x\n", bl.Hash)

		pow := block.NewProofOfWork(bl)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))

		fmt.Println()
	}
}

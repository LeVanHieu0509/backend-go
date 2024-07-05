package main

import (
	"fmt"

	"github.com/LeVanHieu0509/backend-go/blockchain"
)

func main() {
	// Khởi tạo chuỗi khối mới và nhận về con trỏ tới BlockChain.
	chain := blockchain.InitBlockChain()

	// Thêm các khối mới vào chuỗi khối.
	chain.AddBlock("First Block")
	chain.AddBlock("Second Block")
	chain.AddBlock("Three Block")

	//miner
	chain.AddBlock("Four Block")
	chain.AddBlock("Five Block")

	// Lặp qua các con trỏ trong mảng blocks, in ra thông tin của từng khối.
	for _, block := range chain.Blocks {
		fmt.Println("-----------------------------------------------------------------------")
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println("-----------------------------------------------------------------------")

		// pow := blockchain.NewProof(block)
		// fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		// fmt.Println()
	}

}

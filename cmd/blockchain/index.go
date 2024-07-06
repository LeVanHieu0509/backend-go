package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"strconv"

	"github.com/LeVanHieu0509/backend-go/blockchain"
)

type CommandLine struct {
	blockchain *blockchain.BlockChain
}

func (cli *CommandLine) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("add - block BLOCK_DATA - add a block to the chain:")
	fmt.Println("print - Prints the blocks in the chain")

}
func (cli *CommandLine) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		runtime.Goexit()
	}

}

func (cli *CommandLine) addBlock(data string) {
	cli.blockchain.AddBlock((data))

	fmt.Println("Add block!")
}

func (cli *CommandLine) printChain() {
	iter := cli.blockchain.Iterator()

	for {
		block := iter.Next()

		fmt.Printf("Prev. Hash: %x\n", block.PrevHash)
		fmt.Printf("Data. Hash: %s\n", block.Data)
		fmt.Printf("Hash. Hash: %x\n", block.Hash)

		pow := blockchain.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		if len(block.PrevHash) == 0 {
			break
		}

	}
	fmt.Println("Add block!")
}

func (cli *CommandLine) run() {
	cli.validateArgs()

	addBlockCmd := flag.NewFlagSet("add", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("print", flag.ExitOnError)

	addBlockData := addBlockCmd.String("block", "", "Block data")

	switch os.Args[1] {
	case "add":
		err := addBlockCmd.Parse(os.Args[2:])
		blockchain.Handle(err)

	case "print":
		err := printChainCmd.Parse(os.Args[2:])
		blockchain.Handle(err)
	default:
		cli.printUsage()
		runtime.Goexit()
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			runtime.Goexit()
		}
		cli.addBlock(*addBlockData)
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}

}

func main() {
	// Khởi tạo chuỗi khối mới và nhận về con trỏ tới BlockChain.
	// chain := blockchain.InitBlockChain()

	// // Thêm các khối mới vào chuỗi khối.
	// chain.AddBlock("First Block")
	// chain.AddBlock("Second Block")
	// chain.AddBlock("Three Block")

	// //miner
	// chain.AddBlock("Four Block")
	// chain.AddBlock("Five Block")

	// Lặp qua các con trỏ trong mảng blocks, in ra thông tin của từng khối.
	// for _, block := range chain.Blocks {
	// 	fmt.Println("-----------------------------------------------------------------------")
	// 	fmt.Printf("Previous Hash: %x\n", block.PrevHash)
	// 	fmt.Printf("Data in Block: %s\n", block.Data)
	// 	fmt.Printf("Hash: %x\n", block.Hash)
	// 	fmt.Println("-----------------------------------------------------------------------")

	// 	pow := blockchain.NewProof(block)
	// 	fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
	// 	fmt.Println()
	// }

	//
	defer os.Exit(0)
	chain := blockchain.InitBlockChain()
	defer chain.Database.Close()

	cli := CommandLine{chain}
	cli.run()

}

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
	fmt.Println("get balance address ADDRESS - get the balance:")
	fmt.Println("create blockchain - address Address")

	fmt.Println("add - block BLOCK_DATA - add a block to the chain:")
	fmt.Println("print chain - Prints the blocks in the chain")
	println("send - FROM to TO amount")
	fmt.Println("-----------------------------------------------------------------------------")

}
func (cli *CommandLine) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		runtime.Goexit()
	}

}

// func (cli *CommandLine) addBlock(data string) {
// 	cli.blockchain.AddBlock((data))

// 	fmt.Println("Add block!")
// }

func (cli *CommandLine) printChain() {
	chain := blockchain.ContinueBlockchain("")
	defer chain.Database.Close()

	iter := chain.Iterator()

	for {
		block := iter.Next()

		fmt.Printf("Prev Hash: %x\n", block.PrevHash)
		fmt.Printf("Hash Block: %x\n", block.Hash)

		pow := blockchain.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		for _, tx := range block.Transactions {
			fmt.Println("Transaction ID:", tx.ID)
			fmt.Println("Transaction Inputs:", tx.Inputs)
			fmt.Println("Transaction Outputs:", tx.Outputs)
		}
		fmt.Println("-----------------------------------------------------------------------------")

		if len(block.PrevHash) == 0 {
			break
		}

	}

}

func (cli *CommandLine) createBlockChain(address string) {
	chain := blockchain.InitBlockChain(address)
	chain.Database.Close()

	fmt.Println("Finished!")
}

func (cli *CommandLine) getBalance(address string) {
	chain := blockchain.ContinueBlockchain(address)

	defer chain.Database.Close()

	balance := 0
	UTXOs := chain.FindUTXO(address)

	for _, out := range UTXOs {
		balance += out.Value
	}

	fmt.Printf("Balance of %s: %d\n", address, balance)
}

func (cli *CommandLine) send(from, to string, amount int) {
	chain := blockchain.ContinueBlockchain(from)
	defer chain.Database.Close()

	tx := blockchain.NewTransaction(from, to, amount, chain)
	chain.AddBlock([]*blockchain.Transaction{tx})

	fmt.Println("Success!")
}

// Part 3
// func (cli *CommandLine) run() {
// 	cli.validateArgs()

// 	addBlockCmd := flag.NewFlagSet("add", flag.ExitOnError)
// 	printChainCmd := flag.NewFlagSet("print", flag.ExitOnError)

// 	addBlockData := addBlockCmd.String("block", "", "Block data")

// 	switch os.Args[1] {
// 	case "add":
// 		err := addBlockCmd.Parse(os.Args[2:])
// 		blockchain.Handle(err)

// 	case "print":
// 		err := printChainCmd.Parse(os.Args[2:])
// 		blockchain.Handle(err)
// 	default:
// 		cli.printUsage()
// 		runtime.Goexit()
// 	}

// 	if addBlockCmd.Parsed() {
// 		if *addBlockData == "" {
// 			addBlockCmd.Usage()
// 			runtime.Goexit()
// 		}
// 		cli.addBlock(*addBlockData)
// 	}

// 	if printChainCmd.Parsed() {
// 		cli.printChain()
// 	}

// }

// Part 4
func (cli *CommandLine) run() {
	cli.validateArgs() //Hàm validateArgs kiểm tra các tham số đầu vào để đảm bảo rằng chúng hợp lệ.
	fmt.Println("---run---")

	// Khởi tạo các lệnh con (FlagSet) để xử lý các lệnh cụ thể
	getBalanceCmd := flag.NewFlagSet("getbalance", flag.ExitOnError)
	createBlockchainCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)
	sendCmd := flag.NewFlagSet("send", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)

	// Định nghĩa các tham số cho từng lệnh con
	getBalanceAddress := getBalanceCmd.String("address", "", "The address to get balance for")
	createBlockchainAddress := createBlockchainCmd.String("address", "", "create blockchain address")
	sendFrom := sendCmd.String("from", "", "Source wallet address")
	sendTo := sendCmd.String("to", "", "Destination wallet address")
	sendAmount := sendCmd.Int("amount", 0, "Amount to send")

	// Xử lý các lệnh con dựa trên tham số đầu vào
	switch os.Args[1] {
	case "getbalance":
		err := getBalanceCmd.Parse(os.Args[2:])
		blockchain.Handle(err)

	case "createblockchain":
		err := createBlockchainCmd.Parse(os.Args[2:])
		blockchain.Handle(err)
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		blockchain.Handle(err)
	case "send":
		err := sendCmd.Parse(os.Args[2:])
		blockchain.Handle(err)
	default:
		cli.printUsage()
		runtime.Goexit()
	}

	// Gọi phương thức Parse để phân tích các tham số của lệnh con và xử lý lỗi nếu có.
	if getBalanceCmd.Parsed() {
		if *getBalanceAddress == "" {
			getBalanceCmd.Usage()
			runtime.Goexit()
		}
		cli.getBalance(*getBalanceAddress)
	}

	if createBlockchainCmd.Parsed() {
		if *createBlockchainAddress == "" {
			createBlockchainCmd.Usage()
			runtime.Goexit()
		}
		cli.createBlockChain(*createBlockchainAddress)
	}

	if sendCmd.Parsed() {
		if *sendFrom == "" || *sendTo == "" || *sendAmount <= 0 {
			sendCmd.Usage()
			os.Exit(1)
		}
		cli.send(*sendFrom, *sendTo, *sendAmount)
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

	// part 3
	// defer os.Exit(0)
	// chain := blockchain.InitBlockChain()
	// defer chain.Database.Close()

	// cli := CommandLine{chain}
	// cli.run()

	//part 4
	defer os.Exit(0)
	cli := CommandLine{}
	cli.run()

}

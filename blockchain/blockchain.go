package blockchain

import (
	"encoding/hex"
	"fmt"
	"os"
	"runtime"

	"github.com/dgraph-io/badger"
)

const (
	dbPath      = "./tmp/blocks"
	dbFile      = "./tmp/blocks/MANIFEST"
	genesisData = "First Transaction from Genesis"
)

// Cấu trúc chính của blockchain, chứa thông tin về hash của khối cuối cùng và cơ sở dữ liệu.
type BlockChain struct {

	// blocks là một mảng các con trỏ tới các Block. Mỗi phần tử trong mảng này là một con trỏ đến một đối tượng Block.
	// Blocks []*Block

	LashHash []byte     //Lưu trữ hash của khối cuối cùng trong blockchain.
	Database *badger.DB //Cơ sở dữ liệu Badger để lưu trữ các khối.
}

// Cấu trúc giúp duyệt qua các khối trong blockchain.
type BlockChainIterator struct {
	CurrentHash []byte     //Lưu trữ hash của khối hiện tại trong quá trình duyệt.
	Database    *badger.DB //Cơ sở dữ liệu Badger.
}

// Part 4.
func DBexists() bool {
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		return false
	}
	return true
}

// InitBlockChain khởi tạo một chuỗi khối mới với khối gốc.
func InitBlockChain(address string) *BlockChain {
	var lashHash []byte

	if DBexists() {
		fmt.Println("Blockchain already exists")
		runtime.Goexit()
	}

	opts := badger.DefaultOptions(dbPath)

	opts.ValueDir = dbPath

	//Mở cơ sở dữ liệu Badger tại đường dẫn dbPath.
	db, err := badger.Open(opts)

	Handle(err)

	// err = db.Update(func(txn *badger.Txn) error {
	// 	// Kiểm tra xem blockchain có tồn tại trong cơ sở dữ liệu không bằng cách tìm khóa lh.
	// 	// Nếu không tìm thấy, tạo khối gốc (genesis block) và lưu trữ nó

	// 	if _, err := txn.Get([]byte("lh")); err == badger.ErrKeyNotFound {
	// 		fmt.Println("No existing blockchain found")

	// 		genesis := Genesis()

	// 		fmt.Println("genesis proved")

	// 		err := txn.Set(genesis.Hash, genesis.Serialize())

	// 		Handle(err)

	// 		err = txn.Set([]byte("lh"), genesis.Hash)

	// 		lashHash = genesis.Hash

	// 		return err
	// 	} else {
	// 		// lấy hash của khối cuối cùng (lashHash) từ cơ sở dữ liệu
	// 		item, err := txn.Get([]byte("lh"))
	// 		Handle(err)

	// 		lashHash, err = item.ValueCopy(nil)
	// 		return err
	// 	}
	// })

	//Part 4
	err = db.Update(func(txn *badger.Txn) error {
		cbtx := CoinbaseTx(address, genesisData)
		genesis := Genesis(cbtx)

		fmt.Println("Genesis created")
		err := txn.Set(genesis.Hash, genesis.Serialize())

		Handle(err)
		err = txn.Set([]byte("lh"), genesis.Hash)

		lashHash = genesis.Hash

		return err
	})

	Handle(err)

	// Khởi tạo và trả về đối tượng BlockChain.
	blockchain := BlockChain{lashHash, db}

	return &blockchain
	//Genesis tạo và trả về con trỏ tới khối gốc (genesis block).
	// return &BlockChain{[]*Block{Genesis()}}
}

// Part 4
func ContinueBlockchain(address string) *BlockChain {
	if DBexists() == false {
		fmt.Println("No Existing blockchain found, create one")
		runtime.Goexit()
	}

	var lashHash []byte

	opts := badger.DefaultOptions(dbPath)
	opts.Dir = dbPath
	opts.ValueDir = dbPath

	db, err := badger.Open(opts)
	Handle(err)

	err = db.Update(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("lh"))
		Handle(err)
		lashHash, err = item.ValueCopy(nil)

		return err
	})

	Handle(err)

	chain := BlockChain{lashHash, db}

	return &chain
}

// AddBlock thêm một khối mới vào chuỗi khối.
// Nó lấy khối cuối cùng trong chuỗi khối (prevBlock),
// tạo ra một khối mới với dữ liệu mới và hash của khối trước đó, sau đó thêm khối mới vào chuỗi khối.

func (chain *BlockChain) AddBlock(transactions []*Transaction) {
	// prevBlock := chain.Blocks[len(chain.Blocks)-1] //Lấy con trỏ đến khối cuối cùng trong chuỗi khối.
	// new := CreateBlock(data, prevBlock.Hash)       //Tạo một khối mới và trả về con trỏ tới khối đó.
	// chain.Blocks = append(chain.Blocks, new)       //Thêm con trỏ của khối mới vào mảng blocks.

	var lastHash []byte

	//Lấy hash của khối cuối cùng từ cơ sở dữ liệu (lastHash).
	err := chain.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("lh"))

		Handle(err)

		lastHash, err = item.ValueCopy(nil)

		return err

	})

	Handle(err)

	// Tạo khối mới với dữ liệu và hash của khối cuối cùng.
	newBlock := CreateBlock(transactions, lastHash)

	// Cập nhật cơ sở dữ liệu với khối mới và hash của nó.

	// Cập nhật LashHash trong đối tượng BlockChain.
	err = chain.Database.Update(func(txn *badger.Txn) error {
		err := txn.Set(newBlock.Hash, newBlock.Serialize())

		Handle(err)

		err = txn.Set([]byte("lh"), newBlock.Hash)

		chain.LashHash = newBlock.Hash
		return err
	})
	Handle(err)
}

// Các hàm để duyệt qua chuỗi khối.
// Tạo đối tượng BlockChainIterator với hash của khối cuối cùng.
func (chain *BlockChain) Iterator() *BlockChainIterator {
	iter := &BlockChainIterator{chain.LashHash, chain.Database}

	return iter
}

// Các hàm để duyệt qua chuỗi khối.
// Lấy khối hiện tại từ cơ sở dữ liệu và cập nhật CurrentHash để trỏ đến khối trước đó.
func (iter *BlockChainIterator) Next() *Block {
	var block *Block

	err := iter.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get(iter.CurrentHash)

		Handle(err)

		encodedBlock, err := item.ValueCopy(nil)
		block = Deserialize(encodedBlock)

		return err
	})

	Handle(err)
	iter.CurrentHash = block.PrevHash

	return block
}

func (chain *BlockChain) FindUnSpentTransactions(address string) []Transaction {
	var unspentTxs []Transaction

	spentTXOs := make(map[string][]int)

	iter := chain.Iterator()

	for {

		block := iter.Next()
		for _, tx := range block.Transactions {
			txID := hex.EncodeToString(tx.ID)
		Outputs:
			for outIdx, out := range tx.Outputs {
				if spentTXOs[txID] != nil {
					for _, spentOut := range spentTXOs[txID] {
						if spentOut == outIdx {
							continue Outputs
						}
					}
				}

				if out.CanBeUnlocked(address) {
					unspentTxs = append(unspentTxs, *tx)
				}
			}

			if tx.IsCoinbase() == false {
				for _, in := range tx.Inputs {
					if in.CanUnlock(address) {
						inTxID := hex.EncodeToString(in.ID)
						spentTXOs[inTxID] = append(spentTXOs[inTxID], in.Out)
					}
				}
			}
		}

		if len(block.PrevHash) == 0 {
			break
		}
	}

	return unspentTxs
}

func (chain *BlockChain) FindUTXO(address string) []TxOutput {
	var UTXOs []TxOutput
	unspentTransactions := chain.FindUnSpentTransactions(address)

	for _, tx := range unspentTransactions {
		for _, out := range tx.Outputs {
			if out.CanBeUnlocked(address) {
				UTXOs = append(UTXOs, out)
			}
		}
	}
	return UTXOs

}

func (chain *BlockChain) FindSpendableOutputs(address string, amount int) (int, map[string][]int) {
	unspentOuts := make(map[string][]int)
	unspentTxs := chain.FindUnSpentTransactions(address)
	accumulated := 0

Work:
	for _, tx := range unspentTxs {
		txID := hex.EncodeToString(tx.ID)

		for outIdx, out := range tx.Outputs {
			if out.CanBeUnlocked(address) && accumulated < amount {
				accumulated += out.Value
				unspentOuts[txID] = append(unspentOuts[txID], outIdx)

				if accumulated >= amount {
					break Work
				}
			}
		}

	}

	return accumulated, unspentOuts

}

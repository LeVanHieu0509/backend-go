package blockchain

import (
	"fmt"

	"github.com/dgraph-io/badger"
)

const (
	dbPath = "./tmp/blocks"
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

// InitBlockChain khởi tạo một chuỗi khối mới với khối gốc.
func InitBlockChain() *BlockChain {
	var lashHash []byte
	opts := badger.DefaultOptions(dbPath)

	opts.ValueDir = dbPath

	//Mở cơ sở dữ liệu Badger tại đường dẫn dbPath.
	db, err := badger.Open(opts)
	Handle(err)

	err = db.Update(func(txn *badger.Txn) error {
		// Kiểm tra xem blockchain có tồn tại trong cơ sở dữ liệu không bằng cách tìm khóa lh.
		// Nếu không tìm thấy, tạo khối gốc (genesis block) và lưu trữ nó

		if _, err := txn.Get([]byte("lh")); err == badger.ErrKeyNotFound {
			fmt.Println("No existing blockchain found")

			genesis := Genesis()

			fmt.Println("genesis proved")

			err := txn.Set(genesis.Hash, genesis.Serialize())

			Handle(err)

			err = txn.Set([]byte("lh"), genesis.Hash)

			lashHash = genesis.Hash

			return err
		} else {
			// lấy hash của khối cuối cùng (lashHash) từ cơ sở dữ liệu
			item, err := txn.Get([]byte("lh"))
			Handle(err)

			lashHash, err = item.ValueCopy(nil)
			return err
		}
	})

	Handle(err)

	// Khởi tạo và trả về đối tượng BlockChain.
	blockchain := BlockChain{lashHash, db}

	return &blockchain
	//Genesis tạo và trả về con trỏ tới khối gốc (genesis block).
	// return &BlockChain{[]*Block{Genesis()}}
}

// AddBlock thêm một khối mới vào chuỗi khối.
// Nó lấy khối cuối cùng trong chuỗi khối (prevBlock),
// tạo ra một khối mới với dữ liệu mới và hash của khối trước đó, sau đó thêm khối mới vào chuỗi khối.

func (chain *BlockChain) AddBlock(data string) {
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
	newBlock := CreateBlock(data, lastHash)

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

package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

/*
Chuỗi khối là một cấu trúc dữ liệu nơi các khối (blocks) được nối với nhau qua các hàm băm (hash) của khối trước đó
*/

// BlockChain là một chuỗi các khối, được lưu trữ dưới dạng một mảng các con trỏ tới các khối.

type Block struct {
	Hash []byte //Đây là hash của khối hiện tại, được tính từ dữ liệu (Data) và hash của khối trước đó
	// Data     []byte //Dữ liệu của khối hiện tại.

	// Change Data => Transaction in Part 4
	Transactions []*Transaction

	PrevHash []byte //Hash của khối trước đó trong chuỗi khối.
	Nonce    int
}

// DeriveHash tính toán hash của khối bằng cách kết hợp dữ liệu của khối và hash của khối trước đó,
// sau đó sử dụng thuật toán SHA-256 để tính hash và gán nó cho thuộc tính Hash.

//part 2
// func (b *Block) DeriveHash() {
// 	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
// 	hash := sha256.Sum256((info))
// 	b.Hash = hash[:]
// }

// part 4 - Không sài nữa
// func (b *Block) DeriveHash() {
// 	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
// 	hash := sha256.Sum256((info))
// 	b.Hash = hash[:]
// }

// CreateBlock tạo ra một khối mới với dữ liệu (data) và hash của khối trước đó (preHash).
// Sau đó, nó gọi phương thức DeriveHash để tính hash cho khối mới và trả về khối đó.
// func CreateBlock(data string, preHash []byte) *Block {

// 	// CreateBlock tạo một con trỏ tới một Block mới.
// 	// Nó khởi tạo một đối tượng Block, tính toán hash của khối đó bằng cách gọi block.DeriveHash()
// 	// và trả về con trỏ tới đối tượng Block đó.

// 	block := &Block{[]byte{}, []byte(data), preHash, 0}
// 	// block.DeriveHash()

// 	//Video 2
// 	pow := NewProof(block)
// 	nonce, hash := pow.Run()

// 	block.Hash = hash[:]
// 	block.Nonce = nonce

// 	return block
// }

// part 4.
func (b *Block) HashTransactions() []byte {
	var txHashes [][]byte
	var txHash [32]byte

	for _, tx := range b.Transactions {
		txHashes = append(txHashes, tx.ID)
	}
	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))

	return txHash[:]
}

// Transaction part 4
func CreateBlock(tsx []*Transaction, preHash []byte) *Block {

	// CreateBlock tạo một con trỏ tới một Block mới.
	// Nó khởi tạo một đối tượng Block, tính toán hash của khối đó bằng cách gọi block.DeriveHash()
	// và trả về con trỏ tới đối tượng Block đó.

	block := &Block{[]byte{}, tsx, preHash, 0}
	// block.DeriveHash()

	//Video 2
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// Genesis tạo ra khối đầu tiên trong chuỗi khối, gọi là khối gốc (genesis block).
// Khối này có dữ liệu là "Genesis" và không có hash của khối trước đó (vì nó là khối đầu tiên).
// func Genesis() *Block {
// 	return CreateBlock("Genesis", []byte{})
// }

// transaction
func Genesis(coinbase *Transaction) *Block {
	return CreateBlock([]*Transaction{coinbase}, []byte{})
}

// Hàm Serialize thực hiện tuần tự hóa (serialize) đối tượng Block thành một mảng byte
func (b *Block) Serialize() []byte {
	// Tạo một bytes.Buffer để lưu trữ dữ liệu đã tuần tự hóa.
	var res bytes.Buffer

	// Tạo một encoder gob.NewEncoder từ buffer để tuần tự hóa dữ liệu.
	encoder := gob.NewEncoder(&res)

	// Sử dụng encoder để encode đối tượng Block.
	err := encoder.Encode(b)
	Handle(err)

	//Trả về mảng byte từ buffer.
	return res.Bytes()
}

// Hàm Deserialize thực hiện giải tuần tự hóa (deserialize) một mảng byte thành đối tượng Block
func Deserialize(data []byte) *Block {
	// Tạo một đối tượng Block rỗng để lưu trữ dữ liệu giải tuần tự hóa.
	var block Block

	// Tạo một decoder gob.NewDecoder từ một bytes.NewReader được tạo từ mảng byte đầu vào
	decoder := gob.NewDecoder((bytes.NewReader(data)))

	// Sử dụng decoder để decode mảng byte vào đối tượng Block.
	err := decoder.Decode(&block)

	Handle(err)

	// Trả về con trỏ tới đối tượng Block đã giải tuần tự hóa.
	return &block
}

func Handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}

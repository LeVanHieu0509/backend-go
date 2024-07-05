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
type BlockChain struct {

	// blocks là một mảng các con trỏ tới các Block. Mỗi phần tử trong mảng này là một con trỏ đến một đối tượng Block.
	Blocks []*Block
}

type Block struct {
	Hash     []byte //Đây là hash của khối hiện tại, được tính từ dữ liệu (Data) và hash của khối trước đó
	Data     []byte //Dữ liệu của khối hiện tại.
	PrevHash []byte //Hash của khối trước đó trong chuỗi khối.
	Nonce    int
}

// DeriveHash tính toán hash của khối bằng cách kết hợp dữ liệu của khối và hash của khối trước đó,
// sau đó sử dụng thuật toán SHA-256 để tính hash và gán nó cho thuộc tính Hash.
func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256((info))
	b.Hash = hash[:]
}

// CreateBlock tạo ra một khối mới với dữ liệu (data) và hash của khối trước đó (preHash).
// Sau đó, nó gọi phương thức DeriveHash để tính hash cho khối mới và trả về khối đó.
func CreateBlock(data string, preHash []byte) *Block {

	// CreateBlock tạo một con trỏ tới một Block mới.
	// Nó khởi tạo một đối tượng Block, tính toán hash của khối đó bằng cách gọi block.DeriveHash()
	// và trả về con trỏ tới đối tượng Block đó.

	block := &Block{[]byte{}, []byte(data), preHash, 0}
	// block.DeriveHash()

	//Video 2
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// AddBlock thêm một khối mới vào chuỗi khối.
// Nó lấy khối cuối cùng trong chuỗi khối (prevBlock),
// tạo ra một khối mới với dữ liệu mới và hash của khối trước đó, sau đó thêm khối mới vào chuỗi khối.
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1] //Lấy con trỏ đến khối cuối cùng trong chuỗi khối.
	new := CreateBlock(data, prevBlock.Hash)       //Tạo một khối mới và trả về con trỏ tới khối đó.
	chain.Blocks = append(chain.Blocks, new)       //Thêm con trỏ của khối mới vào mảng blocks.
}

// Genesis tạo ra khối đầu tiên trong chuỗi khối, gọi là khối gốc (genesis block).
// Khối này có dữ liệu là "Genesis" và không có hash của khối trước đó (vì nó là khối đầu tiên).
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// InitBlockChain khởi tạo một chuỗi khối mới với khối gốc.
func InitBlockChain() *BlockChain {

	//Genesis tạo và trả về con trỏ tới khối gốc (genesis block).
	return &BlockChain{[]*Block{Genesis()}}
}

func (b *Block) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)

	err := encoder.Encode(b)

	if err != nil {
		log.Panic(err)
	}

	Handle(err)

	return res.Bytes()
}

func Deserialize(data []byte) *Block {
	var block Block

	decoder := gob.NewDecoder((bytes.NewReader(data)))
	err := decoder.Decode(&block)

	if err != nil {
		log.Panic(err)
	}

	Handle(err)

	return &block
}

func Handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}

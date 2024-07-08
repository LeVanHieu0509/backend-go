package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
)

/*
	Lý thuyết:

	- PoW là một thuật toán đồng thuận được sử dụng trong blockchain để đảm bảo tính toàn vẹn và an toàn của dữ liệu
	- PoW yêu cầu các nút trong mạng thực hiện một lượng công việc tính toán để tạo ra một khối hợp lệ

*/

// Take the data from the block

// Create a counter (nonce) which starts at 0

// create a hash of the data plus the counter

// check the hash to see if it meets a set of requirements // Kiểm tra hàm băm xem nó đã đáp ứng đủ yêu cầu hay chưa, lặp để tạo hàm băm đến khi oke thì thôi

// Requirements:
// The first few bytes must contain 0s

const Difficulty = 16 // Độ khó đào xác đinh số bit 0 cần thiết ở đầu của hàm băm hợp lệ.

type ProofOfWork struct {
	Block  *Block
	Target *big.Int //Target là một giá trị lớn mà hàm băm phải nhỏ hơn hoặc bằng để được coi là hợp lệ.
}

// NewProof khởi tạo một đối tượng ProofOfWork mới với khối (Block) và mục tiêu băm (Target)
func NewProof(b *Block) *ProofOfWork {

	//được tính bằng cách dịch trái số 1 với 256 - Difficulty bit. Điều này tạo ra một giá trị mà hash phải nhỏ hơn hoặc bằng để được coi là hợp lệ.
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty))

	pow := &ProofOfWork{b, target}

	return pow
}

// Tạo dữ liệu đầu vào cho hàm băm bằng cách kết hợp hash của khối trước đó (PrevHash), dữ liệu của khối (Data), nonce, và độ khó đào (Difficulty).
func (pow *ProofOfWork) InitData(nonce int) []byte {
	data := bytes.Join([][]byte{
		pow.Block.PrevHash,
		// pow.Block.Data,
		//Part 4: Transaction
		pow.Block.HashTransactions(),
		ToHex(int64(nonce)),
		ToHex(int64(Difficulty)),
	},
		[]byte{})

	return data
}

// ToHex chuyển số nguyên thành mảng byte ở định dạng Big Endian. Đây là cách mã hóa số nguyên thành dạng byte để sử dụng trong tính toán hash.
func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)

	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

// Run là vòng lặp thực hiện công việc tính toán.
// Bắt đầu với nonce bằng 0, tạo dữ liệu đầu vào và tính toán hash.
func (pow *ProofOfWork) Run() (int, []byte) {
	var intHash big.Int
	var hash [32]byte

	nonce := 0

	for nonce < math.MaxInt64 {
		data := pow.InitData(nonce)
		fmt.Println("nonce", nonce, ":", data)
		hash = sha256.Sum256(data)

		//In hash ra màn hình để theo dõi tiến trình.
		fmt.Printf("nonce:%x : \r%x \n", nonce, hash)
		intHash.SetBytes((hash[:]))

		// Kiểm tra xem hash có nhỏ hơn mục tiêu (Target) không.
		// Nếu có, vòng lặp dừng lại và trả về nonce cùng với hash. Nếu không, tăng nonce và lặp lại.
		if intHash.Cmp(pow.Target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Println()

	return nonce, hash[:]
}

// Validate xác nhận khối bằng cách tạo lại dữ liệu đầu vào từ nonce của khối và tính toán hash.
// Kiểm tra xem hash có nhỏ hơn mục tiêu (Target) không. Nếu có, khối là hợp lệ.
func (pow *ProofOfWork) Validate() bool {
	var initHash big.Int
	data := pow.InitData(pow.Block.Nonce)

	hash := sha256.Sum256(data)
	initHash.SetBytes((hash[:]))

	return initHash.Cmp(pow.Target) == -1
}

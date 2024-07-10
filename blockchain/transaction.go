package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"log"
)

/*
Hàm NewTransaction tạo một giao dịch mới từ người gửi đến người nhận, kiểm tra số dư,
tạo các đầu vào và đầu ra cần thiết, và thiết lập ID cho giao dịch.
Nếu số dư không đủ, hàm sẽ gây hoảng với thông báo lỗi.
Nếu có số dư, phần dư sẽ được trả lại cho người gửi.
*/

type Transaction struct {
	ID      []byte     // Mã định danh của giao dịch.
	Inputs  []TxInput  //Các đầu vào của giao dịch.
	Outputs []TxOutput //Các đầu ra của giao dịch.
}

// Đại diện cho một đầu vào của giao dịch với các trường
type TxInput struct {
	ID  []byte //Mã định danh của giao dịch trước đó
	Out int    //Chỉ số của đầu ra trong giao dịch trước đó
	Sig string //Chữ ký của người gửi
}

type TxOutput struct {
	Value  int    //Giá trị của đầu ra.
	PubKey string //Khóa công khai của người nhận
}

// Hàm: SetId tính toán mã định danh của giao dịch
func (tx *Transaction) SetId() {
	var encoded bytes.Buffer //Sử dụng bytes.Buffer để lưu trữ dữ liệu mã hóa của giao dịch.
	var hash [32]byte

	// Sử dụng gob.NewEncoder để mã hóa giao dịch và lưu vào bộ đệm.
	encode := gob.NewEncoder(&encoded)
	err := encode.Encode(tx)
	Handle(err)

	// Sử dụng hàm băm SHA-256 để tính toán mã băm của dữ liệu mã hóa và gán cho tx.ID
	hash = sha256.Sum256(encoded.Bytes())
	tx.ID = hash[:]
}

// Là giao dịch đặc biệt để tạo ra tiền mới, không có đầu vào thực sự.
func CoinbaseTx(to, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Coins to %s", to)
	}

	// Đầu vào: TxInput với ID rỗng, Out là -1, và Sig là dữ liệu tùy chọn
	txin := TxInput{[]byte{}, -1, data}

	// TxOutput với giá trị 100 và khóa công khai là người nhận.
	txout := TxOutput{100, to}

	tx := Transaction{nil, []TxInput{txin}, []TxOutput{txout}}

	// thiết lập mã định danh cho giao dịch.
	tx.SetId()

	return &tx
}

// Hàm: IsCoinbase kiểm tra xem giao dịch có phải là giao dịch coinbase hay không
func (tx *Transaction) IsCoinbase() bool {

	// Giao dịch coinbase có một đầu vào duy nhất với ID rỗng và Out là -1.
	return len(tx.Inputs) == 1 && len(tx.Inputs[0].ID) == 0 && tx.Inputs[0].Out == -1
}

// Kiểm tra xem chữ ký của đầu vào có khớp với dữ liệu không.
func (in *TxInput) CanUnlock(data string) bool {
	return in.Sig == data
}

// Kiểm tra xem khóa công khai của đầu ra có khớp với dữ liệu không
func (out *TxOutput) CanBeUnlocked(data string) bool {
	return out.PubKey == data
}

// Hàm: NewTransaction tạo một giao dịch mới
// from: Địa chỉ của người gửi.
// to: Địa chỉ của người nhận.
// amount: Số tiền giao dịch.
// chain: Chuỗi khối (blockchain) mà giao dịch sẽ được thêm vào.

func NewTransaction(from, to string, amount int, chain *BlockChain) *Transaction {
	var inputs []TxInput
	var outputs []TxOutput

	// Gọi chain.FindSpendableOutputs để tìm các đầu ra có thể chi tiêu của người gửi

	// acc: Tổng số tiền có thể chi tiêu của người gửi.
	// validOutput: Một bản đồ (map) chứa các đầu ra có thể chi tiêu, với mã giao dịch là khóa và các chỉ số của các đầu ra là giá trị.
	acc, validOutput := chain.FindSpendableOutputs(from, amount)

	// Nếu số dư không đủ, gây hoảng với thông báo lỗi
	if acc < amount {
		log.Panic("Error: not enough funds")
	}

	for txid, outs := range validOutput {
		// Chuyển đổi mã giao dịch từ chuỗi hexa thành mảng byte (ixID).
		ixID, err := hex.DecodeString(txid)
		Handle(err)

		for _, out := range outs {
			input := TxInput{ixID, out, from}
			inputs = append(inputs, input)
		}
	}

	// Tạo đầu ra mới với giá trị giao dịch (amount) và địa chỉ người nhận (to).
	outputs = append(outputs, TxOutput{amount, to})

	// Nếu tổng số tiền có thể chi tiêu (acc) lớn hơn số tiền yêu cầu (amount), tạo thêm một đầu ra để trả lại phần tiền dư cho người gửi.
	if acc > amount {
		outputs = append(outputs, TxOutput{acc - amount, from})
	}

	tx := Transaction{nil, inputs, outputs}

	// Thiết lập ID cho giao dịch bằng cách gọi hàm SetId.
	tx.SetId()

	// Trả về con trỏ đến giao dịch mới (&tx).
	return &tx
}

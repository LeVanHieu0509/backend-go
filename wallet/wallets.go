package wallet

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"
)

const walletFile = "../tmp/wallets.data"

/*
Wallets là một cấu trúc để quản lý nhiều ví,
với mỗi ví được lưu trữ trong một bản đồ (map) có khóa là địa chỉ ví và giá trị là con trỏ tới cấu trúc Wallet.

Một địa chỉ ví là duy nhất: Mỗi địa chỉ ví được tạo ra bằng cách sử dụng khóa công khai của ví.
Địa chỉ ví được sử dụng để nhận tiền từ người khác trong hệ thống blockchain.

Một ví duy nhất có một địa chỉ duy nhất: Mỗi ví có một cặp khóa riêng tư và khóa công khai.
Địa chỉ ví được tạo ra từ khóa công khai, do đó mỗi ví có một địa chỉ duy nhất.

Hệ thống quản lý nhiều ví: Trong cấu trúc Wallets, chúng ta có một bản đồ (map) chứa nhiều ví khác nhau.
Mỗi ví có một địa chỉ duy nhất và bản đồ này cho phép chúng ta quản lý nhiều ví trong cùng một hệ thống.
*/

type Wallets struct {
	Wallets map[string]*Wallet
}

/*
Hàm này khởi tạo một đối tượng Wallets mới,
gọi LoadFile để tải dữ liệu từ tệp nếu tồn tại và trả về con trỏ tới Wallets cùng với lỗi nếu có.
*/

func CreateWallets() (*Wallets, error) {
	wallets := Wallets{}

	wallets.Wallets = make(map[string]*Wallet)

	err := wallets.LoadFile()

	return &wallets, err
}

/*
Hàm này nhận vào một địa chỉ ví và trả về ví tương ứng.
Nó tìm kiếm địa chỉ trong bản đồ và trả về giá trị của nó.
*/

func (ws Wallets) GetWallet(address string) Wallet {
	return *ws.Wallets[address]
}

/*
Hàm này trả về danh sách tất cả các địa chỉ ví có trong Wallets.
*/
func (ws Wallets) GetAllAddresses() []string {
	var addresses []string

	for address := range ws.Wallets {
		addresses = append(addresses, address)
	}

	return addresses
}

/*
Hàm này tạo một ví mới, thêm nó vào Wallets, và trả về địa chỉ của ví mới.
*/

func (ws Wallets) AddWallet() string {
	wallet := MakeWallet()

	address := fmt.Sprintf("%s", wallet.Address())

	ws.Wallets[address] = wallet

	return address
}

/*
Hàm này kiểm tra sự tồn tại của tệp walletFile.
Nếu tệp không tồn tại, nó trả về lỗi.
Nếu tệp tồn tại, nó đọc nội dung tệp, giải mã nó thành một bản đồ serializableWallets,
và chuyển đổi từng SerializableWallet thành Wallet trước khi lưu trữ vào Wallets.
*/

func (ws *Wallets) LoadFile() error {
	if _, err := os.Stat(walletFile); os.IsNotExist(err) {
		return err
	}

	var serializableWallets map[string]SerializableWallet

	fileContent, err := ioutil.ReadFile(walletFile)
	if err != nil {
		return err
	}

	// Đăng ký loại elliptic.P256() với gob
	gob.Register(elliptic.P256())

	decoder := gob.NewDecoder(bytes.NewReader(fileContent))
	err = decoder.Decode(&serializableWallets)
	if err != nil {
		return err
	}

	ws.Wallets = make(map[string]*Wallet)
	for address, serializableWallet := range serializableWallets {
		ws.Wallets[address] = FromSerializable(serializableWallet)
	}

	return nil
}

/*
Hàm này mã hóa Wallets thành dạng có thể lưu trữ và ghi vào tệp walletFile.
Nếu thư mục chứa tệp không tồn tại, nó sẽ tạo thư mục trước khi ghi tệp.
*/

func (ws *Wallets) SaveFile() {
	var content bytes.Buffer

	serializableWallets := make(map[string]SerializableWallet)
	for address, wallet := range ws.Wallets {
		serializableWallets[address] = wallet.ToSerializable()
	}

	// Đăng ký loại elliptic.P256() với gob
	gob.Register(elliptic.P256())

	encoder := gob.NewEncoder(&content)
	err := encoder.Encode(serializableWallets)
	if err != nil {
		log.Panic(err)
	}

	// Tạo thư mục nếu chưa tồn tại
	if err := os.MkdirAll("../tmp", os.ModePerm); err != nil {
		log.Panic(err)
	}

	err = ioutil.WriteFile(walletFile, content.Bytes(), 0644)
	if err != nil {
		log.Panic(err)
	}
}

type SerializableWallet struct {
	PrivateKey []byte
	PublicKey  []byte
}

/*
ToSerializable chuyển đổi một Wallet thành dạng có thể lưu trữ (SerializableWallet)
bằng cách lấy các byte của khóa riêng và khóa công khai.
*/

func (w *Wallet) ToSerializable() SerializableWallet {
	privateKeyBytes := w.PrivateKey.D.Bytes()
	curveParams := elliptic.Marshal(w.PrivateKey.Curve, w.PrivateKey.X, w.PrivateKey.Y)
	return SerializableWallet{
		PrivateKey: privateKeyBytes,
		PublicKey:  curveParams,
	}
}

/*
FromSerializable chuyển đổi SerializableWallet thành Wallet bằng cách tái tạo khóa riêng và khóa công khai từ các byte đã lưu trữ.
*/
func FromSerializable(s SerializableWallet) *Wallet {
	privateKey := new(ecdsa.PrivateKey)
	privateKey.PublicKey.Curve = elliptic.P256()
	privateKey.D = new(big.Int).SetBytes(s.PrivateKey)
	privateKey.PublicKey.X, privateKey.PublicKey.Y = elliptic.Unmarshal(elliptic.P256(), s.PublicKey)

	return &Wallet{
		PrivateKey: *privateKey,
		PublicKey:  s.PublicKey,
	}
}

package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"log"

	"golang.org/x/crypto/ripemd160"
)

var (
	checksumLength = 4
	version        = byte(0x00)
)

type Wallet struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey  []byte
}

// function internal
/*
Hàm Address của struct Wallet tạo ra địa chỉ ví từ khóa công khai.
Địa chỉ này được mã hóa bằng Base58 để dễ dàng lưu trữ và truyền tả
*/
func (w Wallet) Address() []byte {
	// Tạo giá trị băm từ khóa công khai của ví.
	// Hàm PublicKeyHash thường sử dụng SHA-256 và RIPEMD-160 để tạo băm.
	pubHash := PublicKeyHash(w.PublicKey)

	// Thêm phiên bản vào giá trị băm
	// Thêm một byte phiên bản vào đầu giá trị băm.
	// Byte phiên bản (version) giúp xác định loại địa chỉ (ví dụ: mainnet hoặc testnet)
	versionedHash := append([]byte{version}, pubHash...)

	// Tạo giá trị kiểm tra (checksum) từ giá trị băm có phiên bản bằng cách sử dụng hàm SHA-256 hai lần
	checksum := CheckSum(versionedHash)

	// Kết hợp giá trị băm có phiên bản với giá trị kiểm tra để tạo ra giá trị băm đầy đủ.
	fullHash := append(versionedHash, checksum...)

	// Mã hóa giá trị băm đầy đủ bằng Base58 để tạo địa chỉ ví.
	// Base58 là một dạng mã hóa không bao gồm các ký tự dễ nhầm lẫn như '0', 'O', 'I', và 'l'.
	address := Base58Encode(fullHash)

	fmt.Printf("pub key: %x\n", w.PublicKey)
	fmt.Printf("pub hash: %x\n", pubHash)
	fmt.Printf("address:: %x\n", address)

	return address
}

/*
Hàm NewKeyPair tạo một cặp khóa sử dụng thuật toán ECDSA (Elliptic Curve Digital Signature Algorithm).
Nó trả về khóa riêng (private key) và khóa công khai (public key) dưới dạng mảng byte.
*/
func NewKeyPair() (ecdsa.PrivateKey, []byte) {
	// Sử dụng đường cong elliptic P-256
	curve := elliptic.P256()

	// Tạo khóa riêng sử dụng đường cong đã chọn và nguồn ngẫu nhiên từ rand.Reader
	private, err := ecdsa.GenerateKey(curve, rand.Reader)

	if err != nil {
		log.Panic(err)
	}

	// Khóa công khai được lấy từ khóa riêng, và nó bao gồm hai thành phần X và Y
	// Nối các thành phần X và Y lại với nhau để tạo thành khóa công khai đầy đủ
	pub := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)

	// Trả về khóa riêng và khóa công khai.
	return *private, pub
}

// Hàm MakeWallet tạo một ví mới sử dụng cặp khóa được tạo bởi hàm NewKeyPair. Nó trả về con trỏ đến đối tượng Wallet.
func MakeWallet() *Wallet {
	// Gọi hàm NewKeyPair để tạo cặp khóa mới, bao gồm khóa riêng và khóa công khai.
	private, public := NewKeyPair()

	// Tạo một đối tượng Wallet mới với các khóa vừa được tạo.
	wallet := Wallet{private, public}

	// Trả về con trỏ đến đối tượng Wallet mới.
	return &wallet
}

/*
Hàm PublicKeyHash thực hiện băm khóa công khai (public key) sử dụng SHA-256 và
RIPEMD-160 để tạo ra một khóa băm công khai (public key hash).
Đây là một phần quan trọng của quá trình tạo địa chỉ ví trong các hệ thống blockchain như Bitcoin

Kết hợp SHA-256 và RIPEMD-160 tăng cường bảo mật.
SHA-256 cung cấp độ an toàn cao với băm 256-bit
trong khi RIPEMD-160 tạo ra băm 160-bit ngắn hơn nhưng vẫn rất an toàn
phù hợp để sử dụng trong địa chỉ ví.
*/
func PublicKeyHash(pubKey []byte) []byte {
	// Sử dụng thuật toán SHA-256 để băm dữ liệu đầu vào pubKey (khóa công khai).
	pubHash := sha256.Sum256(pubKey)

	// Tạo một đối tượng hasher mới sử dụng thuật toán RIPEMD-160.
	hasher := ripemd160.New()

	// Ghi dữ liệu đã được băm bằng SHA-256 vào hasher RIPEMD-160.
	_, err := hasher.Write(pubHash[:])

	if err != nil {
		log.Panic(err)
	}

	// Thực hiện băm trên dữ liệu đã ghi vào hasher và trả về kết quả cuối cùng.
	publicRipMD := hasher.Sum(nil)

	return publicRipMD

}

func CheckSum(payload []byte) []byte {
	// Thực hiện băm dữ liệu đầu vào (payload) bằng thuật toán SHA-256.
	firstHash := sha256.Sum256(payload)

	// Thực hiện băm trên kết quả của băm lần đầu (firstHash) bằng SHA-256 lần nữa.
	secondHash := sha256.Sum256(firstHash[:])

	return secondHash[:checksumLength]

}

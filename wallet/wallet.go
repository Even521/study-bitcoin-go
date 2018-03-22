package wallet

import (
	"bytes"
	"crypto/sha256"
	"crypto/ecdsa"
	"crypto/elliptic"
	"log"
	"github.com/study-bitcoin-go/utils"
	"crypto/rand"
	"github.com/study-bitcoin-go/utils/ripemd160"
)

const version = byte(0x00)  //16进制0 版本号
const walletFile = "db/wallet.dat"
const addressChecksumLen = 4 //地址检查长度4

// 钱包
type Wallet struct {
	/**
	 PrivateKey: ECDSA基于椭圆曲线
	 使用曲线生成私钥，并从私钥生成公钥
	 */
	PrivateKey ecdsa.PrivateKey //私钥
	PublicKey  []byte //公钥
}

// 创建一个新钱包
func NewWallet() *Wallet {
	//公钥私钥生成
	private, public := newKeyPair()
	wallet := Wallet{private, public}
	return &wallet
}

// 得到一个钱包地址
func (w Wallet) GetAddress() []byte {

	pubKeyHash := HashPubKey(w.PublicKey)
	//将版本号+pubKeyHash得到一个散列
	versionedPayload := append([]byte{version}, pubKeyHash...)
	//校验前4个字节的散列
	checksum := checksum(versionedPayload)
	//将校验和附加到version+PubKeyHash组合。
	fullPayload := append(versionedPayload, checksum...)
	//BASE58得到一个钱包地址
	address := utils.Base58Encode(fullPayload)
	return address
}

// 使用RIPEMD160(SHA256(PubKey))哈希算法得到hsahpubkey
func HashPubKey(pubKey []byte) []byte {

	publicSHA256 := sha256.Sum256(pubKey)
	RIPEMD160Hasher := ripemd160.New()
	_, err := RIPEMD160Hasher.Write(publicSHA256[:])
	if err != nil {
		log.Panic(err)
	}
	publicRIPEMD160 := RIPEMD160Hasher.Sum(nil)

	return publicRIPEMD160
}

// 校验地址
func ValidateAddress(address string) bool {
	pubKeyHash := utils.Base58Decode([]byte(address))
	actualChecksum := pubKeyHash[len(pubKeyHash)-addressChecksumLen:]
	version := pubKeyHash[0]
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-addressChecksumLen]
	targetChecksum := checksum(append([]byte{version}, pubKeyHash...))

	return bytes.Compare(actualChecksum, targetChecksum) == 0
}

//SHA256(SHA256(payload))算法返回前4个字节
func checksum(payload []byte) []byte {
	firstSHA := sha256.Sum256(payload)
	secondSHA := sha256.Sum256(firstSHA[:])
	return secondSHA[:addressChecksumLen]
}

//椭圆算法返回私钥与公钥
func newKeyPair() (ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256()
    //获取私钥
	private, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panic(err)
	}
	//在基于椭圆曲线的算法中，公钥是曲线上的点。因此，公钥是X，Y坐标的组合。在比特币中，这些坐标被连接起来形成一个公钥。
	pubKey := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)

	return *private, pubKey
}
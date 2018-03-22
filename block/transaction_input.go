package block

import (
	"bytes"
	"github.com/study-bitcoin-go/wallet"
)

//输入事物
type TXInput struct {
	Txid      []byte //事物hash
	Vout      int    //输出值
	Signature []byte //签名
	PubKey    []byte //公钥
}

//检查输入是否使用特定的键来解锁输出
func (in *TXInput) UsesKey(pubKeyHash []byte) bool {
	lockingHash := wallet.HashPubKey(in.PubKey)

	return bytes.Compare(lockingHash, pubKeyHash) == 0
}

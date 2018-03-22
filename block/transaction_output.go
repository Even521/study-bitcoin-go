package block

import (
	"github.com/study-bitcoin-go/utils"
	"bytes"
)

//一个事物输出
type TXOutput struct {
	Value int       //值
	PubKeyHash []byte //解锁脚本key
}

// Lock只需锁定输出
func (out *TXOutput) Lock(address []byte) {
	pubKeyHash := utils.Base58Decode(address)
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
	out.PubKeyHash = pubKeyHash
}

// 检查提供的公钥散列是否用于锁定输出
func (out *TXOutput) IsLockedWithKey(pubKeyHash []byte) bool {
	return bytes.Compare(out.PubKeyHash, pubKeyHash) == 0
}

// 新的交易输出
func NewTXOutput(value int, address string) *TXOutput {
	txo := &TXOutput{value, nil}
	txo.Lock([]byte(address))

	return txo
}
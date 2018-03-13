package block


import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

//区块结构
type Block struct {
	Hash          []byte //hase值
	Data          []byte //交易数据
	PrevBlockHash []byte //存储前一个区块的Hase值
	Timestamp     int64  //生成区块的时间
}

//生成一个新的区块方法
func NewBlock(data string, prevBlockHash []byte) *Block{
	block := &Block{Timestamp:time.Now().Unix(), Data:[]byte(data), PrevBlockHash:prevBlockHash, Hash:[]byte{}}
	block.SetHash()
	return block
}

//生成区块hase值
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

//创世块方法
func  NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

package block


import (
	"time"
)

//区块结构
type Block struct {
	Hash          []byte //hase值
	Data          []byte //交易数据
	PrevBlockHash []byte //存储前一个区块的Hase值
	Timestamp     int64  //生成区块的时间
	Nonce         int    //工作量证明算法的计数器
}

//生成一个新的区块方法
func NewBlock(data string, prevBlockHash []byte) *Block{
	//GO语言给Block赋值{}里面属性顺序可以打乱，但必须制定元素 如{Timestamp:time.Now().Unix()...}
	block := &Block{Timestamp:time.Now().Unix(), Data:[]byte(data), PrevBlockHash:prevBlockHash, Hash:[]byte{},Nonce:0}

	//工作证明
	pow :=NewProofOfWork(block)
	//工作量证明返回计数器和hash
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}


//区块校验
func (i *Block) Validate() bool {
    return NewProofOfWork(i).Validate()
}

//创世块方法
func  NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}


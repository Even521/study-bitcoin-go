package block



// 区块链
type Blockchain struct {
	Blocks []*Block
}

// 保存区块数据
func (bc *Blockchain) AddBlock(data string) {
	//获取上一个区块
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	//创建一个新的区块
	newBlock := NewBlock(data, prevBlock.Hash)
	//新的区块添加到数组中
	bc.Blocks = append(bc.Blocks, newBlock)
}

// 创建创世块
func NewBlockchain() *Blockchain {
	//go语言&表示获取存储的内存地址
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}


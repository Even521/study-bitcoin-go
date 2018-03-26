package block

import (
	"github.com/boltdb/bolt"
	"log"
)

// 区块链迭代器用于迭代区块
type BlockchainIterator struct {
	currentHash []byte //当前的hash
	db          *bolt.DB //数据库
}

// Next returns next block starting from the tip
func (i *BlockchainIterator) Next() *Block {
	var block *Block

	err := i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodedBlock := b.Get(i.currentHash)
		block = DeserializeBlock(encodedBlock)

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	i.currentHash = block.PrevBlockHash

	return block
}
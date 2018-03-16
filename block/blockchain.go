package block

import (
	"github.com/boltdb/bolt"
	"log"
	"fmt"
)

const dbFile  ="blockchian.db" //定义数据文件名
const blocksBucket="blocks_"//区块桶

// 区块链保持一个序列化
type Blockchain struct {
	tip []byte   //最顶层hash
	db  *bolt.DB //BoltDB数据库
}
// 区块链迭代器用于迭代区块
type BlockchainIterator struct {
	currentHash []byte //当前的hash
	db          *bolt.DB //数据库
}
//添加区块
func (bc *Blockchain) AddBlock(data string)  {
	var lastHash []byte //最后一个区块hash
	//查询数据库中最后一块的hash
	err :=bc.db.View(func(tx *bolt.Tx) error {
		b :=tx.Bucket([]byte(blocksBucket))
		lastHash=b.Get([]byte("1"))//最新的一块hash的key我们知道为"l"
		return nil
	})

	if err!=nil{
		log.Panic(err)
	}
	//利最后的一块hash，挖掘一块新的区块出来
	newBlock :=NewBlock(data,lastHash)
    //在挖掘新块之后，我们将其序列化表示保存到数据块中并更新"l"，该密钥现在存储新块的哈希。
	err=bc.db.Update(func(tx *bolt.Tx) error {
		b :=tx.Bucket([]byte(blocksBucket))
		err :=b.Put(newBlock.Hash,newBlock.Serialize())
		if err!=nil{
			log.Panic(err)
		}
		err = b.Put([]byte("l"), newBlock.Hash)
		if err != nil {
			log.Panic(err)
		}
		bc.tip = newBlock.Hash
		return nil
	})
}
//迭代器
func (bc *Blockchain) Iterator() *BlockchainIterator {
	bci := &BlockchainIterator{bc.tip, bc.db}

	return bci
}

// 迭代下一区块
func (i *BlockchainIterator) Next() *Block {
	var block *Block

	err := i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		//查询区块
		encodedBlock := b.Get(i.currentHash)
		block = DeserializeBlock(encodedBlock)

		return nil
	})

	if err != nil {
		log.Panic(err)
	}
    //将前一个区块
	i.currentHash = block.PrevBlockHash

	return block
}
//关闭方法
func Close(bc *Blockchain) error{
   return bc.db.Close()
}


// 创建一个新的区块链和创世块
func NewBlockchain() *Blockchain {
	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))

		if b == nil {
			fmt.Println("No existing blockchain found. Creating a new one...")
			genesis := NewGenesisBlock()

			b, err := tx.CreateBucket([]byte(blocksBucket))
			if err != nil {
				log.Panic(err)
			}

			err = b.Put(genesis.Hash, genesis.Serialize())
			if err != nil {
				log.Panic(err)
			}

			err = b.Put([]byte("l"), genesis.Hash)
			if err != nil {
				log.Panic(err)
			}
			tip = genesis.Hash
		} else {
			tip = b.Get([]byte("l"))
		}

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	bc := Blockchain{tip, db}

	return &bc
}


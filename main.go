package main

import (
"fmt"


	"strconv"
	"github.com/study-bitcion-go/block"
	"crypto"
)



func main() {

    //创世块
	bc :=block.NewBlockchain()
	//添加第2块
	bc.AddBlock("Send 1 BTC to even")
	//添加第3块
	bc.AddBlock("Send 2 more BTC to even")
    //迭代数组里面的数据
	for _, block := range bc.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Nonce: %d\n",block.Nonce)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(block.Validate()))
		fmt.Println()

	}


}

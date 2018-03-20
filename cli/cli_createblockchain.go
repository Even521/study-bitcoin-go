package cli

import (
	"log"
	"fmt"

	"github.com/study-bitcion-go/wallet"
	"github.com/study-bitcion-go/block"
)
//创建一个区块链
func (cli *CLI) createBlockchain(address string) {
	//校验钱包地址
	if !wallet.ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	//创建一个区块链
	bc := block.CreateBlockchain(address)
	//关闭数据库
	block.Close(bc)
	fmt.Println("Done!")
}
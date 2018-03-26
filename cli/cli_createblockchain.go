package cli

import (
	"log"
	"fmt"

	"github.com/study-bitcoin-go/wallet"
	"github.com/study-bitcoin-go/block"
)
//创建一个区块链
func (cli *CLI) createBlockchain(address string) {
	if !wallet.ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	bc := block.CreateBlockchain(address)

	defer block.Close(bc)

	UTXOSet := block.UTXOSet{bc}
	UTXOSet.Reindex()

	fmt.Println("Done!")
}
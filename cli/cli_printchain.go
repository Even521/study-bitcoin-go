package cli

import (
	"fmt"
	"strconv"
	"github.com/study-bitcoin-go/block"
)

func (cli *CLI) printChain() {
	bc := block.NewBlockchain()
	defer block.Close(bc)

	bci := bc.Iterator()

	for {
		block := bci.Next()

		fmt.Printf("============ Block %x ============\n", block.Hash)
		fmt.Printf("Prev. block: %x\n", block.PrevBlockHash)

		fmt.Printf("PoW: %s\n\n", strconv.FormatBool(block.Validate()))
		for _, tx := range block.Transactions {
			fmt.Println(tx)
		}
		fmt.Printf("\n\n")

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}
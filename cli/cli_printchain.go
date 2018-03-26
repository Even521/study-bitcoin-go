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
		b := bci.Next()

		fmt.Printf("============ Block %x ============\n", b.Hash)
		fmt.Printf("Prev. block: %x\n", b.PrevBlockHash)
		pow := block.NewProofOfWork(b)
		fmt.Printf("PoW: %s\n\n", strconv.FormatBool(pow.Validate()))
		for _, tx := range b.Transactions {
			fmt.Println(tx)
		}
		fmt.Printf("\n\n")

		if len(b.PrevBlockHash) == 0 {
			break
		}
	}
}
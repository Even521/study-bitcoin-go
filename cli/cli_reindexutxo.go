package cli

import (
	"fmt"
	"github.com/study-bitcoin-go/block"

)

func (cli *CLI) reindexUTXO() {
	bc := block.NewBlockchain()
	UTXOSet := block.UTXOSet{bc}
	UTXOSet.Reindex()

	count := UTXOSet.CountTransactions()
	fmt.Printf("Done! There are %d transactions in the UTXO set.\n", count)
}
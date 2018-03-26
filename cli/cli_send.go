package cli

import (
	"log"
	"fmt"
	"github.com/study-bitcoin-go/wallet"
	"github.com/study-bitcoin-go/block"
)

func (cli *CLI) send(from, to string, amount int) {
	if !wallet.ValidateAddress(from) {
		log.Panic("ERROR: Sender address is not valid")
	}
	if !wallet.ValidateAddress(to) {
		log.Panic("ERROR: Recipient address is not valid")
	}
    bc :=block.NewBlockchain()
    if bc!=nil {
		defer block.Close(bc)
	}
	//bc := block.NewBlockchain()
	UTXOSet := block.UTXOSet{bc}



	tx := block.NewUTXOTransaction(from, to, amount, &UTXOSet)
	cbTx := block.NewCoinbaseTX(from, "")
	txs := []*block.Transaction{cbTx, tx}

	newBlock := bc.MineBlock(txs)
	UTXOSet.Update(newBlock)


	fmt.Println("Success!")
}
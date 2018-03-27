package cli

import (
	"log"
	"fmt"
	"github.com/study-bitcoin-go/wallet"
	"github.com/study-bitcoin-go/block"
)
//交易
func (cli *CLI) send(from, to string, amount int) {
    //检验发送钱包地址
	if !wallet.ValidateAddress(from) {
		log.Panic("ERROR: Sender address is not valid")
	}
	//校验接收钱包地址
	if !wallet.ValidateAddress(to) {
		log.Panic("ERROR: Recipient address is not valid")
	}

    bc :=block.NewBlockchain()
    defer block.Close(bc)

	//bc := block.NewBlockchain()
	UTXOSet := block.UTXOSet{bc}

	tx := block.NewUTXOTransaction(from, to, amount, &UTXOSet)
	cbTx := block.NewCoinbaseTX(from, "")
	txs := []*block.Transaction{cbTx, tx}

	newBlock := bc.MineBlock(txs)
	UTXOSet.Update(newBlock)


	fmt.Println("Success!")
}
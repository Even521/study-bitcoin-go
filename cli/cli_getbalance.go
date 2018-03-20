package cli

import (
	"log"
	"fmt"
	"github.com/study-bitcoin-go/wallet"
	"github.com/study-bitcoin-go/block"
	"github.com/study-bitcoin-go/utils"
)

func (cli *CLI) getBalance(address string) {
	if !wallet.ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	bc := block.NewBlockchain(address)
	defer block.Close(bc)

	balance := 0
	pubKeyHash := utils.Base58Decode([]byte(address))
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
	UTXOs := bc.FindUTXO(pubKeyHash)

	for _, out := range UTXOs {
		balance += out.Value
	}

	fmt.Printf("Balance of '%s': %d\n", address, balance)
}
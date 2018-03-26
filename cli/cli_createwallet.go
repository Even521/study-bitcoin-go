package cli

import (
	"github.com/study-bitcoin-go/wallet"
	"fmt"
)

//创建钱包
func (cli *CLI) createWallet() {
	wallets, _ := wallet.NewWallets()
	address := wallets.CreateWallet()
	wallets.SaveToFile()
	fmt.Printf("Your new address: %s\n", address)
}
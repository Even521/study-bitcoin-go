package main

import (
	"github.com/study-bitcion-go/block"
	"github.com/study-bitcion-go/cli"
)



func main() {
	bc := block.NewBlockchain()
    block.Close(bc)
    cli.Start(bc)
}

package cli

import (
	"github.com/study-bitcion-go/block"
	"fmt"
	"os"
	"strconv"
	"flag"
	"log"
)

type CLI struct {
	bc *block.Blockchain //区块链
}
//启动接口函数
func Start(bc *block.Blockchain)interface{}  {
	cl := CLI{bc}
	cl.run()//执行命令方法
	return  nil
}


//打印用法
func (cli *CLI) printUsage()  {
	fmt.Println("Usage:")
	fmt.Println("  addblock -data BLOCK_DATA - add a block to the blockchain")
	fmt.Println("  printchain - print all the blocks of the blockchain")
}
//校验参数
func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}
//添加区块数据
func (cli *CLI) addBlock(data string) {
	cli.bc.AddBlock(data)
	fmt.Println("Success!")
}
//打印区块链上所有区块数据
func (cli *CLI) printChain() {
	bci := cli.bc.Iterator()

	for {
		block := bci.Next()
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		fmt.Printf("PoW: %s\n", strconv.FormatBool(block.Validate()))
		fmt.Println()
		//创世块是没有前一个区块的，所以PrevBlockHash的值是没有的
		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}

// 执行命令方法
func (cli *CLI) run() {
	cli.validateArgs()

	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)

	addBlockData := addBlockCmd.String("data", "", "Block data")

	switch os.Args[1] {
	case "addblock":
		err := addBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		cli.printUsage()
		os.Exit(1)
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			os.Exit(1)
		}
		cli.addBlock(*addBlockData)
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}
}
package main

import (
	"fmt"
	"os"
	"simplechain/blockchain"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Printf("用法：addblock /printchain")
		return
	}

	var diffculty int = 3

	genesis := blockchain.GenerateFirstBlock()
	chain := blockchain.InitBlockchain(genesis)

	cmd := os.Args[1]
	switch cmd {
	case "addblock":
		w := blockchain.NewWallet()
		txs := []blockchain.Transaction{
			{From: "system", To: w.Address, Amount: 100},
		}
		chain.AddBlock(txs, diffculty)
		fmt.Printf("区块已添加")
	case "printchain":
		chain.Print()
	default:
		fmt.Println("未知命令")
	}

}

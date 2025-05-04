package main

import (
	"fmt"
	"simplechain/blockchain"
)

func main() {
	var chain []blockchain.Block

	//定义难度系数
	var diffculty int = 6

	//创建两个钱包
	w1 := blockchain.NewWallet()
	w2 := blockchain.NewWallet()

	fmt.Println("钱包1地址:", w1.Address)
	fmt.Println("钱包2地址:", w2.Address)

	firstblock := blockchain.GenerateFirstBlock()
	chain = append(chain, firstblock)

	txs := []blockchain.Transaction{
		{From: "SYSTEM", To: w1.Address, Amount: 100},
		{From: "SYSTEM", To: w2.Address, Amount: 20},
	}

	newBlock := blockchain.MineBlock(firstblock, txs, diffculty)

	chain = append(chain, newBlock)

	for _, tx := range chain {
		fmt.Printf("高度：%d\n哈希：%s\n", tx.Index, tx.Hash)
		for _, data := range tx.Data {
			fmt.Printf("from:%s to:%s amount:%d\n", data.From, data.To, data.Amount)
		}
		fmt.Println("---")
	}
}

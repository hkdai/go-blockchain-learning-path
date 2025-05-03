package main

import (
	"fmt"
	"simplechain/blockchain"
)

func main() {
	var chain []blockchain.Block
	firstblock := blockchain.GenerateFirstBlock()
	chain = append(chain, firstblock)

	for i := 1; i <= 5; i++ {
		newBlock := blockchain.GenerateNextBlock(chain[len(chain)-1], fmt.Sprintf("区块数据 %d", i))
		chain = append(chain, newBlock)
	}

	for _, block := range chain {
		fmt.Printf("高度: %d\nHash: %s\n上一个: %s\n数据: %s\n---\n",
			block.Index, block.Hash, block.PrevHash, block.Data)
	}
}

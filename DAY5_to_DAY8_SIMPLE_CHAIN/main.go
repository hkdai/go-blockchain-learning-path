package main

import (
	"fmt"
	"simplechain/blockchain"
)

func main() {
	var chain []blockchain.Block
	firstblock := blockchain.GenerateFirstBlock()
	chain = append(chain, firstblock)

	//老版本的直接生成区块
	// for i := 1; i <= 5; i++ {
	// 	newBlock := blockchain.GenerateNextBlock(chain[len(chain)-1], fmt.Sprintf("区块数据 %d", i))
	// 	chain = append(chain, newBlock)
	// }

	//定义难度系数
	var diffculty int = 6
	//新版本挖矿生成区块
	for i := 1; i <= 5; i++ {
		fmt.Printf("正在挖第%d个区块..\n", i)
		newBlock := blockchain.MineBlock(firstblock, fmt.Sprintf("区块数据 %d", i), diffculty)
		chain = append(chain, newBlock)
		fmt.Printf("区块挖出！Nonce ：%d Hash: %s\n\n", newBlock.Nonce, newBlock.Hash)
	}

	for _, block := range chain {
		fmt.Printf("高度: %d\nHash: %s\n上一个: %s\n数据: %s\n---\n",
			block.Index, block.Hash, block.PrevHash, block.Data)
	}
}

package blockchain

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type Block struct {
	Index     int64
	Timestamp int64
	PrevHash  string
	Data      string
	Hash      string
}

// 根据区块信息生成对应的hash
func GenerateHash(block Block) string {
	record := fmt.Sprintf("%d%d%s%s", block.Index, block.Timestamp, block.PrevHash, block.Data)
	h := sha256.New()
	h.Write([]byte(record))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// 生成创世区块
func GenerateFirstBlock() Block {
	block := Block{
		Index:     0,
		Timestamp: time.Now().Unix(),
		PrevHash:  "",
		Data:      "FirstBlock",
	}
	block.Hash = GenerateHash(block)
	return block
}

// 生成下一个区块对象
func GenerateNextBlock(prevBlock Block, data string) Block {
	newBlock := Block{
		Index:     prevBlock.Index + 1,
		Timestamp: time.Now().Unix(),
		PrevHash:  prevBlock.Hash,
		Data:      data,
	}
	newBlock.Hash = GenerateHash(newBlock)
	return newBlock
}

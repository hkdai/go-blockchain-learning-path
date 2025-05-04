package blockchain

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"

	"go.etcd.io/bbolt"
)

const dbFile = "chain.db"
const bucketName = "blocks"
const lastHashKey = "l"

type Blockchain struct {
	Tip []byte //当前链尾（最新区块）hash
	DB  *bbolt.DB
}

// 实现初始化链的方法（创建数据库并写入创世区块）
func InitBlockchain(genesis Block) *Blockchain {
	var lastHash []byte

	db, err := bbolt.Open(dbFile, 0600, nil) //权限是 0600：只允许当前用户访问
	if err != nil {
		log.Fatal(err)
	}
	//用事务创建 bucket 和写入创世块
	db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		if b == nil {
			b, _ = tx.CreateBucket([]byte(bucketName))

			genBytes := serialize(genesis)
			b.Put([]byte(genesis.Hash), genBytes)
			b.Put([]byte(lastHashKey), []byte(genesis.Hash))
			lastHash = []byte(genesis.Hash)
		} else {
			lastHash = b.Get([]byte(lastHashKey))
		}
		return nil
	})
	return &Blockchain{
		Tip: lastHash,
		DB:  db,
	}
}

// 给Blockchain绑定添加区块的方法
func (bc *Blockchain) AddBlock(data []Transaction,difficulty int){
	var lastBlock Block
	bc.DB.View(func(tx *bbolt.Tx) error{
		b := tx.Bucket([]byte(bucketName))
		blockData := b.Get(bc.Tip)
		lastBlock = deserialize(blockData)
		return nil
	})

	newBlock := MineBlock(lastBlock,data,difficulty)
	bc.DB.Update(func(tx *bbolt.Tx) error {
		b:= tx.Bucket([]byte(bucketName))
		b.Put([]byte(newBlock.Hash),serialize(newBlock))
		b.Put([]byte(lastHashKey),[]byte(newBlock.Hash))
		bc.Tip = []byte(newBlock.Hash)
		return nil
	})
}

// 打印整个链（遍历）

func (bc *Blockchain) Print(){
	cur := bc.Tip
	for{
		var block Block
		bc.DB.View(func(tx *bbolt.Tx) error{
			b := tx.Bucket([]byte(bucketName))
			block = deserialize(b.Get(cur))
			return nil
		})

		fmt.Printf(" 高度：%d\nHash: %s\n",block.Index,block.Hash)
		for _,tx := range block.Data{
			fmt.Printf("%s->%s:%d\n",tx.From,tx.To,tx.Amount)
		}
		fmt.Println("---")

		if len(block.PrevHash)==0{
			break
		}
		cur = []byte(block.PrevHash)
	}
}

// 序列化（Block -> []byte）
func serialize(block Block) []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	enc.Encode(block)
	return buf.Bytes()
}

// 反序列化（[]byte -> Block）
func deserialize(data []byte) Block {
	var block Block
	dec := gob.NewDecoder(bytes.NewReader(data))
	dec.Decode(&block)
	return block
}

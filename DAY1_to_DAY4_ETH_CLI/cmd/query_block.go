package cmd


import (
	"os"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"
	"strconv"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"ethcli/internal"
)


func QueryBlock() {
	rpcURL := os.Args[2]
	blockId := os.Args[3]
	client, ctx, cancel := internal.ConnectToEthereum(rpcURL)
	defer cancel()
	defer client.Close()

	var blockNumber *big.Int
	if strings.HasPrefix(blockId, "0x") {
		n, err := hexutil.DecodeBig(blockId)
		if err != nil {
			log.Fatalf("解析区块号失败：%v", err)
		}
		blockNumber = n
	} else if _, err := strconv.Atoi(blockId); err == nil {
		// 如果是十进制数字
		blockNumber = big.NewInt(0)
		blockNumber.SetString(blockId, 10)
	} else {
		blockNumber = nil
	}

	var err error // 声明 err 变量

	var block *types.Block
	block, err = client.BlockByNumber(ctx, blockNumber)
	if err != nil {
		log.Fatalf("获取区块失败：%v", err)
	}
	fmt.Printf("区块Hash：%s\n", block.Hash().Hex())
	fmt.Printf("区块高度：%d\n", block.Number().Uint64())
	fmt.Printf("区块时间戳：%s\n", time.Unix(int64(block.Time()), 0).Format("2006-01-02 15:04:05"))
	fmt.Printf("区块交易数：%d\n", len(block.Transactions()))

}
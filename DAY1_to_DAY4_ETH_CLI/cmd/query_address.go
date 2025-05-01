package cmd

import (
	"ethcli/internal"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func QueryAddress() {
	rpcURL := os.Args[2]
	targetAddress := common.HexToAddress(os.Args[3])
	client, ctx, cancel := internal.ConnectToEthereum(rpcURL)
	defer cancel()
	defer client.Close()

	//找到eth工具类中查询指定区块地址的最近交易记录的方法，直接调用
	latestHeader, _ := client.HeaderByNumber(ctx, nil)
	latestBlockNumber := latestHeader.Number.Uint64()

	// now := time.Now()
	// cutoff := now.Add(-2 * time.Hour)

	querysize := 500

	for i := int64(latestBlockNumber); i > 0 && i > int64(latestBlockNumber)-int64(querysize); i-- {
		//查询区块对象
		block, err := client.BlockByNumber(ctx, big.NewInt(i))
		if err != nil {
			continue
		}
		//遍历区块对象的transactions，打印出我们需要的交易数据
		for _, tx := range block.Transactions() {
			// 获取发送方地址
			chainID := tx.ChainId()
			if tx.ChainId() == nil || chainID.Cmp(big.NewInt(0)) == 0 {
				continue
			}
			from, err := types.Sender(types.LatestSignerForChainID(chainID), tx)
			if err != nil {
				continue
			}

			// 检查接收方地址是否为nil
			var to common.Address
			if tx.To() != nil {
				to = *tx.To()
			}

			if from == targetAddress || (tx.To() != nil && to == targetAddress) {
				fmt.Printf("区块: %d, 交易Hash: %s, 来自: %s, 发往: %s, 金额: %s\n",
					block.Number().Uint64(), tx.Hash().Hex(), from.Hex(),
					to.String(), tx.Value().String())
			}

		}
	}

}

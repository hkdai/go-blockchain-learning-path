package cmd

import (
	"fmt"
	"os"

	"ethcli/internal"
)

func Execute() {
	if len(os.Args) < 2 {
		fmt.Println("请输入以太坊节点RPC地址")
		return
	}
	rpcURL := os.Args[1]

	client, ctx, cancel := internal.ConnectToEthereum(rpcURL)
	defer cancel()
	defer client.Close()

	header, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		fmt.Println("查询区块链失败,", err)
		return
	}

	fmt.Println("最新区块链高度：", header.Number.Uint64)
}

package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("ws://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	// 读取 ABI
	abiBytes, err := ioutil.ReadFile("vote_abi.json")
	if err != nil {
		log.Fatal("读取 ABI 文件失败:", err)
	}
	contractAbi, err := abi.JSON(strings.NewReader(string(abiBytes)))
	if err != nil {
		log.Fatal("ABI 解析失败:", err)
	}
	if err != nil {
		log.Fatal(err)
	}

	contractAddr := common.HexToAddress("0xDc64a140Aa3E981100a9becA4E685f962f0cF6C9")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddr},
		Topics: [][]common.Hash{
			{contractAbi.Events["Voted"].ID},
		},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for err := range sub.Err() {
			log.Println("订阅错误:", err)
		}
	}()

	fmt.Println("🎧 正在监听 Voted 事件...")

	for log := range logs {
		if len(log.Topics) < 2 {
			fmt.Println("⚠️ 事件 Topics 数量不足，跳过")
			continue
		}
		event := struct {
			Voter    common.Address
			Proposal string
		}{}

		err := contractAbi.UnpackIntoInterface(&event, "Voted", log.Data)
		if err != nil {
			fmt.Println("解码失败:", err)
			continue
		}

		// indexed 参数要手动解
		event.Voter = common.HexToAddress(log.Topics[1].Hex())
		fmt.Printf("📢 %s 投票给 %s\n", event.Voter.Hex(), event.Proposal)
	}
}

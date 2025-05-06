package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	fmt.Println("启动合约调用程序")
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal("连接节点失败:", err)
	}
	defer client.Close()
	fmt.Println("成功连接到本地节点")

	abiBytes, err := ioutil.ReadFile("vote_abi.json")
	if err != nil {
		log.Fatal("读取abi失败：", err)
	}
	parsedABI, err := abi.JSON(strings.NewReader(string(abiBytes)))
	if err != nil {
		log.Fatal("解析ABI失败", err)
	}
	fmt.Println("ABI 函数数量:", len(parsedABI.Methods))
	fmt.Println("ABI加载成功")
	contractAddr := common.HexToAddress("0xCf7Ed3AccA5a467e9e704C703E8D87F634fB0Fc9")
	fmt.Println("✅ 合约地址绑定成功:", contractAddr.Hex())

	//构造调用函数，这里是调用votes，查看里面Option A的投票数量请求
	callData, err := parsedABI.Pack("votes", "Option A")
	if err != nil {
		log.Fatal("构造调用数据失败:", err)
	}
	//构造请求后使用ethclient调用合约函数
	msg := ethereum.CallMsg{
		To:   &contractAddr,
		Data: callData,
	}
	result, err := client.CallContract(context.Background(), msg, nil)
	if err != nil {
		log.Fatal("合约调用失败:", err)
	}
	//调用结束后解码返回值
	var voteCount *big.Int
	err = parsedABI.UnpackIntoInterface(&voteCount, "votes", result)
	if err != nil {
		log.Fatal("解析返回数据失败:", err)
	}

	fmt.Printf("📊 Option A 当前得票数：%s\n", voteCount.String())
}

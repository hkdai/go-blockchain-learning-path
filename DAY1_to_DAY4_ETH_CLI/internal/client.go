package internal

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

func ConnectToEthereum(rpcURL string) (*ethclient.Client, context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	var client *ethclient.Client
	var err error

	if rpcURL == "http://127.0.0.1:7545" {
		// 本地连接直接访问
		client, err = ethclient.DialContext(ctx, rpcURL)
	} else if rpcURL == "https://mainnet.infura.io/v3/5e90e80eb50f4d888db6614644bdc875" {
		// 使用代理访问 Infura
		proxyURL, _ := url.Parse("http://127.0.0.1:1080")
		transport := &http.Transport{Proxy: http.ProxyURL(proxyURL)}
		httpClient := &http.Client{Transport: transport}
		rpcClient, err := rpc.DialHTTPWithClient(rpcURL, httpClient)
		if err == nil {
			client = ethclient.NewClient(rpcClient)
		}
	} else {
		log.Fatalf("不支持的 RPC URL: %s", rpcURL)
	}

	if err != nil {
		cancel()
		log.Fatalf("连接Ethereum节点失败：%v", err)
	}

	return client, ctx, cancel
}

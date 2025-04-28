package internal

import (
	"context"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
)

func ConnectToEthereum(rpcURL string) (*ethclient.Client, context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	client, err := ethclient.DialContext(ctx, rpcURL)
	if err != nil {
		cancel()
		log.Fatalf("连接Ethereum节点失败：%v", err)
	}
	return client, ctx, cancel
}

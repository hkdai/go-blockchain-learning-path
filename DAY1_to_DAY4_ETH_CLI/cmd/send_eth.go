package cmd

import (
	"ethcli/internal"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func SendEth() {
	rpcURL := os.Args[2]
	privKey := os.Args[3]
	// 去除可能存在的0x前缀
	if len(privKey) >= 2 && privKey[:2] == "0x" {
		privKey = privKey[2:]
	}
	to := common.HexToAddress(os.Args[4])
	amount, success := big.NewInt(0).SetString(os.Args[5], 10)
	if !success {
		log.Fatalf("无法解析金额: %s", amount)
	}

	client, ctx, cancel := internal.ConnectToEthereum(rpcURL)
	defer cancel()
	defer client.Close()

	privateKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		log.Fatal(err)
	}

	from := crypto.PubkeyToAddress(privateKey.PublicKey)
	nonce, _ := client.PendingNonceAt(ctx, from)

	gasPrice, _ := client.SuggestGasPrice(ctx)

	tx := types.NewTransaction(nonce, to, amount, 21000, gasPrice, nil)
	// chainID, _ := client.NetworkID(ctx)
	chainID := big.NewInt(1337)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(ctx, signedTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("已发送交易: %s\n", signedTx.Hash().Hex())
}

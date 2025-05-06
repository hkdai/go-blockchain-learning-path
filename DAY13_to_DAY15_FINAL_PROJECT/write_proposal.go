package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal("连接失败：", err)
	}
	defer client.Close()
	fmt.Println("✅ 成功连接节点")

	//加载私钥 & 推导账户地址
	privateKey, err := crypto.HexToECDSA("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80") //此处的私钥，要求不含0x
	if err != nil {
		log.Fatal("私钥解析失败：", err)
	}
	publicKey := privateKey.Public().(*ecdsa.PublicKey)
	fromAddress := crypto.PubkeyToAddress(*publicKey)

	fmt.Println("使用账户地址：", fromAddress.Hex())

	//加载 ABI 并构造调用 data
	abiBytes, _ := ioutil.ReadFile("vote_abi.json")
	parsedABI, _ := abi.JSON(strings.NewReader(string(abiBytes)))

	data, err := parsedABI.Pack("addProposal", "Option C")
	if err != nil {
		log.Fatal("构造调用数据失败:", err)
	}
	//构造并签名交易
	to := common.HexToAddress("0xCf7Ed3AccA5a467e9e704C703E8D87F634fB0Fc9")
	nonce, _ := client.PendingNonceAt(context.Background(), fromAddress)
	gasPrice, _ := client.SuggestGasPrice(context.Background())

	tx := types.NewTransaction(
		nonce,
		to,
		big.NewInt(0),  // 无ETH转账
		uint64(300000), // gas limit
		gasPrice,
		data,
	)
	chainID, _ := client.NetworkID(context.Background())
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal("签名失败:", err)
	}

	//发送交易 & 等待打包
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal("发送失败:", err)
	}
	fmt.Println("交易已发送，Hash:", signedTx.Hash().Hex())

	// 7. 等待回执
	receipt, err := waitForReceipt(client, signedTx.Hash())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("交易确认状态:", receipt.Status)
}

// 等待交易回执
func waitForReceipt(client *ethclient.Client, hash common.Hash) (*types.Receipt, error) {
	for {
		receipt, err := client.TransactionReceipt(context.Background(), hash)
		if err == nil {
			return receipt, nil
		}
		time.Sleep(1 * time.Second)
	}
}

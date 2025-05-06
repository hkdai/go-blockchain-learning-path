# 📅 DAY14：Go语言调用合约写操作 + 签名交易发送

## 🎯 今日目标

* 构造一个合约写操作（如 `addProposal()`）
* 使用私钥签名并发送交易
* 设置链ID、Gas参数、Nonce
* 获取并解析交易回执（含事件日志）

## 🧠 思路拆解

| 步骤    | 内容                           |
| ----- | ---------------------------- |
| ABI加载 | 获取函数的 selector 和 encode 参数   |
| 交易构造  | 填入 to、gas、data、nonce、chainID |
| 签名    | 使用私钥完成 ECDSA 签名              |
| 发送交易  | 使用 `SendRawTransaction`      |
| 回执    | 等待交易被打包确认（含 status、logs）     |

## 🛠️ 实战步骤

### ✅ 第1步：编写连接节点方法,连接到hardhat
```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal("连接失败：", err)
	}
	defer client.Close()
	fmt.Println("✅ 成功连接节点")
}
```

### ✅ 第2步：加载私钥 & 推导账户地址
```go
import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/crypto"
)

privateKey, err := crypto.HexToECDSA("你的私钥（不含0x）")
if err != nil {
	log.Fatal("私钥解析失败:", err)
}
publicKey := privateKey.Public().(*ecdsa.PublicKey)
fromAddress := crypto.PubkeyToAddress(*publicKey)

fmt.Println("使用账户地址:", fromAddress.Hex())
```

### ✅ 第3步：加载 ABI 并构造调用 data
```go
import (
	"io/ioutil"
	"strings"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

abiBytes, _ := ioutil.ReadFile("vote_abi.json")
parsedABI, _ := abi.JSON(strings.NewReader(string(abiBytes)))

data, err := parsedABI.Pack("addProposal", "Option C")
if err != nil {
	log.Fatal("构造调用数据失败:", err)
}
```

### ✅ 第4步：构造并签名交易
```GO
import (
	"math/big"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

to := common.HexToAddress("你的合约地址")
nonce, _ := client.PendingNonceAt(context.Background(), fromAddress)
gasPrice, _ := client.SuggestGasPrice(context.Background())

tx := types.NewTransaction(
	nonce,
	to,
	big.NewInt(0),        // 无ETH转账
	uint64(300000),       // gas limit
	gasPrice,
	data,
)

chainID, _ := client.NetworkID(context.Background())
signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
if err != nil {
	log.Fatal("签名失败:", err)
}
```

### ✅ 第5步：发送交易 & 等待打包
```GO
err = client.SendTransaction(context.Background(), signedTx)
if err != nil {
	log.Fatal("发送失败:", err)
}
fmt.Println("交易已发送，Hash:", signedTx.Hash().Hex())
```

### ✅ 第6步：等待交易回执
```GO
import (
	"time"
)

func waitForReceipt(client *ethclient.Client, hash common.Hash) (*types.Receipt, error) {
	for {
		receipt, err := client.TransactionReceipt(context.Background(), hash)
		if err == nil {
			return receipt, nil
		}
		time.Sleep(1 * time.Second)
	}
}
```

```BASH
go run read_vote.go  
```



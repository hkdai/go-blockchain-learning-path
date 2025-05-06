# 📅 DAY13：Go语言调用合约 + ABI解析实战


## 🎯 今日目标

* 使用 Go + `go-ethereum` 连接本地或远程以太坊节点
* 加载合约 ABI 和合约地址
* 在 Go 中调用合约的读取函数（如 `voteCount()`、`proposals()`）
* 了解调用逻辑和数据结构映射（interface→Go struct）

## 🧠 技术要点

| 模块                      | 内容                              |
| ----------------------- | ------------------------------- |
| `ethclient`             | Go-ethereum 提供的以太坊 JSON-RPC 客户端 |
| `bind.NewBoundContract` | 用于绑定合约并进行方法调用                   |
| ABI 文件                  | 用于定义合约函数和事件格式                   |
| 合约地址                    | 从 Hardhat 部署脚本中获取               |
| 网络 RPC                  | 使用本地 Hardhat 节点或 testnet 节点均可   |

## 🛠️ 实战步骤


### ✅ 第1步：准备 Go 项目

```bash
cd DAY13_to_DAY15_FINAL_PROJECT
go mod init voteclient
go get github.com/ethereum/go-ethereum
```

### ✅ 第2步：复制 ABI 文件
将 Day12 中生成的 vote_abi.json 拷贝进项目目录下


### ✅ 第3步：编写合约调用代码

`main.go`
1. 创建 main.go 文件并初始化框架
2. 连接到以太坊节点（本地 Hardhat 节点）
3. 加载合约 ABI 文件
```go
abiBytes, err := ioutil.ReadFile("vote_abi.json")
if err != nil {
	log.Fatal("读取 ABI 文件失败:", err)
}

parsedABI, err := abi.JSON(strings.NewReader(string(abiBytes)))
if err != nil {
	log.Fatal("解析 ABI 失败:", err)
}
fmt.Println("✅ ABI 加载成功")
```
4. 绑定合约地址
```go
import (
	"github.com/ethereum/go-ethereum/common"
)

contractAddr := common.HexToAddress("0x合约地址")
fmt.Println("✅ 合约地址绑定成功:", contractAddr.Hex())
```
5. 构造调用数据 —— votes("Option A")
```go
callData, err := parsedABI.Pack("votes", "Option A")
if err != nil {
	log.Fatal("构造调用数据失败:", err)
}
```
6. 使用ethclient调用合约函数
```GO
import (
	"context"
	"github.com/ethereum/go-ethereum"
)

msg := ethereum.CallMsg{
	To:   &contractAddr,
	Data: callData,
}

result, err := client.CallContract(context.Background(), msg, nil)
if err != nil {
	log.Fatal("合约调用失败:", err)
}
```
7. 解码返回值
```go
import (
	"math/big"
)

var voteCount *big.Int
err = parsedABI.UnpackIntoInterface(&voteCount, "votes", result)
if err != nil {
	log.Fatal("解析返回数据失败:", err)
}

fmt.Printf("📊 Option A 当前得票数：%s\n", voteCount.String())
```

### ✅ 第4步：运行测试
```bash
go run main.go
```
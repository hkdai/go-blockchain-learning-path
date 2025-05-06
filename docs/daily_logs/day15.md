# 📅 DAY15：合约事件监听 + 实时推送机制


## 🎯 今日目标

* 用 Go 语言监听合约事件（如 `Voted`、`Transfer` 等）
* 解析事件参数（如 address、string、uint256）


## 🧠 技术要点

| 模块   | 内容                                        |
| ---- | ----------------------------------------- |
| 日志订阅 | 使用 `client.FilterLogs()` 获取历史日志           |
| 实时订阅 | 使用 `client.SubscribeFilterLogs()` 监听新区块事件 |
| 事件解析 | 使用 ABI 解码 log.Data、log.Topics 中参数         |
| 推送机制 | 将事件打印 / 写入文件 / 发送 WebSocket 消息等          


## 🛠️ 实战步骤

### ✅ 第1步：确认定义合约事件格式（来自 ABI）

📄 合约中有这个事件：

```solidity
event Voted(address indexed voter, string proposal);
```

### ✅ 第2步：监听合约日志
`event_listener.go`

### ✅ 第3步：触发事件进行测试
调用写好的go语言脚本触发投票事件
```go
go run vote.go 
```

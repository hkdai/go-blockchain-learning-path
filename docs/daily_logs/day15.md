# 📅 DAY15：合约事件监听 + 实时推送机制


## 🎯 今日目标

* 用 Go 语言监听合约事件（如 `Voted`、`Transfer` 等）
* 解析事件参数（如 address、string、uint256）
* 将事件推送到日志、WebSocket、数据库等目标（可拓展）
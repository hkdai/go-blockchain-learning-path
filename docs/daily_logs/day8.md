# 📅 DAY8：本地存储（BoltDB）+ CLI交互控制



## 🎯 今日目标

* 使用 BoltDB 持久化区块链
* 实现 CLI 命令：

  * `addblock` ➜ 添加新交易区块
  * `printchain` ➜ 打印区块链
* 实现简单链遍历和链尾记录


## 1.背景
* 目前为止，我写的链数据都只存在于内存中，只要程序已退出，一切都没了。
* 所以要解决两个实际的需求：

| 问题 | 解决方案 |
|------| -------|
| 程序关闭后链会丢失？   | 使用BoltDB持久化到本地文件 |
| 不想写代码就能操作链？ | 使用CLI命令行控制区块添加与打印 |

## 2.关键技术点快速扫盲
### 1.BoltDB是什么
* Go原生支持的嵌入式键值数据库
* 数据结构：Bucket（相当于表） + Key-Value
* 存储文件如 chain.db可直接落盘，无需服务器

### 2.区块如何存储
* 每个区块的Hash作为Key
* 用gob编码区块struct成字节流存到Value
* 使用特殊的"l" key当前链尾Hash

### 3.CLI怎么控制链
* os.args读取命令参数（不引入额外的库）
* 支持命令：addblock、printchain
* 执行时：
```bash
go run main.go addblock
go run main.go printchain
```



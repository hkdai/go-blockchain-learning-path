# 📅 DAY6：实现工作量证明（Proof of Work，PoW）机制

***

## 🎯 今日目标

为你的每一个区块添加一个「挖矿过程」，通过 PoW 找出一个合法的 `nonce`，使得区块的 `Hash` 满足特定的**难度要求**（如：前面有 n 个 0）。

***

## 🧠 思路分解

| 概念             | 说明                                  |
| :------------- | :---------------------------------- |
| nonce          | 为了让 hash 符合难度，每次加 1 不断尝试            |
| 难度（difficulty） | 要求 hash 前面有几个零（比如 2 个零：`00abcd...`） |
| 哈希函数           | 用 SHA256(block内容 + nonce) 来算        |
| 挖矿             | 指的就是尝试不同 nonce，直到算出合法哈希为止           |

***

### 第1步：更新Block结构
```go
type Block struct{
    Index     int64
	Timestamp int64
	PrevHash  string
	Data      string
	Hash      string
	Nonce     int64    // 新增字段
}
```

### 第2步： 实现挖矿函数（找到符合条件的哈希）
原理就是将Block的各个字段拼接起来，Nonce无限递增去计算hash，直到符合diffculty（挖矿难度）的哈希值算出来

### 第3步：在main.go 中调用PoW挖矿



## 区块链知识


## 🧠 1. 挖矿是在解一个“数学难题”

PoW 挖矿其实就是在解这个问题：

> 找一个合适的数字（我们称之为 `Nonce`），
> 把它和区块内容拼在一起计算 SHA256 哈希，
> 直到这个哈希值“够小”——即它的前面有 N 个 `0`。

---

## 📐 举个例子

假设我们有一段数据要打包进区块，比如这样：

```
Index:     2
Timestamp: 1714860000
PrevHash:  abcdefg...
Data:      "这是区块2"
```

然后我们来加个 `Nonce`：

```
Nonce:     0
```

我们将它们拼接起来形成一段字符串作为“挖矿材料”：

```go
record := fmt.Sprintf("%d%d%s%s%d", index, timestamp, prevHash, data, nonce)
```

这时候：

```go
record = "21714860000abcdefg...这是区块20"
```

我们拿这段字符串去做 SHA256 哈希：

```go
h := sha256.Sum256([]byte(record))
```

比如算出来是：

```
hash = 3e9f52a6c8aa11234b9c7f...
```

但它不是我们要的结果。我们说“**必须前3个字符是0**”，例如：

```
目标hash必须像这样开头：000xxxxxx
```

---

## 🎯 那怎么办？

我们只能**不断修改 Nonce**：

```go
Nonce = 0, hash = 3e9f...
Nonce = 1, hash = 84ac...
Nonce = 2, hash = 9b23...
Nonce = 3, hash = 000abc123456789...
```

当我们试到 Nonce = 3 时，突然发现哈希满足要求了：`以000开头`

这时候我们就说：

> 💡「挖矿成功」，记录这个 nonce，作为这个区块的工作量证明。

---

## 🧩 2. difficulty 是什么意思？

```go
difficulty := 3
prefix := strings.Repeat("0", difficulty)
```

这个就是我们定义的“挖矿难度”：

| difficulty 值 | 代表要求          | 难度级别       |
| ------------ | ------------- | ---------- |
| 1            | hash 前面 1 个 0 | 轻松         |
| 2            | 前面 2 个 0      | 稍难         |
| 3            | 前面 3 个 0      | 挖到手软😅     |
| 4\~5+        | 非常困难          | 挖矿要几秒甚至几十秒 |

因为哈希是伪随机的，要满足 `000...` 这个前缀，需要试很多很多次（每多一个 0，难度就高一倍以上）。

---

## 🔁 3. 完整的挖矿逻辑就是：

```go
for nonce := 0; ; nonce++ {
    record := 区块数据 + nonce
    hash := SHA256(record)
    if hash 前缀有 n 个 0:
        返回这个 nonce + hash
}
```

---

## ✅ 最后简化理解

* 你在试一个数字（Nonce），直到满足“哈希值前n位为0”
* difficulty 就是这个“n”，控制你要试多少次才能成功
* 成功后，这个区块就有了“不可伪造的证明” —— 因为别人要复制你这区块也要试出同样的 nonce，非常耗时

---

## 🔧 提问你几个小练习巩固一下

1. 如果我设置 `difficulty = 1`，你觉得挖矿平均耗时会怎样？
2. 为什么 SHA256 的输出是随机分布的？为什么别人很难直接伪造 hash？
3. 如果我们在 block 中修改了 `Data` 字段，会影响挖矿吗？为什么？

---



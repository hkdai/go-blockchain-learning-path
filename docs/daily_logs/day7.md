# 📅 DAY7：钱包系统 + 模拟交易写入区块（Go）

---

## 🎯 今日目标

* 实现钱包地址系统（使用 ECDSA 公私钥生成地址）
* 构造一个“交易”（转账信息）
* 把交易信息写入区块的 `Data` 字段
* 每个区块支持多个交易

## 步骤

### 第1步：实现钱包地址生成
新建一个wallet.go 

### 第2步：定义交易结构
新建tx.go

### 第3步：修改Block，支持交易数组作为数据
将Data从原本的string改为 []Transaction

同时更新GenerateHash() ,

### 第4步：挖矿时写入交易数据


### 第5步：main.go中模拟用户&转账
1. 创建2个钱包
2. 创世区块
3. 添加交易区块
4. 交易信息打包进挖矿，生成新的区块
5. 输出链内容


## GO知识点
Block里把Data从string改成了[]Transaction，所以这边强化一下切片的知识

## 🧠 1. 切片（slice）是什么？

切片是 Go 中的“变长数组”，用法非常灵活。

它可以：

* 是空的（没有元素）
* 初始化时就指定内容
* 后期通过 `append()` 添加元素

---

## ✅ 基本语法总结

| 写法                                     | 说明                |
| -------------------------------------- | ----------------- |
| `var txs []Transaction`                | 声明一个空的切片变量，值是 nil |
| `txs := []Transaction{}`               | 初始化为空切片（非 nil）    |
| `txs := []Transaction{tx1, tx2}`       | 初始化时指定多个元素        |
| `txs := append(txs, Transaction{...})` | 追加一个元素到已有切片中      |

---

## 🧪 回到你的写法问题

你尝试了三种写法，我们一一来说明：

---

### ❌ 错误写法1：

```go
firstTx := []Transaction
```

> ❌ 这是不合法的语法，因为你只是写了类型，没有给它赋初始值。

Go 的完整声明必须是：

```go
firstTx := []Transaction{}                 // 空值
firstTx := []Transaction{tx1, tx2, ...}   // 带内容
```

---

### ✅ 正确写法1：空切片

```go
firstTx := []Transaction{}
```

> 这是初始化一个空的 `[]Transaction`，可以用来 `append` 添加元素。

---

### ✅ 正确写法2：一次创建一个含有一笔交易的切片

```go
firstTx := []Transaction{
    {
        From:   "SYSTEM",
        To:     "创世用户",
        Amount: 100000,
    },
}
```

注意这里必须加 **`{}` 包住单个 Transaction**，因为切片字面量本质上是“包含多个结构体的数组”。

---

### ✅ 正确写法3：先写结构体，再装进切片

```go
tx1 := Transaction{From: "A", To: "B", Amount: 10}
tx2 := Transaction{From: "B", To: "C", Amount: 20}

txs := []Transaction{tx1, tx2}
```

---

### ✅ 正确写法4：先空切片，再 append

```go
txs := []Transaction{}
txs = append(txs, Transaction{From: "X", To: "Y", Amount: 1})
```

> 非常适合在循环里动态添加交易！

---

## 🧠 延伸理解：为什么切片必须这样写？

因为切片本质上是“引用 + 长度 + 容量”的三元结构，所以你必须明确：

* 创建它（`[]Transaction{}`）
* 给它数据（`append` 或 `{...}` 初始化）

---

## 📝 总结口诀：

> ✅ `[]类型{}`：创建空切片
> ✅ `[]类型{元素1, 元素2}`：带内容初始化
> ✅ `append(切片, 元素)`：动态添加元素

---

你可以试试以下代码练练手：

```go
package main

import "fmt"

type Transaction struct {
	From   string
	To     string
	Amount int
}

func main() {
	// 方法1
	txs := []Transaction{}

	// 方法2
	tx1 := Transaction{"A", "B", 10}
	txs = append(txs, tx1)

	// 方法3
	txs = append(txs, Transaction{From: "B", To: "C", Amount: 20})

	// 方法4（直接初始化）
	txs2 := []Transaction{
		{From: "SYSTEM", To: "Alice", Amount: 100},
		{From: "Alice", To: "Bob", Amount: 30},
	}

	fmt.Println(txs)
	fmt.Println(txs2)
}
```

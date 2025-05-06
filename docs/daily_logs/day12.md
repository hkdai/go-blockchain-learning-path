# 📅 DAY12：综合合约项目实战（投票系统 or NFT）+ 跨语言调用准备


## 🎯 今日目标

* 实现一个完整的智能合约项目（含状态管理、权限控制、事件通知）
* 撰写部署 + 脚本调用 + 自动测试
* 输出 ABI & 合约地址，为 Go / Java / Python 调用打基础

###  第1步：智能合约编写
`contracts/Vote.sol`


### 第2步：部署脚本
`scripts/deploy_vote.js`
```bash
npx hardhat run scripts/deploy_vote.js --network localhost
```

### 第3步：调用脚本（添加提案 + 投票）
`scripts/vote_actions.js`
```bash
npx hardhat run scripts/vote_actions.js --network localhost
```

### 第4步：测试 Vote 合约
`test/vote_test.js`
```bash
npx hardhat test
```

### 第5步：导出 ABI + 地址
`scripts/export_abi.js`

```bash
npx hardhat run scripts/export_abi.js
```

# 知识点



## ✅ 什么是 ABI？为什么你需要它？

**ABI（Application Binary Interface）** 是合约和外部调用者之间的“通信协议”。

你可以理解成“合约说明书”：

| 概念       | 说明                                       |
| -------- | ---------------------------------------- |
| ABI 是什么？ | 合约的函数名、参数类型、事件结构的清单，供链下程序（Go/JS）调用时参考    |
| 有什么用？    | 让链下程序知道你写的合约有哪些函数、接收哪些参数、返回什么结果          |
| 谁用 ABI？  | 不是给用户看的，是给链下程序、前端、RPC 客户端看的              |
| 文件格式？    | 一个 `.json` 文件，里面是一个数组，每个元素表示一个函数、事件、构造器等 |

---

## 🧠 举个例子：看一段 ABI 项

你导出的 `vote_abi.json` 里面可能包含这样一段：

```json
{
  "inputs": [
    { "internalType": "string", "name": "name", "type": "string" }
  ],
  "name": "addProposal",
  "outputs": [],
  "stateMutability": "nonpayable",
  "type": "function"
}
```

这意味着：

* 合约有个叫 `addProposal` 的函数
* 它接收一个字符串（提案名称）
* 不返回值
* 是“非支付型调用”（不能带ETH）

---

## ✅ ABI 主要用于哪里？

| 场景                  | ABI 用途                  |
| ------------------- | ----------------------- |
| 前端调用合约（ethers/web3） | ABI 决定你调用哪个函数、传什么参数     |
| 后端（Go、Java）调用合约     | 要通过 ABI 解析函数签名和参数编码格式   |
| 监听事件                | ABI 告诉你合约有哪些事件，如何解码事件数据 |

---

## ✅ 如何配合使用 ABI + 地址？

你调用合约通常需要两样东西：

| 元素   | 说明              |
| ---- | --------------- |
| ABI  | 用于理解合约函数结构      |
| 合约地址 | 用于定位部署在链上的哪一个合约 |

---

## ✅ 总结口诀

> ABI 就是“链下程序调用链上合约的接口说明书”，不是代码，但比代码还关键！

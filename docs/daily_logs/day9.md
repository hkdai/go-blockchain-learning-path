# 📅 DAY9：Solidity 合约开发入门 + Hardhat 初体验

## 🎯 今日目标

* 搭建智能合约开发环境（Hardhat）
* 编写第一个 Solidity 合约（简单代币或存储器）
* 本地编译、部署合约
* 使用脚本或命令行调用合约

### 第1步：初始化Hardhat项目

```bash
cd DAY9_to_DAY12_DAPP_PROJECT
npm init -y # 创建一个 package.json 文件，Node.js 所有依赖都记录在这里。
npm install --save-dev hardhat # --save-dev 表示这是开发依赖，不是生产环境用的。
npx hardhat #初始化hardhat项目结构
```

### 第2步：编写第一个solidity合约
1. 创建合约文件contracts/Storage.sol
```solidity
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

contract Storage {
    uint256 private value;

    function store(uint256 _value) public {
        value = _value;
    }

    function retrieve() public view returns (uint256) {
        return value;
    }
}
```
2. 编写一个简单的set/get函数的合约，能读能写

### 第3步：部署合约脚本
1. 创建部署脚本scripts/deploy.js
```js
const hre = require("hardhat");

async function main() {
  const Storage = await hre.ethers.getContractFactory("Storage");
  const storage = await Storage.deploy();
  await storage.deployed();

  console.log("✅ 合约已部署，地址:", storage.address);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
```
* getContractFactory("Storage") 会编译并加载合约
* .deploy() 会部署到链上
* .deployed() 会等部署完成后再继续执行
### 第4步：运行本地链并部署
1. 启动本地链
```bash
npx hardhat node
```
将看到 Hardhat 启动了本地区块链，并生成了20个测试账户。

2. 另开一个终端，部署合约
```bash
npx hardhat run scripts/deploy.js --network localhost
```
应该会看到,合约部署成功提示
### 第5步：使用控制台调用合约
1. 启动控制台
```bash
npx hardhat console --network localhost
```
将进入 Hardhat 的 JavaScript 控制台环境，可以直接调用合约函数。

2. 在控制台中输入以下命令
```js
const [owner] = await ethers.getSigners()
const Storage = await ethers.getContractFactory("Storage")
const storage = await Storage.attach("0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0") //这里替换为你的合约地址
```
.attach() 表示连接到已经部署的合约实例。

3. 读写合约
```js
await storage.store(123)           // 写入
(await storage.retrieve()).toString()   // 读取并转字符串
```
应该会看到输出：
```bash
'123'
```
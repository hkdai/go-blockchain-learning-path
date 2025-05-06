# 📅 DAY11：权限控制合约 + 脚本自动调用

## 🎯 今日目标

* 编写一个具备「管理员权限」的智能合约
* 合约可添加/删除白名单成员
* 只有白名单成员能调用关键函数
* 编写脚本批量操作合约函数（添加成员/调用功能）


### 第1步：合约结构设计
`contracts/WhitelistVault.sol`


### 第2步：部署合约
`scripts/deploy_whitelist.js`

```bash
npx hardhat run scripts/deploy_whitelist.js --network localhost
```

### 第3步：调用脚本
`scripts/claim_and_manage.js`


运行：
```bash
npx hardhat run scripts/claim_and_manage.js --network localhost
```

### 第4步：测试合约逻辑
`test/whitelist_test.js`

```bash
npx hardhat test
```
# 📅 DAY10：ERC20合约实现 + 单元测试 + 事件机制


## 🎯 今日目标

* 编写并部署一个最小化版本的 ERC20 合约
* 学会使用 `emit` 事件追踪合约行为
* 使用 Hardhat 测试框架进行合约自动化测试

### 第1步：创建ERC20合约
`contracts/MyToken.sol`

```solidity
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

contract MyToken{
    string public name = "MyToken";
    string public symbol = "MTK";
    uint8 public decimals = 18;
    uint256 public totalSupply;

    mapping(address =>uint256) public balanceOf;

    event Transfer(address indexed from, address indexed to,uint256 value);

    constructor(uint256 _initialSupply){
        totalSupply = _initialSupply;
        balanceOf[msg.sender] = _initialSupply;
        emit Transfer( address(0),msg.sender,_initialSupply);
    }

    function transfer(address to,uint256 amount) public returns(bool){
        require(balanceOf[msg.sender] >=amount, "balance not enough");
        balanceOf[msg.sender]-=amount;
        balanceOf[to]+=amount;
        emit Transfer(msg.sender, to, amount);
        return true;
    }

}
```

编写完以后可以运行编译命令查看是否通过编译
```bash
npx hardhat compile
```

### 第2步：编写部署脚本
`scripts/deploy_token.js`

运行部署
```bash
npx hardhat run scripts/deploy_token.js --network localhost
```


### 第3步：编写单元测试
`test/token_test.js`

运行测试
```bash
npx hardhat test
```


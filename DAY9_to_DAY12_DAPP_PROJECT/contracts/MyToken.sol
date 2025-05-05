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
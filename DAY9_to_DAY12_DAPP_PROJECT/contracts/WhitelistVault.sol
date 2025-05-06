// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

contract WhitelistVault {
    address public owner;
    mapping(address=>bool) public whitelist;
    mapping(address=>uint256) public claimed;

    event Added(address indexed user);
    event Removed(address indexed user);
    event Claimed(address indexed user, uint256 amount);

    modifier onlyOwner(){
        require(msg.sender==owner,"only admin");
        _;
    }

    modifier onlyWhitelisted(){
        require(whitelist[msg.sender],"not in whitelist");
        _;
    }

    constructor(){
        owner = msg.sender;
    }

    function addToWhitelist(address user) public onlyOwner{
        whitelist[user] = true;
        emit Added(user);
    }

    function removeFromeWhitelist(address user) public onlyOwner{
        whitelist[user] = false;
        emit Removed(user);
    }

    function claimReward() public onlyWhitelisted{
        require(claimed[msg.sender]==0,"claimed yet");
        claimed[msg.sender] = 100;
        emit Claimed(msg.sender, 100);
    }

}
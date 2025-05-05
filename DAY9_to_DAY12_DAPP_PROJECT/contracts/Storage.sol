// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

contract Storage {
    uint256 private value;

    function store(uint256 _value) public{
        value = _value;
    }

    function retrieve() public view returns(uint256){
        return value;
    }
}
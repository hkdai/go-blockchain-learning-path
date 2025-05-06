// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

contract Vote {
    address public owner;
    string[] public proposals;
    mapping(string=>uint256) public votes;
    mapping(address=> bool) public hasVoted;

    event ProposalAdded(string proposal);
    event Voted(address voter,string proposal);

    modifier onlyOwner(){
        require(msg.sender==owner,"only admin");
        _;
    }

    constructor(){
        owner = msg.sender;
    }

    function addProposal(string memory name) public onlyOwner {
        proposals.push(name);
        emit ProposalAdded(name);
    }

    function vote(string memory name) public {
        require(!hasVoted[msg.sender], "voted");
        votes[name] += 1;
        hasVoted[msg.sender] = true;
        emit Voted(msg.sender, name);
    }

    function getProposalCount() public view returns (uint256) {
        return proposals.length;
    }
}
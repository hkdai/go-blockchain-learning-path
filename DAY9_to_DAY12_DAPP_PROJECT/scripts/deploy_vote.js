const hre = require("hardhat");

async function main() {
  const Vote = await hre.ethers.getContractFactory("Vote");
  const vote = await Vote.deploy();
  await vote.waitForDeployment(); 
  console.log("✅ 合约部署成功，地址:", await vote.getAddress());
}

main().catch(console.error);
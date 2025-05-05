const hre = require("hardhat");

async function main() {
  const [deployer] = await hre.ethers.getSigners();
  const Token = await hre.ethers.getContractFactory("MyToken");
  const token = await Token.deploy("1000000000000000000000"); // 1000 MTK
  await token.waitForDeployment(); 
  console.log("✅ 合约部署成功，地址:", await token.getAddress());
  console.log("部署者余额:", (await token.balanceOf(deployer.address)).toString());
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
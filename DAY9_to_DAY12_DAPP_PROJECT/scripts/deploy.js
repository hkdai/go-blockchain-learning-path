const hre = require("hardhat");

async function main() {
  const Storage = await hre.ethers.getContractFactory("Storage");
  const storage = await Storage.deploy();

  console.log("等待合约部署确认...");
  await storage.waitForDeployment(); // ✅ 替代 .deployed()
  console.log("✅ 合约部署成功，地址:", await storage.getAddress());

}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
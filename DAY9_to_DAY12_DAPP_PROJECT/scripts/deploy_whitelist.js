const hre = require("hardhat");

async function main() {
  const Vault = await hre.ethers.getContractFactory("WhitelistVault");
  const vault = await Vault.deploy();

  await vault.waitForDeployment(); 
  console.log("✅ 合约部署成功，地址:", await vault.getAddress());
}

main().catch((err) => {
  console.error(err);
  process.exit(1);
});
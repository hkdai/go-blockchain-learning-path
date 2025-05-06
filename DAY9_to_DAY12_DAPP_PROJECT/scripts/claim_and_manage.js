const hre = require("hardhat");

async function main() {
  const [owner, user1] = await hre.ethers.getSigners();
  const contractAddr = "0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512";

  const Vault = await hre.ethers.getContractFactory("WhitelistVault");
  const vault = await Vault.attach(contractAddr);

  // 添加 user1 到白名单
  await vault.addToWhitelist(user1.address);
  console.log("添加白名单:", user1.address);

  // user1 调用 claimReward（模拟用 user1 身份）
  const vaultFromUser = vault.connect(user1);
  await vaultFromUser.claimReward();

  const amount = await vaultFromUser.claimed(user1.address);
  console.log("user1 已领取:", amount.toString());
}

main().catch(console.error);
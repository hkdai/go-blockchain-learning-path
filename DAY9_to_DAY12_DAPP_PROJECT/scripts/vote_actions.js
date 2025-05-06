const hre = require("hardhat");

async function main() {
  const [owner, user1] = await hre.ethers.getSigners();
  const addr = "0xCf7Ed3AccA5a467e9e704C703E8D87F634fB0Fc9";

  const Vote = await hre.ethers.getContractFactory("Vote");
  const vote = await Vote.attach(addr);

  await vote.addProposal("Option A");
  await vote.addProposal("Option B");

  const voteAsUser1 = vote.connect(user1);
  await voteAsUser1.vote("Option B");

  const count = await vote.votes("Option B");
  console.log("Option B 得票数:", count.toString());
}

main().catch(console.error);
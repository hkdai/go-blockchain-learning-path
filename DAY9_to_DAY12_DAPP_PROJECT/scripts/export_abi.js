const fs = require("fs");

async function main() {
  const artifact = await hre.artifacts.readArtifact("Vote");
  fs.writeFileSync("vote_abi.json", JSON.stringify(artifact.abi, null, 2));
  console.log("✅ ABI 导出完成");
}

main();
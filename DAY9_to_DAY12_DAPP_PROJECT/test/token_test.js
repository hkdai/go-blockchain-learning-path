const { expect } = require("chai");

describe("MyToken", function () {
  it("应正确部署并给部署者初始代币", async function () {
    const [owner, addr1] = await ethers.getSigners();
    const Token = await ethers.getContractFactory("MyToken");
    const token = await Token.deploy("1000");

    await token.waitForDeployment(); 
    expect(await token.balanceOf(owner.address)).to.equal(1000);
  });

  it("应正确完成转账", async function () {
    const [owner, addr1] = await ethers.getSigners();
    const Token = await ethers.getContractFactory("MyToken");
    const token = await Token.deploy("1000");

    await token.transfer(addr1.address, 200);
    expect(await token.balanceOf(addr1.address)).to.equal(200);
    expect(await token.balanceOf(owner.address)).to.equal(800);
  });

  it("应触发 Transfer 事件", async function () {
    const [owner, addr1] = await ethers.getSigners();
    const Token = await ethers.getContractFactory("MyToken");
    const token = await Token.deploy("1000");

    await expect(token.transfer(addr1.address, 100))
      .to.emit(token, "Transfer")
      .withArgs(owner.address, addr1.address, 100);
  });
});
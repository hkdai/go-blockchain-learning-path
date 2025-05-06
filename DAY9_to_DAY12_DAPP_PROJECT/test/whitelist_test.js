describe("WhitelistVault", () => {
    it("应限制非白名单用户 claim", async () => {
      const [owner, user1] = await ethers.getSigners();
      const Vault = await ethers.getContractFactory("WhitelistVault");
      const vault = await Vault.deploy();
      await vault.waitForDeployment(); 
      await expect(
        vault.connect(user1).claimReward()
      ).to.be.revertedWith("不在白名单");
  
      await vault.addToWhitelist(user1.address);
      await vault.connect(user1).claimReward();
      const claimed = await vault.claimed(user1.address);
      expect(claimed).to.equal(100);
    });
  });
describe("Vote", function () {
    it("应正确添加提案和投票", async function () {
      const [owner, voter] = await ethers.getSigners();
      const Vote = await ethers.getContractFactory("Vote");
      const vote = await Vote.deploy();
      await vote.waitForDeployment(); 
  
      await vote.addProposal("Alice");
      await vote.connect(voter).vote("Alice");
      const count = await vote.votes("Alice");
      expect(count).to.equal(1);
    });
  });
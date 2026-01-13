const { expect } = require("chai");
const { ethers } = require("hardhat");

describe("ApprovalService", function () {
  it("approve and revoke", async function () {
    const [admin, user] = await ethers.getSigners();
    const AS = await ethers.getContractFactory("ApprovalService");
    const as = await AS.deploy(admin.address);
    await as.deployed();

    expect(await as.check(user.address)).to.equal(false);
    await as.connect(admin).approve(user.address);
    expect(await as.check(user.address)).to.equal(true);
    await as.connect(admin).revoke(user.address);
    expect(await as.check(user.address)).to.equal(false);
  });
});

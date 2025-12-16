const { expect } = require("chai");
const { ethers } = require("hardhat");

describe("PlatformRegistry", function () {
  it("deploys and sets addresses", async function () {
    const [admin, other] = await ethers.getSigners();
    const Registry = await ethers.getContractFactory("PlatformRegistry");
    const registry = await Registry.deploy(admin.address);
    await registry.deployed();

    expect(await registry.getPropertyFactory()).to.equal(ethers.constants.AddressZero);

    await expect(registry.connect(other).setPropertyFactory(other.address)).to.be.reverted;
    await registry.connect(admin).setPropertyFactory(other.address);
    expect(await registry.getPropertyFactory()).to.equal(other.address);
  });
});

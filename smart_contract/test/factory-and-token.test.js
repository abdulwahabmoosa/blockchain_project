const { expect } = require("chai");
const { ethers } = require("hardhat");

describe("Factory + Asset + Token flow", function () {
  it("creates asset and token, mints tokens", async function () {
    const [admin, owner, investor] = await ethers.getSigners();

    const Registry = await ethers.getContractFactory("PlatformRegistry");
    const registry = await Registry.deploy(admin.address);

    const Approval = await ethers.getContractFactory("ApprovalService");
    const approval = await Approval.deploy(admin.address);
    await registry.connect(admin).setApprovalService(approval.address);

    const Factory = await ethers.getContractFactory("PropertyFactory");
    const factory = await Factory.deploy(registry.address, admin.address);
    await registry.connect(admin).setPropertyFactory(factory.address);

    // Create property
    await factory.connect(admin).createProperty(
      owner.address,
      "AssetName",
      "AST",
      "QmHash",
      1000000,
      ethers.utils.parseUnits("1000", 18),
      "TokenName",
      "TKN"
    );

    // More asserts: events, balances etc. (left to complete)
  });
});

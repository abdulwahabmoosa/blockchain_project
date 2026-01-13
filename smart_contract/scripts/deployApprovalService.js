const hre = require("hardhat");

async function main() {
  // Use the specific admin address (the one the user is logged in with)
  const adminAddress = "0x98e5a749e25c56e19c28008505df75aff4988049";

  const Approval = await hre.ethers.getContractFactory("ApprovalService");
  const approval = await Approval.deploy(adminAddress);
  await approval.deployed();

  console.log("ApprovalService:", approval.address);
  console.log("Admin granted permissions:", adminAddress);

  const registryAddr = process.env.REGISTRY;
  if (!registryAddr) {
    console.log("REGISTRY env variable not set");
    return;
  }

  const Registry = await hre.ethers.getContractAt("PlatformRegistry", registryAddr);
  await Registry.setApprovalService(approval.address);

  console.log("Registry updated with ApprovalService");
}

// main().catch(err => {
//   console.error(err);
//   process.exitCode = 1;
// });
main()
  .then(() => process.exit(0))
  .catch((err) => {
    console.error(err);
    process.exit(1);
  });

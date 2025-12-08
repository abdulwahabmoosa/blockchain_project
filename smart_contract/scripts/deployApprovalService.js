const hre = require("hardhat");

async function main() {
  const [admin] = await hre.ethers.getSigners();

  const Approval = await hre.ethers.getContractFactory("ApprovalService");
  const approval = await Approval.deploy(admin.address);
  await approval.deployed();

  console.log("ApprovalService:", approval.address);

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

const hre = require("hardhat");

async function main() {
  const [deployer] = await hre.ethers.getSigners();
  console.log("Deploying from", deployer.address);

  const Registry = await hre.ethers.getContractFactory("PlatformRegistry");
  const registry = await Registry.deploy(deployer.address);
  await registry.deployed();
  console.log("PlatformRegistry:", registry.address);
}

// main().catch((error) => {
//   console.error(error);
//   process.exitCode = 1;
// });

main()
  .then(() => process.exit(0))
  .catch((err) => {
    console.error(err);
    process.exit(1);
  });

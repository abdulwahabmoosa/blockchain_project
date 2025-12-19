const hre = require("hardhat");

async function main() {
  const [deployer] = await hre.ethers.getSigners();
  console.log("Deploying from", deployer.address);
  console.log("Network:", hre.network.name);

  // Use PRIVATE_KEY_2 address as admin (0x98E5a749E25C56e19C28008505DF75aFf4988049)
  // This ensures PRIVATE_KEY_2 has admin role in the registry
  const adminAddress = '0x98E5a749E25C56e19C28008505DF75aFf4988049';
  console.log("Admin address (PRIVATE_KEY_2):", adminAddress);

  const Registry = await hre.ethers.getContractFactory("PlatformRegistry");
  const registry = await Registry.deploy(adminAddress);
  await registry.deployed();
  console.log("✅ PlatformRegistry deployed at:", registry.address);
  console.log("✅ Admin (PRIVATE_KEY_2) has DEFAULT_ADMIN_ROLE");
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

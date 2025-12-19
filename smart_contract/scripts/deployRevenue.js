const hre = require('hardhat');

async function main() {
  const [deployer] = await hre.ethers.getSigners();
  console.log("Deploying from", deployer.address);
  console.log("Network:", hre.network.name);

  // Use PRIVATE_KEY_2 address as admin (0x98E5a749E25C56e19C28008505DF75aFf4988049)
  // This ensures PRIVATE_KEY_2 has admin and distributor roles
  const adminAddress = '0x98E5a749E25C56e19C28008505DF75aFf4988049';
  console.log("Admin address (PRIVATE_KEY_2):", adminAddress);

  const RD = await hre.ethers.getContractFactory('RevenueDistribution');
  const rd = await RD.deploy(adminAddress);
  await rd.deployed();
  console.log('✅ RevenueDistribution deployed at:', rd.address);
  console.log('✅ Admin (PRIVATE_KEY_2) has DEFAULT_ADMIN_ROLE and DISTRIBUTOR_ROLE');
}

// main().catch((err) => {
//   console.error(err);
//   process.exitCode = 1;
// })
main()
  .then(() => process.exit(0))
  .catch((err) => {
    console.error(err);
    process.exit(1);
  });

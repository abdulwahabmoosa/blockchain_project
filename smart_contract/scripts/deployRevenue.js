const hre = require('hardhat');

async function main() {
  const [deployer] = await hre.ethers.getSigners();

  const RD = await hre.ethers.getContractFactory('RevenueDistribution');
  const rd = await RD.deploy(deployer.address);
  await rd.deployed();
  console.log('RevenueDistribution:', rd.address);
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

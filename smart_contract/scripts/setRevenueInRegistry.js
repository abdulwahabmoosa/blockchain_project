require('dotenv').config();
const hre = require('hardhat');

async function main() {
  const [admin] = await hre.ethers.getSigners();

  const registryAddr = process.env.REGISTRY;
  const revenueAddr = process.env.REVENUE; // store this in .env

  if (!registryAddr || !revenueAddr) {
    console.error('Missing REGISTRY or REVENUE address in .env');
    return;
  }

  const Registry = await hre.ethers.getContractAt(
    'PlatformRegistry',
    registryAddr
  );

  console.log('Setting RevenueDistribution in registry...');
  const tx = await Registry.connect(admin).setRevenueDistribution(revenueAddr);
  await tx.wait();

  console.log('✔ RevenueDistribution set to:', revenueAddr);
}

main()
  .then(() => process.exit(0))
  .catch((err) => {
    console.error(err);
    process.exit(1);
  });

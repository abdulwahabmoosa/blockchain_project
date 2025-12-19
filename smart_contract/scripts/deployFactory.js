// const hre = require("hardhat");

// async function main() {
//   const [deployer] = await hre.ethers.getSigners();
//   const registryAddr = process.env.REGISTRY || ""; // set env or paste address

//   if (!registryAddr) {
//     console.error("Set REGISTRY env variable to registry address.");
//     return;
//   }

//   const Factory = await hre.ethers.getContractFactory("PropertyFactory");
//   const factory = await Factory.deploy(registryAddr, deployer.address);
//   await factory.deployed();
//   console.log("PropertyFactory:", factory.address);

//   // optionally set in registry if deployed by same admin
//   const Registry = await hre.ethers.getContractAt("PlatformRegistry", registryAddr);
//   await Registry.setPropertyFactory(factory.address);
//   console.log("Registry updated with factory address");
// }

// // main().catch((error) => {
// //   console.error(error);
// //   process.exitCode = 1;
// // });
// main()
//   .then(() => process.exit(0))
//   .catch((err) => {
//     console.error(err);
//     process.exit(1);
//   });
require('dotenv').config({ path: '../.env' });
const hre = require('hardhat');

async function main() {
  // Get the deployer signer (will be PRIVATE_KEY_2 for sepolia network)
  const [deployer] = await hre.ethers.getSigners();
  
  console.log('Deploying from address:', deployer.address);
  console.log('Network:', hre.network.name);

  // Load registry address from environment (Sepolia registry)
  const registryAddr = process.env.REGISTRY;
  
  if (!registryAddr) {
    console.error('❌ REGISTRY environment variable not found!');
    console.error('   Make sure REGISTRY is set in the .env file');
    process.exit(1);
  }

  console.log('Using Registry:', registryAddr);

  // Deploy factory with deployer.address as admin
  // This will give deployer (PRIVATE_KEY_2) both DEFAULT_ADMIN_ROLE and CREATOR_ROLE
  console.log('Deploying PropertyFactory...');
  const Factory = await hre.ethers.getContractFactory('PropertyFactory');
  const factory = await Factory.deploy(registryAddr, deployer.address);
  await factory.deployed();

  console.log('✅ PropertyFactory deployed at:', factory.address);
  console.log('✅ Admin address (PRIVATE_KEY_2):', deployer.address);

  // Attach registry
  const Registry = await hre.ethers.getContractAt(
    'PlatformRegistry',
    registryAddr,
    deployer
  );

  // Update registry with new factory address (IMPORTANT: wait for tx)
  console.log('Updating PlatformRegistry with new factory address...');
  const tx = await Registry.setPropertyFactory(factory.address);
  await tx.wait();

  console.log('✅ Registry updated with Factory address');

  // Verify registry value
  const result = await Registry.getPropertyFactory();
  console.log('✅ Registry now stores Factory:', result);
  
  console.log('\n📋 IMPORTANT: Update the following files with the new factory address:');
  console.log('   1. frontend/src/deployments/sepolia/contract-addresses.json');
  console.log('   2. backend/.env (if REGISTRY needs updating)');
  console.log('\n   New PropertyFactory address:', factory.address);
}

main().catch((err) => {
  console.error(err);
  process.exit(1);
});

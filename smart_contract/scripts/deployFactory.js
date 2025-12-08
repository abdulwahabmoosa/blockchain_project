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
require('dotenv').config();
const hre = require('hardhat');

async function main() {
  const [admin] = await hre.ethers.getSigners();

  // Load registry address from env
  const registryAddr = process.env.REGISTRY;
  if (!registryAddr) throw new Error('❌ REGISTRY env variable is missing!');

  console.log('Using Registry:', registryAddr);

  // Deploy factory
  const Factory = await hre.ethers.getContractFactory('PropertyFactory');
  const factory = await Factory.deploy(registryAddr, admin.address);
  await factory.deployed();

  console.log('PropertyFactory deployed at:', factory.address);

  // Attach registry
  const Registry = await hre.ethers.getContractAt(
    'PlatformRegistry',
    registryAddr
  );

  // Update registry (IMPORTANT: wait for tx)
  const tx = await Registry.setPropertyFactory(factory.address);
  await tx.wait();

  console.log('Registry updated with Factory address');

  // Verify registry value
  const result = await Registry.getPropertyFactory();
  console.log('Registry now stores Factory:', result);
}

main().catch((err) => {
  console.error(err);
  process.exit(1);
});

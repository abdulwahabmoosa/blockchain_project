const hre = require('hardhat');
const fs = require('fs');
const path = require('path');

async function main() {
  const [admin] = await hre.ethers.getSigners();

  const addressesPath = path.join(__dirname, '../deployments/sepolia/contract-addresses.json');
  if (!fs.existsSync(addressesPath)) {
    console.error('contract-addresses.json not found at', addressesPath);
    return;
  }

  const data = JSON.parse(fs.readFileSync(addressesPath, 'utf8'));
  const properties = data.properties || [];

  if (properties.length === 0) {
    console.log('No properties found in contract-addresses.json');
    return;
  }

  const index = process.env.PROPERTY_INDEX
    ? parseInt(process.env.PROPERTY_INDEX, 10)
    : 0;

  if (index < 0 || index >= properties.length) {
    console.error('PROPERTY_INDEX out of range. Available properties:', properties.length);
    return;
  }

  const property = properties[index];
  const assetAddr = property.asset;

  console.log(`Rejecting property at index ${index}, asset address: ${assetAddr}`);

  const Asset = await hre.ethers.getContractAt('PropertyAsset', assetAddr);

  const tx = await Asset.connect(admin).rejectProperty();
  console.log('Transaction submitted:', tx.hash);
  await tx.wait();

  console.log('Property rejected on-chain. Asset address:', assetAddr);
}

main()
  .then(() => process.exit(0))
  .catch((err) => {
    console.error(err);
    process.exit(1);
  });



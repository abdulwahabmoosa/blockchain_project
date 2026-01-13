const fs = require('fs');
const path = require('path');
const hre = require('hardhat');

async function main() {
  const addressesPath = path.join(
    __dirname,
    '../deployments/sepolia/contract-addresses.json'
  );
  const data = JSON.parse(fs.readFileSync(addressesPath, 'utf8'));

  const factoryAddr = data.PropertyFactory;
  if (!factoryAddr) throw new Error('PropertyFactory address missing');

  const Factory = await hre.ethers.getContractAt(
    'PropertyFactory',
    factoryAddr
  );

  // Scan last 5000 blocks
  const latest = await hre.ethers.provider.getBlockNumber();
  const fromBlock = latest - 5000 > 0 ? latest - 5000 : 0;

  console.log(`Scanning blocks ${fromBlock} → ${latest}...`);

  const events = await Factory.queryFilter(
    Factory.filters.PropertyRegistered(),
    fromBlock,
    latest
  );

  if (events.length === 0) {
    console.log('NO PropertyRegistered events found in last 5000 blocks.');
    return;
  }

  const evt = events[events.length - 1];

  const newProperty = {
    asset: evt.args.propertyAsset,
    token: evt.args.propertyToken,
    owner: evt.args.owner,
    valuation: evt.args.valuation.toString(),
    metadataHash: evt.args.propertyDataHash,
    createdAtBlock: evt.blockNumber,
  };

  data.properties.push(newProperty);

  fs.writeFileSync(addressesPath, JSON.stringify(data, null, 2));

  console.log('✓ Property saved:');
  console.log(newProperty);
}

main()
  .then(() => process.exit(0))
  .catch((err) => {
    console.error(err);
    process.exit(1);
  });

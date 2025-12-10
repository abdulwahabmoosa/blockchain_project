const hre = require("hardhat");
const fs = require("fs");
const path = require("path");

async function main() {
  const [admin, owner] = await hre.ethers.getSigners();

  const addressesPath = path.join(__dirname, "../deployments/sepolia/contract-addresses.json");
  const data = JSON.parse(fs.readFileSync(addressesPath, "utf8"));

  const registryAddr = data.PlatformRegistry;
  const Registry = await hre.ethers.getContractAt("PlatformRegistry", registryAddr);

  const factoryAddr = data.PropertyFactory;
  const Factory = await hre.ethers.getContractAt("PropertyFactory", factoryAddr);

  console.log("Creating property...");


  //for testing purpose but frontend need to fill it later
  const tx = await Factory.createProperty(
    owner.address,
    "DemoProperty",
    "DPR",
    "QmHashHere",
    1000000,
    hre.ethers.utils.parseUnits("1000", 18),
    "DemoToken",
    "DTK"
  );

  const receipt = await tx.wait();
  console.log("Property created!");

  // Extract event directly from receipt
  const evt = receipt.events.find(e => e.event === "PropertyRegistered");

  if (!evt) {
    console.log("ERROR: PropertyRegistered event not found.");
    return;
  }

  const newProperty = {
    asset: evt.args.propertyAsset,
    token: evt.args.propertyToken,
    owner: evt.args.owner,
    valuation: evt.args.valuation.toString(),
    metadataHash: evt.args.propertyDataHash,
    createdAtBlock: receipt.blockNumber,
    txHash: receipt.transactionHash
  };

  data.properties.push(newProperty);

  fs.writeFileSync(addressesPath, JSON.stringify(data, null, 2));

  console.log("Property saved to contract-addresses.json:");
  console.log(newProperty);
}

main()
  .then(() => process.exit(0))
  .catch((err) => {
    console.error(err);
    process.exit(1);
  });
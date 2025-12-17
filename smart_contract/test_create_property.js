const { ethers } = require('hardhat');

async function main() {
  // Get the admin signer (should be the deployer)
  const [admin] = await ethers.getSigners();
  console.log('Admin address:', admin.address);

  // Get deployed contracts
  const factoryAddress = '0x31531Cea335B72E05493924af6d97Ef3C89C49bd';
  const Factory = await ethers.getContractFactory('PropertyFactory');
  const factory = Factory.attach(factoryAddress);

  // Check if admin has CREATOR_ROLE
  const creatorRole = await factory.CREATOR_ROLE();
  const hasCreatorRole = await factory.hasRole(creatorRole, admin.address);
  console.log('Admin has CREATOR_ROLE:', hasCreatorRole);

  if (!hasCreatorRole) {
    console.log('Granting CREATOR_ROLE to admin...');
    const tx = await factory.grantRole(creatorRole, admin.address);
    await tx.wait();
    console.log('CREATOR_ROLE granted');
  }

  // Test parameters (use small values to avoid gas issues)
  const owner = admin.address; // Use admin as owner for testing
  const name = 'Test Property';
  const symbol = 'TST';
  const dataHash = 'QmTestHash123456789012345678901234567890123456789012345678901234567890';
  const valuation = ethers.utils.parseEther('10'); // 10 ETH in wei (smaller for testing)
  const tokenSupply = ethers.utils.parseEther('1000'); // 1000 tokens in wei (smaller for testing)
  const tokenName = 'Test Property Token';
  const tokenSymbol = 'TPT';

  console.log('Creating property with params:');
  console.log('Owner:', owner);
  console.log('Name:', name);
  console.log('Symbol:', symbol);
  console.log('Data Hash:', dataHash);
  console.log('Valuation:', valuation.toString());
  console.log('Token Supply:', tokenSupply.toString());
  console.log('Token Name:', tokenName);
  console.log('Token Symbol:', tokenSymbol);

  try {
    const tx = await factory.createProperty(
      owner,
      name,
      symbol,
      dataHash,
      valuation,
      tokenSupply,
      tokenName,
      tokenSymbol
    );

    console.log('Transaction sent:', tx.hash);
    const receipt = await tx.wait();
    console.log('Transaction confirmed in block:', receipt.blockNumber);

    // Parse events
    const events = receipt.events;
    console.log('Events:', events.length);
    events.forEach((event, index) => {
      console.log(`Event ${index}:`, event.event, event.args);
    });

  } catch (error) {
    console.error('Error creating property:', error);
    if (error.error && error.error.message) {
      console.error('Revert reason:', error.error.message);
    }
  }
}

main().catch(console.error);
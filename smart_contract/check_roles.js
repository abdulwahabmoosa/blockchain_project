const { ethers } = require('ethers');
// Load environment variables from smart_contract/.env
require('dotenv').config();

async function checkRoles() {
  console.log('Available env vars:');
  console.log('SEPOLIA_RPC:', process.env.SEPOLIA_RPC ? 'set' : 'not set');
  console.log('PRIVATE_KEY:', process.env.PRIVATE_KEY ? 'set' : 'not set');
  console.log('PRIVATE_KEY_2:', process.env.PRIVATE_KEY_2 ? 'set' : 'not set');

  const provider = new ethers.providers.JsonRpcProvider(process.env.SEPOLIA_RPC);
  const privateKey = process.env.PRIVATE_KEY_2 || process.env.PRIVATE_KEY;

  if (!privateKey) {
    console.error('No private key found!');
    return;
  }

  console.log('Using private key:', privateKey.substring(0, 10) + '...');

  const signer = new ethers.Wallet(privateKey, provider);

  console.log('Signer address:', signer.address);
  console.log('Admin address from deployment:', '0x98E5a749E25C56e19C28008505DF75aFf4988049');

  // PropertyFactory contract (from registry)
  const factoryAddress = '0x9Aa0a2Ee2Cb91587484b0F7F2BE992E1F4484F1e';
  const factoryAbi = [
    "function hasRole(bytes32 role, address account) view returns (bool)",
    "function CREATOR_ROLE() view returns (bytes32)",
    "function DEFAULT_ADMIN_ROLE() view returns (bytes32)"
  ];

  const factory = new ethers.Contract(factoryAddress, factoryAbi, provider);

  const creatorRole = await factory.CREATOR_ROLE();
  const adminRole = await factory.DEFAULT_ADMIN_ROLE();

  console.log('\nRole checks:');
  console.log('CREATOR_ROLE hash:', creatorRole);
  console.log('DEFAULT_ADMIN_ROLE hash:', adminRole);

  const hasCreatorRole = await factory.hasRole(creatorRole, signer.address);
  const hasAdminRole = await factory.hasRole(adminRole, signer.address);

  console.log('Signer has CREATOR_ROLE:', hasCreatorRole);
  console.log('Signer has DEFAULT_ADMIN_ROLE:', hasAdminRole);

  // Also check the admin address
  const adminAddress = '0x98E5a749E25C56e19C28008505DF75aFf4988049';
  const adminHasCreatorRole = await factory.hasRole(creatorRole, adminAddress);
  const adminHasAdminRole = await factory.hasRole(adminRole, adminAddress);

  console.log('\nAdmin address role checks:');
  console.log('Admin has CREATOR_ROLE:', adminHasCreatorRole);
  console.log('Admin has DEFAULT_ADMIN_ROLE:', adminHasAdminRole);
}

checkRoles().catch(console.error);
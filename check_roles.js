const { ethers } = require('ethers');
// Load environment variables from smart_contract/.env
require('dotenv').config({ path: './smart_contract/.env' });

async function checkRoles() {
  const provider = new ethers.providers.JsonRpcProvider(process.env.SEPOLIA_RPC);
  const signer = new ethers.Wallet(process.env.PRIVATE_KEY_2 || process.env.PRIVATE_KEY, provider);

  console.log('Signer address:', signer.address);
  console.log('Admin address from deployment:', '0x98E5a749E25C56e19C28008505DF75aFf4988049');

  // PropertyFactory contract
  const factoryAddress = '0x31531Cea335B72E05493924af6d97Ef3C89C49bd';
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



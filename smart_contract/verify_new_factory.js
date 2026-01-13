const { ethers } = require('ethers');
const path = require('path');
require('dotenv').config({ path: path.join(__dirname, '..', '.env') });

async function verifyNewFactory() {
  console.log('🔍 Verifying new PropertyFactory roles...\n');

  let rpcUrl = process.env.SEPOLIA_RPC_BC || process.env.SEPOLIA_RPC;
  
  if (!rpcUrl) {
    console.error('❌ SEPOLIA_RPC or SEPOLIA_RPC_BC not found!');
    return;
  }

  let provider;
  if (rpcUrl.startsWith('wss://') || rpcUrl.startsWith('ws://')) {
    provider = new ethers.providers.WebSocketProvider(rpcUrl);
  } else {
    provider = new ethers.providers.JsonRpcProvider(rpcUrl);
  }

  // New factory address
  const factoryAddress = '0x4064A228951F39e3bFCC17B9A241a49558FaC4B6';
  const backendWallet = '0x98E5a749E25C56e19C28008505DF75aFf4988049';

  const factoryAbi = [
    "function hasRole(bytes32 role, address account) view returns (bool)",
    "function CREATOR_ROLE() view returns (bytes32)",
    "function DEFAULT_ADMIN_ROLE() view returns (bytes32)"
  ];

  const factory = new ethers.Contract(factoryAddress, factoryAbi, provider);

  const creatorRole = await factory.CREATOR_ROLE();
  const adminRole = await factory.DEFAULT_ADMIN_ROLE();

  console.log('📋 Role Hashes:');
  console.log('DEFAULT_ADMIN_ROLE:', adminRole);
  console.log('CREATOR_ROLE:', creatorRole);

  // Check backend wallet (PRIVATE_KEY_2)
  const hasAdmin = await factory.hasRole(adminRole, backendWallet);
  const hasCreator = await factory.hasRole(creatorRole, backendWallet);

  console.log(`\n📋 Backend Wallet (${backendWallet}):`);
  console.log(`   DEFAULT_ADMIN_ROLE: ${hasAdmin}`);
  console.log(`   CREATOR_ROLE: ${hasCreator}`);

  if (hasCreator) {
    console.log('\n✅ SUCCESS! Backend wallet has CREATOR_ROLE!');
    console.log('   The backend can now create properties on the blockchain.');
  } else {
    console.log('\n❌ Backend wallet does NOT have CREATOR_ROLE!');
    console.log('   This should not happen - the constructor should have granted it.');
  }

  if (provider && provider._websocket) {
    await provider.destroy();
  }
}

verifyNewFactory().catch((error) => {
  console.error('Fatal error:', error);
  process.exit(1);
});


const { ethers } = require('ethers');
const path = require('path');
require('dotenv').config({ path: path.join(__dirname, '..', '.env') });

async function findAdmin() {
  console.log('🔍 Finding which wallet has DEFAULT_ADMIN_ROLE...\n');

  // Use HTTP RPC if available
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

  const factoryAddress = '0x9Aa0a2Ee2Cb91587484b0F7F2BE992E1F4484F1e';
  const factoryAbi = [
    "function hasRole(bytes32 role, address account) view returns (bool)",
    "function CREATOR_ROLE() view returns (bytes32)",
    "function DEFAULT_ADMIN_ROLE() view returns (bytes32)",
    "function getRoleMemberCount(bytes32 role) view returns (uint256)",
    "function getRoleMember(bytes32 role, uint256 index) view returns (address)"
  ];

  const factory = new ethers.Contract(factoryAddress, factoryAbi, provider);

  const adminRole = await factory.DEFAULT_ADMIN_ROLE();
  const creatorRole = await factory.CREATOR_ROLE();

  console.log('📋 Role Hashes:');
  console.log('DEFAULT_ADMIN_ROLE:', adminRole);
  console.log('CREATOR_ROLE:', creatorRole);

  // Get all addresses with DEFAULT_ADMIN_ROLE
  try {
    const adminCount = await factory.getRoleMemberCount(adminRole);
    console.log(`\n👥 Found ${adminCount.toString()} address(es) with DEFAULT_ADMIN_ROLE:`);
    
    for (let i = 0; i < adminCount; i++) {
      const adminAddress = await factory.getRoleMember(adminRole, i);
      console.log(`  ${i + 1}. ${adminAddress}`);
      
      // Check if this address also has CREATOR_ROLE
      const hasCreator = await factory.hasRole(creatorRole, adminAddress);
      console.log(`     Has CREATOR_ROLE: ${hasCreator}`);
    }
  } catch (error) {
    console.log('⚠️ Could not enumerate role members (contract might not support it)');
    console.log('   Trying to check known wallets...\n');
    
    // Check known wallets
    const walletsToCheck = [
      { name: 'PRIVATE_KEY', key: process.env.PRIVATE_KEY },
      { name: 'PRIVATE_KEY_2', key: process.env.PRIVATE_KEY_2 },
    ];

    for (const wallet of walletsToCheck) {
      if (wallet.key) {
        const signer = new ethers.Wallet(wallet.key, provider);
        const address = signer.address;
        const hasAdmin = await factory.hasRole(adminRole, address);
        const hasCreator = await factory.hasRole(creatorRole, address);
        
        console.log(`📋 ${wallet.name} (${address}):`);
        console.log(`   DEFAULT_ADMIN_ROLE: ${hasAdmin}`);
        console.log(`   CREATOR_ROLE: ${hasCreator}\n`);
      }
    }
  }

  // Check backend wallet
  const backendWallet = '0x98E5a749E25C56e19C28008505DF75aFf4988049';
  const backendHasAdmin = await factory.hasRole(adminRole, backendWallet);
  const backendHasCreator = await factory.hasRole(creatorRole, backendWallet);
  
  console.log(`📋 Backend Wallet (${backendWallet}):`);
  console.log(`   DEFAULT_ADMIN_ROLE: ${backendHasAdmin}`);
  console.log(`   CREATOR_ROLE: ${backendHasCreator}`);

  if (provider && provider._websocket) {
    await provider.destroy();
  }
}

findAdmin().catch((error) => {
  console.error('Fatal error:', error);
  process.exit(1);
});

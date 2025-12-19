const { ethers } = require('ethers');
const path = require('path');
// Load environment variables from parent directory .env
require('dotenv').config({ path: path.join(__dirname, '..', '.env') });

async function grantCreatorRole() {
  console.log('🔧 Granting CREATOR_ROLE to backend wallet...\n');

  // Use HTTP RPC if available, otherwise try WebSocket
  let rpcUrl = process.env.SEPOLIA_RPC_BC || process.env.SEPOLIA_RPC;
  
  if (!rpcUrl) {
    console.error('❌ SEPOLIA_RPC or SEPOLIA_RPC_BC not found in .env file!');
    return;
  }

  // If it's a WebSocket URL, use WebSocketProvider, otherwise use JsonRpcProvider
  let provider;
  if (rpcUrl.startsWith('wss://') || rpcUrl.startsWith('ws://')) {
    console.log('📡 Using WebSocket provider...');
    provider = new ethers.providers.WebSocketProvider(rpcUrl);
  } else {
    console.log('📡 Using HTTP provider...');
    provider = new ethers.providers.JsonRpcProvider(rpcUrl);
  }
  
  // Use PRIVATE_KEY_2 (the deployer wallet that has DEFAULT_ADMIN_ROLE)
  const deployerPrivateKey = process.env.PRIVATE_KEY_2;
  
  if (!deployerPrivateKey) {
    console.error('❌ PRIVATE_KEY_2 not found! This should be the deployer wallet with DEFAULT_ADMIN_ROLE.');
    return;
  }

  const deployerSigner = new ethers.Wallet(deployerPrivateKey, provider);
  console.log('👤 Deployer address:', deployerSigner.address);

  // Backend wallet that needs CREATOR_ROLE
  const backendWallet = '0x98E5a749E25C56e19C28008505DF75aFf4988049';
  console.log('🎯 Backend wallet to grant role:', backendWallet);

  // PropertyFactory contract address
  const factoryAddress = '0x9Aa0a2Ee2Cb91587484b0F7F2BE992E1F4484F1e';
  const factoryAbi = [
    "function hasRole(bytes32 role, address account) view returns (bool)",
    "function CREATOR_ROLE() view returns (bytes32)",
    "function DEFAULT_ADMIN_ROLE() view returns (bytes32)",
    "function grantRole(bytes32 role, address account)"
  ];

  const factory = new ethers.Contract(factoryAddress, factoryAbi, deployerSigner);

  // Check current roles
  const creatorRole = await factory.CREATOR_ROLE();
  const adminRole = await factory.DEFAULT_ADMIN_ROLE();

  console.log('\n📋 Checking current roles...');
  const deployerHasAdminRole = await factory.hasRole(adminRole, deployerSigner.address);
  const backendHasCreatorRole = await factory.hasRole(creatorRole, backendWallet);

  console.log('Deployer has DEFAULT_ADMIN_ROLE:', deployerHasAdminRole);
  console.log('Backend wallet has CREATOR_ROLE:', backendHasCreatorRole);

  if (!deployerHasAdminRole) {
    console.error('❌ Deployer wallet does not have DEFAULT_ADMIN_ROLE! Cannot grant role.');
    return;
  }

  if (backendHasCreatorRole) {
    console.log('✅ Backend wallet already has CREATOR_ROLE! No action needed.');
    return;
  }

  console.log('\n🚀 Granting CREATOR_ROLE to backend wallet...');
  try {
    const tx = await factory.grantRole(creatorRole, backendWallet);
    console.log('📝 Transaction submitted:', tx.hash);
    console.log('⏳ Waiting for confirmation...');
    
    const receipt = await tx.wait();
    console.log('✅ Transaction confirmed in block:', receipt.blockNumber);
    console.log('✅ CREATOR_ROLE granted successfully!');
    
    // Verify
    const verified = await factory.hasRole(creatorRole, backendWallet);
    console.log('\n🔍 Verification - Backend wallet has CREATOR_ROLE:', verified);
    
    // Close WebSocket connection if used
    if (provider && provider._websocket) {
      await provider.destroy();
    }
    
  } catch (error) {
    console.error('❌ Failed to grant role:', error.message);
    if (error.reason) {
      console.error('Reason:', error.reason);
    }
    // Close WebSocket connection on error
    if (provider && provider._websocket) {
      await provider.destroy();
    }
  }
}

grantCreatorRole().catch((error) => {
  console.error('Fatal error:', error);
  process.exit(1);
});

const { ethers } = require('ethers');
const path = require('path');
require('dotenv').config({ path: path.join(__dirname, '..', '.env') });

async function checkDeployment() {
  console.log('🔍 Checking PropertyFactory deployment...\n');

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
  
  // Get contract creation transaction
  const code = await provider.getCode(factoryAddress);
  if (code === '0x') {
    console.error('❌ No contract found at this address!');
    return;
  }

  console.log('✅ Contract exists at:', factoryAddress);
  console.log('📦 Contract bytecode length:', code.length);

  // Try to get the deployment transaction
  // We'll need to check Etherscan or find it in deployment logs
  console.log('\n📋 To find the admin address:');
  console.log('1. Go to: https://sepolia.etherscan.io/address/' + factoryAddress);
  console.log('2. Click on "Contract" tab');
  console.log('3. Click on "Read Contract"');
  console.log('4. Find "DEFAULT_ADMIN_ROLE" and "getRoleMember" functions');
  console.log('5. Or check the deployment transaction to see what admin address was used\n');

  // Try to read the registry address from the contract
  const factoryAbi = [
    "function registry() view returns (address)",
    "function DEFAULT_ADMIN_ROLE() view returns (bytes32)",
    "function CREATOR_ROLE() view returns (bytes32)",
    "function hasRole(bytes32 role, address account) view returns (bool)"
  ];

  try {
    const factory = new ethers.Contract(factoryAddress, factoryAbi, provider);
    const registryAddr = await factory.registry();
    console.log('📋 Registry address:', registryAddr);
    
    // Check if DEFAULT_ADMIN_ROLE is actually 0x00... (which would be wrong)
    const adminRoleHash = await factory.DEFAULT_ADMIN_ROLE();
    console.log('📋 DEFAULT_ADMIN_ROLE hash:', adminRoleHash);
    
    if (adminRoleHash === '0x0000000000000000000000000000000000000000000000000000000000000000') {
      console.log('⚠️ WARNING: DEFAULT_ADMIN_ROLE is 0x00... This is incorrect!');
      console.log('   The contract may not be properly initialized.');
      console.log('   You may need to redeploy PropertyFactory with the correct admin address.');
    }
  } catch (error) {
    console.error('❌ Error reading contract:', error.message);
  }

  if (provider && provider._websocket) {
    await provider.destroy();
  }
}

checkDeployment().catch((error) => {
  console.error('Fatal error:', error);
  process.exit(1);
});


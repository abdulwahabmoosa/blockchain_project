const { ethers } = require('ethers');
const path = require('path');
require('dotenv').config({ path: path.join(__dirname, '..', '.env') });

/**
 * Script to update PlatformRegistry with new PropertyFactory address
 * This needs to be run by the wallet that has admin role in PlatformRegistry
 */

async function updateRegistry() {
  console.log('🔧 Updating PlatformRegistry with new PropertyFactory address...\n');

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

  const registryAddress = process.env.REGISTRY || '0x113FdC450314d59eAbEEfA2ff9F232DD7b08ad38';
  const newFactoryAddress = '0x4064A228951F39e3bFCC17B9A241a49558FaC4B6';

  console.log('Registry:', registryAddress);
  console.log('New Factory:', newFactoryAddress);

  // Try with PRIVATE_KEY first (the original deployer)
  const privateKey = process.env.PRIVATE_KEY;
  
  if (!privateKey) {
    console.error('❌ PRIVATE_KEY not found!');
    console.log('\n📋 Alternative: Update via Etherscan');
    console.log('1. Go to: https://sepolia.etherscan.io/address/' + registryAddress);
    console.log('2. Click "Contract" -> "Write Contract"');
    console.log('3. Connect the wallet that deployed PlatformRegistry');
    console.log('4. Find "setPropertyFactory" function');
    console.log('5. Enter new factory address:', newFactoryAddress);
    console.log('6. Click "Write" and confirm transaction');
    return;
  }

  const signer = new ethers.Wallet(privateKey, provider);
  console.log('Using wallet:', signer.address);

  const registryAbi = [
    "function setPropertyFactory(address factory)",
    "function getPropertyFactory() view returns (address)",
    "function hasRole(bytes32 role, address account) view returns (bool)",
    "function DEFAULT_ADMIN_ROLE() view returns (bytes32)"
  ];

  const registry = new ethers.Contract(registryAddress, registryAbi, signer);

  // Check if signer has admin role
  try {
    const adminRole = await registry.DEFAULT_ADMIN_ROLE();
    const hasAdmin = await registry.hasRole(adminRole, signer.address);
    
    if (!hasAdmin) {
      console.error('❌ Wallet does not have DEFAULT_ADMIN_ROLE in PlatformRegistry!');
      console.log('\n📋 Update via Etherscan instead:');
      console.log('1. Go to: https://sepolia.etherscan.io/address/' + registryAddress);
      console.log('2. Click "Contract" -> "Write Contract"');
      console.log('3. Connect the wallet that has admin role');
      console.log('4. Find "setPropertyFactory" function');
      console.log('5. Enter:', newFactoryAddress);
      console.log('6. Click "Write" and confirm');
      return;
    }

    console.log('✅ Wallet has admin role, updating registry...');
    
    // Get current factory
    const currentFactory = await registry.getPropertyFactory();
    console.log('Current factory:', currentFactory);
    
    // Update to new factory
    const tx = await registry.setPropertyFactory(newFactoryAddress);
    console.log('📝 Transaction submitted:', tx.hash);
    console.log('⏳ Waiting for confirmation...');
    
    const receipt = await tx.wait();
    console.log('✅ Transaction confirmed in block:', receipt.blockNumber);
    
    // Verify
    const updatedFactory = await registry.getPropertyFactory();
    console.log('✅ Registry updated! New factory:', updatedFactory);
    
  } catch (error) {
    console.error('❌ Error:', error.message);
    if (error.reason) {
      console.error('Reason:', error.reason);
    }
  }

  if (provider && provider._websocket) {
    await provider.destroy();
  }
}

updateRegistry().catch((error) => {
  console.error('Fatal error:', error);
  process.exit(1);
});


const { ethers } = require('ethers');
const path = require('path');
require('dotenv').config({ path: path.join(__dirname, '..', '.env') });

/**
 * Script to grant SNAPSHOT_ROLE to RevenueDistribution for existing properties
 * using the PropertyFactory's grantSnapshotRoleToRevenue() function
 * 
 * This script calls PropertyFactory.grantSnapshotRoleToRevenue() for each property token.
 * PropertyFactory must be redeployed with the new grantSnapshotRoleToRevenue() function first.
 */

async function grantSnapshotForExistingProperties() {
  console.log('🔧 Granting SNAPSHOT_ROLE to RevenueDistribution for existing properties...\n');

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

  // Contract addresses
  const propertyFactoryAddress = '0x4064A228951F39e3bFCC17B9A241a49558FaC4B6'; // Update after redeployment
  
  // PropertyToken addresses - ADD YOUR PROPERTY TOKEN ADDRESSES HERE
  // Query your database: SELECT onchain_token_address FROM properties;
  const propertyTokens = [
    // Example: '0x805fdBbD19974e3145861836A2c71CBbff898D0A',
    // Add all your property token addresses here
  ];

  if (propertyTokens.length === 0) {
    console.error('❌ No PropertyToken addresses provided!');
    console.error('   Please add PropertyToken addresses to the propertyTokens array in this script.');
    console.error('   Query your database: SELECT onchain_token_address FROM properties;');
    return;
  }

  // PropertyFactory ABI (minimal)
  const propertyFactoryABI = [
    "function grantSnapshotRoleToRevenue(address tokenAddress)"
  ];

  const adminPrivateKey = process.env.PRIVATE_KEY || process.env.PRIVATE_KEY_2;
  if (!adminPrivateKey) {
    console.error('❌ PRIVATE_KEY or PRIVATE_KEY_2 not found in .env!');
    console.error('   You need the admin private key that has DEFAULT_ADMIN_ROLE on PropertyFactory.');
    return;
  }

  const adminSigner = new ethers.Wallet(adminPrivateKey, provider);
  console.log(`👤 Admin wallet: ${adminSigner.address}\n`);

  const factoryContract = new ethers.Contract(propertyFactoryAddress, propertyFactoryABI, adminSigner);

  // For each property token, call PropertyFactory.grantSnapshotRoleToRevenue()
  for (const tokenAddress of propertyTokens) {
    console.log(`📋 Processing PropertyToken: ${tokenAddress}`);
    
    try {
      console.log(`   🚀 Calling PropertyFactory.grantSnapshotRoleToRevenue(${tokenAddress})...`);
      const tx = await factoryContract.grantSnapshotRoleToRevenue(tokenAddress);
      console.log(`   📝 Transaction: ${tx.hash}`);
      console.log(`   ⏳ Waiting for confirmation...`);
      
      const receipt = await tx.wait();
      console.log(`   ✅ Confirmed in block: ${receipt.blockNumber}\n`);
      
    } catch (error) {
      console.error(`   ❌ Error: ${error.message}`);
      if (error.reason) {
        console.error(`   Reason: ${error.reason}`);
      }
      console.error(`   📝 Make sure:\n`);
      console.error(`      1. PropertyFactory is redeployed with grantSnapshotRoleToRevenue() function`);
      console.error(`      2. Admin wallet has DEFAULT_ADMIN_ROLE on PropertyFactory`);
      console.error(`      3. RevenueDistribution is set in PlatformRegistry\n`);
    }
  }

  console.log('📋 Summary:');
  console.log('   After running this script, RevenueDistribution should be able to create snapshots for existing properties.');
  console.log('   New properties will automatically have the role granted by PropertyFactory.\n');

  if (provider && provider._websocket) {
    await provider.destroy();
  }
}

grantSnapshotForExistingProperties().catch((error) => {
  console.error('Fatal error:', error);
  process.exit(1);
});


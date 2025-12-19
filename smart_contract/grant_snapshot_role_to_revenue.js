const { ethers } = require('ethers');
const path = require('path');
require('dotenv').config({ path: path.join(__dirname, '..', '.env') });

/**
 * Script to grant SNAPSHOT_ROLE to RevenueDistribution contract on all PropertyToken contracts
 * 
 * This is required because RevenueDistribution needs to call snapshotNow() on PropertyTokens
 * to create snapshots before distributing revenue.
 */

async function grantSnapshotRoleToRevenue() {
  console.log('🔧 Granting SNAPSHOT_ROLE to RevenueDistribution on all PropertyTokens...\n');

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
  const revenueDistributionAddress = '0xA08e097D421bD4a23b9db42a7356945a373a4666';
  const propertyFactoryAddress = '0x4064A228951F39e3bFCC17B9A241a49558FaC4B6';
  
  // IMPORTANT: PropertyTokens are created with PropertyFactory as their admin (address(this))
  // So PropertyFactory is the DEFAULT_ADMIN_ROLE holder, not a wallet.
  // However, since contracts can't sign transactions, we need to check:
  // 1. If PropertyFactory has DEFAULT_ADMIN_ROLE, we need the wallet that can control PropertyFactory
  // 2. OR, we need to modify PropertyFactory to have a function that grants roles
  //
  // For now, try using the wallet that deployed PropertyFactory (should have DEFAULT_ADMIN_ROLE on PropertyFactory)
  // But PropertyTokens have PropertyFactory as admin, so this might not work directly.
  //
  // PropertyToken addresses - ADD YOUR PROPERTY TOKEN ADDRESSES HERE
  // You can find these from:
  // 1. Your database (properties table, OnchainTokenAddress column)  
  // 2. Query your database: SELECT onchain_token_address FROM properties WHERE id = 'ddc9a63c-adf2-4e68-a09d-101b6097fd00'
  // 3. Etherscan - search for PropertyRegistered events from PropertyFactory at 0x4064A228951F39e3bFCC17B9A241a49558FaC4B6
  const propertyTokens = [
    // Example: '0x805fdBbD19974e3145861836A2c71CBbff898D0A',
    // Add your property token addresses here - one per line
    // For the property in your error (ddc9a63c-adf2-4e68-a09d-101b6097fd00),
    // you need to find its OnchainTokenAddress from your database
  ];

  // PropertyToken ABI (minimal - just what we need)
  const propertyTokenABI = [
    "function SNAPSHOT_ROLE() view returns (bytes32)",
    "function hasRole(bytes32 role, address account) view returns (bool)",
    "function grantRole(bytes32 role, address account)"
  ];

  // PropertyFactory ABI to get admin address (or we can use PRIVATE_KEY from env)
  // Actually, PropertyFactory is the admin of PropertyTokens, so we can use PRIVATE_KEY
  // which should be the PropertyFactory deployer/admin

  const adminPrivateKey = process.env.PRIVATE_KEY || process.env.PRIVATE_KEY_2;
  if (!adminPrivateKey) {
    console.error('❌ PRIVATE_KEY or PRIVATE_KEY_2 not found in .env!');
    console.error('   You need the admin private key to grant roles.');
    return;
  }

  const adminSigner = new ethers.Wallet(adminPrivateKey, provider);
  console.log(`👤 Admin wallet: ${adminSigner.address}\n`);
  console.log(`⚠️  NOTE: PropertyTokens have PropertyFactory (${propertyFactoryAddress}) as their admin.`);
  console.log(`   If this script fails, PropertyFactory needs to grant the role, not a wallet.\n`);

  if (propertyTokens.length === 0) {
    console.error('❌ No PropertyToken addresses provided!');
    console.error('   Please add PropertyToken addresses to the propertyTokens array in this script.');
    console.error('   You can find them in your database (properties table, OnchainTokenAddress column).');
    return;
  }

  // For each property token, grant SNAPSHOT_ROLE to RevenueDistribution
  for (const tokenAddress of propertyTokens) {
    console.log(`📋 Processing PropertyToken: ${tokenAddress}`);
    
    const tokenContract = new ethers.Contract(tokenAddress, propertyTokenABI, adminSigner);
    
    try {
      // Check who the admin is (should be PropertyFactory)
      const DEFAULT_ADMIN_ROLE = await tokenContract.DEFAULT_ADMIN_ROLE();
      const adminHasRole = await tokenContract.hasRole(DEFAULT_ADMIN_ROLE, adminSigner.address);
      const factoryHasRole = await tokenContract.hasRole(DEFAULT_ADMIN_ROLE, propertyFactoryAddress);
      
      console.log(`   Admin wallet has DEFAULT_ADMIN_ROLE: ${adminHasRole}`);
      console.log(`   PropertyFactory has DEFAULT_ADMIN_ROLE: ${factoryHasRole}`);
      
      if (!adminHasRole && factoryHasRole) {
        console.error(`   ⚠️  WARNING: PropertyFactory is the admin, not a wallet!`);
        console.error(`   This script may not work. You may need to modify PropertyFactory to grant roles.`);
        console.error(`   OR use a wallet that has DEFAULT_ADMIN_ROLE on PropertyFactory and can control it.\n`);
        continue;
      }
      
      // Get SNAPSHOT_ROLE constant  
      const snapshotRole = await tokenContract.SNAPSHOT_ROLE();
      console.log(`   SNAPSHOT_ROLE: ${snapshotRole}`);
      
      // Check if RevenueDistribution already has the role
      const hasRole = await tokenContract.hasRole(snapshotRole, revenueDistributionAddress);
      console.log(`   RevenueDistribution has role: ${hasRole}`);
      
      if (hasRole) {
        console.log(`   ✅ RevenueDistribution already has SNAPSHOT_ROLE!`);
        continue;
      }
      
      // Grant the role
      console.log(`   🚀 Granting SNAPSHOT_ROLE to RevenueDistribution (${revenueDistributionAddress})...`);
      const tx = await tokenContract.grantRole(snapshotRole, revenueDistributionAddress);
      console.log(`   📝 Transaction: ${tx.hash}`);
      console.log(`   ⏳ Waiting for confirmation...`);
      
      const receipt = await tx.wait();
      console.log(`   ✅ Confirmed in block: ${receipt.blockNumber}`);
      
      // Verify
      const verified = await tokenContract.hasRole(snapshotRole, revenueDistributionAddress);
      console.log(`   ✅ Verification: ${verified ? 'SUCCESS' : 'FAILED'}\n`);
      
    } catch (error) {
      console.error(`   ❌ Error: ${error.message}`);
      if (error.reason) {
        console.error(`   Reason: ${error.reason}`);
      }
      if (error.message.includes('AccessControl') || error.message.includes('not admin')) {
        console.error(`   ⚠️  This PropertyToken's admin is PropertyFactory (a contract), not a wallet.`);
        console.error(`   You need to modify PropertyFactory to grant this role, or add a helper function.\n`);
      } else {
        console.error(`   📝 Make sure the admin wallet has DEFAULT_ADMIN_ROLE on this PropertyToken\n`);
      }
    }
  }

  console.log('\n📋 Summary:');
  console.log('   After running this script, RevenueDistribution should be able to create snapshots.');
  console.log('   If you add new properties, you may need to run this script again for new PropertyTokens.');
  console.log('\n⚠️  IMPORTANT: If the script failed because PropertyFactory is the admin:');
  console.log('   1. The PropertyFactory contract has been modified to auto-grant the role for new properties');
  console.log('   2. For EXISTING properties, you need to:');
  console.log('      a. Redeploy PropertyFactory with the new grantSnapshotRoleToRevenue() function');
  console.log('      b. Call grantSnapshotRoleToRevenue(tokenAddress) on PropertyFactory for each existing token');
  console.log('      c. OR manually grant the role using the wallet that controls PropertyFactory');
  console.log('\n💡 Alternative: Check if your PropertyFactory already has a way to grant roles on tokens.');

  if (provider && provider._websocket) {
    await provider.destroy();
  }
}

grantSnapshotRoleToRevenue().catch((error) => {
  console.error('Fatal error:', error);
  process.exit(1);
});

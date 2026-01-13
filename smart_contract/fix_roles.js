const { ethers } = require('ethers');
const path = require('path');
require('dotenv').config({ path: path.join(__dirname, '..', '.env') });

/**
 * This script attempts to fix the PropertyFactory roles issue.
 * Since no wallet has DEFAULT_ADMIN_ROLE, we need to either:
 * 1. Find the original deployer wallet
 * 2. Or the contract needs to be redeployed
 * 
 * First, let's check Etherscan to find the deployment transaction.
 */

async function fixRoles() {
  console.log('🔧 Attempting to fix PropertyFactory roles...\n');

  const factoryAddress = '0x9Aa0a2Ee2Cb91587484b0F7F2BE992E1F4484F1e';
  const backendWallet = '0x98E5a749E25C56e19C28008505DF75aFf4988049';

  console.log('📋 Instructions:');
  console.log('1. Go to: https://sepolia.etherscan.io/address/' + factoryAddress);
  console.log('2. Click on "Contract" tab');
  console.log('3. Click on "Read Contract"');
  console.log('4. Scroll to "getRoleMember" function');
  console.log('5. Enter role: 0x0000000000000000000000000000000000000000000000000000000000000000');
  console.log('6. Enter index: 0');
  console.log('7. This will show you which address has DEFAULT_ADMIN_ROLE\n');

  console.log('OR check the deployment transaction:');
  console.log('1. Go to: https://sepolia.etherscan.io/address/' + factoryAddress);
  console.log('2. Click on "Contract" tab');
  console.log('3. Click on "Contract Creation Code"');
  console.log('4. Find the transaction that created this contract');
  console.log('5. Check the "From" address - that is likely the admin\n');

  console.log('📝 Alternative: If you have access to the deployer wallet,');
  console.log('   you can grant DEFAULT_ADMIN_ROLE to PRIVATE_KEY_2, then');
  console.log('   use PRIVATE_KEY_2 to grant CREATOR_ROLE to the backend wallet.\n');

  console.log('⚠️  If no wallet has DEFAULT_ADMIN_ROLE, the contract is locked.');
  console.log('   You may need to redeploy PropertyFactory with the correct admin address.');
}

fixRoles().catch((error) => {
  console.error('Fatal error:', error);
  process.exit(1);
});


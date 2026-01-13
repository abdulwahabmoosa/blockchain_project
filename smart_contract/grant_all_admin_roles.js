const { ethers } = require('ethers');
const path = require('path');
require('dotenv').config({ path: path.join(__dirname, '..', '.env') });

/**
 * Script to grant PRIVATE_KEY_2 admin roles in all contracts
 * This will try to grant DEFAULT_ADMIN_ROLE to PRIVATE_KEY_2 in:
 * - PlatformRegistry
 * - ApprovalService  
 * - RevenueDistribution
 * - PropertyFactory (already has it, but will verify)
 */

async function grantAllAdminRoles() {
  console.log('🔧 Granting PRIVATE_KEY_2 admin roles in all contracts...\n');

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

  const targetWallet = '0x98E5a749E25C56e19C28008505DF75aFf4988049'; // PRIVATE_KEY_2
  const contracts = [
    {
      name: 'PlatformRegistry',
      address: process.env.REGISTRY || '0x113FdC450314d59eAbEEfA2ff9F232DD7b08ad38',
      adminKey: process.env.PRIVATE_KEY, // Try PRIVATE_KEY first (original deployer)
    },
    {
      name: 'ApprovalService',
      address: '0xB88F19F1074E6d7c23B6951360c00AEd1C4Ac635',
      adminKey: process.env.PRIVATE_KEY,
    },
    {
      name: 'RevenueDistribution',
      address: process.env.REVENUE || '0xD7ebec79B2e8D218F59A3777fa615eae762Bb7f6',
      adminKey: process.env.PRIVATE_KEY,
    },
    {
      name: 'PropertyFactory',
      address: '0x4064A228951F39e3bFCC17B9A241a49558FaC4B6', // New factory
      adminKey: process.env.PRIVATE_KEY_2, // Already has PRIVATE_KEY_2 as admin
    },
  ];

  const abi = [
    "function hasRole(bytes32 role, address account) view returns (bool)",
    "function DEFAULT_ADMIN_ROLE() view returns (bytes32)",
    "function grantRole(bytes32 role, address account)"
  ];

  for (const contract of contracts) {
    console.log(`\n📋 Processing ${contract.name}...`);
    console.log(`   Address: ${contract.address}`);

    if (!contract.adminKey) {
      console.log(`   ⚠️  No admin key available for this contract`);
      console.log(`   📝 Manual update needed via Etherscan`);
      continue;
    }

    const adminSigner = new ethers.Wallet(contract.adminKey, provider);
    const contractInstance = new ethers.Contract(contract.address, abi, adminSigner);

    try {
      // Check if admin signer has admin role
      const adminRole = await contractInstance.DEFAULT_ADMIN_ROLE();
      const adminHasRole = await contractInstance.hasRole(adminRole, adminSigner.address);
      const targetHasRole = await contractInstance.hasRole(adminRole, targetWallet);

      console.log(`   Admin wallet: ${adminSigner.address}`);
      console.log(`   Admin has role: ${adminHasRole}`);
      console.log(`   Target has role: ${targetHasRole}`);

      if (!adminHasRole) {
        console.log(`   ❌ Admin wallet does not have DEFAULT_ADMIN_ROLE!`);
        console.log(`   📝 Need to find the correct admin wallet or redeploy contract`);
        continue;
      }

      if (targetHasRole) {
        console.log(`   ✅ Target wallet already has admin role!`);
        continue;
      }

      // Grant role
      console.log(`   🚀 Granting DEFAULT_ADMIN_ROLE to ${targetWallet}...`);
      const tx = await contractInstance.grantRole(adminRole, targetWallet);
      console.log(`   📝 Transaction: ${tx.hash}`);
      console.log(`   ⏳ Waiting for confirmation...`);
      
      const receipt = await tx.wait();
      console.log(`   ✅ Confirmed in block: ${receipt.blockNumber}`);

      // Verify
      const verified = await contractInstance.hasRole(adminRole, targetWallet);
      console.log(`   ✅ Verification: ${verified ? 'SUCCESS' : 'FAILED'}`);

    } catch (error) {
      console.error(`   ❌ Error: ${error.message}`);
      if (error.reason) {
        console.error(`   Reason: ${error.reason}`);
      }
      console.log(`   📝 May need to update manually via Etherscan`);
    }
  }

  console.log('\n📋 Summary:');
  console.log('   After running this script, verify all roles with: node find_admin.js');
  console.log('   If some contracts failed, you may need to redeploy them with PRIVATE_KEY_2 as admin.');

  if (provider && provider._websocket) {
    await provider.destroy();
  }
}

grantAllAdminRoles().catch((error) => {
  console.error('Fatal error:', error);
  process.exit(1);
});


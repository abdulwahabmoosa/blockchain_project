const { ethers } = require('ethers');
const path = require('path');
require('dotenv').config({ path: path.join(__dirname, '..', '.env') });

/**
 * Comprehensive script to verify and grant all necessary roles to PRIVATE_KEY_2
 * This checks and grants:
 * - PropertyFactory: CREATOR_ROLE (needed to create properties)
 * - PlatformRegistry: DEFAULT_ADMIN_ROLE (needed to update contract addresses)
 * - ApprovalService: ADMIN_ROLE and DEFAULT_ADMIN_ROLE (needed to approve users)
 * - RevenueDistribution: DISTRIBUTOR_ROLE and DEFAULT_ADMIN_ROLE (needed to distribute revenue)
 */

async function verifyAndGrantAllRoles() {
  console.log('🔧 Verifying and granting all roles to PRIVATE_KEY_2...\n');

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
  
  // Contract addresses from the updated contract-addresses.json
  const contracts = [
    {
      name: 'PropertyFactory',
      address: '0x4064A228951F39e3bFCC17B9A241a49558FaC4B6',
      roles: [
        { name: 'CREATOR_ROLE', method: 'CREATOR_ROLE', required: true },
        { name: 'DEFAULT_ADMIN_ROLE', method: 'DEFAULT_ADMIN_ROLE', required: true }
      ],
      adminKey: process.env.PRIVATE_KEY_2, // PRIVATE_KEY_2 is already admin
    },
    {
      name: 'PlatformRegistry',
      address: '0x42886E2091615c831252bA88C640641d5cfab43F', // New deployment
      roles: [
        { name: 'DEFAULT_ADMIN_ROLE', method: 'DEFAULT_ADMIN_ROLE', required: true }
      ],
      adminKey: process.env.PRIVATE_KEY_2, // Should be admin from new deployment
    },
    {
      name: 'ApprovalService',
      address: '0xB88F19F1074E6d7c23B6951360c00AEd1C4Ac635',
      roles: [
        { name: 'ADMIN_ROLE', method: 'ADMIN_ROLE', required: true },
        { name: 'DEFAULT_ADMIN_ROLE', method: 'DEFAULT_ADMIN_ROLE', required: true }
      ],
      adminKey: process.env.PRIVATE_KEY, // Try PRIVATE_KEY first
    },
    {
      name: 'RevenueDistribution',
      address: '0xA08e097D421bD4a23b9db42a7356945a373a4666', // New deployment
      roles: [
        { name: 'DISTRIBUTOR_ROLE', method: 'DISTRIBUTOR_ROLE', required: true },
        { name: 'DEFAULT_ADMIN_ROLE', method: 'DEFAULT_ADMIN_ROLE', required: true }
      ],
      adminKey: process.env.PRIVATE_KEY_2, // Should be admin from new deployment
    },
  ];

  const baseAbi = [
    "function hasRole(bytes32 role, address account) view returns (bool)",
    "function DEFAULT_ADMIN_ROLE() view returns (bytes32)",
    "function grantRole(bytes32 role, address account)"
  ];

  const results = {
    verified: [],
    granted: [],
    failed: []
  };

  for (const contract of contracts) {
    console.log(`\n${'='.repeat(60)}`);
    console.log(`📋 ${contract.name}`);
    console.log(`   Address: ${contract.address}`);

    if (!contract.adminKey) {
      console.log(`   ⚠️  No admin key available`);
      results.failed.push({ contract: contract.name, reason: 'No admin key' });
      continue;
    }

    const adminSigner = new ethers.Wallet(contract.adminKey, provider);
    console.log(`   Admin wallet: ${adminSigner.address}`);

    // Build ABI with role-specific methods
    const roleMethods = contract.roles.map(r => `function ${r.method}() view returns (bytes32)`);
    const abi = [...baseAbi, ...roleMethods];

    const contractInstance = new ethers.Contract(contract.address, abi, adminSigner);

    try {
      // Check each required role
      for (const role of contract.roles) {
        const roleHash = await contractInstance[role.method]();
        const adminHasRole = await contractInstance.hasRole(roleHash, adminSigner.address);
        const targetHasRole = await contractInstance.hasRole(roleHash, targetWallet);

        console.log(`\n   🔑 ${role.name}:`);
        console.log(`      Admin has role: ${adminHasRole}`);
        console.log(`      Target has role: ${targetHasRole}`);

        if (!adminHasRole) {
          console.log(`      ❌ Admin wallet does not have ${role.name}!`);
          console.log(`      📝 Cannot grant role - need correct admin wallet`);
          results.failed.push({ 
            contract: contract.name, 
            role: role.name, 
            reason: 'Admin wallet does not have role' 
          });
          continue;
        }

        if (targetHasRole) {
          console.log(`      ✅ Target already has ${role.name}!`);
          results.verified.push({ contract: contract.name, role: role.name });
          continue;
        }

        // Grant role
        console.log(`      🚀 Granting ${role.name}...`);
        try {
          const tx = await contractInstance.grantRole(roleHash, targetWallet);
          console.log(`      📝 Transaction: ${tx.hash}`);
          console.log(`      ⏳ Waiting for confirmation...`);
          
          const receipt = await tx.wait();
          console.log(`      ✅ Confirmed in block: ${receipt.blockNumber}`);

          // Verify
          const verified = await contractInstance.hasRole(roleHash, targetWallet);
          if (verified) {
            console.log(`      ✅ Verification: SUCCESS`);
            results.granted.push({ contract: contract.name, role: role.name, txHash: tx.hash });
          } else {
            console.log(`      ❌ Verification: FAILED`);
            results.failed.push({ contract: contract.name, role: role.name, reason: 'Verification failed' });
          }
        } catch (grantError) {
          console.error(`      ❌ Failed to grant: ${grantError.message}`);
          results.failed.push({ 
            contract: contract.name, 
            role: role.name, 
            reason: grantError.message 
          });
        }
      }
    } catch (error) {
      console.error(`   ❌ Error: ${error.message}`);
      results.failed.push({ contract: contract.name, reason: error.message });
    }
  }

  // Summary
  console.log(`\n${'='.repeat(60)}`);
  console.log('📊 SUMMARY');
  console.log(`${'='.repeat(60)}`);
  console.log(`✅ Verified (already has role): ${results.verified.length}`);
  results.verified.forEach(r => console.log(`   - ${r.contract}: ${r.role}`));
  
  console.log(`\n🚀 Granted (just granted): ${results.granted.length}`);
  results.granted.forEach(r => console.log(`   - ${r.contract}: ${r.role} (${r.txHash})`));
  
  console.log(`\n❌ Failed: ${results.failed.length}`);
  results.failed.forEach(r => {
    console.log(`   - ${r.contract}${r.role ? ': ' + r.role : ''} - ${r.reason}`);
  });

  if (results.failed.length > 0) {
    console.log(`\n📝 For failed grants, you may need to:`);
    console.log(`   1. Find the correct admin wallet for those contracts`);
    console.log(`   2. Or redeploy those contracts with PRIVATE_KEY_2 as admin`);
  }

  if (provider && provider._websocket) {
    await provider.destroy();
  }
}

verifyAndGrantAllRoles().catch((error) => {
  console.error('Fatal error:', error);
  process.exit(1);
});

const { ethers } = require('ethers');
const path = require('path');
require('dotenv').config({ path: path.join(__dirname, '..', '.env') });

/**
 * Script to update PlatformRegistry with all contract addresses
 * This should be run after redeploying contracts to ensure registry points to correct addresses
 */

async function setupRegistry() {
  console.log('🔧 Setting up PlatformRegistry with contract addresses...\n');

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

  // Contract addresses from contract-addresses.json
  const registryAddress = '0x42886E2091615c831252bA88C640641d5cfab43F'; // New PlatformRegistry
  const factoryAddress = '0x4064A228951F39e3bFCC17B9A241a49558FaC4B6';
  const approvalAddress = '0xB88F19F1074E6d7c23B6951360c00AEd1C4Ac635';
  const revenueAddress = '0xA08e097D421bD4a23b9db42a7356945a373a4666'; // New RevenueDistribution

  // Use PRIVATE_KEY_2 (should be admin in new registry)
  const adminKey = process.env.PRIVATE_KEY_2;
  
  if (!adminKey) {
    console.error('❌ PRIVATE_KEY_2 not found!');
    return;
  }

  const adminSigner = new ethers.Wallet(adminKey, provider);
  console.log('Using admin wallet:', adminSigner.address);

  const registryAbi = [
    "function setPropertyFactory(address factory)",
    "function setApprovalService(address approval)",
    "function setRevenueDistribution(address revenue)",
    "function getPropertyFactory() view returns (address)",
    "function getApprovalService() view returns (address)",
    "function getRevenueDistribution() view returns (address)",
    "function hasRole(bytes32 role, address account) view returns (bool)",
    "function DEFAULT_ADMIN_ROLE() view returns (bytes32)"
  ];

  const registry = new ethers.Contract(registryAddress, registryAbi, adminSigner);

  try {
    // Verify admin role
    const adminRole = await registry.DEFAULT_ADMIN_ROLE();
    const hasAdmin = await registry.hasRole(adminRole, adminSigner.address);
    
    if (!hasAdmin) {
      console.error('❌ Admin wallet does not have DEFAULT_ADMIN_ROLE in PlatformRegistry!');
      console.log('   You may need to redeploy PlatformRegistry with PRIVATE_KEY_2 as admin');
      return;
    }

    console.log('✅ Admin wallet has DEFAULT_ADMIN_ROLE\n');

    // Check current values
    const currentFactory = await registry.getPropertyFactory();
    const currentApproval = await registry.getApprovalService();
    const currentRevenue = await registry.getRevenueDistribution();

    console.log('📋 Current registry values:');
    console.log(`   PropertyFactory: ${currentFactory}`);
    console.log(`   ApprovalService: ${currentApproval}`);
    console.log(`   RevenueDistribution: ${currentRevenue}\n`);

    // Update if needed
    const updates = [];

    if (currentFactory.toLowerCase() !== factoryAddress.toLowerCase()) {
      console.log('🔄 Updating PropertyFactory...');
      const tx1 = await registry.setPropertyFactory(factoryAddress);
      await tx1.wait();
      console.log('✅ PropertyFactory updated');
      updates.push('PropertyFactory');
    } else {
      console.log('✅ PropertyFactory already correct');
    }

    if (currentApproval.toLowerCase() !== approvalAddress.toLowerCase()) {
      console.log('🔄 Updating ApprovalService...');
      const tx2 = await registry.setApprovalService(approvalAddress);
      await tx2.wait();
      console.log('✅ ApprovalService updated');
      updates.push('ApprovalService');
    } else {
      console.log('✅ ApprovalService already correct');
    }

    if (currentRevenue.toLowerCase() !== revenueAddress.toLowerCase()) {
      console.log('🔄 Updating RevenueDistribution...');
      const tx3 = await registry.setRevenueDistribution(revenueAddress);
      await tx3.wait();
      console.log('✅ RevenueDistribution updated');
      updates.push('RevenueDistribution');
    } else {
      console.log('✅ RevenueDistribution already correct');
    }

    if (updates.length === 0) {
      console.log('\n✅ All registry values are already correct!');
    } else {
      console.log(`\n✅ Updated ${updates.length} contract address(es) in registry`);
    }

    // Verify final values
    console.log('\n📋 Final registry values:');
    const finalFactory = await registry.getPropertyFactory();
    const finalApproval = await registry.getApprovalService();
    const finalRevenue = await registry.getRevenueDistribution();
    console.log(`   PropertyFactory: ${finalFactory}`);
    console.log(`   ApprovalService: ${finalApproval}`);
    console.log(`   RevenueDistribution: ${finalRevenue}`);

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

setupRegistry().catch((error) => {
  console.error('Fatal error:', error);
  process.exit(1);
});


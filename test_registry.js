const ethers = require('ethers');

// Check what addresses the registry contract returns
async function checkRegistry() {
  try {
    // Registry contract address
    const registryAddress = '0x113FdC450314d59eAbEEfA2ff9F232DD7b08ad38';
    const provider = new ethers.JsonRpcProvider('https://ethereum-sepolia-rpc.publicnode.com');

    // Minimal ABI for registry functions
    const abi = [
      "function getApprovalService() view returns (address)",
      "function getPropertyFactory() view returns (address)",
      "function getRevenueDistribution() view returns (address)"
    ];

    const registry = new ethers.Contract(registryAddress, abi, provider);

    console.log('Checking registry contract at:', registryAddress);

    const approvalAddr = await registry.getApprovalService();
    const factoryAddr = await registry.getPropertyFactory();
    const revenueAddr = await registry.getRevenueDistribution();

    console.log('\nRegistry returns:');
    console.log('ApprovalService:', approvalAddr);
    console.log('PropertyFactory:', factoryAddr);
    console.log('RevenueDistribution:', revenueAddr);

    // Check if these match our expected addresses
    const expectedApproval = '0xB88F19F1074E6d7c23B6951360c00AEd1C4Ac635';
    const expectedRevenue = '0xD7ebec79B2e8D218F59A3777fa615eae762Bb7f6';

    console.log('\nComparison:');
    console.log('Approval matches expected:', approvalAddr.toLowerCase() === expectedApproval.toLowerCase());
    console.log('Revenue matches expected:', revenueAddr.toLowerCase() === expectedRevenue.toLowerCase());

  } catch (error) {
    console.error('Error checking registry:', error);
  }
}

checkRegistry();


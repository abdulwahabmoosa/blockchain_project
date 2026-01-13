const ethers = require('ethers');

// Check approval status directly
async function checkApproval() {
  try {
    // Contract details - use the updated address from smart_contract folder
    const contractAddress = '0xB88F19F1074E6d7c23B6951360c00AEd1C4Ac635';
    const provider = new ethers.JsonRpcProvider('https://ethereum-sepolia-rpc.publicnode.com');

    // Minimal ABI for checking approvals
    const abi = [
      "function check(address user) view returns (bool)"
    ];

    const contract = new ethers.Contract(contractAddress, abi, provider);

    // Test addresses from your screenshots
    const adminAddress = '0x98E5a749E25C56e19C28008505DF75aFf4988049';
    const userAddress = '0x2b4f7f1e6d8b9c5a3e2f8d7a9b6c4e5f1a8b9c2d';

    console.log('Checking approval status for:');
    console.log('Admin:', adminAddress);
    console.log('User:', userAddress);

    const adminApproved = await contract.check(adminAddress);
    const userApproved = await contract.check(userAddress);

    console.log('\nResults:');
    console.log('Admin approved:', adminApproved);
    console.log('User approved:', userApproved);

  } catch (error) {
    console.error('Error checking approval:', error);
  }
}

checkApproval();


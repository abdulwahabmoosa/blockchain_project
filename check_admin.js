const ethers = require('ethers');

// Check current admin status on the contract
async function checkAdminStatus() {
  try {
    // Contract details
    const contractAddress = '0xB88F19F1074E6d7c23B6951360c00AEd1C4Ac635';
    const provider = new ethers.JsonRpcProvider('https://ethereum-sepolia-rpc.publicnode.com');

    // Minimal ABI for checking roles
    const abi = [
      "function hasRole(bytes32 role, address account) view returns (bool)",
      "function ADMIN_ROLE() view returns (bytes32)"
    ];

    const contract = new ethers.Contract(contractAddress, abi, provider);

    // Get ADMIN_ROLE
    const adminRole = await contract.ADMIN_ROLE();
    console.log('ADMIN_ROLE:', adminRole);

    // Check current admin wallets
    const oldAdmin = '0x8626f6940E2eb28930eFb4CeF49B2d1F2C9C1199';
    const newAdmin = '0x98E5a749E25C56e19C28008505DF75aFf4988049';

    const oldAdminHasRole = await contract.hasRole(adminRole, oldAdmin);
    const newAdminHasRole = await contract.hasRole(adminRole, newAdmin);

    console.log('Old admin (0x8626...) has role:', oldAdminHasRole);
    console.log('New admin (0x98E5...) has role:', newAdminHasRole);

    if (newAdminHasRole) {
      console.log('✅ New admin wallet has admin privileges');
    } else {
      console.log('❌ New admin wallet does NOT have admin privileges');
    }

    if (oldAdminHasRole) {
      console.log('⚠️ Old admin wallet still has admin privileges');
    } else {
      console.log('✅ Old admin wallet has been revoked');
    }

  } catch (error) {
    console.error('Error checking admin status:', error);
  }
}

checkAdminStatus();



const ethers = require('ethers');

// Check what address corresponds to the frontend admin private key
const adminPrivateKey = '0x65f24585f0bafc504a5d88a5cc4b7eb8b10ff71aab7f36284f16161414c15b6b';
const expectedAdminAddress = '0x98E5a749E25C56e19C28008505DF75aFf4988049';

try {
  const wallet = new ethers.Wallet(adminPrivateKey);
  console.log('Frontend admin private key address:', wallet.address);
  console.log('Expected admin address:', expectedAdminAddress);
  console.log('Addresses match:', wallet.address.toLowerCase() === expectedAdminAddress.toLowerCase());
} catch (error) {
  console.error('Error:', error);
}

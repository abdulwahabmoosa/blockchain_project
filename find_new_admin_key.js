const ethers = require('ethers');

// New admin address that should be used
const newAdminAddress = '0x98E5a749E25C56e19C28008505DF75aFf4988049';

// Test different private keys to find which one controls the new admin address
console.log('Looking for private key that controls:', newAdminAddress);
console.log('');

// Test PRIVATE_KEY from environment or common test keys
const testKeys = [
  '0x700273c7efb8c191a5f6f501ea27e179f124002e201f0948f8c2145829c15e57', // Old key
  '0x65f24585f0bafc504a5d88a5cc4b7eb8b10ff71aab7f36284f16161414c15b6b', // Old key 2
  // Add more keys here if you have them
];

testKeys.forEach((key, index) => {
  try {
    const wallet = new ethers.Wallet(key);
    const isMatch = wallet.address.toLowerCase() === newAdminAddress.toLowerCase();
    console.log(`Key ${index + 1}: ${wallet.address} ${isMatch ? '✅ MATCH!' : '❌ No match'}`);
    if (isMatch) {
      console.log(`   Private key: ${key}`);
    }
  } catch (error) {
    console.log(`Key ${index + 1}: Invalid key format`);
  }
});

// Generate a new key pair for testing if needed
console.log('\nGenerating a new key pair for reference:');
const newWallet = ethers.Wallet.createRandom();
console.log('New random address:', newWallet.address);
console.log('New random private key:', newWallet.privateKey);
console.log('\n⚠️  IMPORTANT: Replace your environment PRIVATE_KEY with the correct one for the new admin address!');

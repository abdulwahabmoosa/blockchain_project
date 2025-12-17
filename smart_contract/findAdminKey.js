const ethers = require('ethers');

// Test different private keys to find which one controls the admin address
const adminAddress = '0x8626f6940E2eb28930eFb4CeF49B2d1F2C9C1199';

// Test PRIVATE_KEY
const pk1 = '0x700273c7efb8c191a5f6f501ea27e179f124002e201f0948f8c2145829c15e57';
const wallet1 = new ethers.Wallet(pk1);
console.log('PRIVATE_KEY address:', wallet1.address);
console.log('Is admin?', wallet1.address.toLowerCase() === adminAddress.toLowerCase());

// Test PRIVATE_KEY_2
const pk2 = '0x65f24585f0bafc504a5d88a5cc4b7eb8b10ff71aab7f36284f16161414c15b6b';
const wallet2 = new ethers.Wallet(pk2);
console.log('PRIVATE_KEY_2 address:', wallet2.address);
console.log('Is admin?', wallet2.address.toLowerCase() === adminAddress.toLowerCase());

console.log('\nExpected admin address:', adminAddress);
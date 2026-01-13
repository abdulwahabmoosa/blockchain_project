const fs = require('fs');
const path = require('path');

/**
 * Helper script to update contract-addresses.json after redeploying PropertyFactory
 * Usage: node update_addresses.js <new_factory_address>
 */

const newFactoryAddress = process.argv[2];

if (!newFactoryAddress) {
  console.error('❌ Usage: node update_addresses.js <new_factory_address>');
  console.error('   Example: node update_addresses.js 0x1234567890123456789012345678901234567890');
  process.exit(1);
}

// Validate address format
if (!/^0x[a-fA-F0-9]{40}$/.test(newFactoryAddress)) {
  console.error('❌ Invalid address format! Must be a valid Ethereum address (0x followed by 40 hex characters)');
  process.exit(1);
}

const addressesPath = path.join(__dirname, '..', 'frontend', 'src', 'deployments', 'sepolia', 'contract-addresses.json');

try {
  // Read existing addresses
  const addresses = JSON.parse(fs.readFileSync(addressesPath, 'utf8'));
  
  // Update PropertyFactory address
  const oldAddress = addresses.PropertyFactory;
  addresses.PropertyFactory = newFactoryAddress;
  
  // Write back
  fs.writeFileSync(addressesPath, JSON.stringify(addresses, null, 2));
  
  console.log('✅ Updated contract-addresses.json');
  console.log(`   Old PropertyFactory: ${oldAddress}`);
  console.log(`   New PropertyFactory: ${newFactoryAddress}`);
  
} catch (error) {
  console.error('❌ Error updating addresses file:', error.message);
  process.exit(1);
}


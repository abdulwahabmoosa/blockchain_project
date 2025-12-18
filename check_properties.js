const ethers = require('ethers');

// Check what properties exist
async function checkProperties() {
  try {
    const provider = new ethers.JsonRpcProvider('https://ethereum-sepolia-rpc.publicnode.com');

    // Factory contract address from registry
    const factoryAddress = '0x9Aa0a2Ee2Cb91587484b0F7F2BE992E1F4484F1e';

    // Minimal ABI for factory functions
    const abi = [
      "function getPropertiesCount() view returns (uint256)",
      "function getProperty(uint256) view returns (address asset, address token, address owner, string memory name, string memory symbol, uint256 valuation, uint256 tokenSupply, string memory dataHash, uint8 status)"
    ];

    const factory = new ethers.Contract(factoryAddress, abi, provider);

    console.log('Checking factory contract at:', factoryAddress);

    const count = await factory.getPropertiesCount();
    console.log('Total properties:', count.toString());

    // Get details of first few properties
    for (let i = 0; i < Math.min(count, 5); i++) {
      try {
        const property = await factory.getProperty(i);
        console.log(`\nProperty ${i}:`);
        console.log('Asset Address:', property.asset);
        console.log('Token Address:', property.token);
        console.log('Owner:', property.owner);
        console.log('Name:', property.name);
        console.log('Symbol:', property.symbol);
        console.log('Valuation:', ethers.formatEther(property.valuation));
        console.log('Token Supply:', ethers.formatEther(property.tokenSupply));
        console.log('Data Hash:', property.dataHash);
        console.log('Status:', property.status);
      } catch (err) {
        console.error(`Error getting property ${i}:`, err.message);
      }
    }

  } catch (error) {
    console.error('Error checking properties:', error);
  }
}

checkProperties();

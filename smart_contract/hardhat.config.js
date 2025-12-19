// require("@nomicfoundation/hardhat-toolbox");

/** @type import('hardhat/config').HardhatUserConfig */
require('@nomiclabs/hardhat-waffle');
require('@nomiclabs/hardhat-ethers');
require("dotenv").config({ path: "../.env" });
const { SEPOLIA_RPC_BC, PRIVATE_KEY, PRIVATE_KEY_2, REGISTRY, REVENUE, PINATA_API_KEY, PINATA_API_SECRET, PINATA_JWT_TOKEN } = process.env;

// Build networks object
const networks = {
  localhost: {
    url: 'http://127.0.0.1:8545',
  },
};

// Add sepolia network if environment variables are available
if (SEPOLIA_RPC_BC && PRIVATE_KEY_2) {
  networks.sepolia = {
    url: SEPOLIA_RPC_BC,
    accounts: [PRIVATE_KEY_2],
  };
} else {
  console.warn('⚠️  SEPOLIA_RPC_BC or PRIVATE_KEY_2 not found. Sepolia network will not be available.');
}

module.exports = {
  solidity: {
    compilers: [
      {
        version: '0.8.19',
        settings: {
          optimizer: {
            enabled: true,
            runs: 200,
          },
        },
      },
    ],
  },
  networks,

  paths: {
    sources: './contracts',
    tests: './test',
    cache: './cache',
    artifacts: './artifacts',
  },
};

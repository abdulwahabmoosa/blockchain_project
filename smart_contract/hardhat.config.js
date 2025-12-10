// require("@nomicfoundation/hardhat-toolbox");

/** @type import('hardhat/config').HardhatUserConfig */
require('@nomiclabs/hardhat-waffle');
require('@nomiclabs/hardhat-ethers');
require('dotenv').config();
const { PRIVATE_KEY, SEPOLIA_RPC, PRIVATE_KEY_2 } = process.env;

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

  networks: {
      sepolia: {
        url: SEPOLIA_RPC,
        accounts: [PRIVATE_KEY, PRIVATE_KEY_2],
      },
      localhost: {
        url: 'http://127.0.0.1:8545',
      }
    },

  paths: {
    sources: './contracts',
    tests: './test',
    cache: './cache',
    artifacts: './artifacts',
  },
};

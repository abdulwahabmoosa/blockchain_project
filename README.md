# blockchain_project

#for the smart contract part
![alt text](image.png)

This project implements the full blockchain layer for a real estate tokenization platform using Solidity and Hardhat. The system includes a PlatformRegistry (address hub), ApprovalService (admin-controlled user authorization), PropertyFactory (deploys ERC721/ ERC20 for each property), and RevenueDistribution (handles rental payouts). All contracts are deployed to Sepolia, and the deployment workflow automatically records contract addresses and property details into deployments/sepolia/contract-addresses.json. Clean ABI files are provided for frontend and backend integration. The blockchain layer is fully complete and ready to be consumed by the rest of the application.

# blockchain_project

## Smart Contract

This project implements the full blockchain layer for a real estate tokenization platform using Solidity and Hardhat. The system includes a PlatformRegistry (address hub), ApprovalService (admin-controlled user authorization), PropertyFactory (deploys ERC721/ ERC20 for each property), and RevenueDistribution (handles rental payouts). All contracts are deployed to Sepolia, and the deployment workflow automatically records contract addresses and property details into deployments/sepolia/contract-addresses.json. Clean ABI files are provided for frontend and backend integration. The blockchain layer is fully complete and ready to be consumed by the rest of the application.

## Backend

Follow these steps to run the smart contracts, database and backend in one go locally.

### First time setup

1. Install [docker](https://docs.docker.com/engine/install/) or [podman](https://podman.io/docs/installation)

> Avoid using desktop/GUI version.

### Run

To run the full backend stack:

```bash
podman compose up -d
```

To stop:

```bash
podman compose down
```

> If using docker, use "docker" instead of "podman" in the command, the rest of the command is the same.

> Note that these commands need to be ran at the top level directory of the project where [the compose file](compose.yaml) is found.
> For example: ~/dev/blockchain_project or C:\Users\user\dev\blockchain_project

# Real Estate Tokenization Backend

This backend acts as the secure bridge between the Postgres database and the Ethereum Blockchain. It manages user identities (JWT), handles IPFS file uploads, and executes administrative blockchain transactions (minting and revenue deposits).

## Core Architecture

- **Language:** Go (Golang)
- **Database:** PostgreSQL (with GORM)
- **Blockchain:** Ethereum (Sepolia Testnet) via go-ethereum
- **Storage:** IPFS (via Pinata API)
- **Auth:** JWT for users

## Quick Start

### 1. Prerequisites

- **Podman** or **Docker**
- **Sepolia RPC URL** (WebSocket `wss://...` is recommended for event listeners)
- **Pinata API Token**

### 2. Environment Variables

Create a `.env` file in the root directory. These must match the variables used in your code:

```env
# Database Connection
DB_URL="postgres://user:password@db:5432/realestate?sslmode=disable"

# Blockchain Configuration
SEPOLIA_RPC="wss://[eth-sepolia.g.alchemy.com/v2/YOUR_KEY](https://eth-sepolia.g.alchemy.com/v2/YOUR_KEY)"
PRIVATE_KEY="0xYourWalletPrivateKey"
REGISTRY="0xYourRegistryContractAddress"

# IPFS Storage
PINATA_JWT_TOKEN="your_pinata_jwt_token"

# Admin Seeding (Optional - Defaults are defined in db.go)
ADMIN_EMAIL="admin@admin.com"
ADMIN_PASSWORD="password"
```

### 3\. Run with Podman/Docker

To start the database and backend services:

```bash
podman compose up -d
```

To stop:

```bash
podman compose down
```

To view logs (essential for debugging blockchain events):

```bash
# Get container name from 'podman ps' or compose output (e.g., blockchain-backend-1)
podman logs <your_container_name>
```

The server will start on **port 3000**.

## Frontend Integration Guide

This API uses standard REST conventions. Authentication is handled via JWT tokens, which must be included in the header for all protected routes.

**Base URL:** `http://localhost:3000`

### 1\. Authentication Headers

Include this header in all requests except `/login`, `/register`, and `/ping`.

```http
Authorization: Bearer <your_token_here>
Content-Type: application/json
```

### 2\. API Reference

#### User Management

**POST** `/register`
Registers a new user. The wallet address must be a valid Ethereum address.

- **Body:**
    ```json
    {
        "email": "user@example.com",
        "password": "password123",
        "name": "John Doe",
        "wallet_address": "0x123..."
    }
    ```

**POST** `/login`
Authenticates a user and returns a JWT token.

- **Body:**
    ```json
    {
        "email": "user@example.com",
        "password": "password123"
    }
    ```
- **Response:**
    ```json
    {
        "token": "eyJhbGci...",
        "user": { "id": "...", "email": "...", "role": "user" }
    }
    ```

#### Properties

**GET** `/properties`
Fetches all properties stored in the database.

- **Security:** Protected (Requires JWT)
- **Response:** A JSON array of property objects.

**GET** `/properties/{id}`
Fetches details for a single property by its UUID.

- **Security:** Protected (Requires JWT)

**POST** `/upload`
Uploads a file to IPFS via Pinata. This is the first step in the property creation flow.

- **Security:** Protected (Requires JWT)
- **Format:** `multipart/form-data`
- **Field:** `file` (The binary file content)
- **Response:**
    ```json
    {
        "ipfs_hash": "QmHash...",
        "url": "[https://gateway.pinata.cloud/ipfs/QmHash](https://gateway.pinata.cloud/ipfs/QmHash)..."
    }
    ```

#### Admin Actions

These endpoints require the user to have the `admin` role.

**POST** `/approve-user`
Whitelists a user on the smart contract, allowing them to participate in the platform.

- **Security:** Admin JWT Only
- **Body:**
    ```json
    {
        "wallet_address": "0xUserWallet..."
    }
    ```

**POST** `/properties`
Mints a new Property NFT and Fractional Tokens on the blockchain.

- **Security:** Admin JWT Only
- **Body:**
    ```json
    {
        "owner_address": "0xOwner...",
        "name": "Villa Sunset",
        "symbol": "VIL",
        "data_hash": "QmHashFromUploadStep...",
        "valuation": 500000,
        "token_supply": 1000
    }
    ```

**POST** `/revenue/distribute`
Deposits stablecoins into the revenue distribution contract.

- **Security:** Admin JWT Only
- **Body:**
    ```json
    {
        "token_address": "0xPropertyTokenAddress...",
        "stablecoin_address": "0xUSDCAddress...",
        "amount": 1000
    }
    ```

### 3\. Data Types (TypeScript Interfaces)

Use these interfaces to strongly type your frontend API responses and requests.

#### Enums

```typescript
export enum UserRole {
    Admin = "admin",
    User = "user",
}

export enum PropertyStatus {
    Active = "Active",
    Paused = "Paused",
    Disputed = "Disputed",
    Closed = "Closed",
}
```

#### Core Models

**User**
(Response from `/login` inside the `user` object)

```typescript
export interface User {
    ID: string; // UUID
    WalletAddress: string; // Ethereum Address (0x...)
    Email: string;
    Name: string;
    Role: UserRole; // "admin" | "user"
    IsApproved: boolean; // True if whitelisted on-chain
    CreatedAt: string; // ISO 8601 Date String
    UpdatedAt: string;
}
```

**Property**
(Response from `GET /properties`)

```typescript
export interface Property {
    ID: string; // UUID
    OnchainAssetAddress: string; // ERC721 Address
    OnchainTokenAddress: string; // ERC20 Address
    OwnerWallet: string; // 0x...
    MetadataHash: string; // IPFS CID (Qm...)
    Valuation: number; // Stored as float in DB
    Status: PropertyStatus; // "Active" | "Paused" | ...
    CreatedAt: string; // ISO 8601 Date String
}
```

#### Request Payloads

**Auth**

```typescript
// POST /register
export interface RegisterPayload {
    email: string;
    password: string;
    name: string;
    wallet_address: string;
}

// POST /login
export interface LoginPayload {
    email: string;
    password: string;
}

// Login Response
export interface LoginResponse {
    token: string;
    user: User;
}
```

**Property Creation**
_Note: `valuation` and `token_supply` are integers._

```typescript
// POST /properties (Admin Only)
export interface CreatePropertyPayload {
    owner_address: string;
    name: string;
    symbol: string; // 3-4 chars (e.g. "SVL")
    data_hash: string; // IPFS Hash from /upload endpoint
    valuation: number; // Integer (e.g. 500000 cents)
    token_supply: number; // Integer (e.g. 10000)
}
```

**Revenue & Admin**

```typescript
// POST /revenue/distribute (Admin Only)
export interface DistributeRevenuePayload {
    token_address: string; // The Property Token Address
    stablecoin_address: string; // The Stablecoin Address (USDC)
    amount: number; // Amount to distribute
}

// POST /approve-user (Admin Only)
export interface ApproveUserPayload {
    wallet_address: string;
}
```

**File Upload**
Response from `POST /upload`.

```typescript
export interface UploadResponse {
    ipfs_hash: string; // Use this as 'data_hash' in CreateProperty
    url: string; // Public gateway URL for viewing
}
```

### 4\. Web3 Frontend Logic

The backend does not expose an endpoint for contract configuration. The frontend application must import contract addresses from a local JSON file (generated by your deployment script) or environment variables.

#### Connecting Wallet

The frontend should use a library like `ethers.js` to connect to the user's wallet.

```javascript
import contractAddresses from "./deployments/sepolia/contract-addresses.json";
import { ethers } from "ethers";

const connectWallet = async () => {
    if (!window.ethereum) return;

    const provider = new ethers.BrowserProvider(window.ethereum);
    await provider.send("eth_requestAccounts", []);

    // Access contracts using the imported addresses
    const registryAddress = contractAddresses.PlatformRegistry;
};
```

#### Creating a Property (Frontend Flow)

1.  **Upload:** Send the file to `POST /upload`. Receive the `ipfs_hash`.
2.  **Create:** Send the property details and the `ipfs_hash` to `POST /properties`.

<!-- end list -->

```javascript
const createPropertyFlow = async (file, data) => {
    const token = localStorage.getItem("token");

    // 1. Upload File
    const formData = new FormData();
    formData.append("file", file);

    const uploadRes = await fetch("http://localhost:3000/upload", {
        method: "POST",
        headers: { Authorization: `Bearer ${token}` },
        body: formData,
    });
    const { ipfs_hash } = await uploadRes.json();

    // 2. Create Property
    const createRes = await fetch("http://localhost:3000/properties", {
        method: "POST",
        headers: {
            Authorization: `Bearer ${token}`,
            "Content-Type": "application/json",
        },
        body: JSON.stringify({ ...data, data_hash: ipfs_hash }),
    });

    const result = await createRes.json();
    console.log("Transaction Hash:", result.tx_hash);
};
```

## Testing

A bash script is provided to test the full API lifecycle.

```bash
chmod +x test_endpoints.sh
USER_WALLET="0xAddr..." ./test_endpoints.sh
```

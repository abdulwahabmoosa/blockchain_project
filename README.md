# Real Estate Tokenization Platform (RWA)

A robust backend for tokenizing Real World Assets (Real Estate) on the Ethereum blockchain. This system bridges the gap between Web2 (PostgreSQL, IPFS) and Web3 (Smart Contracts), handling user identity, document verification, and asset minting.

## 🚀 Key Features

- **User Identity**:
    - JWT-based authentication linked to Ethereum Wallet addresses.
    - **Profile Management**: Users can update details, reset passwords, or delete accounts (Hard Delete).
- **Role-Based Access Control (RBAC)**:
    - **Users**: Register, submit properties, view assets.
    - **Admins**: Review and Approve/Reject users and property listings via protected endpoints.
- **Smart Property Creation**:
    - **Atomic Multipart Upload**: Uploads files (Deeds, Images) and metadata in a single request.
    - **Auto-Inference**: Automatically detects document types (e.g., "Deed", "Inspection") based on filenames.
    - **IPFS**: Decentralized storage using Pinata.
- **Blockchain Integration**:
    - **Event Listeners**: Background workers sync on-chain events (Minting, Approvals) with the local database.
    - **Smart Contracts**: Integration with Property Factory and Compliance contracts (Sepolia/Hardhat).
- **Infrastructure**: Fully containerized with Docker & Docker Compose.

---

## 🛠 Tech Stack

- **Backend**: Go (Golang) using Chi Router
- **Database**: PostgreSQL (GORM ORM)
- **Blockchain**: Hardhat (Local Node) or Sepolia Testnet
- **Storage**: IPFS (Pinata)
- **Containerization**: Docker & Docker Compose

---

## ⚡ Quick Start (Docker)

The easiest way to run the entire stack (Blockchain Node, Database, API) is via Docker Compose.

### 1. Configuration

Create a `.env` file in the root directory:

```ini
# Server Config
PORT=3000
JWT_SECRET="super-secret-key-change-me"

# Database (Service name is 'db')
DB_URL="host=db user=blockchain-db password=password dbname=db port=5432 sslmode=disable"

# Blockchain Configuration
# Use "http://blockchain:8545" if running the local hardhat node via Docker
SEPOLIA_RPC="http://blockchain:8545"
PRIVATE_KEY="0x..." # Admin Wallet Private Key (Deployer)
REGISTRY="0x..."    # Registry Contract Address

# IPFS (Pinata)
# IMPORTANT: PINATA_JWT_TOKEN must be set in your environment for the backend service
# to upload files to IPFS. Ensure this variable is available to the backend container.
# If using Docker Compose, add it to your .env file and ensure compose.yaml references it.
PINATA_JWT_TOKEN="your-pinata-jwt-token"

# Admin Seeding (Created on first run)
ADMIN_EMAIL="admin@admin.com"
ADMIN_PASSWORD="password"

```

### 2. Run Services

```bash
docker compose up -d

```

- **API**: `http://localhost:3000`
- **Blockchain Node**: `http://localhost:8545`
- **Postgres**: Port `5433` (mapped locally)

---

## 📡 API Documentation

### 🔓 Public Routes

#### Health Check

`GET /ping`

- Returns: `pong`

#### Register

`POST /register`

```json
{
    "email": "user@example.com",
    "password": "password123",
    "name": "John Doe",
    "wallet_address": "0x123..."
}
```

#### Login

`POST /login`

```json
{
    "email": "user@example.com",
    "password": "password123"
}
```

- **Returns**: `{"token": "..."}`.
- **Auth**: Include `Authorization: Bearer <token>` in all subsequent requests.

---

### 👤 User Profile (Authenticated)

#### Update Profile

`PUT /users/me`

- Update your own Name or Email.

```json
{ "name": "Jane Doe", "email": "new@example.com" }
```

#### Reset Password

`POST /users/me/reset-password`

```json
{ "old_password": "oldpass", "new_password": "newpass" }
```

#### Delete Account

`DELETE /users/me`

- **Action**: Permanently removes the user account from the database.

---

### 🏠 Property Routes (Authenticated)

#### Create Property (Multipart)

`POST /properties`

- **Headers**: `Content-Type: multipart/form-data`
- **Description**: Uploads files to IPFS, saves metadata to DB, and mints on-chain.

**Form Data Fields:**

1. **`files`** (Binary): Select one or more files (PDFs, JPGs).
2. **`data`** (Text/JSON String):

```json
{
    "owner_address": "0xUserWallet...",
    "name": "Sunset Villa",
    "symbol": "SVL",
    "valuation": 500000,
    "token_supply": 1000
}
```

#### Get Properties

`GET /properties`

- Returns list of all active properties.

#### Get Property Details

`GET /properties/{id}`

- Returns details and document URLs for a specific property.

---

### 🛡 Admin Routes (Role: Admin)

#### User Approval

`POST /users/approval`

- Triggers on-chain transaction to whitelist or ban a user.

```json
{
    "wallet_address": "0xUserWallet...",
    "status": "approved" // or "rejected"
}
```

#### Property Approval

`POST /properties/approval`

- Updates the listing status in the database.

```json
{
    "property_id": "uuid-of-property",
    "status": "approved" // or "rejected"
}
```

---

## 🧪 Testing

A standalone script is available to test the full lifecycle (Registration -> Admin Approval -> Property Creation -> Profile Updates).

```bash
# 1. Make executable
chmod +x test_endpoints.sh

# 2. Run against localhost
./test_endpoints.sh
```

## 📂 Project Structure

```
├── api/             # HTTP Handlers (Controllers)
├── auth/            # JWT Middleware & Utils
├── blockchain/      # Smart Contract Bindings & Event Listeners
├── db/              # Database Models & GORM Logic
├── ipfs/            # Pinata Integration
├── compose.yaml     # Docker Services
└── main.go          # Application Entry Point
```

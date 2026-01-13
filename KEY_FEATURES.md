## 🚀 Key Features

### **User Identity & Authentication**
- **JWT-based Authentication**: Token-based authentication system linked to Ethereum Wallet addresses
- **Wallet Address Integration**: Users register and login with Ethereum wallet addresses
- **Profile Management**: 
  - Users can view their profile information
  - Update profile details (name, email)
  - Reset password (requires old password verification)
  - Delete account (Hard Delete - permanent account deletion)

### **Role-Based Access Control (RBAC)**
- **Users**: 
  - Register accounts with wallet addresses
  - Submit property upload requests
  - View property listings and details
  - Purchase property tokens
  - Transfer tokens to other approved users
  - Claim revenue distributions
  - View their token purchases and investments
  
- **Admins**: 
  - Review and approve/reject user registrations via protected endpoints
  - Review and approve/reject property upload requests
  - Approve/reject existing property listings
  - Create properties directly (bypassing upload request flow)
  - Distribute revenue to property token holders
  - Manage all users and properties

### **Property Management System**

#### **Property Upload Request Workflow**
- **User Submission**: Authenticated users can submit property upload requests with:
  - Property details (name, symbol, valuation, token supply)
  - Multiple document files (deeds, inspection reports, valuation reports, photos, etc.)
- **Admin Review**: Admins can:
  - View all pending property upload requests
  - Approve requests (creates property on blockchain and database)
  - Reject requests with optional rejection reason
- **Document Storage**: All property documents are uploaded to IPFS via Pinata

#### **Property Features**
- **Property Listing**: Properties include:
  - On-chain asset address (ERC721 PropertyAsset contract)
  - On-chain token address (ERC20 PropertyToken contract)
  - Owner wallet address
  - Valuation and token supply
  - Status tracking (Active, Paused, Disputed, Closed)
  - IPFS metadata hash for documents
  - Transaction hashes for audit trail
- **Property Documents**: Multiple document types supported (Deeds, Inspection Reports, Valuation Reports, Photos, Legal Contracts)
- **Property Metadata**: Fetch and display property metadata from IPFS

### **Blockchain Integration**

#### **Smart Contracts**
- **PropertyFactory**: Creates new property assets and tokens on-chain
- **ApprovalService**: Manages user approval whitelist (only approved users can receive tokens)
- **PropertyAsset (ERC721)**: Non-fungible token representing the property itself
- **PropertyToken (ERC20)**: Fungible token representing property shares with snapshot capabilities
- **RevenueDistribution**: Handles automated revenue distribution to token holders via snapshots
- **PlatformRegistry**: Central registry for managing platform contracts and configurations

#### **Blockchain Event Listeners**
- **Property Creation Events**: Automatically syncs newly created properties from blockchain to database
- **Revenue Distribution Events**: Tracks revenue deposits and claims on-chain
- **User Approval Events**: Syncs user approval status from blockchain to database

#### **Network Support**
- Sepolia Testnet integration
- Hardhat local development support

### **Token Management**

#### **Token Purchase Workflow**
- **Purchase Request**: Buyers can submit token purchase requests with:
  - Payment transaction hash (ETH payment)
  - Token amount desired
  - Purchase price
- **Owner Approval**: Property owners can:
  - View pending token purchase requests
  - Approve purchases (transfers tokens to buyer)
  - Update purchase records with token transfer transaction hash
- **Purchase Tracking**: Database tracks all token purchases with transaction hashes

#### **Token Operations**
- **Token Balance**: Check token balance for any wallet address
- **Token Statistics**: View total supply, sold amount, available tokens, and percentage sold
- **Token Transfer**: Transfer property tokens to other approved wallet addresses
- **Token Purchases**: Users can view their purchase history

### **Revenue Distribution System**
- **Revenue Deposit** (Admin-only):
  - Admins deposit stablecoin revenue for properties
  - System automatically creates token snapshot at deposit time
  - Snapshot freezes token balances for fair distribution calculation
- **Revenue Claiming**:
  - Token holders can claim their proportional share of revenue
  - Share calculated based on token balance at snapshot time
  - Prevents double-claiming via on-chain tracking
  - Revenue distributed in stablecoin (USDC/USDT)
- **Distribution Tracking**: Database tracks all distributions and claims with transaction hashes

### **IPFS Integration**
- **Decentralized Storage**: All property documents stored on IPFS via Pinata
- **Document Upload**: Multi-file upload support with automatic IPFS hashing
- **Metadata Storage**: Property metadata stored as IPFS hashes on-chain
- **IPFS Gateway**: Documents accessible via Pinata gateway

### **Database & Data Management**
- **PostgreSQL Database**: 
  - User management with approval status
  - Property and property upload request tracking
  - Token purchase records
  - Revenue distribution and claim tracking
  - Transaction audit logs
  - Property document metadata
- **Hard Delete**: Complete account deletion with cascade deletion of related records
- **Transaction Tracking**: All blockchain transactions logged with status (pending, confirmed, failed)

### **Infrastructure**
- **Containerization**: Fully containerized with Docker & Docker Compose
- **Services**:
  - PostgreSQL database container
  - Go backend API server
  - React/TypeScript frontend with Vite
- **Database Migrations**: Automatic schema initialization
- **Environment Configuration**: Environment variable-based configuration for blockchain, database, and IPFS credentials

### **Additional Features**
- **Wallet Connection**: MetaMask wallet integration for frontend interactions
- **Transaction Monitoring**: Real-time blockchain transaction status tracking
- **User Investment Tracking**: Users can view all their property token investments
- **Pending Transfers Management**: Property owners can view and manage pending token purchase requests
- **Property Status Management**: Properties can be marked as Active, Paused, Disputed, or Closed
- **Search & Filtering**: Admin interfaces include search and filter capabilities for properties and users
- **Error Handling & Logging**: Comprehensive error handling and logging throughout the application

## 🛠 Tech Stack

- **Backend**: Go (Golang) using Chi Router
- **Database**: PostgreSQL (GORM ORM)
- **Blockchain**: Hardhat (Local Node) or Sepolia Testnet
- **Storage**: IPFS (Pinata)
- **Containerization**: Docker & Docker Compose
- **Frontend**: React + TypeScript, Vite, Tailwind CSS, `ethers` v6, `react-router-dom`, `lucide-react`, `tailwind-merge`


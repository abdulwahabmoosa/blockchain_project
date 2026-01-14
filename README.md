# Real Estate Tokenization Platform (RWA)

A robust backend for tokenizing Real World Assets (Real Estate) on the Ethereum blockchain. This system bridges the gap between Web2 (PostgreSQL, IPFS) and Web3 (Smart Contracts), handling user identity, document verification, and asset minting.
---


## API Documentation

### Public Routes

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

### User Profile (Authenticated)

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

### Property Routes (Authenticated)

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

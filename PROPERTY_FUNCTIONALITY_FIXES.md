# Property Functionality Fixes - Summary

This document summarizes all the fixes and improvements made to property-related functionality across the project.

## ✅ Completed Fixes

### 1. Backend API Endpoints

#### Added Endpoints:
- **GET `/properties/{id}/metadata`** - Fetches property metadata from IPFS
- **GET `/properties/{id}/token-balance/{wallet}`** - Gets token balance for a wallet address
- **POST `/properties/{id}/transfer`** - Transfers tokens (requires backend wallet to have tokens)

#### Updated Endpoints:
- **POST `/properties/approval`** - Now calls blockchain to approve/reject properties on-chain using `PropertyAsset.SetStatus()`

### 2. Blockchain Service Methods

Added new methods to `backend/blockchain/service.go`:
- `ApproveProperty(propertyAssetAddrStr string)` - Sets property status to Active (0) on-chain
- `RejectProperty(propertyAssetAddrStr string)` - Sets property status to Closed (3) on-chain
- `GetTokenBalance(tokenAddrStr, walletAddrStr string)` - Gets ERC20 token balance
- `TransferTokens(tokenAddrStr, toAddrStr string, amount *big.Int)` - Transfers tokens

### 3. Database Model Updates

- Added `Name` field to `Property` model in `backend/db/models/models.go`
- Updated property listener to fetch property name from PropertyAsset contract when properties are created

### 4. Frontend Updates

- Fixed `AdminManageProperties.tsx` to use the correct API endpoint (`api.updatePropertyApproval`)
- Added `updatePropertyApproval` method to `frontend/src/lib/api.ts`
- All property-related API calls now properly handle errors and display transaction hashes

### 5. IPFS Integration

- Backend can now fetch metadata from IPFS via the `/properties/{id}/metadata` endpoint
- PropertyDetails page displays metadata including images, description, and location

## ⚠️ Required Configuration

### PINATA_JWT_TOKEN Environment Variable

**IMPORTANT**: The backend service needs access to `PINATA_JWT_TOKEN` environment variable to upload files to IPFS.

**Current Status**: The `compose.yaml` file does not pass `PINATA_JWT_TOKEN` to the backend service (it's only available to the blockchain service).

**Solution**: You need to add `PINATA_JWT_TOKEN` to the backend service environment in `compose.yaml`:

```yaml
backend:
    # ... existing config ...
    environment:
        DB_URL: postgres://blockchain-db:TP075164@db:5432/db?sslmode=disable
        PINATA_JWT_TOKEN: ${PINATA_JWT_TOKEN}  # Add this line
```

**Note**: Since `compose.yaml` cannot be modified per requirements, you'll need to add this manually. Alternatively, ensure `PINATA_JWT_TOKEN` is available in your host environment when running `docker compose up`.

## 🔄 Property Workflow

### Creating Properties

1. **User Upload**: User uploads property document via `/upload` endpoint
2. **IPFS Storage**: Document is uploaded to IPFS via Pinata, returns IPFS hash
3. **Blockchain Creation**: Property is created on-chain via PropertyFactory
4. **Event Listener**: Background worker listens for `PropertyRegistered` event
5. **Database Storage**: Property is saved to database with name fetched from PropertyAsset contract

### Approving/Rejecting Properties

1. **Admin Action**: Admin clicks approve/reject in AdminManageProperties page
2. **Backend API**: Calls `/properties/approval` endpoint
3. **Blockchain Update**: Backend calls `PropertyAsset.SetStatus()` on-chain
4. **Database Update**: Property status is updated in database
5. **Response**: Transaction hash is returned to frontend

### Viewing Property Details

1. **Property Fetch**: Frontend calls `/properties/{id}` to get property data
2. **Metadata Fetch**: Frontend calls `/properties/{id}/metadata` to get IPFS metadata
3. **Token Balance**: Frontend calls smart contract directly to get token balance
4. **Display**: All information is displayed in PropertyDetails page

### Token Operations

- **Balance Check**: Done via smart contract calls (frontend) or `/properties/{id}/token-balance/{wallet}` (backend)
- **Transfer**: Done via smart contract (frontend) or `/properties/{id}/transfer` (backend)
- **Revenue Claim**: Done via smart contract directly from frontend

## 📝 Testing Checklist

- [ ] Property upload with file upload to IPFS works
- [ ] Property creation on blockchain succeeds
- [ ] Property appears in database with correct name
- [ ] Admin can approve properties (both DB and blockchain updated)
- [ ] Admin can reject properties (both DB and blockchain updated)
- [ ] Property metadata is fetched from IPFS correctly
- [ ] Token balances are displayed correctly
- [ ] Token transfers work (if backend wallet has tokens)
- [ ] Revenue claiming works from PropertyDetails page

## 🐛 Known Limitations

1. **Token Transfers via Backend**: The `/properties/{id}/transfer` endpoint requires the backend wallet to have tokens. In practice, users should transfer tokens directly from their wallets via the frontend.

2. **Property Name**: Property name is fetched from the PropertyAsset contract after creation. If the contract doesn't have a name set, it will be empty in the database.

3. **Metadata Format**: The metadata endpoint expects JSON format from IPFS. Ensure uploaded metadata files are valid JSON with fields like `name`, `description`, `images`, `location`.

## 🔧 Files Modified

### Backend
- `backend/api/handler.go` - Added new endpoints and updated approval handler
- `backend/blockchain/service.go` - Added PropertyAsset interaction methods
- `backend/db/models/models.go` - Added Name field to Property model
- `backend/blockchain/worker/listener.go` - Updated to fetch property name

### Frontend
- `frontend/src/lib/api.ts` - Added `updatePropertyApproval` method
- `frontend/src/Pages/AdminManageProperties.tsx` - Fixed to use correct API endpoint

### Documentation
- `README.md` - Added note about PINATA_JWT_TOKEN requirement
- `PROPERTY_FUNCTIONALITY_FIXES.md` - This file

# Property API Integration - Summary

This document summarizes the integration of `property.go` APIs with the frontend.

## ✅ Changes Made

### 1. Backend API Updates (`backend/api/property.go`)

#### Updated Functions:
- **`CreateProperty`**: Now handles multipart form data with files and JSON payload
  - Accepts multiple files via `files` field
  - Accepts JSON payload via `data` field
  - Uploads files to IPFS automatically
  - Creates property on blockchain
  - Returns transaction hash and file count

- **`UpdatePropertyApproval`**: Added blockchain integration
  - Calls `ApproveProperty()` or `RejectProperty()` on-chain
  - Updates database
  - Returns transaction hash if blockchain call succeeds

#### Helper Functions:
- `parseCreatePropertyRequest`: Parses multipart form and extracts files + JSON payload
- `processPropertyFiles`: Uploads files to IPFS and creates PropertyDocument records
- `createPropertyOnChain`: Calls blockchain service to create property
- `inferDocType`: Automatically detects document type from filename

### 2. Frontend API Client Updates (`frontend/src/lib/api.ts`)

- **`createProperty`**: Updated to send multipart form data
  - Accepts `files: File[]` parameter
  - Creates FormData with files and JSON payload
  - Sends to `/properties` endpoint

### 3. Frontend Page Updates

#### AdminCreateProperty.tsx
- ✅ Changed from single file upload to multiple file support
- ✅ Removed separate IPFS upload step
- ✅ Files are now sent directly with property creation
- ✅ Updated to use new API signature: `api.createProperty(payload, files)`

#### UploadProperty.tsx
- ✅ Changed from single file upload to multiple file support
- ✅ Removed separate IPFS upload step
- ✅ Files are now sent directly with property creation
- ✅ Updated to use new API signature: `api.createProperty(payload, files)`

#### CreateProperty.tsx
- ✅ Simplified from two-step process to single form
- ✅ Changed from single file upload to multiple file support
- ✅ Files are now sent directly with property creation
- ✅ Updated to use new API signature: `api.createProperty(payload, files)`

#### ViewProperties.tsx
- ✅ Updated to display property name when available

#### PropertyDetails.tsx
- ✅ Updated to use property name and metadata description

### 4. Backend Handler Cleanup (`backend/api/handler.go`)

- ✅ Removed duplicate `GetProperties`, `GetProperty`, `CreateProperty`, and `UpdatePropertyApproval` functions
- ✅ These functions are now only in `property.go`
- ✅ Routes still point to the same handler methods (they're in property.go)

## 📋 API Format

### Create Property Request

**Endpoint**: `POST /properties` (Admin only)

**Content-Type**: `multipart/form-data`

**Form Fields**:
- `files`: Array of files (multiple files supported)
- `data`: JSON string containing:
  ```json
  {
    "owner_address": "0x...",
    "name": "Property Name",
    "symbol": "ETH",
    "valuation": 100,
    "token_supply": 10
  }
  ```

**Response**:
```json
{
  "status": "success",
  "tx_hash": "0x...",
  "files_count": 2,
  "message": "Transaction submitted. Property will be created on blockchain and saved to database via event listener."
}
```

### Update Property Approval

**Endpoint**: `POST /properties/approval` (Admin only)

**Request Body**:
```json
{
  "property_id": "uuid",
  "status": "approved" // or "rejected"
}
```

**Response**:
```json
{
  "status": "success",
  "new_status": "approved",
  "property_id": "uuid",
  "tx_hash": "0x...", // Optional, if blockchain call succeeded
  "message": "Transaction submitted. DB updated." // Optional
}
```

## 🔄 Property Creation Workflow

1. **Frontend**: User selects multiple files and fills property details
2. **Frontend**: Calls `api.createProperty(payload, files)`
3. **Backend**: Parses multipart form data
4. **Backend**: Uploads each file to IPFS via Pinata
5. **Backend**: Creates PropertyDocument records (not saved to DB yet - will be linked when property is created)
6. **Backend**: Uses first file's IPFS hash as main hash
7. **Backend**: Calls blockchain to create property via PropertyFactory
8. **Blockchain**: PropertyFactory emits PropertyRegistered event
9. **Backend Listener**: Listens for event and saves property to database
10. **Response**: Returns transaction hash to frontend

## ⚠️ Important Notes

1. **Property Documents**: Documents are created in memory but not saved to database immediately. They will be linked to the property when the blockchain event is processed by the listener. If you need documents saved immediately, you may need to add that functionality.

2. **File Upload**: All files are uploaded to IPFS before blockchain transaction. If IPFS upload fails, the entire operation fails.

3. **Multiple Files**: The backend supports multiple files. The first file's hash is used as the main metadata hash for the blockchain.

4. **Document Type Inference**: Document types are automatically inferred from filenames (Deed, Inspection Report, Valuation Report, Photo, Legal Contract, General Document).

## 🧪 Testing

To test property creation:

1. **Admin Create Property**:
   - Navigate to Admin Dashboard → Create Property
   - Select owner wallet (must be blockchain-approved)
   - Enter property name and valuation
   - Upload one or more files
   - Submit

2. **User Upload Property**:
   - Navigate to Dashboard → Upload Property
   - Enter property name and valuation
   - Upload one or more files
   - Submit (uses logged-in user's wallet)

3. **Verify**:
   - Check backend logs for IPFS uploads
   - Check blockchain transaction hash
   - Wait for event listener to process
   - Verify property appears in database
   - Check property details page shows name and metadata

## 🔧 Files Modified

### Backend
- `backend/api/property.go` - Added blockchain calls to UpdatePropertyApproval, improved error handling
- `backend/api/handler.go` - Removed duplicate property functions (kept comments pointing to property.go)

### Frontend
- `frontend/src/lib/api.ts` - Updated createProperty to send multipart form data
- `frontend/src/Pages/AdminCreateProperty.tsx` - Updated to use new API with multiple files
- `frontend/src/Pages/UploadProperty.tsx` - Updated to use new API with multiple files
- `frontend/src/Pages/CreateProperty.tsx` - Updated to use new API with multiple files
- `frontend/src/Pages/ViewProperties.tsx` - Updated to display property name
- `frontend/src/Pages/PropertyDetails.tsx` - Updated to use property name and metadata
- `frontend/src/types/index.ts` - Updated CreatePropertyPayload comment

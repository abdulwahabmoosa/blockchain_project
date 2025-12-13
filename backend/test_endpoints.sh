#!/usr/bin/env bash

# ==========================================
# CONFIGURATION
# ==========================================
BASE_URL="http://localhost:3000"

# 1. Standard User (We create this one fresh every time)
USER_TIMESTAMP=$(date +%s)
USER_EMAIL="user_${USER_TIMESTAMP}@example.com"
USER_PASS="password123"

# 2. Hardcoded Admin (Matches your db.go seedAdmin function)
ADMIN_EMAIL="admin@admin.com"
ADMIN_PASS="password"

# Colors for output
GREEN='\033[0;32m'
CYAN='\033[0;36m'
RED='\033[0;31m'
NC='\033[0m'

echo -e "${CYAN}🚀 Starting System Test...${NC}"
echo "Target: $BASE_URL"

# ==========================================
# 1. HEALTH CHECK
# ==========================================
echo -e "\n${GREEN}[1] Checking Server Health...${NC}"
PING_RES=$(curl -s "$BASE_URL/ping")
if [[ "$PING_RES" != *"pong"* ]]; then
  echo -e "${RED}❌ Server not reachable. Is it running on port 3000?${NC}"
  exit 1
fi
echo "✅ Server is UP."

# ==========================================
# 2. STANDARD USER FLOW
# ==========================================
echo -e "\n${GREEN}[2] Registering Standard User ($USER_EMAIL)...${NC}"
curl -s -X POST "$BASE_URL/register" \
  -H "Content-Type: application/json" \
  -d "{\"email\": \"$USER_EMAIL\", \"password\": \"$USER_PASS\", \"name\": \"Test User\", \"wallet_address\": \"$USER_WALLET\"}" > /dev/null

echo -e "\n${GREEN}[3] Logging in as Standard User...${NC}"
USER_LOGIN=$(curl -s -X POST "$BASE_URL/login" \
  -H "Content-Type: application/json" \
  -d "{\"email\": \"$USER_EMAIL\", \"password\": \"$USER_PASS\"}")

# Extract Token
USER_TOKEN=$(echo $USER_LOGIN | grep -o '"token":"[^"]*' | grep -o '[^"]*$' | tr -d '"')

if [ -z "$USER_TOKEN" ]; then
    echo -e "${RED}❌ Login Failed. Response:${NC} $USER_LOGIN"
    exit 1
fi
echo -e "${CYAN}Token Acquired!${NC}"

# ==========================================
# 3. SECURITY CHECK (Middleware Test)
# ==========================================
echo -e "\n${GREEN}[4] Testing Security (Standard User trying Admin Route)...${NC}"
# User tries to hit the Admin-only approve route
FAIL_RES=$(curl -s -o /dev/null -w "%{http_code}" -X POST "$BASE_URL/approve-user" \
  -H "Authorization: Bearer $USER_TOKEN" \
  -H "Content-Type: application/json" \
  -d "{\"wallet_address\": \"$USER_WALLET\"}")

if [ "$FAIL_RES" -eq 403 ] || [ "$FAIL_RES" -eq 401 ]; then
  echo -e "✅ Security Check Passed: Access Forbidden ($FAIL_RES)"
else
  echo -e "${RED}❌ Security Check Failed: Expected 403, got $FAIL_RES${NC}"
fi

# ==========================================
# 4. ADMIN FLOW
# ==========================================
echo -e "\n${GREEN}[5] Logging in as SEEDED ADMIN...${NC}"
ADMIN_LOGIN=$(curl -s -X POST "$BASE_URL/login" \
  -H "Content-Type: application/json" \
  -d "{\"email\": \"$ADMIN_EMAIL\", \"password\": \"$ADMIN_PASS\"}")

ADMIN_TOKEN=$(echo $ADMIN_LOGIN | grep -o '"token":"[^"]*' | grep -o '[^"]*$' | tr -d '"')

if [ -z "$ADMIN_TOKEN" ]; then
    echo -e "${RED}❌ Admin Login Failed. Did the DB seed run?${NC}"
    echo "Check db.go seedAdmin() credentials."
    exit 1
fi

echo -e "\n${GREEN}[6] Admin Action: Approving User...${NC}"
curl -s -X POST "$BASE_URL/approve-user" \
  -H "Authorization: Bearer $ADMIN_TOKEN" \
  -H "Content-Type: application/json" \
  -d "{\"wallet_address\": \"$USER_WALLET\"}"
echo " (Request Sent)"

# ==========================================
# 5. FILE UPLOAD (IPFS)
# ==========================================
echo -e "\n${GREEN}[7] Uploading Document to IPFS...${NC}"
echo "This is a dummy deed file for testing" > dummy_deed.txt

UPLOAD_RES=$(curl -s -X POST "$BASE_URL/upload" \
  -H "Authorization: Bearer $ADMIN_TOKEN" \
  -F "file=@dummy_deed.txt")

# Extract IPFS Hash
IPFS_HASH=$(echo $UPLOAD_RES | grep -o '"ipfs_hash":"[^"]*' | grep -o '[^"]*$' | tr -d '"')

if [ -z "$IPFS_HASH" ]; then
    echo -e "${RED}❌ Upload Failed. Response:${NC}"
    echo "$UPLOAD_RES"
    # Fallback to keep script running if Pinata isn't set up
    IPFS_HASH="QmDummyHashForTesting"
else
    echo -e "✅ Upload Success! Hash: ${CYAN}$IPFS_HASH${NC}"
fi
rm dummy_deed.txt

# ==========================================
# 6. CREATE PROPERTY
# ==========================================
echo -e "\n${GREEN}[8] Creating Property (Admin)...${NC}"
# Use a random wallet for the owner address field
OWNER_WALLET="0x71C7656EC7ab88b098defB751B7401B5f6d8976F"

CREATE_RES=$(curl -s -X POST "$BASE_URL/properties" \
  -H "Authorization: Bearer $ADMIN_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "owner_address": "'"$OWNER_WALLET"'",
    "name": "Luxury Villa",
    "symbol": "LXV",
    "data_hash": "'"$IPFS_HASH"'",
    "valuation": 500000,
    "token_supply": 100
  }')

echo "Response: $CREATE_RES"

# ==========================================
# 7. READ DATA
# ==========================================
echo -e "\n${GREEN}[9] Fetching All Properties...${NC}"
curl -s -X GET "$BASE_URL/properties" \
  -H "Authorization: Bearer $USER_TOKEN" | head -c 200
echo "..."

echo -e "\n${CYAN}✅ Test Complete.${NC}"

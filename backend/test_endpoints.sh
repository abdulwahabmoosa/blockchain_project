#!/bin/bash

# ==========================================
# CONFIGURATION
# ==========================================
BASE_URL="http://localhost:3000"

# 1. Random User Credentials
USER_TIMESTAMP=$(date +%s)
USER_EMAIL="user_${USER_TIMESTAMP}@example.com"
USER_PASS="password123"
NEW_USER_PASS="newpassword456" # For testing password reset
# Generate a random 42-char Hex Wallet Address
USER_WALLET="0x$(openssl rand -hex 20)"

# 2. Hardcoded Admin (Matches your db.go seedAdmin)
ADMIN_EMAIL="admin@admin.com"
ADMIN_PASS="password"

# Colors
GREEN='\033[0;32m'
CYAN='\033[0;36m'
RED='\033[0;31m'
NC='\033[0m'

echo -e "${CYAN}🚀 Starting Full System Test...${NC}"
echo "Target: $BASE_URL"

# ==========================================
# 1. HEALTH CHECK
# ==========================================
echo -e "\n${GREEN}[1] Checking Server...${NC}"
if curl -s "$BASE_URL/ping" | grep -q "pong"; then
    echo "✅ Server is UP."
else
    echo -e "${RED}❌ Server is DOWN.${NC}"
    exit 1
fi

# ==========================================
# 2. REGISTER USER
# ==========================================
echo -e "\n${GREEN}[2] Registering User ($USER_EMAIL)...${NC}"
REG_RES=$(curl -s -X POST "$BASE_URL/register" \
  -H "Content-Type: application/json" \
  -d '{
    "email": "'"$USER_EMAIL"'",
    "password": "'"$USER_PASS"'",
    "name": "Test User",
    "wallet_address": "'"$USER_WALLET"'"
  }')

if [[ "$REG_RES" == *"token"* ]] || [[ "$REG_RES" == *"success"* ]]; then
    echo "✅ Registration Successful."
else
    echo -e "${RED}❌ Registration Failed:${NC} $REG_RES"
    exit 1
fi

# ==========================================
# 3. USER LOGIN (Get Token)
# ==========================================
echo -e "\n${GREEN}[3] Logging in as User...${NC}"
USER_LOGIN_RES=$(curl -s -X POST "$BASE_URL/login" \
  -H "Content-Type: application/json" \
  -d "{\"email\": \"$USER_EMAIL\", \"password\": \"$USER_PASS\"}")

USER_TOKEN=$(echo $USER_LOGIN_RES | grep -o '"token":"[^"]*' | grep -o '[^"]*$' | tr -d '"')

if [ -z "$USER_TOKEN" ] || [ ${#USER_TOKEN} -lt 10 ]; then
    echo -e "${RED}❌ User Login Failed.${NC}"
    exit 1
fi
echo "✅ User Logged In."

# ==========================================
# 4. ADMIN LOGIN (Get Token)
# ==========================================
echo -e "\n${GREEN}[4] Logging in as Admin...${NC}"
ADMIN_LOGIN_RES=$(curl -s -X POST "$BASE_URL/login" \
  -H "Content-Type: application/json" \
  -d "{\"email\": \"$ADMIN_EMAIL\", \"password\": \"$ADMIN_PASS\"}")

ADMIN_TOKEN=$(echo $ADMIN_LOGIN_RES | grep -o '"token":"[^"]*' | grep -o '[^"]*$' | tr -d '"')

if [ -z "$ADMIN_TOKEN" ] || [ ${#ADMIN_TOKEN} -lt 10 ]; then
    echo -e "${RED}❌ Admin Login Failed.${NC}"
    exit 1
fi
echo "✅ Admin Logged In."

# ==========================================
# 5. APPROVE USER (Admin Action)
# ==========================================
echo -e "\n${GREEN}[5] Admin Approving User...${NC}"
APPROVE_RES=$(curl -s -X POST "$BASE_URL/users/approval" \
  -H "Authorization: Bearer $ADMIN_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "wallet_address": "'"$USER_WALLET"'",
    "status": "approved"
  }')

if [[ "$APPROVE_RES" == *"tx_hash"* ]]; then
    echo "✅ User Approval Tx Sent."
else
    echo -e "${RED}❌ User Approval Failed:${NC} $APPROVE_RES"
fi

# ==========================================
# 6. UPDATE PROFILE (User Action)
# ==========================================
echo -e "\n${GREEN}[6] User Updating Profile (PUT /users/me)...${NC}"
UPDATE_RES=$(curl -s -X PUT "$BASE_URL/users/me" \
  -H "Authorization: Bearer $USER_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"name": "Updated Name Awesome"}')

echo "$UPDATE_RES"
if [[ "$UPDATE_RES" == *"success"* ]]; then
    echo "✅ Profile Updated."
else
    echo -e "${RED}❌ Update Failed.${NC}"
fi

# ==========================================
# 7. RESET PASSWORD (User Action)
# ==========================================
echo -e "\n${GREEN}[7] User Resetting Password...${NC}"
RESET_RES=$(curl -s -X POST "$BASE_URL/users/me/reset-password" \
  -H "Authorization: Bearer $USER_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "old_password": "'"$USER_PASS"'",
    "new_password": "'"$NEW_USER_PASS"'"
  }')

echo "$RESET_RES"
if [[ "$RESET_RES" == *"success"* ]]; then
    echo "✅ Password Reset Successful."
else
    echo -e "${RED}❌ Password Reset Failed.${NC}"
    exit 1
fi

# ==========================================
# 8. VERIFY NEW PASSWORD (Login Again)
# ==========================================
echo -e "\n${GREEN}[8] Verifying New Password Login...${NC}"
NEW_LOGIN_RES=$(curl -s -X POST "$BASE_URL/login" \
  -H "Content-Type: application/json" \
  -d "{\"email\": \"$USER_EMAIL\", \"password\": \"$NEW_USER_PASS\"}")

NEW_USER_TOKEN=$(echo $NEW_LOGIN_RES | grep -o '"token":"[^"]*' | grep -o '[^"]*$' | tr -d '"')

if [ -n "$NEW_USER_TOKEN" ]; then
    echo "✅ Login with NEW password successful."
    USER_TOKEN=$NEW_USER_TOKEN # Update token for next step
else
    echo -e "${RED}❌ Login with new password failed.${NC}"
    exit 1
fi

# ==========================================
# 9. CREATE PROPERTY
# ==========================================
echo -e "\n${GREEN}[9] Creating Property (Multipart)...${NC}"
echo "Deed Content" > deed.pdf
echo "Photo Content" > photo.jpg

PROP_RES=$(curl -s -X POST "$BASE_URL/properties" \
  -H "Authorization: Bearer $ADMIN_TOKEN" \
  -F "files=@deed.pdf" \
  -F "files=@photo.jpg" \
  -F 'data={"owner_address": "'"$USER_WALLET"'", "name": "Sunset Villa", "symbol": "SVL", "valuation": 750000, "token_supply": 1000}')

rm deed.pdf photo.jpg

PROP_ID=$(echo $PROP_RES | grep -o '"property_id":"[^"]*' | grep -o '[^"]*$' | tr -d '"')

if [ -n "$PROP_ID" ] && [ "$PROP_ID" != "null" ]; then
    echo "✅ Property Created: $PROP_ID"
else
    echo -e "${RED}❌ Property Creation Failed.${NC}"
fi

# ==========================================
# 10. DELETE ACCOUNT (User Action)
# ==========================================
echo -e "\n${GREEN}[10] Deleting User Account (Hard Delete)...${NC}"
DELETE_RES=$(curl -s -X DELETE "$BASE_URL/users/me" \
  -H "Authorization: Bearer $USER_TOKEN")

echo "$DELETE_RES"
if [[ "$DELETE_RES" == *"success"* ]]; then
    echo "✅ Account Deleted."
else
    echo -e "${RED}❌ Account Deletion Failed.${NC}"
fi

# ==========================================
# 11. VERIFY DELETION
# ==========================================
echo -e "\n${GREEN}[11] Verifying Deletion (Login Should Fail)...${NC}"
FAIL_LOGIN_RES=$(curl -s -w "%{http_code}" -X POST "$BASE_URL/login" \
  -H "Content-Type: application/json" \
  -d "{\"email\": \"$USER_EMAIL\", \"password\": \"$NEW_USER_PASS\"}")

# We expect a 401 or 404/500 depending on how your error handler wraps "User Not Found"
if [[ "$FAIL_LOGIN_RES" == *"401"* ]] || [[ "$FAIL_LOGIN_RES" == *"404"* ]] || [[ "$FAIL_LOGIN_RES" == *"Unauthorized"* ]]; then
    echo -e "${CYAN}✅ SUCCESS: User cannot login (Account is gone).${NC}"
else
    echo -e "${RED}❌ FAILURE: User still exists or login succeeded.${NC}"
    echo "Response: $FAIL_LOGIN_RES"
fi

import type {
  CreatePropertyPayload,
  LoginResponse,
  Property,
  UploadResponse,
  User,
} from "../types";

// Use backend service name for Docker, localhost for local development
const BASE_URL = import.meta.env.DEV ? "/api" : "/api";

const getHeaders = () => {
  const token = localStorage.getItem("token");
  const headers: HeadersInit = {
    "Content-Type": "application/json",
  };
  if (token) {
    headers["Authorization"] = `Bearer ${token}`;
  }
  return headers;
};

// Retry helper for API calls
const fetchWithRetry = async (
  url: string,
  options: RequestInit,
  maxRetries = 5,
  retryDelay = 2000
): Promise<Response> => {
  for (let i = 0; i < maxRetries; i++) {
    try {
      const response = await fetch(url, options);
      // If it's a 500 error and we're in the first few retries, it might be backend still starting
      if (!response.ok && response.status === 500 && i < maxRetries - 1) {
        // Check if it's a connection-related 500 (backend not ready)
        const text = await response.text();
        if (text.includes("connection") || text.includes("refused") || text.length === 0) {
          console.log(`⏳ Backend may still be starting (500 error), retrying in ${retryDelay}ms... (attempt ${i + 1}/${maxRetries})`);
          await new Promise(resolve => setTimeout(resolve, retryDelay));
          retryDelay *= 1.5; // Exponential backoff
          continue;
        }
      }
      // If it's a connection error (status 0 or network error), retry
      if (!response.ok && response.status === 0) {
        throw new Error("Connection refused");
      }
      return response;
    } catch (error: any) {
      // If it's the last retry, throw the error
      if (i === maxRetries - 1) {
        throw error;
      }
      // If it's a connection error, wait and retry
      if (error.message?.includes("ECONNREFUSED") || 
          error.message?.includes("Connection refused") || 
          error.message?.includes("Failed to fetch") ||
          error.message?.includes("NetworkError") ||
          error.name === "TypeError") {
        console.log(`⏳ Backend not ready, retrying in ${retryDelay}ms... (attempt ${i + 1}/${maxRetries})`);
        await new Promise(resolve => setTimeout(resolve, retryDelay));
        retryDelay *= 1.5; // Exponential backoff
        continue;
      }
      // For other errors, throw immediately
      throw error;
    }
  }
  throw new Error("Max retries exceeded");
};

const handleResponse = async <T>(response: Response): Promise<T> => {
  if (!response.ok) {
    const errorText = await response.text();
    throw new Error(errorText || `HTTP error! status: ${response.status}`);
  }

  // Check for token refresh
  const newToken = response.headers.get("X-Access-Token");
  if (newToken) {
    localStorage.setItem("token", newToken);
  }

  return response.json() as Promise<T>;
};

export const api = {
  login: async (email: string, password: string): Promise<LoginResponse> => {
    const response = await fetchWithRetry(`${BASE_URL}/login`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ email, password }),
    });
    const data = await handleResponse<LoginResponse>(response);
    localStorage.setItem("token", data.token);
    localStorage.setItem("user", JSON.stringify(data.user));
    // Set legacy auth flag for existing code compatibility if needed
    localStorage.setItem("auth", "true");
    return data;
  },

  register: async (
    email: string,
    password: string,
    name: string,
    wallet_address: string
  ): Promise<{ message: string }> => {
    const response = await fetch(`${BASE_URL}/register`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ email, password, name, wallet_address }),
    });
    return handleResponse(response);
  },

  getProperties: async (): Promise<Property[]> => {
    const response = await fetchWithRetry(`${BASE_URL}/properties`, {
      headers: getHeaders(),
    });
    return handleResponse<Property[]>(response);
  },

  getProperty: async (id: string): Promise<Property> => {
    const response = await fetchWithRetry(`${BASE_URL}/properties/${id}`, {
      headers: getHeaders(),
    });
    return handleResponse<Property>(response);
  },

  uploadFile: async (file: File): Promise<UploadResponse> => {
    // Temporarily remove auth requirement for testing
    const formData = new FormData();
    formData.append("file", file);

    const response = await fetch(`${BASE_URL}/upload`, {
      method: "POST",
      headers: {
        // Do not set Content-Type for FormData, browser sets it with boundary
      },
      body: formData,
    });
    return handleResponse<UploadResponse>(response);
  },

  createProperty: async (
    payload: CreatePropertyPayload,
    files: File[]
  ): Promise<{ status: string; tx_hash: string; files_count?: number }> => {
    // Create FormData for multipart request
    const formData = new FormData();
    
    // Add files (backend expects "files" field with multiple files)
    files.forEach((file) => {
      formData.append("files", file);
    });

    // Add JSON payload as "data" field
    formData.append("data", JSON.stringify({
      owner_address: payload.owner_address,
      name: payload.name,
      symbol: payload.symbol,
      valuation: Math.floor(payload.valuation), // Convert to int64
      token_supply: payload.token_supply,
    }));

    // Get auth token for headers (but don't set Content-Type - browser will set it with boundary)
    const token = localStorage.getItem("token");
    const headers: HeadersInit = {};
    if (token) {
      headers["Authorization"] = `Bearer ${token}`;
    }

    const response = await fetch(`${BASE_URL}/properties`, {
      method: "POST",
      headers: headers,
      body: formData,
    });
    return handleResponse(response);
  },

  getCurrentUser: async (): Promise<User> => {
    const response = await fetchWithRetry(`${BASE_URL}/users/me`, {
      headers: getHeaders(),
    });
    return handleResponse<User>(response);
  },

  getUsers: async (): Promise<any[]> => {
    const response = await fetch(`${BASE_URL}/users`, {
      headers: getHeaders(),
    });
    return handleResponse<any[]>(response);
  },

  approveUser: async (
    wallet_address: string
  ): Promise<{ status: string; tx_hash: string; approved?: boolean }> => {
    // Use /users/approval endpoint which waits for mining and updates the database
    const response = await fetch(`${BASE_URL}/users/approval`, {
      method: "POST",
      headers: getHeaders(),
      body: JSON.stringify({ wallet_address, status: "approved" }),
    });
    return handleResponse(response);
  },

  rejectUser: async (
    wallet_address: string
  ): Promise<{ status: string; message: string }> => {
    // Use /users/approval endpoint which waits for mining and updates the database
    const response = await fetch(`${BASE_URL}/users/approval`, {
      method: "POST",
      headers: getHeaders(),
      body: JSON.stringify({ wallet_address, status: "rejected" }),
    });
    return handleResponse(response);
  },

  distributeRevenue: async (
    token_address: string,
    stablecoin_address: string,
    amount: number
  ): Promise<{ status: string; tx_hash: string }> => {
    const response = await fetch(`${BASE_URL}/revenue/distribute`, {
      method: "POST",
      headers: getHeaders(),
      body: JSON.stringify({
        token_address,
        stablecoin_address,
        amount,
      }),
    });
    return handleResponse(response);
  },

  getTokenBalance: async (
    propertyId: string,
    walletAddress: string
  ): Promise<{ balance: string }> => {
    const response = await fetch(`${BASE_URL}/properties/${propertyId}/token-balance/${walletAddress}`, {
      headers: getHeaders(),
    });
    return handleResponse(response);
  },

  getPropertyTokenStats: async (
    propertyId: string
  ): Promise<{
    total: number;
    sold: number;
    available: number;
    percentage_sold: number;
  }> => {
    const response = await fetch(
      `${BASE_URL}/properties/${propertyId}/token-stats`,
      {
        headers: getHeaders(),
      }
    );
    return handleResponse<{
      total: number;
      sold: number;
      available: number;
      percentage_sold: number;
    }>(response);
  },

  createTokenPurchase: async (
    propertyId: string,
    buyerWallet: string,
    amount: string,
    paymentTxHash: string,
    purchasePrice: string
  ): Promise<{ status: string; purchase_id: string; message: string }> => {
    const response = await fetch(`${BASE_URL}/properties/${propertyId}/purchase`, {
      method: "POST",
      headers: getHeaders(),
      body: JSON.stringify({
        buyer_wallet: buyerWallet,
        amount,
        payment_tx_hash: paymentTxHash,
        purchase_price: purchasePrice,
      }),
    });
    return handleResponse(response);
  },

  getPendingTokenPurchases: async (
    propertyId: string
  ): Promise<Array<{
    id: string;
    property_id: string;
    buyer_wallet: string;
    amount: string;
    payment_tx_hash: string;
    token_tx_hash: string;
    purchase_price: string;
    created_at: string;
  }>> => {
    const response = await fetch(
      `${BASE_URL}/properties/${propertyId}/pending-purchases`,
      {
        headers: getHeaders(),
      }
    );
    return handleResponse(response);
  },

  updateTokenPurchaseTxHash: async (
    propertyId: string,
    purchaseId: string,
    tokenTxHash: string
  ): Promise<{ status: string; message: string }> => {
    const response = await fetch(
      `${BASE_URL}/properties/${propertyId}/purchases/${purchaseId}/update-tx`,
      {
        method: "POST",
        headers: getHeaders(),
        body: JSON.stringify({
          token_tx_hash: tokenTxHash,
        }),
      }
    );
    return handleResponse(response);
  },

  getMyTokenPurchases: async (): Promise<Array<{
    id: string;
    property_id: string;
    buyer_wallet: string;
    amount: string;
    payment_tx_hash: string;
    token_tx_hash: string;
    purchase_price: string;
    created_at: string;
  }>> => {
    const response = await fetch(`${BASE_URL}/users/me/purchases`, {
      headers: getHeaders(),
    });
    return handleResponse(response);
  },

  getMyPendingTransfers: async (): Promise<Array<{
    id: string;
    property_id: string;
    buyer_wallet: string;
    amount: string;
    payment_tx_hash: string;
    token_tx_hash: string;
    purchase_price: string;
    created_at: string;
  }>> => {
    const response = await fetch(`${BASE_URL}/users/me/pending-transfers`, {
      headers: getHeaders(),
    });
    return handleResponse(response);
  },

  transferTokens: async (
    propertyId: string,
    toAddress: string,
    amount: string
  ): Promise<{ tx_hash: string; status: string }> => {
    const response = await fetch(`${BASE_URL}/properties/${propertyId}/transfer`, {
      method: "POST",
      headers: getHeaders(),
      body: JSON.stringify({
        to_address: toAddress,
        amount: amount,
      }),
    });
    return handleResponse(response);
  },

  getPropertyMetadata: async (
    propertyId: string
  ): Promise<any> => {
    const response = await fetch(`${BASE_URL}/properties/${propertyId}/metadata`, {
      headers: getHeaders(),
    });
    return handleResponse(response);
  },

  updatePassword: async (
    oldPassword: string,
    newPassword: string
  ): Promise<{ status: string; message: string }> => {
    const response = await fetch(`${BASE_URL}/users/me/reset-password`, {
      method: "POST",
      headers: getHeaders(),
      body: JSON.stringify({
        old_password: oldPassword,
        new_password: newPassword,
      }),
    });
    return handleResponse(response);
  },

  deleteAccount: async (): Promise<{ status: string; message: string }> => {
    const response = await fetch(`${BASE_URL}/users/me`, {
      method: "DELETE",
      headers: getHeaders(),
    });
    return handleResponse(response);
  },

  updatePropertyApproval: async (
    propertyId: string,
    status: "approved" | "rejected"
  ): Promise<{ status: string; new_status: string; property_id: string; tx_hash?: string; message?: string }> => {
    const response = await fetch(`${BASE_URL}/properties/approval`, {
      method: "POST",
      headers: getHeaders(),
      body: JSON.stringify({
        property_id: propertyId,
        status: status,
      }),
    });
    return handleResponse(response);
  },

  // ============ Property Upload Requests ============

  createPropertyUploadRequest: async (
    payload: { name: string; symbol: string; valuation: number; token_supply: number },
    files: File[]
  ): Promise<{ status: string; message: string; request_id: string; files_count: number }> => {
    console.log("📤 Creating property upload request:", payload.name);

    // Create FormData for multipart request
    const formData = new FormData();

    // Add files
    files.forEach((file) => {
      formData.append("files", file);
    });

    // Add JSON payload as "data" field
    formData.append("data", JSON.stringify({
      name: payload.name,
      symbol: payload.symbol,
      valuation: Math.floor(payload.valuation),
      token_supply: payload.token_supply,
    }));

    // Get auth token for headers (but don't set Content-Type - browser will set it with boundary)
    const token = localStorage.getItem("token");
    const headers: HeadersInit = {};
    if (token) {
      headers["Authorization"] = `Bearer ${token}`;
    }

    const response = await fetch(`${BASE_URL}/property-upload-requests`, {
      method: "POST",
      headers: headers,
      body: formData,
    });
    return handleResponse(response);
  },

  getPropertyUploadRequests: async (): Promise<any[]> => {
    console.log("📋 Fetching property upload requests");
    const response = await fetch(`${BASE_URL}/property-upload-requests`, {
      headers: getHeaders(),
    });
    return handleResponse<any[]>(response);
  },

  getMyPropertyUploadRequests: async (): Promise<any[]> => {
    console.log("📋 Fetching my property upload requests");
    const response = await fetch(`${BASE_URL}/property-upload-requests/user`, {
      headers: getHeaders(),
    });
    return handleResponse<any[]>(response);
  },

  getPropertyUploadRequest: async (id: string): Promise<any> => {
    console.log("📋 Fetching property upload request:", id);
    const response = await fetch(`${BASE_URL}/property-upload-requests/${id}`, {
      headers: getHeaders(),
    });
    return handleResponse(response);
  },

  approvePropertyUploadRequest: async (
    id: string
  ): Promise<{ status: string; message: string; request_id: string; property_id: string; tx_hash: string; asset_address: string; token_address: string }> => {
    console.log("✅ Approving property upload request:", id);
    const response = await fetch(`${BASE_URL}/property-upload-requests/${id}/approve`, {
      method: "POST",
      headers: getHeaders(),
    });
    return handleResponse(response);
  },

  rejectPropertyUploadRequest: async (
    id: string,
    reason: string
  ): Promise<{ status: string; message: string; request_id: string }> => {
    console.log("❌ Rejecting property upload request:", id);
    const response = await fetch(`${BASE_URL}/property-upload-requests/${id}/reject`, {
      method: "POST",
      headers: getHeaders(),
      body: JSON.stringify({ reason }),
    });
    return handleResponse(response);
  },
};

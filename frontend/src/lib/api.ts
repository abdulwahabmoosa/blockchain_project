import type {
  CreatePropertyPayload,
  LoginResponse,
  Property,
  UploadResponse,
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
    const response = await fetch(`${BASE_URL}/login`, {
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
    const response = await fetch(`${BASE_URL}/properties`, {
      headers: getHeaders(),
    });
    return handleResponse<Property[]>(response);
  },

  getProperty: async (id: string): Promise<Property> => {
    const response = await fetch(`${BASE_URL}/properties/${id}`, {
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

  getUsers: async (): Promise<any[]> => {
    const response = await fetch(`${BASE_URL}/users`, {
      headers: getHeaders(),
    });
    return handleResponse<any[]>(response);
  },

  approveUser: async (
    wallet_address: string
  ): Promise<{ status: string; tx_hash: string; approved: boolean }> => {
    const response = await fetch(`${BASE_URL}/approve-user`, {
      method: "POST",
      headers: getHeaders(),
      body: JSON.stringify({ wallet_address }),
    });
    return handleResponse(response);
  },

  rejectUser: async (
    wallet_address: string
  ): Promise<{ status: string; message: string }> => {
    const response = await fetch(`${BASE_URL}/approve-user`, {
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
};

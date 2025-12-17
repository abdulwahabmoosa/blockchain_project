import type {
  CreatePropertyPayload,
  LoginResponse,
  Property,
  UploadResponse,
} from "../types";

const BASE_URL = "/api";

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
    const token = localStorage.getItem("token");
    const formData = new FormData();
    formData.append("file", file);

    const response = await fetch(`${BASE_URL}/upload`, {
      method: "POST",
      headers: {
        Authorization: token ? `Bearer ${token}` : "",
        // Do not set Content-Type for FormData, browser sets it with boundary
      },
      body: formData,
    });
    return handleResponse<UploadResponse>(response);
  },

  createProperty: async (
    payload: CreatePropertyPayload
  ): Promise<{ status: string; tx_hash: string }> => {
    const response = await fetch(`${BASE_URL}/properties`, {
      method: "POST",
      headers: getHeaders(),
      body: JSON.stringify(payload),
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
};

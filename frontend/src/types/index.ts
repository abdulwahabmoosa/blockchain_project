export type UserRole = "admin" | "user";

export type PropertyStatus = "Active" | "Paused" | "Disputed" | "Closed";

export interface User {
  ID: string;
  WalletAddress: string;
  Email: string;
  Name: string;
  Role: UserRole;
  IsApproved?: boolean; // Computed from approval_statuss
  approval_status?: "pending" | "approved" | "rejected"; // From backend
  CreatedAt: string;
  UpdatedAt: string;
}

export interface Property {
  ID: string;
  Name?: string;
  OnchainAssetAddress: string;
  OnchainTokenAddress: string;
  OwnerWallet: string;
  MetadataHash: string;
  Valuation: number;
  Status: PropertyStatus;
  CreatedAt: string;
}

export interface PropertyMetadata {
  name: string;
  description: string;
  images: string[];
  location?: string;
  // Add other metadata fields as needed
}

export interface PropertyWithMetadata extends Property {
  metadata?: PropertyMetadata;
}

export interface LoginResponse {
  token: string;
  user: User;
}

export interface UploadResponse {
  ipfs_hash: string;
  url: string;
}

export interface CreatePropertyPayload {
  owner_address: string;
  name: string;
  symbol: string;
  data_hash: string; // Not used in new API - files are sent directly
  valuation: number;
  token_supply: number;
}

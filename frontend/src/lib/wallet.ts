import { ethers } from "ethers";
import contractAddresses from "../deployments/sepolia/contract-addresses.json";

/**
 * Wallet Connection Utility
 * 
 * This module provides functions to connect to MetaMask and interact with the blockchain.
 * It follows the architecture described in README.md where the frontend connects directly
 * to the blockchain using ethers.js, rather than going through the backend.
 */

// Type definition for the window.ethereum object (MetaMask injection)
interface EthereumProvider {
  request: (args: { method: string; params?: any[] }) => Promise<any>;
  on: (event: string, callback: (...args: any[]) => void) => void;
  removeListener: (event: string, callback: (...args: any[]) => void) => void;
  isMetaMask?: boolean;
}

declare global {
  interface Window {
    ethereum?: EthereumProvider;
  }
}

/**
 * Wallet connection state type
 */
export interface WalletState {
  isConnected: boolean;
  address: string | null;
  provider: ethers.BrowserProvider | null;
  signer: ethers.JsonRpcSigner | null;
  chainId: number | null;
}

/**
 * Check if MetaMask is installed
 */
export const isMetaMaskInstalled = (): boolean => {
  return typeof window !== "undefined" && typeof window.ethereum !== "undefined";
};

/**
 * Connect to MetaMask wallet
 * 
 * This function:
 * 1. Checks if MetaMask is installed
 * 2. Creates a BrowserProvider from window.ethereum
 * 3. Requests account access from the user
 * 4. Returns the provider, signer, and connected address
 * 
 * @returns Promise with wallet connection details
 */
export const connectWallet = async (forceSelection: boolean = false): Promise<{
  address: string;
  provider: ethers.BrowserProvider;
  signer: ethers.JsonRpcSigner;
  chainId: number;
}> => {
  // Check if MetaMask is installed
  if (!isMetaMaskInstalled()) {
    throw new Error("MetaMask is not installed. Please install MetaMask to continue.");
  }

  // Create a provider from window.ethereum
  // BrowserProvider is the ethers.js v6 way to connect to browser-based wallets like MetaMask
  const provider = new ethers.BrowserProvider(window.ethereum!);

  if (forceSelection) {
    console.log('🔄 Forcing wallet account selection...');
    // Try to revoke permissions first to force account selection
    try {
      await window.ethereum!.request({
        method: 'wallet_revokePermissions',
        params: [{ eth_accounts: {} }]
      });
      console.log('✅ Wallet permissions revoked, forcing account selection');
    } catch (err) {
      // wallet_revokePermissions might not be supported, continue normally
      console.log('⚠️ Could not revoke permissions, proceeding with normal connection');
    }
  }

  // Request account access from the user
  // This will trigger MetaMask to pop up asking for permission
  // The user must approve this request
  console.log('🔓 Requesting account access...');
  await provider.send("eth_requestAccounts", []);

  // Get the signer (the account that will sign transactions)
  // The signer is connected to the first account returned by MetaMask
  console.log('✍️ Getting signer...');
  const signer = await provider.getSigner();

  // Get the connected wallet address
  console.log('🏠 Getting address...');
  const address = await signer.getAddress();

  // Get the chain ID to verify we're on the correct network
  console.log('🌐 Getting network...');
  const network = await provider.getNetwork();
  const chainId = Number(network.chainId);

  return {
    address,
    provider,
    signer,
    chainId,
  };
};


/**
 * Get the current wallet connection state
 * 
 * This checks if a wallet is already connected without requesting permission
 * Useful for checking connection status on page load
 * 
 * @returns Promise with current wallet state or null if not connected
 */
export const getWalletState = async (): Promise<{
  address: string;
  provider: ethers.BrowserProvider;
  signer: ethers.JsonRpcSigner;
  chainId: number;
} | null> => {
  if (!isMetaMaskInstalled()) {
    console.log('❌ MetaMask not installed');
    return null;
  }

  try {
    console.log('🔗 Creating provider...');
    const provider = new ethers.BrowserProvider(window.ethereum!);

    console.log('👤 Getting accounts...');
    // Get accounts without requesting (will return empty if not connected)
    const accounts = await provider.listAccounts();
    console.log('📋 Accounts found:', accounts.length);

    if (accounts.length === 0) {
      console.log('❌ No accounts connected');
      return null; // No accounts connected
    }

    console.log('✍️ Getting signer...');
    const signer = await provider.getSigner();

    console.log('🏠 Getting address...');
    const address = await signer.getAddress();

    console.log('🌐 Getting network...');
    const network = await provider.getNetwork();
    const chainId = Number(network.chainId);

    console.log('✅ Wallet state retrieved:', { address, chainId });

    return {
      address,
      provider,
      signer,
      chainId,
    };
  } catch (error) {
    console.error("❌ Error getting wallet state:", error);
    return null;
  }
};

/**
 * Disconnect wallet (note: MetaMask doesn't have a true disconnect, but we can clear local state)
 * 
 * @param setWalletState Optional callback to update wallet state in your component
 */
export const disconnectWallet = (): void => {
  // MetaMask doesn't provide a disconnect method
  // We just clear any local state here
  // The actual "disconnect" happens when user switches accounts in MetaMask
  console.log("Wallet disconnected (local state cleared)");
};

/**
 * Switch to Sepolia network
 * 
 * This function attempts to switch the user's MetaMask to Sepolia testnet
 * If the network is not added, it will try to add it
 * 
 * @returns Promise that resolves when network is switched
 */
export const switchToSepolia = async (): Promise<void> => {
  if (!isMetaMaskInstalled()) {
    throw new Error("MetaMask is not installed");
  }

  const sepoliaChainId = "0xaa36a7"; // 11155111 in hex (Sepolia testnet)

  try {
    // Try to switch to Sepolia
    await window.ethereum!.request({
      method: "wallet_switchEthereumChain",
      params: [{ chainId: sepoliaChainId }],
    });
  } catch (switchError: any) {
    // If the network doesn't exist, add it
    if (switchError.code === 4902) {
      try {
        await window.ethereum!.request({
          method: "wallet_addEthereumChain",
          params: [
            {
              chainId: sepoliaChainId,
              chainName: "Sepolia Test Network",
              nativeCurrency: {
                name: "Ether",
                symbol: "ETH",
                decimals: 18,
              },
              rpcUrls: ["https://sepolia.infura.io/v3/"],
              blockExplorerUrls: ["https://sepolia.etherscan.io"],
            },
          ],
        });
      } catch (addError) {
        throw new Error("Failed to add Sepolia network to MetaMask");
      }
    } else {
      throw switchError;
    }
  }
};

/**
 * Verify that the connected wallet matches the user's registered wallet address
 * 
 * This is important for security - we want to make sure the wallet the user
 * connected to MetaMask is the same one they registered with on the platform
 * 
 * @param connectedAddress The address from MetaMask
 * @param registeredAddress The address from the user's account in the database
 * @returns boolean indicating if addresses match
 */
export const verifyWalletMatch = (
  connectedAddress: string,
  registeredAddress: string
): boolean => {
  // Compare addresses (convert to lowercase for case-insensitive comparison)
  return connectedAddress.toLowerCase() === registeredAddress.toLowerCase();
};

/**
 * Get contract addresses from the deployment file
 * 
 * These addresses are imported from the contract deployment JSON file
 * as described in README.md section 4
 */
/**
 * Create admin wallet connection for backend operations
 * This uses the admin private key for direct blockchain operations
 */
export const createAdminWalletConnection = async () => {
  try {
    // Use localhost for development, Sepolia for production
    const isDevelopment = import.meta.env.DEV;
    const rpcUrl = isDevelopment
      ? 'http://localhost:8545'  // Local Hardhat node
      : 'https://ethereum-sepolia-rpc.publicnode.com';  // Sepolia for production

    const provider = new ethers.JsonRpcProvider(rpcUrl);

    // Admin private key
    const adminPrivateKey = '0x65f24585f0bafc504a5d88a5cc4b7eb8b10ff71aab7f36284f16161414c15b6b';
    const adminWallet = new ethers.Wallet(adminPrivateKey, provider);

    console.log(`🔑 Admin wallet connected (${isDevelopment ? 'localhost' : 'Sepolia'}):`, adminWallet.address);

    return {
      address: adminWallet.address,
      provider: provider,
      signer: adminWallet,
      chainId: isDevelopment ? 31337 : 11155111, // Local Hardhat (31337) or Sepolia (11155111)
    };
  } catch (error) {
    console.error('❌ Failed to create admin wallet connection:', error);
    throw error;
  }
};

/**
 * Get admin wallet state (always available for admin operations)
 */
export const getAdminWalletState = async () => {
  return await createAdminWalletConnection();
};

export const getContractAddresses = () => {
  const isDevelopment = import.meta.env.DEV;

  if (isDevelopment) {
    // For local development, use localhost deployment
    console.log('🏠 Using localhost contract addresses for development');
    try {
      const localhostAddresses = require("../deployments/localhost/contract-addresses.json");
      return localhostAddresses;
    } catch (error) {
      console.warn('⚠️ Localhost contract addresses not found, using placeholders');
      return {
        network: "localhost",
        adminAddress: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
        PlatformRegistry: "0x5FbDB2315678afecb367f032d93F642f64180aa3",
        ApprovalService: "0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512",
        PropertyFactory: "0xa513E6E4b8f2a923D98304ec87F64353C4D5C853",
        RevenueDistribution: "0x8A791620dd6260079BF849Dc5567aDC3F2FdC318",
        "PropertyAsset (ERC721)": "0x0000000000000000000000000000000000000000",
        "PropertyToken (ERC20)": "0x0000000000000000000000000000000000000000",
      };
    }
  }

  console.log('🌐 Using Sepolia contract addresses for production');
  return contractAddresses;
};

/**
 * Listen for account changes in MetaMask
 * 
 * This sets up event listeners to detect when the user switches accounts
 * or when they disconnect their wallet
 * 
 * @param onAccountChange Callback when account changes
 * @param onDisconnect Callback when wallet disconnects
 * @returns Function to remove the listeners
 */
export const setupWalletListeners = (
  onAccountChange: (accounts: string[]) => void,
  onDisconnect: () => void
): (() => void) => {
  if (!isMetaMaskInstalled()) {
    return () => {}; // Return empty cleanup function
  }

  const handleAccountsChanged = (accounts: string[]) => {
    if (accounts.length === 0) {
      onDisconnect();
    } else {
      onAccountChange(accounts);
    }
  };

  const handleDisconnect = () => {
    onDisconnect();
  };

  // Set up listeners
  window.ethereum!.on("accountsChanged", handleAccountsChanged);
  window.ethereum!.on("disconnect", handleDisconnect);

  // Return cleanup function to remove listeners
  return () => {
    if (window.ethereum) {
      window.ethereum.removeListener("accountsChanged", handleAccountsChanged);
      window.ethereum.removeListener("disconnect", handleDisconnect);
    }
  };
};








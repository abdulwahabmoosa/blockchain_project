import { createContext, useContext, useEffect, useState } from "react";
import type { ReactNode } from "react";
import { ethers } from "ethers";
import type { BrowserProvider, JsonRpcSigner } from "ethers";
import {
  connectWallet,
  getWalletState,
  isMetaMaskInstalled,
  setupWalletListeners,
  switchToSepolia,
  verifyWalletMatch,
} from "../lib/wallet";
import { getTokenBalance } from "../lib/contracts";
import { api } from "../lib/api";
import { formatTokenBalance } from "../utils/format";

interface WalletContextType {
  isConnected: boolean;
  address: string | null;
  provider: BrowserProvider | null;
  signer: JsonRpcSigner | null;
  chainId: number | null;
  isMetaMaskAvailable: boolean;
  isLoading: boolean;
  error: string | null;
  totalTokenBalance: string | null;
  connect: () => Promise<void>;
  disconnect: () => Promise<void>;
  resetLoading: () => void;
  checkWallet: () => Promise<void>;
  switchNetwork: () => Promise<void>;
  verifyWallet: (registeredAddress: string) => boolean;
  connectRegisteredWallet: (registeredAddress: string) => Promise<void>;
  clearError: () => void;
  updateTotalTokenBalance: () => Promise<void>;
}

const WalletContext = createContext<WalletContextType | undefined>(undefined);

export const useWallet = () => {
  const context = useContext(WalletContext);
  if (context === undefined) {
    throw new Error("useWallet must be used within a WalletProvider");
  }
  return context;
};

interface WalletProviderProps {
  children: ReactNode;
}

export const WalletProvider = ({ children }: WalletProviderProps) => {
  const [isConnected, setIsConnected] = useState(false);
  const [address, setAddress] = useState<string | null>(null);
  const [provider, setProvider] = useState<BrowserProvider | null>(null);
  const [signer, setSigner] = useState<JsonRpcSigner | null>(null);
  const [chainId, setChainId] = useState<number | null>(null);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [totalTokenBalance, setTotalTokenBalance] = useState<string | null>(null);
  const [isMetaMaskAvailable, setIsMetaMaskAvailable] = useState(false);
  const [wrongWalletConnected, setWrongWalletConnected] = useState(false);

  const checkWallet = async () => {
    console.log('🔍 Checking wallet state...');
    setIsLoading(true);
    setError(null);
    try {
      const available = isMetaMaskInstalled();
      setIsMetaMaskAvailable(available);
      console.log('📱 MetaMask available:', available);

      if (!available) {
        console.log('❌ MetaMask not available');
        setIsConnected(false);
        return;
      }

      // Don't auto-reconnect if user explicitly disconnected or wrong wallet was connected
      const wasDisconnected = localStorage.getItem('wallet_disconnected') === 'true';
      console.log('🔌 Was disconnected:', wasDisconnected);
      console.log('❌ Wrong wallet flag:', wrongWalletConnected);
      if (wasDisconnected || wrongWalletConnected) {
        console.log('🚫 Skipping auto-reconnect due to previous disconnect or wrong wallet');
        setIsConnected(false);
        setAddress(null);
        setProvider(null);
        setSigner(null);
        setChainId(null);
        return;
      }

      console.log('⏳ Getting wallet state...');
      // Add timeout to prevent hanging
      const statePromise = getWalletState();
      const timeoutPromise = new Promise<null>((_, reject) =>
        setTimeout(() => reject(new Error('Wallet state check timeout')), 5000)
      );

      const state = await Promise.race([statePromise, timeoutPromise]);
      console.log('✅ Wallet state result:', state ? 'connected' : 'not connected');

      if (state) {
        // Double-check that this wallet matches the registered wallet before connecting
        const userStr = localStorage.getItem("user");
        const user = userStr ? JSON.parse(userStr) : null;
        const registeredWallet = user?.WalletAddress;

        if (registeredWallet && state.address.toLowerCase() !== registeredWallet.toLowerCase()) {
          console.log('⚠️ Found connected wallet but it\'s not the registered one, skipping auto-connect');
          setWrongWalletConnected(true);
          setIsConnected(false);
          setAddress(null);
          setProvider(null);
          setSigner(null);
          setChainId(null);
          return;
        }

        setIsConnected(true);
        setAddress(state.address);
        setProvider(state.provider);
        setSigner(state.signer);
        setChainId(state.chainId);

        // Update total token balance after successful wallet state check
        await updateTotalTokenBalance();
      } else {
        setIsConnected(false);
        setAddress(null);
        setProvider(null);
        setSigner(null);
        setChainId(null);
      }
    } catch (err: any) {
      console.error('❌ Error checking wallet:', err);
      setError(err.message);
      setIsConnected(false);
    } finally {
      console.log('🏁 Wallet check completed');
      setIsLoading(false);
    }
  };

  const updateTotalTokenBalance = async () => {
    if (!provider || !address) {
      setTotalTokenBalance(null);
      return;
    }

    try {
      console.log('🔄 Calculating total token balance...');
      // Get user's properties
      const properties = await api.getProperties();

      let totalBalance = 0n;

      // Sum balances from all property tokens
      for (const property of properties) {
        if (property.OnchainTokenAddress) {
          try {
            const balance = await getTokenBalance(
              property.OnchainTokenAddress,
              address,
              provider
            );
            totalBalance += balance;
          } catch (err) {
            console.warn(`Failed to get balance for property ${property.ID}:`, err);
          }
        }
      }

      // Format and set the total balance
      const formattedBalance = formatTokenBalance(totalBalance);
      setTotalTokenBalance(formattedBalance);
      console.log('✅ Total token balance updated:', formattedBalance);
    } catch (err: any) {
      console.error('❌ Failed to update total token balance:', err);
      setTotalTokenBalance("0");
    }
  };

  const connect = async () => {
    setError(null);
    setIsLoading(true);
    console.log('🔗 Starting wallet connection...');

    try {
      // Get user data to check registered wallet
      const userStr = localStorage.getItem("user");
      const user = userStr ? JSON.parse(userStr) : null;
      const registeredWallet = user?.WalletAddress;

      // Force account selection if we were previously disconnected
      const wasDisconnected = localStorage.getItem('wallet_disconnected') === 'true';
      console.log('🔄 Force selection:', wasDisconnected);
      console.log('👤 User registered wallet:', registeredWallet);

      // Add timeout to prevent hanging
      const connectPromise = connectWallet(wasDisconnected);
      const timeoutPromise = new Promise<never>((_, reject) =>
        setTimeout(() => reject(new Error('Wallet connection timeout')), 15000)
      );

      const { address, provider, signer, chainId } = await Promise.race([connectPromise, timeoutPromise]);
      console.log('✅ Wallet connected:', address);

      setIsConnected(true);
      setAddress(address);
      setProvider(provider);
      setSigner(signer);
      setChainId(chainId);

      // Clear the disconnected flag since user manually connected
      localStorage.removeItem('wallet_disconnected');

      // Check if connected wallet matches registered wallet
      if (registeredWallet && address.toLowerCase() !== registeredWallet.toLowerCase()) {
        console.log('⚠️ Connected to different wallet than registered');
        console.log('Registered:', registeredWallet);
        console.log('Connected:', address);

        // Set flag to prevent auto-reconnection and disconnect
        setWrongWalletConnected(true);
        await disconnect(); // Force disconnect

        throw new Error(`Wrong wallet connected. Please connect with your registered wallet address: ${registeredWallet.substring(0, 6)}...${registeredWallet.substring(38)}`);
      } else {
        // Clear the wrong wallet flag on successful connection
        setWrongWalletConnected(false);
        // Check if on Sepolia, if not, prompt to switch
        if (chainId !== 11155111) {
          setError("Please switch to Sepolia network");
        }

        // Update total token balance after successful connection
        await updateTotalTokenBalance();
      }
    } catch (err: any) {
      console.error('❌ Wallet connection failed:', err);
      setError(err.message || "Failed to connect wallet");
      setIsConnected(false);
    } finally {
      setIsLoading(false);
    }
  };

  const connectRegisteredWallet = async (registeredAddress: string) => {
    setError(null);
    setIsLoading(true);
    console.log('🔗 Auto-connecting to registered wallet:', registeredAddress);

    try {
      if (!isMetaMaskInstalled()) {
        throw new Error("MetaMask is not installed");
      }

      // Validate that user is logged in and has a registered wallet
      const userStr = localStorage.getItem("user");
      const user = userStr ? JSON.parse(userStr) : null;

      if (!user || !user.WalletAddress) {
        throw new Error("Please log in to connect your wallet");
      }

      // Ensure the requested address matches the user's registered address
      if (registeredAddress.toLowerCase() !== user.WalletAddress.toLowerCase()) {
        throw new Error("Wallet address mismatch. Please use your registered wallet.");
      }

      // Create provider
      const provider = new ethers.BrowserProvider(window.ethereum!);
      console.log('🔌 Provider created');

      // Always request account access to ensure MetaMask opens and user can select account
      console.log('🔓 Requesting account access to ensure MetaMask opens...');
      const accounts = await provider.send("eth_requestAccounts", []);
      console.log('✅ Accounts after request:', accounts);

      if (accounts.length === 0) {
        throw new Error("No accounts available. Please connect your wallet.");
      }

      // Check if the registered address is among available accounts
      const matchingAccount = accounts.find((acc: string) => acc.toLowerCase() === registeredAddress.toLowerCase());
      console.log('🔍 Looking for registered address:', registeredAddress.toLowerCase());
      console.log('📋 Available accounts (lowercase):', accounts.map((acc: string) => acc.toLowerCase()));
      console.log('✅ Matching account found:', matchingAccount);

      // Get the signer for the currently selected account
      const signer = await provider.getSigner();
      const currentAddress = await signer.getAddress();
      const network = await provider.getNetwork();
      const chainId = Number(network.chainId);
      console.log('🌐 Network:', chainId);
      console.log('👤 Current selected account:', currentAddress);

      if (matchingAccount || currentAddress.toLowerCase() === registeredAddress.toLowerCase()) {
        // Either found in available accounts or user selected it manually
        console.log('✅ Using registered wallet account');
        setIsConnected(true);
        setAddress(currentAddress);
        setProvider(provider);
        setSigner(signer);
        setChainId(chainId);

        // Clear the disconnected flag and wrong wallet flag
        localStorage.removeItem('wallet_disconnected');
        setWrongWalletConnected(false);

        // Check if on Sepolia
        if (chainId !== 11155111) {
          setError("Please switch to Sepolia network");
        }

        // Update total token balance after successful connection
        await updateTotalTokenBalance();
      } else {
        // Registered wallet not selected - user needs to switch manually
        console.log('⚠️ Registered wallet not selected');
        throw new Error(`Please select your registered wallet address in MetaMask: ${registeredAddress}. Click the account selector in MetaMask and choose your registered account.`);
      }
    } catch (err: any) {
      console.error('❌ Registered wallet connection failed:', err);
      setError(err.message || "Failed to connect to registered wallet");
      setIsConnected(false);
      // Re-throw to let the calling component handle it
      throw err;
    } finally {
      setIsLoading(false);
    }
  };


  const disconnect = async () => {
    console.log('🔌 Starting wallet disconnect...');

    // First clear the app state to prevent any race conditions
    setIsConnected(false);
    setAddress(null);
    setProvider(null);
    setSigner(null);
    setChainId(null);
    setError(null);
    setTotalTokenBalance(null);
    setIsLoading(false);
    setWrongWalletConnected(false);

    // Clear any cached wallet state to force fresh connection
    localStorage.removeItem('wallet_connected');
    localStorage.removeItem('wallet_address');
    localStorage.removeItem('wallet_chainId');

    // Set disconnected flag to prevent auto-reconnection
    localStorage.setItem('wallet_disconnected', 'true');

    try {
      // Try to revoke MetaMask permissions if possible
      if (window.ethereum) {
        await window.ethereum.request({
          method: 'wallet_revokePermissions',
          params: [{ eth_accounts: {} }],
        });
        console.log('✅ MetaMask permissions revoked');
      }
    } catch (err) {
      console.log('⚠️ Could not revoke MetaMask permissions:', err);
      // Continue even if permission revocation fails
    }

    console.log('✅ Wallet disconnected - app state cleared and permissions revoked');
  };

  const resetLoading = () => {
    console.log('🔄 Manually resetting loading state');
    setIsLoading(false);
    setError(null);
    setIsConnected(false);
    setAddress(null);
    setProvider(null);
    setSigner(null);
    setChainId(null);
    setWrongWalletConnected(false);
    // Clear all wallet-related localStorage
    localStorage.removeItem('wallet_disconnected');
    localStorage.removeItem('wallet_connected');
    localStorage.removeItem('wallet_address');
    localStorage.removeItem('wallet_chainId');
  };

  const switchNetwork = async () => {
    setError(null);
    try {
      await switchToSepolia();
      // Refresh wallet state after network switch
      setTimeout(() => {
        checkWallet();
      }, 1000);
    } catch (err: any) {
      setError(err.message || "Failed to switch network");
    }
  };

  const verifyWallet = (registeredAddress: string): boolean => {
    if (!address) return false;
    return verifyWalletMatch(address, registeredAddress);
  };

  useEffect(() => {
    let mounted = true;

    const initWallet = async () => {
      // Only check wallet if we haven't been explicitly disconnected
      const wasDisconnected = localStorage.getItem('wallet_disconnected') === 'true';
      console.log('🚀 Initializing wallet, wasDisconnected:', wasDisconnected);

      if (!wasDisconnected && mounted) {
        await checkWallet();
      } else if (mounted) {
        // If disconnected, ensure loading is false
        setIsLoading(false);
      }
    };

    initWallet();

    // Set up listeners for account changes
    if (isMetaMaskInstalled()) {
      const cleanup = setupWalletListeners(
        async (accounts) => {
          if (accounts.length > 0) {
            // Clear the disconnected flag when user manually connects
            localStorage.removeItem('wallet_disconnected');
            checkWallet();
          } else {
            await disconnect();
          }
        },
        async () => {
          await disconnect();
        }
      );

      return () => {
        mounted = false;
        cleanup();
      };
    }

    return () => {
      mounted = false;
    };
  }, []);

  const clearError = () => {
    setError(null);
  };

  const value: WalletContextType = {
    isConnected,
    address,
    provider,
    signer,
    chainId,
    isMetaMaskAvailable,
    isLoading,
    error,
    totalTokenBalance,
    connect,
    disconnect,
    resetLoading,
    checkWallet,
    switchNetwork,
    verifyWallet,
    connectRegisteredWallet,
    clearError,
    updateTotalTokenBalance,
  };

  return <WalletContext.Provider value={value}>{children}</WalletContext.Provider>;
};






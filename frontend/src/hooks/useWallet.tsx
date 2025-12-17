import { createContext, useContext, useEffect, useState } from "react";
import type { ReactNode } from "react";
import type { BrowserProvider, JsonRpcSigner } from "ethers";
import {
  connectWallet,
  getWalletState,
  isMetaMaskInstalled,
  setupWalletListeners,
  switchToSepolia,
  verifyWalletMatch,
} from "../lib/wallet";

interface WalletContextType {
  isConnected: boolean;
  address: string | null;
  provider: BrowserProvider | null;
  signer: JsonRpcSigner | null;
  chainId: number | null;
  isMetaMaskAvailable: boolean;
  isLoading: boolean;
  error: string | null;
  connect: () => Promise<void>;
  disconnect: () => void;
  resetLoading: () => void;
  checkWallet: () => Promise<void>;
  switchNetwork: () => Promise<void>;
  verifyWallet: (registeredAddress: string) => boolean;
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
  const [isMetaMaskAvailable, setIsMetaMaskAvailable] = useState(false);

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

      // Don't auto-reconnect if user explicitly disconnected
      const wasDisconnected = localStorage.getItem('wallet_disconnected') === 'true';
      console.log('🔌 Was disconnected:', wasDisconnected);
      if (wasDisconnected) {
        console.log('🚫 Skipping auto-reconnect due to previous disconnect');
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
        setIsConnected(true);
        setAddress(state.address);
        setProvider(state.provider);
        setSigner(state.signer);
        setChainId(state.chainId);
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

  const connect = async () => {
    setError(null);
    setIsLoading(true);
    console.log('🔗 Starting wallet connection...');

    try {
      // Force account selection if we were previously disconnected
      const wasDisconnected = localStorage.getItem('wallet_disconnected') === 'true';
      console.log('🔄 Force selection:', wasDisconnected);

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

      // Check if on Sepolia, if not, prompt to switch
      if (chainId !== 11155111) {
        setError("Please switch to Sepolia network");
      }
    } catch (err: any) {
      console.error('❌ Wallet connection failed:', err);
      setError(err.message || "Failed to connect wallet");
      setIsConnected(false);
    } finally {
      setIsLoading(false);
    }
  };


  const disconnect = () => {
    // Clear all wallet state
    setIsConnected(false);
    setAddress(null);
    setProvider(null);
    setSigner(null);
    setChainId(null);
    setError(null);
    setIsLoading(false);

    // Clear any cached wallet state to force fresh connection
    localStorage.removeItem('wallet_connected');
    localStorage.removeItem('wallet_address');
    localStorage.removeItem('wallet_chainId');

    // Set disconnected flag to prevent auto-reconnection
    localStorage.setItem('wallet_disconnected', 'true');

    console.log('Wallet disconnected - will require manual reconnection');
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
        (accounts) => {
          if (accounts.length > 0) {
            // Clear the disconnected flag when user manually connects
            localStorage.removeItem('wallet_disconnected');
            checkWallet();
          } else {
            disconnect();
          }
        },
        () => {
          disconnect();
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

  const value: WalletContextType = {
    isConnected,
    address,
    provider,
    signer,
    chainId,
    isMetaMaskAvailable,
    isLoading,
    error,
    connect,
    disconnect,
    resetLoading,
    checkWallet,
    switchNetwork,
    verifyWallet,
  };

  return <WalletContext.Provider value={value}>{children}</WalletContext.Provider>;
};





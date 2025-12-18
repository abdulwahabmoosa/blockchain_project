import React, { createContext, useContext, useState, useEffect } from 'react';
import type { ReactNode } from 'react';
import { ethers } from 'ethers';
import {
  connectWallet,
  getWalletState,
  disconnectWallet as disconnectWalletUtil,
  setupWalletListeners,
  switchToSepolia,
  verifyWalletMatch
} from '../lib/wallet';
import type { WalletState } from '../lib/wallet';

/**
 * Wallet context interface
 */
interface WalletContextType extends WalletState {
  connect: (forceSelection?: boolean) => Promise<void>;
  disconnect: () => void;
  connectRegisteredWallet: (registeredAddress: string) => Promise<void>;
  switchToSepolia: () => Promise<void>;
  verifyWalletMatch: (registeredAddress: string) => boolean;
}

/**
 * Wallet context
 */
const WalletContext = createContext<WalletContextType | undefined>(undefined);

/**
 * Wallet provider props
 */
interface WalletProviderProps {
  children: ReactNode;
}

/**
 * Wallet provider component
 */
export const WalletProvider: React.FC<WalletProviderProps> = ({ children }) => {
  const [isConnected, setIsConnected] = useState(false);
  const [address, setAddress] = useState<string | null>(null);
  const [provider, setProvider] = useState<ethers.BrowserProvider | null>(null);
  const [signer, setSigner] = useState<ethers.JsonRpcSigner | null>(null);
  const [chainId, setChainId] = useState<number | null>(null);

  // Initialize wallet state on mount
  useEffect(() => {
    const initializeWallet = async () => {
      try {
        const walletState = await getWalletState();
        if (walletState) {
          setIsConnected(true);
          setAddress(walletState.address);
          setProvider(walletState.provider);
          setSigner(walletState.signer);
          setChainId(walletState.chainId);
        }
      } catch (error) {
        console.error('Failed to initialize wallet:', error);
      }
    };

    initializeWallet();
  }, []);

  // Set up wallet listeners
  useEffect(() => {
    const cleanup = setupWalletListeners(
      (accounts) => {
        if (accounts.length === 0) {
          // Disconnected
          setIsConnected(false);
          setAddress(null);
          setProvider(null);
          setSigner(null);
          setChainId(null);
        } else {
          // Account changed - reinitialize
          getWalletState().then(walletState => {
            if (walletState) {
              setIsConnected(true);
              setAddress(walletState.address);
              setProvider(walletState.provider);
              setSigner(walletState.signer);
              setChainId(walletState.chainId);
            }
          }).catch(console.error);
        }
      },
      () => {
        // Disconnected
        setIsConnected(false);
        setAddress(null);
        setProvider(null);
        setSigner(null);
        setChainId(null);
      }
    );

    return cleanup;
  }, []);

  // Connect wallet
  const connect = async (forceSelection: boolean = false) => {
    try {
      const walletState = await connectWallet(forceSelection);
      setIsConnected(true);
      setAddress(walletState.address);
      setProvider(walletState.provider);
      setSigner(walletState.signer);
      setChainId(walletState.chainId);
    } catch (error) {
      console.error('Failed to connect wallet:', error);
      throw error;
    }
  };

  // Connect to a specific registered wallet address
  const connectRegisteredWallet = async (registeredAddress: string) => {
    try {
      // First try to connect normally
      await connect();

      // Then verify the connected address matches the registered one
      if (address && !verifyWalletMatch(address, registeredAddress)) {
        throw new Error(`Connected wallet (${address}) doesn't match registered wallet (${registeredAddress}). Please connect with the correct wallet.`);
      }
    } catch (error) {
      console.error('Failed to connect registered wallet:', error);
      throw error;
    }
  };

  // Disconnect wallet
  const disconnect = () => {
    disconnectWalletUtil();
    setIsConnected(false);
    setAddress(null);
    setProvider(null);
    setSigner(null);
    setChainId(null);
  };

  // Switch to Sepolia network
  const switchToSepoliaNetwork = async () => {
    try {
      await switchToSepolia();
      // Update chain ID after switching
      if (provider) {
        const network = await provider.getNetwork();
        setChainId(Number(network.chainId));
      }
    } catch (error) {
      console.error('Failed to switch to Sepolia:', error);
      throw error;
    }
  };

  // Verify wallet match
  const checkWalletMatch = (registeredAddress: string): boolean => {
    return address ? verifyWalletMatch(address, registeredAddress) : false;
  };

  const contextValue: WalletContextType = {
    isConnected,
    address,
    provider,
    signer,
    chainId,
    connect,
    disconnect,
    connectRegisteredWallet,
    switchToSepolia: switchToSepoliaNetwork,
    verifyWalletMatch: checkWalletMatch,
  };

  return (
    <WalletContext.Provider value={contextValue}>
      {children}
    </WalletContext.Provider>
  );
};

/**
 * useWallet hook
 *
 * Provides access to wallet state and methods throughout the app
 */
export const useWallet = (): WalletContextType => {
  const context = useContext(WalletContext);
  if (context === undefined) {
    throw new Error('useWallet must be used within a WalletProvider');
  }
  return context;
};
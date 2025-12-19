import { useCallback, useEffect, useState } from "react";
import { Outlet, useNavigate } from "react-router-dom";
import { Sidebar } from "./Sidebar";
import { useWallet } from "../../hooks/useWallet";
import { api } from "../../lib/api";
import { Button } from "../atoms/Button";
import type { User } from "../../types";

export const DashboardLayout: React.FC = () => {
  const { disconnect, address, verifyWalletMatch, connect, isLoading: walletLoading, isManuallyConnected, provider } = useWallet();
  const navigate = useNavigate();
  const [user, setUser] = useState<User | null>(null);
  const [loading, setLoading] = useState(true);
  const [walletMismatch, setWalletMismatch] = useState(false);
  const [connectError, setConnectError] = useState<string | null>(null);
  const [justConnected, setJustConnected] = useState(false);
  const [showWalletModal, setShowWalletModal] = useState(false);
  const [hasAttemptedConnection, setHasAttemptedConnection] = useState(false);

  // Fetch user from database (not localStorage)
  useEffect(() => {
    const fetchUser = async () => {
      try {
        const userData = await api.getCurrentUser();
        setUser(userData);
        if (userData.Role === "admin") {
          navigate("/admin");
          return;
        }
      } catch (err) {
        console.error("Failed to fetch user from database:", err);
        // Fallback to localStorage if API fails
        const userStr = localStorage.getItem("user");
        if (userStr) {
          const userData = JSON.parse(userStr);
          setUser(userData);
          if (userData.Role === "admin") {
            navigate("/admin");
            return;
          }
        } else {
          navigate("/login");
          return;
        }
      } finally {
        setLoading(false);
      }
    };
    fetchUser();
  }, [navigate]);

  // Handle wallet connection with validation
  const handleConnectWallet = async () => {
    if (!user || !user.WalletAddress) {
      setConnectError("User data not loaded. Please refresh the page.");
      return;
    }

    setConnectError(null);
    setJustConnected(true);
    try {
      // Connect wallet (opens MetaMask for user to select account)
      await connect(true); // forceSelection = true to allow account selection
      
      // Wait a moment for address to update, then verify
      setTimeout(() => {
        if (address && !verifyWalletMatch(address, user.WalletAddress)) {
          setConnectError(`Wrong wallet! You connected ${address} but your registered wallet is ${user.WalletAddress}. Please try again.`);
          disconnect();
          setJustConnected(false);
        }
      }, 500);
    } catch (err: any) {
      setConnectError(err.message || "Failed to connect wallet. Please try again.");
      console.error("Failed to connect wallet:", err);
      setJustConnected(false);
    }
  };

  // Auto-prompt wallet connection on first load (after login)
  useEffect(() => {
    if (!user || !user.WalletAddress || loading || hasAttemptedConnection) return;
    
    // If wallet is not connected, show modal to force connection
    if (!isManuallyConnected) {
      setShowWalletModal(true);
      setHasAttemptedConnection(true);
    }
  }, [user, loading, isManuallyConnected, hasAttemptedConnection]);

  // Validate wallet when it's connected
  useEffect(() => {
    if (!user || !user.WalletAddress || loading) return;

    // Only check if wallet is already connected
    if (address) {
      const matches = verifyWalletMatch(address, user.WalletAddress);
      if (!matches) {
        setWalletMismatch(true);
        setShowWalletModal(true);
        // If we just connected and it's wrong, show error and disconnect
        if (justConnected) {
          setConnectError(`Wrong wallet! You connected ${address} but your registered wallet is ${user.WalletAddress}. Please connect with the correct wallet.`);
          disconnect();
          setJustConnected(false);
        }
      } else {
        setWalletMismatch(false);
        setConnectError(null);
        setJustConnected(false);
        setShowWalletModal(false); // Close modal when correct wallet is connected
      }
    } else {
      setWalletMismatch(false);
      setJustConnected(false);
    }
  }, [user, address, verifyWalletMatch, loading, justConnected, disconnect]);

  const handleLogout = useCallback(async () => {
    // Reuse existing auth clearing semantics
    localStorage.removeItem("auth");
    localStorage.removeItem("token");
    localStorage.removeItem("user");
    await disconnect();
    window.location.href = "/";
  }, [disconnect]);

  if (loading || !user) {
    return (
      <div className="min-h-screen w-full bg-[#0f0f0f] text-white flex items-center justify-center">
        <p className="text-gray-400">Loading dashboard...</p>
      </div>
    );
  }

  return (
    <div className="min-h-screen w-full bg-[#0f0f0f] text-white">
      <Sidebar onLogout={handleLogout} />

      {/* Main content area */}
      <main className="md:ml-64 min-h-screen">
        {/* Connect Wallet Button - Show when not manually connected OR wallet doesn't match */}
        {(!isManuallyConnected || walletMismatch) && (
          <div className="mx-4 mt-4 p-4 bg-blue-500/10 border border-blue-500/20 rounded-xl">
            <div className="flex items-center justify-between">
              <p className="text-blue-400 text-sm">
                {!isManuallyConnected 
                  ? "Please connect your wallet to access all features."
                  : "Please connect with the correct wallet to access all features."}
              </p>
              <Button
                variant="primary"
                size="sm"
                onClick={handleConnectWallet}
                disabled={walletLoading}
                className="bg-[#6d41ff] hover:bg-[#5b2fff] text-white"
              >
                {walletLoading ? "Connecting..." : "Connect Wallet"}
              </Button>
            </div>
          </div>
        )}

        {/* Modal for wallet connection (mandatory) */}
        {showWalletModal && (
          <div className="fixed inset-0 bg-black/80 z-50 flex items-center justify-center p-4">
            <div className="bg-[#111111] border border-[#1f1f1f] rounded-2xl p-6 max-w-md w-full space-y-4">
              <h2 className="text-xl font-semibold">Connect Your Wallet</h2>
              <p className="text-sm text-gray-400">
                You must connect your registered wallet to access the dashboard.
              </p>
              {user?.WalletAddress && (
                <div className="p-3 bg-[#0f0f0f] rounded-lg border border-[#1f1f1f]">
                  <p className="text-xs text-gray-500">Registered Wallet:</p>
                  <p className="font-mono text-sm mt-1">{user.WalletAddress}</p>
                </div>
              )}
              
              {connectError && (
                <div className="p-3 bg-red-500/10 border border-red-500/20 rounded-lg">
                  <p className="text-red-400 text-sm">{connectError}</p>
                </div>
              )}
              
              <div className="flex gap-3">
                <Button
                  variant="primary"
                  size="md"
                  onClick={handleConnectWallet}
                  disabled={walletLoading}
                  className="flex-1 bg-[#6d41ff] hover:bg-[#5b2fff] text-white"
                >
                  {walletLoading ? "Connecting..." : "Connect Wallet"}
                </Button>
                {!isManuallyConnected && (
                  <Button
                    variant="outline"
                    size="md"
                    onClick={() => {
                      // Logout if user refuses to connect
                      handleLogout();
                    }}
                    className="border-[#262626] text-gray-400 hover:bg-[#1a1a1a]"
                  >
                    Logout
                  </Button>
                )}
              </div>
            </div>
          </div>
        )}

        {/* Error popup for wrong wallet (non-modal) */}
        {connectError && !showWalletModal && (
          <div className="mx-4 mt-4 p-4 bg-red-500/10 border border-red-500/20 rounded-xl">
            <div className="flex items-center justify-between">
              <p className="text-red-400 text-sm">{connectError}</p>
              <button
                onClick={() => {
                  setConnectError(null);
                  setShowWalletModal(true); // Show modal to retry
                }}
                className="text-red-400 hover:text-red-300 text-lg font-bold"
              >
                ×
              </button>
            </div>
            <Button
              variant="outline"
              size="sm"
              onClick={handleConnectWallet}
              disabled={walletLoading}
              className="mt-2 border-red-500/50 bg-red-500/10 text-red-400 hover:bg-red-500/20"
            >
              {walletLoading ? "Connecting..." : "Try Again"}
            </Button>
          </div>
        )}

        {/* Warning for wallet mismatch */}
        {walletMismatch && address && (
          <div className="mx-4 mt-4 p-4 bg-yellow-500/10 border border-yellow-500/20 rounded-xl">
            <div className="flex items-center justify-between">
              <p className="text-yellow-400 text-sm">
                ⚠️ Connected wallet ({address}) doesn't match your registered wallet ({user.WalletAddress}). 
                Please connect with the correct wallet.
              </p>
              <Button
                variant="outline"
                size="sm"
                onClick={handleConnectWallet}
                disabled={walletLoading}
                className="border-yellow-500/50 bg-yellow-500/10 text-yellow-400 hover:bg-yellow-500/20"
              >
                {walletLoading ? "Connecting..." : "Reconnect"}
              </Button>
            </div>
          </div>
        )}

        <div className="max-w-6xl mx-auto px-4 py-8 pt-16 md:pt-8">
          {/* Only show content if wallet is manually connected and matches */}
          {isManuallyConnected && !walletMismatch ? (
            <Outlet />
          ) : (
            <div className="text-center py-16">
              <p className="text-gray-400 text-lg">
                Please connect your wallet to view your dashboard.
              </p>
            </div>
          )}
        </div>
      </main>
    </div>
  );
};




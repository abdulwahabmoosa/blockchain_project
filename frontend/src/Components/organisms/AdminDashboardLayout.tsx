import { useCallback, useEffect, useState } from "react";
import { Outlet, useNavigate } from "react-router-dom";
import { AdminSidebar } from "./AdminSidebar";
import type { User } from "../../types";
import { useWallet } from "../../hooks/useWallet";
import { api } from "../../lib/api";
import { Button } from "../atoms/Button";

export const AdminDashboardLayout: React.FC = () => {
  const [user, setUser] = useState<User | null>(null);
  const [loading, setLoading] = useState(true);
  const navigate = useNavigate();
  const { disconnect, address, verifyWalletMatch, connect, isLoading: walletLoading, isManuallyConnected, provider } = useWallet();
  const [walletMismatch, setWalletMismatch] = useState(false);
  const [connectError, setConnectError] = useState<string | null>(null);
  const [justConnected, setJustConnected] = useState(false);

  // Fetch user from database (not localStorage)
  useEffect(() => {
    const fetchUser = async () => {
      try {
        const userData = await api.getCurrentUser();
        setUser(userData);
        if (userData.Role !== "admin") {
          navigate("/dashboard");
          return;
        }
      } catch (err) {
        console.error("Failed to fetch user from database:", err);
        // Fallback to localStorage if API fails
        const userStr = localStorage.getItem("user");
        if (userStr) {
          const userData = JSON.parse(userStr);
          setUser(userData);
          if (userData.Role !== "admin") {
            navigate("/dashboard");
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
    } catch (err: any) {
      setConnectError(err.message || "Failed to connect wallet. Please try again.");
      console.error("Failed to connect wallet:", err);
      setJustConnected(false);
    }
  };

  // Validate wallet when it's connected (but don't auto-connect)
  useEffect(() => {
    if (!user || !user.WalletAddress || loading) return;

    // Only check if wallet is already connected
    if (address) {
      const matches = verifyWalletMatch(address, user.WalletAddress);
      if (!matches) {
        setWalletMismatch(true);
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
      }
    } else {
      setWalletMismatch(false);
      setJustConnected(false);
    }
  }, [user, address, verifyWalletMatch, loading, justConnected, disconnect]);

  const handleLogout = useCallback(async () => {
    localStorage.removeItem("auth");
    localStorage.removeItem("token");
    localStorage.removeItem("user");
    await disconnect();
    navigate("/login");
  }, [disconnect, navigate]);

  if (loading || !user) {
    return (
      <div className="min-h-screen w-full bg-[#0f0f0f] text-white flex items-center justify-center">
        <p className="text-gray-400">Loading admin dashboard...</p>
      </div>
    );
  }

  return (
    <div className="min-h-screen w-full bg-[#0f0f0f] text-white">
      <AdminSidebar onLogout={handleLogout} />

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

        {/* Error popup for wrong wallet */}
        {connectError && (
          <div className="mx-4 mt-4 p-4 bg-red-500/10 border border-red-500/20 rounded-xl">
            <div className="flex items-center justify-between">
              <p className="text-red-400 text-sm">{connectError}</p>
              <button
                onClick={() => setConnectError(null)}
                className="text-red-400 hover:text-red-300 text-lg font-bold"
              >
                ×
              </button>
            </div>
          </div>
        )}

        {/* Warning for wallet mismatch */}
        {walletMismatch && address && (
          <div className="mx-4 mt-4 p-4 bg-yellow-500/10 border border-yellow-500/20 rounded-xl">
            <div className="flex items-center justify-between">
              <p className="text-yellow-400 text-sm">
                ⚠️ Connected wallet ({address}) doesn't match your registered wallet ({user?.WalletAddress}). 
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



import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { Navbar } from "../Components/organisms/Navbar";
import { Button } from "../Components/atoms/Button";
import { api } from "../lib/api";
import type { Property, User } from "../types";
import { useWallet } from "../hooks/useWallet";
import { checkApprovalStatus, getTokenBalance } from "../lib/contracts";
import { formatTokenBalance } from "../utils/format";
import contractAddresses from "../deployments/sepolia/contract-addresses.json";

const Card = ({
  title,
  value,
  sub,
  accent = "bg-blue-500",
}: {
  title: string;
  value: string;
  sub?: string;
  accent?: string;
}) => (
  <div className="bg-[#1A1A1A] border border-[#262626] rounded-xl p-6">
    <div className="flex items-center gap-3 mb-4">
      <p className="text-sm text-gray-400">{title}</p>
      <div className={`h-8 w-8 rounded-full ${accent}`} />
    </div>
    <p className="text-2xl font-semibold">{value}</p>
    {sub && <p className="text-xs text-gray-400">{sub}</p>}
  </div>
);

interface PropertyWithBalance extends Property {
  tokenBalance?: string;
  balanceLoading?: boolean;
}

function AdminDashboard() {
  const [properties, setProperties] = useState<PropertyWithBalance[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");
  const [totalBalance, setTotalBalance] = useState<string>("0");
  const [totalBalanceLoading, setTotalBalanceLoading] = useState(false);
  const [isApproved, setIsApproved] = useState<boolean | null>(null);
  const [totalUsers, setTotalUsers] = useState<number>(0);
  const [user, setUser] = useState<User | null>(null);
  const [autoConnectAttempted, setAutoConnectAttempted] = useState(false);
  const navigate = useNavigate();
  const { isConnected, address, provider, connectRegisteredWallet } = useWallet();

  // Check if user is admin
  useEffect(() => {
    const userStr = localStorage.getItem("user");
    if (userStr) {
      const userData = JSON.parse(userStr);
      setUser(userData);
      if (userData.Role !== "admin") {
        navigate("/dashboard"); // Redirect non-admins to regular dashboard
        return;
      }
    }
  }, [navigate]);

  // Auto-connect wallet for admin users (only attempt once)
  useEffect(() => {
    const autoConnectWallet = async () => {
      if (user && user.WalletAddress && !autoConnectAttempted) {
        setAutoConnectAttempted(true);

        // Check if we're already connected to the correct wallet
        const isCorrectWallet = address && address.toLowerCase() === user.WalletAddress.toLowerCase();

        if (!isCorrectWallet) {
          console.log('🔗 Auto-connecting admin wallet for user:', user.Email);
          try {
            await connectRegisteredWallet(user.WalletAddress);
            console.log('✅ Admin wallet auto-connection successful');
          } catch (err) {
            console.error('❌ Admin wallet auto-connect failed:', err);
            // Don't retry, just log the error
          }
        } else {
          console.log('✅ Admin already connected to correct wallet');
        }
      }
    };

    autoConnectWallet();
  }, [user, autoConnectAttempted, address, connectRegisteredWallet]);

  useEffect(() => {
    const fetchProperties = async () => {
      try {
        const data = await api.getProperties();
        const propertiesWithLoading = data.map((p) => ({
          ...p,
          balanceLoading: true,
          tokenBalance: undefined,
        }));
        setProperties(propertiesWithLoading);
      } catch (err: any) {
        console.error("Failed to fetch properties:", err);
        setError("Failed to load properties");
      } finally {
        setLoading(false);
      }
    };

    fetchProperties();
  }, []);

  // Fetch total user count (excluding admins)
  useEffect(() => {
    const fetchUserCount = async () => {
      try {
        const users = await api.getUsers();
        // Filter out admin users, only count regular users
        const regularUsers = users.filter((u: any) => u.Role !== 'admin');
        setTotalUsers(regularUsers.length);
      } catch (err) {
        console.error("Failed to fetch user count:", err);
        setTotalUsers(0);
      }
    };

    fetchUserCount();
  }, []);

  // Check approval status when wallet is connected
  useEffect(() => {
    const checkApproval = async () => {
      if (isConnected && address && provider) {
        try {
          const approved = await checkApprovalStatus(address, provider);
          setIsApproved(approved);
        } catch (err) {
          console.error("Failed to check approval status:", err);
          setIsApproved(null);
        }
      } else {
        setIsApproved(null);
      }
    };

    checkApproval();
  }, [isConnected, address, provider]);

  // Also check approval status when component mounts (in case user just got approved)
  useEffect(() => {
    if (user && isConnected && address && provider) {
      checkApprovalStatus(address, provider).then(setIsApproved).catch(() => setIsApproved(null));
    }
  }, [user]); // This will trigger when user data changes

  // Fetch token balances when wallet is connected and properties are loaded
  useEffect(() => {
    if (!isConnected || !address || !provider || properties.length === 0) {
      return;
    }

    const fetchBalances = async () => {
      setTotalBalanceLoading(true);
      let total = 0n;

      const updatedProperties = await Promise.all(
        properties.map(async (property) => {
          try {
            const balance = await getTokenBalance(
              property.OnchainTokenAddress,
              address,
              provider
            );
            total += balance;
            return {
              ...property,
              tokenBalance: formatTokenBalance(balance),
              balanceLoading: false,
            };
          } catch (err) {
            console.error(`Failed to fetch balance for property ${property.ID}:`, err);
            return {
              ...property,
              tokenBalance: "0",
              balanceLoading: false,
            };
          }
        })
      );

      setProperties(updatedProperties);
      setTotalBalance(formatTokenBalance(total));
      setTotalBalanceLoading(false);
    };

    fetchBalances();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [isConnected, address, provider]);

  if (!user || user.Role !== "admin") {
    return null; // Will redirect in useEffect
  }

  return (
    <div className="min-h-screen w-full bg-[#0f0f0f] text-white">
      <Navbar currentPage="/admin" />

      <main className="max-w-6xl mx-auto px-4 pt-24 pb-16 space-y-8">
        <header className="flex items-center justify-between">
          <div>
            <p className="text-sm text-gray-400">Admin Panel</p>
            <h1 className="text-3xl font-semibold">Dashboard</h1>
          </div>
          <Button
            variant="primary"
            size="md"
            className="bg-[#6d41ff] hover:bg-[#5b2fff] text-white"
            onClick={() => navigate("/create-property")}
          >
            Create Property
          </Button>
        </header>

        {/* Admin Stats */}
        <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
          <Card
            title="Total Properties"
            value={properties.length.toString()}
            accent="bg-blue-500"
          />
          <Card
            title="Your Token Holdings"
            value={totalBalanceLoading ? "Loading..." : totalBalance}
            sub="Across all properties"
            accent="bg-green-500"
          />
          <Card
            title="Total registered users"
            value={totalUsers.toString()}
            accent="bg-purple-500"
          />
        </div>

        {/* Approval Status Section */}
        {isConnected && (
          <div className="bg-[#1A1A1A] border border-[#262626] rounded-xl p-6">
            <div className="flex items-center justify-between">
              <div className="space-y-1">
                <h3 className="text-lg font-semibold">Your Admin Status</h3>
                <p className="text-sm text-gray-400">
                  {isApproved === null
                    ? "Checking approval status..."
                    : isApproved
                    ? `✅ Your wallet (${address?.substring(0, 6)}...${address?.substring(38)}) is approved and can perform all admin functions`
                    : `❌ Your wallet (${address?.substring(0, 6)}...${address?.substring(38)}) is not approved. You can still use admin panel but cannot receive tokens.`}
                </p>
                {!isApproved && (
                  <div className="text-xs text-yellow-400 mt-1 space-y-1">
                    <p>💡 Go to "Manage User Approvals" to approve your wallet address</p>
                    <p><strong>Expected admin address:</strong> {contractAddresses.adminAddress}</p>
                    {address && (
                      <p><strong>Your connected address:</strong> {address}</p>
                    )}
                  </div>
                )}
              </div>
            </div>
          </div>
        )}

        {/* Admin Actions */}
        <div className="bg-[#1A1A1A] border border-[#262626] rounded-xl p-6">
          <h3 className="text-lg font-semibold mb-4">Admin Actions</h3>
          <div className="grid grid-cols-1 md:grid-cols-1 gap-4 max-w-xs">
            <Button
              variant="outline"
              size="md"
              onClick={() => navigate("/admin/approve-users")}
              className="border-[#262626] bg-[#0f0f0f] text-white hover:bg-[#262626]"
            >
              Manage User Approvals
            </Button>
          </div>
        </div>

        {error && (
          <div className="text-red-500 bg-red-500/10 p-4 rounded-xl border border-red-500/20">
            {error}
          </div>
        )}

        {/* Properties Section */}
        <div className="space-y-4">
          <h2 className="text-xl font-semibold">Platform Properties</h2>

          {loading ? (
            <div className="text-center py-8">
              <p>Loading properties...</p>
            </div>
          ) : properties.length === 0 ? (
            <div className="text-center py-8">
              <p className="text-gray-400">No properties found.</p>
            </div>
          ) : (
            <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
              {properties.map((property) => (
                <div
                  key={property.ID}
                  className="bg-[#1A1A1A] border border-[#262626] rounded-lg p-6 hover:border-[#6d41ff]/50 transition-colors"
                >
                  <div className="flex items-start justify-between mb-4">
                    <div className="space-y-1">
                      <h3 className="font-semibold">
                        Property #{property.ID.substring(0, 8)}
                      </h3>
                      <p className="text-sm text-gray-400">
                        Owner: {property.OwnerWallet.substring(0, 6)}...{property.OwnerWallet.substring(38)}
                      </p>
                    </div>
                    <div className="text-right">
                      <p className="text-sm font-semibold">
                        {property.Valuation} ETH
                      </p>
                      <p className="text-xs text-gray-400">
                        {property.Status}
                      </p>
                    </div>
                  </div>

                  {isConnected && address && (
                    <div className="flex justify-between items-center text-sm mb-4">
                      <span className="text-gray-400">Your holdings:</span>
                      <span className="font-semibold">
                        {property.balanceLoading ? "Loading..." : property.tokenBalance || "0"} tokens
                      </span>
                    </div>
                  )}

                  <div className="flex gap-2">
                    <Button
                      variant="outline"
                      size="sm"
                      onClick={() => navigate(`/properties/${property.ID}`)}
                      className="flex-1 border-[#262626] bg-[#0f0f0f] text-white hover:bg-[#262626]"
                    >
                      View Details
                    </Button>
                    <Button
                      variant="outline"
                      size="sm"
                      onClick={() => navigate(`/admin/property/${property.ID}`)}
                      className="flex-1 border-[#262626] bg-[#0f0f0f] text-white hover:bg-[#262626]"
                    >
                      Admin View
                    </Button>
                  </div>
                </div>
              ))}
            </div>
          )}
        </div>
      </main>
    </div>
  );
}

export default AdminDashboard;

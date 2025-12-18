import "../index.css";
import { useEffect, useState } from "react";
import { Navbar } from "../Components/organisms/Navbar";
import { Button } from "../Components/atoms/Button";
import { api } from "../lib/api";
import type { Property } from "../types";
import { useNavigate } from "react-router-dom";
import { useWallet } from "../hooks/useWallet";
import { getTokenBalance, checkApprovalStatus, getClaimableDistributions, claimRevenue } from "../lib/contracts";
import { formatTokenBalance } from "../utils/format";

const Card = ({
  title,
  value,
  sub,
  accent = "bg-[#6d41ff]",
}: {
  title: string;
  value: string;
  sub?: string;
  accent?: string;
}) => (
  <div className="rounded-2xl border border-[#1f1f1f] bg-[#111111] p-4 space-y-2">
    <div className="flex items-center justify-between">
      <p className="text-sm text-gray-400">{title}</p>
      <div className={`h-8 w-8 rounded-full ${accent}`} />
    </div>
    <p className="text-2xl font-semibold">{value}</p>
    {sub && <p className="text-xs text-gray-400">{sub}</p>}
  </div>
);

const GradientThumb = () => (
  <div className="rounded-2xl border border-[#1f1f1f] bg-linear-to-br from-[#191919] via-[#1f1f1f] to-[#121212] w-full h-full" />
);

interface PropertyWithBalance extends Property {
  tokenBalance?: string;
  balanceLoading?: boolean;
  claimableDistributions?: Array<{
    distributionId: number;
    amount: bigint;
    stablecoin: string;
  }>;
  claiming?: boolean;
}

function DashboardPage() {
  const [properties, setProperties] = useState<PropertyWithBalance[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");
  const [totalBalance, setTotalBalance] = useState<string>("0");
  const [totalBalanceLoading, setTotalBalanceLoading] = useState(false);
  const [isApproved, setIsApproved] = useState<boolean | null>(null);
  const [isCheckingApproval, setIsCheckingApproval] = useState(false);
  const [user, setUser] = useState<any>(null);
  const [autoConnectAttempted, setAutoConnectAttempted] = useState(false);
  const navigate = useNavigate();
  const { isConnected, address, provider, signer, connectRegisteredWallet } = useWallet();

  // Debug wallet state
  useEffect(() => {
    console.log("🔗 Wallet state:", {
      isConnected,
      address: address?.substring(0, 10) + "...",
      hasProvider: !!provider,
      hasSigner: !!signer
    });
  }, [isConnected, address, provider, signer]);


  // Check if user is admin and redirect if needed
  useEffect(() => {
    const userStr = localStorage.getItem("user");
    if (userStr) {
      const userData = JSON.parse(userStr);
      setUser(userData);
      if (userData.Role === "admin") {
        navigate("/admin"); // Redirect admins to admin dashboard
        return;
      }
    }
  }, [navigate]);

  const handleClaimRevenue = async (propertyId: string, distributionId: number) => {
    if (!signer) {
      setError("Wallet not connected");
      return;
    }

    // Set claiming state for this property
    setProperties(prev => prev.map(p =>
      p.ID === propertyId ? { ...p, claiming: true } : p
    ));

    try {
      const tx = await claimRevenue(distributionId, signer);
      await tx.wait();

      // Refresh claimable distributions for this property
      if (address && provider) {
        const property = properties.find(p => p.ID === propertyId);
        if (property) {
          const claimable = await getClaimableDistributions(
            property.OnchainTokenAddress,
            address,
            provider
          );

          setProperties(prev => prev.map(p =>
            p.ID === propertyId
              ? { ...p, claimableDistributions: claimable, claiming: false }
              : p
          ));

          // Also refresh token balance
          const balance = await getTokenBalance(
            property.OnchainTokenAddress,
            address,
            provider
          );
          setProperties(prev => prev.map(p =>
            p.ID === propertyId
              ? { ...p, tokenBalance: formatTokenBalance(balance) }
              : p
          ));
        }
      }

      alert("Revenue claimed successfully!");
    } catch (err: any) {
      setError(err.message || "Failed to claim revenue");
    } finally {
      setProperties(prev => prev.map(p =>
        p.ID === propertyId ? { ...p, claiming: false } : p
      ));
    }
  };

  // Auto-connect wallet for registered users (only attempt once)
  useEffect(() => {
    const autoConnectWallet = async () => {
      if (user && user.WalletAddress && !autoConnectAttempted) {
        setAutoConnectAttempted(true);

        console.log('🔍 Auto-connect check:', { user: user?.Email, walletAddress: user?.WalletAddress, currentAddress: address, isConnected });

        // Check if we're already connected to the correct wallet
        const isCorrectWallet = address && address.toLowerCase() === user.WalletAddress.toLowerCase();

        console.log('🔗 Wallet match check:', { isCorrectWallet, expected: user.WalletAddress, current: address });

        if (!isCorrectWallet) {
          console.log('🔗 Auto-connecting registered wallet for user:', user.Email);
          try {
            await connectRegisteredWallet(user.WalletAddress);
            console.log('✅ Auto-connection successful');
          } catch (err) {
            console.error('❌ Auto-connect failed:', err);
            // Don't retry, just log the error
          }
        } else {
          console.log('✅ Already connected to correct wallet');
        }
      } else {
        console.log('⏭️ Skipping auto-connect: no user or wallet address');
      }
    };

    autoConnectWallet();
  }, [user, autoConnectAttempted, address, connectRegisteredWallet]);

  useEffect(() => {
    const fetchProperties = async () => {
      try {
        console.log("🔍 Fetching properties from API...");
        const data = await api.getProperties();
        console.log("📊 Properties received:", data.length, data);
        const propertiesWithLoading = data.map((p) => ({
          ...p,
          balanceLoading: true,
          tokenBalance: undefined,
        }));
        setProperties(propertiesWithLoading);
        console.log("✅ Properties loaded:", propertiesWithLoading.length);
      } catch (err: any) {
        console.error("❌ Failed to fetch properties:", err);
        setError("Failed to load properties");
      } finally {
        setLoading(false);
      }
    };

    fetchProperties();
  }, []);

  // Check approval status when wallet is connected
  useEffect(() => {
    const checkApproval = async () => {
      if (isConnected && address && provider) {
        setIsCheckingApproval(true);
        try {
          // Add timeout to prevent hanging
          const approvalPromise = checkApprovalStatus(address, provider);
          const timeoutPromise = new Promise<boolean>((_, reject) =>
            setTimeout(() => reject(new Error('Approval check timeout')), 15000)
          );

          const approved = await Promise.race([approvalPromise, timeoutPromise]);
          setIsApproved(approved);
        } catch (err) {
          console.error("Failed to check approval status:", err);
          // On timeout or error, set to null and let user know they can still use the app
          setIsApproved(null);
        } finally {
          setIsCheckingApproval(false);
        }
      } else {
        setIsApproved(null);
        setIsCheckingApproval(false);
      }
    };

    checkApproval();
  }, [isConnected, address, provider]);

  // Fetch token balances when wallet is connected and properties are loaded
  useEffect(() => {
    const fetchBalances = async () => {
      if (!isConnected || !address || !provider || properties.length === 0) {
        console.log("⏭️ Skipping balance fetch - not connected or no properties", { isConnected, address: address?.substring(0, 10), propertiesCount: properties.length });
        return;
      }

      // Skip if already loading or if balances are already loaded
      const needsLoading = properties.some((p) => p.balanceLoading === true);
      if (!needsLoading) {
        console.log("⏭️ Skipping balance fetch - all balances already loaded");
        return;
      }

      console.log("🔄 Fetching token balances for", properties.length, "properties");
      console.log("👤 User address:", address);
      setTotalBalanceLoading(true);
      let total = 0n;

      // Fetch balance for each property
      const updatedProperties = await Promise.all(
        properties.map(async (property) => {
          console.log(`🔍 Checking balance for property ${property.ID} (${property.OnchainTokenAddress})`);

          // If already loaded, reuse the value
          if (property.tokenBalance !== undefined && !property.balanceLoading) {
            try {
              const balance = parseFloat(property.tokenBalance.replace(/[^\d.]/g, "")) * 10**18;
              total += BigInt(Math.floor(balance));
              console.log(`✅ Reusing existing balance for ${property.ID}: ${property.tokenBalance}`);
            } catch {
              console.log(`⚠️ Failed to parse existing balance for ${property.ID}`);
            }
            return property;
          }

          try {
            console.log(`📡 Calling getTokenBalance for ${property.ID}...`);
            const balance = await getTokenBalance(
              property.OnchainTokenAddress,
              address,
              provider
            );
            console.log(`✅ Balance for ${property.ID}:`, balance.toString(), "wei");
            total += balance;

            // If user has tokens in this property, fetch claimable distributions
            let claimableDistributions = undefined;
            if (balance > 0n) {
              console.log(`💰 User has tokens in ${property.ID}, fetching claimable distributions...`);
              try {
                claimableDistributions = await getClaimableDistributions(
                  property.OnchainTokenAddress,
                  address,
                  provider
                );
                console.log(`✅ Found ${claimableDistributions.length} claimable distributions for ${property.ID}`);
              } catch (claimErr) {
                console.error(`❌ Failed to fetch claimable distributions for property ${property.ID}:`, claimErr);
              }
            } else {
              console.log(`ℹ️ User has no tokens in ${property.ID}`);
            }

            return {
              ...property,
              tokenBalance: formatTokenBalance(balance),
              balanceLoading: false,
              claimableDistributions,
            };
          } catch (err) {
            console.error(`❌ Failed to fetch balance for property ${property.ID}:`, err);
            return {
              ...property,
              tokenBalance: "0",
              balanceLoading: false,
              claimableDistributions: [],
            };
          }
        })
      );

      console.log("🎯 Total balance calculated:", total.toString(), "wei");
      setProperties(updatedProperties);
      setTotalBalance(formatTokenBalance(total));
      setTotalBalanceLoading(false);
      console.log("✅ Balance fetching completed");
    };

    fetchBalances();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [isConnected, address, provider]);

  return (
    <div className="min-h-screen w-full bg-[#0f0f0f] text-white">
      <Navbar currentPage="/dashboard" />
      <main className="max-w-6xl mx-auto px-4 pt-24 pb-16 space-y-8">
        <header className="flex items-center justify-between">
          <div>
            <p className="text-sm text-gray-400">Welcome back</p>
            <h1 className="text-3xl font-semibold">Dashboard</h1>
          </div>
          <Button
            variant="primary"
            size="md"
            className="bg-[#6d41ff] hover:bg-[#5b2fff] text-white"
            onClick={() => navigate("/create-property")}
          >
            New Investment
          </Button>
        </header>

        {/* Approval Status Section */}
        {isConnected && (
          <div className="bg-[#1A1A1A] border border-[#262626] rounded-xl p-6">
            <div className="flex items-center justify-between">
              <div className="space-y-1">
                <h3 className="text-lg font-semibold">Investment Status</h3>
                <p className="text-sm text-gray-400">
                  {isCheckingApproval
                    ? "⏳ Checking approval status... (may take 10-15 seconds on slow networks)"
                    : isApproved === null
                    ? "⚠️ Unable to check approval status. You can still browse properties."
                    : isApproved
                    ? "✅ You are approved to invest in properties and receive tokens"
                    : "❌ You need approval to invest. Contact an administrator."}
                </p>
                {isCheckingApproval && (
                  <p className="text-xs text-yellow-400">
                    If this takes too long, you can still browse properties but may need approval to invest.
                  </p>
                )}
              </div>
              <div className="flex gap-2">
                {isApproved === null && !isCheckingApproval && (
                  <Button
                    variant="outline"
                    size="sm"
                    onClick={async () => {
                      setIsCheckingApproval(true);
                      // Re-run the approval check
                      if (isConnected && address && provider) {
                        try {
                          const approved = await checkApprovalStatus(address, provider);
                          setIsApproved(approved);
                        } catch (err) {
                          console.error("Failed to refresh approval status:", err);
                          setIsApproved(null);
                        } finally {
                          setIsCheckingApproval(false);
                        }
                      }
                    }}
                    className="border-blue-500/50 bg-blue-500/10 text-blue-400 hover:bg-blue-500/20"
                    disabled={isCheckingApproval}
                  >
                    {isCheckingApproval ? "⏳ Checking..." : "🔄 Refresh"}
                  </Button>
                )}
                {isApproved === false && (
                  <Button
                    variant="outline"
                    size="sm"
                    onClick={() => window.location.href = "mailto:admin@propchain.com?subject=Request Property Investment Approval"}
                    className="border-yellow-500/50 bg-yellow-500/10 text-yellow-400 hover:bg-yellow-500/20"
                  >
                    Request Approval
                  </Button>
                )}
              </div>
            </div>
          </div>
        )}

        {error && (
          <div className="text-red-500 bg-red-500/10 p-4 rounded-xl border border-red-500/20">
            {error}
          </div>
        )}

        {/* Top grid */}
        <section className="grid md:grid-cols-3 gap-4">
          {/* Active Properties Highlight */}
          <div className="md:col-span-2 rounded-3xl border border-[#1f1f1f] bg-[#111111] p-5 space-y-4">
            <div className="flex items-center justify-between">
              <p className="text-sm text-gray-400 uppercase tracking-[0.2em]">
                Active Properties
              </p>
              <span className="text-xs text-gray-500">Latest</span>
            </div>
            <div className="grid md:grid-cols-2 gap-4">
              {properties.length > 0 ? (
                <div className="space-y-3">
                  <div className="h-48 rounded-2xl border border-[#1f1f1f] overflow-hidden">
                    <GradientThumb />
                  </div>
                  <div className="space-y-1">
                    <h3 className="text-xl font-semibold">Property #{properties[0].ID.substring(0, 8)}</h3>
                    <p className="text-sm text-gray-400">
                      {properties[0].Valuation} ETH · {properties[0].Status}
                    </p>
                  </div>
                </div>
              ) : (
                <div className="space-y-3 flex items-center justify-center border border-[#1f1f1f] rounded-2xl h-full min-h-[200px]">
                   <p className="text-gray-500">No properties available</p>
                </div>
              )}

              <div className="space-y-3">
                <div className="grid grid-cols-2 gap-3">
                  <Card title="Total Properties" value={properties.length.toString()} sub="Platform" />
                  <Card title="Active" value={properties.filter(p => p.Status === "Active").length.toString()} sub="Status" accent="bg-[#2e2e2e]" />
                  <Card title="Avg Price" value="-- ETH" sub="Est." accent="bg-[#2c1f5e]" />
                  <Card title="Watchlist" value="0" sub="Saved" accent="bg-[#1f3a5e]" />
                </div>
                <Button
                  variant="outline"
                  size="md"
                  className="border-[#262626] bg-[#111111] text-white hover:bg-[#1a1a1a] w-full"
                  onClick={() => navigate("/properties")}
                >
                  View All Properties
                </Button>
              </div>
            </div>
          </div>

          {/* Balance card */}
          <div className="rounded-3xl border border-[#1f1f1f] bg-linear-to-br from-[#6d41ff] via-[#5a3ccc] to-[#1b132f] p-5 space-y-3">
            <p className="text-sm text-white/80">Total Token Holdings</p>
            <p className="text-3xl font-semibold">
              {isConnected && address
                ? totalBalanceLoading
                  ? "Loading..."
                  : `${totalBalance} Tokens`
                : "Connect Wallet"}
            </p>
            <p className="text-sm text-white/80">Properties Owned</p>
            <p className="text-2xl font-semibold">
              {isConnected && address
                ? properties.filter((p) => p.tokenBalance && parseFloat(p.tokenBalance) > 0).length
                : "--"}
            </p>
            <div className="flex gap-2">
              {isConnected && address && (
                <span className="rounded-full bg-white/10 px-3 py-1 text-xs">
                  {address.substring(0, 6)}...{address.substring(38)}
                </span>
              )}
            </div>
            {!isConnected && (
              <p className="text-xs text-white/60 text-center">
                Connect your wallet to see token balances
              </p>
            )}
          </div>
        </section>

        {/* Mid grid */}
        <section className="grid md:grid-cols-3 gap-4">
          {/* Properties List */}
          <div className="rounded-3xl border border-[#1f1f1f] bg-[#111111] p-5 space-y-4">
            <div className="flex items-center justify-between">
              <p className="font-semibold">Recent Properties</p>
              <span className="text-xs text-gray-500">Portfolio</span>
            </div>
            <div className="space-y-3">
              {properties.slice(0, 4).map((item) => (
                <div
                  key={item.ID}
                  className="flex items-center justify-between rounded-2xl border border-[#1f1f1f] bg-[#0f0f0f] px-3 py-3"
                >
                  <div>
                    <p className="text-sm font-semibold truncate max-w-[150px]">Property #{item.ID.substring(0, 8)}</p>
                    <p className="text-xs text-gray-400">{item.Valuation} ETH</p>
                  </div>
                  <p className="text-xs text-gray-300">{item.Status}</p>
                </div>
              ))}
              {properties.length === 0 && !loading && (
                 <p className="text-sm text-gray-500 text-center py-4">No properties found</p>
              )}
            </div>
          </div>

          {/* Investment stats */}
          <div className="rounded-3xl border border-[#1f1f1f] bg-[#111111] p-5 space-y-4 md:col-span-2">
            <div className="flex items-center justify-between">
              <p className="font-semibold">Investment Stats</p>
              <span className="text-xs text-gray-500">Weekly</span>
            </div>
            <div className="grid md:grid-cols-2 gap-4">
              <div className="space-y-3">
                <Card 
                  title="Total Tokens" 
                  value={isConnected && address ? totalBalance : "--"} 
                  sub={isConnected ? "All Properties" : "Connect Wallet"} 
                />
                <Card 
                  title="Properties with Balance" 
                  value={isConnected && address 
                    ? properties.filter((p) => p.tokenBalance && parseFloat(p.tokenBalance) > 0).length.toString()
                    : "--"
                  } 
                  sub={isConnected ? "Owned" : "Connect Wallet"} 
                />
                <Card 
                  title="Wallet Status" 
                  value={isConnected ? "Connected" : "Not Connected"} 
                  sub={isConnected && address ? address.substring(0, 6) + "..." : "Click Connect"} 
                />
              </div>
              <div className="rounded-2xl border border-[#1f1f1f] bg-[#0f0f0f] p-4 flex items-center justify-center">
                <div className="h-32 w-full">
                  <GradientThumb />
                </div>
              </div>
            </div>
          </div>
        </section>

        {/* Portfolio & Distributions */}
        <section className="grid md:grid-cols-3 gap-4">
          <div className="rounded-3xl border border-[#1f1f1f] bg-[#111111] p-5 space-y-3 md:col-span-2">
            <div className="flex items-center justify-between">
              <p className="font-semibold">My Portfolio</p>
              <span className="text-xs text-gray-500">Active</span>
            </div>
            <div className="flex gap-3 overflow-x-auto pb-2">
              {properties.length > 0 ? properties.map((item) => {
                const hasBalance = item.tokenBalance && parseFloat(item.tokenBalance) > 0;
                return (
                  <div
                    key={item.ID}
                    className="min-w-[180px] rounded-2xl border border-[#1f1f1f] bg-[#0f0f0f] p-3 space-y-2 cursor-pointer hover:border-[#6d41ff]/50 transition-colors"
                    onClick={() => navigate(`/properties/${item.ID}`)}
                  >
                    <div className="h-24 rounded-xl border border-[#1f1f1f] overflow-hidden">
                        <GradientThumb />
                    </div>
                    <p className="text-sm font-semibold truncate">Property #{item.ID.substring(0, 8)}</p>
                    <div className="space-y-1">
                      <p className="text-xs text-gray-400">Valuation: {item.Valuation} ETH</p>
                      {isConnected && address && (
                        <p className="text-xs text-gray-300">
                          {item.balanceLoading
                            ? "Loading..."
                            : `Balance: ${item.tokenBalance || "0"} tokens`}
                        </p>
                      )}
                      {isConnected && hasBalance && (
                        <span className="inline-block text-xs px-2 py-0.5 rounded bg-green-500/20 text-green-400 border border-green-500/30">
                          Owned
                        </span>
                      )}
                      {isConnected && item.claimableDistributions && item.claimableDistributions.length > 0 && (
                        <div className="space-y-1 pt-1 border-t border-gray-700">
                          <p className="text-xs text-yellow-400 font-semibold">
                            💰 Claimable: {formatTokenBalance(item.claimableDistributions.reduce((sum, dist) => sum + dist.amount, 0n))}
                          </p>
                          <Button
                            variant="outline"
                            size="sm"
                            onClick={() => {
                              // For simplicity, claim from the first distribution
                              // In production, you might want to show all distributions
                              if (item.claimableDistributions && item.claimableDistributions.length > 0) {
                                handleClaimRevenue(item.ID, item.claimableDistributions[0].distributionId);
                              }
                            }}
                            disabled={item.claiming}
                            className="w-full text-xs py-1 h-6 border-yellow-500/50 bg-yellow-500/10 text-yellow-400 hover:bg-yellow-500/20"
                          >
                            {item.claiming ? "Claiming..." : "Claim Revenue"}
                          </Button>
                        </div>
                      )}
                    </div>
                  </div>
                );
              }) : (
                 <div className="w-full text-center py-8 text-gray-500">
                    {loading ? "Loading properties..." : "No properties available"}
                 </div>
              )}
            </div>
          </div>
          <div className="rounded-3xl border border-[#1f1f1f] bg-[#111111] p-5 space-y-3">
            <div className="flex items-center justify-between">
              <p className="font-semibold">Top Picks</p>
              <span className="text-xs text-gray-500">Today</span>
            </div>
            <div className="space-y-3">
               {/* Placeholder for top picks until we have a logic for it */}
               <p className="text-sm text-gray-500 text-center py-4">Coming soon</p>
            </div>
          </div>
        </section>
      </main>
    </div>
  );
}

export default DashboardPage;


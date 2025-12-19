import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { Button } from "../Components/atoms/Button";
import { api } from "../lib/api";
import { useWallet } from "../hooks/useWallet";
import type { Property, User } from "../types";
import { ethers } from "ethers";
import { getTokenBalance, transferTokens } from "../lib/contracts";

const BALANCE_MULTIPLIER = 1000000; // 1 SepoliaETH = 1,000,000 ETH in system

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

function DashboardHome() {
  const [properties, setProperties] = useState<Property[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");
  const [walletBalance, setWalletBalance] = useState<string>("0");
  const [user, setUser] = useState<User | null>(null);
  const navigate = useNavigate();
  const { provider, isManuallyConnected, signer } = useWallet();
  const [allPendingTransfers, setAllPendingTransfers] = useState<Array<{
    id: string;
    property_id: string;
    buyer_wallet: string;
    amount: string;
    payment_tx_hash: string;
    token_tx_hash: string;
    purchase_price: string;
    created_at: string;
  }>>([]);
  const [approvingTransfer, setApprovingTransfer] = useState<string | null>(null);

  // Load user data
  useEffect(() => {
    const userStr = localStorage.getItem("user");
    if (userStr) {
      setUser(JSON.parse(userStr));
    }
  }, []);

  // Fetch pending transfers for owner
  useEffect(() => {
    const fetchPendingTransfers = async () => {
      if (!user) return;
      try {
        const transfers = await api.getMyPendingTransfers();
        setAllPendingTransfers(transfers);
        console.log(`📋 Found ${transfers.length} pending transfer(s) across all properties`);
      } catch (err) {
        console.error("Failed to fetch pending transfers:", err);
      }
    };

    fetchPendingTransfers();
    // Poll every 5 seconds
    const interval = setInterval(fetchPendingTransfers, 5000);
    return () => clearInterval(interval);
  }, [user]);

  // Handle approving transfer from dashboard
  const handleApproveTransfer = async (transfer: any) => {
    if (!signer || !provider) {
      alert("Please connect your wallet to approve transfers");
      return;
    }

    setApprovingTransfer(transfer.id);

    try {
      // Get property to get token address
      const property = await api.getProperty(transfer.property_id);
      
      console.log(`✅ Owner approving transfer: ${transfer.amount} tokens to ${transfer.buyer_wallet}`);
      
      // Convert amount to wei
      const tokenAmountWei = ethers.parseUnits(transfer.amount, 18);
      
      // Transfer tokens from owner to buyer
      const transferTx = await transferTokens(
        property.OnchainTokenAddress,
        transfer.buyer_wallet,
        tokenAmountWei,
        signer
      );

      console.log(`📤 Token transfer transaction sent: ${transferTx.hash}`);
      const receipt = await transferTx.wait();
      console.log(`✅ Token transfer confirmed in block: ${receipt?.blockNumber}`);

      // Update database with token transfer hash
      try {
        await api.updateTokenPurchaseTxHash(
          transfer.property_id,
          transfer.id,
          transferTx.hash
        );
        console.log(`✅ Purchase record updated with token transfer hash`);
      } catch (updateErr: any) {
        console.error("❌ Failed to update purchase record:", updateErr);
      }

      // Refresh pending transfers
      const updatedTransfers = await api.getMyPendingTransfers();
      setAllPendingTransfers(updatedTransfers);

      alert(`✅ Transfer Approved!\n\n${transfer.amount} tokens transferred to ${transfer.buyer_wallet}\n\nToken TX: ${transferTx.hash}\nPayment TX: ${transfer.payment_tx_hash}`);
    } catch (err: any) {
      console.error("Failed to approve transfer:", err);
      alert(`Failed to approve transfer: ${err.message || "Unknown error"}`);
    } finally {
      setApprovingTransfer(null);
    }
  };

  useEffect(() => {
    const fetchProperties = async () => {
      if (!user || !user.WalletAddress) {
        setLoading(false);
        return;
      }

      try {
        // Get all properties
        const allProperties = await api.getProperties();
        
        // Use user's registered wallet address from database
        const userWalletAddress = user.WalletAddress.toLowerCase();
        
        console.log("🔍 DashboardHome: Filtering properties for wallet:", userWalletAddress);
        console.log("📋 Total properties found:", allProperties.length);

        // Filter properties where user is owner OR has token balance > 0
        const filteredProperties: Property[] = [];

        for (const property of allProperties) {
          const propertyOwnerWallet = property.OwnerWallet?.toLowerCase() || "";
          const isOwner = propertyOwnerWallet === userWalletAddress;
          
          // If user is owner, always include by default (they created it)
          // But check token balance if wallet is manually connected
          if (isOwner) {
            // Only check token balance if wallet is manually connected
            if (isManuallyConnected && provider && property.OnchainTokenAddress) {
              try {
                const balance = await getTokenBalance(
                  property.OnchainTokenAddress,
                  userWalletAddress,
                  provider
                );
                // Only exclude if user has explicitly sold ALL tokens (balance is exactly 0)
                if (balance > 0n) {
                  filteredProperties.push(property);
                }
              } catch (err) {
                // If balance check fails, include it anyway since user is owner
                filteredProperties.push(property);
              }
            } else {
              // If wallet not manually connected, still show owner properties
              filteredProperties.push(property);
            }
          } else {
            // User is not owner, check if they have tokens (investor)
            // Only check if wallet is manually connected
            if (isManuallyConnected && provider && property.OnchainTokenAddress) {
              try {
                const balance = await getTokenBalance(
                  property.OnchainTokenAddress,
                  userWalletAddress,
                  provider
                );
                // Include if user has tokens
                if (balance > 0n) {
                  filteredProperties.push(property);
                }
              } catch (err) {
                // If balance check fails, don't include it (not owner, so need proof of investment)
              }
            }
            // If wallet not manually connected, don't show investor properties
          }
        }
        
        console.log(`✅ DashboardHome: Filtered properties: ${filteredProperties.length} out of ${allProperties.length}`);

        setProperties(filteredProperties);
      } catch (err: any) {
        console.error("Failed to fetch properties:", err);
        setError("Failed to load properties");
      } finally {
        setLoading(false);
      }
    };

    fetchProperties();
  }, [user, isManuallyConnected, provider]);

  // Fetch wallet balance - ALWAYS use user's registered wallet address
  useEffect(() => {
    const fetchWalletBalance = async () => {
      const userAddress = user?.WalletAddress;
      
      if (!userAddress) {
        console.log("ℹ️ No user wallet address available for balance check");
        setWalletBalance("0");
        return;
      }
      
      // Only fetch balance if wallet is manually connected
      if (isManuallyConnected && provider) {
        try {
          console.log(`💰 Fetching ETH balance for user's registered wallet: ${userAddress}`);
          const balance = await provider.getBalance(userAddress);
          const onChainBalanceEth = Number(ethers.formatEther(balance));
          const scaledBalance = (onChainBalanceEth * BALANCE_MULTIPLIER).toFixed(2);
          console.log(
            `✅ Balance fetched: ${onChainBalanceEth} SepoliaETH, scaled: ${scaledBalance} ETH`
          );
          setWalletBalance(scaledBalance);
        } catch (err) {
          console.error("❌ Failed to fetch wallet balance:", err);
          setWalletBalance("0");
        }
      } else {
        // Don't fetch balance until wallet is manually connected
        setWalletBalance("0");
      }
    };

    fetchWalletBalance();
  }, [user, isManuallyConnected, provider]);

  // Compute dynamic values from existing data
  const avgPrice = properties.length > 0
    ? (properties.reduce((sum, p) => sum + p.Valuation, 0) / properties.length).toFixed(2)
    : "0.00";

  const totalInvestment = properties.reduce((sum, p) => sum + p.Valuation, 0).toFixed(2);

  return (
    <div className="space-y-8">
      <header className="flex items-center justify-between">
        <div>
          <p className="text-sm text-gray-400">Welcome back</p>
          <h1 className="text-3xl font-semibold">Dashboard</h1>
        </div>
        <Button
          variant="primary"
          size="md"
          className="bg-[#6d41ff] hover:bg-[#5b2fff] text-white"
          onClick={() => navigate("/dashboard/upload")}
        >
          New Investment
        </Button>
      </header>

      {error && (
        <div className="text-red-500 bg-red-500/10 p-4 rounded-xl border border-red-500/20">
          {error}
        </div>
      )}

      {/* Pending Transfers Notification (Owner Only) */}
      {allPendingTransfers.length > 0 && (
        <div className="bg-yellow-900/20 border border-yellow-600/50 rounded-xl p-6 space-y-4 mb-6">
          <div className="flex items-center gap-2">
            <span className="text-yellow-400 font-semibold text-lg">⚠️ Pending Token Transfers</span>
            <span className="bg-yellow-600 text-black px-2 py-1 rounded text-sm font-bold">
              {allPendingTransfers.length}
            </span>
          </div>
          <p className="text-gray-300 text-sm">
            Buyers have sent payments and are waiting for you to approve token transfers. Click on a property to view details or approve below.
          </p>
          <div className="space-y-3 mt-4 max-h-96 overflow-y-auto">
            {allPendingTransfers.map((transfer: any) => {
              // Find property name
              const property = properties.find(p => p.ID === transfer.property_id);
              const propertyName = property?.Name || `Property #${transfer.property_id.substring(0, 8)}`;
              
              return (
                <div
                  key={transfer.id}
                  className="bg-[#1a1a1a] border border-gray-700 rounded-lg p-4 space-y-3"
                >
                  <div className="flex justify-between items-start">
                    <div className="space-y-1">
                      <p className="text-sm text-gray-400">Property</p>
                      <p 
                        className="font-semibold cursor-pointer hover:text-[#6d41ff]"
                        onClick={() => navigate(`/properties/${transfer.property_id}`)}
                      >
                        {propertyName}
                      </p>
                    </div>
                    <div className="text-right space-y-1">
                      <p className="text-sm text-gray-400">Amount</p>
                      <p className="font-semibold">{transfer.amount} tokens</p>
                    </div>
                  </div>
                  <div className="grid grid-cols-2 gap-4 text-sm">
                    <div>
                      <p className="text-gray-400">Buyer</p>
                      <p className="font-mono text-xs">{transfer.buyer_wallet}</p>
                    </div>
                    <div>
                      <p className="text-gray-400">Payment Received</p>
                      <p className="font-mono text-xs">{transfer.purchase_price} ETH</p>
                    </div>
                    <div>
                      <p className="text-gray-400">Payment TX</p>
                      <a
                        href={`https://sepolia.etherscan.io/tx/${transfer.payment_tx_hash}`}
                        target="_blank"
                        rel="noopener noreferrer"
                        className="text-[#6d41ff] hover:underline font-mono text-xs"
                      >
                        {transfer.payment_tx_hash.substring(0, 10)}...
                      </a>
                    </div>
                  </div>
                  <div className="flex gap-2 pt-2">
                    <Button
                      onClick={() => navigate(`/properties/${transfer.property_id}`)}
                      variant="outline"
                      size="sm"
                      className="flex-1 border-[#262626] text-gray-400 hover:bg-[#1a1a1a]"
                    >
                      View Property
                    </Button>
                    <Button
                      onClick={() => handleApproveTransfer(transfer)}
                      disabled={approvingTransfer === transfer.id}
                      variant="primary"
                      size="sm"
                      className="flex-1 bg-green-600 hover:bg-green-700"
                    >
                      {approvingTransfer === transfer.id ? "Approving..." : "✅ Approve"}
                    </Button>
                  </div>
                </div>
              );
            })}
          </div>
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
                  <h3 className="text-xl font-semibold">
                    Property #{properties[0].ID.substring(0, 8)}
                  </h3>
                  <p className="text-sm text-gray-400">
                    {properties[0].Valuation} ETH · {properties[0].Status}
                  </p>
                </div>
              </div>
            ) : (
              <div className="space-y-3 flex items-center justify-center border border-[#1f1f1f] rounded-2xl h-full min-h-[200px]">
                <p className="text-gray-500">
                  {loading ? "Loading properties..." : "No properties available"}
                </p>
              </div>
            )}

            <div className="space-y-3">
              <div className="grid grid-cols-2 gap-3">
                <Card
                  title="Total Properties"
                  value={properties.length.toString()}
                  sub="My Properties"
                />
                <Card
                  title="Active"
                  value={properties
                    .filter((p) => p.Status === "Active")
                    .length.toString()}
                  sub="Status"
                  accent="bg-[#2e2e2e]"
                />
                <Card
                  title="Avg Price"
                  value={`${avgPrice} ETH`}
                  sub="Est."
                  accent="bg-[#2c1f5e]"
                />
                <Card
                  title="Watchlist"
                  value="0"
                  sub="Saved"
                  accent="bg-[#1f3a5e]"
                />
              </div>
              <Button
                variant="outline"
                size="md"
                className="border-[#262626] bg-[#111111] text-white hover:bg-[#1a1a1a] w-full"
                onClick={() => navigate("/dashboard/properties")}
              >
                View All Properties
              </Button>
            </div>
          </div>
        </div>

        {/* Balance card (UI only, no on-chain logic here) */}
        <div className="rounded-3xl border border-[#1f1f1f] bg-linear-to-br from-[#6d41ff] via-[#5a3ccc] to-[#1b132f] p-5 space-y-3">
          <p className="text-sm text-white/80">My Wallet</p>
          <p className="text-3xl font-semibold">{walletBalance} ETH</p>
          <p className="text-sm text-white/80">Monthly Profit</p>
          <p className="text-2xl font-semibold">0.00 ETH</p>
          <div className="flex gap-2">
            <span className="rounded-full bg-white/10 px-3 py-1 text-xs">
              +0%
            </span>
          </div>
          <Button
            variant="outline"
            size="md"
            className="border-white/30 text-white hover:bg-white/10"
            onClick={() => navigate("/dashboard/wallet")}
          >
            Manage Wallet
          </Button>
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
                  <p className="text-sm font-semibold truncate max-w-[150px]">
                    Property #{item.ID.substring(0, 8)}
                  </p>
                  <p className="text-xs text-gray-400">{item.Valuation} ETH</p>
                </div>
                <p className="text-xs text-gray-300">{item.Status}</p>
              </div>
            ))}
            {properties.length === 0 && !loading && (
              <p className="text-sm text-gray-500 text-center py-4">
                No properties found
              </p>
            )}
          </div>
        </div>

        {/* Investment stats (UI only) */}
        <div className="rounded-3xl border border-[#1f1f1f] bg-[#111111] p-5 space-y-4 md:col-span-2">
          <div className="flex items-center justify-between">
            <p className="font-semibold">Investment Stats</p>
            <span className="text-xs text-gray-500">Weekly</span>
          </div>
          <div className="grid md:grid-cols-2 gap-4">
            <div className="space-y-3">
              <Card title="Total Investment" value={`${totalInvestment} ETH`} sub="Active" />
              <Card title="Weekly Returns" value="0.00 ETH" sub="Last 7 days" />
              <Card title="Expenses" value="0.00 ETH" sub="Gas / Fees" />
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
            {properties.length > 0 ? (
              properties.map((item) => (
                <div
                  key={item.ID}
                  className="min-w-[180px] rounded-2xl border border-[#1f1f1f] bg-[#0f0f0f] p-3 space-y-2"
                >
                  <div className="h-24 rounded-xl border border-[#1f1f1f] overflow-hidden">
                    <GradientThumb />
                  </div>
                  <p className="text-sm font-semibold truncate">
                    Property #{item.ID.substring(0, 8)}
                  </p>
                  <p className="text-xs text-gray-400">
                    {item.Valuation} ETH
                  </p>
                </div>
              ))
            ) : (
              <div className="w-full text-center py-8 text-gray-500">
                You don't own any properties yet.
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
            <p className="text-sm text-gray-500 text-center py-4">
              Coming soon
            </p>
          </div>
        </div>
      </section>
    </div>
  );
}

export default DashboardHome;



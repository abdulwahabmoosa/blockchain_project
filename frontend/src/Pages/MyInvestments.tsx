import { useEffect, useState, useCallback } from "react";
import { useNavigate } from "react-router-dom";
import { TrendingUp, TrendingDown, ArrowUpRight } from "lucide-react";
import { Button } from "../Components/atoms/Button";
import { api } from "../lib/api";
import { getTokenBalance, getClaimableDistributions, claimRevenue } from "../lib/contracts";
import { useWallet } from "../hooks/useWallet";
import { formatTokenBalance } from "../utils/format";
import { ethers } from "ethers";
import type { Property, User } from "../types";

const GradientThumb = () => (
  <div className="rounded-xl border border-[#1f1f1f] bg-linear-to-br from-[#191919] via-[#1f1f1f] to-[#121212] w-full h-full" />
);

const StatCard = ({
  title,
  value,
  change,
  positive = true,
}: {
  title: string;
  value: string;
  change?: string;
  positive?: boolean;
}) => (
  <div className="rounded-2xl border border-[#1f1f1f] bg-[#111111] p-4 space-y-2">
    <p className="text-sm text-gray-400">{title}</p>
    <p className="text-2xl font-semibold">{value}</p>
    {change && (
      <div
        className={`flex items-center gap-1 text-sm ${
          positive ? "text-green-400" : "text-red-400"
        }`}
      >
        {positive ? <TrendingUp size={16} /> : <TrendingDown size={16} />}
        <span>{change}</span>
      </div>
    )}
  </div>
);

interface PropertyWithBalance extends Property {
  tokenBalance?: string;
  tokenValueETH?: number;
  claimableDistributions?: Array<{
    distributionId: number;
    amount: bigint;
    stablecoin: string;
  }>;
  claimingRent?: boolean;
}

function MyInvestments() {
  const [properties, setProperties] = useState<PropertyWithBalance[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");
  const navigate = useNavigate();
  const { address, provider, isManuallyConnected, signer, connect } = useWallet();
  const [user, setUser] = useState<User | null>(null);
  const [refreshing, setRefreshing] = useState(false);

  // Load user data
  useEffect(() => {
    const userStr = localStorage.getItem("user");
    if (userStr) {
      setUser(JSON.parse(userStr));
    }
  }, []);

  const fetchProperties = useCallback(async () => {
      if (!user || !user.WalletAddress) {
        setLoading(false);
        return;
      }

      setLoading(true);
      try {
        // Get all properties
        const allProperties = await api.getProperties();
        
        // Get user's token purchases from database (handle failure gracefully)
        let myPurchases: any[] = [];
        let purchasedPropertyIds = new Set<string>();
        try {
          myPurchases = await api.getMyTokenPurchases();
          console.log("📦 My token purchases from database:", myPurchases.length);
          
          // Create a map of property IDs that user has purchased tokens for
          myPurchases.forEach((purchase: any) => {
            purchasedPropertyIds.add(purchase.property_id);
          });
          console.log("📋 Properties with purchases:", Array.from(purchasedPropertyIds));
        } catch (purchaseErr) {
          // If purchases API fails, continue without purchase data (still show owner properties)
          console.warn("Warning: Failed to fetch token purchases (non-critical):", purchaseErr);
        }
        
        // Use user's registered wallet address from database (more reliable than connected wallet)
        const userWalletAddress = user.WalletAddress.toLowerCase();
        
        console.log("MyInvestments: Filtering properties for wallet:", userWalletAddress);
        console.log("📋 Total properties found:", allProperties.length);

        // Filter properties where user is owner OR has token balance > 0 OR has purchases recorded
        const filteredProperties: PropertyWithBalance[] = [];

        for (const property of allProperties) {
          let propertyBalance: bigint = 0n;
          const propertyOwnerWallet = property.OwnerWallet?.toLowerCase() || "";
          const isOwner = propertyOwnerWallet === userWalletAddress;
          
          console.log(`Checking property ${property.ID}:`, {
            propertyOwner: propertyOwnerWallet,
            userWallet: userWalletAddress,
            isOwner,
            hasTokenAddress: !!property.OnchainTokenAddress,
            hasProvider: !!provider,
            isManuallyConnected
          });
          
          // If user is owner, ALWAYS include (they created it, regardless of token balance)
          if (isOwner) {
            console.log(`✅ Including property ${property.ID} (user is owner)`);
            // For owners, still check balance for display purposes
            let tokenBalanceStr = "0";
            let tokenValueETH = 0;
            if (provider && property.OnchainTokenAddress) {
              try {
                propertyBalance = await getTokenBalance(
                  property.OnchainTokenAddress,
                  userWalletAddress,
                  provider
                );
                tokenBalanceStr = formatTokenBalance(propertyBalance);
                console.log(`💰 Owner token balance for property ${property.ID}:`, tokenBalanceStr);
                
                // Calculate ETH value: tokenSupply = valuation/10, so each token = 10 ETH worth
                const balanceFloat = Number(tokenBalanceStr);
                tokenValueETH = balanceFloat * 10;

                // Fetch claimable distributions if user has tokens
                if (provider && propertyBalance > 0n) {
                  try {
                    const claimable = await getClaimableDistributions(
                      property.OnchainTokenAddress,
                      userWalletAddress,
                      provider
                    );
                    filteredProperties.push({ ...property, tokenBalance: tokenBalanceStr, tokenValueETH, claimableDistributions: claimable });
                  } catch (err) {
                    console.warn(`Warning: Could not fetch claimable distributions for property ${property.ID}`);
                    filteredProperties.push({ ...property, tokenBalance: tokenBalanceStr, tokenValueETH });
                  }
                } else {
                  filteredProperties.push({ ...property, tokenBalance: tokenBalanceStr, tokenValueETH });
                }
              } catch (err) {
                console.warn(`Warning: Could not fetch owner balance for property ${property.ID}`);
                filteredProperties.push(property);
              }
            } else {
              filteredProperties.push(property);
            }
          } else {
            // User is not owner, check if they have tokens (investor)
            // First check if user has a purchase record for this property
            const hasPurchaseRecord = purchasedPropertyIds.has(property.ID);
            
            if (hasPurchaseRecord) {
              // User has purchased tokens for this property (even if not yet transferred)
              console.log(`✅ Including property ${property.ID} (investor with purchase record in database)`);
              // Get token balance to calculate value
              if (provider && property.OnchainTokenAddress) {
                try {
                  const balance = await getTokenBalance(
                    property.OnchainTokenAddress,
                    userWalletAddress,
                    provider
                  );
                  const tokenBalanceStr = formatTokenBalance(balance);
                  const tokenBalanceNum = parseFloat(tokenBalanceStr);
                  const tokenValueETH = tokenBalanceNum * 10; // Each token = 10 ETH worth (valuation/10 tokens)
                  
                  // Fetch claimable distributions
                  if (provider && balance > 0n) {
                    try {
                      const claimable = await getClaimableDistributions(
                        property.OnchainTokenAddress,
                        userWalletAddress,
                        provider
                      );
                      filteredProperties.push({ ...property, tokenBalance: tokenBalanceStr, tokenValueETH, claimableDistributions: claimable });
                    } catch (err) {
                      console.warn(`Warning: Could not fetch claimable distributions for property ${property.ID}`);
                      filteredProperties.push({ ...property, tokenBalance: tokenBalanceStr, tokenValueETH });
                    }
                  } else {
                    filteredProperties.push({ ...property, tokenBalance: tokenBalanceStr, tokenValueETH });
                  }
                } catch (err) {
                  console.error(`Failed to get balance for property ${property.ID}:`, err);
                  filteredProperties.push(property);
                }
              } else {
                filteredProperties.push(property);
              }
              continue; // Skip blockchain balance check since we have purchase record
            } else if (provider && property.OnchainTokenAddress) {
              // No purchase record, but check blockchain balance as fallback
              try {
                // Check balance for registered wallet address
                console.log(`Checking token balance for property ${property.ID}:`, {
                  tokenAddress: property.OnchainTokenAddress,
                  registeredWallet: userWalletAddress,
                  connectedWallet: address?.toLowerCase(),
                });
                
                let balance = await getTokenBalance(
                  property.OnchainTokenAddress,
                  userWalletAddress,
                  provider
                );
                console.log(`💰 Balance in registered wallet (${userWalletAddress}):`, formatTokenBalance(balance));
                
                // Also check currently connected address if different
                if (address && address.toLowerCase() !== userWalletAddress) {
                  const connectedBalance = await getTokenBalance(
                    property.OnchainTokenAddress,
                    address.toLowerCase(),
                    provider
                  );
                  console.log(`💰 Balance in connected wallet (${address.toLowerCase()}):`, formatTokenBalance(connectedBalance));
                  // Use the higher balance (tokens might be in either address)
                  if (connectedBalance > balance) {
                    balance = connectedBalance;
                    console.log(`✅ Found tokens in connected address (${address}) instead of registered address`);
                  }
                }
                
                console.log(`💰 Final token balance for property ${property.ID}:`, balance.toString(), `(${formatTokenBalance(balance)} tokens)`);
                // Include if user has tokens
                if (balance > 0n) {
                  const tokenBalanceStr = formatTokenBalance(balance);
                  const tokenBalanceNum = parseFloat(tokenBalanceStr);
                  const tokenValueETH = tokenBalanceNum * 10; // Each token = 10 ETH worth (valuation/10 tokens)
                  console.log(`✅ Including property ${property.ID} (investor with ${tokenBalanceStr} tokens, worth ${tokenValueETH.toFixed(2)} ETH)`);
                  
                  // Fetch claimable distributions
                  if (provider && balance > 0n) {
                    try {
                      const claimable = await getClaimableDistributions(
                        property.OnchainTokenAddress,
                        userWalletAddress,
                        provider
                      );
                      filteredProperties.push({ ...property, tokenBalance: tokenBalanceStr, tokenValueETH, claimableDistributions: claimable });
                    } catch (err) {
                      console.warn(`Warning: Could not fetch claimable distributions for property ${property.ID}`);
                      filteredProperties.push({ ...property, tokenBalance: tokenBalanceStr, tokenValueETH });
                    }
                  } else {
                    filteredProperties.push({ ...property, tokenBalance: tokenBalanceStr, tokenValueETH });
                  }
                } else {
                  console.log(`Warning: Excluding property ${property.ID} (investor but no tokens found)`);
                }
              } catch (err) {
                console.error(`❌ Failed to check balance for property ${property.ID}:`, err);
                // If balance check fails, don't include it (not owner, so need proof of investment)
              }
            } else {
              // If wallet not connected and no purchase record, don't show investor properties
              console.log(`Warning: Excluding property ${property.ID} (investor, no purchase record and wallet not connected)`);
            }
          }
        }
        
        console.log(`✅ Filtered properties: ${filteredProperties.length} out of ${allProperties.length}`);

        setProperties(filteredProperties);
      } catch (err: any) {
        console.error("Failed to fetch properties:", err);
        setError("Failed to load investments");
      } finally {
        setLoading(false);
      }
  }, [user, isManuallyConnected, provider, address]);

  useEffect(() => {
    fetchProperties();
  }, [fetchProperties]);

  // Manual refresh function
  const handleRefresh = async () => {
    setRefreshing(true);
    setError("");
    try {
      await fetchProperties();
    } catch (err) {
      console.error("Failed to refresh investments:", err);
    } finally {
      setRefreshing(false);
    }
  };

  // Claim all rent for a property
  const handleClaimAllRent = async (property: PropertyWithBalance) => {
    if (!signer) {
      alert("Please connect your wallet first");
      try {
        await connect();
      } catch (err) {
        setError("Failed to connect wallet");
      }
      return;
    }

    if (!property.claimableDistributions || property.claimableDistributions.length === 0) {
      alert("No claimable rent available");
      return;
    }

    // Set claiming state
    setProperties(prev => prev.map(p =>
      p.ID === property.ID ? { ...p, claimingRent: true } : p
    ));
    setError("");

    try {
      // Claim all distributions sequentially
      let claimedCount = 0;
      for (const dist of property.claimableDistributions) {
        try {
          console.log(`💰 Claiming distribution ${dist.distributionId} for property ${property.ID}...`);
          const tx = await claimRevenue(dist.distributionId, signer);
          await tx.wait();
          console.log(`✅ Claimed distribution ${dist.distributionId}. TX: ${tx.hash}`);
          claimedCount++;
        } catch (err: any) {
          console.error(`❌ Failed to claim distribution ${dist.distributionId}:`, err);
          // Continue with other distributions even if one fails
        }
      }

      if (claimedCount > 0) {
        alert(`✅ Successfully claimed ${claimedCount} out of ${property.claimableDistributions.length} distribution(s)!`);
        // Refresh properties to update claimable distributions
        await fetchProperties();
      } else {
        throw new Error("Failed to claim any distributions");
      }

    } catch (err: any) {
      console.error("❌ Failed to claim rent:", err);
      setError(err.message || "Failed to claim rent");
      alert(`❌ Failed to claim rent: ${err.message || "Unknown error"}`);
    } finally {
      setProperties(prev => prev.map(p =>
        p.ID === property.ID ? { ...p, claimingRent: false } : p
      ));
    }
  };

  const totalValue = properties.reduce((sum, p) => sum + (p.tokenValueETH || 0), 0);
  const activeCount = properties.filter((p) => p.Status === "Active").length;

  return (
    <div className="space-y-6">
      <header className="flex flex-col md:flex-row md:items-center justify-between gap-4">
        <div>
          <h1 className="text-3xl font-semibold">My Investments</h1>
          <p className="text-sm text-gray-400">
            Track and manage your property investments
          </p>
        </div>
        <div className="flex gap-2">
          <Button
            variant="outline"
            size="md"
            className="border-[#262626] bg-[#1A1A1A] text-white hover:bg-[#262626]"
            onClick={handleRefresh}
            disabled={refreshing || loading}
          >
            {refreshing ? "Refreshing..." : "Refresh"}
          </Button>
          <Button
            variant="primary"
            size="md"
            className="bg-[#6d41ff] hover:bg-[#5b2fff] text-white"
            onClick={() => navigate("/dashboard/properties")}
          >
            Browse Properties
          </Button>
        </div>
      </header>

      {error && (
        <div className="text-red-500 bg-red-500/10 p-4 rounded-xl border border-red-500/20">
          {error}
        </div>
      )}

      {/* Stats Overview */}
      <div className="grid grid-cols-2 md:grid-cols-4 gap-4">
        <StatCard
          title="Total Value"
          value={`${totalValue.toFixed(2)} ETH`}
          change="+0%"
          positive={true}
        />
        <StatCard
          title="Total Investments"
          value={properties.length.toString()}
        />
        <StatCard title="Active" value={activeCount.toString()} />
        <StatCard
          title="Returns"
          value="0.00 ETH"
          change="+0%"
          positive={true}
        />
      </div>

      {/* Investment List */}
      <div className="rounded-3xl border border-[#1f1f1f] bg-[#111111] p-6 space-y-4">
        <div className="flex items-center justify-between">
          <h2 className="font-semibold text-lg">Your Investments</h2>
          <span className="text-xs text-gray-500">
            {properties.length} properties
          </span>
        </div>

        {loading ? (
          <div className="space-y-3">
            {[...Array(3)].map((_, i) => (
              <div
                key={i}
                className="flex items-center gap-4 p-4 rounded-xl border border-[#1f1f1f] bg-[#0f0f0f] animate-pulse"
              >
                <div className="h-16 w-16 rounded-xl bg-[#1a1a1a]" />
                <div className="flex-1 space-y-2">
                  <div className="h-4 rounded bg-[#1a1a1a] w-1/3" />
                  <div className="h-3 rounded bg-[#1a1a1a] w-1/4" />
                </div>
              </div>
            ))}
          </div>
        ) : properties.length > 0 ? (
          <div className="space-y-3">
            {properties.map((property) => (
              <div
                key={property.ID}
                onClick={() => navigate(`/properties/${property.ID}`)}
                className="flex items-center gap-4 p-4 rounded-xl border border-[#1f1f1f] bg-[#0f0f0f] cursor-pointer hover:border-[#703BF7] transition-colors"
              >
                <div className="h-16 w-16 rounded-xl overflow-hidden flex-shrink-0 bg-gradient-to-br from-[#1f1f1f] via-[#2a2a2a] to-[#111111]">
                  {property.MetadataHash ? (
                    <img
                      src={`https://gateway.pinata.cloud/ipfs/${property.MetadataHash}`}
                      alt="Property"
                      className="w-full h-full object-cover"
                      onError={(e) => {
                        e.currentTarget.style.display = 'none';
                        e.currentTarget.parentElement!.innerHTML = '<div class="w-full h-full flex items-center justify-center"><span class="text-gray-500 text-xs">No img</span></div>';
                      }}
                    />
                  ) : (
                    <div className="w-full h-full flex items-center justify-center">
                      <span className="text-gray-500 text-xs">No img</span>
                    </div>
                  )}
                </div>
                <div className="flex-1 min-w-0">
                  <h3 className="font-semibold truncate">
                    {property.Name || `Property #${property.ID.substring(0, 8)}`}
                  </h3>
                  <p className="text-sm text-gray-400">
                    {property.tokenBalance ? `${property.tokenBalance} tokens` : "0 tokens"}
                  </p>
                  <p className="text-xs text-gray-500">
                    {property.tokenValueETH ? `≈ ${property.tokenValueETH.toFixed(2)} ETH value` : "—"}
                  </p>
                  {property.claimableDistributions && property.claimableDistributions.length > 0 && (
                    <div className="mt-2 pt-2 border-t border-gray-700" style={{ display: 'none' }}>
                      <p className="text-xs text-yellow-400 font-semibold">
                        💰 Claimable Rent: {formatTokenBalance(property.claimableDistributions.reduce((sum, dist) => sum + dist.amount, 0n))} WETH
                      </p>
                      <p className="text-xs text-gray-500">
                        {property.claimableDistributions.length} distribution(s) available
                      </p>
                    </div>
                  )}
                </div>
                <div className="text-right flex-shrink-0 flex flex-col items-end gap-2">
                  <span
                    className={`text-xs px-2 py-1 rounded-full ${
                      property.Status === "Active"
                        ? "bg-green-500/20 text-green-400"
                        : property.Status === "Paused"
                        ? "bg-yellow-500/20 text-yellow-400"
                        : "bg-gray-500/20 text-gray-400"
                    }`}
                  >
                    {property.Status}
                  </span>
                  {property.claimableDistributions && property.claimableDistributions.length > 0 && (
                    <Button
                      variant="primary"
                      size="sm"
                      onClick={(e) => {
                        e.stopPropagation();
                        handleClaimAllRent(property);
                      }}
                      disabled={property.claimingRent || !signer}
                      className="bg-green-600 hover:bg-green-700 text-white text-xs disabled:opacity-50"
                      style={{ display: 'none' }}
                    >
                      {property.claimingRent ? "Claiming..." : "Claim Rent"}
                    </Button>
                  )}
                  <p className="text-xs text-gray-500">
                    {new Date(property.CreatedAt).toLocaleDateString()}
                  </p>
                </div>
                <ArrowUpRight className="text-gray-500" size={20} />
              </div>
            ))}
          </div>
        ) : (
          <div className="text-center py-12">
            <p className="text-gray-400 text-lg">No investments yet</p>
            <p className="text-gray-500 text-sm mt-2">
              Start investing in properties to build your portfolio
            </p>
            <Button
              variant="primary"
              size="md"
              className="bg-[#6d41ff] hover:bg-[#5b2fff] text-white mt-4"
              onClick={() => navigate("/dashboard/properties")}
            >
              Browse Properties
            </Button>
          </div>
        )}
      </div>

    </div>
  );
}

export default MyInvestments;



import "../index.css";
import { useEffect, useState, useCallback } from "react";
import { useParams, Link as RouterLink, useNavigate } from "react-router-dom";
import { Navbar } from "../Components/organisms/Navbar";
import { Button } from "../Components/atoms/Button";
import { api } from "../lib/api";
import type { Property, PropertyMetadata } from "../types";
import { useWallet } from "../hooks/useWallet";
import {
  getTokenBalance,
  checkApprovalStatus,
  transferTokens,
  getTotalSupply,
} from "../lib/contracts";
import { formatTokenBalance } from "../utils/format";
import { ethers } from "ethers";
import { Home, MapPin, Building2, Calendar, TrendingUp } from "lucide-react";

// Conversion factor: 1,000,000 system ETH = 1 SepoliaETH
const SEPOLIA_CONVERSION_FACTOR = 1000000;

const GradientThumb = () => (
  <div className="w-full h-full bg-linear-to-br from-[#1f1f1f] via-[#2a2a2a] to-[#111111]" />
);

export default function PropertyDetailsPage() {
  const { id } = useParams<{ id: string }>();
  const [property, setProperty] = useState<Property | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");
  const [tokenBalance, setTokenBalance] = useState<string>("");
  const [balanceLoading, setBalanceLoading] = useState(false);
  const [isApproved, setIsApproved] = useState<boolean | null>(null);
  const [transferTo, setTransferTo] = useState(""); // Transfer recipient
  const [transferAmount, setTransferAmount] = useState(""); // Transfer amount
  const [transferring, setTransferring] = useState(false); // Transfer loading
  const [autoConnectAttempted, setAutoConnectAttempted] = useState(false);
  const [metadata, setMetadata] = useState<PropertyMetadata | null>(null);
  const [metadataLoading, setMetadataLoading] = useState(false);
  const { isConnected, address, provider, signer, connectRegisteredWallet, isManuallyConnected } = useWallet();
  const navigate = useNavigate();
  const [user, setUser] = useState<any>(null);
  
  // Buy tokens state
  const [buyTokenAmount, setBuyTokenAmount] = useState(""); // Amount of tokens to buy
  const [buying, setBuying] = useState(false); // Buying loading state
  const [totalSupply, setTotalSupply] = useState<bigint | null>(null); // Total token supply (from blockchain, for price calculations)
  const [ownerBalance, setOwnerBalance] = useState<bigint | null>(null); // Owner's token balance (kept for backward compatibility)
  // Token stats from API (accurate sold/available counts)
  const [tokenStats, setTokenStats] = useState<{
    total: number;
    sold: number;
    available: number;
    percentage_sold: number;
  } | null>(null);
  
  // Owner approval state
  const [pendingTransfers, setPendingTransfers] = useState<Array<{
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

  // Check for pending transfers if user is owner
  useEffect(() => {
    if (!property || !user || !address) return;
    
    const isOwner = address.toLowerCase() === property.OwnerWallet?.toLowerCase();
    if (!isOwner) return;

    const fetchPendingTransfers = async () => {
      try {
        const transfers = await api.getPendingTokenPurchases(property.ID);
        setPendingTransfers(transfers);
        console.log(`📋 Found ${transfers.length} pending transfer(s) for this property`);
      } catch (err) {
        console.error("Failed to fetch pending transfers:", err);
      }
    };

    // Fetch immediately
    fetchPendingTransfers();

    // Poll periodically to check for new pending transfers
    const interval = setInterval(fetchPendingTransfers, 5000);
    
    return () => {
      clearInterval(interval);
    };
  }, [property, user, address]);

  // Auto-connect wallet for registered users (only attempt once)
  useEffect(() => {
    const autoConnectWallet = async () => {
      if (user && user.WalletAddress && !autoConnectAttempted) {
        setAutoConnectAttempted(true);

        // Check if we're already connected to the correct wallet
        const isCorrectWallet = address && address.toLowerCase() === user.WalletAddress.toLowerCase();

        if (!isCorrectWallet) {
          console.log('🔗 Auto-connecting registered wallet for user:', user.Email);
          try {
            await connectRegisteredWallet(user.WalletAddress);
          } catch (err) {
            console.error('❌ Auto-connect failed:', err);
            // Don't retry, just log the error
          }
        }
      }
    };

    autoConnectWallet();
  }, [user, autoConnectAttempted, address, connectRegisteredWallet]);

  useEffect(() => {
    const fetchProperty = async () => {
      if (!id) return;
      try {
        const data = await api.getProperty(id);
        setProperty(data);
      } catch (err: any) {
        console.error("Failed to fetch property:", err);
        setError("Failed to load property details");
      } finally {
        setLoading(false);
      }
    };
    fetchProperty();
  }, [id]);

  // Function to refresh token availability data
  const refreshTokenAvailability = useCallback(async () => {
    if (!property || !property.ID) {
      return;
    }

    try {
      console.log("🔄 Refreshing token availability data...");
      
      // Fetch token stats from API (accurate sold/available counts from database)
      try {
        const stats = await api.getPropertyTokenStats(property.ID);
        setTokenStats(stats);
        console.log("✅ Token stats from API:", stats);
      } catch (apiErr) {
        console.error("Failed to fetch token stats from API:", apiErr);
        // Fallback: set default values if API fails
        setTokenStats({ total: 0, sold: 0, available: 0, percentage_sold: 0 });
      }

      // Still fetch total supply from blockchain for price calculations
      if (provider && property.OnchainTokenAddress) {
        const supply = await getTotalSupply(property.OnchainTokenAddress, provider);
        setTotalSupply(supply);

        // Also fetch owner's balance for backward compatibility (if needed elsewhere)
        if (property.OwnerWallet) {
          const ownerBal = await getTokenBalance(
            property.OnchainTokenAddress,
            property.OwnerWallet,
            provider
          );
          setOwnerBalance(ownerBal);
        }
      }
    } catch (err) {
      console.error("Failed to refresh token availability:", err);
    }
  }, [property, provider]);

  // Fetch token balance when wallet is connected
  useEffect(() => {
    const fetchBalance = async () => {
      if (!property || !isConnected || !address || !provider) {
        setTokenBalance("");
        return;
      }

      setBalanceLoading(true);
      try {
        const balance = await getTokenBalance(
          property.OnchainTokenAddress,
          address,
          provider
        );
        const formattedBalance = formatTokenBalance(balance);
        setTokenBalance(formattedBalance);
        console.log(`💰 User token balance for property ${property.ID}:`, {
          raw: balance.toString(),
          formatted: formattedBalance,
          address: address,
        });
      } catch (err) {
        console.error("Failed to fetch token balance:", err);
        setTokenBalance("0");
      } finally {
        setBalanceLoading(false);
      }
    };

    fetchBalance();
    
    // Poll for token balance changes every 10 seconds (in case owner transfers tokens)
    const pollInterval = setInterval(() => {
      if (property && isConnected && address && provider) {
        fetchBalance();
        // Also refresh token availability to update sold count
        refreshTokenAvailability();
      }
    }, 10000); // Check every 10 seconds

    return () => clearInterval(pollInterval);
  }, [property, isConnected, address, provider, refreshTokenAvailability]);

  // Fetch total supply and owner balance for buy tokens calculation
  useEffect(() => {
    refreshTokenAvailability();
  }, [refreshTokenAvailability]);

  // Refresh token availability when page regains focus (user might have received tokens)
  useEffect(() => {
    const handleFocus = () => {
      if (property && provider) {
        refreshTokenAvailability();
      }
    };

    window.addEventListener('focus', handleFocus);
    return () => window.removeEventListener('focus', handleFocus);
  }, [property, provider, refreshTokenAvailability]);

  // Check approval status
  useEffect(() => {
    const checkApproval = async () => {
      if (!isConnected || !address || !provider) {
        setIsApproved(null);
        return;
      }

      try {
        const approved = await checkApprovalStatus(address, provider);
        setIsApproved(approved);
      } catch (err) {
        console.error("Failed to check approval status:", err);
        setIsApproved(false);
      }
    };

    checkApproval();
  }, [isConnected, address, provider]);

  // Fetch property metadata
  useEffect(() => {
    const fetchMetadata = async () => {
      if (!property?.ID) return;

      setMetadataLoading(true);
      try {
        const metadataData = await api.getPropertyMetadata(property.ID);
        setMetadata(metadataData);
      } catch (err) {
        console.log("No metadata available or failed to fetch:", err);
        setMetadata(null);
      } finally {
        setMetadataLoading(false);
      }
    };

    fetchMetadata();
  }, [property?.ID]);

  const handleTransfer = async () => {
    if (!signer || !property || !transferTo || !transferAmount) {
      setError("Please fill all transfer fields and connect wallet");
      return;
    }

    setTransferring(true);
    setError("");
    try {
      // Convert amount to proper units (assuming 18 decimals)
      const amount = ethers.parseUnits(transferAmount, 18);

      const tx = await transferTokens(
        property.OnchainTokenAddress,
        transferTo,
        amount,
        signer
      );

      // Wait for transaction
      await tx.wait();

      // Refresh balance
      if (address && provider) {
        const balance = await getTokenBalance(
          property.OnchainTokenAddress,
          address,
          provider
        );
        setTokenBalance(formatTokenBalance(balance));
      }

      // Clear form
      setTransferTo("");
      setTransferAmount("");

      alert("Transfer successful!");
    } catch (err: any) {
      setError(err.message || "Failed to transfer tokens");
      console.error("Failed to transfer tokens:", err);
    } finally {
      setTransferring(false);
    }
  };

  const handleBuyTokens = async () => {
    if (!signer || !property || !buyTokenAmount || !address || !provider) {
      setError("Please enter token amount and connect wallet");
      return;
    }

    if (!isManuallyConnected) {
      setError("Please connect your wallet first");
      return;
    }

    // Check if user is the owner
    const isOwner = address.toLowerCase() === property.OwnerWallet?.toLowerCase();
    if (isOwner) {
      setError("You cannot buy tokens from your own property");
      return;
    }

    // Check if tokens are available for purchase
    if (tokenStats && tokenStats.available <= 0) {
      setError("No tokens available for purchase. All tokens have been sold.");
      setBuying(false);
      return;
    }

    setBuying(true);
    setError("");

    try {
      // Convert token amount to BigInt (assuming 18 decimals)
      const tokenAmount = ethers.parseUnits(buyTokenAmount, 18);

      // Validate owner has enough tokens
      if (ownerBalance !== null && tokenAmount > ownerBalance) {
        setError(`Owner only has ${formatTokenBalance(ownerBalance)} tokens available`);
        setBuying(false);
        return;
      }

      // Calculate price: (valuation / totalSupply) * tokenAmount
      // Valuation is in ETH, we need to convert to wei
      const valuationStr = toFixedString(property.Valuation);
      const valuationWei = ethers.parseEther(valuationStr);
      if (totalSupply === null || totalSupply === 0n) {
        setError("Unable to calculate price: total supply not available");
        setBuying(false);
        return;
      }

      // Price per token = valuation / totalSupply
      // Total price = (valuation * tokenAmount) / totalSupply
      // This is more accurate: (valuationWei * tokenAmount) / totalSupply
      const totalPriceSystemETH = (valuationWei * tokenAmount) / totalSupply;
      
      // Convert from system ETH to SepoliaETH: divide by conversion factor
      // totalPriceSystemETH is in wei, so we need to convert properly
      const totalPriceSepoliaETH = totalPriceSystemETH / BigInt(SEPOLIA_CONVERSION_FACTOR);

      console.log(`💰 Purchase Details:`, {
        tokenAmount: buyTokenAmount,
        tokenAmountWei: tokenAmount.toString(),
        valuation: property.Valuation,
        valuationWei: valuationWei.toString(),
        totalSupply: totalSupply.toString(),
        totalPriceSystemETH: ethers.formatEther(totalPriceSystemETH),
        totalPriceSepoliaETH: ethers.formatEther(totalPriceSepoliaETH),
        conversionFactor: SEPOLIA_CONVERSION_FACTOR,
        sellerAddress: property.OwnerWallet,
      });

      // Send SepoliaETH to seller's wallet address
      console.log(`💸 Sending ${ethers.formatEther(totalPriceSepoliaETH)} SepoliaETH to seller: ${property.OwnerWallet}`);
      
      const tx = await signer.sendTransaction({
        to: property.OwnerWallet,
        value: totalPriceSepoliaETH,
      });

      console.log(`📤 Transaction sent: ${tx.hash}`);
      console.log(`⏳ Waiting for transaction confirmation...`);

      // Wait for transaction to be confirmed
      const receipt = await tx.wait();
      console.log(`✅ Transaction confirmed in block: ${receipt?.blockNumber}`);

      // Record purchase in database
      try {
        await api.createTokenPurchase(
          property.ID,
          address, // Buyer's wallet
          buyTokenAmount, // Token amount
          tx.hash, // Payment transaction hash
          ethers.formatEther(totalPriceSepoliaETH) // Purchase price in ETH
        );
        console.log(`✅ Purchase recorded in database. Waiting for owner approval...`);
      } catch (purchaseErr: any) {
        console.error("❌ Failed to record purchase in database:", purchaseErr);
        // Still show success message even if DB recording fails
      }
      
      // Show success message to buyer
      alert(`✅ Payment successful! ${ethers.formatEther(totalPriceSepoliaETH)} SepoliaETH has been sent to the property owner.\n\nTransaction Hash: ${tx.hash}\n\n⏳ Waiting for owner approval...\n\nThe property owner will receive a notification to approve the token transfer. Once approved, ${buyTokenAmount} tokens will be transferred to your wallet.\n\nYour wallet address: ${address}`);
      
      // Clear form
      setBuyTokenAmount("");

      // Refresh all balances and token availability data
      if (address && provider && property) {
        try {
          // Refresh user's token balance
          const balance = await getTokenBalance(
            property.OnchainTokenAddress,
            address,
            provider
          );
          setTokenBalance(formatTokenBalance(balance));
          console.log("✅ User token balance refreshed:", formatTokenBalance(balance));

          // Refresh token availability (total supply and owner balance)
          await refreshTokenAvailability();
        } catch (err) {
          console.error("Failed to refresh token data after purchase:", err);
        }
      }
    } catch (err: any) {
      setError(err.message || "Failed to purchase tokens");
      console.error("Failed to purchase tokens:", err);
    } finally {
      setBuying(false);
    }
  };

  // Handle owner approval of token transfer
  const handleApproveTransfer = async (transfer: any) => {
    if (!signer || !property || !provider) {
      setError("Please connect your wallet to approve transfers");
      return;
    }

    // Verify owner is connected
    const isOwner = address?.toLowerCase() === property.OwnerWallet?.toLowerCase();
    if (!isOwner) {
      setError("Only the property owner can approve transfers");
      return;
    }

    setApprovingTransfer(transfer.id);

    try {
      console.log(`✅ Owner approving transfer: ${transfer.amount} tokens to ${transfer.buyer_wallet}`);
      
      // Convert amount to wei (amount is in token units with 18 decimals)
      const tokenAmountWei = ethers.parseUnits(transfer.amount, 18);
      
      // Transfer tokens from owner to buyer
      const transferTx = await transferTokens(
        property.OnchainTokenAddress,
        transfer.buyer_wallet,
        tokenAmountWei,
        signer
      );

      console.log(`📤 Token transfer transaction sent: ${transferTx.hash}`);
      console.log(`⏳ Waiting for confirmation...`);

      // Wait for confirmation
      const receipt = await transferTx.wait();
      console.log(`✅ Token transfer confirmed in block: ${receipt?.blockNumber}`);

      // Update database with token transfer hash
      try {
        await api.updateTokenPurchaseTxHash(
          property.ID,
          transfer.id,
          transferTx.hash
        );
        console.log(`✅ Purchase record updated with token transfer hash`);
      } catch (updateErr: any) {
        console.error("❌ Failed to update purchase record:", updateErr);
        // Continue even if DB update fails - tokens are already transferred
      }

      // Refresh pending transfers list
      const updatedTransfers = await api.getPendingTokenPurchases(property.ID);
      setPendingTransfers(updatedTransfers);

      alert(`✅ Transfer Approved!\n\n${transfer.amount} tokens have been transferred to ${transfer.buyer_wallet}\n\nToken Transfer TX: ${transferTx.hash}\nPayment TX: ${transfer.payment_tx_hash}`);

      // Refresh token balances
      await refreshTokenAvailability();
    } catch (err: any) {
      console.error("Failed to approve transfer:", err);
      setError(`Failed to approve transfer: ${err.message || "Unknown error"}`);
    } finally {
      setApprovingTransfer(null);
    }
  };

  // Helper function to convert number to string without scientific notation
  const toFixedString = (num: number): string => {
    if (num === 0) return "0";
    // Use toFixed with enough decimals, then remove trailing zeros
    const str = num.toFixed(18);
    return str.replace(/\.?0+$/, '');
  };

  // Calculate price per token and total price
  const calculatePrice = () => {
    if (!property || !totalSupply || totalSupply === 0n || !buyTokenAmount) {
      return { pricePerToken: "0", totalPrice: "0", pricePerTokenSepolia: "0", totalPriceSepolia: "0" };
    }

    try {
      const valuationStr = toFixedString(property.Valuation);
      const valuationWei = ethers.parseEther(valuationStr);
      const tokenAmount = ethers.parseUnits(buyTokenAmount, 18);
      
      // Price per token = valuation / totalSupply (in system ETH)
      const pricePerTokenWei = valuationWei / totalSupply;
      const pricePerToken = Number(ethers.formatEther(pricePerTokenWei));
      
      // Total price = (valuation * tokenAmount) / totalSupply (in system ETH)
      const totalPriceWei = (valuationWei * tokenAmount) / totalSupply;
      const totalPrice = ethers.formatEther(totalPriceWei);

      // Convert to SepoliaETH: divide by conversion factor
      const pricePerTokenSepolia = pricePerToken / SEPOLIA_CONVERSION_FACTOR;
      const totalPriceSepolia = Number(totalPrice) / SEPOLIA_CONVERSION_FACTOR;

      return { 
        pricePerToken: pricePerToken.toFixed(8), 
        totalPrice,
        pricePerTokenSepolia: pricePerTokenSepolia.toFixed(8),
        totalPriceSepolia: totalPriceSepolia.toFixed(8)
      };
    } catch (err) {
      console.error("Error calculating price:", err);
      return { pricePerToken: "0", totalPrice: "0", pricePerTokenSepolia: "0", totalPriceSepolia: "0" };
    }
  };

  const priceInfo = calculatePrice();
  const isOwner = address && property?.OwnerWallet && address.toLowerCase() === property.OwnerWallet.toLowerCase();

  if (loading) {
    return (
      <div className="min-h-screen bg-[#0f0f0f] text-white flex items-center justify-center">
        <Navbar currentPage="/properties" />
        <p>Loading...</p>
      </div>
    );
  }

  if (error || !property) {
    return (
      <div className="min-h-screen bg-[#0f0f0f] text-white flex items-center justify-center">
        <Navbar currentPage="/properties" />
        <div className="text-center">
            <h1 className="text-2xl font-bold mb-4">Property Not Found</h1>
            <p className="text-gray-400 mb-6">{error || "The property you are looking for does not exist."}</p>
            <RouterLink to="/dashboard">
                <Button>Go to Dashboard</Button>
            </RouterLink>
        </div>
      </div>
    );
  }

  // Use property name if available, otherwise use ID
  const title = property.Name || metadata?.name || `Property #${property.ID.substring(0, 8)}`;
  const description = metadata?.description || "No description available for this property.";
  const location = metadata?.location || "Location not specified";
  const features = ["Tokenized Asset", "Blockchain Verified", "Instant Settlement"];
  const details = {
    "Property Type": "Tokenized Real Estate",
    "Status": property.Status,
    "Created": new Date(property.CreatedAt).toLocaleDateString(),
    "Owner": `${property.OwnerWallet.substring(0, 6)}...${property.OwnerWallet.substring(38)}`,
    "Token Address": `${property.OnchainTokenAddress.substring(0, 6)}...${property.OnchainTokenAddress.substring(38)}`,
  };

  // Calculate token availability dynamically
  // Use API stats if available (more accurate), otherwise fallback to blockchain calculation
  const tokenTotal = tokenStats?.total || (totalSupply ? parseFloat(formatTokenBalance(totalSupply)) : 0);
  const tokenSold = tokenStats?.sold || (tokenTotal > 0 && ownerBalance !== null ? Math.max(0, tokenTotal - parseFloat(formatTokenBalance(ownerBalance))) : 0);
  const tokenAvailable = tokenStats?.available || (ownerBalance !== null ? parseFloat(formatTokenBalance(ownerBalance)) : 0);
  const tokenPercentage = tokenStats?.percentage_sold || (tokenTotal > 0 ? (tokenSold / tokenTotal) * 100 : 0);
  
  // Calculate price per token in system ETH (for display)
  const pricePerTokenETH = priceInfo.pricePerToken !== "0" ? priceInfo.pricePerToken : 
    (property && totalSupply && totalSupply > 0n 
      ? (() => {
          try {
            const valuationStr = toFixedString(property.Valuation);
            const valuationWei = ethers.parseEther(valuationStr);
            const pricePerTokenWei = valuationWei / totalSupply;
            return Number(ethers.formatEther(pricePerTokenWei)).toFixed(8);
          } catch (err) {
            console.error("Error calculating price per token:", err);
            return "0";
          }
        })()
      : "0");

  // Calculate ROI (placeholder - you can enhance this with actual data)
  const currentValuation = property.Valuation;
  const roi = "+10%"; // This can be calculated from actual data if available

  return (
    <div className="min-h-screen w-full bg-[#0f0f0f] text-white font-sans selection:bg-[#6d41ff] selection:text-white">
      <Navbar currentPage="/properties" />

      <main className="max-w-7xl mx-auto px-4 pt-28 pb-16 space-y-10">
        {/* Pending Transfers Notification (Owner Only) */}
        {isOwner && pendingTransfers.length > 0 && (
          <section className="bg-yellow-900/20 border border-yellow-600/50 rounded-lg p-6 space-y-4">
            <div className="flex items-center gap-2">
              <span className="text-yellow-400 font-semibold text-lg">⚠️ Pending Token Transfers</span>
              <span className="bg-yellow-600 text-black px-2 py-1 rounded text-sm font-bold">
                {pendingTransfers.length}
              </span>
            </div>
            <p className="text-gray-300 text-sm">
              Buyers have sent payments and are waiting for you to approve token transfers. Please review and approve each transfer.
            </p>
            <div className="space-y-3 mt-4">
              {pendingTransfers.map((transfer: any) => (
                <div
                  key={transfer.id}
                  className="bg-[#1a1a1a] border border-gray-700 rounded-lg p-4 space-y-3"
                >
                  <div className="flex justify-between items-start">
                    <div className="space-y-1">
                      <p className="text-sm text-gray-400">Buyer Address</p>
                      <p className="font-mono text-sm">{transfer.buyer_wallet}</p>
                    </div>
                    <div className="text-right space-y-1">
                      <p className="text-sm text-gray-400">Amount</p>
                      <p className="font-semibold">{transfer.amount} tokens</p>
                    </div>
                  </div>
                  <div className="grid grid-cols-2 gap-4 text-sm">
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
                      onClick={() => handleApproveTransfer(transfer)}
                      disabled={approvingTransfer === transfer.id}
                      className="flex-1 bg-green-600 hover:bg-green-700"
                    >
                      {approvingTransfer === transfer.id ? "Approving..." : "✅ Approve Transfer"}
                    </Button>
                  </div>
                </div>
              ))}
            </div>
          </section>
        )}

        {/* Header Section */}
        <section className="space-y-6">
          <div className="flex flex-col md:flex-row md:items-end justify-between gap-4">
            <div className="space-y-3">
              <h1 className="text-3xl md:text-4xl font-semibold">{title}</h1>
              <div className="flex items-center gap-2 text-gray-400 text-sm">
                <MapPin size={16} className="text-[#6d41ff]" />
                <span>{location}</span>
              </div>
            </div>
            <div className="text-left md:text-right space-y-1">
              <p className="text-gray-400 text-sm">Total Property Value</p>
              <p className="text-3xl font-semibold">{property.Valuation} ETH</p>
            </div>
          </div>

          {/* Image Gallery */}
          {metadata?.images && metadata.images.length > 0 ? (
            <div className="grid grid-cols-1 md:grid-cols-4 gap-4 h-[400px] md:h-[500px]">
              <div className="md:col-span-2 md:row-span-2 rounded-3xl overflow-hidden border border-[#1f1f1f] relative group">
                <img
                  src={metadata.images[0]}
                  alt="Property main"
                  className="w-full h-full object-cover group-hover:scale-105 transition-transform duration-500"
                  onError={(e) => {
                    e.currentTarget.style.display = 'none';
                    e.currentTarget.nextElementSibling!.classList.remove('hidden');
                  }}
                />
                <div className="absolute top-4 left-4 bg-[#6d41ff] px-3 py-1.5 rounded-full text-xs font-semibold">
                  Featured
                </div>
                <div className="absolute inset-0 flex items-center justify-center pointer-events-none hidden">
                  <GradientThumb />
                  <span className="absolute bg-black/50 px-4 py-2 rounded-lg backdrop-blur-sm">Image not available</span>
                </div>
              </div>
              {metadata.images.slice(1, 4).map((image, index) => (
                <div
                  key={index}
                  className="rounded-3xl overflow-hidden border border-[#1f1f1f] relative group"
                >
                  <img
                    src={image}
                    alt={`Property ${index + 2}`}
                    className="w-full h-full object-cover group-hover:scale-105 transition-transform duration-500"
                    onError={(e) => {
                      e.currentTarget.style.display = 'none';
                      e.currentTarget.nextElementSibling!.classList.remove('hidden');
                    }}
                  />
                  <div className="absolute inset-0 flex items-center justify-center pointer-events-none hidden">
                    <GradientThumb />
                  </div>
                </div>
              ))}
            </div>
          ) : (
            <div className="grid grid-cols-1 md:grid-cols-4 gap-4 h-[400px] md:h-[500px]">
              <div className="md:col-span-2 md:row-span-2 rounded-3xl overflow-hidden border border-[#1f1f1f] relative group">
                <GradientThumb />
                <div className="absolute inset-0 flex items-center justify-center pointer-events-none">
                  <span className="bg-black/50 px-4 py-2 rounded-lg backdrop-blur-sm">
                    {metadataLoading ? "Loading images..." : "No images available"}
                  </span>
                </div>
              </div>
              <div className="rounded-3xl overflow-hidden border border-[#1f1f1f] relative group">
                <GradientThumb />
              </div>
              <div className="rounded-3xl overflow-hidden border border-[#1f1f1f] relative group">
                <GradientThumb />
              </div>
              <div className="rounded-3xl overflow-hidden border border-[#1f1f1f] relative group">
                <GradientThumb />
              </div>
            </div>
          )}
        </section>

        {/* Details Grid */}
        <section className="grid md:grid-cols-3 gap-8">
          {/* Left Column: Description & Features */}
          <div className="md:col-span-2 space-y-8">
            {/* Current Valuation Badge */}
            <div className="rounded-3xl border border-[#1f1f1f] bg-gradient-to-br from-[#6d41ff]/10 via-[#111111] to-[#111111] p-6">
              <div className="flex items-center justify-between">
                <div className="space-y-1">
                  <p className="text-sm text-gray-400">Current Valuation</p>
                  <p className="text-3xl font-semibold">{currentValuation} ETH</p>
                </div>
                <div className="text-right space-y-1">
                  <p className="text-sm text-gray-400">Return on Investment</p>
                  <p className="text-2xl font-semibold text-green-400 flex items-center gap-2">
                    <TrendingUp size={24} />
                    {roi}
                  </p>
                </div>
              </div>
            </div>

            {/* Description */}
            <div className="space-y-4 rounded-3xl border border-[#1f1f1f] bg-[#111111] p-6">
              <div className="flex items-center gap-2">
                <Building2 size={20} className="text-[#6d41ff]" />
                <h2 className="text-xl font-semibold">Description</h2>
              </div>
              <p className="text-gray-300 leading-relaxed">{description}</p>
            </div>

            {/* Key Features */}
            <div className="space-y-4 rounded-3xl border border-[#1f1f1f] bg-[#111111] p-6">
              <h2 className="text-xl font-semibold">Key Features</h2>
              <ul className="grid sm:grid-cols-2 gap-3">
                {features.map((feature, index) => (
                  <li
                    key={index}
                    className="flex items-center gap-3 text-gray-300 bg-[#1a1a1a] p-3 rounded-xl border border-[#262626] hover:border-[#6d41ff]/50 transition-colors"
                  >
                    <span className="h-2 w-2 rounded-full bg-[#6d41ff]" />
                    {feature}
                  </li>
                ))}
              </ul>
            </div>

            {/* Details & Stats */}
            <div className="space-y-4 rounded-3xl border border-[#1f1f1f] bg-[#111111] p-6">
              <h2 className="text-xl font-semibold">Details & Stats</h2>
              <div className="grid sm:grid-cols-2 gap-x-8 gap-y-4">
                {Object.entries(details).map(([key, value], index) => (
                  <div
                    key={index}
                    className="flex justify-between items-center border-b border-[#262626] pb-3 last:border-0"
                  >
                    <span className="text-gray-400 text-sm">{key}</span>
                    <span className="font-medium text-sm">{value}</span>
                  </div>
                ))}
              </div>
            </div>
          </div>

          {/* Right Column: Investment Card with Token Availability */}
          <div className="space-y-6">
            <div className="rounded-3xl border border-[#1f1f1f] bg-[#111111] p-6 space-y-6 sticky top-28">
              {/* Token Availability Section */}
              <div className="space-y-4">
                <div className="flex items-center justify-between">
                  <h3 className="text-lg font-semibold">Token Availability</h3>
                  <div className="bg-[#6d41ff]/20 text-[#6d41ff] px-3 py-1 rounded-full text-xs font-semibold">
                    {tokenPercentage.toFixed(1)}% Sold
                  </div>
                </div>

                {/* Token Stats */}
                <div className="grid grid-cols-3 gap-3">
                  <div className="bg-[#1a1a1a] rounded-xl p-3 border border-[#262626]">
                    <p className="text-xs text-gray-400 mb-1">Total</p>
                    <p className="text-lg font-semibold">{tokenTotal.toLocaleString()}</p>
                  </div>
                  <div className="bg-[#1a1a1a] rounded-xl p-3 border border-green-500/20">
                    <p className="text-xs text-gray-400 mb-1">Sold</p>
                    <p className="text-lg font-semibold text-green-400">
                      {tokenSold.toLocaleString()}
                    </p>
                  </div>
                  <div className="bg-[#1a1a1a] rounded-xl p-3 border border-[#6d41ff]/20">
                    <p className="text-xs text-gray-400 mb-1">Available</p>
                    <p className="text-lg font-semibold text-[#6d41ff]">
                      {tokenAvailable.toLocaleString()}
                    </p>
                  </div>
                </div>

                {/* Animated Progress Bar */}
                <div className="space-y-2">
                  <div className="flex justify-between text-xs text-gray-400">
                    <span>Token Sale Progress</span>
                    <span>{tokenSold.toLocaleString()} / {tokenTotal.toLocaleString()}</span>
                  </div>
                  <div className="relative h-3 bg-[#1a1a1a] rounded-full overflow-hidden border border-[#262626]">
                    {/* Animated gradient background */}
                    <div className="absolute inset-0 opacity-20">
                      <div
                        className="h-full bg-gradient-to-r from-[#6d41ff] via-[#8b5cf6] to-[#6d41ff] animate-pulse"
                        style={{
                          backgroundSize: "200% 100%",
                          animation: "gradient 3s ease infinite",
                        }}
                      />
                    </div>
                    {/* Progress fill with fluid gradient */}
                    <div
                      className="relative h-full bg-gradient-to-r from-[#6d41ff] via-[#8b5cf6] via-[#a78bfa] to-[#6d41ff] rounded-full transition-all duration-1000 ease-out"
                      style={{
                        width: `${tokenPercentage}%`,
                        backgroundSize: "200% 100%",
                        animation: "gradient 4s ease infinite",
                      }}
                    >
                      {/* Shimmer effect */}
                      <div className="absolute inset-0 bg-gradient-to-r from-transparent via-white/20 to-transparent animate-shimmer" />
                    </div>
                  </div>
                </div>

                {/* Price per token */}
                <div className="bg-[#1a1a1a] rounded-xl p-4 border border-[#262626]">
                  <div className="flex justify-between items-center">
                    <span className="text-sm text-gray-400">Price per Token</span>
                    <span className="text-xl font-semibold">{pricePerTokenETH} ETH</span>
                  </div>
                </div>
              </div>

              <div className="h-px bg-[#262626]" />

              {/* Investment Info */}
              <div className="space-y-3">
                <div className="flex justify-between text-sm">
                  <span className="text-gray-400">Minimum Investment</span>
                  <span className="font-semibold">10 tokens ({pricePerTokenETH !== "0" ? (parseFloat(pricePerTokenETH) * 10).toFixed(6) : "0"} ETH)</span>
                </div>
                <div className="flex justify-between text-sm">
                  <span className="text-gray-400">Expected Annual Yield</span>
                  <span className="font-semibold text-green-400">8-12%</span>
                </div>
                <div className="flex justify-between text-sm">
                  <span className="text-gray-400">Investment Type</span>
                  <span className="font-semibold">Tokenized Real Estate</span>
                </div>
              </div>

              {/* User-specific sections */}
              {isConnected && address && (
                <>
                  <div className="h-px bg-[#262626]" />
                  <div className="space-y-3">
                    <div className="flex justify-between text-sm">
                      <span className="text-gray-400">Your Token Balance</span>
                      <span className="font-semibold">
                        {balanceLoading ? "Loading..." : tokenBalance || "0"} tokens
                      </span>
                    </div>
                    <div className="flex justify-between text-sm">
                      <span className="text-gray-400">Approval Status</span>
                      <span>
                        {isApproved === null
                          ? "--"
                          : isApproved
                          ? "✅ Approved"
                          : "❌ Not Approved"}
                      </span>
                    </div>
                  </div>
                </>
              )}

              {/* Buy Tokens Section - Only show if user is NOT the owner */}
              {!isOwner && isManuallyConnected && (
                <>
                  <div className="h-px bg-[#262626]" />
                  <div className="space-y-3">
                    <p className="text-sm font-semibold text-gray-300">Purchase Tokens</p>
                    <div className="space-y-2">
                      <input
                        type="number"
                        placeholder="Token amount to buy"
                        className="w-full rounded-lg border border-[#262626] bg-[#0f0f0f] px-3 py-2 text-sm text-white"
                        value={buyTokenAmount}
                        onChange={(e) => setBuyTokenAmount(e.target.value)}
                        min="0"
                        step="0.000001"
                      />
                      {buyTokenAmount && priceInfo.totalPrice !== "0" && (
                        <div className="text-xs text-gray-400 space-y-1 bg-[#1a1a1a] p-2 rounded-lg">
                          <div>Price per token: {priceInfo.pricePerToken} ETH</div>
                          <div className="font-semibold text-white">
                            Total Cost: {priceInfo.totalPrice} ETH
                          </div>
                          <div className="text-xs text-gray-500 mt-1">
                            (Will send {priceInfo.totalPriceSepolia} SepoliaETH to seller)
                          </div>
                          <div className="text-yellow-400">
                            Will be sent to: {property.OwnerWallet.substring(0, 6)}...{property.OwnerWallet.substring(38)}
                          </div>
                        </div>
                      )}
                      <p className="text-xs text-yellow-400">
                        ⚠️ Note: After payment is sent, the property owner will need to transfer the tokens to your address ({address?.substring(0, 6)}...{address?.substring(38)}).
                      </p>
                    </div>
                  </div>
                </>
              )}

              {/* Transfer Tokens Section - Only show if user owns tokens */}
              {parseFloat(tokenBalance || "0") > 0 && (
                <>
                  <div className="h-px bg-[#262626]" />
                  <div className="space-y-2">
                    <p className="text-sm font-semibold text-gray-300">Transfer Tokens</p>
                    <div className="space-y-2">
                      <input
                        type="text"
                        placeholder="Recipient address (0x...)"
                        className="w-full rounded-lg border border-[#262626] bg-[#0f0f0f] px-3 py-2 text-sm text-white"
                        value={transferTo}
                        onChange={(e) => setTransferTo(e.target.value)}
                      />
                      <input
                        type="number"
                        placeholder="Amount"
                        className="w-full rounded-lg border border-[#262626] bg-[#0f0f0f] px-3 py-2 text-sm text-white"
                        value={transferAmount}
                        onChange={(e) => setTransferAmount(e.target.value)}
                      />
                      <Button
                        variant="outline"
                        size="sm"
                        onClick={handleTransfer}
                        disabled={transferring}
                        className="w-full border-[#262626] bg-[#1A1A1A] text-white hover:bg-[#262626]"
                      >
                        {transferring ? "Transferring..." : "Transfer Tokens"}
                      </Button>
                    </div>
                  </div>
                </>
              )}

              {!isConnected && (
                <div className="p-3 rounded-lg border border-yellow-500/30 bg-yellow-500/10">
                  <p className="text-xs text-yellow-400 text-center">
                    Connect your wallet to view token balance
                  </p>
                </div>
              )}

              {error && (
                <div className="p-3 rounded-lg border border-red-500/30 bg-red-500/10">
                  <p className="text-xs text-red-400">{error}</p>
                </div>
              )}

              <div className="h-px bg-[#262626]" />

              <Button
                size="lg"
                className="w-full bg-gradient-to-r from-[#6d41ff] to-[#8b5cf6] hover:from-[#5b36d6] hover:to-[#7c3aed] text-white font-semibold shadow-lg shadow-[#6d41ff]/20"
                onClick={!isOwner && isManuallyConnected && buyTokenAmount && parseFloat(buyTokenAmount) > 0 ? handleBuyTokens : () => navigate(`/dashboard`)}
                disabled={!isOwner && isManuallyConnected && (buying || !buyTokenAmount || parseFloat(buyTokenAmount) <= 0)}
              >
                {buying ? "Processing..." : "Invest Now"}
              </Button>

              <p className="text-xs text-gray-400 text-center">
                Property is blockchain verified and tokenized
              </p>
            </div>
          </div>
        </section>

        {/* Additional Info Section */}
        <section className="rounded-3xl border border-[#1f1f1f] bg-[#111111] p-8">
          <div className="flex items-center gap-3 mb-6">
            <Home size={24} className="text-[#6d41ff]" />
            <h2 className="text-2xl font-semibold">Why Invest in This Property?</h2>
          </div>
          <div className="grid md:grid-cols-3 gap-6">
            <div className="space-y-2">
              <div className="bg-[#6d41ff]/10 w-12 h-12 rounded-xl flex items-center justify-center mb-3">
                <TrendingUp className="text-[#6d41ff]" size={24} />
              </div>
              <h3 className="font-semibold">Prime Location</h3>
              <p className="text-sm text-gray-400">
                Located in a desirable neighborhood with excellent growth potential
              </p>
            </div>
            <div className="space-y-2">
              <div className="bg-green-500/10 w-12 h-12 rounded-xl flex items-center justify-center mb-3">
                <Calendar className="text-green-400" size={24} />
              </div>
              <h3 className="font-semibold">Rental Income</h3>
              <p className="text-sm text-gray-400">
                Generate passive income through monthly rental distributions to token holders
              </p>
            </div>
            <div className="space-y-2">
              <div className="bg-blue-500/10 w-12 h-12 rounded-xl flex items-center justify-center mb-3">
                <Building2 className="text-blue-400" size={24} />
              </div>
              <h3 className="font-semibold">Fractional Ownership</h3>
              <p className="text-sm text-gray-400">
                Own a piece of luxury real estate without the hassle of property management
              </p>
            </div>
          </div>
        </section>
      </main>

      <style>{`
        @keyframes gradient {
          0%, 100% {
            background-position: 0% 50%;
          }
          50% {
            background-position: 100% 50%;
          }
        }

        @keyframes shimmer {
          0% {
            transform: translateX(-100%);
          }
          100% {
            transform: translateX(100%);
          }
        }

        .animate-shimmer {
          animation: shimmer 2s infinite;
        }
      `}</style>
    </div>
  );
}

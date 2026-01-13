import { useEffect, useState } from "react";
import { Button } from "../Components/atoms/Button";
import { api } from "../lib/api";
import { CheckCircle, XCircle, Clock, RefreshCw, FileText, Coins, Wallet } from "lucide-react";
import { useWallet } from "../hooks/useWallet";
import { ethers } from "ethers";
import { transferTokens } from "../lib/contracts";

interface PropertyUploadRequest {
  id: string;
  name: string;
  symbol: string;
  valuation: number;
  token_supply: number;
  status: "pending" | "approved" | "rejected";
  rejection_reason?: string;
  created_at: string;
  documents?: Array<{
    id: string;
    file_url: string;
    file_hash: string;
    name: string;
    type: string;
  }>;
}

interface TokenPurchase {
  id: string;
  property_id: string;
  buyer_wallet: string;
  amount: string;
  payment_tx_hash: string;
  token_tx_hash: string;
  purchase_price: string;
  created_at: string;
  property_name?: string;
}

function Messages() {
  const [requests, setRequests] = useState<PropertyUploadRequest[]>([]);
  const [pendingPurchases, setPendingPurchases] = useState<TokenPurchase[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");
  const [approvingPurchaseId, setApprovingPurchaseId] = useState<string | null>(null);
  const { signer, provider } = useWallet();

  useEffect(() => {
    loadAllMessages();
  }, []);

  const loadAllMessages = async () => {
    try {
      setLoading(true);
      setError("");
      
      // Load property upload requests
      try {
        const uploadRequests = await api.getMyPropertyUploadRequests();
        setRequests(uploadRequests);
      } catch (err: any) {
        console.warn("Failed to load property upload requests:", err);
      }

      // Load pending token purchases (for property owners)
      try {
        const purchases = await api.getMyPendingTransfers();
        console.log(`📦 Found ${purchases.length} pending token purchases`);
        
        // Fetch property names for each purchase
        const purchasesWithNames = await Promise.all(
          purchases.map(async (purchase: TokenPurchase) => {
            try {
              const property = await api.getProperty(purchase.property_id);
              return { ...purchase, property_name: property.Name || `Property ${purchase.property_id.substring(0, 8)}` };
            } catch (err) {
              console.warn(`Failed to fetch property ${purchase.property_id} for purchase ${purchase.id}:`, err);
              return { ...purchase, property_name: `Property ${purchase.property_id.substring(0, 8)}` };
            }
          })
        );
        
        setPendingPurchases(purchasesWithNames);
      } catch (err: any) {
        // 404 means endpoint doesn't exist or user has no pending purchases
        if (err.message?.includes("404")) {
          console.log("⚠️ Pending transfers endpoint returned 404 - either endpoint doesn't exist or no pending purchases");
          setPendingPurchases([]); // Set empty array instead of leaving it undefined
        } else {
          console.warn("Failed to load pending token purchases:", err);
        }
      }
    } catch (err: any) {
      setError(err.message || "Failed to load messages");
    } finally {
      setLoading(false);
    }
  };

  const handleApprovePurchase = async (purchase: TokenPurchase) => {
    if (!signer || !provider) {
      alert("Please connect your wallet to approve token transfers");
      return;
    }

    setApprovingPurchaseId(purchase.id);

    try {
      // Get property to get token address
      const property = await api.getProperty(purchase.property_id);
      
      console.log(`✅ Owner approving transfer: ${purchase.amount} tokens to ${purchase.buyer_wallet}`);
      
      // Convert amount to wei
      const tokenAmountWei = ethers.parseUnits(purchase.amount, 18);
      
      // Transfer tokens from owner to buyer
      const transferTx = await transferTokens(
        property.OnchainTokenAddress,
        purchase.buyer_wallet,
        tokenAmountWei,
        signer
      );

      console.log(`📤 Token transfer transaction sent: ${transferTx.hash}`);
      const receipt = await transferTx.wait();
      console.log(`✅ Token transfer confirmed in block: ${receipt?.blockNumber}`);

      // Update database with token transfer hash
      try {
        await api.updateTokenPurchaseTxHash(
          purchase.property_id,
          purchase.id,
          transferTx.hash
        );
        console.log(`✅ Purchase record updated with token transfer hash`);
      } catch (updateErr: any) {
        console.error("❌ Failed to update purchase record:", updateErr);
        alert("Tokens transferred but failed to update record. Please contact support.");
        return;
      }

      // Refresh messages
      await loadAllMessages();

      alert(`✅ Transfer Approved!\n\n${purchase.amount} tokens transferred to ${purchase.buyer_wallet}\n\nToken TX: ${transferTx.hash}\nPayment TX: ${purchase.payment_tx_hash}`);
    } catch (err: any) {
      console.error("Failed to approve transfer:", err);
      alert(`Failed to approve transfer: ${err.message || "Unknown error"}`);
    } finally {
      setApprovingPurchaseId(null);
    }
  };

  const getStatusBadge = (status: string) => {
    switch (status) {
      case 'pending':
        return (
          <span className="px-3 py-1 rounded-full text-xs font-medium bg-yellow-500/20 text-yellow-400 border border-yellow-500/30 flex items-center gap-1">
            <Clock size={12} />
            Pending Review
          </span>
        );
      case 'approved':
        return (
          <span className="px-3 py-1 rounded-full text-xs font-medium bg-green-500/20 text-green-400 border border-green-500/30 flex items-center gap-1">
            <CheckCircle size={12} />
            Approved
          </span>
        );
      case 'rejected':
        return (
          <span className="px-3 py-1 rounded-full text-xs font-medium bg-red-500/20 text-red-400 border border-red-500/30 flex items-center gap-1">
            <XCircle size={12} />
            Rejected
          </span>
        );
      default:
        return null;
    }
  };

  const getStatusMessage = (request: PropertyUploadRequest) => {
    switch (request.status) {
      case 'pending':
        return "Your property upload request is pending admin review. You will be notified once a decision is made.";
      case 'approved':
        return "Your property upload request has been approved! The property has been created on the blockchain.";
      case 'rejected':
        return request.rejection_reason 
          ? `Your property upload request was rejected. Reason: ${request.rejection_reason}`
          : "Your property upload request was rejected. Please contact support for more information.";
      default:
        return "";
    }
  };

  const formatDate = (dateString: string) => {
    return new Date(dateString).toLocaleDateString('en-US', {
      year: 'numeric',
      month: 'short',
      day: 'numeric',
      hour: '2-digit',
      minute: '2-digit',
    });
  };

  if (loading) {
    return (
      <div className="min-h-screen bg-[#0f0f0f] text-white flex items-center justify-center">
        <p className="text-gray-400">Loading your messages...</p>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-[#0f0f0f] text-white">
      <main className="max-w-4xl mx-auto px-6 pt-8 pb-16">
        <div className="space-y-8">
          {/* Header */}
          <div className="flex items-center justify-between">
            <div>
              <h1 className="text-3xl font-semibold">Messages</h1>
              <p className="text-gray-400 mt-2">
                View property upload requests and pending token purchase approvals
              </p>
            </div>
            <Button
              variant="outline"
              size="md"
              onClick={loadAllMessages}
              className="border-[#262626] bg-[#1A1A1A] text-white hover:bg-[#262626]"
            >
              <RefreshCw size={16} className="mr-2" />
              Refresh
            </Button>
          </div>

          {/* Error Display */}
          {error && (
            <div className="text-red-500 bg-red-500/10 p-4 rounded-xl border border-red-500/20">
              {error}
            </div>
          )}

          {/* Pending Token Purchases Section (for Property Owners) */}
          {pendingPurchases.length > 0 && (
            <div className="space-y-4">
              <h2 className="text-xl font-semibold flex items-center gap-2">
                <Coins size={20} />
                Pending Token Purchases ({pendingPurchases.length})
              </h2>
              <p className="text-gray-400 text-sm">
                Buyers have purchased tokens for your properties. Approve to transfer tokens to them.
              </p>
              {pendingPurchases.map((purchase) => (
                <div
                  key={purchase.id}
                  className="bg-[#1A1A1A] border border-[#262626] rounded-xl p-6 space-y-4"
                >
                  <div className="flex items-start justify-between">
                    <div className="flex-1">
                      <div className="flex items-center gap-3 mb-3">
                        <h3 className="text-xl font-semibold">{purchase.property_name}</h3>
                        <span className="px-3 py-1 rounded-full text-xs font-medium bg-yellow-500/20 text-yellow-400 border border-yellow-500/30 flex items-center gap-1">
                          <Clock size={12} />
                          Pending Approval
                        </span>
                      </div>
                      
                      <div className="grid md:grid-cols-2 gap-4 mb-4 text-sm">
                        <div>
                          <span className="text-gray-400">Buyer Wallet:</span>
                          <p className="text-gray-300 mt-1 font-mono text-xs">{purchase.buyer_wallet}</p>
                        </div>
                        <div>
                          <span className="text-gray-400">Amount:</span>
                          <p className="text-gray-300 mt-1">{purchase.amount} tokens</p>
                        </div>
                        <div>
                          <span className="text-gray-400">Purchase Price:</span>
                          <p className="text-gray-300 mt-1">{purchase.purchase_price} ETH</p>
                        </div>
                        <div>
                          <span className="text-gray-400">Payment TX:</span>
                          <p className="text-gray-300 mt-1 font-mono text-xs break-all">{purchase.payment_tx_hash}</p>
                        </div>
                        <div>
                          <span className="text-gray-400">Purchase Date:</span>
                          <p className="text-gray-300 mt-1">{formatDate(purchase.created_at)}</p>
                        </div>
                      </div>

                      <div className="p-4 rounded-lg border bg-yellow-500/10 border-yellow-500/20 mb-4">
                        <p className="text-sm text-yellow-400">
                          A buyer has purchased {purchase.amount} tokens for this property. Click "Approve & Transfer Tokens" to send the tokens to the buyer's wallet.
                        </p>
                      </div>

                      <Button
                        variant="primary"
                        size="md"
                        onClick={() => handleApprovePurchase(purchase)}
                        disabled={approvingPurchaseId === purchase.id}
                        className="bg-[#6d41ff] hover:bg-[#5b2fff] text-white"
                      >
                        {approvingPurchaseId === purchase.id ? (
                          <>
                            <RefreshCw size={16} className="mr-2 animate-spin" />
                            Processing...
                          </>
                        ) : (
                          <>
                            <Wallet size={16} className="mr-2" />
                            Approve & Transfer Tokens
                          </>
                        )}
                      </Button>
                    </div>
                  </div>
                </div>
              ))}
            </div>
          )}

          {/* Property Upload Requests Section */}
          <div className="space-y-4">
            <h2 className="text-xl font-semibold flex items-center gap-2">
              <FileText size={20} />
              Property Upload Requests ({requests.length})
            </h2>
            {requests.length === 0 && pendingPurchases.length === 0 ? (
              <div className="bg-[#1A1A1A] border border-[#262626] rounded-xl p-12 text-center">
                <FileText size={48} className="mx-auto text-gray-600 mb-4" />
                <p className="text-gray-400 text-lg mb-2">No messages yet</p>
                <p className="text-gray-500 text-sm">
                  Submit a property upload request or receive token purchase requests to see messages here.
                </p>
              </div>
            ) : requests.length === 0 ? (
              <div className="bg-[#1A1A1A] border border-[#262626] rounded-xl p-8 text-center">
                <p className="text-gray-400">No property upload requests</p>
              </div>
            ) : (
              requests.map((request) => (
                <div
                  key={request.id}
                  className="bg-[#1A1A1A] border border-[#262626] rounded-xl p-6 space-y-4"
                >
                  <div className="flex items-start justify-between">
                    <div className="flex-1">
                      <div className="flex items-center gap-3 mb-3">
                        <h3 className="text-xl font-semibold">{request.name}</h3>
                        {getStatusBadge(request.status)}
                      </div>
                      
                      <div className="grid md:grid-cols-2 gap-4 mb-4 text-sm">
                        <div>
                          <span className="text-gray-400">Symbol:</span>
                          <p className="text-gray-300 mt-1">{request.symbol}</p>
                        </div>
                        <div>
                          <span className="text-gray-400">Valuation:</span>
                          <p className="text-gray-300 mt-1">{request.valuation} ETH</p>
                        </div>
                        <div>
                          <span className="text-gray-400">Token Supply:</span>
                          <p className="text-gray-300 mt-1">{request.token_supply.toLocaleString()}</p>
                        </div>
                        <div>
                          <span className="text-gray-400">Submitted:</span>
                          <p className="text-gray-300 mt-1">{formatDate(request.created_at)}</p>
                        </div>
                      </div>

                      {/* Status Message */}
                      <div className={`p-4 rounded-lg border ${
                        request.status === 'pending' 
                          ? 'bg-yellow-500/10 border-yellow-500/20'
                          : request.status === 'approved'
                          ? 'bg-green-500/10 border-green-500/20'
                          : 'bg-red-500/10 border-red-500/20'
                      }`}>
                        <p className={`text-sm ${
                          request.status === 'pending'
                            ? 'text-yellow-400'
                            : request.status === 'approved'
                            ? 'text-green-400'
                            : 'text-red-400'
                        }`}>
                          {getStatusMessage(request)}
                        </p>
                      </div>
                    </div>
                  </div>
                </div>
              ))
            )}
          </div>
        </div>
      </main>
    </div>
  );
}

export default Messages;


import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { Upload, X, FileText, CheckCircle } from "lucide-react";
import { Button } from "../Components/atoms/Button";
import { api } from "../lib/api";
import { checkApprovalStatus } from "../lib/contracts";
import { useWallet } from "../hooks/useWallet";
import { getAdminWalletState } from "../lib/wallet";
import { ethers } from "ethers";
import type { User } from "../types";

function AdminCreateProperty() {
  const navigate = useNavigate();
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");
  const [success, setSuccess] = useState(false);
  const [txHash, setTxHash] = useState("");

  // Form state
  const [name, setName] = useState("");
  const [valuation, setValuation] = useState("");
  const [ownerAddress, setOwnerAddress] = useState("");
  const [files, setFiles] = useState<File[]>([]);
  const [uploading, setUploading] = useState(false);

  // Approved wallets state
  const [approvedWallets, setApprovedWallets] = useState<User[]>([]);
  const [walletsLoading, setWalletsLoading] = useState(true);
  const [isCheckingBlockchain, setIsCheckingBlockchain] = useState(false);
  const { isConnected, provider } = useWallet();

  // Check approval status for users - uses blockchain first, falls back to database
  const checkBlockchainApproval = async (users: User[]) => {
    console.log('🔄 Starting approval check for create property (v2.0 - with DB fallback)...');
    setIsCheckingBlockchain(true);

    let blockchainProvider: ethers.BrowserProvider | ethers.AbstractProvider | null = provider;
    let useBlockchain = true;

    // If MetaMask is not connected, use admin wallet for blockchain operations
    if (!isConnected || !provider) {
      console.log('ℹ️ MetaMask not connected, using admin wallet for blockchain checks...');
      try {
        const adminWallet = await getAdminWalletState();
        blockchainProvider = adminWallet.provider;
        console.log('✅ Admin wallet connected for blockchain operations');
      } catch (error) {
        console.error('❌ Failed to connect admin wallet, will use database fallback:', error);
        useBlockchain = false;
      }
    }

    if (users.length === 0) {
      console.log('ℹ️ No users to check');
      setIsCheckingBlockchain(false);
      return users;
    }

    try {
      const approvedUsers: User[] = [];

      for (const user of users) {
        // Skip users with empty wallet addresses
        if (!user.WalletAddress || user.WalletAddress.trim() === "") {
          console.log(`⏭️ Skipping ${user.Email}: No wallet address set`);
          continue;
        }

        let isApproved = false;

        // First, try blockchain check if available
        if (useBlockchain && blockchainProvider) {
          try {
            console.log(`🔍 Checking blockchain approval for ${user.Email} (${user.WalletAddress})`);

            // Add timeout to blockchain check
            const checkPromise = checkApprovalStatus(user.WalletAddress, blockchainProvider);
            const timeoutPromise = new Promise<boolean>((_, reject) =>
              setTimeout(() => reject(new Error('Blockchain check timeout')), 3000)
            );

            isApproved = await Promise.race([checkPromise, timeoutPromise]);
            console.log(`📡 Blockchain approval status for ${user.Email}: ${isApproved}`);

          } catch (err) {
            console.warn(`⚠️ Blockchain check failed for ${user.Email}, using database fallback:`, err);
            isApproved = false; // Will fall through to database check
          }
        }

        // If blockchain check didn't confirm approval, check database as fallback
        if (!isApproved) {
          // Check database approval_status field
          const dbApproved = user.approval_status === 'approved';
          console.log(`📋 Database approval status for ${user.Email}: ${user.approval_status} -> ${dbApproved}`);
          
          if (dbApproved) {
            console.log(`✅ ${user.Email} approved via database (blockchain unavailable or returned false)`);
            isApproved = true;
          }
        }

        if (isApproved) {
          approvedUsers.push(user);
          console.log(`✅ ${user.Email} added to approved wallets list`);
        } else {
          console.log(`❌ ${user.Email} not approved (neither blockchain nor database)`);
        }
      }

      console.log(`✅ Found ${approvedUsers.length} approved wallets (blockchain + DB fallback)`);
      return approvedUsers;

    } catch (err) {
      console.error('❌ Approval check failed:', err);
      
      // Last resort fallback: return users approved in database
      console.log('🔄 Using database-only fallback for approved wallets...');
      const dbApprovedUsers = users.filter(user => user.approval_status === 'approved');
      console.log(`📋 Found ${dbApprovedUsers.length} database-approved wallets as fallback`);
      return dbApprovedUsers;
    } finally {
      setIsCheckingBlockchain(false);
    }
  };

  // Fetch approved wallets
  useEffect(() => {
    const fetchApprovedWallets = async () => {
      try {
        // First get all users from database
        const users: User[] = await api.getUsers();

        // Filter to users with wallet addresses (exclude admins and users without wallets)
        const usersWithWallets = users.filter(user =>
          user.WalletAddress &&
          user.WalletAddress.trim() !== "" &&
          user.Role !== 'admin'
        );

        console.log(`📋 Found ${usersWithWallets.length} users with wallet addresses`);

        // Check blockchain approval status for these users
        const blockchainApprovedUsers = await checkBlockchainApproval(usersWithWallets);

        setApprovedWallets(blockchainApprovedUsers);
      } catch (err) {
        console.error("Failed to fetch approved wallets:", err);
        setApprovedWallets([]);
      } finally {
        setWalletsLoading(false);
      }
    };

    fetchApprovedWallets();
  }, [isConnected, provider]);

  // Computed values
  const symbol = "ETH";
  const tokenSupply = valuation ? Math.floor(parseFloat(valuation) / 10) : 0;


  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const selectedFiles = Array.from(e.target.files || []);
    if (selectedFiles.length === 0) return;

    setFiles(selectedFiles);
    setError("");
  };

  const removeFile = (index: number) => {
    setFiles(files.filter((_, i) => i !== index));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError("");
    setLoading(true);

    try {
      if (!ownerAddress) {
        throw new Error("Please enter an owner wallet address");
      }

      if (files.length === 0) {
        throw new Error("Please upload at least one document");
      }

      if (tokenSupply < 1) {
        throw new Error("Valuation must be at least 10 ETH to create tokens");
      }

      setUploading(true);
      const response = await api.createProperty(
        {
          owner_address: ownerAddress,
          name,
          symbol,
          data_hash: "", // Not needed - files are sent directly
          valuation: parseFloat(valuation),
          token_supply: tokenSupply,
        },
        files
      );

      setTxHash(response.tx_hash);
      setSuccess(true);
    } catch (err: any) {
      setError(err.message || "Failed to create property");
    } finally {
      setLoading(false);
      setUploading(false);
    }
  };

  if (success) {
    return (
      <div className="max-w-xl mx-auto space-y-6">
        <div className="rounded-3xl border border-[#1f1f1f] bg-[#111111] p-8 text-center space-y-4">
          <div className="flex justify-center">
            <CheckCircle className="text-green-500" size={64} />
          </div>
          <h2 className="text-2xl font-semibold">Property Created!</h2>
          <p className="text-gray-400">
            The property has been successfully tokenized on the blockchain.
          </p>
          {txHash && (
            <div className="bg-[#0f0f0f] rounded-xl p-4 text-left">
              <p className="text-xs text-gray-500 mb-1">Transaction Hash</p>
              <p className="text-sm font-mono text-gray-300 break-all">
                {txHash}
              </p>
            </div>
          )}
          <div className="flex gap-3 justify-center pt-4">
            <Button
              variant="outline"
              size="md"
              className="border-[#262626] text-white hover:bg-[#1a1a1a]"
              onClick={() => navigate("/admin")}
            >
              Back to Dashboard
            </Button>
            <Button
              variant="outline"
              size="md"
              className="border-[#6d41ff] text-[#6d41ff] hover:bg-[#6d41ff]/10"
              onClick={() => navigate("/admin/properties")}
            >
              View Properties
            </Button>
            <Button
              variant="primary"
              size="md"
              className="bg-[#6d41ff] hover:bg-[#5b2fff] text-white"
              onClick={() => {
                setSuccess(false);
                setName("");
                setValuation("");
                setOwnerAddress("");
                setFiles([]);
                setTxHash("");
              }}
            >
              Create Another
            </Button>
          </div>
        </div>
      </div>
    );
  }

  return (
    <div className="max-w-2xl mx-auto space-y-6">
      <header>
        <p className="text-sm text-gray-400">Admin</p>
        <h1 className="text-3xl font-semibold">Create Property</h1>
      </header>

      {error && (
        <div className="text-red-500 bg-red-500/10 p-4 rounded-xl border border-red-500/20">
          {error}
        </div>
      )}

      <form onSubmit={handleSubmit} className="space-y-6">
        {/* Owner Address - Dropdown of approved wallets (UI only for now) */}
        <div className="rounded-3xl border border-[#1f1f1f] bg-[#111111] p-6 space-y-4">
          <h2 className="font-semibold text-lg">Owner Information</h2>
          <div className="space-y-2">
            <label className="text-sm text-gray-400">Owner Wallet Address</label>
            <select
              value={ownerAddress}
              onChange={(e) => setOwnerAddress(e.target.value)}
              required
              className="w-full px-4 py-3 rounded-xl border border-[#1f1f1f] bg-[#0f0f0f] text-white focus:outline-none focus:border-[#6d41ff] cursor-pointer"
            >
              <option value="">
                {walletsLoading || isCheckingBlockchain ? "Loading approved wallets..." : "Select an approved wallet..."}
              </option>
              {approvedWallets.map((user) => (
                <option key={user.WalletAddress} value={user.WalletAddress}>
                  {user.Name} ({user.WalletAddress.substring(0, 6)}...
                  {user.WalletAddress.substring(38)})
                </option>
              ))}
            </select>
            <p className="text-xs text-gray-500">
              Only approved wallets can own properties
            </p>
          </div>
        </div>

        {/* Property Details */}
        <div className="rounded-3xl border border-[#1f1f1f] bg-[#111111] p-6 space-y-4">
          <h2 className="font-semibold text-lg">Property Details</h2>

          <div className="space-y-2">
            <label className="text-sm text-gray-400">Property Name</label>
            <input
              type="text"
              value={name}
              onChange={(e) => setName(e.target.value)}
              placeholder="e.g., Luxury Villa Miami"
              required
              className="w-full px-4 py-3 rounded-xl border border-[#1f1f1f] bg-[#0f0f0f] text-white placeholder-gray-500 focus:outline-none focus:border-[#6d41ff]"
            />
          </div>

          <div className="space-y-2">
            <label className="text-sm text-gray-400">Valuation (ETH)</label>
            <input
              type="number"
              value={valuation}
              onChange={(e) => setValuation(e.target.value)}
              placeholder="e.g., 100"
              step="0.01"
              min="0"
              required
              className="w-full px-4 py-3 rounded-xl border border-[#1f1f1f] bg-[#0f0f0f] text-white placeholder-gray-500 focus:outline-none focus:border-[#6d41ff]"
            />
          </div>

          <div className="grid md:grid-cols-2 gap-4">
            <div className="space-y-2">
              <label className="text-sm text-gray-400">Token Symbol</label>
              <div className="w-full px-4 py-3 rounded-xl border border-[#1f1f1f] bg-[#1a1a1a] text-gray-300">
                {symbol}
              </div>
              <p className="text-xs text-gray-500">Fixed token symbol</p>
            </div>

            <div className="space-y-2">
              <label className="text-sm text-gray-400">Token Supply</label>
              <div className="w-full px-4 py-3 rounded-xl border border-[#1f1f1f] bg-[#1a1a1a] text-gray-300">
                {tokenSupply || "—"}
              </div>
              <p className="text-xs text-gray-500">Calculated as valuation ÷ 10</p>
            </div>
          </div>
        </div>

        {/* Document Upload */}
        <div className="rounded-3xl border border-[#1f1f1f] bg-[#111111] p-6 space-y-4">
          <h2 className="font-semibold text-lg">Property Documents</h2>
          <p className="text-sm text-gray-400">
            Upload property documentation (deed, certificate, images, etc.). Multiple files are supported.
          </p>

          <label className="flex flex-col items-center justify-center h-40 border-2 border-dashed border-[#1f1f1f] rounded-2xl cursor-pointer hover:border-[#6d41ff] transition-colors">
            <Upload className="text-gray-400 mb-2" size={32} />
            <span className="text-gray-400">Click to upload files</span>
            <span className="text-xs text-gray-500 mt-1">Multiple files supported</span>
            <input
              type="file"
              onChange={handleFileChange}
              multiple
              className="hidden"
            />
          </label>

          {files.length > 0 && (
            <div className="space-y-2">
              <p className="text-sm text-gray-400">Selected files ({files.length}):</p>
              {files.map((file, index) => (
                <div key={index} className="flex items-center justify-between p-3 rounded-xl border border-[#1f1f1f] bg-[#0f0f0f]">
                  <div className="flex items-center gap-3">
                    <FileText className="text-[#6d41ff]" size={20} />
                    <div>
                      <p className="text-sm font-medium truncate max-w-[200px]">
                        {file.name}
                      </p>
                      <p className="text-xs text-gray-500">
                        {(file.size / 1024).toFixed(2)} KB
                      </p>
                    </div>
                  </div>
                  <button
                    type="button"
                    onClick={() => removeFile(index)}
                    className="p-2 rounded-lg hover:bg-[#1a1a1a] text-gray-400 hover:text-white transition-colors"
                  >
                    <X size={18} />
                  </button>
                </div>
              ))}
            </div>
          )}
        </div>

        {/* Submit */}
        <div className="flex gap-3">
          <Button
            type="button"
            variant="outline"
            size="md"
            className="border-[#262626] text-white hover:bg-[#1a1a1a]"
            onClick={() => navigate("/admin")}
          >
            Cancel
          </Button>
          <Button
            type="submit"
            variant="primary"
            size="md"
            className="bg-[#6d41ff] hover:bg-[#5b2fff] text-white flex-1"
            disabled={loading || uploading || files.length === 0}
          >
            {loading ? "Creating Property..." : "Create Property"}
          </Button>
        </div>
      </form>
    </div>
  );
}

export default AdminCreateProperty;



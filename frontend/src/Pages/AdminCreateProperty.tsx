import { useState, useEffect, useRef } from "react";
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
  const [file, setFile] = useState<File | null>(null);
  const [uploadedHash, setUploadedHash] = useState("");
  const [uploading, setUploading] = useState(false);

  // Approved wallets state
  const [approvedWallets, setApprovedWallets] = useState<User[]>([]);
  const [walletsLoading, setWalletsLoading] = useState(true);
  const [isCheckingBlockchain, setIsCheckingBlockchain] = useState(false);
  const { isConnected, provider } = useWallet();

  // Check blockchain approval status for users
  const checkBlockchainApproval = async (users: User[]) => {
    console.log('🔄 Starting blockchain approval check for create property...');
    setIsCheckingBlockchain(true);

    let blockchainProvider: ethers.BrowserProvider | ethers.AbstractProvider | null = provider;

    // If MetaMask is not connected, use admin wallet for blockchain operations
    if (!isConnected || !provider) {
      console.log('ℹ️ MetaMask not connected, using admin wallet for blockchain checks...');
      try {
        const adminWallet = await getAdminWalletState();
        blockchainProvider = adminWallet.provider;
        console.log('✅ Admin wallet connected for blockchain operations');
      } catch (error) {
        console.error('❌ Failed to connect admin wallet:', error);
        setIsCheckingBlockchain(false);
        return users;
      }
    }

    if (users.length === 0) {
      console.log('ℹ️ No users to check');
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

        try {
          console.log(`🔍 Checking approval status for ${user.Email} (${user.WalletAddress})`);

          // Add timeout to blockchain check
          const checkPromise = checkApprovalStatus(user.WalletAddress, blockchainProvider!);
          const timeoutPromise = new Promise<boolean>((_, reject) =>
            setTimeout(() => reject(new Error('Blockchain check timeout')), 3000)
          );

          const isApproved = await Promise.race([checkPromise, timeoutPromise]);
          console.log(`✅ Approval status for ${user.Email}: ${isApproved}`);

          if (isApproved) {
            approvedUsers.push(user);
          }

        } catch (err) {
          console.error(`❌ Failed to check approval for ${user.Email}:`, err);
          // Skip users we can't check - only include confirmed approved users
        }
      }

      console.log(`✅ Found ${approvedUsers.length} blockchain-approved wallets`);
      return approvedUsers;

    } catch (err) {
      console.error('❌ Blockchain approval check failed:', err);
      // Return empty array if blockchain check fails completely
      return [];
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


  const handleFileChange = async (e: React.ChangeEvent<HTMLInputElement>) => {
    const selectedFile = e.target.files?.[0];
    if (!selectedFile) return;

    setFile(selectedFile);
    setUploading(true);
    setError("");

    try {
      const response = await api.uploadFile(selectedFile);
      setUploadedHash(response.ipfs_hash);
    } catch (err: any) {
      setError("Failed to upload file: " + err.message);
      setFile(null);
    } finally {
      setUploading(false);
    }
  };

  const removeFile = () => {
    setFile(null);
    setUploadedHash("");
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError("");
    setLoading(true);

    try {
      if (!ownerAddress) {
        throw new Error("Please enter an owner wallet address");
      }

      if (!uploadedHash) {
        throw new Error("Please upload a document first");
      }

      if (tokenSupply < 1) {
        throw new Error("Valuation must be at least 10 ETH to create tokens");
      }

      const response = await api.createProperty({
        owner_address: ownerAddress,
        name,
        symbol,
        data_hash: uploadedHash,
        valuation: parseFloat(valuation),
        token_supply: tokenSupply,
      });

      setTxHash(response.tx_hash);
      setSuccess(true);
    } catch (err: any) {
      setError(err.message || "Failed to create property");
    } finally {
      setLoading(false);
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
              variant="primary"
              size="md"
              className="bg-[#6d41ff] hover:bg-[#5b2fff] text-white"
              onClick={() => {
                setSuccess(false);
                setName("");
                setValuation("");
                setOwnerAddress("");
                setFile(null);
                setUploadedHash("");
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
                {walletsLoading || isCheckingBlockchain ? "Loading blockchain-approved wallets..." : "Select a blockchain-approved wallet..."}
              </option>
              {approvedWallets.map((user) => (
                <option key={user.WalletAddress} value={user.WalletAddress}>
                  {user.Name} ({user.WalletAddress.substring(0, 6)}...
                  {user.WalletAddress.substring(38)})
                </option>
              ))}
            </select>
            <p className="text-xs text-gray-500">
              Only blockchain-approved wallets can own properties
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
          <h2 className="font-semibold text-lg">Property Document</h2>
          <p className="text-sm text-gray-400">
            Upload property documentation (deed, certificate, etc.)
          </p>

          {!file ? (
            <label className="flex flex-col items-center justify-center h-40 border-2 border-dashed border-[#1f1f1f] rounded-2xl cursor-pointer hover:border-[#6d41ff] transition-colors">
              <Upload className="text-gray-400 mb-2" size={32} />
              <span className="text-gray-400">Click to upload</span>
              <span className="text-xs text-gray-500 mt-1">PNG files only</span>
              <input
                type="file"
                onChange={handleFileChange}
                accept=".png"
                className="hidden"
              />
            </label>
          ) : (
            <div className="flex items-center justify-between p-4 rounded-xl border border-[#1f1f1f] bg-[#0f0f0f]">
              <div className="flex items-center gap-3">
                <FileText className="text-[#6d41ff]" size={24} />
                <div>
                  <p className="text-sm font-medium truncate max-w-[200px]">
                    {file.name}
                  </p>
                  <p className="text-xs text-gray-500">
                    {uploading
                      ? "Uploading to IPFS..."
                      : uploadedHash
                      ? "Uploaded successfully"
                      : "Ready"}
                  </p>
                </div>
              </div>
              <button
                type="button"
                onClick={removeFile}
                className="p-2 rounded-lg hover:bg-[#1a1a1a] text-gray-400 hover:text-white transition-colors"
              >
                <X size={20} />
              </button>
            </div>
          )}

          {uploadedHash && (
            <div className="bg-[#0f0f0f] rounded-xl p-3">
              <p className="text-xs text-gray-500 mb-1">IPFS Hash</p>
              <p className="text-xs font-mono text-gray-400 break-all">
                {uploadedHash}
              </p>
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
            disabled={loading || uploading || !uploadedHash}
          >
            {loading ? "Creating Property..." : "Create Property"}
          </Button>
        </div>
      </form>
    </div>
  );
}

export default AdminCreateProperty;



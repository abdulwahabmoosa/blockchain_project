import { useEffect, useState } from "react";
import { RefreshCw, Search, CheckCircle, XCircle, AlertTriangle, Building2 } from "lucide-react";
import { Button } from "../Components/atoms/Button";
import { api } from "../lib/api";
import { useWallet } from "../hooks/useWallet";
import { WETH_ADDRESS, checkWETHBalance, wrapETH, approveWETH, checkWETHAllowance } from "../lib/contracts";
import { getContractAddresses } from "../lib/wallet";
import { ethers } from "ethers";
import type { Property } from "../types";

interface PropertyWithStatus extends Property {
  approving?: boolean;
  rejecting?: boolean;
  distributingRent?: boolean;
}

const CONVERSION_FACTOR = 1000000; // 1 SepoliaETH = 1,000,000 system ETH

function AdminManageProperties() {
  const [properties, setProperties] = useState<PropertyWithStatus[]>([]);
  const [filteredProperties, setFilteredProperties] = useState<PropertyWithStatus[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");
  const [filter, setFilter] = useState<"all" | "active" | "closed">("all");
  const [searchQuery, setSearchQuery] = useState("");
  const { signer, provider, address, connect } = useWallet();

  useEffect(() => {
    loadProperties();
  }, []);

  // Filter properties based on status
  useEffect(() => {
    let filtered = properties;

    switch (filter) {
      case 'active':
        filtered = properties.filter(property => property.Status === "Active");
        break;
      case 'closed':
        filtered = properties.filter(property => property.Status === "Closed");
        break;
      case 'all':
      default:
        filtered = properties;
        break;
    }

    if (searchQuery) {
      filtered = filtered.filter((property) =>
        (property.Name?.toLowerCase().includes(searchQuery.toLowerCase()) ?? false) ||
        property.OwnerWallet.toLowerCase().includes(searchQuery.toLowerCase())
      );
    }

    setFilteredProperties(filtered);
  }, [properties, filter, searchQuery]);

  const loadProperties = async () => {
    setLoading(true);
    setError("");

    try {
      const data = await api.getProperties();
      const propertiesWithStatus = data.map(prop => ({
        ...prop,
        approving: false,
        rejecting: false,
        distributingRent: false
      }));
      setProperties(propertiesWithStatus);
    } catch (err: any) {
      console.error("Failed to load properties:", err);
      setError(err.message || "Failed to load properties");
    } finally {
      setLoading(false);
    }
  };

  const handleApproveProperty = async (property: PropertyWithStatus) => {
    setProperties(prev => prev.map(p =>
      p.ID === property.ID ? { ...p, approving: true } : p
    ));

    try {
      console.log(`🔄 Calling backend API to approve property ${property.ID}`);

      const result = await api.updatePropertyApproval(property.ID, "approved");
      console.log('✅ Backend API response:', result);

      // Update property status
      setProperties(prev => prev.map(p =>
        p.ID === property.ID ? { ...p, Status: "Active", approving: false } : p
      ));

      alert(`✅ Property approval successful!${result.tx_hash ? ` Transaction: ${result.tx_hash}` : ''}`);

    } catch (err: any) {
      console.error("❌ Approval failed:", err);
      setError(err.message || "Failed to approve property");

      setProperties(prev => prev.map(p =>
        p.ID === property.ID ? { ...p, approving: false } : p
      ));

      alert(`❌ Approval failed: ${err.message}`);
    }
  };

  const handleRejectProperty = async (property: PropertyWithStatus) => {
    setProperties(prev => prev.map(p =>
      p.ID === property.ID ? { ...p, rejecting: true } : p
    ));

    try {
      console.log(`🔄 Calling backend API to reject property ${property.ID}`);

      const result = await api.updatePropertyApproval(property.ID, "rejected");
      console.log('✅ Backend API response:', result);

      // Update property status
      setProperties(prev => prev.map(p =>
        p.ID === property.ID ? { ...p, Status: "Closed", rejecting: false } : p
      ));

      alert(`✅ Property rejection successful!${result.tx_hash ? ` Transaction: ${result.tx_hash}` : ''}`);

    } catch (err: any) {
      console.error("❌ Rejection failed:", err);
      setError(err.message || "Failed to reject property");

      setProperties(prev => prev.map(p =>
        p.ID === property.ID ? { ...p, rejecting: false } : p
      ));

      alert(`❌ Rejection failed: ${err.message}`);
    }
  };

  const handleDistributeRent = async (property: PropertyWithStatus) => {
    if (!signer || !provider || !address) {
      alert("Please connect your wallet first");
      try {
        await connect();
      } catch (err) {
        setError("Failed to connect wallet");
      }
      return;
    }

    // Set distributing state
    setProperties(prev => prev.map(p =>
      p.ID === property.ID ? { ...p, distributingRent: true } : p
    ));
    setError("");

    try {
      // Calculate rent: valuation / 50 (in system ETH)
      const rentSystemETH = property.Valuation / 50;
      
      // Convert to SepoliaETH: rent / 1,000,000
      const sepoliaETHAmount = rentSystemETH / CONVERSION_FACTOR;
      
      // Convert to wei (for contract interaction)
      const amountInWei = ethers.parseEther(sepoliaETHAmount.toString());
      
      console.log(`💰 Distributing rent for property ${property.ID}:`);
      console.log(`   Valuation: ${property.Valuation} ETH (system)`);
      console.log(`   Rent: ${rentSystemETH} ETH (system) = ${sepoliaETHAmount} SepoliaETH`);
      console.log(`   Amount in wei: ${amountInWei.toString()}`);

      // Get RevenueDistribution contract address
      const addresses = getContractAddresses();
      const revenueDistributionAddress = addresses.RevenueDistribution;

      // Check WETH balance
      const wethBalance = await checkWETHBalance(address, provider);
      console.log(`💵 WETH balance: ${ethers.formatEther(wethBalance)} WETH`);

      // If not enough WETH, wrap ETH
      if (wethBalance < amountInWei) {
        const ethNeeded = amountInWei - wethBalance;
        console.log(`⚠️ Not enough WETH. Need to wrap ${ethers.formatEther(ethNeeded)} ETH`);
        
        // Get ETH balance
        const ethBalance = await provider.getBalance(address);
        if (ethBalance < ethNeeded) {
          throw new Error(`Insufficient balance. Need ${ethers.formatEther(ethNeeded)} ETH but have ${ethers.formatEther(ethBalance)} ETH`);
        }

        // Wrap ETH to WETH
        console.log(`🔄 Wrapping ${ethers.formatEther(ethNeeded)} ETH to WETH...`);
        const wrapTx = await wrapETH(ethNeeded, signer);
        await wrapTx.wait();
        console.log(`✅ Wrapped ETH to WETH. TX: ${wrapTx.hash}`);
      }

      // Check and approve WETH allowance
      const currentAllowance = await checkWETHAllowance(address, revenueDistributionAddress, provider);
      console.log(`✅ Current WETH allowance: ${ethers.formatEther(currentAllowance)} WETH`);

      if (currentAllowance < amountInWei) {
        console.log(`🔄 Approving WETH spending...`);
        const approveTx = await approveWETH(revenueDistributionAddress, ethers.MaxUint256, signer);
        await approveTx.wait();
        console.log(`✅ WETH approved. TX: ${approveTx.hash}`);
      }

      // Call backend endpoint to distribute revenue
      console.log(`🚀 Calling backend to distribute revenue...`);
      // Convert BigInt to number (safe for SepoliaETH amounts which are small)
      const amountNumber = Number(amountInWei);
      const result = await api.distributeRevenue(
        property.OnchainTokenAddress,
        WETH_ADDRESS,
        amountNumber
      );

      console.log(`✅ Revenue distribution submitted. TX: ${result.tx_hash}`);
      alert(`✅ Rent distribution successful!\n\nTransaction: ${result.tx_hash}\n\nA snapshot has been taken and token holders can now claim their proportional share.`);

      // Refresh properties
      await loadProperties();

    } catch (err: any) {
      console.error("❌ Failed to distribute rent:", err);
      setError(err.message || "Failed to distribute rent");
      alert(`❌ Failed to distribute rent: ${err.message || "Unknown error"}`);
    } finally {
      setProperties(prev => prev.map(p =>
        p.ID === property.ID ? { ...p, distributingRent: false } : p
      ));
    }
  };

  const getStatusBadge = (property: PropertyWithStatus) => {
    if (property.approving) return (
      <span className="flex items-center gap-1 text-xs px-2 py-1 rounded-full bg-blue-500/20 text-blue-400">
        <RefreshCw size={12} className="animate-spin" />
        Approving...
      </span>
    );
    if (property.rejecting) return (
      <span className="flex items-center gap-1 text-xs px-2 py-1 rounded-full bg-red-500/20 text-red-400">
        <RefreshCw size={12} className="animate-spin" />
        Rejecting...
      </span>
    );
    if (property.Status === "Active") return (
      <span className="flex items-center gap-1 text-xs px-2 py-1 rounded-full bg-green-500/20 text-green-400">
        <CheckCircle size={12} />
        Active
      </span>
    );
    if (property.Status === "Closed") return (
      <span className="flex items-center gap-1 text-xs px-2 py-1 rounded-full bg-red-500/20 text-red-400">
        <XCircle size={12} />
        Closed
      </span>
    );
    return (
      <span className="flex items-center gap-1 text-xs px-2 py-1 rounded-full bg-gray-500/20 text-gray-400">
        {property.Status}
      </span>
    );
  };

  return (
    <div className="space-y-6">
      <header className="flex flex-col md:flex-row md:items-center justify-between gap-4">
        <div>
          <h1 className="text-3xl font-semibold">Manage Properties</h1>
          <p className="text-sm text-gray-400">
            Approve or reject property listings
          </p>
        </div>
        <Button
          variant="outline"
          size="md"
          onClick={loadProperties}
          disabled={loading}
          className="border-[#262626] text-white hover:bg-[#1a1a1a]"
        >
          <RefreshCw
            size={18}
            className={`mr-2 ${loading ? "animate-spin" : ""}`}
          />
          Refresh
        </Button>
      </header>

      {error && (
        <div className="text-red-500 bg-red-500/10 p-4 rounded-xl border border-red-500/20 flex items-center gap-2">
          <AlertTriangle size={18} />
          {error}
        </div>
      )}

      {!signer && (
        <div className="text-yellow-500 bg-yellow-500/10 p-4 rounded-xl border border-yellow-500/20 flex items-center gap-2">
          <AlertTriangle size={18} />
          <span>Please connect your wallet to distribute rent. Click "Distribute Rent" to connect automatically.</span>
        </div>
      )}

      {/* Search and Filter */}
      <div className="flex flex-col md:flex-row gap-4">
        <div className="relative flex-1">
          <Search
            className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400"
            size={20}
          />
          <input
            type="text"
            placeholder="Search by property name or owner..."
            value={searchQuery}
            onChange={(e) => setSearchQuery(e.target.value)}
            className="w-full pl-10 pr-4 py-3 rounded-xl border border-[#1f1f1f] bg-[#111111] text-white placeholder-gray-500 focus:outline-none focus:border-[#6d41ff]"
          />
        </div>
        <div className="flex gap-2">
          <button
            onClick={() => setFilter("all")}
            className={`px-4 py-2 rounded-xl text-sm font-medium transition-colors ${
              filter === "all"
                ? "bg-[#703BF7] text-white"
                : "bg-[#111111] border border-[#1f1f1f] text-gray-400 hover:text-white"
            }`}
          >
            All ({properties.length})
          </button>
          <button
            onClick={() => setFilter("active")}
            className={`px-4 py-2 rounded-xl text-sm font-medium transition-colors ${
              filter === "active"
                ? "bg-green-600 text-white"
                : "bg-[#111111] border border-[#1f1f1f] text-gray-400 hover:text-white"
            }`}
          >
            Active ({properties.filter(p => p.Status === "Active").length})
          </button>
          <button
            onClick={() => setFilter("closed")}
            className={`px-4 py-2 rounded-xl text-sm font-medium transition-colors ${
              filter === "closed"
                ? "bg-red-600 text-white"
                : "bg-[#111111] border border-[#1f1f1f] text-gray-400 hover:text-white"
            }`}
          >
            Closed ({properties.filter(p => p.Status === "Closed").length})
          </button>
        </div>
      </div>

      {/* Properties List */}
      <div className="rounded-3xl border border-[#1f1f1f] bg-[#111111] overflow-hidden">
        <div className="p-4 border-b border-[#1f1f1f] flex items-center justify-between">
          <h2 className="font-semibold">Properties</h2>
          <span className="text-xs text-gray-500">
            {filteredProperties.length} of {properties.length} properties
          </span>
        </div>

        {loading ? (
          <div className="p-8 text-center">
            <RefreshCw
              size={24}
              className="animate-spin mx-auto text-gray-400 mb-2"
            />
            <p className="text-gray-400">Loading properties...</p>
          </div>
        ) : filteredProperties.length === 0 ? (
          <div className="p-8 text-center text-gray-400">
            {searchQuery || filter !== "all"
              ? "No properties match your search criteria"
              : "No properties found"}
          </div>
        ) : (
          <div className="divide-y divide-[#1f1f1f]">
            {filteredProperties.map((property) => (
              <div
                key={property.ID}
                className="p-4 hover:bg-[#0f0f0f]/50 transition-colors"
              >
                <div className="flex flex-col md:flex-row md:items-center justify-between gap-4">
                  <div className="flex-1 min-w-0">
                    <div className="flex items-center gap-2 mb-1">
                      <h3 className="font-semibold truncate">{property.Name}</h3>
                      <span className="text-xs px-2 py-0.5 rounded-full bg-blue-500/20 text-blue-400">
                        ID: {property.ID.substring(0, 8)}
                      </span>
                    </div>
                    <p className="text-sm text-gray-400">
                      Valuation: {property.Valuation} ETH
                    </p>
                    <p className="text-xs text-gray-500 font-mono mt-1">
                      Owner: {property.OwnerWallet.substring(0, 10)}...
                      {property.OwnerWallet.substring(38)}
                    </p>
                  </div>

                  <div className="flex items-center gap-3">
                    {getStatusBadge(property)}

                    {property.Status === "Active" ? (
                      <div className="flex gap-2 items-center">
                        <span className="text-xs text-green-400">Active</span>
                        <Button
                          variant="primary"
                          size="sm"
                          onClick={() => handleDistributeRent(property)}
                          disabled={property.distributingRent}
                          className="hidden bg-green-600 hover:bg-green-700 text-white disabled:opacity-50"
                          title={!signer ? "Connect your wallet to distribute rent" : ""}
                        >
                          {property.distributingRent ? "Distributing..." : "Distribute Rent"}
                        </Button>
                      </div>
                    ) : property.Status === "Closed" ? (
                      <span className="text-xs text-red-400">Already Closed</span>
                    ) : (
                      <div className="flex gap-2">
                        <Button
                          variant="primary"
                          size="sm"
                          onClick={() => handleApproveProperty(property)}
                          disabled={property.approving || property.rejecting}
                          className="bg-[#6d41ff] hover:bg-[#5b2fff] text-white disabled:opacity-50"
                        >
                          {property.approving ? "Approving..." : "Approve"}
                        </Button>
                        <Button
                          variant="outline"
                          size="sm"
                          onClick={() => handleRejectProperty(property)}
                          disabled={property.approving || property.rejecting}
                          className="border-red-500/50 bg-red-500/10 text-red-400 hover:bg-red-500/20 disabled:opacity-50"
                        >
                          {property.rejecting ? "Rejecting..." : "Reject"}
                        </Button>
                      </div>
                    )}
                  </div>
                </div>
              </div>
            ))}
          </div>
        )}
      </div>

      {/* Info Section */}
      <div className="rounded-2xl border border-blue-500/20 bg-blue-500/5 p-4">
        <h3 className="text-blue-400 font-medium mb-2 flex items-center gap-2">
          <Building2 size={16} />
          Property Management
        </h3>
        <ul className="text-sm text-gray-400 space-y-1">
          <li>Active properties are visible to all users</li>
          <li>Closed properties are hidden from public listings</li>
          <li>Property approval/rejection affects token trading permissions</li>
        </ul>
      </div>
    </div>
  );
}

export default AdminManageProperties;

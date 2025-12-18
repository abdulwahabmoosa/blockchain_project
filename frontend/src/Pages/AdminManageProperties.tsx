import { useEffect, useState } from "react";
import { RefreshCw, Search, CheckCircle, XCircle, AlertTriangle, Building2 } from "lucide-react";
import { Button } from "../Components/atoms/Button";
import { api } from "../lib/api";
import type { Property } from "../types";

interface PropertyWithStatus extends Property {
  approving?: boolean;
  rejecting?: boolean;
}

function AdminManageProperties() {
  const [properties, setProperties] = useState<PropertyWithStatus[]>([]);
  const [filteredProperties, setFilteredProperties] = useState<PropertyWithStatus[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");
  const [filter, setFilter] = useState<"all" | "active" | "closed">("all");
  const [searchQuery, setSearchQuery] = useState("");

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
        property.Name.toLowerCase().includes(searchQuery.toLowerCase()) ||
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
        rejecting: false
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
                      <span className="text-xs text-green-400">Already Active</span>
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

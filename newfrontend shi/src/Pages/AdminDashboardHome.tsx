import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { Building2, Coins, Users, Shield } from "lucide-react";
import { Button } from "../Components/atoms/Button";
import { api } from "../lib/api";
import type { Property, User } from "../types";

const StatCard = ({
  title,
  value,
  icon: Icon,
  accent = "bg-[#6d41ff]",
}: {
  title: string;
  value: string;
  icon: React.ElementType;
  accent?: string;
}) => (
  <div className="rounded-2xl border border-[#1f1f1f] bg-[#111111] p-6 space-y-3">
    <div className="flex items-center justify-between">
      <p className="text-sm text-gray-400">{title}</p>
      <div className={`h-10 w-10 rounded-full ${accent} flex items-center justify-center`}>
        <Icon size={20} />
      </div>
    </div>
    <p className="text-3xl font-semibold">{value}</p>
  </div>
);

function AdminDashboardHome() {
  const [properties, setProperties] = useState<Property[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");
  const [totalUsers] = useState<number>(0);
  const [user, setUser] = useState<User | null>(null);
  const navigate = useNavigate();

  useEffect(() => {
    const userStr = localStorage.getItem("user");
    if (userStr) {
      const userData = JSON.parse(userStr);
      setUser(userData);
    }
  }, []);

  useEffect(() => {
    const fetchProperties = async () => {
      try {
        const data = await api.getProperties();
        setProperties(data);
      } catch (err: any) {
        console.error("Failed to fetch properties:", err);
        setError("Failed to load properties");
      } finally {
        setLoading(false);
      }
    };
    fetchProperties();
  }, []);

  // Calculate total token supply across all properties
  const totalTokens = properties.reduce((sum, p) => sum + (p.Valuation || 0), 0);

  return (
    <div className="space-y-6">
      <header className="flex flex-col md:flex-row md:items-center justify-between gap-4">
        <div>
          <p className="text-sm text-gray-400">Admin Panel</p>
          <h1 className="text-3xl font-semibold">Dashboard</h1>
        </div>
        <Button
          variant="primary"
          size="md"
          className="bg-[#6d41ff] hover:bg-[#5b2fff] text-white"
          onClick={() => navigate("/admin/create-property")}
        >
          Create Property
        </Button>
      </header>

      {error && (
        <div className="text-red-500 bg-red-500/10 p-4 rounded-xl border border-red-500/20">
          {error}
        </div>
      )}

      {/* Stats Grid */}
      <div className="grid md:grid-cols-3 gap-4">
        <StatCard
          title="Total Properties"
          value={loading ? "..." : properties.length.toString()}
          icon={Building2}
          accent="bg-blue-500/20 text-blue-400"
        />
        <StatCard
          title="Total Value (ETH)"
          value={loading ? "..." : totalTokens.toFixed(2)}
          icon={Coins}
          accent="bg-green-500/20 text-green-400"
        />
        <StatCard
          title="Total Registered Users"
          value={totalUsers.toString()}
          icon={Users}
          accent="bg-purple-500/20 text-purple-400"
        />
      </div>

      {/* Admin Info */}
      {user && (
        <div className="rounded-3xl border border-[#1f1f1f] bg-[#111111] p-6 space-y-3">
          <div className="flex items-center gap-3">
            <div className="h-10 w-10 rounded-full bg-[#703BF7]/20 flex items-center justify-center">
              <Shield size={20} className="text-[#703BF7]" />
            </div>
            <div>
              <h2 className="font-semibold text-lg">Admin Account</h2>
              <p className="text-sm text-gray-400">{user.Email}</p>
            </div>
          </div>
          <div className="grid md:grid-cols-2 gap-4 mt-4">
            <div className="p-3 rounded-xl border border-[#1f1f1f] bg-[#0f0f0f]">
              <p className="text-xs text-gray-500">Wallet Address</p>
              <p className="text-sm font-mono mt-1 truncate">{user.WalletAddress || "Not set"}</p>
            </div>
            <div className="p-3 rounded-xl border border-[#1f1f1f] bg-[#0f0f0f]">
              <p className="text-xs text-gray-500">Role</p>
              <p className="text-sm font-medium mt-1 capitalize">{user.Role}</p>
            </div>
          </div>
        </div>
      )}

      {/* Quick Actions */}
      <div className="rounded-3xl border border-[#1f1f1f] bg-[#111111] p-6 space-y-4">
        <h2 className="font-semibold text-lg">Quick Actions</h2>
        <div className="flex flex-wrap gap-3">
          <Button
            variant="outline"
            size="md"
            onClick={() => navigate("/admin/users")}
            className="border-[#262626] bg-[#0f0f0f] text-white hover:bg-[#1a1a1a]"
          >
            <Users size={18} className="mr-2" />
            Manage User Approvals
          </Button>
          <Button
            variant="outline"
            size="md"
            onClick={() => navigate("/admin/create-property")}
            className="border-[#262626] bg-[#0f0f0f] text-white hover:bg-[#1a1a1a]"
          >
            <Building2 size={18} className="mr-2" />
            Create New Property
          </Button>
        </div>
      </div>

      {/* Properties Section */}
      <div className="rounded-3xl border border-[#1f1f1f] bg-[#111111] p-6 space-y-4">
        <div className="flex items-center justify-between">
          <h2 className="font-semibold text-lg">Platform Properties</h2>
          <span className="text-xs text-gray-500">{properties.length} total</span>
        </div>

        {loading ? (
          <div className="grid md:grid-cols-2 gap-4">
            {[...Array(4)].map((_, i) => (
              <div key={i} className="rounded-xl border border-[#1f1f1f] bg-[#0f0f0f] p-4 animate-pulse">
                <div className="h-4 bg-[#1a1a1a] rounded w-1/2 mb-3" />
                <div className="h-3 bg-[#1a1a1a] rounded w-1/3" />
              </div>
            ))}
          </div>
        ) : properties.length === 0 ? (
          <div className="text-center py-8">
            <p className="text-gray-400">No properties found</p>
            <Button
              variant="primary"
              size="md"
              className="bg-[#6d41ff] hover:bg-[#5b2fff] text-white mt-4"
              onClick={() => navigate("/admin/create-property")}
            >
              Create First Property
            </Button>
          </div>
        ) : (
          <div className="grid md:grid-cols-2 gap-4">
            {properties.map((property) => (
              <div
                key={property.ID}
                className="rounded-xl border border-[#1f1f1f] bg-[#0f0f0f] p-4 hover:border-[#703BF7] transition-colors"
              >
                <div className="flex items-start justify-between mb-3">
                  <div>
                    <h3 className="font-semibold">Property #{property.ID.substring(0, 8)}</h3>
                    <p className="text-xs text-gray-400">
                      Owner: {property.OwnerWallet.substring(0, 6)}...{property.OwnerWallet.substring(38)}
                    </p>
                  </div>
                  <div className="text-right">
                    <p className="text-sm font-semibold">{property.Valuation} ETH</p>
                    <span className={`text-xs px-2 py-0.5 rounded-full ${
                      property.Status === "Active"
                        ? "bg-green-500/20 text-green-400"
                        : "bg-gray-500/20 text-gray-400"
                    }`}>
                      {property.Status}
                    </span>
                  </div>
                </div>

                <div className="flex gap-2">
                  <Button
                    variant="outline"
                    size="sm"
                    onClick={() => navigate(`/properties/${property.ID}`)}
                    className="flex-1 border-[#262626] text-white hover:bg-[#1a1a1a] text-xs"
                  >
                    View Details
                  </Button>
                </div>
              </div>
            ))}
          </div>
        )}
      </div>
    </div>
  );
}

export default AdminDashboardHome;

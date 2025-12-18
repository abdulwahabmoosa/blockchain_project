import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { TrendingUp, TrendingDown, ArrowUpRight } from "lucide-react";
import { Button } from "../Components/atoms/Button";
import { api } from "../lib/api";
import type { Property } from "../types";

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

function MyInvestments() {
  const [properties, setProperties] = useState<Property[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");
  const navigate = useNavigate();

  useEffect(() => {
    const fetchProperties = async () => {
      try {
        const data = await api.getProperties();
        setProperties(data);
      } catch (err: any) {
        console.error("Failed to fetch properties:", err);
        setError("Failed to load investments");
      } finally {
        setLoading(false);
      }
    };

    fetchProperties();
  }, []);

  const totalValue = properties.reduce((sum, p) => sum + p.Valuation, 0);
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
        <Button
          variant="primary"
          size="md"
          className="bg-[#6d41ff] hover:bg-[#5b2fff] text-white"
          onClick={() => navigate("/dashboard/properties")}
        >
          Browse Properties
        </Button>
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
                <div className="h-16 w-16 rounded-xl overflow-hidden flex-shrink-0">
                  <GradientThumb />
                </div>
                <div className="flex-1 min-w-0">
                  <h3 className="font-semibold truncate">
                    Property #{property.ID.substring(0, 8)}
                  </h3>
                  <p className="text-sm text-gray-400">
                    {property.Valuation} ETH
                  </p>
                </div>
                <div className="text-right flex-shrink-0">
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
                  <p className="text-xs text-gray-500 mt-1">
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

      {/* Performance Chart Placeholder */}
      <div className="rounded-3xl border border-[#1f1f1f] bg-[#111111] p-6 space-y-4">
        <div className="flex items-center justify-between">
          <h2 className="font-semibold text-lg">Performance</h2>
          <div className="flex gap-2">
            <button className="px-3 py-1 text-xs rounded-lg bg-[#703BF7] text-white">
              1M
            </button>
            <button className="px-3 py-1 text-xs rounded-lg text-gray-400 hover:bg-[#1a1a1a]">
              3M
            </button>
            <button className="px-3 py-1 text-xs rounded-lg text-gray-400 hover:bg-[#1a1a1a]">
              1Y
            </button>
            <button className="px-3 py-1 text-xs rounded-lg text-gray-400 hover:bg-[#1a1a1a]">
              All
            </button>
          </div>
        </div>
        <div className="h-48 rounded-xl border border-[#1f1f1f] bg-[#0f0f0f] flex items-center justify-center">
          <p className="text-gray-500 text-sm">Performance chart coming soon</p>
        </div>
      </div>
    </div>
  );
}

export default MyInvestments;

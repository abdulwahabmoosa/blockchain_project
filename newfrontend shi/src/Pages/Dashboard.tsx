import "../index.css";
import { useEffect, useState } from "react";
import { Navbar } from "../Components/organisms/Navbar";
import { Button } from "../Components/atoms/Button";
import { api } from "../lib/api";
import type { Property } from "../types";
import { useNavigate } from "react-router-dom";

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

function DashboardPage() {
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
        setError("Failed to load properties");
      } finally {
        setLoading(false);
      }
    };

    fetchProperties();
  }, []);

  return (
    <div className="min-h-screen w-full bg-[#0f0f0f] text-white">
      <Navbar currentPage="/dashboard" />
      <main className="max-w-6xl mx-auto px-4 pt-24 pb-16 space-y-8">
        <header className="flex items-center justify-between">
          <div>
            <p className="text-sm text-gray-400">Welcome back</p>
            <h1 className="text-3xl font-semibold">Dashboard</h1>
          </div>
          <Button
            variant="primary"
            size="md"
            className="bg-[#6d41ff] hover:bg-[#5b2fff] text-white"
            onClick={() => navigate("/create-property")}
          >
            New Investment
          </Button>
        </header>

        {error && (
          <div className="text-red-500 bg-red-500/10 p-4 rounded-xl border border-red-500/20">
            {error}
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
                    <h3 className="text-xl font-semibold">Property #{properties[0].ID.substring(0, 8)}</h3>
                    <p className="text-sm text-gray-400">
                      {properties[0].Valuation} ETH · {properties[0].Status}
                    </p>
                  </div>
                </div>
              ) : (
                <div className="space-y-3 flex items-center justify-center border border-[#1f1f1f] rounded-2xl h-full min-h-[200px]">
                   <p className="text-gray-500">No properties available</p>
                </div>
              )}

              <div className="space-y-3">
                <div className="grid grid-cols-2 gap-3">
                  <Card title="Total Properties" value={properties.length.toString()} sub="Platform" />
                  <Card title="Active" value={properties.filter(p => p.Status === "Active").length.toString()} sub="Status" accent="bg-[#2e2e2e]" />
                  <Card title="Avg Price" value="0.5 ETH" sub="Est." accent="bg-[#2c1f5e]" />
                  <Card title="Watchlist" value="0" sub="Saved" accent="bg-[#1f3a5e]" />
                </div>
                <Button
                  variant="outline"
                  size="md"
                  className="border-[#262626] bg-[#111111] text-white hover:bg-[#1a1a1a] w-full"
                  onClick={() => navigate("/properties")}
                >
                  View All Properties
                </Button>
              </div>
            </div>
          </div>

          {/* Balance card */}
          <div className="rounded-3xl border border-[#1f1f1f] bg-linear-to-br from-[#6d41ff] via-[#5a3ccc] to-[#1b132f] p-5 space-y-3">
            <p className="text-sm text-white/80">My Wallet</p>
            <p className="text-3xl font-semibold">0.00 ETH</p>
            <p className="text-sm text-white/80">Monthly Profit</p>
            <p className="text-2xl font-semibold">0.00 ETH</p>
            <div className="flex gap-2">
              <span className="rounded-full bg-white/10 px-3 py-1 text-xs">+0%</span>
            </div>
            <Button
              variant="outline"
              size="md"
              className="border-white/30 text-white hover:bg-white/10"
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
                    <p className="text-sm font-semibold truncate max-w-[150px]">Property #{item.ID.substring(0, 8)}</p>
                    <p className="text-xs text-gray-400">{item.Valuation} ETH</p>
                  </div>
                  <p className="text-xs text-gray-300">{item.Status}</p>
                </div>
              ))}
              {properties.length === 0 && !loading && (
                 <p className="text-sm text-gray-500 text-center py-4">No properties found</p>
              )}
            </div>
          </div>

          {/* Investment stats */}
          <div className="rounded-3xl border border-[#1f1f1f] bg-[#111111] p-5 space-y-4 md:col-span-2">
            <div className="flex items-center justify-between">
              <p className="font-semibold">Investment Stats</p>
              <span className="text-xs text-gray-500">Weekly</span>
            </div>
            <div className="grid md:grid-cols-2 gap-4">
              <div className="space-y-3">
                <Card title="Total Investment" value="0.00 ETH" sub="Active" />
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
              {properties.length > 0 ? properties.map((item) => (
                <div
                  key={item.ID}
                  className="min-w-[180px] rounded-2xl border border-[#1f1f1f] bg-[#0f0f0f] p-3 space-y-2"
                >
                  <div className="h-24 rounded-xl border border-[#1f1f1f] overflow-hidden">
                      <GradientThumb />
                  </div>
                  <p className="text-sm font-semibold truncate">Property #{item.ID.substring(0, 8)}</p>
                  <p className="text-xs text-gray-400">{item.Valuation} ETH</p>
                </div>
              )) : (
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
               {/* Placeholder for top picks until we have a logic for it */}
               <p className="text-sm text-gray-500 text-center py-4">Coming soon</p>
            </div>
          </div>
        </section>
      </main>
    </div>
  );
}

export default DashboardPage;


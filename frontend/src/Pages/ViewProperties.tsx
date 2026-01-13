import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { Search, Filter } from "lucide-react";
import { Button } from "../Components/atoms/Button";
import { api } from "../lib/api";
import type { Property } from "../types";

const GradientThumb = () => (
  <div className="rounded-2xl border border-[#1f1f1f] bg-linear-to-br from-[#191919] via-[#1f1f1f] to-[#121212] w-full h-full" />
);

const PropertyCard = ({
  property,
  onClick,
}: {
  property: Property;
  onClick: () => void;
}) => (
  <div
    onClick={onClick}
    className="rounded-2xl border border-[#1f1f1f] bg-[#111111] p-4 space-y-3 cursor-pointer hover:border-[#703BF7] transition-colors"
  >
    <div className="h-40 rounded-xl overflow-hidden">
      <GradientThumb />
    </div>
    <div className="space-y-1">
      <h3 className="text-lg font-semibold truncate">
        {property.Name || `Property #${property.ID.substring(0, 8)}`}
      </h3>
      <p className="text-sm text-gray-400">{property.Valuation} ETH</p>
    </div>
    <div className="flex items-center justify-between">
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
      <span className="text-xs text-gray-500">
        {new Date(property.CreatedAt).toLocaleDateString()}
      </span>
    </div>
  </div>
);

function ViewProperties() {
  const [properties, setProperties] = useState<Property[]>([]);
  const [filteredProperties, setFilteredProperties] = useState<Property[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");
  const [searchQuery, setSearchQuery] = useState("");
  const [statusFilter, setStatusFilter] = useState<string>("all");
  const navigate = useNavigate();

  useEffect(() => {
    const fetchProperties = async () => {
      try {
        const data = await api.getProperties();
        setProperties(data);
        setFilteredProperties(data);
      } catch (err: any) {
        console.error("Failed to fetch properties:", err);
        setError("Failed to load properties");
      } finally {
        setLoading(false);
      }
    };

    fetchProperties();
  }, []);

  useEffect(() => {
    let filtered = properties;

    if (searchQuery) {
      filtered = filtered.filter((p) =>
        p.ID.toLowerCase().includes(searchQuery.toLowerCase())
      );
    }

    if (statusFilter !== "all") {
      filtered = filtered.filter((p) => p.Status === statusFilter);
    }

    setFilteredProperties(filtered);
  }, [searchQuery, statusFilter, properties]);

  return (
    <div className="space-y-6">
      <header className="flex flex-col md:flex-row md:items-center justify-between gap-4">
        <div>
          <h1 className="text-3xl font-semibold">Properties</h1>
          <p className="text-sm text-gray-400">
            Browse and explore available properties
          </p>
        </div>
        <Button
          variant="primary"
          size="md"
          className="bg-[#6d41ff] hover:bg-[#5b2fff] text-white"
          onClick={() => navigate("/dashboard/upload")}
        >
          Upload Property
        </Button>
      </header>

      {/* Search and Filter */}
      <div className="flex flex-col md:flex-row gap-4">
        <div className="relative flex-1">
          <Search
            className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400"
            size={20}
          />
          <input
            type="text"
            placeholder="Search by property ID..."
            value={searchQuery}
            onChange={(e) => setSearchQuery(e.target.value)}
            className="w-full pl-10 pr-4 py-3 rounded-xl border border-[#1f1f1f] bg-[#111111] text-white placeholder-gray-500 focus:outline-none focus:border-[#6d41ff]"
          />
        </div>
        <div className="relative">
          <Filter
            className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400"
            size={20}
          />
          <select
            value={statusFilter}
            onChange={(e) => setStatusFilter(e.target.value)}
            className="pl-10 pr-8 py-3 rounded-xl border border-[#1f1f1f] bg-[#111111] text-white appearance-none cursor-pointer focus:outline-none focus:border-[#6d41ff]"
          >
            <option value="all">All Status</option>
            <option value="Active">Active</option>
            <option value="Paused">Paused</option>
            <option value="Disputed">Disputed</option>
            <option value="Closed">Closed</option>
          </select>
        </div>
      </div>

      {error && (
        <div className="text-red-500 bg-red-500/10 p-4 rounded-xl border border-red-500/20">
          {error}
        </div>
      )}

      {loading ? (
        <div className="grid md:grid-cols-2 lg:grid-cols-3 gap-4">
          {[...Array(6)].map((_, i) => (
            <div
              key={i}
              className="rounded-2xl border border-[#1f1f1f] bg-[#111111] p-4 space-y-3 animate-pulse"
            >
              <div className="h-40 rounded-xl bg-[#1a1a1a]" />
              <div className="h-6 rounded bg-[#1a1a1a] w-3/4" />
              <div className="h-4 rounded bg-[#1a1a1a] w-1/2" />
            </div>
          ))}
        </div>
      ) : filteredProperties.length > 0 ? (
        <div className="grid md:grid-cols-2 lg:grid-cols-3 gap-4">
          {filteredProperties.map((property) => (
            <PropertyCard
              key={property.ID}
              property={property}
              onClick={() => navigate(`/properties/${property.ID}`)}
            />
          ))}
        </div>
      ) : (
        <div className="rounded-3xl border border-[#1f1f1f] bg-[#111111] p-12 text-center">
          <p className="text-gray-400 text-lg">No properties found</p>
          <p className="text-gray-500 text-sm mt-2">
            {searchQuery || statusFilter !== "all"
              ? "Try adjusting your search or filter"
              : "Be the first to upload a property!"}
          </p>
          {!searchQuery && statusFilter === "all" && (
            <Button
              variant="primary"
              size="md"
              className="bg-[#6d41ff] hover:bg-[#5b2fff] text-white mt-4"
              onClick={() => navigate("/dashboard/upload")}
            >
              Upload Property
            </Button>
          )}
        </div>
      )}
    </div>
  );
}

export default ViewProperties;




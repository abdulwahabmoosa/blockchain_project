import { useEffect, useState } from "react";
import { Link as RouterLink } from "react-router-dom";
import { Navbar } from "../Components/organisms/Navbar";
import { Button } from "../Components/atoms/Button";
import { api } from "../lib/api";
import { useWallet } from "../hooks/useWallet";
import type { Property } from "../types";

function PropertiesPage() {
  const [properties, setProperties] = useState<Property[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");
  const { isConnected, connectRegisteredWallet } = useWallet();
  const [user, setUser] = useState<any>(null);

  // Load user data
  useEffect(() => {
    const userStr = localStorage.getItem("user");
    if (userStr) {
      setUser(JSON.parse(userStr));
    }
  }, []);

  // Auto-connect wallet for registered users
  useEffect(() => {
    const autoConnectWallet = async () => {
      if (user && user.WalletAddress && !isConnected) {
        console.log('🔗 Auto-connecting registered wallet for user:', user.Email);
        try {
          await connectRegisteredWallet(user.WalletAddress);
        } catch (err) {
          console.error('❌ Auto-connect failed:', err);
        }
      }
    };

    autoConnectWallet();
  }, [user, isConnected, connectRegisteredWallet]);

  useEffect(() => {
    const fetchProperties = async () => {
      try {
        const data = await api.getProperties();
        setProperties(data);
      } catch (err: any) {
        setError(err.message || "Failed to load properties");
      } finally {
        setLoading(false);
      }
    };

    fetchProperties();
  }, []);

  if (loading) {
    return (
      <div className="min-h-screen bg-[#0f0f0f] text-white">
        <Navbar currentPage="/properties" />
        <div className="flex items-center justify-center pt-32">
          <p>Loading properties...</p>
        </div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="min-h-screen bg-[#0f0f0f] text-white">
        <Navbar currentPage="/properties" />
        <div className="flex items-center justify-center pt-32">
          <p className="text-red-500">{error}</p>
        </div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-[#0f0f0f] text-white">
      <Navbar currentPage="/properties" />

      <main className="max-w-7xl mx-auto px-6 pt-28 pb-16">
        <div className="space-y-8">
          {/* Header */}
          <div className="text-center space-y-4">
            <h1 className="text-4xl font-bold">Available Properties</h1>
            <p className="text-gray-400 max-w-2xl mx-auto">
              Discover tokenized real estate properties. Invest in fractions of high-value assets
              and earn passive income through rental revenue distributions.
            </p>
          </div>


          {/* Properties Grid */}
          {properties.length === 0 ? (
            <div className="text-center py-16">
              <p className="text-gray-400">No properties available yet.</p>
            </div>
          ) : (
            <div className="grid md:grid-cols-2 lg:grid-cols-3 gap-6">
              {properties.map((property) => (
                <div
                  key={property.ID}
                  className="bg-[#1A1A1A] border border-[#262626] rounded-lg overflow-hidden hover:border-[#6d41ff]/50 transition-colors"
                >
                  {/* Property Image Placeholder */}
                  <div className="aspect-video bg-gradient-to-br from-[#1f1f1f] via-[#2a2a2a] to-[#111111] flex items-center justify-center">
                    <span className="text-gray-500">Property Image</span>
                  </div>

                  <div className="p-6 space-y-4">
                    <div>
                      <h3 className="text-xl font-semibold mb-2">
                        Property #{property.ID.substring(0, 8)}
                      </h3>
                      <p className="text-sm text-gray-400">
                        Owner: {property.OwnerWallet.substring(0, 6)}...{property.OwnerWallet.substring(38)}
                      </p>
                    </div>

                    <div className="space-y-2">
                      <div className="flex justify-between text-sm">
                        <span className="text-gray-400">Valuation:</span>
                        <span className="font-semibold">{property.Valuation} ETH</span>
                      </div>
                      <div className="flex justify-between text-sm">
                        <span className="text-gray-400">Status:</span>
                        <span className={`font-semibold ${
                          property.Status === 'Active' ? 'text-green-400' :
                          property.Status === 'Paused' ? 'text-yellow-400' :
                          property.Status === 'Disputed' ? 'text-red-400' : 'text-gray-400'
                        }`}>
                          {property.Status}
                        </span>
                      </div>
                      <div className="flex justify-between text-sm">
                        <span className="text-gray-400">Created:</span>
                        <span>{new Date(property.CreatedAt).toLocaleDateString()}</span>
                      </div>
                    </div>

                    <RouterLink to={`/properties/${property.ID}`}>
                      <Button
                        variant="primary"
                        size="sm"
                        className="w-full bg-[#6d41ff] hover:bg-[#5b2fff] text-white"
                      >
                        View Details
                      </Button>
                    </RouterLink>
                  </div>
                </div>
              ))}
            </div>
          )}
        </div>
      </main>
    </div>
  );
}

export default PropertiesPage;

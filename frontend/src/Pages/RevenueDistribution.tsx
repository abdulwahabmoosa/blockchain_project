import { useEffect, useState } from "react";
import { Navbar } from "../Components/organisms/Navbar";
import { Button } from "../Components/atoms/Button";
import { api } from "../lib/api";
import type { Property } from "../types";

interface PropertyWithDetails extends Property {
  tokenSymbol?: string;
}

function RevenueDistributionPage() {
  const [properties, setProperties] = useState<PropertyWithDetails[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");
  const [selectedProperty, setSelectedProperty] = useState<PropertyWithDetails | null>(null);
  const [amount, setAmount] = useState("");
  const [stablecoinAddress, setStablecoinAddress] = useState("");
  const [distributing, setDistributing] = useState(false);

  // Common stablecoin addresses on Sepolia (example addresses)
  const commonStablecoins = [
    { name: "USDC (Mock)", address: "0x1234567890123456789012345678901234567890" },
    { name: "USDT (Mock)", address: "0x0987654321098765432109876543210987654321" },
    { name: "Custom Address", address: "" }
  ];

  useEffect(() => {
    loadProperties();
  }, []);

  const loadProperties = async () => {
    try {
      const data = await api.getProperties();
      setProperties(data);
    } catch (err: any) {
      setError(err.message || "Failed to load properties");
    } finally {
      setLoading(false);
    }
  };

  const handleDistributeRevenue = async () => {
    if (!selectedProperty) {
      setError("Please select a property");
      return;
    }

    if (!amount || parseFloat(amount) <= 0) {
      setError("Please enter a valid amount");
      return;
    }

    if (!stablecoinAddress) {
      setError("Please enter a stablecoin address");
      return;
    }

    setDistributing(true);
    setError("");

    try {
      const response = await api.distributeRevenue(
        selectedProperty.OnchainTokenAddress,
        stablecoinAddress,
        parseFloat(amount)
      );

      alert(`Revenue distributed successfully!\nTransaction Hash: ${response.tx_hash}`);

      // Reset form
      setSelectedProperty(null);
      setAmount("");
      setStablecoinAddress("");
    } catch (err: any) {
      setError(err.message || "Failed to distribute revenue");
    } finally {
      setDistributing(false);
    }
  };

  if (loading) {
    return (
      <div className="min-h-screen bg-[#0f0f0f] text-white">
        <Navbar currentPage="/admin/distribute-revenue" />
        <div className="flex items-center justify-center pt-32">
          <p>Loading properties...</p>
        </div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-[#0f0f0f] text-white">
      <Navbar currentPage="/admin/distribute-revenue" />

      <main className="max-w-4xl mx-auto px-6 pt-28 pb-16">
        <div className="space-y-8">
          {/* Header */}
          <div className="text-center space-y-4">
            <h1 className="text-3xl font-bold">Revenue Distribution</h1>
            <p className="text-gray-400 max-w-2xl mx-auto">
              Distribute rental income or property revenue to token holders. Funds will be proportionally distributed based on token ownership at the time of the last snapshot.
            </p>
          </div>

          {/* Error Display */}
          {error && (
            <div className="text-red-500 bg-red-500/10 p-4 rounded-xl border border-red-500/20">
              {error}
            </div>
          )}

          {/* Distribution Form */}
          <div className="bg-[#1A1A1A] border border-[#262626] rounded-xl p-8">
            <div className="space-y-6">
              <div>
                <h2 className="text-2xl font-semibold mb-2">Distribute Revenue</h2>
                <p className="text-gray-400">
                  Select a property and specify the revenue amount and stablecoin to distribute.
                </p>
              </div>

              {/* Property Selection */}
              <div>
                <label className="block text-sm font-medium mb-3">Select Property</label>
                <div className="grid gap-3">
                  {properties.map((property) => (
                    <div
                      key={property.ID}
                      onClick={() => setSelectedProperty(property)}
                      className={`p-4 rounded-lg border cursor-pointer transition-colors ${
                        selectedProperty?.ID === property.ID
                          ? 'border-[#6d41ff] bg-[#6d41ff]/10'
                          : 'border-[#262626] bg-[#0f0f0f] hover:border-[#262626]'
                      }`}
                    >
                      <div className="flex justify-between items-start">
                        <div>
                          <h3 className="font-semibold">Property #{property.ID.substring(0, 8)}</h3>
                          <p className="text-sm text-gray-400">
                            Valuation: {property.Valuation} ETH • Status: {property.Status}
                          </p>
                          <p className="text-xs text-gray-500 font-mono mt-1">
                            Token: {property.OnchainTokenAddress.substring(0, 10)}...
                          </p>
                        </div>
                        {selectedProperty?.ID === property.ID && (
                          <div className="text-[#6d41ff]">✓</div>
                        )}
                      </div>
                    </div>
                  ))}
                </div>
              </div>

              {/* Amount Input */}
              <div>
                <label className="block text-sm font-medium mb-2">Revenue Amount</label>
                <input
                  type="number"
                  placeholder="e.g., 1000"
                  step="0.01"
                  value={amount}
                  onChange={(e) => setAmount(e.target.value)}
                  className="w-full p-3 rounded-lg border border-[#262626] bg-[#0f0f0f] text-white focus:border-[#6d41ff] outline-none"
                />
                <p className="text-xs text-gray-400 mt-1">Amount of stablecoins to distribute to token holders</p>
              </div>

              {/* Stablecoin Selection */}
              <div>
                <label className="block text-sm font-medium mb-2">Stablecoin Contract</label>
                <div className="space-y-3">
                  {commonStablecoins.map((coin) => (
                    <div key={coin.name} className="flex items-center space-x-3">
                      <input
                        type="radio"
                        name="stablecoin"
                        value={coin.address}
                        checked={stablecoinAddress === coin.address}
                        onChange={(e) => setStablecoinAddress(e.target.value)}
                        className="accent-[#6d41ff]"
                      />
                      <div>
                        <div className="text-sm font-medium">{coin.name}</div>
                        {coin.address && (
                          <div className="text-xs text-gray-400 font-mono">
                            {coin.address.substring(0, 10)}...{coin.address.substring(38)}
                          </div>
                        )}
                      </div>
                    </div>
                  ))}

                  {/* Custom Address Input */}
                  <div className="flex items-center space-x-3">
                    <input
                      type="radio"
                      name="stablecoin"
                      checked={stablecoinAddress !== "" && !commonStablecoins.some(c => c.address === stablecoinAddress)}
                      onChange={() => setStablecoinAddress("")}
                      className="accent-[#6d41ff]"
                    />
                    <input
                      type="text"
                      placeholder="Enter custom stablecoin address"
                      value={stablecoinAddress && !commonStablecoins.some(c => c.address === stablecoinAddress) ? stablecoinAddress : ""}
                      onChange={(e) => setStablecoinAddress(e.target.value)}
                      className="flex-1 p-2 rounded border border-[#262626] bg-[#0f0f0f] text-white text-sm focus:border-[#6d41ff] outline-none"
                    />
                  </div>
                </div>
              </div>

              {/* Distribution Summary */}
              {selectedProperty && amount && stablecoinAddress && (
                <div className="bg-blue-500/10 border border-blue-500/20 rounded-lg p-4">
                  <h3 className="text-blue-400 font-semibold mb-2">📊 Distribution Summary</h3>
                  <div className="text-sm text-gray-300 space-y-1">
                    <p><strong>Property:</strong> #{selectedProperty.ID.substring(0, 8)}</p>
                    <p><strong>Token Address:</strong> {selectedProperty.OnchainTokenAddress.substring(0, 10)}...</p>
                    <p><strong>Amount:</strong> {amount} stablecoins</p>
                    <p><strong>Stablecoin:</strong> {stablecoinAddress.substring(0, 10)}...</p>
                  </div>
                  <p className="text-xs text-blue-300 mt-2">
                    💡 Revenue will be distributed proportionally to all token holders based on their current balance.
                  </p>
                </div>
              )}

              {/* Distribute Button */}
              <Button
                variant="primary"
                size="lg"
                onClick={handleDistributeRevenue}
                disabled={distributing || !selectedProperty || !amount || !stablecoinAddress}
                className="w-full bg-[#6d41ff] hover:bg-[#5b2fff] text-white disabled:opacity-50"
              >
                {distributing ? "Distributing Revenue..." : "Distribute Revenue"}
              </Button>
            </div>
          </div>

          {/* How It Works */}
          <div className="bg-green-500/10 border border-green-500/20 rounded-xl p-6">
            <h3 className="text-green-400 font-semibold mb-3">ℹ️ How Revenue Distribution Works</h3>
            <div className="text-sm text-gray-300 space-y-2">
              <p>
                <strong>1. Snapshot Creation:</strong> When distributing revenue, the system takes a snapshot of token balances.
              </p>
              <p>
                <strong>2. Proportional Distribution:</strong> Revenue is distributed based on token ownership at snapshot time.
              </p>
              <p>
                <strong>3. Claiming Process:</strong> Token holders can claim their share through the property details page.
              </p>
              <p>
                <strong>4. Smart Contract:</strong> All distributions are recorded on-chain and managed by the RevenueDistribution contract.
              </p>
            </div>
          </div>
        </div>
      </main>
    </div>
  );
}

export default RevenueDistributionPage;




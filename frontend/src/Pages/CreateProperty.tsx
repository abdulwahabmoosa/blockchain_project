import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { Navbar } from "../Components/organisms/Navbar";
import { Button } from "../Components/atoms/Button";
import { api } from "../lib/api";
import { useWallet } from "../hooks/useWallet";

function CreatePropertyPage() {
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");
  const [success, setSuccess] = useState(false);
  const [txHash, setTxHash] = useState("");

  // Form fields
  const [selectedFiles, setSelectedFiles] = useState<File[]>([]);
  const [name, setName] = useState("");
  const [valuation, setValuation] = useState("");

  // Calculated fields
  const symbol = "ETH"; // Always ETH
  const tokenSupply = valuation ? (parseFloat(valuation) / 10).toString() : "";

  const navigate = useNavigate();
  const { address } = useWallet();

  const handleFileSelect = (event: React.ChangeEvent<HTMLInputElement>) => {
    const files = Array.from(event.target.files || []);
    if (files.length > 0) {
      setSelectedFiles(files);
      setError("");
    }
  };

  const removeFile = (index: number) => {
    setSelectedFiles(selectedFiles.filter((_, i) => i !== index));
  };

  const handleCreateProperty = async () => {
    if (selectedFiles.length === 0) {
      setError("Please select at least one file to upload");
      return;
    }

    if (!address) {
      setError("Please connect your wallet");
      return;
    }

    if (!name || !valuation || !tokenSupply) {
      setError("Please fill all required fields");
      return;
    }

    setLoading(true);
    setError("");

    try {
      const response = await api.createProperty(
        {
          owner_address: address,
          name,
          symbol: symbol.toUpperCase(),
          data_hash: "", // Not needed - files are sent directly
          valuation: parseFloat(valuation),
          token_supply: parseInt(tokenSupply)
        },
        selectedFiles
      );
      
      setTxHash(response.tx_hash);
      setSuccess(true);
    } catch (err: any) {
      setError(err.message || "Failed to create property");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="min-h-screen bg-[#0f0f0f] text-white">
      <Navbar currentPage="/create-property" />

      <main className="max-w-4xl mx-auto px-6 pt-28 pb-16">
        <div className="space-y-8">
          {/* Header */}
          <div className="text-center space-y-4">
            <h1 className="text-4xl font-bold">Create New Property</h1>
            <p className="text-gray-400 max-w-2xl mx-auto">
              Tokenize a real estate property by uploading documentation and creating fractional ownership tokens on the blockchain.
            </p>
          </div>

          {success && (
            <div className="bg-green-500/10 border border-green-500/20 rounded-xl p-6 text-center">
              <h2 className="text-2xl font-semibold text-green-400 mb-2">Property Created!</h2>
              <p className="text-gray-400 mb-4">Your property has been successfully tokenized on the blockchain.</p>
              {txHash && (
                <div className="bg-[#0f0f0f] rounded-lg p-3 mb-4">
                  <p className="text-xs text-gray-500 mb-1">Transaction Hash</p>
                  <p className="text-sm font-mono text-gray-300 break-all">{txHash}</p>
                </div>
              )}
              <Button
                variant="primary"
                size="lg"
                onClick={() => navigate("/dashboard")}
                className="bg-[#6d41ff] hover:bg-[#5b2fff] text-white"
              >
                Go to Dashboard
              </Button>
            </div>
          )}

          {!success && (
            <>
              {/* Error Display */}
              {error && (
                <div className="text-red-500 bg-red-500/10 p-4 rounded-xl border border-red-500/20">
                  {error}
                </div>
              )}

              {/* Property Creation Form */}
              <div className="bg-[#1A1A1A] border border-[#262626] rounded-xl p-8">
                <div className="space-y-6">
                  <div>
                    <h2 className="text-2xl font-semibold mb-2">Property Details</h2>
                    <p className="text-gray-400">
                      Configure the property tokenization parameters and upload documents.
                    </p>
                  </div>

                  {/* File Upload Section */}
                  <div className="space-y-4">
                    <div>
                      <label className="block text-sm font-medium mb-2">Property Documents</label>
                      <input
                        type="file"
                        accept="image/*,.pdf,.doc,.docx"
                        onChange={handleFileSelect}
                        multiple
                        className="w-full p-3 rounded-lg border border-[#262626] bg-[#0f0f0f] text-white file:mr-4 file:py-2 file:px-4 file:rounded-lg file:border-0 file:bg-[#6d41ff] file:text-white file:cursor-pointer hover:file:bg-[#5b2fff]"
                      />
                      <p className="text-xs text-gray-400 mt-2">Multiple files supported</p>
                    </div>

                    {selectedFiles.length > 0 && (
                      <div className="space-y-2">
                        <p className="text-sm text-gray-400">Selected files ({selectedFiles.length}):</p>
                        {selectedFiles.map((file, index) => (
                          <div key={index} className="flex items-center justify-between p-3 rounded-lg border border-[#262626] bg-[#0f0f0f]">
                            <span className="text-sm text-gray-300">{file.name} ({(file.size / 1024).toFixed(2)} KB)</span>
                            <button
                              type="button"
                              onClick={() => removeFile(index)}
                              className="text-red-400 hover:text-red-300 text-sm"
                            >
                              Remove
                            </button>
                          </div>
                        ))}
                      </div>
                    )}
                  </div>

                  <div className="grid md:grid-cols-2 gap-6">
                    <div className="space-y-4">
                      <div>
                        <label className="block text-sm font-medium mb-2">Property Name</label>
                        <input
                          type="text"
                          placeholder="e.g., Downtown Office Building"
                          value={name}
                          onChange={(e) => setName(e.target.value)}
                          className="w-full p-3 rounded-lg border border-[#262626] bg-[#0f0f0f] text-white focus:border-[#6d41ff] outline-none"
                        />
                      </div>
                    </div>

                    <div className="space-y-4">
                      <div>
                        <label className="block text-sm font-medium mb-2">Property Valuation (ETH)</label>
                        <input
                          type="number"
                          placeholder="e.g., 500000"
                          step="0.01"
                          value={valuation}
                          onChange={(e) => setValuation(e.target.value)}
                          className="w-full p-3 rounded-lg border border-[#262626] bg-[#0f0f0f] text-white focus:border-[#6d41ff] outline-none"
                        />
                        <p className="text-xs text-gray-400 mt-1">Total property value in ETH</p>
                      </div>

                      {/* Auto-calculated Token Symbol */}
                      <div>
                        <label className="block text-sm font-medium mb-2">Token Symbol</label>
                        <div className="w-full p-3 rounded-lg border border-[#262626] bg-[#0f0f0f] text-gray-300">
                          ETH
                        </div>
                        <p className="text-xs text-gray-400 mt-1">Fixed as ETH for all properties</p>
                      </div>

                      {/* Auto-calculated Token Supply */}
                      <div>
                        <label className="block text-sm font-medium mb-2">Token Supply</label>
                        <div className="w-full p-3 rounded-lg border border-[#262626] bg-[#0f0f0f] text-gray-300">
                          {tokenSupply || "0"}
                        </div>
                        <p className="text-xs text-gray-400 mt-1">Automatically calculated as 10% of valuation</p>
                      </div>
                    </div>
                  </div>

                  {/* Owner Address Display */}
                  <div className="bg-[#0f0f0f] border border-[#262626] rounded-lg p-4">
                    <div className="flex justify-between items-center">
                      <span className="text-sm text-gray-400">Owner Address:</span>
                      <span className="text-sm font-mono">
                        {address?.substring(0, 6)}...{address?.substring(38)}
                      </span>
                    </div>
                  </div>

                  <div className="flex gap-4">
                    <Button
                      variant="outline"
                      size="lg"
                      onClick={() => navigate("/dashboard")}
                      className="flex-1 border-[#262626] bg-[#0f0f0f] text-white hover:bg-[#262626]"
                    >
                      Cancel
                    </Button>
                    <Button
                      variant="primary"
                      size="lg"
                      onClick={handleCreateProperty}
                      disabled={loading || !name || !valuation || !tokenSupply || selectedFiles.length === 0}
                      className="flex-1 bg-[#6d41ff] hover:bg-[#5b2fff] text-white disabled:opacity-50"
                    >
                      {loading ? "Creating Property..." : "Create Property"}
                    </Button>
                  </div>
                </div>
              </div>
            </>
          )}
        </div>
      </main>
    </div>
  );
}

export default CreatePropertyPage;

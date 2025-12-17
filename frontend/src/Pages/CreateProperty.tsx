import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { Navbar } from "../Components/organisms/Navbar";
import { Button } from "../Components/atoms/Button";
import { api } from "../lib/api";
import type { CreatePropertyPayload, UploadResponse } from "../types";
import { useWallet } from "../hooks/useWallet";

function CreatePropertyPage() {
  const [step, setStep] = useState<'upload' | 'details'>('upload');
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");
  const [uploadResponse, setUploadResponse] = useState<UploadResponse | null>(null);

  // Form fields
  const [selectedFile, setSelectedFile] = useState<File | null>(null);
  const [name, setName] = useState("");
  const [valuation, setValuation] = useState("");

  // Calculated fields
  const symbol = "ETH"; // Always ETH
  const tokenSupply = valuation ? (parseFloat(valuation) / 10).toString() : "";

  const navigate = useNavigate();
  const { address } = useWallet();

  const handleFileSelect = (event: React.ChangeEvent<HTMLInputElement>) => {
    const file = event.target.files?.[0];
    if (file) {
      setSelectedFile(file);
      setError("");
    }
  };

  const handleUpload = async () => {
    if (!selectedFile) {
      setError("Please select a file to upload");
      return;
    }

    setLoading(true);
    setError("");

    try {
      const response = await api.uploadFile(selectedFile);
      setUploadResponse(response);
      setStep('details');
    } catch (err: any) {
      setError(err.message || "Failed to upload file");
    } finally {
      setLoading(false);
    }
  };

  const handleCreateProperty = async () => {
    if (!uploadResponse) {
      setError("Please upload a file first");
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
      const payload: CreatePropertyPayload = {
        owner_address: address,
        name,
        symbol: symbol.toUpperCase(),
        data_hash: uploadResponse.ipfs_hash,
        valuation: parseFloat(valuation),
        token_supply: parseInt(tokenSupply)
      };

      const response = await api.createProperty(payload);
      alert(`Property created successfully!\nTransaction Hash: ${response.tx_hash}`);
      navigate("/admin");
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

          {/* Progress Indicator */}
          <div className="flex items-center justify-center space-x-4">
            <div className={`flex items-center space-x-2 ${step === 'upload' ? 'text-[#6d41ff]' : 'text-gray-600'}`}>
              <div className={`w-8 h-8 rounded-full flex items-center justify-center border-2 ${
                step === 'upload' ? 'border-[#6d41ff] bg-[#6d41ff]/20' : uploadResponse ? 'border-green-500 bg-green-500/20' : 'border-gray-600'
              }`}>
                1
              </div>
              <span className="text-sm">Upload Documents</span>
            </div>
            <div className="w-16 h-px bg-gray-600"></div>
            <div className={`flex items-center space-x-2 ${step === 'details' ? 'text-[#6d41ff]' : 'text-gray-600'}`}>
              <div className={`w-8 h-8 rounded-full flex items-center justify-center border-2 ${
                step === 'details' ? 'border-[#6d41ff] bg-[#6d41ff]/20' : 'border-gray-600'
              }`}>
                2
              </div>
              <span className="text-sm">Property Details</span>
            </div>
          </div>

          {/* Error Display */}
          {error && (
            <div className="text-red-500 bg-red-500/10 p-4 rounded-xl border border-red-500/20">
              {error}
            </div>
          )}

          {/* Step 1: File Upload */}
          {step === 'upload' && (
            <div className="bg-[#1A1A1A] border border-[#262626] rounded-xl p-8">
              <div className="space-y-6">
                <div>
                  <h2 className="text-2xl font-semibold mb-2">Upload Property Documents</h2>
                  <p className="text-gray-400">
                    Upload property documents, legal agreements, or images to IPFS for permanent storage.
                  </p>
                </div>

                <div className="space-y-4">
                  <div>
                    <label className="block text-sm font-medium mb-2">Select File</label>
                    <input
                      type="file"
                      accept="image/*,.pdf,.doc,.docx"
                      onChange={handleFileSelect}
                      className="w-full p-3 rounded-lg border border-[#262626] bg-[#0f0f0f] text-white file:mr-4 file:py-2 file:px-4 file:rounded-lg file:border-0 file:bg-[#6d41ff] file:text-white file:cursor-pointer hover:file:bg-[#5b2fff]"
                    />
                    {selectedFile && (
                      <p className="text-sm text-gray-400 mt-2">
                        Selected: {selectedFile.name} ({(selectedFile.size / 1024 / 1024).toFixed(2)} MB)
                      </p>
                    )}
                  </div>

                  <Button
                    variant="primary"
                    size="lg"
                    onClick={handleUpload}
                    disabled={!selectedFile || loading}
                    className="w-full bg-[#6d41ff] hover:bg-[#5b2fff] text-white disabled:opacity-50"
                  >
                    {loading ? "Uploading..." : "Upload to IPFS"}
                  </Button>
                </div>
              </div>
            </div>
          )}

          {/* Step 2: Property Details */}
          {step === 'details' && uploadResponse && (
            <div className="bg-[#1A1A1A] border border-[#262626] rounded-xl p-8">
              <div className="space-y-6">
                <div>
                  <h2 className="text-2xl font-semibold mb-2">Property Details</h2>
                  <p className="text-gray-400">
                    Configure the property tokenization parameters.
                  </p>
                </div>

                {/* IPFS Upload Summary */}
                <div className="bg-green-500/10 border border-green-500/20 rounded-lg p-4">
                  <h3 className="text-green-400 font-semibold mb-2">✅ File Uploaded Successfully</h3>
                  <p className="text-sm text-gray-400 break-all">
                    IPFS Hash: {uploadResponse.ipfs_hash}
                  </p>
                  <a
                    href={uploadResponse.url}
                    target="_blank"
                    rel="noopener noreferrer"
                    className="text-[#6d41ff] hover:text-[#5b2fff] text-sm underline"
                  >
                    View on IPFS Gateway
                  </a>
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
                    onClick={() => setStep('upload')}
                    className="flex-1 border-[#262626] bg-[#0f0f0f] text-white hover:bg-[#262626]"
                  >
                    Back
                  </Button>
                  <Button
                    variant="primary"
                    size="lg"
                    onClick={handleCreateProperty}
                    disabled={loading || !name || !valuation || !tokenSupply}
                    className="flex-1 bg-[#6d41ff] hover:bg-[#5b2fff] text-white disabled:opacity-50"
                  >
                    {loading ? "Creating Property..." : "Create Property"}
                  </Button>
                </div>
              </div>
            </div>
          )}
        </div>
      </main>
    </div>
  );
}

export default CreatePropertyPage;

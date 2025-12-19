import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { Upload, X, FileText, CheckCircle } from "lucide-react";
import { Button } from "../Components/atoms/Button";
import { api } from "../lib/api";

function UploadProperty() {
  const navigate = useNavigate();
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");
  const [success, setSuccess] = useState(false);

  // Form state
  const [name, setName] = useState("");
  const [valuation, setValuation] = useState("");
  const [files, setFiles] = useState<File[]>([]);
  const [uploading, setUploading] = useState(false);

  // Computed values
  const symbol = "ETH";
  const tokenSupply = valuation ? Math.floor(parseFloat(valuation) / 10) : 0;

  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const selectedFiles = Array.from(e.target.files || []);
    if (selectedFiles.length === 0) return;

    setFiles(selectedFiles);
    setError("");
  };

  const removeFile = (index: number) => {
    setFiles(files.filter((_, i) => i !== index));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError("");
    setLoading(true);

    try {
      // Get user wallet from localStorage (reuse existing auth data)
      const userStr = localStorage.getItem("user");
      if (!userStr) {
        throw new Error("User not found. Please log in again.");
      }

      if (files.length === 0) {
        throw new Error("Please upload at least one document");
      }

      if (tokenSupply < 1) {
        throw new Error("Valuation must be at least 10 ETH to create tokens");
      }

      setUploading(true);
      await api.createPropertyUploadRequest(
        {
          name,
          symbol,
          valuation: parseFloat(valuation),
          token_supply: tokenSupply,
        },
        files
      );

      setSuccess(true);
    } catch (err: any) {
      setError(err.message || "Failed to submit property upload request");
    } finally {
      setLoading(false);
      setUploading(false);
    }
  };

  if (success) {
    return (
      <div className="max-w-xl mx-auto space-y-6">
        <div className="rounded-3xl border border-[#1f1f1f] bg-[#111111] p-8 text-center space-y-4">
          <div className="flex justify-center">
            <CheckCircle className="text-green-500" size={64} />
          </div>
          <h2 className="text-2xl font-semibold">Request Submitted!</h2>
          <p className="text-gray-400">
            Your property upload request has been submitted successfully. An admin will review it shortly.
          </p>
          <p className="text-sm text-gray-500">
            You can check the status of your request in the Messages section.
          </p>
          <div className="flex gap-3 justify-center pt-4">
            <Button
              variant="outline"
              size="md"
              className="border-[#262626] text-white hover:bg-[#1a1a1a]"
              onClick={() => navigate("/dashboard/messages")}
            >
              View Messages
            </Button>
            <Button
              variant="primary"
              size="md"
              className="bg-[#6d41ff] hover:bg-[#5b2fff] text-white"
              onClick={() => {
                setSuccess(false);
                setName("");
                setValuation("");
                setFiles([]);
              }}
            >
              Submit Another
            </Button>
          </div>
        </div>
      </div>
    );
  }

  return (
    <div className="max-w-2xl mx-auto space-y-6">
      <header>
        <h1 className="text-3xl font-semibold">Upload Property</h1>
        <p className="text-sm text-gray-400">
          Submit a property upload request for admin approval
        </p>
      </header>

      {error && (
        <div className="text-red-500 bg-red-500/10 p-4 rounded-xl border border-red-500/20">
          {error}
        </div>
      )}

      <form onSubmit={handleSubmit} className="space-y-6">
        {/* Property Details */}
        <div className="rounded-3xl border border-[#1f1f1f] bg-[#111111] p-6 space-y-4">
          <h2 className="font-semibold text-lg">Property Details</h2>

          <div className="space-y-2">
            <label className="text-sm text-gray-400">Property Name</label>
            <input
              type="text"
              value={name}
              onChange={(e) => setName(e.target.value)}
              placeholder="e.g., Luxury Villa Miami"
              required
              className="w-full px-4 py-3 rounded-xl border border-[#1f1f1f] bg-[#0f0f0f] text-white placeholder-gray-500 focus:outline-none focus:border-[#6d41ff]"
            />
          </div>

          <div className="space-y-2">
            <label className="text-sm text-gray-400">Valuation (ETH)</label>
            <input
              type="number"
              value={valuation}
              onChange={(e) => setValuation(e.target.value)}
              placeholder="e.g., 100"
              step="0.01"
              min="0"
              required
              className="w-full px-4 py-3 rounded-xl border border-[#1f1f1f] bg-[#0f0f0f] text-white placeholder-gray-500 focus:outline-none focus:border-[#6d41ff]"
            />
          </div>

          <div className="grid md:grid-cols-2 gap-4">
            <div className="space-y-2">
              <label className="text-sm text-gray-400">Token Symbol</label>
              <div className="w-full px-4 py-3 rounded-xl border border-[#1f1f1f] bg-[#1a1a1a] text-gray-300">
                {symbol}
              </div>
              <p className="text-xs text-gray-500">Fixed token symbol</p>
            </div>

            <div className="space-y-2">
              <label className="text-sm text-gray-400">Token Supply</label>
              <div className="w-full px-4 py-3 rounded-xl border border-[#1f1f1f] bg-[#1a1a1a] text-gray-300">
                {tokenSupply || "—"}
              </div>
              <p className="text-xs text-gray-500">Calculated as valuation ÷ 10</p>
            </div>
          </div>
        </div>

        {/* Document Upload */}
        <div className="rounded-3xl border border-[#1f1f1f] bg-[#111111] p-6 space-y-4">
          <h2 className="font-semibold text-lg">Property Documents</h2>
          <p className="text-sm text-gray-400">
            Upload property documentation (deed, certificate, images, etc.). Multiple files are supported.
          </p>

          <label className="flex flex-col items-center justify-center h-40 border-2 border-dashed border-[#1f1f1f] rounded-2xl cursor-pointer hover:border-[#6d41ff] transition-colors">
            <Upload className="text-gray-400 mb-2" size={32} />
            <span className="text-gray-400">Click to upload files</span>
            <span className="text-xs text-gray-500 mt-1">Multiple files supported</span>
            <input
              type="file"
              onChange={handleFileChange}
              multiple
              className="hidden"
            />
          </label>

          {files.length > 0 && (
            <div className="space-y-2">
              <p className="text-sm text-gray-400">Selected files ({files.length}):</p>
              {files.map((file, index) => (
                <div key={index} className="flex items-center justify-between p-3 rounded-xl border border-[#1f1f1f] bg-[#0f0f0f]">
                  <div className="flex items-center gap-3">
                    <FileText className="text-[#6d41ff]" size={20} />
                    <div>
                      <p className="text-sm font-medium truncate max-w-[200px]">
                        {file.name}
                      </p>
                      <p className="text-xs text-gray-500">
                        {(file.size / 1024).toFixed(2)} KB
                      </p>
                    </div>
                  </div>
                  <button
                    type="button"
                    onClick={() => removeFile(index)}
                    className="p-2 rounded-lg hover:bg-[#1a1a1a] text-gray-400 hover:text-white transition-colors"
                  >
                    <X size={18} />
                  </button>
                </div>
              ))}
            </div>
          )}
        </div>

        {/* Submit */}
        <div className="flex gap-3">
          <Button
            type="button"
            variant="outline"
            size="md"
            className="border-[#262626] text-white hover:bg-[#1a1a1a]"
            onClick={() => navigate("/dashboard")}
          >
            Cancel
          </Button>
          <Button
            type="submit"
            variant="primary"
            size="md"
            className="bg-[#6d41ff] hover:bg-[#5b2fff] text-white flex-1"
            disabled={loading || uploading || files.length === 0}
          >
            {loading ? "Submitting Request..." : "Submit Request"}
          </Button>
        </div>
      </form>
    </div>
  );
}

export default UploadProperty;




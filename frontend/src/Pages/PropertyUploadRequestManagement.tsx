import { useEffect, useState } from "react";
import { Button } from "../Components/atoms/Button";
import { api } from "../lib/api";
import { CheckCircle, XCircle, Clock, RefreshCw } from "lucide-react";

interface PropertyUploadRequest {
  id: string;
  wallet_address: string;
  name: string;
  symbol: string;
  valuation: number;
  token_supply: number;
  metadata_hash: string;
  status: "pending" | "approved" | "rejected";
  rejection_reason?: string;
  created_at: string;
  documents?: Array<{
    id: string;
    file_url: string;
    file_hash: string;
    name: string;
    type: string;
  }>;
}

function PropertyUploadRequestManagement() {
  const [requests, setRequests] = useState<PropertyUploadRequest[]>([]);
  const [filteredRequests, setFilteredRequests] = useState<PropertyUploadRequest[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");
  const [filter, setFilter] = useState<'all' | 'pending' | 'approved' | 'rejected'>('all');
  const [approving, setApproving] = useState<string | null>(null);
  const [rejecting, setRejecting] = useState<string | null>(null);
  const [rejectReason, setRejectReason] = useState("");
  const [showRejectModal, setShowRejectModal] = useState<string | null>(null);

  useEffect(() => {
    loadRequests();
  }, []);

  useEffect(() => {
    let filtered = requests;

    switch (filter) {
      case 'pending':
        filtered = requests.filter(req => req.status === 'pending');
        break;
      case 'approved':
        filtered = requests.filter(req => req.status === 'approved');
        break;
      case 'rejected':
        filtered = requests.filter(req => req.status === 'rejected');
        break;
      case 'all':
      default:
        filtered = requests;
        break;
    }

    setFilteredRequests(filtered);
  }, [requests, filter]);

  const loadRequests = async () => {
    try {
      setLoading(true);
      setError("");
      const data = await api.getPropertyUploadRequests();
      setRequests(data);
    } catch (err: any) {
      setError(err.message || "Failed to load property upload requests");
    } finally {
      setLoading(false);
    }
  };

  const handleApprove = async (id: string) => {
    try {
      setApproving(id);
      setError("");
      await api.approvePropertyUploadRequest(id);
      await loadRequests();
    } catch (err: any) {
      setError(err.message || "Failed to approve request");
    } finally {
      setApproving(null);
    }
  };

  const handleReject = async (id: string) => {
    if (!rejectReason.trim()) {
      setError("Please provide a rejection reason");
      return;
    }

    try {
      setRejecting(id);
      setError("");
      await api.rejectPropertyUploadRequest(id, rejectReason);
      setShowRejectModal(null);
      setRejectReason("");
      await loadRequests();
    } catch (err: any) {
      setError(err.message || "Failed to reject request");
    } finally {
      setRejecting(null);
    }
  };

  const getStatusBadge = (status: string) => {
    switch (status) {
      case 'pending':
        return (
          <span className="px-3 py-1 rounded-full text-xs font-medium bg-yellow-500/20 text-yellow-400 border border-yellow-500/30 flex items-center gap-1">
            <Clock size={12} />
            Pending
          </span>
        );
      case 'approved':
        return (
          <span className="px-3 py-1 rounded-full text-xs font-medium bg-green-500/20 text-green-400 border border-green-500/30 flex items-center gap-1">
            <CheckCircle size={12} />
            Approved
          </span>
        );
      case 'rejected':
        return (
          <span className="px-3 py-1 rounded-full text-xs font-medium bg-red-500/20 text-red-400 border border-red-500/30 flex items-center gap-1">
            <XCircle size={12} />
            Rejected
          </span>
        );
      default:
        return null;
    }
  };

  const formatDate = (dateString: string) => {
    return new Date(dateString).toLocaleDateString('en-US', {
      year: 'numeric',
      month: 'short',
      day: 'numeric',
      hour: '2-digit',
      minute: '2-digit',
    });
  };

  if (loading && requests.length === 0) {
    return (
      <div className="min-h-screen bg-[#0f0f0f] text-white flex items-center justify-center">
        <p className="text-gray-400">Loading property upload requests...</p>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-[#0f0f0f] text-white">
      <main className="max-w-6xl mx-auto px-6 pt-8 pb-16">
        <div className="space-y-8">
          {/* Header */}
          <div className="flex items-center justify-between">
            <div>
              <h1 className="text-3xl font-semibold">Property Upload Requests</h1>
              <p className="text-gray-400 mt-2">
                Review and approve or reject property upload requests from users.
              </p>
            </div>
            <Button
              variant="outline"
              size="md"
              onClick={loadRequests}
              className="border-[#262626] bg-[#1A1A1A] text-white hover:bg-[#262626]"
            >
              <RefreshCw size={16} className="mr-2" />
              Refresh
            </Button>
          </div>

          {/* Filter Section */}
          <div className="flex items-center justify-between bg-[#1A1A1A] border border-[#262626] rounded-xl p-4">
            <div className="flex items-center gap-4">
              <span className="text-sm font-medium text-gray-300">Filter by status:</span>
              <div className="flex gap-2">
                <button
                  onClick={() => setFilter('all')}
                  className={
                    filter === 'all'
                      ? 'px-3 py-1 rounded-lg text-sm font-medium transition-colors bg-[#6d41ff] text-white'
                      : 'px-3 py-1 rounded-lg text-sm font-medium transition-colors bg-[#262626] text-gray-300 hover:bg-[#333333]'
                  }
                >
                  All ({requests.length})
                </button>
                <button
                  onClick={() => setFilter('pending')}
                  className={
                    filter === 'pending'
                      ? 'px-3 py-1 rounded-lg text-sm font-medium transition-colors bg-yellow-600 text-white'
                      : 'px-3 py-1 rounded-lg text-sm font-medium transition-colors bg-[#262626] text-gray-300 hover:bg-[#333333]'
                  }
                >
                  Pending ({requests.filter(r => r.status === 'pending').length})
                </button>
                <button
                  onClick={() => setFilter('approved')}
                  className={
                    filter === 'approved'
                      ? 'px-3 py-1 rounded-lg text-sm font-medium transition-colors bg-green-600 text-white'
                      : 'px-3 py-1 rounded-lg text-sm font-medium transition-colors bg-[#262626] text-gray-300 hover:bg-[#333333]'
                  }
                >
                  Approved ({requests.filter(r => r.status === 'approved').length})
                </button>
                <button
                  onClick={() => setFilter('rejected')}
                  className={
                    filter === 'rejected'
                      ? 'px-3 py-1 rounded-lg text-sm font-medium transition-colors bg-red-600 text-white'
                      : 'px-3 py-1 rounded-lg text-sm font-medium transition-colors bg-[#262626] text-gray-300 hover:bg-[#333333]'
                  }
                >
                  Rejected ({requests.filter(r => r.status === 'rejected').length})
                </button>
              </div>
            </div>
            <div className="text-sm text-gray-400">
              Showing {filteredRequests.length} of {requests.length} requests
            </div>
          </div>

          {/* Error Display */}
          {error && (
            <div className="text-red-500 bg-red-500/10 p-4 rounded-xl border border-red-500/20">
              {error}
            </div>
          )}

          {/* Requests List */}
          <div className="space-y-4">
            {filteredRequests.length === 0 ? (
              <div className="bg-[#1A1A1A] border border-[#262626] rounded-xl p-12 text-center">
                <p className="text-gray-400">No property upload requests found.</p>
              </div>
            ) : (
              filteredRequests.map((request) => (
                <div
                  key={request.id}
                  className="bg-[#1A1A1A] border border-[#262626] rounded-xl p-6 space-y-4"
                >
                  <div className="flex items-start justify-between">
                    <div className="flex-1">
                      <div className="flex items-center gap-3 mb-2">
                        <h3 className="text-xl font-semibold">{request.name}</h3>
                        {getStatusBadge(request.status)}
                      </div>
                      <div className="grid md:grid-cols-2 gap-4 mt-4 text-sm">
                        <div>
                          <span className="text-gray-400">Wallet Address:</span>
                          <p className="text-gray-300 font-mono text-xs mt-1">{request.wallet_address}</p>
                        </div>
                        <div>
                          <span className="text-gray-400">Symbol:</span>
                          <p className="text-gray-300 mt-1">{request.symbol}</p>
                        </div>
                        <div>
                          <span className="text-gray-400">Valuation:</span>
                          <p className="text-gray-300 mt-1">{request.valuation} ETH</p>
                        </div>
                        <div>
                          <span className="text-gray-400">Token Supply:</span>
                          <p className="text-gray-300 mt-1">{request.token_supply.toLocaleString()}</p>
                        </div>
                        <div>
                          <span className="text-gray-400">Submitted:</span>
                          <p className="text-gray-300 mt-1">{formatDate(request.created_at)}</p>
                        </div>
                        <div>
                          <span className="text-gray-400">IPFS Hash:</span>
                          <p className="text-gray-300 font-mono text-xs mt-1 break-all">{request.metadata_hash}</p>
                        </div>
                      </div>
                      {request.rejection_reason && (
                        <div className="mt-4 p-3 bg-red-500/10 border border-red-500/20 rounded-lg">
                          <span className="text-sm font-medium text-red-400">Rejection Reason:</span>
                          <p className="text-sm text-gray-300 mt-1">{request.rejection_reason}</p>
                        </div>
                      )}
                    </div>
                  </div>

                  {/* Actions */}
                  {request.status === 'pending' && (
                    <div className="flex gap-3 pt-4 border-t border-[#262626]">
                      <Button
                        variant="primary"
                        size="md"
                        onClick={() => handleApprove(request.id)}
                        disabled={approving === request.id}
                        className="bg-green-600 hover:bg-green-700 text-white"
                      >
                        {approving === request.id ? "Approving..." : "Approve"}
                      </Button>
                      <Button
                        variant="outline"
                        size="md"
                        onClick={() => setShowRejectModal(request.id)}
                        disabled={rejecting === request.id}
                        className="border-red-500 text-red-400 hover:bg-red-500/10"
                      >
                        Reject
                      </Button>
                    </div>
                  )}
                </div>
              ))
            )}
          </div>
        </div>
      </main>

      {/* Reject Modal */}
      {showRejectModal && (
        <div className="fixed inset-0 bg-black/50 flex items-center justify-center z-50 p-4">
          <div className="bg-[#1A1A1A] border border-[#262626] rounded-xl p-6 max-w-md w-full space-y-4">
            <h3 className="text-xl font-semibold">Reject Property Upload Request</h3>
            <p className="text-gray-400 text-sm">
              Please provide a reason for rejecting this request. The user will be notified.
            </p>
            <textarea
              value={rejectReason}
              onChange={(e) => setRejectReason(e.target.value)}
              placeholder="Enter rejection reason..."
              className="w-full px-4 py-3 rounded-xl border border-[#1f1f1f] bg-[#0f0f0f] text-white placeholder-gray-500 focus:outline-none focus:border-[#6d41ff] min-h-[100px]"
            />
            <div className="flex gap-3">
              <Button
                variant="outline"
                size="md"
                onClick={() => {
                  setShowRejectModal(null);
                  setRejectReason("");
                }}
                className="border-[#262626] text-white hover:bg-[#1a1a1a] flex-1"
              >
                Cancel
              </Button>
              <Button
                variant="primary"
                size="md"
                onClick={() => handleReject(showRejectModal)}
                disabled={!rejectReason.trim() || rejecting === showRejectModal}
                className="bg-red-600 hover:bg-red-700 text-white flex-1"
              >
                {rejecting === showRejectModal ? "Rejecting..." : "Reject"}
              </Button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
}

export default PropertyUploadRequestManagement;

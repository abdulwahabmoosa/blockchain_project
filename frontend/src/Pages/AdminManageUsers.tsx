import { useEffect, useState, useRef } from "react";
import { RefreshCw, Search, CheckCircle, XCircle, AlertTriangle } from "lucide-react";
import { Button } from "../Components/atoms/Button";
import { api } from "../lib/api";
import { checkApprovalStatus } from "../lib/contracts";
import { ethers } from "ethers";
import { useWallet } from "../hooks/useWallet";
import { getAdminWalletState } from "../lib/wallet";

interface UserData {
  ID: string;
  Email: string;
  Name: string;
  WalletAddress: string;
  Role: string;
  IsApproved: boolean;
}

interface ApiUser {
  ID: string;
  Email: string;
  Name: string;
  WalletAddress: string;
  Role: string;
  IsApproved: boolean;
}

interface UserWithStatus extends UserData {
  blockchainApproved: boolean | null;
  approving: boolean;
  rejecting: boolean;
  pendingConfirmation: boolean;
}

function AdminManageUsers() {
  const [users, setUsers] = useState<UserWithStatus[]>([]);
  const [filteredUsers, setFilteredUsers] = useState<UserWithStatus[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");
  const [filter, setFilter] = useState<"all" | "approved" | "not-approved">("all");
  const [searchQuery, setSearchQuery] = useState("");
  const [isCheckingBlockchain, setIsCheckingBlockchain] = useState(false);
  const { provider, isConnected } = useWallet();
  const loadUsersRunning = useRef(false);

  // Load pending approvals from localStorage
  const getPendingApprovals = () => {
    try {
      const pending = localStorage.getItem("pendingApprovals");
      return pending ? JSON.parse(pending) : {};
    } catch {
      return {};
    }
  };

  // Save pending approvals to localStorage
  const setPendingApprovals = (pending: Record<string, boolean>) => {
    localStorage.setItem("pendingApprovals", JSON.stringify(pending));
  };

  // Load persisted approval status from localStorage
  const getPersistedApprovals = (): Record<string, boolean> => {
    try {
      const persisted = localStorage.getItem("userApprovals");
      return persisted ? JSON.parse(persisted) : {};
    } catch {
      return {};
    }
  };

  // Save approval status to localStorage
  const setPersistedApprovals = (approvals: Record<string, boolean>) => {
    localStorage.setItem("userApprovals", JSON.stringify(approvals));
  };

  // Clear all persisted approvals (useful for debugging)
  const clearPersistedApprovals = () => {
    console.log("🗑️ Clearing persisted approval data...");
    console.log("📚 Before clearing - userApprovals:", localStorage.getItem("userApprovals"));
    console.log("⏳ Before clearing - pendingApprovals:", localStorage.getItem("pendingApprovals"));
    localStorage.removeItem("userApprovals");
    localStorage.removeItem("pendingApprovals");
    console.log("✅ Cleared all persisted approval data");
    console.log("📚 After clearing - userApprovals:", localStorage.getItem("userApprovals"));
    console.log("⏳ After clearing - pendingApprovals:", localStorage.getItem("pendingApprovals"));
  };

  useEffect(() => {
    loadUsers();
  }, []);

  // Filter users based on approval status
  useEffect(() => {
    let filtered = users;

    switch (filter) {
      case 'approved':
        filtered = users.filter(user => user.blockchainApproved === true);
        break;
      case 'not-approved':
        filtered = users.filter(user => user.blockchainApproved === false);
        break;
      case 'all':
      default:
        filtered = users;
        break;
    }

    if (searchQuery) {
      filtered = filtered.filter((user) =>
        user.Name.toLowerCase().includes(searchQuery.toLowerCase()) ||
        user.Email.toLowerCase().includes(searchQuery.toLowerCase()) ||
        (user.WalletAddress &&
          user.WalletAddress.toLowerCase().includes(searchQuery.toLowerCase()))
      );
    }

    setFilteredUsers(filtered);
  }, [users, filter, searchQuery]);

  // Re-check blockchain status when wallet connects (only if we have users loaded and need updates)
  useEffect(() => {
    if (isConnected && provider && users.length > 0 && !loadUsersRunning.current) {
      // Only re-check if we have users that haven't been checked yet (blockchainApproved is null)
      // or if we're using a different provider than before
      const needsBlockchainCheck = users.some(user =>
        user.blockchainApproved === null && user.WalletAddress && user.WalletAddress.trim() !== ""
      );

      if (needsBlockchainCheck) {
        const usersNeedingCheck = users.filter(user =>
          user.blockchainApproved === null && user.WalletAddress && user.WalletAddress.trim() !== ""
        );
        console.log(`🔄 Wallet connected, checking blockchain status for ${usersNeedingCheck.length} unchecked users...`);
        checkBlockchainStatus(usersNeedingCheck).then(checkedUsers => {
          console.log('✅ Blockchain status updated after wallet connection');
          // Update only the users that were checked
          setUsers(prevUsers => prevUsers.map(user => {
            const checkedUser = checkedUsers.find(cu => cu.ID === user.ID);
            return checkedUser || user;
          }));
        }).catch(err => {
          console.error('❌ Failed to update blockchain status:', err);
        });
      } else {
        console.log('ℹ️ All users have definitive blockchain status, skipping re-check');
      }
    }
  }, [isConnected, provider, users.length]);

  const checkBlockchainStatus = async (currentUsers: UserWithStatus[]) => {
    console.log('🔄 Starting blockchain status check...');
    console.log('🔗 Wallet connected:', isConnected, 'Provider available:', !!provider);
    console.log('👥 Users to check:', currentUsers.map(u => ({ email: u.Email, address: u.WalletAddress })));
    setIsCheckingBlockchain(true);

    let blockchainProvider: ethers.BrowserProvider | ethers.AbstractProvider | null = provider;

    // If MetaMask is not connected, use admin wallet for blockchain operations
    if (!isConnected || !provider) {
      console.log('ℹ️ MetaMask not connected, using admin wallet for blockchain checks...');
      try {
        const adminWallet = await getAdminWalletState();
        blockchainProvider = adminWallet.provider;
        console.log('✅ Admin wallet connected for blockchain operations');
      } catch (error) {
        console.error('❌ Failed to connect admin wallet:', error);
        setIsCheckingBlockchain(false);
        return currentUsers;
      }
    }

    if (currentUsers.length === 0) {
      console.log('ℹ️ No users to check');
      setIsCheckingBlockchain(false);
      return currentUsers;
    }

    try {
      const pendingApprovals = getPendingApprovals();
      console.log('📋 Pending approvals:', pendingApprovals);

      // Ensure we have a valid provider
      if (!blockchainProvider) {
        console.error('❌ No blockchain provider available');
        setIsCheckingBlockchain(false);
        return currentUsers;
      }

      // Check users in parallel for faster results (with concurrency limit)
      const checkUserStatus = async (user: UserWithStatus): Promise<UserWithStatus> => {
        // Skip users with empty wallet addresses
        if (!user.WalletAddress || user.WalletAddress.trim() === "") {
          console.log(`⏭️ Skipping ${user.Email}: No wallet address set`);
          return { ...user, blockchainApproved: null, pendingConfirmation: false };
        }

        try {
          console.log(`🔍 Checking approval status for ${user.Email} (${user.WalletAddress})`);

          // Add timeout to blockchain check (reduced from 3s to 2s)
          const checkPromise = checkApprovalStatus(user.WalletAddress, blockchainProvider!);
          const timeoutPromise = new Promise<boolean>((_, reject) =>
            setTimeout(() => reject(new Error('Blockchain check timeout')), 2000)
          );

          const blockchainApproved = await Promise.race([checkPromise, timeoutPromise]);
          console.log(`✅ Approval status for ${user.Email}: ${blockchainApproved}`);

          // CRITICAL: Once approved on blockchain, status should never revert to false
          const wasPreviouslyApproved = user.blockchainApproved === true;

          if (wasPreviouslyApproved && !blockchainApproved) {
            console.warn(`⚠️ WARNING: User ${user.Email} was previously approved but blockchain check returned false. Keeping approved status.`);
            return { ...user, blockchainApproved: true, pendingConfirmation: false };
          } else if (pendingApprovals[user.WalletAddress] && !blockchainApproved) {
            console.log(`⏳ Keeping ${user.Email} in pending state`);
            return { ...user, blockchainApproved: null, pendingConfirmation: true };
          } else {
            // Update persisted approvals when status is confirmed
            if (blockchainApproved === true && user.WalletAddress) {
              const persistedApprovals = getPersistedApprovals();
              persistedApprovals[user.WalletAddress] = true;
              setPersistedApprovals(persistedApprovals);
            }
            return { ...user, blockchainApproved, pendingConfirmation: false };
          }
        } catch (err) {
          console.error(`❌ Failed to check approval for ${user.Email}:`, err);
          // If there's a pending approval, show as pending even if check fails
          if (pendingApprovals[user.WalletAddress]) {
            return { ...user, blockchainApproved: null, pendingConfirmation: true };
          }
          return user;
        }
      };

      // Process users in parallel batches (max 5 concurrent checks)
      const batchSize = 5;
      const results: UserWithStatus[] = [];
      
      for (let i = 0; i < currentUsers.length; i += batchSize) {
        const batch = currentUsers.slice(i, i + batchSize);
        const batchResults = await Promise.all(batch.map(checkUserStatus));
        results.push(...batchResults);
        
        // Update UI incrementally as batches complete
        setUsers(prevUsers => {
          const updated = [...prevUsers];
          batchResults.forEach(checkedUser => {
            const index = updated.findIndex(u => u.ID === checkedUser.ID);
            if (index !== -1) {
              updated[index] = checkedUser;
            }
          });
          return updated;
        });
      }

      console.log('✅ Blockchain status check completed');
      return results;

    } catch (err) {
      console.error('❌ Blockchain status check failed:', err);
      // Return users as-is if blockchain check fails completely
      return currentUsers;
    } finally {
      setIsCheckingBlockchain(false);
    }
  };

  const loadUsers = async () => {
    // Prevent multiple simultaneous calls
    if (loadUsersRunning.current) {
      console.log('⚠️ loadUsers already running, skipping');
      return;
    }

    console.log('🔄 Starting loadUsers...');
    loadUsersRunning.current = true;
    setLoading(true);
    setError("");

    try {
      // Clear persisted approvals to ensure fresh blockchain checks
      clearPersistedApprovals();

      // Fetch real users from the backend API
      const usersResponse = await api.getUsers();

      // Convert backend user format to frontend format
      const sampleUsers: ApiUser[] = usersResponse.map((user: any) => ({
        ID: user.ID,
        Email: user.Email,
        Name: user.Name,
        WalletAddress: user.WalletAddress,
        Role: user.Role,
        IsApproved: user.IsApproved
      }));

      console.log('📋 Created sample users:', sampleUsers.length);

      // Create users with initial state - always start with null blockchain status
      // This ensures we always check the blockchain for the most current approval status
      const usersWithInitialState: UserWithStatus[] = sampleUsers.map(user => {
        // Try to find existing user with the same ID to preserve UI state only
        const existingUser = users.find(u => u.ID === user.ID);

        return {
          ...user,
          blockchainApproved: null, // Always start with null to force blockchain check
          approving: false, // Reset approving state
          rejecting: false, // Reset rejecting state
          pendingConfirmation: existingUser?.pendingConfirmation ?? false
        };
      });

      console.log('🔗 Checking blockchain status...');

      console.log('📋 Setting initial users without blockchain status');
      setUsers(usersWithInitialState);

      // Check blockchain status asynchronously - use admin wallet if MetaMask not connected
      const usersNeedingCheck = usersWithInitialState.filter(user =>
        user.WalletAddress &&
        user.WalletAddress.trim() !== "" &&
        user.blockchainApproved === null
      );

      if (usersNeedingCheck.length > 0) {
        console.log(`🌐 Starting async blockchain status check for ${usersNeedingCheck.length} users...`);
        // Start blockchain check immediately (will use admin wallet if MetaMask not connected)
        checkBlockchainStatus(usersNeedingCheck).then(checkedUsers => {
          console.log('✅ Async blockchain check completed, updating users');
          // Update the main users list with the checked results
          setUsers(prevUsers => prevUsers.map(user => {
            const checkedUser = checkedUsers.find(cu => cu.ID === user.ID);
            return checkedUser || user;
          }));
        }).catch(err => {
          console.error('❌ Async blockchain check failed:', err);
          // Keep users as-is if blockchain check fails
        });
      } else {
        console.log('ℹ️ No users need blockchain status check');
      }

      // Force loading to complete after a maximum time
      setTimeout(() => {
        if (loading) {
          console.log('⏰ Force completing loading state');
          loadUsersRunning.current = false;
          setLoading(false);
        }
      }, 30000); // 30 second absolute timeout
    } catch (err: any) {
      console.error('❌ loadUsers failed:', err);
      setError(err.message || "Failed to load users");

      // Still show users even if blockchain check fails - use empty array as fallback
      const fallbackUsers: UserWithStatus[] = [];
      setUsers(fallbackUsers);
    } finally {
      console.log('🏁 loadUsers completed');
      loadUsersRunning.current = false;
      setLoading(false);
    }
  };

  const handleRetryApproval = (user: UserWithStatus) => {
    // Clear pending state and allow retry
    const pendingApprovals = getPendingApprovals();
    delete pendingApprovals[user.WalletAddress];
    setPendingApprovals(pendingApprovals);

    setUsers(prev => prev.map(u =>
      u.ID === user.ID ? { ...u, pendingConfirmation: false } : u
    ));

    // Also clear any error messages
    setError("");
  };

  const handleApproveUser = async (user: UserWithStatus) => {
    // Check if user is authenticated
    const token = localStorage.getItem("token");
    if (!token) {
      alert("You are not logged in. Please log in as an admin first.");
      return;
    }

    // Check if user has a valid wallet address
    if (!user.WalletAddress || user.WalletAddress.trim() === "") {
      alert("Cannot approve user: No wallet address is set for this user.");
      return;
    }

    setUsers(prev => prev.map(u =>
      u.ID === user.ID ? { ...u, approving: true } : u
    ));

    try {
      console.log(`🔄 Calling backend API to approve user ${user.WalletAddress}`);
      console.log(`📋 User details: ${user.Email} (${user.WalletAddress})`);

      const response = await api.approveUser(user.WalletAddress);
      console.log('✅ Backend API response:', response);

      // Mark as pending in localStorage
      const pendingApprovals = getPendingApprovals();
      pendingApprovals[user.WalletAddress] = true;
      setPendingApprovals(pendingApprovals);

      // Set to pending confirmation state
      setUsers(prev => prev.map(u =>
        u.ID === user.ID ? { ...u, approving: false, pendingConfirmation: true } : u
      ));

      alert(`Approval transaction submitted!\nTransaction Hash: ${response.tx_hash}\n\nTransaction confirmed on blockchain!`);

      // Backend already confirmed the transaction, use the response
      console.log(`✅ Backend confirmed approval status: ${response.approved}`);

      // Remove from pending approvals since it's confirmed
      const pendingApprovals2 = getPendingApprovals();
      delete pendingApprovals2[user.WalletAddress];
      setPendingApprovals(pendingApprovals2);

      // Persist approval status if approved
      if (response.approved && user.WalletAddress) {
        const persistedApprovals = getPersistedApprovals();
        persistedApprovals[user.WalletAddress] = true;
        setPersistedApprovals(persistedApprovals);
        console.log(`💾 Persisted approval status for ${user.Email} after backend confirmation`);
      }

      setUsers(prev => prev.map(u =>
        u.ID === user.ID ? {
          ...u,
          blockchainApproved: response.approved,
          pendingConfirmation: false
        } : u
      ));

      alert(`✅ Approval ${response.approved ? 'successful' : 'completed but verification failed'}!`);

    } catch (err: any) {
      console.error("❌ Approval failed:", err);
      console.error("❌ Error details:", err.message);
      
      // Extract error message - handle both string errors and Error objects
      let errorMessage = err.message || "Failed to approve user";
      
      // Provide user-friendly error messages
      if (errorMessage.includes("approval contract not deployed")) {
        errorMessage = "Blockchain Error: Approval contract not deployed. Please deploy contracts to enable blockchain functionality.";
      } else if (errorMessage.includes("Blockchain Error")) {
        // Keep the blockchain error message as-is
        errorMessage = errorMessage;
      }
      
      setError(errorMessage);

      // Reset pending state on error
      const pendingApprovals = getPendingApprovals();
      delete pendingApprovals[user.WalletAddress];
      setPendingApprovals(pendingApprovals);

      setUsers(prev => prev.map(u =>
        u.ID === user.ID ? { ...u, approving: false, pendingConfirmation: false } : u
      ));

      alert(`❌ Approval failed: ${errorMessage}`);
    }
  };

  const handleRejectUser = async (user: UserWithStatus) => {
    // Check if user is authenticated
    const token = localStorage.getItem("token");
    if (!token) {
      alert("You are not logged in. Please log in as an admin first.");
      return;
    }

    // Check if user has a valid wallet address
    if (!user.WalletAddress || user.WalletAddress.trim() === "") {
      alert("Cannot reject user: No wallet address is set for this user.");
      return;
    }

    setUsers(prev => prev.map(u =>
      u.ID === user.ID ? { ...u, rejecting: true } : u
    ));

    try {
      console.log(`🔄 Calling backend API to reject user ${user.WalletAddress}`);
      console.log(`📋 User details: ${user.Email} (${user.WalletAddress})`);

      const response = await api.rejectUser(user.WalletAddress);
      console.log('✅ Backend API response:', response);

      // Update user status immediately since rejection is database-only
      setUsers(prev => prev.map(u =>
        u.ID === user.ID ? {
          ...u,
          blockchainApproved: false,
          rejecting: false,
          pendingConfirmation: false
        } : u
      ));

      alert(`✅ User rejection successful!`);

    } catch (err: any) {
      console.error("❌ Rejection failed:", err);
      console.error("❌ Error details:", err.message);
      setError(err.message || "Failed to reject user");

      setUsers(prev => prev.map(u =>
        u.ID === user.ID ? { ...u, rejecting: false } : u
      ));

      alert(`❌ Rejection failed: ${err.message}`);
    }
  };

  const getStatusBadge = (user: UserWithStatus) => {
    if (user.approving) return (
      <span className="flex items-center gap-1 text-xs px-2 py-1 rounded-full bg-blue-500/20 text-blue-400">
        <RefreshCw size={12} className="animate-spin" />
        Approving...
      </span>
    );
    if (user.rejecting) return (
      <span className="flex items-center gap-1 text-xs px-2 py-1 rounded-full bg-red-500/20 text-red-400">
        <RefreshCw size={12} className="animate-spin" />
        Rejecting...
      </span>
    );
    if (user.pendingConfirmation) return (
      <span className="flex items-center gap-1 text-xs px-2 py-1 rounded-full bg-yellow-500/20 text-yellow-400">
        ⏳ Pending Confirmation
      </span>
    );
    if (user.blockchainApproved === null) return (
      <span className="flex items-center gap-1 text-xs px-2 py-1 rounded-full bg-gray-500/20 text-gray-400">
        🔍 Checking blockchain...
      </span>
    );
    if (user.blockchainApproved) return (
      <span className="flex items-center gap-1 text-xs px-2 py-1 rounded-full bg-green-500/20 text-green-400">
        <CheckCircle size={12} />
        Approved
      </span>
    );
    return (
      <span className="flex items-center gap-1 text-xs px-2 py-1 rounded-full bg-red-500/20 text-red-400">
        <XCircle size={12} />
        Rejected
      </span>
    );
  };

  return (
    <div className="space-y-6">
      <header className="flex flex-col md:flex-row md:items-center justify-between gap-4">
        <div>
          <h1 className="text-3xl font-semibold">Manage Users</h1>
          <p className="text-sm text-gray-400">
            Approve users to enable platform participation
          </p>
        </div>
        <div className="flex gap-2">
          <Button
            variant="outline"
            size="md"
            onClick={async () => {
              // Clear ALL cached data aggressively
              console.log('🔄 Manual refresh triggered - clearing ALL cache');
              clearPersistedApprovals();

              // Also clear any cached approval status in memory
              setUsers(prevUsers => prevUsers.map(u => ({
                ...u,
                blockchainApproved: null,
                pendingConfirmation: false,
                approving: false,
                rejecting: false
              })));

              // Force reload all users
              await loadUsers();
            }}
            disabled={loading || isCheckingBlockchain}
            className="border-[#262626] text-white hover:bg-[#1a1a1a]"
          >
            <RefreshCw
              size={18}
              className={`mr-2 ${loading || isCheckingBlockchain ? "animate-spin" : ""}`}
            />
            {isCheckingBlockchain ? "Checking..." : "Refresh"}
          </Button>

          <Button
            variant="outline"
            size="md"
            onClick={() => {
              console.log('🧹 Clearing all cached approval data...');
              clearPersistedApprovals();
              alert('✅ Cleared all cached approval data. Click Refresh to reload fresh data.');
            }}
            className="border-red-500/50 bg-red-500/10 text-red-400 hover:bg-red-500/20"
          >
            Clear Cache
          </Button>
        </div>
      </header>

      {error && (
        <div className="text-red-500 bg-red-500/10 p-4 rounded-xl border border-red-500/20 flex items-center gap-2">
          <AlertTriangle size={18} />
          {error}
        </div>
      )}

      {/* Search and Filter */}
      <div className="flex flex-col md:flex-row gap-4">
        <div className="relative flex-1">
          <Search
            className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400"
            size={20}
          />
          <input
            type="text"
            placeholder="Search by name, email, or wallet..."
            value={searchQuery}
            onChange={(e) => setSearchQuery(e.target.value)}
            className="w-full pl-10 pr-4 py-3 rounded-xl border border-[#1f1f1f] bg-[#111111] text-white placeholder-gray-500 focus:outline-none focus:border-[#6d41ff]"
          />
        </div>
        <div className="flex gap-2">
          <button
            onClick={() => setFilter("all")}
            className={`px-4 py-2 rounded-xl text-sm font-medium transition-colors ${
              filter === "all"
                ? "bg-[#703BF7] text-white"
                : "bg-[#111111] border border-[#1f1f1f] text-gray-400 hover:text-white"
            }`}
          >
            All ({users.length})
          </button>
          <button
            onClick={() => setFilter("approved")}
            className={`px-4 py-2 rounded-xl text-sm font-medium transition-colors ${
              filter === "approved"
                ? "bg-green-600 text-white"
                : "bg-[#111111] border border-[#1f1f1f] text-gray-400 hover:text-white"
            }`}
          >
            Approved ({users.filter(u => u.blockchainApproved === true).length})
          </button>
          <button
            onClick={() => setFilter("not-approved")}
            className={`px-4 py-2 rounded-xl text-sm font-medium transition-colors ${
              filter === "not-approved"
                ? "bg-red-600 text-white"
                : "bg-[#111111] border border-[#1f1f1f] text-gray-400 hover:text-white"
            }`}
          >
            Not Approved ({users.filter(u => u.blockchainApproved === false).length})
          </button>
        </div>
      </div>

      {/* Users List */}
      <div className="rounded-3xl border border-[#1f1f1f] bg-[#111111] overflow-hidden">
        <div className="p-4 border-b border-[#1f1f1f] flex items-center justify-between">
          <h2 className="font-semibold">Registered Users</h2>
          <div className="flex items-center gap-4">
            {isCheckingBlockchain && (
              <span className="text-blue-400 text-xs flex items-center gap-1">
                🔄 Checking blockchain status...
              </span>
            )}
            <span className="text-xs text-gray-500">
              {filteredUsers.length} of {users.length} users
            </span>
          </div>
        </div>

        {loading ? (
          <div className="p-8 text-center">
            <RefreshCw
              size={24}
              className="animate-spin mx-auto text-gray-400 mb-2"
            />
            <p className="text-gray-400">Loading users...</p>
          </div>
        ) : filteredUsers.length === 0 ? (
          <div className="p-8 text-center text-gray-400">
            {searchQuery || filter !== "all"
              ? "No users match your search criteria"
              : "No users found"}
          </div>
        ) : (
          <div className="divide-y divide-[#1f1f1f]">
            {filteredUsers.map((user) => (
              <div
                key={user.ID}
                className="p-4 hover:bg-[#0f0f0f]/50 transition-colors"
              >
                <div className="flex flex-col md:flex-row md:items-center justify-between gap-4">
                  <div className="flex-1 min-w-0">
                    <div className="flex items-center gap-2 mb-1">
                      <h3 className="font-medium truncate">{user.Name}</h3>
                      <span
                        className={`text-xs px-2 py-0.5 rounded-full ${
                          user.Role === "admin"
                            ? "bg-purple-500/20 text-purple-400"
                            : "bg-blue-500/20 text-blue-400"
                        }`}
                      >
                        {user.Role}
                      </span>
                    </div>
                    <p className="text-sm text-gray-400 truncate">
                      {user.Email}
                    </p>
                    <p className="text-xs text-gray-500 font-mono mt-1">
                      {user.WalletAddress
                        ? `${user.WalletAddress.substring(
                            0,
                            10
                          )}...${user.WalletAddress.substring(38)}`
                        : "No wallet"}
                    </p>
                  </div>

                  <div className="flex items-center gap-3">
                    {getStatusBadge(user)}

                    {user.blockchainApproved === true ? (
                      <span className="text-xs text-green-400">Already Approved</span>
                    ) : user.blockchainApproved === false ? (
                      <span className="text-xs text-red-400">Rejected</span>
                    ) : user.pendingConfirmation ? (
                      <div className="flex gap-2 items-center">
                        <span className="text-yellow-400 text-sm">Transaction Pending...</span>
                        <Button
                          variant="outline"
                          size="sm"
                          onClick={() => handleRetryApproval(user)}
                          className="border-yellow-500/50 bg-yellow-500/10 text-yellow-400 hover:bg-yellow-500/20 text-xs"
                        >
                          Retry
                        </Button>
                      </div>
                    ) : user.WalletAddress && user.WalletAddress.trim() !== "" ? (
                      <div className="flex gap-2">
                        <Button
                          variant="primary"
                          size="sm"
                          onClick={() => handleApproveUser(user)}
                          disabled={user.approving || user.rejecting}
                          className="bg-[#6d41ff] hover:bg-[#5b2fff] text-white disabled:opacity-50"
                        >
                          {user.approving ? "Approving..." : "Approve"}
                        </Button>
                        <Button
                          variant="outline"
                          size="sm"
                          onClick={() => handleRejectUser(user)}
                          disabled={user.approving || user.rejecting}
                          className="border-red-500/50 bg-red-500/10 text-red-400 hover:bg-red-500/20 disabled:opacity-50"
                        >
                          {user.rejecting ? "Rejecting..." : "Reject"}
                        </Button>
                      </div>
                    ) : (
                      <span className="text-gray-500 text-sm">No wallet address</span>
                    )}
                  </div>
                </div>
              </div>
            ))}
          </div>
        )}
      </div>

      {/* Info Section */}
      <div className="rounded-2xl border border-blue-500/20 bg-blue-500/5 p-4">
        <h3 className="text-blue-400 font-medium mb-2 flex items-center gap-2">
          <AlertTriangle size={16} />
          How User Approval Works
        </h3>
        <ul className="text-sm text-gray-400 space-y-1">
          <li>Users must be approved to participate on the platform</li>
          <li>Each approval creates a transaction on the blockchain</li>
          <li>Transactions may take 15-30 seconds to confirm</li>
        </ul>
      </div>
    </div>
  );
}

export default AdminManageUsers;



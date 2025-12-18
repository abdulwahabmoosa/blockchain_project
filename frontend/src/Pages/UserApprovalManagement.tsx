import { useEffect, useState, useRef } from "react";
import { Navbar } from "../Components/organisms/Navbar";
import { Button } from "../Components/atoms/Button";
import { api } from "../lib/api";
import { checkApprovalStatus } from "../lib/contracts";
import { ethers } from "ethers";
import { useWallet } from "../hooks/useWallet";
import { getContractAddresses, getAdminWalletState } from "../lib/wallet";

interface User {
  ID: string;
  Email: string;
  Name: string;
  WalletAddress: string;
  Role: string;
  IsApproved: boolean;
}

interface UserWithApprovalStatus extends User {
  blockchainApproved: boolean | null;
  approving: boolean;
  pendingConfirmation: boolean;
}

function UserApprovalManagementPage() {
  const [users, setUsers] = useState<UserWithApprovalStatus[]>([]);
  const [filteredUsers, setFilteredUsers] = useState<UserWithApprovalStatus[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");
  const [filter, setFilter] = useState<'all' | 'approved' | 'not-approved'>('all');
  const [isCheckingBlockchain, setIsCheckingBlockchain] = useState(false);
  const { provider, isConnected, connect, address } = useWallet();
  const loadUsersRunning = useRef(false);

  // Get admin address from contract deployment
  const contractAddresses = getContractAddresses();
  const adminAddress = contractAddresses.adminAddress;

  // Clear any wallet disconnection flags on mount to allow auto-connection
  useEffect(() => {
    localStorage.removeItem('wallet_disconnected');
  }, []);

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

  // Load users on component mount
  useEffect(() => {
    loadUsers();
  }, []); // Only run once on mount

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

    setFilteredUsers(filtered);
  }, [users, filter]);

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

  const checkBlockchainStatus = async (currentUsers: UserWithApprovalStatus[]) => {
    console.log('🔄 Starting blockchain status check...');
    console.log('🔗 Wallet connected:', isConnected, 'Provider available:', !!provider);
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
      return currentUsers;
    }

    try {
      const pendingApprovals = getPendingApprovals();
      console.log('📋 Pending approvals:', pendingApprovals);

      // Check users sequentially to avoid overwhelming the RPC endpoint
      const checkAllPromise = (async () => {
        const results: UserWithApprovalStatus[] = [];

        // Ensure we have a valid provider
        if (!blockchainProvider) {
          console.error('❌ No blockchain provider available');
          return currentUsers;
        }

        for (const user of currentUsers) {
          // Skip users with empty wallet addresses
          if (!user.WalletAddress || user.WalletAddress.trim() === "") {
            console.log(`⏭️ Skipping ${user.Email}: No wallet address set`);
            results.push({ ...user, blockchainApproved: null, pendingConfirmation: false });
            continue;
          }

          try {
            console.log(`🔍 Checking approval status for ${user.Email} (${user.WalletAddress})`);
            console.log(`📋 Previous approval status: ${user.blockchainApproved}`);

            // Add timeout to blockchain check - reduced from 3000ms to 1000ms for faster feedback
            const checkPromise = checkApprovalStatus(user.WalletAddress, blockchainProvider!);
            const timeoutPromise = new Promise<boolean>((_, reject) =>
              setTimeout(() => reject(new Error('Blockchain check timeout')), 1000)
            );

            const blockchainApproved = await Promise.race([checkPromise, timeoutPromise]);
            console.log(`✅ Approval status for ${user.Email}: ${blockchainApproved} (was: ${user.blockchainApproved})`);

            // CRITICAL: Once approved on blockchain, status should never revert to false
            // This prevents approved users from appearing as "Not approved" due to RPC issues
            const wasPreviouslyApproved = user.blockchainApproved === true;

            if (wasPreviouslyApproved && !blockchainApproved) {
              console.warn(`⚠️ WARNING: User ${user.Email} was previously approved but blockchain check returned false. This may indicate an RPC issue. Keeping approved status.`);
              results.push({ ...user, blockchainApproved: true, pendingConfirmation: false });
            } else if (pendingApprovals[user.WalletAddress] && !blockchainApproved) {
              console.log(`⏳ Keeping ${user.Email} in pending state (transaction not yet mined)`);
              results.push({ ...user, blockchainApproved: null, pendingConfirmation: true });
            } else {
              // Update persisted approvals when status is confirmed
              if (blockchainApproved === true && user.WalletAddress) {
                const persistedApprovals = getPersistedApprovals();
                persistedApprovals[user.WalletAddress] = true;
                setPersistedApprovals(persistedApprovals);
                console.log(`💾 Persisted approval status for ${user.Email}`);
              }
              results.push({ ...user, blockchainApproved, pendingConfirmation: false });
            }

            // Reduced delay between checks for faster processing
            await new Promise(resolve => setTimeout(resolve, 100));
          } catch (err) {
            console.error(`❌ Failed to check approval for ${user.Email}:`, err);
            // If blockchain check fails, fall back to database-only approval status
            console.log(`ℹ️ Blockchain check failed for ${user.Email}, using database status: ${user.IsApproved}`);
            results.push({ ...user, blockchainApproved: user.IsApproved ? true : false, pendingConfirmation: false });
          }
        }

        return results;
      })();

      const overallTimeoutPromise = new Promise<UserWithApprovalStatus[]>((_, reject) =>
        setTimeout(() => reject(new Error('Overall blockchain check timeout')), 15000)
      );

      const updatedUsers = await Promise.race([checkAllPromise, overallTimeoutPromise]);
      console.log('✅ Blockchain status check completed');
      return updatedUsers;

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
      // Fetch real users from the backend API
      const usersResponse = await api.getUsers();

      // Convert backend user format to frontend format
      const sampleUsers: User[] = usersResponse.map((user: any) => ({
        ID: user.ID,
        Email: user.Email,
        Name: user.Name,
        WalletAddress: user.WalletAddress,
        Role: user.Role,
        IsApproved: user.IsApproved
      }));

      console.log('📋 Created sample users:', sampleUsers.length);

      // Load persisted approval status
      const persistedApprovals = getPersistedApprovals();
      console.log('📚 Loaded persisted approvals:', persistedApprovals);

      // Create users with initial state, preserving existing blockchain approval status if available
      const usersWithInitialState: UserWithApprovalStatus[] = sampleUsers.map(user => {
        // Try to find existing user with the same ID to preserve blockchain status
        const existingUser = users.find(u => u.ID === user.ID);

        // First check existing state, then persisted state, then default to null
        let preservedStatus = existingUser?.blockchainApproved ?? null;

        // If no existing state, check persisted approvals
        if (preservedStatus === null && user.WalletAddress) {
          preservedStatus = persistedApprovals[user.WalletAddress] ?? null;
          if (preservedStatus === true) {
            console.log(`💾 Restored approved status for ${user.Email} from localStorage`);
          }
        } else if (existingUser && existingUser.blockchainApproved === true) {
          console.log(`🛡️ Preserving approved status for ${user.Email}`);
        }

        return {
          ...user,
          blockchainApproved: preservedStatus,
          approving: false, // Reset approving state
          pendingConfirmation: existingUser?.pendingConfirmation ?? false
        };
      });

      console.log('🔗 Checking blockchain status...');

      console.log('📋 Setting initial users without blockchain status');
      setUsers(usersWithInitialState);

      // Check blockchain status asynchronously if wallet is available
      // Prioritize users that don't have persisted approval status
      const usersNeedingCheck = usersWithInitialState.filter(user =>
        user.WalletAddress &&
        user.WalletAddress.trim() !== "" &&
        user.blockchainApproved === null
      );

      if (usersNeedingCheck.length > 0 && isConnected && provider) {
        console.log(`🌐 Starting async blockchain status check for ${usersNeedingCheck.length} users without persisted status...`);
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
      } else if (usersNeedingCheck.length === 0) {
        console.log('ℹ️ All users have persisted approval status, skipping blockchain check');
      } else {
        console.log('⚠️ Wallet not connected, blockchain status will update when connected');
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
      const fallbackUsers: UserWithApprovalStatus[] = [];
      setUsers(fallbackUsers);
    } finally {
      console.log('🏁 loadUsers completed');
      loadUsersRunning.current = false;
      setLoading(false);
    }
  };

  const handleRetryApproval = (user: UserWithApprovalStatus) => {
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

  const handleApproveUser = async (user: UserWithApprovalStatus) => {
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
      setError(err.message || "Failed to approve user");

      // Reset pending state on error
      const pendingApprovals = getPendingApprovals();
      delete pendingApprovals[user.WalletAddress];
      setPendingApprovals(pendingApprovals);

      setUsers(prev => prev.map(u =>
        u.ID === user.ID ? { ...u, approving: false, pendingConfirmation: false } : u
      ));

      alert(`❌ Approval failed: ${err.message}`);
    }
  };

  const getApprovalStatus = (user: UserWithApprovalStatus) => {
    if (user.approving) return { status: "🔄 Approving...", color: "text-blue-400" };
    if (user.pendingConfirmation) return { status: "⏳ Pending Confirmation (Transaction sent, awaiting blockchain)", color: "text-yellow-400" };
    if (user.blockchainApproved === null) return { status: "🔍 Checking blockchain...", color: "text-gray-400" };
    if (user.blockchainApproved) return { status: "✅ Approved", color: "text-green-400" };
    return { status: "❌ Not Approved", color: "text-red-400" };
  };

  if (loading) {
    return (
      <div className="min-h-screen bg-[#0f0f0f] text-white">
        <Navbar currentPage="/admin/approve-users" />
        <div className="flex items-center justify-center pt-32">
          <p>Loading users...</p>
        </div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-[#0f0f0f] text-white">
      <Navbar currentPage="/admin/approve-users" />

      <main className="max-w-6xl mx-auto px-6 pt-28 pb-16">
        <div className="space-y-8">
          {/* Header */}
          <div className="flex items-center justify-between">
            <div>
              <h1 className="text-3xl font-semibold">User Approval Management</h1>
              <p className="text-gray-400 mt-2">
                Approve registered users to enable them to participate in the platform and receive tokens.
              </p>
            </div>
            <div className="flex gap-2">
              <Button
                variant="outline"
                size="md"
              onClick={async () => {
                if (users.length > 0) {
                  console.log('🔄 Manual refresh triggered');
                  try {
                    const updatedUsers = await checkBlockchainStatus(users);
                    console.log('✅ Manual refresh completed');
                    setUsers(updatedUsers);
                  } catch (err) {
                    console.error('❌ Manual refresh failed:', err);
                  }
                } else {
                  loadUsers();
                }
              }}
                className="border-[#262626] bg-[#1A1A1A] text-white hover:bg-[#262626]"
              >
                Refresh
              </Button>
            </div>
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
                  All ({users.length})
                </button>
                <button
                  onClick={() => setFilter('approved')}
                  className={
                    filter === 'approved'
                      ? 'px-3 py-1 rounded-lg text-sm font-medium transition-colors bg-green-600 text-white'
                      : 'px-3 py-1 rounded-lg text-sm font-medium transition-colors bg-[#262626] text-gray-300 hover:bg-[#333333]'
                  }
                >
                  Approved ({users.filter(u => u.blockchainApproved === true).length})
                </button>
                <button
                  onClick={() => setFilter('not-approved')}
                  className={
                    filter === 'not-approved'
                      ? 'px-3 py-1 rounded-lg text-sm font-medium transition-colors bg-red-600 text-white'
                      : 'px-3 py-1 rounded-lg text-sm font-medium transition-colors bg-[#262626] text-gray-300 hover:bg-[#333333]'
                  }
                >
                  Not Approved ({users.filter(u => u.blockchainApproved === false).length})
                </button>
              </div>
            </div>
            <div className="text-sm text-gray-400">
              Showing {filteredUsers.length} of {users.length} users
              {isCheckingBlockchain && (
                <span className="ml-2 text-blue-400">
                  🔄 Checking blockchain status...
                </span>
              )}
            </div>
          </div>

          {/* Error Display */}
          {error && (
            <div className="text-red-500 bg-red-500/10 p-4 rounded-xl border border-red-500/20">
              {error}
            </div>
          )}

          {/* Users Table */}
          <div className="bg-[#1A1A1A] border border-[#262626] rounded-xl overflow-hidden">
            <div className="px-6 py-4 border-b border-[#262626] flex justify-between items-center">
              <h2 className="text-xl font-semibold">Registered Users</h2>
            </div>

            {filteredUsers.length === 0 ? (
              <div className="p-8 text-center text-gray-400">
                {filter === 'all'
                  ? "No users found. New user registrations will appear here automatically."
                  : `No users found with ${filter === 'approved' ? 'approved' : 'not approved'} status.`
                }
              </div>
            ) : (
              <div className="overflow-x-auto">
                <table className="w-full">
                  <thead className="bg-[#0f0f0f]">
                    <tr>
                      <th className="px-6 py-3 text-left text-xs font-medium text-gray-400 uppercase tracking-wider">
                        User
                      </th>
                      <th className="px-6 py-3 text-left text-xs font-medium text-gray-400 uppercase tracking-wider">
                        Wallet Address
                      </th>
                      <th className="px-6 py-3 text-left text-xs font-medium text-gray-400 uppercase tracking-wider">
                        Role
                      </th>
                      <th className="px-6 py-3 text-left text-xs font-medium text-gray-400 uppercase tracking-wider">
                        Blockchain Status
                      </th>
                      <th className="px-6 py-3 text-left text-xs font-medium text-gray-400 uppercase tracking-wider">
                        Actions
                      </th>
                    </tr>
                  </thead>
                  <tbody className="divide-y divide-[#262626]">
                    {filteredUsers.map((user) => {
                      const approvalStatus = getApprovalStatus(user);
                      return (
                        <tr key={user.ID} className="hover:bg-[#0f0f0f]/50">
                          <td className="px-6 py-4 whitespace-nowrap">
                            <div>
                              <div className="text-sm font-medium text-white">{user.Name}</div>
                              <div className="text-sm text-gray-400">{user.Email}</div>
                            </div>
                          </td>
                          <td className="px-6 py-4 whitespace-nowrap">
                            <div className="text-sm font-mono text-gray-300">
                              {user.WalletAddress && user.WalletAddress.length > 10
                                ? `${user.WalletAddress.substring(0, 6)}...${user.WalletAddress.substring(38)}`
                                : user.WalletAddress || "Not set"
                              }
                            </div>
                          </td>
                          <td className="px-6 py-4 whitespace-nowrap">
                            <span className={`inline-flex px-2 py-1 text-xs font-semibold rounded-full ${
                              user.Role === 'admin'
                                ? 'bg-purple-500/20 text-purple-400'
                                : 'bg-blue-500/20 text-blue-400'
                            }`}>
                              {user.Role}
                            </span>
                          </td>
                          <td className="px-6 py-4 whitespace-nowrap">
                            <span className={`text-sm ${approvalStatus.color}`}>
                              {approvalStatus.status}
                            </span>
                          </td>
                          <td className="px-6 py-4 whitespace-nowrap">
                            {user.blockchainApproved ? (
                              <span className="text-green-400 text-sm">Already Approved</span>
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
                                <Button
                                  variant="outline"
                                  size="sm"
                                  onClick={async () => {
                                    try {
                                      const updatedUsers = await checkBlockchainStatus([user]);
                                      setUsers(prev => prev.map(u =>
                                        u.ID === user.ID ? updatedUsers.find(uu => uu.ID === user.ID) || u : u
                                      ));
                                    } catch (err) {
                                      console.error('❌ Individual status check failed:', err);
                                    }
                                  }}
                                  className="border-blue-500/50 bg-blue-500/10 text-blue-400 hover:bg-blue-500/20 text-xs"
                                >
                                  Check Status
                                </Button>
                              </div>
                            ) : user.WalletAddress && user.WalletAddress.trim() !== "" ? (
                              <Button
                                variant="primary"
                                size="sm"
                                onClick={() => handleApproveUser(user)}
                                disabled={user.approving}
                                className="bg-[#6d41ff] hover:bg-[#5b2fff] text-white disabled:opacity-50"
                              >
                                {user.approving ? "Approving..." : "Approve"}
                              </Button>
                            ) : (
                              <span className="text-gray-500 text-sm">No wallet address</span>
                            )}
                          </td>
                        </tr>
                      );
                    })}
                  </tbody>
                </table>
              </div>
            )}
          </div>

          {/* Controls */}
          <div className="flex gap-4 mb-6 flex-wrap">
            <Button
              variant="outline"
              size="sm"
              onClick={() => loadUsers()}
              disabled={loading || loadUsersRunning.current}
              className="border-gray-500 bg-gray-500/10 text-gray-300 hover:bg-gray-500/20"
            >
              {loading ? "Refreshing..." : "🔄 Refresh Users"}
            </Button>
            <Button
              variant="outline"
              size="sm"
              onClick={async () => {
                if (users.length > 0) {
                  setIsCheckingBlockchain(true);
                  try {
                    const usersNeedingCheck = users.filter(user =>
                      user.WalletAddress && user.WalletAddress.trim() !== ""
                    );
                    const updatedUsers = await checkBlockchainStatus(usersNeedingCheck);
                    setUsers(prevUsers => prevUsers.map(user => {
                      const checkedUser = updatedUsers.find(cu => cu.ID === user.ID);
                      return checkedUser || user;
                    }));
                    console.log('✅ Manual blockchain status check completed');
                  } catch (err) {
                    console.error('❌ Manual blockchain check failed:', err);
                  } finally {
                    setIsCheckingBlockchain(false);
                  }
                }
              }}
              disabled={isCheckingBlockchain || users.length === 0}
              className="border-blue-500 bg-blue-500/10 text-blue-400 hover:bg-blue-500/20"
            >
              {isCheckingBlockchain ? "Checking..." : "🔍 Check Blockchain Status"}
            </Button>
            <Button
              variant="outline"
              size="sm"
              onClick={() => {
                if (confirm('Clear all persisted approval data? This will reset approval status on page refresh.')) {
                  localStorage.removeItem('userApprovals');
                  console.log('🗑️ Cleared persisted approvals');
                  // Reload users to reset their status
                  loadUsers();
                }
              }}
              className="border-red-500 bg-red-500/10 text-red-400 hover:bg-red-500/20"
            >
              🗑️ Clear Cache
            </Button>
          </div>

          {/* Info Section */}
          <div className="bg-blue-500/10 border border-blue-500/20 rounded-xl p-6">
            <h3 className="text-blue-400 font-semibold mb-2">ℹ️ How User Approval Works</h3>
            <ul className="text-sm text-gray-300 space-y-1">
              <li>• <strong>Registration:</strong> Users register with email/password and wallet address</li>
              <li>• <strong>Approval Required:</strong> Investor users must be approved on the blockchain to receive tokens</li>
              <li>• <strong>Admin Action:</strong> Only admins can approve users via this interface</li>
              <li>• <strong>Admin Role:</strong> Admins have special privileges and don't need approval</li>
              <li>• <strong>Auto-population:</strong> All registered users from the database appear here automatically</li>
              <li>• <strong>Blockchain Transaction:</strong> Each approval creates a transaction on Sepolia</li>
              <li>• <strong>Smart Contract:</strong> Calls <code>ApprovalService.approve(userAddress)</code> on Sepolia</li>
              <li>• <strong>Status Checking:</strong> Reads <code>ApprovalService.check(userAddress)</code></li>
              <li>• <strong>Sepolia Network:</strong> Transactions may take 15-30 seconds to confirm</li>
              <li>• <strong>Refresh Status:</strong> Use the refresh button to check latest approval status</li>
            </ul>
          </div>

          {/* Debug Section */}
          <div className="bg-yellow-500/10 border border-yellow-500/20 rounded-xl p-6">
            <h3 className="text-yellow-400 font-semibold mb-2">🔍 Debug Information</h3>
            <div className="text-sm text-gray-300 space-y-1">
              <p><strong>Contract Address:</strong> {contractAddresses.ApprovalService}</p>
              <p><strong>Network:</strong> Sepolia Testnet</p>
              <p><strong>Admin Address:</strong> {adminAddress}</p>
              <p><strong>Wallet Connected:</strong> {isConnected ? '✅ Yes' : '❌ No'}</p>
              {isConnected && address && (
                <p><strong>Connected Address:</strong> <code className="bg-gray-800 px-2 py-1 rounded text-xs">{address}</code></p>
              )}
              <p><strong>Authenticated:</strong> {localStorage.getItem("token") ? '✅ Yes' : '❌ No'}</p>
              {localStorage.getItem("token") && (
                <p><strong>Token Preview:</strong> {localStorage.getItem("token")?.substring(0, 20)}...</p>
              )}
              {!isConnected && (
                <div className="mt-2">
                  <Button
                    variant="outline"
                    size="sm"
                    onClick={connect}
                    className="border-blue-500/50 bg-blue-500/10 text-blue-400 hover:bg-blue-500/20"
                  >
                    🔗 Connect Wallet
                  </Button>
                </div>
              )}
              <p><strong>Note:</strong> The admin address ({adminAddress}) is the signer for approval transactions. Individual users need to be approved to receive tokens.</p>
              {isConnected && address && (
                <p className={`text-sm ${address.toLowerCase() === adminAddress.toLowerCase() ? 'text-green-400' : 'text-yellow-400'}`}>
                  {address.toLowerCase() === adminAddress.toLowerCase()
                    ? '✅ Connected with admin wallet'
                    : '⚠️ Connected with different wallet - switch to admin wallet to approve it'}
                </p>
              )}
              <p><strong>Check browser console for detailed logs when clicking "Refresh Status"</strong></p>
            </div>
          </div>
        </div>
      </main>
    </div>
  );
}

export default UserApprovalManagementPage;

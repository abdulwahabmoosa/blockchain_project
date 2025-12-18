import { useEffect, useState } from "react";
import { RefreshCw, Search, CheckCircle, XCircle, AlertTriangle } from "lucide-react";
import { Button } from "../Components/atoms/Button";
import { api } from "../lib/api";

interface UserData {
  ID: string;
  Email: string;
  Name: string;
  WalletAddress: string;
  Role: string;
  IsApproved: boolean;
}

interface UserWithStatus extends UserData {
  approving: boolean;
}

function AdminManageUsers() {
  const [users, setUsers] = useState<UserWithStatus[]>([]);
  const [filteredUsers, setFilteredUsers] = useState<UserWithStatus[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");
  const [filter, setFilter] = useState<'all' | 'approved' | 'not-approved'>('all');
  const [searchQuery, setSearchQuery] = useState("");

  useEffect(() => {
    loadUsers();
  }, []);

  useEffect(() => {
    let filtered = users;

    if (searchQuery) {
      filtered = filtered.filter(user =>
        user.Name.toLowerCase().includes(searchQuery.toLowerCase()) ||
        user.Email.toLowerCase().includes(searchQuery.toLowerCase()) ||
        (user.WalletAddress && user.WalletAddress.toLowerCase().includes(searchQuery.toLowerCase()))
      );
    }

    switch (filter) {
      case 'approved':
        filtered = filtered.filter(user => user.IsApproved === true);
        break;
      case 'not-approved':
        filtered = filtered.filter(user => user.IsApproved === false);
        break;
    }

    setFilteredUsers(filtered);
  }, [users, filter, searchQuery]);

  const loadUsers = async () => {
    setLoading(true);
    setError("");

    try {
      // Note: You'll need to add getUsers to your api.ts
      // For now, this will show an empty list if the API doesn't exist
      const response = await fetch("/api/users", {
        headers: {
          "Authorization": `Bearer ${localStorage.getItem("token")}`,
          "Content-Type": "application/json",
        },
      });

      if (!response.ok) {
        throw new Error("Failed to fetch users");
      }

      const usersData = await response.json();
      const usersWithStatus: UserWithStatus[] = usersData.map((user: UserData) => ({
        ...user,
        approving: false,
      }));

      setUsers(usersWithStatus);
    } catch (err: any) {
      console.error("Failed to load users:", err);
      setError(err.message || "Failed to load users");
      setUsers([]);
    } finally {
      setLoading(false);
    }
  };

  const handleApproveUser = async (user: UserWithStatus) => {
    if (!user.WalletAddress || user.WalletAddress.trim() === "") {
      alert("Cannot approve user: No wallet address is set.");
      return;
    }

    setUsers(prev => prev.map(u =>
      u.ID === user.ID ? { ...u, approving: true } : u
    ));

    try {
      const response = await api.approveUser(user.WalletAddress);

      setUsers(prev => prev.map(u =>
        u.ID === user.ID ? {
          ...u,
          IsApproved: true,
          approving: false
        } : u
      ));

      alert(`Approval successful! TX: ${response.tx_hash}`);
    } catch (err: any) {
      console.error("Approval failed:", err);
      setError(err.message || "Failed to approve user");

      setUsers(prev => prev.map(u =>
        u.ID === user.ID ? { ...u, approving: false } : u
      ));
    }
  };

  const getStatusBadge = (user: UserWithStatus) => {
    if (user.approving) {
      return (
        <span className="flex items-center gap-1 text-xs px-2 py-1 rounded-full bg-blue-500/20 text-blue-400">
          <RefreshCw size={12} className="animate-spin" />
          Approving...
        </span>
      );
    }
    if (user.IsApproved) {
      return (
        <span className="flex items-center gap-1 text-xs px-2 py-1 rounded-full bg-green-500/20 text-green-400">
          <CheckCircle size={12} />
          Approved
        </span>
      );
    }
    return (
      <span className="flex items-center gap-1 text-xs px-2 py-1 rounded-full bg-red-500/20 text-red-400">
        <XCircle size={12} />
        Not Approved
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
        <Button
          variant="outline"
          size="md"
          onClick={() => loadUsers()}
          disabled={loading}
          className="border-[#262626] text-white hover:bg-[#1a1a1a]"
        >
          <RefreshCw size={18} className={`mr-2 ${loading ? 'animate-spin' : ''}`} />
          Refresh
        </Button>
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
          <Search className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400" size={20} />
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
            onClick={() => setFilter('all')}
            className={`px-4 py-2 rounded-xl text-sm font-medium transition-colors ${
              filter === 'all'
                ? 'bg-[#703BF7] text-white'
                : 'bg-[#111111] border border-[#1f1f1f] text-gray-400 hover:text-white'
            }`}
          >
            All ({users.length})
          </button>
          <button
            onClick={() => setFilter('approved')}
            className={`px-4 py-2 rounded-xl text-sm font-medium transition-colors ${
              filter === 'approved'
                ? 'bg-green-600 text-white'
                : 'bg-[#111111] border border-[#1f1f1f] text-gray-400 hover:text-white'
            }`}
          >
            Approved ({users.filter(u => u.IsApproved === true).length})
          </button>
          <button
            onClick={() => setFilter('not-approved')}
            className={`px-4 py-2 rounded-xl text-sm font-medium transition-colors ${
              filter === 'not-approved'
                ? 'bg-red-600 text-white'
                : 'bg-[#111111] border border-[#1f1f1f] text-gray-400 hover:text-white'
            }`}
          >
            Pending ({users.filter(u => u.IsApproved === false).length})
          </button>
        </div>
      </div>

      {/* Users List */}
      <div className="rounded-3xl border border-[#1f1f1f] bg-[#111111] overflow-hidden">
        <div className="p-4 border-b border-[#1f1f1f] flex items-center justify-between">
          <h2 className="font-semibold">Registered Users</h2>
          <span className="text-xs text-gray-500">
            {filteredUsers.length} of {users.length} users
          </span>
        </div>

        {loading ? (
          <div className="p-8 text-center">
            <RefreshCw size={24} className="animate-spin mx-auto text-gray-400 mb-2" />
            <p className="text-gray-400">Loading users...</p>
          </div>
        ) : filteredUsers.length === 0 ? (
          <div className="p-8 text-center text-gray-400">
            {searchQuery || filter !== 'all'
              ? "No users match your search criteria"
              : "No users found"}
          </div>
        ) : (
          <div className="divide-y divide-[#1f1f1f]">
            {filteredUsers.map((user) => (
              <div key={user.ID} className="p-4 hover:bg-[#0f0f0f]/50 transition-colors">
                <div className="flex flex-col md:flex-row md:items-center justify-between gap-4">
                  <div className="flex-1 min-w-0">
                    <div className="flex items-center gap-2 mb-1">
                      <h3 className="font-medium truncate">{user.Name}</h3>
                      <span className={`text-xs px-2 py-0.5 rounded-full ${
                        user.Role === 'admin'
                          ? 'bg-purple-500/20 text-purple-400'
                          : 'bg-blue-500/20 text-blue-400'
                      }`}>
                        {user.Role}
                      </span>
                    </div>
                    <p className="text-sm text-gray-400 truncate">{user.Email}</p>
                    <p className="text-xs text-gray-500 font-mono mt-1">
                      {user.WalletAddress
                        ? `${user.WalletAddress.substring(0, 10)}...${user.WalletAddress.substring(38)}`
                        : "No wallet"}
                    </p>
                  </div>

                  <div className="flex items-center gap-3">
                    {getStatusBadge(user)}

                    {!user.IsApproved && !user.approving && user.WalletAddress && (
                      <Button
                        variant="primary"
                        size="sm"
                        onClick={() => handleApproveUser(user)}
                        className="bg-[#6d41ff] hover:bg-[#5b2fff] text-white"
                      >
                        Approve
                      </Button>
                    )}

                    {user.IsApproved && (
                      <span className="text-xs text-green-400">Already Approved</span>
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

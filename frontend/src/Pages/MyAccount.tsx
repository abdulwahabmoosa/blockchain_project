import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import {
  User,
  Lock,
  Trash2,
  Wallet,
  Copy,
  CheckCircle,
  AlertTriangle,
  ExternalLink,
} from "lucide-react";
import { Button } from "../Components/atoms/Button";
import type { User as UserType } from "../types";
import { api } from "../lib/api";

function MyAccount() {
  const navigate = useNavigate();
  const [user, setUser] = useState<UserType | null>(null);
  const [copied, setCopied] = useState(false);

  // Password form state
  const [currentPassword, setCurrentPassword] = useState("");
  const [newPassword, setNewPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");
  const [passwordLoading, setPasswordLoading] = useState(false);
  const [passwordError, setPasswordError] = useState("");
  const [passwordSuccess, setPasswordSuccess] = useState(false);

  // Delete account state
  const [showDeleteConfirm, setShowDeleteConfirm] = useState(false);
  const [deleteConfirmText, setDeleteConfirmText] = useState("");
  const [deleteLoading, setDeleteLoading] = useState(false);
  const [deleteError, setDeleteError] = useState("");

  useEffect(() => {
    const userStr = localStorage.getItem("user");
    if (userStr) {
      setUser(JSON.parse(userStr));
    }
  }, []);

  const copyAddress = () => {
    if (user?.WalletAddress) {
      navigator.clipboard.writeText(user.WalletAddress);
      setCopied(true);
      setTimeout(() => setCopied(false), 2000);
    }
  };

  const handlePasswordUpdate = async (e: React.FormEvent) => {
    e.preventDefault();
    setPasswordError("");
    setPasswordSuccess(false);

    if (newPassword !== confirmPassword) {
      setPasswordError("New passwords do not match");
      return;
    }

    if (newPassword.length < 6) {
      setPasswordError("Password must be at least 6 characters");
      return;
    }

    setPasswordLoading(true);

    try {
      await api.updatePassword(currentPassword, newPassword);

      setPasswordSuccess(true);
      setCurrentPassword("");
      setNewPassword("");
      setConfirmPassword("");
    } catch (err: any) {
      setPasswordError(err.message || "Failed to update password");
    } finally {
      setPasswordLoading(false);
    }
  };

  const handleDeleteAccount = async () => {
    if (deleteConfirmText !== "DELETE") {
      setDeleteError("Please type DELETE to confirm");
      return;
    }

    setDeleteLoading(true);
    setDeleteError("");

    try {
      await api.deleteAccount();

      // Clear local storage and redirect
      localStorage.removeItem("auth");
      localStorage.removeItem("token");
      localStorage.removeItem("user");
      navigate("/");
    } catch (err: any) {
      setDeleteError(err.message || "Failed to delete account");
    } finally {
      setDeleteLoading(false);
    }
  };

  return (
    <div className="space-y-6 max-w-3xl">
      <header>
        <h1 className="text-3xl font-semibold">My Account</h1>
        <p className="text-sm text-gray-400">
          Manage your account settings and preferences
        </p>
      </header>

      {/* Profile Overview */}
      <div className="rounded-3xl border border-[#1f1f1f] bg-[#111111] p-6 space-y-4">
        <div className="flex items-center gap-3">
          <div className="h-10 w-10 rounded-full bg-[#6d41ff] flex items-center justify-center">
            <User size={20} />
          </div>
          <h2 className="font-semibold text-lg">Profile Information</h2>
        </div>

        <div className="grid md:grid-cols-2 gap-4">
          <div className="p-4 rounded-xl border border-[#1f1f1f] bg-[#0f0f0f]">
            <p className="text-xs text-gray-500">Name</p>
            <p className="font-medium mt-1">{user?.Name || "N/A"}</p>
          </div>
          <div className="p-4 rounded-xl border border-[#1f1f1f] bg-[#0f0f0f]">
            <p className="text-xs text-gray-500">Email</p>
            <p className="font-medium mt-1">{user?.Email || "N/A"}</p>
          </div>
          <div className="p-4 rounded-xl border border-[#1f1f1f] bg-[#0f0f0f]">
            <p className="text-xs text-gray-500">Role</p>
            <p className="font-medium mt-1 capitalize">{user?.Role || "N/A"}</p>
          </div>
          <div className="p-4 rounded-xl border border-[#1f1f1f] bg-[#0f0f0f]">
            <p className="text-xs text-gray-500">Member Since</p>
            <p className="font-medium mt-1">
              {user?.CreatedAt
                ? new Date(user.CreatedAt).toLocaleDateString()
                : "N/A"}
            </p>
          </div>
        </div>
      </div>

      {/* Wallet Section */}
      <div className="rounded-3xl border border-[#1f1f1f] bg-[#111111] p-6 space-y-4">
        <div className="flex items-center gap-3">
          <div className="h-10 w-10 rounded-full bg-[#2c1f5e] flex items-center justify-center">
            <Wallet size={20} />
          </div>
          <h2 className="font-semibold text-lg">Wallet</h2>
        </div>

        <div className="flex items-center justify-between p-4 rounded-xl border border-[#1f1f1f] bg-[#0f0f0f]">
          <div>
            <p className="text-xs text-gray-500">Wallet Address</p>
            <p className="font-mono text-sm mt-1">
              {user?.WalletAddress || "Not connected"}
            </p>
          </div>
          <div className="flex gap-2">
            <button
              onClick={copyAddress}
              className="p-2 rounded-lg hover:bg-[#1a1a1a] text-gray-400 hover:text-white transition-colors"
              title="Copy address"
            >
              {copied ? (
                <CheckCircle size={18} className="text-green-500" />
              ) : (
                <Copy size={18} />
              )}
            </button>
            <a
              href={`https://etherscan.io/address/${user?.WalletAddress}`}
              target="_blank"
              rel="noopener noreferrer"
              className="p-2 rounded-lg hover:bg-[#1a1a1a] text-gray-400 hover:text-white transition-colors"
              title="View on Etherscan"
            >
              <ExternalLink size={18} />
            </a>
          </div>
        </div>

        <div className="flex items-center gap-2">
          <span
            className={`text-xs px-2 py-1 rounded-full ${
              user?.IsApproved
                ? "bg-green-500/20 text-green-400"
                : "bg-yellow-500/20 text-yellow-400"
            }`}
          >
            {user?.IsApproved ? "Verified" : "Pending Verification"}
          </span>
        </div>

        <Button
          variant="outline"
          size="md"
          className="border-[#262626] text-white hover:bg-[#1a1a1a]"
          onClick={() => navigate("/dashboard/wallet")}
        >
          View Full Wallet Details
        </Button>
      </div>

      {/* Update Password */}
      <div className="rounded-3xl border border-[#1f1f1f] bg-[#111111] p-6 space-y-4">
        <div className="flex items-center gap-3">
          <div className="h-10 w-10 rounded-full bg-[#1f3a5e] flex items-center justify-center">
            <Lock size={20} />
          </div>
          <h2 className="font-semibold text-lg">Update Password</h2>
        </div>

        {passwordError && (
          <div className="text-red-500 bg-red-500/10 p-3 rounded-xl border border-red-500/20 text-sm">
            {passwordError}
          </div>
        )}

        {passwordSuccess && (
          <div className="text-green-500 bg-green-500/10 p-3 rounded-xl border border-green-500/20 text-sm flex items-center gap-2">
            <CheckCircle size={16} />
            Password updated successfully
          </div>
        )}

        <form onSubmit={handlePasswordUpdate} className="space-y-4">
          <div className="space-y-2">
            <label className="text-sm text-gray-400">Current Password</label>
            <input
              type="password"
              value={currentPassword}
              onChange={(e) => setCurrentPassword(e.target.value)}
              placeholder="Enter current password"
              required
              className="w-full px-4 py-3 rounded-xl border border-[#1f1f1f] bg-[#0f0f0f] text-white placeholder-gray-500 focus:outline-none focus:border-[#6d41ff]"
            />
          </div>

          <div className="grid md:grid-cols-2 gap-4">
            <div className="space-y-2">
              <label className="text-sm text-gray-400">New Password</label>
              <input
                type="password"
                value={newPassword}
                onChange={(e) => setNewPassword(e.target.value)}
                placeholder="Enter new password"
                required
                className="w-full px-4 py-3 rounded-xl border border-[#1f1f1f] bg-[#0f0f0f] text-white placeholder-gray-500 focus:outline-none focus:border-[#6d41ff]"
              />
            </div>

            <div className="space-y-2">
              <label className="text-sm text-gray-400">Confirm New Password</label>
              <input
                type="password"
                value={confirmPassword}
                onChange={(e) => setConfirmPassword(e.target.value)}
                placeholder="Confirm new password"
                required
                className="w-full px-4 py-3 rounded-xl border border-[#1f1f1f] bg-[#0f0f0f] text-white placeholder-gray-500 focus:outline-none focus:border-[#6d41ff]"
              />
            </div>
          </div>

          <Button
            type="submit"
            variant="primary"
            size="md"
            className="bg-[#6d41ff] hover:bg-[#5b2fff] text-white"
            disabled={passwordLoading}
          >
            {passwordLoading ? "Updating..." : "Update Password"}
          </Button>
        </form>
      </div>

      {/* Delete Account */}
      <div className="rounded-3xl border border-red-500/20 bg-[#111111] p-6 space-y-4">
        <div className="flex items-center gap-3">
          <div className="h-10 w-10 rounded-full bg-red-500/20 flex items-center justify-center text-red-500">
            <Trash2 size={20} />
          </div>
          <div>
            <h2 className="font-semibold text-lg text-red-500">Danger Zone</h2>
            <p className="text-sm text-gray-400">
              Permanently delete your account
            </p>
          </div>
        </div>

        {!showDeleteConfirm ? (
          <Button
            variant="danger"
            size="md"
            className="bg-red-500/10 border border-red-500/30 text-red-500 hover:bg-red-500/20"
            onClick={() => setShowDeleteConfirm(true)}
          >
            Delete Account
          </Button>
        ) : (
          <div className="space-y-4 p-4 rounded-xl border border-red-500/20 bg-red-500/5">
            <div className="flex items-start gap-3">
              <AlertTriangle className="text-red-500 flex-shrink-0 mt-0.5" size={20} />
              <div>
                <p className="text-sm font-medium text-red-500">
                  Are you sure you want to delete your account?
                </p>
                <p className="text-xs text-gray-400 mt-1">
                  This action cannot be undone. All your data, including properties
                  and investments, will be permanently deleted.
                </p>
              </div>
            </div>

            {deleteError && (
              <div className="text-red-500 bg-red-500/10 p-3 rounded-xl border border-red-500/20 text-sm">
                {deleteError}
              </div>
            )}

            <div className="space-y-2">
              <label className="text-sm text-gray-400">
                Type <span className="font-mono text-red-400">DELETE</span> to confirm
              </label>
              <input
                type="text"
                value={deleteConfirmText}
                onChange={(e) => setDeleteConfirmText(e.target.value)}
                placeholder="DELETE"
                className="w-full px-4 py-3 rounded-xl border border-red-500/30 bg-[#0f0f0f] text-white placeholder-gray-500 focus:outline-none focus:border-red-500"
              />
            </div>

            <div className="flex gap-3">
              <Button
                variant="outline"
                size="md"
                className="border-[#262626] text-white hover:bg-[#1a1a1a]"
                onClick={() => {
                  setShowDeleteConfirm(false);
                  setDeleteConfirmText("");
                  setDeleteError("");
                }}
              >
                Cancel
              </Button>
              <Button
                variant="danger"
                size="md"
                className="bg-red-600 hover:bg-red-700 text-white"
                onClick={handleDeleteAccount}
                disabled={deleteLoading || deleteConfirmText !== "DELETE"}
              >
                {deleteLoading ? "Deleting..." : "Delete My Account"}
              </Button>
            </div>
          </div>
        )}
      </div>
    </div>
  );
}

export default MyAccount;




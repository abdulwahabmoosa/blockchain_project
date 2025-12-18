import { useEffect, useState } from "react";
import {
  Wallet,
  Copy,
  ExternalLink,
  ArrowUpRight,
  ArrowDownLeft,
  CheckCircle,
} from "lucide-react";
import { Button } from "../Components/atoms/Button";
import type { User } from "../types";

const StatCard = ({
  title,
  value,
  icon: Icon,
  accent = "bg-[#6d41ff]",
}: {
  title: string;
  value: string;
  icon?: React.ElementType;
  accent?: string;
}) => (
  <div className="rounded-2xl border border-[#1f1f1f] bg-[#111111] p-4 space-y-2">
    <div className="flex items-center justify-between">
      <p className="text-sm text-gray-400">{title}</p>
      {Icon && (
        <div className={`h-8 w-8 rounded-full ${accent} flex items-center justify-center`}>
          <Icon size={16} />
        </div>
      )}
    </div>
    <p className="text-2xl font-semibold">{value}</p>
  </div>
);

function WalletStatus() {
  const [user, setUser] = useState<User | null>(null);
  const [copied, setCopied] = useState(false);

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

  const formatAddress = (address: string) => {
    return `${address.substring(0, 6)}...${address.substring(address.length - 4)}`;
  };

  // Placeholder transactions
  const transactions = [
    {
      id: "1",
      type: "receive",
      amount: "0.00 ETH",
      from: "0x0000...0000",
      date: "No transactions",
    },
  ];

  return (
    <div className="space-y-6">
      <header>
        <h1 className="text-3xl font-semibold">Wallet</h1>
        <p className="text-sm text-gray-400">
          Manage your wallet and view transactions
        </p>
      </header>

      {/* Wallet Overview */}
      <div className="rounded-3xl border border-[#1f1f1f] bg-linear-to-br from-[#6d41ff] via-[#5a3ccc] to-[#1b132f] p-6 space-y-4">
        <div className="flex items-center justify-between">
          <div className="flex items-center gap-3">
            <div className="h-12 w-12 rounded-full bg-white/10 flex items-center justify-center">
              <Wallet size={24} />
            </div>
            <div>
              <p className="text-sm text-white/80">Connected Wallet</p>
              <p className="font-semibold">
                {user?.WalletAddress
                  ? formatAddress(user.WalletAddress)
                  : "Not connected"}
              </p>
            </div>
          </div>
          {user?.IsApproved && (
            <span className="flex items-center gap-1 text-xs text-green-400 bg-green-500/20 px-2 py-1 rounded-full">
              <CheckCircle size={12} />
              Verified
            </span>
          )}
        </div>

        <div className="pt-4 border-t border-white/10">
          <p className="text-sm text-white/80">Total Balance</p>
          <p className="text-4xl font-bold mt-1">0.00 ETH</p>
          <p className="text-sm text-white/60 mt-1">$0.00 USD</p>
        </div>

        <div className="flex gap-3">
          <Button
            variant="outline"
            size="md"
            className="border-white/30 text-white hover:bg-white/10 flex-1"
          >
            <ArrowDownLeft size={18} className="mr-2" />
            Receive
          </Button>
          <Button
            variant="outline"
            size="md"
            className="border-white/30 text-white hover:bg-white/10 flex-1"
          >
            <ArrowUpRight size={18} className="mr-2" />
            Send
          </Button>
        </div>
      </div>

      {/* Wallet Details */}
      <div className="rounded-3xl border border-[#1f1f1f] bg-[#111111] p-6 space-y-4">
        <h2 className="font-semibold text-lg">Wallet Details</h2>

        <div className="space-y-3">
          <div className="flex items-center justify-between p-4 rounded-xl border border-[#1f1f1f] bg-[#0f0f0f]">
            <div>
              <p className="text-xs text-gray-500">Wallet Address</p>
              <p className="font-mono text-sm mt-1">
                {user?.WalletAddress || "Not available"}
              </p>
            </div>
            <div className="flex gap-2">
              <button
                onClick={copyAddress}
                className="p-2 rounded-lg hover:bg-[#1a1a1a] text-gray-400 hover:text-white transition-colors"
                title="Copy address"
              >
                {copied ? <CheckCircle size={18} className="text-green-500" /> : <Copy size={18} />}
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

          <div className="grid md:grid-cols-2 gap-3">
            <div className="p-4 rounded-xl border border-[#1f1f1f] bg-[#0f0f0f]">
              <p className="text-xs text-gray-500">Account Name</p>
              <p className="font-medium mt-1">{user?.Name || "N/A"}</p>
            </div>
            <div className="p-4 rounded-xl border border-[#1f1f1f] bg-[#0f0f0f]">
              <p className="text-xs text-gray-500">Email</p>
              <p className="font-medium mt-1">{user?.Email || "N/A"}</p>
            </div>
          </div>

          <div className="grid md:grid-cols-2 gap-3">
            <div className="p-4 rounded-xl border border-[#1f1f1f] bg-[#0f0f0f]">
              <p className="text-xs text-gray-500">Status</p>
              <p className="font-medium mt-1">
                {user?.IsApproved ? (
                  <span className="text-green-400">Approved</span>
                ) : (
                  <span className="text-yellow-400">Pending Approval</span>
                )}
              </p>
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
      </div>

      {/* Stats */}
      <div className="grid md:grid-cols-3 gap-4">
        <StatCard title="Total Invested" value="0.00 ETH" icon={ArrowUpRight} />
        <StatCard
          title="Total Earned"
          value="0.00 ETH"
          icon={ArrowDownLeft}
          accent="bg-green-500/20"
        />
        <StatCard
          title="Transactions"
          value="0"
          icon={ExternalLink}
          accent="bg-[#2c1f5e]"
        />
      </div>

      {/* Transaction History */}
      <div className="rounded-3xl border border-[#1f1f1f] bg-[#111111] p-6 space-y-4">
        <div className="flex items-center justify-between">
          <h2 className="font-semibold text-lg">Transaction History</h2>
          <span className="text-xs text-gray-500">Recent</span>
        </div>

        <div className="space-y-3">
          {transactions.length > 0 && transactions[0].date !== "No transactions" ? (
            transactions.map((tx) => (
              <div
                key={tx.id}
                className="flex items-center justify-between p-4 rounded-xl border border-[#1f1f1f] bg-[#0f0f0f]"
              >
                <div className="flex items-center gap-3">
                  <div
                    className={`h-10 w-10 rounded-full flex items-center justify-center ${
                      tx.type === "receive"
                        ? "bg-green-500/20 text-green-400"
                        : "bg-red-500/20 text-red-400"
                    }`}
                  >
                    {tx.type === "receive" ? (
                      <ArrowDownLeft size={20} />
                    ) : (
                      <ArrowUpRight size={20} />
                    )}
                  </div>
                  <div>
                    <p className="font-medium capitalize">{tx.type}</p>
                    <p className="text-xs text-gray-500">From: {tx.from}</p>
                  </div>
                </div>
                <div className="text-right">
                  <p
                    className={`font-medium ${
                      tx.type === "receive" ? "text-green-400" : "text-red-400"
                    }`}
                  >
                    {tx.type === "receive" ? "+" : "-"}
                    {tx.amount}
                  </p>
                  <p className="text-xs text-gray-500">{tx.date}</p>
                </div>
              </div>
            ))
          ) : (
            <div className="text-center py-8">
              <p className="text-gray-400">No transactions yet</p>
              <p className="text-gray-500 text-sm mt-1">
                Your transaction history will appear here
              </p>
            </div>
          )}
        </div>
      </div>
    </div>
  );
}

export default WalletStatus;

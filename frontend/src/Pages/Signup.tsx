import "../index.css";
import { useState } from "react";
import { Link } from "../Components/atoms/Link";
import { Button } from "../Components/atoms/Button";
import { Navbar } from "../Components/organisms/Navbar";
import { useNavigate } from "react-router-dom";
import { api } from "../lib/api";

const GradientCard = ({ className = "" }: { className?: string }) => (
  <div
    className={`bg-linear-to-br from-[#151515] via-[#1d1d1d] to-[#111111] border border-[#262626] ${className}`}
  />
);

function SignupPage() {
  const navigate = useNavigate();
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");
  const [walletAddress, setWalletAddress] = useState("");
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);

  const handleSignup = async () => {
    setError("");
    if (password !== confirmPassword) {
      setError("Passwords do not match");
      return;
    }
    if (!walletAddress.startsWith("0x")) {
      setError("Invalid wallet address (must start with 0x)");
      return;
    }

    setLoading(true);
    try {
      await api.register(email, password, name, walletAddress);
      navigate("/login");
    } catch (err: any) {
      setError(err.message || "Registration failed");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="min-h-screen w-full bg-[#0f0f0f] text-white">
      <Navbar currentPage="/signup" />

      <main className="max-w-4xl mx-auto px-4 pt-28 pb-16">
        <div className="grid md:grid-cols-2 gap-10 items-center">
          <div className="space-y-4">
            <p className="uppercase tracking-[0.25em] text-xs text-gray-400">
              Join Propchain
            </p>
            <h1 className="text-4xl font-semibold">Create your account</h1>
            <p className="text-sm text-gray-300 leading-relaxed">
              Sign up to explore properties, manage investments, and access
              personalized real estate insights with Propchain.
            </p>
            <GradientCard className="h-48 rounded-3xl" />
          </div>

          <div className="rounded-3xl border border-[#1f1f1f] bg-[#111111] p-6 space-y-6">
            <div className="space-y-2">
              <h2 className="text-2xl font-semibold">Sign Up</h2>
              <p className="text-sm text-gray-400">
                Enter your details to get started.
              </p>
            </div>

            {error && (
              <div className="text-red-500 text-sm bg-red-500/10 p-3 rounded-lg border border-red-500/20">
                {error}
              </div>
            )}

            <div className="space-y-4">
              <div className="space-y-2">
                <label className="text-sm text-gray-300">Full Name</label>
                <input
                  type="text"
                  placeholder="Jane Doe"
                  className="w-full rounded-xl border border-[#262626] bg-[#0f0f0f] px-4 py-3 text-sm outline-none focus:border-[#6d41ff]"
                  value={name}
                  onChange={(e) => setName(e.target.value)}
                />
              </div>
              <div className="space-y-2">
                <label className="text-sm text-gray-300">Wallet Address</label>
                <input
                  type="text"
                  placeholder="0x..."
                  className="w-full rounded-xl border border-[#262626] bg-[#0f0f0f] px-4 py-3 text-sm outline-none focus:border-[#6d41ff]"
                  value={walletAddress}
                  onChange={(e) => setWalletAddress(e.target.value)}
                />
              </div>
              <div className="space-y-2">
                <label className="text-sm text-gray-300">Email</label>
                <input
                  type="email"
                  placeholder="you@example.com"
                  className="w-full rounded-xl border border-[#262626] bg-[#0f0f0f] px-4 py-3 text-sm outline-none focus:border-[#6d41ff]"
                  value={email}
                  onChange={(e) => setEmail(e.target.value)}
                />
              </div>
              <div className="space-y-2">
                <label className="text-sm text-gray-300">Password</label>
                <input
                  type="password"
                  placeholder="••••••••"
                  className="w-full rounded-xl border border-[#262626] bg-[#0f0f0f] px-4 py-3 text-sm outline-none focus:border-[#6d41ff]"
                  value={password}
                  onChange={(e) => setPassword(e.target.value)}
                />
              </div>
              <div className="space-y-2">
                <label className="text-sm text-gray-300">Confirm Password</label>
                <input
                  type="password"
                  placeholder="••••••••"
                  className="w-full rounded-xl border border-[#262626] bg-[#0f0f0f] px-4 py-3 text-sm outline-none focus:border-[#6d41ff]"
                  value={confirmPassword}
                  onChange={(e) => setConfirmPassword(e.target.value)}
                />
              </div>
              <div className="flex items-center justify-between text-sm text-gray-400">
                <label className="flex items-center gap-2">
                  <input type="checkbox" className="accent-[#6d41ff]" />
                  I agree to the Terms &amp; Conditions
                </label>
              </div>
              <Button
                variant="primary"
                size="md"
                className="w-full bg-[#6d41ff] hover:bg-[#5b2fff] text-white disabled:opacity-50"
                onClick={handleSignup}
                disabled={loading}
              >
                {loading ? "Creating Account..." : "Create Account"}
              </Button>
            </div>

            <div className="text-center text-sm text-gray-400">
              Already have an account?{" "}
              <Link href="/login" className="text-[#8b6dff] hover:opacity-80">
                Login
              </Link>
            </div>
          </div>
        </div>
      </main>
    </div>
  );
}

export default SignupPage;


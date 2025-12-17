// Components/organisms/Navbar.tsx
import React, { useState, useEffect } from "react";
import { useLocation, useNavigate } from "react-router-dom";
import { Brand } from "../molecules/Brand";
import { NavLink } from "../molecules/NavLink";
import { HamburgerButton } from "../molecules/Hamburger";
import { MobileMenu } from "../molecules/MobileMenu";
import { Link } from "../atoms/Link";
import { Button } from "../atoms/Button";
import { useWallet } from "../../hooks/useWallet";

interface NavbarProps {
  currentPage?: string;
}

export const Navbar: React.FC<NavbarProps> = ({ currentPage = "/" }) => {
  const [isMobileMenuOpen, setIsMobileMenuOpen] = useState(false);
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [user, setUser] = useState<any>(null);
  const location = useLocation();
  const navigate = useNavigate();
  const activeHref = currentPage || location.pathname;
  const { isConnected, address, connect, disconnect, resetLoading, isLoading, chainId, switchNetwork } = useWallet();

  const isAdmin = user?.Role === "admin";

  // Check if user is logged in
  useEffect(() => {
    const checkAuth = () => {
      const auth = localStorage.getItem("auth") === "true";
      const userStr = localStorage.getItem("user");
      setIsLoggedIn(auth);
      setUser(userStr ? JSON.parse(userStr) : null);
    };

    checkAuth();
    // Listen for storage changes (e.g., when login happens in another tab)
    window.addEventListener("storage", checkAuth);
    // Listen for custom auth change event (when login happens in same tab)
    window.addEventListener("authChange", checkAuth);
    return () => {
      window.removeEventListener("storage", checkAuth);
      window.removeEventListener("authChange", checkAuth);
    };
  }, []);

  const handleLogout = () => {
    // Clear authentication data
    localStorage.removeItem("auth");
    localStorage.removeItem("token");
    localStorage.removeItem("user");

    // Disconnect wallet to force fresh connection on next login
    disconnect();

    setIsLoggedIn(false);
    setUser(null);
    navigate("/");
  };

  const navLinks = [
    { label: "Home", href: "/" },
    { label: "About Us", href: "/about" },
    { label: "Properties", href: "/properties" },
    { label: "Services", href: "/services" },
    ...(isAdmin ? [{ label: "Admin", href: "/admin" }] : []),
  ];

  return (
    <header className="bg-[#1A1A1A] fixed top-0 left-0 right-0 z-50 border-b border-[#262626]">
      <div className="max-w-7xl mx-auto px-6">
        <div className="flex items-center justify-between h-20 relative">
          {/* LEFT: Brand wrapped in Link */}
          <Link href="/" className="flex items-center gap-3 no-underline">
            <Brand
              logoSrc="src/assets/blockchain_House-removebg-preview.svg"
              companyName="Propchain"
              size="lg"
            />
          </Link>

          {/* CENTER: Desktop Navigation */}
          <NavLink
            links={navLinks}
            activeHref={activeHref}
            className="hidden md:flex"
          />

          {/* RIGHT: Desktop Login/Wallet Buttons */}
          <div className="hidden md:flex items-center gap-3">
            {/* Wallet Connect Section */}
            {isConnected ? (
              <div className="flex items-center gap-2">
                {chainId !== 11155111 && (
                  <Button
                    variant="outline"
                    size="sm"
                    onClick={switchNetwork}
                    className="border-yellow-500/50 bg-yellow-500/10 text-yellow-400 hover:bg-yellow-500/20 text-xs"
                  >
                    Switch to Sepolia
                  </Button>
                )}
                <div className="flex items-center gap-2 px-3 py-1.5 rounded-lg border border-[#262626] bg-[#1A1A1A]">
                  <div className="h-2 w-2 rounded-full bg-green-500" />
                  <span className="text-sm font-mono">
                    {address?.substring(0, 6)}...{address?.substring(38)}
                  </span>
                  <Button
                    variant="ghost"
                    size="sm"
                    onClick={disconnect}
                    className="text-red-400 hover:text-red-300 hover:bg-red-500/10 text-xs px-2 py-1"
                  >
                    ✕
                  </Button>
                </div>
              </div>
            ) : (
              <div className="flex items-center gap-2">
                <Button
                  variant="outline"
                  size="sm"
                  onClick={isLoading ? resetLoading : connect}
                  disabled={false}
                  className="border-[#262626] bg-[#1A1A1A] text-white hover:bg-[#262626]"
                >
                  {isLoading ? "Stuck? Click to clear & retry" : "Connect Wallet"}
                </Button>
                {localStorage.getItem('wallet_disconnected') === 'true' && (
                  <span className="text-xs text-yellow-400" title="Wallet was disconnected - manual reconnection required">
                    🔒
                  </span>
                )}
              </div>
            )}

            {/* Login/User Section */}
            {isLoggedIn ? (
              <div className="flex items-center gap-2">
                <span className="text-sm text-gray-300">{user?.Name || user?.Email}</span>
                <Link href="/dashboard">
                  <Button
                    variant="primary"
                    size="sm"
                    className="bg-[#6d41ff] hover:bg-[#5b2fff] text-white"
                  >
                    Dashboard
                  </Button>
                </Link>
                <Button
                  variant="outline"
                  size="sm"
                  onClick={handleLogout}
                  className="border-[#262626] bg-[#1A1A1A] text-white hover:bg-[#262626]"
                >
                  Logout
                </Button>
              </div>
            ) : (
              <Link href="/login">
                <Button
                  variant="primary"
                  size="sm"
                  className="bg-[#6d41ff] hover:bg-[#5b2fff] text-white"
                >
                  Login
                </Button>
              </Link>
            )}
          </div>

          {/* MOBILE: Hamburger Button */}
          <HamburgerButton
            isOpen={isMobileMenuOpen}
            onClick={() => setIsMobileMenuOpen(!isMobileMenuOpen)}
            className="md:hidden"
          />
        </div>

        {/* MOBILE: Dropdown Menu */}
        <MobileMenu
          links={navLinks}
          isOpen={isMobileMenuOpen}
          ctaButton={
            <div className="w-full space-y-2">
              {/* Wallet Section */}
              {isConnected ? (
                <>
                  {chainId !== 11155111 && (
                    <Button
                      variant="outline"
                      size="md"
                      onClick={switchNetwork}
                      className="border-yellow-500/50 bg-yellow-500/10 text-yellow-400 hover:bg-yellow-500/20 w-full text-xs"
                    >
                      Switch to Sepolia
                    </Button>
                  )}
                  <div className="w-full px-3 py-2 rounded-lg border border-[#262626] bg-[#1A1A1A] text-center">
                    <div className="flex items-center justify-center gap-2">
                      <div className="h-2 w-2 rounded-full bg-green-500" />
                      <span className="text-sm font-mono">
                        {address?.substring(0, 6)}...{address?.substring(38)}
                      </span>
                      <Button
                        variant="ghost"
                        size="sm"
                        onClick={disconnect}
                        className="text-red-400 hover:text-red-300 hover:bg-red-500/10 text-xs px-2 py-1 ml-2"
                      >
                        ✕
                      </Button>
                    </div>
                  </div>
                </>
              ) : (
                <div className="w-full space-y-1">
                <Button
                  variant="outline"
                  size="md"
                  onClick={isLoading ? resetLoading : connect}
                  disabled={false}
                  className="border-[#262626] bg-[#1A1A1A] text-white w-full"
                >
                  {isLoading ? "Stuck? Click to clear & retry" : "Connect Wallet"}
                </Button>
                  {localStorage.getItem('wallet_disconnected') === 'true' && (
                    <p className="text-xs text-yellow-400 text-center">
                      🔒 Wallet disconnected - requires manual reconnection
                    </p>
                  )}
                </div>
              )}

              {/* Login Section */}
              {isLoggedIn ? (
                <>
                  <Link href="/dashboard" className="w-full">
                    <Button
                      variant="primary"
                      size="md"
                      className="bg-[#6d41ff] hover:bg-[#5b2fff] text-white w-full"
                    >
                      Dashboard
                    </Button>
                  </Link>
                  <Button
                    variant="outline"
                    size="md"
                    onClick={handleLogout}
                    className="border-[#262626] bg-[#1A1A1A] text-white w-full"
                  >
                    Logout
                  </Button>
                </>
              ) : (
                <Link href="/login" className="w-full">
                  <Button
                    variant="primary"
                    size="md"
                    className="bg-[#6d41ff] hover:bg-[#5b2fff] text-white w-full"
                  >
                    Login
                  </Button>
                </Link>
              )}
            </div>
          }
          className="bg-[#1A1A1A]"
        />
      </div>
    </header>
  );
};

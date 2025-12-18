import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { twMerge } from "tailwind-merge";
import {
  Home,
  Building2,
  Upload,
  TrendingUp,
  Wallet,
  UserCog,
  LogOut,
  Menu,
  X,
} from "lucide-react";
import { SidebarItem } from "../molecules/SidebarItem";
import { Brand } from "../molecules/Brand";
import { Link } from "../atoms/Link";
import logo from "../../assets/blockchain_House-removebg-preview.svg";

const navItems = [
  { icon: Home, label: "Home", href: "/dashboard" },
  { icon: Building2, label: "View Properties", href: "/dashboard/properties" },
  { icon: Upload, label: "Upload Property", href: "/dashboard/upload" },
  { icon: TrendingUp, label: "My Investments", href: "/dashboard/investments" },
  { icon: Wallet, label: "Wallet", href: "/dashboard/wallet" },
  { icon: UserCog, label: "My Account", href: "/dashboard/account" },
];

interface SidebarProps {
  className?: string;
}

export const Sidebar: React.FC<SidebarProps> = ({ className }) => {
  const [isMobileOpen, setIsMobileOpen] = useState(false);
  const navigate = useNavigate();

  const handleLogout = () => {
    localStorage.removeItem("auth");
    localStorage.removeItem("token");
    navigate("/login");
  };

  const closeMobileMenu = () => {
    setIsMobileOpen(false);
  };

  const sidebarContent = (
    <>
      {/* Brand */}
      <div className="p-4 border-b border-[#1f1f1f]">
        <Link href="/" className="flex items-center gap-3">
          <Brand logoSrc={logo} companyName="Propchain" size="md" />
        </Link>
      </div>

      {/* Navigation */}
      <nav className="flex-1 p-4 space-y-2">
        {navItems.map((item) => (
          <SidebarItem
            key={item.href}
            icon={item.icon}
            label={item.label}
            href={item.href}
            onClick={closeMobileMenu}
          />
        ))}
      </nav>

      {/* Logout */}
      <div className="p-4 border-t border-[#1f1f1f]">
        <button
          onClick={handleLogout}
          className="flex items-center gap-3 px-4 py-3 rounded-xl transition-all text-sm font-medium text-gray-400 hover:bg-[#1a1a1a] hover:text-white w-full"
        >
          <LogOut size={20} />
          <span>Logout</span>
        </button>
      </div>
    </>
  );

  return (
    <>
      {/* Mobile hamburger button */}
      <button
        onClick={() => setIsMobileOpen(!isMobileOpen)}
        className="md:hidden fixed top-4 left-4 z-50 p-2 rounded-lg bg-[#111111] border border-[#1f1f1f] text-white"
      >
        {isMobileOpen ? <X size={24} /> : <Menu size={24} />}
      </button>

      {/* Mobile overlay */}
      {isMobileOpen && (
        <div
          className="md:hidden fixed inset-0 bg-black/60 z-40"
          onClick={closeMobileMenu}
        />
      )}

      {/* Mobile sidebar */}
      <aside
        className={twMerge(
          "md:hidden fixed top-0 left-0 h-full w-64 bg-[#111111] border-r border-[#1f1f1f] z-50 flex flex-col transition-transform duration-300",
          isMobileOpen ? "translate-x-0" : "-translate-x-full"
        )}
      >
        {sidebarContent}
      </aside>

      {/* Desktop sidebar */}
      <aside
        className={twMerge(
          "hidden md:flex fixed top-0 left-0 h-full w-64 bg-[#111111] border-r border-[#1f1f1f] flex-col z-40",
          className
        )}
      >
        {sidebarContent}
      </aside>
    </>
  );
};

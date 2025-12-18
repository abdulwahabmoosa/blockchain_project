import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { twMerge } from "tailwind-merge";
import {
  LayoutDashboard,
  Users,
  Plus,
  Wallet,
  LogOut,
  Menu,
  X,
  Shield,
} from "lucide-react";
import { SidebarItem } from "../molecules/SidebarItem";
import { Brand } from "../molecules/Brand";
import { Link } from "../atoms/Link";
import logo from "../../assets/blockchain_House-removebg-preview.svg";

const navItems = [
  { icon: LayoutDashboard, label: "Dashboard", href: "/admin" },
  { icon: Plus, label: "Create Property", href: "/admin/create-property" },
  { icon: Users, label: "Manage Users", href: "/admin/users" },
  { icon: Wallet, label: "Wallet", href: "/admin/wallet" },
];

interface AdminSidebarProps {
  className?: string;
  onLogout?: () => void;
}

export const AdminSidebar: React.FC<AdminSidebarProps> = ({
  className,
  onLogout,
}) => {
  const [isMobileOpen, setIsMobileOpen] = useState(false);
  const navigate = useNavigate();

  const closeMobileMenu = () => {
    setIsMobileOpen(false);
  };

  const handleLogoutClick = () => {
    if (onLogout) {
      onLogout();
    } else {
      // Fallback that mirrors existing logout semantics
      localStorage.removeItem("auth");
      localStorage.removeItem("token");
      localStorage.removeItem("user");
      navigate("/login");
    }
    closeMobileMenu();
  };

  const sidebarContent = (
    <>
      {/* Brand */}
      <div className="p-4 border-b border-[#1f1f1f]">
        <Link href="/" className="flex items-center gap-3">
          <Brand logoSrc={logo} companyName="Propchain" size="md" />
        </Link>
        <div className="flex items-center gap-2 mt-3 px-2">
          <Shield size={14} className="text-[#703BF7]" />
          <span className="text-xs text-[#703BF7] font-medium">Admin Panel</span>
        </div>
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
          onClick={handleLogoutClick}
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



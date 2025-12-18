import { Link, useLocation } from "react-router-dom";
import { twMerge } from "tailwind-merge";
import type { LucideIcon } from "lucide-react";

interface SidebarItemProps {
  icon: LucideIcon;
  label: string;
  href: string;
  onClick?: () => void;
}

export const SidebarItem: React.FC<SidebarItemProps> = ({
  icon: Icon,
  label,
  href,
  onClick,
}) => {
  const location = useLocation();

  // For root paths like /dashboard or /admin, only match exactly
  // For sub-paths like /dashboard/properties, use startsWith
  const isRootPath = href === "/dashboard" || href === "/admin";
  const isActive = isRootPath
    ? location.pathname === href
    : location.pathname.startsWith(href);

  return (
    <Link
      to={href}
      onClick={onClick}
      className={twMerge(
        "flex items-center gap-3 px-4 py-3 rounded-xl transition-all text-sm font-medium",
        isActive
          ? "bg-[#703BF7] text-white"
          : "text-gray-400 hover:bg-[#1a1a1a] hover:text-white"
      )}
    >
      <Icon size={20} />
      <span>{label}</span>
    </Link>
  );
};



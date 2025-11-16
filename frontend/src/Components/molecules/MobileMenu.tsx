import React from "react";
import { Link } from "../atoms/Link";
import { twMerge } from "tailwind-merge";

interface MobileMenuProps {
  links: Array<{ label: string; href: string }>;
  isOpen: boolean;
  ctaButton?: React.ReactNode;
  className?: string;
}

export const MobileMenu: React.FC<MobileMenuProps> = ({
  links,
  isOpen,
  ctaButton,
  className = "",
}) => {
  if (!isOpen) return null;

  return (
    <div
      className={twMerge(
        "absolute top-full left-0 w-full bg-gray-900 border-t border-gray-800 md:hidden",
        className
      )}
    >
      <nav className="flex flex-col p-4 gap-4">
        {links.map((link, index) => (
          <Link
            key={index}
            href={link.href}
            className="text-white text-base font-medium py-2"
          >
            {link.label}
          </Link>
        ))}

        {ctaButton && <div className="pt-2">{ctaButton}</div>}
      </nav>
    </div>
  );
};

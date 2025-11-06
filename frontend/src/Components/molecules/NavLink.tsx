import React from "react";
import { twMerge } from "tailwind-merge";
import { Link } from "../atoms/Link";

interface NavLinksProps {
  links: Array<{ label: string; href: string }>;
  className?: string;
}

export const NavLink: React.FC<NavLinksProps> = ({ links, className = "" }) => {
  return (
    <nav className={twMerge("flex items-center gap-8", className)}>
      {links.map((link, index) => (
        <Link
          key={index}
          href={link.href}
          className="text-white text-base font-medium"
        >
          {link.label}
        </Link>
      ))}
    </nav>
  );
};

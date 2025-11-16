// Components/molecules/NavLinks.tsx
import React from "react";
import { twMerge } from "tailwind-merge";
import { Link } from "../atoms/Link";

interface NavLinksProps {
  links: Array<{ label: string; href: string }>;
  className?: string;
  activeHref?: string;
}

export const NavLink: React.FC<NavLinksProps> = ({
  links,
  className = "",
  activeHref,
}) => {
  return (
    <nav className={twMerge("flex items-center gap-2", className)}>
      {links.map((link, index) => {
        const isActive = activeHref === link.href;

        return (
          <Link
            key={index}
            href={link.href}
            className={twMerge(
              "text-white text-base font-medium px-5 py-2.5 rounded-lg transition-all",
              isActive
                ? "bg-[#703BF7] hover:bg-[#5F32D4]"
                : "hover:bg-[#1A1A1A]"
            )}
          >
            {link.label}
          </Link>
        );
      })}
    </nav>
  );
};

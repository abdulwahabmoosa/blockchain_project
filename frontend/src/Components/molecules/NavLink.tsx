import React from "react";
import { twMerge } from "tailwind-merge";
import { Text } from "../atoms/Text";
import { Link } from "../atoms/Link";

interface NavLinkProps {
  href: string;
  label: string;
  isActive?: boolean;
  className?: string;
}

export const NavLink: React.FC<NavLinkProps> = ({
  href,
  label,
  isActive = false,
  className,
}) => {
  return (
    <Link
      href={href}
      className={twMerge(
        "px-3 py-2 transition-opacity duration-200",
        className
      )}
    >
      <Text
        as="span"
        size="md"
        className={twMerge(
          isActive
            ? "text-white font-semibold"
            : "text-gray-300 hover:text-white opacity-75 hover:opacity-100"
        )}
      >
        {label}
      </Text>
    </Link>
  );
};

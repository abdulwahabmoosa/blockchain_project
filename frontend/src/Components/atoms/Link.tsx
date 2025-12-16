import React from "react";
import { Link as RouterLink } from "react-router-dom";
import { twMerge } from "tailwind-merge";

interface LinkProps {
  children: React.ReactNode;
  href: string;
  className?: string;
  external?: boolean;
}

export const Link: React.FC<LinkProps> = ({
  children,
  href,
  className = "",
  external = false,
}) => {
  const classes = twMerge(
    "text-inherit hover:opacity-80 transition-opacity",
    className
  );

  if (external) {
    return (
      <a
        href={href}
        className={classes}
        target="_blank"
        rel="noopener noreferrer"
      >
        {children}
      </a>
    );
  }

  return (
    <RouterLink to={href} className={classes}>
      {children}
    </RouterLink>
  );
};

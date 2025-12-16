import React from "react";
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

  return (
    <a
      href={href}
      className={classes}
      target={external ? "_blank" : undefined}
      rel={external ? "noopener noreferrer" : undefined}
    >
      {children}
    </a>
  );
};

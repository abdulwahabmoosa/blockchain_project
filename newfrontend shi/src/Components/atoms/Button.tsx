import React from "react";
import { twMerge } from "tailwind-merge";

interface ButtonProps {
  children: React.ReactNode;
  as?: "button" | "a";
  onClick?: () => void;
  className?: string;
  color?: string;
  variant?: "primary" | "secondary" | "outline" | "ghost" | "danger";
  href?: string;
  size?: "sm" | "md" | "lg";
  type?: "button" | "submit" | "reset";
  rounding?: "none" | "sm" | "md" | "lg" | "full";
  disabled?: boolean;
}

export const Button: React.FC<ButtonProps> = ({
  children,
  onClick,
  className = "",
  color,
  size = "lg",
  href,
  type = "button",
  variant = "primary",
  rounding = "md",
  disabled = false,
  as: Component = "button",
}) => {
  const variantClasses = {
    primary: "bg-blue-600 text-white hover:bg-blue-700 active:bg-blue-800",
    secondary: "bg-gray-600 text-white hover:bg-gray-700 active:bg-gray-800",
    outline:
      "border-2 border-blue-600 text-blue-600 hover:bg-blue-50 active:bg-blue-100",
    ghost: "text-blue-600 hover:bg-blue-50 active:bg-blue-100",
    danger: "bg-red-600 text-white hover:bg-red-700 active:bg-red-800",
  };

  const sizeClasses = {
    sm: "px-2 py-1 text-sm",
    md: "px-4 py-2 text-base",
    lg: "px-6 py-3 text-lg",
  };

  const roundingClasses = {
    none: "rounded-none",
    sm: "rounded-md",
    md: "rounded-xl",
    lg: "rounded-3xl",
    full: "rounded-full",
  };

  const classes = twMerge(
    variantClasses[variant],
    sizeClasses[size],
    roundingClasses[rounding],
    disabled && "opacity-50 cursor-not-allowed pointer-events-none",
    className
  );

  if (Component === "a" && href) {
    return (
      <a href={href} className={classes}>
        {children}
      </a>
    );
  }

  return (
    <Component
      className={classes}
      onClick={onClick}
      style={color ? { color } : undefined}
      type={type}
      disabled={disabled}
    >
      {children}
    </Component>
  );
};

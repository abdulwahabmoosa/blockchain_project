import React from "react";
import { twMerge } from "tailwind-merge";

interface TextProps {
  children: React.ReactNode;
  className?: string;
  as?: "p" | "span" | "div" | "label";
  size?: "sm" | "md" | "lg";
  weight?: "normal" | "medium" | "semibold" | "bold";
  color?: string;
}

export const Text: React.FC<TextProps> = ({
  children,
  className = "",
  as: Component = "p",
  size = "md",
  weight = "normal",
  color,
}) => {
  const sizeClasses = {
    sm: "text-sm",
    md: "text-base",
    lg: "text-lg",
  };

  const weightClasses = {
    normal: "font-normal",
    medium: "font-medium",
    semibold: "font-semibold",
    bold: "font-bold",
  };

  const classes = twMerge(sizeClasses[size], weightClasses[weight], className);

  return (
    <Component className={classes} style={color ? { color } : undefined}>
      {children}
    </Component>
  );
};

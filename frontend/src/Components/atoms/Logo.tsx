import React from "react";
import { twMerge } from "tailwind-merge";

interface LogoProps {
  className?: string;
  size?: "sm" | "md" | "lg";
  src?: string;
  alt?: string;
}

export const Logo: React.FC<LogoProps> = ({
  className = "",
  size = "md",
  src,
  alt = "Src missing",
}) => {
  const sizeClasses = {
    sm: "h-6 w-6",
    md: "h-10 w-10",
    lg: "h-16 w-16",
  };
  const classes = twMerge(sizeClasses[size], className);

  return (
    <div className={classes}>
      {src ? (
        <img src={src} alt={alt} className="w-full h-full object-contain" />
      ) : (
        alt
      )}
    </div>
  );
};

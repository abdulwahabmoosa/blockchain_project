import React from "react";
import { twMerge } from "tailwind-merge";

interface HamburgerButtonProps {
  isOpen: boolean;
  onClick: () => void;
  className?: string;
}

export const HamburgerButton: React.FC<HamburgerButtonProps> = ({
  isOpen,
  onClick,
  className = "",
}) => {
  return (
    <button
      onClick={onClick}
      className={twMerge(
        "flex flex-col gap-1 p-2 bg-transparent border-none cursor-pointer",
        className
      )}
      aria-label="Toggle menu"
    >
      <span
        className={twMerge(
          "block w-6 h-0.5 bg-white transition-transform",
          isOpen && "rotate-45 translate-y-1.5"
        )}
      />
      <span
        className={twMerge(
          "block w-6 h-0.5 bg-white transition-opacity",
          isOpen && "opacity-0"
        )}
      />
      <span
        className={twMerge(
          "block w-6 h-0.5 bg-white transition-transform",
          isOpen && "-rotate-45 -translate-y-1.5"
        )}
      />
    </button>
  );
};

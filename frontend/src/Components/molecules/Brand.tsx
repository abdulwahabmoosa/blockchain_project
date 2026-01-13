// Components/molecules/Brand.tsx
import React from "react";
import { Logo } from "../atoms/Logo";
import { Text } from "../atoms/Text";

interface BrandProps {
  logoSrc: string;
  companyName: string;
  size?: "sm" | "md" | "lg";
}

export const Brand: React.FC<BrandProps> = ({ logoSrc, companyName, size }) => {
  return (
    <>
      <Logo src={logoSrc} alt={companyName} size={size} />
      <Text as="span" size="lg" weight="semibold" className="text-white">
        {companyName}
      </Text>
    </>
  );
};

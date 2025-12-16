// Components/organisms/Navbar.tsx
import React, { useState } from "react";
import { useLocation } from "react-router-dom";
import { Brand } from "../molecules/Brand";
import { NavLink } from "../molecules/NavLink";
import { HamburgerButton } from "../molecules/Hamburger";
import { MobileMenu } from "../molecules/MobileMenu";
import { Link } from "../atoms/Link";
import { Button } from "../atoms/Button";

interface NavbarProps {
  currentPage?: string;
}

export const Navbar: React.FC<NavbarProps> = ({ currentPage = "/" }) => {
  const [isMobileMenuOpen, setIsMobileMenuOpen] = useState(false);
  const location = useLocation();
  const activeHref = currentPage || location.pathname;

  const navLinks = [
    { label: "Home", href: "/" },
    { label: "About Us", href: "/about" },
    { label: "Properties", href: "/properties" },
    { label: "Services", href: "/services" },
  ];

  return (
    <header className="bg-[#1A1A1A] fixed top-0 left-0 right-0 z-50 border-b border-[#262626]">
      <div className="max-w-7xl mx-auto px-6">
        <div className="flex items-center justify-between h-20 relative">
          {/* LEFT: Brand wrapped in Link */}
          <Link href="/" className="flex items-center gap-3 no-underline">
            <Brand
              logoSrc="src/assets/blockchain_House-removebg-preview.svg"
              companyName="Propchain"
              size="lg"
            />
          </Link>

          {/* CENTER: Desktop Navigation */}
          <NavLink
            links={navLinks}
            activeHref={activeHref}
            className="hidden md:flex"
          />

          {/* RIGHT: Desktop CTA Button */}
          <div className="hidden md:block">
            <Link href="/login">
              <Button
                variant="outline"
                size="md"
                className="border-[#262626] bg-[#1A1A1A] text-white hover:bg-[#262626]"
              >
                Login
              </Button>
            </Link>
          </div>

          {/* MOBILE: Hamburger Button */}
          <HamburgerButton
            isOpen={isMobileMenuOpen}
            onClick={() => setIsMobileMenuOpen(!isMobileMenuOpen)}
            className="md:hidden"
          />
        </div>

        {/* MOBILE: Dropdown Menu */}
        <MobileMenu
          links={navLinks}
          isOpen={isMobileMenuOpen}
          ctaButton={
              <Link href="/login" className="w-full">
                <Button
                  variant="outline"
                  size="md"
                  className="border-[#262626] bg-[#1A1A1A] text-white  w-full"
                >
                  Login
                </Button>
              </Link>
          }
          className="bg-[#1A1A1A]"
        />
      </div>
    </header>
  );
};

import "../index.css";
import { Text } from "../Components/atoms/Text";
import { Button } from "../Components/atoms/Button";
import { Logo } from "../Components/atoms/Logo";
import { NavLink } from "../Components/molecules/NavLink";

function LandingPage() {
  return (
    <div className="flex h-screen w-screen items-center justify-center">
      <div className="text-center space-y-4">
        <nav className="flex gap-6 bg-blue-900 text-white px-6 py-2 rounded">
          <NavLink href="/" label="Home" isActive />
          <NavLink href="/about" label="About" />
          <NavLink href="/pricing" label="Pricing" />
          <NavLink href="/contact" label="Contact" />
        </nav>

        <Text as="p" size="lg" className="text-9xl text-gray-900">
          Hello world
        </Text>
      </div>
      <div className="flex flex-col gap-4 p-6 ml-10">
        <Button variant="primary" rounding="sm" className="">
          Get Started
        </Button>
        <Button variant="secondary" className="mt-4 rounded-full">
          Learn More
        </Button>
        <Button variant="outline" rounding="md" className="mt-4">
          Contact Us
        </Button>
        <Button variant="ghost" className="mt-4">
          Help
        </Button>
        <Button variant="danger" rounding="lg" className="mt-4">
          Delete Account
        </Button>
      </div>

      <div className="flex gap-6 items-center justify-center mt-10">
        {/* 1️⃣ src provided → renders <img> */}
        <Logo src="/images/oceanic-logo.svg" size="lg" />
        <Logo size="md" />
        <Logo size="sm" />
      </div>

      <div className="flex flex-col gap-4 p-6"></div>
    </div>
  );
}

export default LandingPage;

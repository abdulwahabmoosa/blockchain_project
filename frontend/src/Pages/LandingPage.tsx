import "../index.css";
import { Text } from "../Components/atoms/Text";
import { Button } from "../Components/atoms/Button";
import { Logo } from "../Components/atoms/Logo";
import { Navbar } from "../Components/organisms/Navbar";

function LandingPage() {
  return (
    <div className="min-h-screen w-screen bg-gray-100">
      {/* Navbar Organism */}
      <Navbar currentPage="/about" />

      {/* Main Content */}
      <main className="flex flex-col items-center justify-center pt-24 px-4">
        {/* Hero Section */}
        <div className="text-center space-y-4 mb-12">
          <Text as="p" size="sm" className="text-9xl text-gray-900">
            Hello
          </Text>
        </div>

        {/* Buttons Section */}
        <div className="flex flex-col gap-4 mb-12">
          <Button variant="primary" rounding="sm">
            Get Started
          </Button>
          <Button variant="secondary" rounding="full">
            Learn More
          </Button>
          <Button variant="outline" rounding="md">
            Contact Us
          </Button>
          <Button variant="ghost">Help</Button>
          <Button variant="danger" rounding="lg">
            Delete Account
          </Button>
        </div>

        {/* Logos Section */}
        <div className="flex gap-6 items-center justify-center">
          <Logo src="/images/oceanic-logo.svg" size="lg" />
          <Logo size="md" />
          <Logo size="sm" />
        </div>
      </main>
    </div>
  );
}

export default LandingPage;

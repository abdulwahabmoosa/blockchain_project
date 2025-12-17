import { Navbar } from "../Components/organisms/Navbar";

function ServicesPage() {
  return (
    <div className="min-h-screen bg-[#0f0f0f] text-white">
      <Navbar currentPage="/services" />

      <main className="max-w-7xl mx-auto px-6 pt-28 pb-16">
        <div className="text-center space-y-4">
          <h1 className="text-4xl font-bold">Our Services</h1>
          <p className="text-gray-400 max-w-2xl mx-auto">
            Professional real estate tokenization and investment management services.
          </p>
        </div>

        <div className="mt-16 grid md:grid-cols-3 gap-8">
          <div className="bg-[#1A1A1A] border border-[#262626] rounded-lg p-6">
            <h3 className="text-xl font-semibold mb-4">Property Tokenization</h3>
            <p className="text-gray-400">
              Convert real estate assets into blockchain-based tokens for fractional ownership and investment.
            </p>
          </div>

          <div className="bg-[#1A1A1A] border border-[#262626] rounded-lg p-6">
            <h3 className="text-xl font-semibold mb-4">Revenue Distribution</h3>
            <p className="text-gray-400">
              Automated distribution of rental income and property returns to token holders.
            </p>
          </div>

          <div className="bg-[#1A1A1A] border border-[#262626] rounded-lg p-6">
            <h3 className="text-xl font-semibold mb-4">Investment Management</h3>
            <p className="text-gray-400">
              Professional management of tokenized property portfolios with real-time tracking.
            </p>
          </div>
        </div>
      </main>
    </div>
  );
}

export default ServicesPage;



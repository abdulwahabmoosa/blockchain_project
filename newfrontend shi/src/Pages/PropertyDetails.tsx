import "../index.css";
import { useEffect, useState } from "react";
import { useParams, Link as RouterLink } from "react-router-dom";
import { Navbar } from "../Components/organisms/Navbar";
import { Button } from "../Components/atoms/Button";
import { api } from "../lib/api";
import type { Property } from "../types";

const GradientThumb = () => (
  <div className="w-full h-full bg-linear-to-br from-[#1f1f1f] via-[#2a2a2a] to-[#111111]" />
);

export default function PropertyDetailsPage() {
  const { id } = useParams<{ id: string }>();
  const [property, setProperty] = useState<Property | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");

  useEffect(() => {
    const fetchProperty = async () => {
      if (!id) return;
      try {
        const data = await api.getProperty(id);
        setProperty(data);
      } catch (err: any) {
        console.error("Failed to fetch property:", err);
        setError("Failed to load property details");
      } finally {
        setLoading(false);
      }
    };
    fetchProperty();
  }, [id]);

  if (loading) {
    return (
      <div className="min-h-screen bg-[#0f0f0f] text-white flex items-center justify-center">
        <Navbar currentPage="/properties" />
        <p>Loading...</p>
      </div>
    );
  }

  if (error || !property) {
    return (
      <div className="min-h-screen bg-[#0f0f0f] text-white flex items-center justify-center">
        <Navbar currentPage="/properties" />
        <div className="text-center">
            <h1 className="text-2xl font-bold mb-4">Property Not Found</h1>
            <p className="text-gray-400 mb-6">{error || "The property you are looking for does not exist."}</p>
            <RouterLink to="/dashboard">
                <Button>Go to Dashboard</Button>
            </RouterLink>
        </div>
      </div>
    );
  }

  // Placeholder data for missing fields
  const title = `Property #${property.ID.substring(0, 8)}`;
  const description = "No description available for this property.";
  const features = ["Tokenized Asset", "Blockchain Verified", "Instant Settlement"];
  const amenities = [
    `Valuation: ${property.Valuation} ETH`,
    `Status: ${property.Status}`,
    `Created: ${new Date(property.CreatedAt).toLocaleDateString()}`,
    `Owner: ${property.OwnerWallet.substring(0, 6)}...${property.OwnerWallet.substring(38)}`,
  ];

  return (
    <div className="min-h-screen w-full bg-[#0f0f0f] text-white font-sans selection:bg-[#6d41ff] selection:text-white">
      <Navbar currentPage="/properties" />

      <main className="max-w-7xl mx-auto px-4 pt-28 pb-16 space-y-10">
        {/* Header Section */}
        <section className="space-y-6">
          <div className="flex flex-col md:flex-row md:items-end justify-between gap-4">
            <div className="space-y-2">
              <h1 className="text-3xl md:text-4xl font-semibold">{title}</h1>
              <div className="flex items-center gap-2 text-gray-400 text-sm">
                <span className="bg-[#1f1f1f] px-2 py-1 rounded-md text-white/80">
                  {property.Status}
                </span>
                <span>{property.OnchainAssetAddress}</span>
              </div>
            </div>
            <div className="text-left md:text-right space-y-1">
              <p className="text-gray-400 text-sm">Property Value</p>
              <p className="text-3xl font-semibold">{property.Valuation} ETH</p>
            </div>
          </div>

          {/* Image Gallery */}
          <div className="grid grid-cols-1 md:grid-cols-4 gap-4 h-[400px] md:h-[500px]">
            <div className="md:col-span-2 md:row-span-2 rounded-3xl overflow-hidden border border-[#1f1f1f] relative group">
               <GradientThumb />
               <div className="absolute inset-0 flex items-center justify-center pointer-events-none">
                  <span className="bg-black/50 px-4 py-2 rounded-lg backdrop-blur-sm">Main Image</span>
               </div>
            </div>
            <div className="rounded-3xl overflow-hidden border border-[#1f1f1f] relative group">
               <GradientThumb />
            </div>
            <div className="rounded-3xl overflow-hidden border border-[#1f1f1f] relative group">
               <GradientThumb />
            </div>
             <div className="rounded-3xl overflow-hidden border border-[#1f1f1f] relative group">
               <GradientThumb />
            </div>
             <div className="rounded-3xl overflow-hidden border border-[#1f1f1f] relative group">
               <GradientThumb />
            </div>
          </div>
        </section>

        {/* Details Grid */}
        <section className="grid md:grid-cols-3 gap-8">
          {/* Left Column: Description & Features */}
          <div className="md:col-span-2 space-y-8">
            <div className="space-y-4 rounded-3xl border border-[#1f1f1f] bg-[#111111] p-6">
              <h2 className="text-xl font-semibold">Description</h2>
              <p className="text-gray-300 leading-relaxed">
                {description}
              </p>
            </div>

            <div className="space-y-4 rounded-3xl border border-[#1f1f1f] bg-[#111111] p-6">
              <h2 className="text-xl font-semibold">Key Features</h2>
              <ul className="grid sm:grid-cols-2 gap-3">
                {features.map((feature, index) => (
                  <li key={index} className="flex items-center gap-2 text-gray-300">
                    <span className="h-1.5 w-1.5 rounded-full bg-[#6d41ff]" />
                    {feature}
                  </li>
                ))}
              </ul>
            </div>
            
             <div className="space-y-4 rounded-3xl border border-[#1f1f1f] bg-[#111111] p-6">
              <h2 className="text-xl font-semibold">Details & Stats</h2>
              <div className="grid sm:grid-cols-2 gap-y-4 gap-x-8">
                {amenities.map((item, index) => (
                   <div key={index} className="flex justify-between border-b border-[#262626] pb-2 last:border-0">
                      <span className="text-gray-400 text-sm">{item.split(":")[0]}</span>
                      <span className="font-medium text-sm">{item.split(":")[1]}</span>
                   </div>
                ))}
              </div>
            </div>
          </div>

          {/* Right Column: Investment Card */}
          <div className="space-y-6">
            <div className="rounded-3xl border border-[#1f1f1f] bg-[#111111] p-6 space-y-6 sticky top-28">
              <div className="space-y-2">
                <p className="text-sm text-gray-400">Current Valuation</p>
                <div className="flex items-baseline gap-2">
                  <span className="text-3xl font-semibold">{property.Valuation} ETH</span>
                </div>
              </div>

              <div className="h-px bg-[#262626]" />

              <div className="space-y-3">
                <div className="flex justify-between text-sm">
                   <span className="text-gray-400">Min. Investment</span>
                   <span>0.01 ETH</span>
                </div>
                 <div className="flex justify-between text-sm">
                   <span className="text-gray-400">Total Tokens</span>
                   <span>--</span>
                </div>
                 <div className="flex justify-between text-sm">
                   <span className="text-gray-400">Available</span>
                   <span>--</span>
                </div>
              </div>

              <Button size="lg" className="w-full bg-[#6d41ff] hover:bg-[#5b36d6]">
                Invest Now
              </Button>
               <p className="text-xs text-center text-gray-500">
                Investment is powered by smart contracts.
              </p>
            </div>
          </div>
        </section>
      </main>
    </div>
  );
}

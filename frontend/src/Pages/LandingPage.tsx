import "../index.css";
import { useRef } from "react";
import { Link } from "../Components/atoms/Link";
import { Button } from "../Components/atoms/Button";
import { Navbar } from "../Components/organisms/Navbar";

const stats = [
  { label: "Happy Customers", value: "200+" },
  { label: "Properties For Clients", value: "10k+" },
  { label: "Years of Experience", value: "16+" },
];

const featureCards = [
  { title: "Find Your Dream Home" },
  { title: "Unlock Property Value" },
  { title: "Effortless Property Management" },
  { title: "Smart Investments, Informed Decisions" },
];

const properties = [
  {
    id: "seaside-serenity-villa",
    title: "Seaside Serenity Villa",
    location: "Malibu, California",
    description:
      "A stunning 4-bedroom, 3-bathroom villa in a peaceful suburban neighborhood... Read More",
    specs: ["4-Bedroom", "3-Bathroom", "Villa"],
    price: "$1,250,000",
  },
  {
    id: "metropolitan-haven",
    title: "Metropolitan Haven",
    location: "New York, New York",
    description:
      "A chic and fully-furnished 2-bedroom apartment with panoramic city views... Read More",
    specs: ["2-Bedroom", "2-Bathroom", "Apartment"],
    price: "$950,000",
  },
  {
    id: "rustic-retreat-cottage",
    title: "Rustic Retreat Cottage",
    location: "Austin, Texas",
    description:
      "An elegant 3-bedroom, 2.5-bathroom townhouse in a gated community... Read More",
    specs: ["3-Bedroom", "3-Bathroom", "Townhouse"],
    price: "$780,000",
  },
];

const testimonials = [
  {
    title: "Exceptional Service!",
    body: "Our experience with Propchain was outstanding. Their team's dedication and professionalism made finding our dream home a breeze. Highly recommended!",
    name: "Wade Warren",
    location: "USA, California",
  },
  {
    title: "Efficient and Reliable",
    body: "Propchain provided us with top-notch service. They helped us sell our property quickly and at a great price. We couldn't be happier with the results.",
    name: "Emelie Thomson",
    location: "USA, Florida",
  },
  {
    title: "Trusted Advisors",
    body: "The Propchain team guided us through the entire buying process. Their knowledge and commitment to our needs were impressive. Thank you for your support!",
    name: "John Mans",
    location: "USA, Nevada",
  },
];

const faqs = [
  {
    question: "How do I search for properties on Propchain?",
    answer:
      "Learn how to use our user-friendly search tools to find properties that match your criteria.",
  },
  {
    question: "What documents do I need to sell my property through Propchain?",
    answer:
      "Find out about the necessary documentation for listing your property with us.",
  },
  {
    question: "How can I contact a Propchain agent?",
    answer:
      "Discover the different ways you can get in touch with our experienced agents.",
  },
];

const socialIcons = ["facebook", "linkedin", "twitter", "youtube"];

function scrollCarousel(
  ref: React.RefObject<HTMLDivElement | null>,
  direction: 1 | -1
) {
  if (!ref.current) return;
  const amount = ref.current.clientWidth * 0.9 * direction;
  ref.current.scrollBy({ left: amount, behavior: "smooth" });
}

const PlaceholderImage = () => (
  <div className="w-full h-full rounded-2xl bg-linear-to-br from-[#1f1f1f] via-[#262626] to-[#1a1a1a] border border-[#262626]" />
);

function LandingPage() {
  const propertyRef = useRef<HTMLDivElement>(null);
  const testimonialRef = useRef<HTMLDivElement>(null);
  const faqRef = useRef<HTMLDivElement>(null);

  return (
    <div className="min-h-screen w-full bg-[#0e0e0e] text-white">
      <Navbar currentPage="/" />

      <main className="max-w-6xl mx-auto px-4 pt-28 pb-16 space-y-16">
        {/* Hero */}
        <section className="grid md:grid-cols-2 gap-10 items-center">
          <div className="space-y-6">
            <p className="uppercase tracking-[0.25em] text-xs text-gray-400">
              Real Estate Tokenization
            </p>
            <h1 className="text-4xl md:text-5xl font-semibold leading-tight">
              Discover Your Dream Property with Propchain
            </h1>
            <p className="text-gray-300 leading-relaxed">
              Your journey to finding the perfect property begins here. Explore
              our listings to find the home that matches your dreams.
            </p>
            <div className="flex gap-3">
              <Button
                variant="outline"
                size="md"
                className="bg-[#1a1a1a] border-[#262626] text-white hover:bg-[#262626]"
              >
                Learn More
              </Button>
              <Button
                variant="primary"
                size="md"
                className="bg-[#6d41ff] hover:bg-[#5b2fff] text-white"
              >
                Browse Properties
              </Button>
            </div>
            <div className="grid grid-cols-3 gap-3 pt-4">
              {stats.map((stat) => (
                <div
                  key={stat.label}
                  className="rounded-2xl border border-[#1f1f1f] bg-[#111111] px-4 py-4"
                >
                  <div className="text-2xl font-semibold">{stat.value}</div>
                  <p className="text-sm text-gray-400">{stat.label}</p>
                </div>
              ))}
            </div>
          </div>

          <div className="relative">
            <div className="overflow-hidden rounded-3xl border border-[#262626] bg-[#0a0a0a]">
              <PlaceholderImage />
            </div>
            <div className="absolute -left-10 top-1/2 -translate-y-1/2 hidden lg:block">
              <div className="h-36 w-36 border border-[#262626] rounded-full flex items-center justify-center text-xs uppercase tracking-[0.25em]">
                Propchain
              </div>
            </div>
          </div>
        </section>

        {/* Feature tiles */}
        <section className="grid md:grid-cols-4 gap-4">
          {featureCards.map((card) => (
            <div
              key={card.title}
              className="rounded-2xl border border-[#1f1f1f] bg-[#111111] p-5 flex flex-col gap-4"
            >
              <div className="h-12 w-12 rounded-full border border-[#262626] bg-[#0e0e0e] flex items-center justify-center text-lg text-gray-300">
                ↗
              </div>
              <p className="text-lg text-gray-100">{card.title}</p>
              <div className="h-28 rounded-xl border border-[#1f1f1f] bg-linear-to-br from-[#161616] via-[#1d1d1d] to-[#161616]" />
            </div>
          ))}
        </section>

        {/* Featured Properties */}
        <section className="space-y-6">
          <div className="flex items-center justify-between">
            <div>
              <h2 className="text-3xl font-semibold">Featured Properties</h2>
              <p className="text-sm text-gray-400">
                Explore our handpicked selection of featured properties.
              </p>
            </div>
              <Button
                variant="outline"
                size="md"
                className="border-[#262626] bg-[#111111] text-white hover:bg-[#1a1a1a]"
              >
                View All Properties
              </Button>
          </div>

          <div className="relative">
            <div
              ref={propertyRef}
              className="flex gap-4 overflow-x-auto scroll-smooth pb-2"
            >
              {properties.map((property) => (
                  <div
                    key={property.title}
                    className="min-w-[320px] rounded-2xl border border-[#1f1f1f] bg-[#111111] p-4 flex flex-col gap-4"
                  >
                  <Link
                    href={`/properties/${property.id}`}
                    className="space-y-3"
                  >
                    <div className="h-52 rounded-xl border border-[#1f1f1f] bg-linear-to-br from-[#161616] via-[#1d1d1d] to-[#161616]" />
                    <div>
                      <h3 className="text-lg font-semibold">{property.title}</h3>
                      <p className="text-xs text-gray-400">
                        {property.location}
                      </p>
                      <p className="text-sm text-gray-400 mt-1">
                        {property.description}
                      </p>
                    </div>
                  </Link>
                  <div className="flex flex-wrap gap-2">
                    {property.specs.map((spec) => (
                      <span
                        key={spec}
                        className="inline-flex items-center gap-2 px-3 py-2 rounded-full bg-[#151515] border border-[#262626] text-sm text-gray-200"
                      >
                        {spec}
                      </span>
                    ))}
                  </div>
                  <div className="flex items-center justify-between">
                    <div>
                      <p className="text-xs text-gray-400">Price</p>
                      <p className="text-lg font-semibold">
                        {property.price}
                      </p>
                    </div>
                    <Link href={`/properties/${property.id}`}>
                      <Button
                        variant="primary"
                        size="md"
                        className="bg-[#6d41ff] hover:bg-[#5b2fff] text-white"
                      >
                        View Property Details
                      </Button>
                    </Link>
                  </div>
                </div>
              ))}
            </div>
            <div className="flex gap-3 justify-end pt-3">
              <Button
                variant="outline"
                size="md"
                className="border-[#262626] bg-[#111111] text-white hover:bg-[#1a1a1a]"
                onClick={() => scrollCarousel(propertyRef, -1)}
              >
                ←
              </Button>
              <Button
                variant="outline"
                size="md"
                className="border-[#262626] bg-[#111111] text-white hover:bg-[#1a1a1a]"
                onClick={() => scrollCarousel(propertyRef, 1)}
              >
                →
              </Button>
            </div>
          </div>
          <p className="text-xs text-gray-500">01 of 60</p>
        </section>

        {/* Testimonials */}
        <section className="space-y-6">
          <div className="flex items-center justify-between">
            <div>
              <h2 className="text-3xl font-semibold">What Our Clients Say</h2>
              <p className="text-sm text-gray-400 max-w-xl">
                Read the success stories and heartfelt testimonials from our
                valued clients.
              </p>
            </div>
            <Button
              variant="outline"
              size="md"
              className="border-[#262626] bg-[#111111] text-white hover:bg-[#1a1a1a]"
            >
              View All Testimonials
            </Button>
          </div>

          <div className="relative">
            <div
              ref={testimonialRef}
              className="flex gap-4 overflow-x-auto scroll-smooth pb-2"
            >
              {testimonials.map((testimony) => (
                <div
                  key={testimony.name}
                  className="min-w-[320px] rounded-2xl border border-[#1f1f1f] bg-[#111111] p-5 flex flex-col gap-4"
                >
                  <div className="flex items-center gap-2 text-yellow-400 text-lg">
                    {"★".repeat(5)}
                  </div>
                  <div>
                    <h3 className="text-lg font-semibold">{testimony.title}</h3>
                    <p className="text-sm text-gray-400 mt-2">
                      {testimony.body}
                    </p>
                  </div>
                  <div className="flex items-center gap-3 pt-2">
                    <div className="h-10 w-10 rounded-full bg-linear-to-br from-[#161616] via-[#1d1d1d] to-[#161616]" />
                    <div>
                      <p className="text-sm font-semibold">{testimony.name}</p>
                      <p className="text-xs text-gray-400">
                        {testimony.location}
                      </p>
                    </div>
                  </div>
                </div>
              ))}
            </div>
            <div className="flex gap-3 justify-end pt-3">
              <Button
                variant="outline"
                size="md"
                className="border-[#262626] bg-[#111111] text-white hover:bg-[#1a1a1a]"
                onClick={() => scrollCarousel(testimonialRef, -1)}
              >
                ←
              </Button>
              <Button
                variant="outline"
                size="md"
                className="border-[#262626] bg-[#111111] text-white hover:bg-[#1a1a1a]"
                onClick={() => scrollCarousel(testimonialRef, 1)}
              >
                →
              </Button>
            </div>
          </div>
          <p className="text-xs text-gray-500">01 of 10</p>
        </section>

        {/* FAQ */}
        <section className="space-y-6">
          <div className="flex items-center justify-between">
            <div>
              <h2 className="text-3xl font-semibold">
                Frequently Asked Questions
              </h2>
              <p className="text-sm text-gray-400 max-w-xl">
                Find answers to common questions about Propchain&apos;s services,
                property listings, and the real estate process.
              </p>
            </div>
            <Button
              variant="outline"
              size="md"
              className="border-[#262626] bg-[#111111] text-white hover:bg-[#1a1a1a]"
            >
              View All FAQ&apos;s
            </Button>
          </div>

          <div className="relative">
            <div ref={faqRef} className="flex gap-4 overflow-x-auto scroll-smooth pb-2">
              {faqs.map((faq) => (
                <div
                  key={faq.question}
                  className="min-w-[320px] rounded-2xl border border-[#1f1f1f] bg-[#111111] p-5 flex flex-col gap-3"
                >
                  <h3 className="text-lg font-semibold text-gray-100">
                    {faq.question}
                  </h3>
                  <p className="text-sm text-gray-400">{faq.answer}</p>
                  <Button
                    variant="outline"
                    size="md"
                    className="mt-2 border-[#262626] bg-[#111111] text-white hover:bg-[#1a1a1a]"
                  >
                    Read More
                  </Button>
                </div>
              ))}
            </div>
            <div className="flex gap-3 justify-end pt-3">
              <Button
                variant="outline"
                size="md"
                className="border-[#262626] bg-[#111111] text-white hover:bg-[#1a1a1a]"
                onClick={() => scrollCarousel(faqRef, -1)}
              >
                ←
              </Button>
              <Button
                variant="outline"
                size="md"
                className="border-[#262626] bg-[#111111] text-white hover:bg-[#1a1a1a]"
                onClick={() => scrollCarousel(faqRef, 1)}
              >
                →
              </Button>
            </div>
          </div>
          <p className="text-xs text-gray-500">01 of 10</p>
        </section>
      </main>

      {/* CTA + Footer */}
      <section className="bg-[#0f0f0f] border-t border-[#1a1a1a]">
        <div className="max-w-6xl mx-auto px-4 py-12 space-y-10">
          <div className="rounded-3xl border border-[#1f1f1f] bg-[#0e0e0e] p-10 flex flex-col md:flex-row items-start md:items-center justify-between gap-6">
            <div className="space-y-2 max-w-3xl">
              <h3 className="text-3xl font-semibold">
                Start Your Real Estate Journey Today
              </h3>
              <p className="text-sm text-gray-300 leading-relaxed">
                Your dream property is just a click away. Whether you&apos;re looking
                for a new home, a strategic investment, or expert real estate
                advice, Propchain is here to assist you every step of the way.
              </p>
            </div>
            <Button
              variant="primary"
              size="md"
              className="bg-[#6d41ff] hover:bg-[#5b2fff] text-white"
            >
              Explore Properties
            </Button>
          </div>

          <div className="grid md:grid-cols-5 gap-8 text-sm text-gray-300">
            <div className="space-y-4">
              <div className="flex items-center gap-3">
                <div className="h-10 w-10 rounded-full bg-[#6d41ff]" />
                <span className="text-lg font-semibold text-white">
                  Propchain
                </span>
              </div>
              <div className="flex items-center gap-2 rounded-xl border border-[#1f1f1f] bg-[#0e0e0e] px-3 py-2">
                <span className="text-gray-400">✉</span>
                <input
                  className="bg-transparent outline-none text-sm w-full placeholder:text-gray-500"
                  placeholder="Enter Your Email"
                />
                <span>✈</span>
              </div>
            </div>

            <div className="space-y-2">
              <p className="text-white font-semibold">Home</p>
              <p>Hero Section</p>
              <p>Features</p>
              <p>Properties</p>
              <p>Testimonials</p>
              <p>FAQ&apos;s</p>
            </div>

            <div className="space-y-2">
              <p className="text-white font-semibold">About Us</p>
              <p>Our Story</p>
              <p>Our Works</p>
              <p>How It Works</p>
              <p>Our Team</p>
              <p>Our Clients</p>
            </div>

            <div className="space-y-2">
              <p className="text-white font-semibold">Properties</p>
              <p>Portfolio</p>
              <p>Categories</p>
            </div>

            <div className="space-y-2">
              <p className="text-white font-semibold">Services</p>
              <p>Valuation Mastery</p>
              <p>Strategic Marketing</p>
              <p>Negotiation Wizardry</p>
              <p>Closing Success</p>
              <p>Property Management</p>
            </div>
          </div>

          <div className="flex flex-col md:flex-row items-center justify-between gap-4 text-sm text-gray-400 border-t border-[#1a1a1a] pt-6">
            <p>@2023 Propchain. All Rights Reserved.</p>
            <p>Terms &amp; Conditions</p>
            <div className="flex gap-3">
              {socialIcons.map((icon) => (
                <div
                  key={icon}
                  className="h-10 w-10 rounded-full border border-[#1f1f1f] bg-[#0e0e0e] flex items-center justify-center"
                >
                  {icon === "facebook" && "f"}
                  {icon === "linkedin" && "in"}
                  {icon === "twitter" && "t"}
                  {icon === "youtube" && "▶"}
                </div>
              ))}
            </div>
          </div>
        </div>
      </section>
    </div>
  );
}

export default LandingPage;

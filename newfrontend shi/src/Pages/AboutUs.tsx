import "../index.css";
import { Button } from "../Components/atoms/Button";
import { Navbar } from "../Components/organisms/Navbar";

const stats = [
  { label: "Happy Customers", value: "200+" },
  { label: "Properties For Clients", value: "10k+" },
  { label: "Years of Experience", value: "16+" },
];

const values = [
  {
    title: "Trust",
    body: "Trust is the cornerstone of every successful real estate transaction.",
  },
  {
    title: "Excellence",
    body: "We set the bar high for ourselves. From the properties we list to the services we provide.",
  },
  {
    title: "Client-Centric",
    body: "Your dreams and needs are at the center of our universe. We listen, understand.",
  },
  {
    title: "Our Commitment",
    body: "We are dedicated to providing you with the highest level of service, professionalism, and support.",
  },
];

const achievements = [
  {
    title: "3+ Years of Excellence",
    body: "With over 3 years in the industry, we’ve amassed a wealth of knowledge and experience, becoming a go-to resource for all things real estate.",
  },
  {
    title: "Happy Clients",
    body: "Our greatest achievement is the satisfaction of our clients. Their success stories fuel our passion for what we do.",
  },
  {
    title: "Industry Recognition",
    body: "We’ve earned the respect of our peers and industry leaders, with accolades and awards that reflect our commitment to excellence.",
  },
];

const steps = [
  {
    title: "Discover a World of Possibilities",
    body: "Your journey begins with exploring our carefully curated property listings.",
  },
  {
    title: "Narrowing Down Your Choices",
    body: "Save properties to your account or shortlist and revisit your favorites.",
  },
  {
    title: "Personalized Guidance",
    body: "Our dedicated team of real estate experts is just a call or message away.",
  },
  {
    title: "See It for Yourself",
    body: "Arrange viewings and get a firsthand look at your potential new home.",
  },
  {
    title: "Making Informed Decisions",
    body: "We assist you with due diligence, inspections, legal checks, and market analysis.",
  },
  {
    title: "Getting the Best Deal",
    body: "We negotiate the best terms and secure the property at the right price.",
  },
];

const team = [
  { name: "Max Mitchell", role: "Founder" },
  { name: "Sarah Johnson", role: "Chief Real Estate Officer" },
  { name: "David Brown", role: "Head of Property Management" },
  { name: "Michael Turner", role: "Legal Counsel" },
];

const GradientBox = ({ className = "" }: { className?: string }) => (
  <div
    className={`bg-linear-to-br from-[#1a1a1a] via-[#1f1f1f] to-[#141414] border border-[#262626] ${className}`}
  />
);

function AboutUsPage() {
  return (
    <div className="min-h-screen w-full bg-[#0f0f0f] text-white">
      <Navbar currentPage="/about" />

      <main className="max-w-6xl mx-auto px-4 pt-28 pb-16 space-y-16">
        {/* Hero / Journey */}
        <section className="grid md:grid-cols-2 gap-10 items-center">
          <div className="space-y-5">
            <p className="uppercase tracking-[0.25em] text-xs text-gray-400">
              About Propchain
            </p>
            <h1 className="text-4xl md:text-5xl font-semibold">Our Journey</h1>
            <p className="text-sm text-gray-300 leading-relaxed">
              Our story is one of continuous growth and evolution. We started as
              a small team with big dreams, determined to create a real estate
              platform that transcended the ordinary. Over the years, we&apos;ve
              expanded our reach, forged valuable partnerships, and gained the
              trust of countless clients.
            </p>
            <div className="grid grid-cols-3 gap-3 pt-2">
              {stats.map((stat) => (
                <div
                  key={stat.label}
                  className="rounded-2xl border border-[#1f1f1f] bg-[#121212] px-4 py-4"
                >
                  <div className="text-2xl font-semibold">{stat.value}</div>
                  <p className="text-xs text-gray-400 mt-1">{stat.label}</p>
                </div>
              ))}
            </div>
          </div>

          <div className="rounded-3xl overflow-hidden border border-[#1f1f1f] bg-[#0c0c0c]">
            <GradientBox className="w-full h-[320px] rounded-none" />
          </div>
        </section>

        {/* Values */}
        <section className="grid md:grid-cols-[1fr,1.4fr] gap-8 items-center">
          <div className="space-y-3">
            <h2 className="text-3xl font-semibold">Our Values</h2>
            <p className="text-sm text-gray-300 leading-relaxed">
              Our story is one of continuous growth and evolution. We started as
              a small team with big dreams, determined to create a real estate
              platform that transcended the ordinary.
            </p>
          </div>
          <div className="rounded-3xl border border-[#1f1f1f] bg-[#0f0f0f] p-4 grid md:grid-cols-2 gap-4">
            {values.map((value) => (
              <div
                key={value.title}
                className="rounded-2xl border border-[#1f1f1f] bg-[#111111] p-4 space-y-2"
              >
                <div className="flex items-center gap-3">
                  <div className="h-10 w-10 rounded-full border border-[#262626] bg-[#1b1b1b] flex items-center justify-center text-lg text-purple-300">
                    ★
                  </div>
                  <p className="font-semibold">{value.title}</p>
                </div>
                <p className="text-sm text-gray-400">{value.body}</p>
              </div>
            ))}
          </div>
        </section>

        {/* Achievements */}
        <section className="space-y-4">
          <div className="space-y-2">
            <h2 className="text-3xl font-semibold">Our Achievements</h2>
            <p className="text-sm text-gray-300 leading-relaxed">
              We started as a small team with big dreams, determined to create a
              real estate platform that transcended the ordinary.
            </p>
          </div>
          <div className="grid md:grid-cols-3 gap-4">
            {achievements.map((item) => (
              <div
                key={item.title}
                className="rounded-2xl border border-[#1f1f1f] bg-[#111111] p-5 space-y-3"
              >
                <h3 className="text-lg font-semibold">{item.title}</h3>
                <p className="text-sm text-gray-400 leading-relaxed">
                  {item.body}
                </p>
              </div>
            ))}
          </div>
        </section>

        {/* Steps */}
        <section className="space-y-4">
          <div className="space-y-2">
            <h2 className="text-3xl font-semibold">
              Navigating the Propchain Experience
            </h2>
            <p className="text-sm text-gray-300 leading-relaxed">
              A straightforward process to help you find and purchase your dream
              property with ease.
            </p>
          </div>
          <div className="grid md:grid-cols-3 gap-4">
            {steps.map((step, idx) => (
              <div
                key={step.title}
                className="rounded-2xl border border-[#1f1f1f] bg-[#111111] p-5 space-y-3"
              >
                <p className="text-xs text-gray-400">Step {String(idx + 1).padStart(2, "0")}</p>
                <h3 className="text-lg font-semibold">{step.title}</h3>
                <p className="text-sm text-gray-400 leading-relaxed">
                  {step.body}
                </p>
              </div>
            ))}
          </div>
        </section>

        {/* Team */}
        <section className="space-y-4">
          <div className="space-y-2">
            <h2 className="text-3xl font-semibold">Meet the Propchain Team</h2>
            <p className="text-sm text-gray-300 leading-relaxed">
              Get to know the people behind our mission to make your real estate
              dreams a reality.
            </p>
          </div>
          <div className="grid md:grid-cols-4 gap-4">
            {team.map((member) => (
              <div
                key={member.name}
                className="rounded-2xl border border-[#1f1f1f] bg-[#111111] p-5 space-y-4"
              >
                <div className="flex justify-center">
                  <GradientBox className="h-24 w-24 rounded-full" />
                </div>
                <div className="text-center space-y-1">
                  <p className="text-lg font-semibold">{member.name}</p>
                  <p className="text-xs text-gray-400">{member.role}</p>
                </div>
                <Button
                  variant="outline"
                  size="md"
                  className="w-full border-[#262626] bg-[#111111] text-white hover:bg-[#1a1a1a]"
                >
                  Say Hello 👋
                </Button>
              </div>
            ))}
          </div>
        </section>
      </main>
    </div>
  );
}

export default AboutUsPage;


import { Outlet } from "react-router-dom";
import { Sidebar } from "./Sidebar";

export const DashboardLayout: React.FC = () => {
  return (
    <div className="min-h-screen w-full bg-[#0f0f0f] text-white">
      <Sidebar />

      {/* Main content area */}
      <main className="md:ml-64 min-h-screen">
        <div className="max-w-6xl mx-auto px-4 py-8 pt-16 md:pt-8">
          <Outlet />
        </div>
      </main>
    </div>
  );
};

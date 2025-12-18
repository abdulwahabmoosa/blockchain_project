import { useCallback } from "react";
import { Outlet } from "react-router-dom";
import { Sidebar } from "./Sidebar";
import { useWallet } from "../../hooks/useWallet";

export const DashboardLayout: React.FC = () => {
  const { disconnect } = useWallet();

  const handleLogout = useCallback(async () => {
    // Reuse existing auth clearing semantics
    localStorage.removeItem("auth");
    localStorage.removeItem("token");
    localStorage.removeItem("user");
    await disconnect();
    window.location.href = "/";
  }, [disconnect]);

  return (
    <div className="min-h-screen w-full bg-[#0f0f0f] text-white">
      <Sidebar onLogout={handleLogout} />

      {/* Main content area */}
      <main className="md:ml-64 min-h-screen">
        <div className="max-w-6xl mx-auto px-4 py-8 pt-16 md:pt-8">
          <Outlet />
        </div>
      </main>
    </div>
  );
};




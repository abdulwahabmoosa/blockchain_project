import { useCallback, useEffect, useState } from "react";
import { Outlet, useNavigate } from "react-router-dom";
import { AdminSidebar } from "./AdminSidebar";
import type { User } from "../../types";
import { useWallet } from "../../hooks/useWallet";

export const AdminDashboardLayout: React.FC = () => {
  const [user, setUser] = useState<User | null>(null);
  const [loading, setLoading] = useState(true);
  const navigate = useNavigate();
  const { disconnect } = useWallet();

  useEffect(() => {
    const userStr = localStorage.getItem("user");
    if (userStr) {
      const userData = JSON.parse(userStr);
      setUser(userData);
      if (userData.Role !== "admin") {
        navigate("/dashboard");
        return;
      }
    } else {
      navigate("/login");
      return;
    }
    setLoading(false);
  }, [navigate]);

  const handleLogout = useCallback(async () => {
    localStorage.removeItem("auth");
    localStorage.removeItem("token");
    localStorage.removeItem("user");
    await disconnect();
    navigate("/login");
  }, [disconnect, navigate]);

  if (loading || !user) {
    return (
      <div className="min-h-screen w-full bg-[#0f0f0f] text-white flex items-center justify-center">
        <p className="text-gray-400">Loading admin dashboard...</p>
      </div>
    );
  }

  return (
    <div className="min-h-screen w-full bg-[#0f0f0f] text-white">
      <AdminSidebar onLogout={handleLogout} />

      {/* Main content area */}
      <main className="md:ml-64 min-h-screen">
        <div className="max-w-6xl mx-auto px-4 py-8 pt-16 md:pt-8">
          <Outlet />
        </div>
      </main>
    </div>
  );
};



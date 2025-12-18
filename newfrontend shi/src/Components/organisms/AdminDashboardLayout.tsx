import { Outlet } from "react-router-dom";
import { AdminSidebar } from "./AdminSidebar";

export const AdminDashboardLayout: React.FC = () => {
  // TODO: Re-enable admin check after testing
  // const [user, setUser] = useState<User | null>(null);
  // const [loading, setLoading] = useState(true);
  // const navigate = useNavigate();
  // useEffect(() => {
  //   const userStr = localStorage.getItem("user");
  //   if (userStr) {
  //     const userData = JSON.parse(userStr);
  //     setUser(userData);
  //     if (userData.Role !== "admin") {
  //       navigate("/dashboard");
  //       return;
  //     }
  //   } else {
  //     navigate("/login");
  //     return;
  //   }
  //   setLoading(false);
  // }, [navigate]);

  return (
    <div className="min-h-screen w-full bg-[#0f0f0f] text-white">
      <AdminSidebar />

      {/* Main content area */}
      <main className="md:ml-64 min-h-screen">
        <div className="max-w-6xl mx-auto px-4 py-8 pt-16 md:pt-8">
          <Outlet />
        </div>
      </main>
    </div>
  );
};

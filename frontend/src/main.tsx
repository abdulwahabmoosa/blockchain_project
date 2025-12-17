import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import "./index.css";
import LandingPage from "./Pages/LandingPage.tsx";
import AboutUsPage from "./Pages/AboutUs.tsx";
import LoginPage from "./Pages/Login.tsx";
import SignupPage from "./Pages/Signup.tsx";
import PropertiesPage from "./Pages/Properties.tsx";
import PropertyDetailsPage from "./Pages/PropertyDetails.tsx";
import ServicesPage from "./Pages/Services.tsx";
import DashboardPage from "./Pages/Dashboard.tsx";
import AdminDashboard from "./Pages/AdminDashboard.tsx";
import CreatePropertyPage from "./Pages/CreateProperty.tsx";
import UserApprovalManagementPage from "./Pages/UserApprovalManagement.tsx";
import { ProtectedRoute } from "./Components/ProtectedRoute.tsx";
import { WalletProvider } from "./hooks/useWallet.tsx";

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <WalletProvider>
      <BrowserRouter>
        <Routes>
        <Route path="/" element={<LandingPage />} />
        <Route path="/about" element={<AboutUsPage />} />
        <Route path="/login" element={<LoginPage />} />
        <Route path="/signup" element={<SignupPage />} />
        <Route path="/properties" element={<PropertiesPage />} />
        <Route path="/services" element={<ServicesPage />} />
        <Route path="/properties/:id" element={<PropertyDetailsPage />} />
        <Route
          path="/dashboard"
          element={
            <ProtectedRoute>
              <DashboardPage />
            </ProtectedRoute>
          }
        />
        <Route
          path="/admin"
          element={
            <ProtectedRoute>
              <AdminDashboard />
            </ProtectedRoute>
          }
        />
        <Route
          path="/create-property"
          element={
            <ProtectedRoute>
              <CreatePropertyPage />
            </ProtectedRoute>
          }
        />
        <Route
          path="/admin/approve-users"
          element={
            <ProtectedRoute>
              <UserApprovalManagementPage />
            </ProtectedRoute>
          }
        />
      </Routes>
      </BrowserRouter>
    </WalletProvider>
  </StrictMode>
);

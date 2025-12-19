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
import { DashboardLayout } from "./Components/organisms/DashboardLayout";
import DashboardHome from "./Pages/DashboardHome";
import ViewProperties from "./Pages/ViewProperties";
import UploadProperty from "./Pages/UploadProperty";
import MyInvestments from "./Pages/MyInvestments";
import WalletStatus from "./Pages/WalletStatus";
import MyAccount from "./Pages/MyAccount";
import { AdminDashboardLayout } from "./Components/organisms/AdminDashboardLayout";
import AdminDashboardHome from "./Pages/AdminDashboardHome";
import AdminManageUsers from "./Pages/AdminManageUsers";
import AdminManageProperties from "./Pages/AdminManageProperties";
import AdminCreateProperty from "./Pages/AdminCreateProperty";
import PropertyUploadRequestManagement from "./Pages/PropertyUploadRequestManagement";
import Messages from "./Pages/Messages";

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
                <DashboardLayout />
              </ProtectedRoute>
            }
          >
            <Route index element={<DashboardHome />} />
            <Route path="properties" element={<ViewProperties />} />
            <Route path="upload" element={<UploadProperty />} />
            <Route path="investments" element={<MyInvestments />} />
            <Route path="wallet" element={<WalletStatus />} />
            <Route path="account" element={<MyAccount />} />
            <Route path="messages" element={<Messages />} />
          </Route>

          <Route
            path="/admin"
            element={
              <ProtectedRoute>
                <AdminDashboardLayout />
              </ProtectedRoute>
            }
          >
            <Route index element={<AdminDashboardHome />} />
            <Route path="create-property" element={<AdminCreateProperty />} />
            <Route path="users" element={<AdminManageUsers />} />
            <Route path="properties" element={<AdminManageProperties />} />
            <Route path="property-requests" element={<PropertyUploadRequestManagement />} />
            <Route path="wallet" element={<WalletStatus />} />
          </Route>

          {/* Existing routes kept for compatibility */}
          <Route
            path="/dashboard-legacy"
            element={
              <ProtectedRoute>
                <DashboardPage />
              </ProtectedRoute>
            }
          />
          <Route
            path="/admin-legacy"
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

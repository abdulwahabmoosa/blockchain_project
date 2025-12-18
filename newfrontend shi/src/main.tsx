import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import "./index.css";
import LandingPage from "./Pages/LandingPage.tsx";
import AboutUsPage from "./Pages/AboutUs.tsx";
import LoginPage from "./Pages/Login.tsx";
import SignupPage from "./Pages/Signup.tsx";
import PropertyDetailsPage from "./Pages/PropertyDetails.tsx";
import { ProtectedRoute } from "./Components/ProtectedRoute.tsx";
import { DashboardLayout } from "./Components/organisms/DashboardLayout.tsx";
import DashboardHome from "./Pages/DashboardHome.tsx";
import ViewProperties from "./Pages/ViewProperties.tsx";
import UploadProperty from "./Pages/UploadProperty.tsx";
import MyInvestments from "./Pages/MyInvestments.tsx";
import WalletStatus from "./Pages/WalletStatus.tsx";
import MyAccount from "./Pages/MyAccount.tsx";
import { AdminDashboardLayout } from "./Components/organisms/AdminDashboardLayout.tsx";
import AdminDashboardHome from "./Pages/AdminDashboardHome.tsx";
import AdminManageUsers from "./Pages/AdminManageUsers.tsx";
import AdminCreateProperty from "./Pages/AdminCreateProperty.tsx";

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<LandingPage />} />
        <Route path="/about" element={<AboutUsPage />} />
        <Route path="/login" element={<LoginPage />} />
        <Route path="/signup" element={<SignupPage />} />
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
        </Route>
        {/* TODO: Re-enable ProtectedRoute after testing */}
        <Route path="/admin" element={<AdminDashboardLayout />}>
          <Route index element={<AdminDashboardHome />} />
          <Route path="create-property" element={<AdminCreateProperty />} />
          <Route path="users" element={<AdminManageUsers />} />
        </Route>
      </Routes>
    </BrowserRouter>
  </StrictMode>
);

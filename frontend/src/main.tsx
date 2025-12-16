import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import "./index.css";
import LandingPage from "./Pages/LandingPage.tsx";
import AboutUsPage from "./Pages/AboutUs.tsx";
import LoginPage from "./Pages/Login.tsx";
import SignupPage from "./Pages/Signup.tsx";
import PropertyDetailsPage from "./Pages/PropertyDetails.tsx";
import DashboardPage from "./Pages/Dashboard.tsx";
import { ProtectedRoute } from "./Components/ProtectedRoute.tsx";

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
              <DashboardPage />
            </ProtectedRoute>
          }
        />
      </Routes>
    </BrowserRouter>
  </StrictMode>
);

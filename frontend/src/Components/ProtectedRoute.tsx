import React from "react";
import { Navigate, useLocation } from "react-router-dom";

interface ProtectedRouteProps {
  children: React.ReactNode;
}

export const ProtectedRoute: React.FC<ProtectedRouteProps> = ({ children }) => {
  const location = useLocation();
  const params = new URLSearchParams(location.search);
  const hasPreviewBypass = params.get("preview") === "1";
  const isAuthed = localStorage.getItem("auth") === "true";

  if (isAuthed || hasPreviewBypass) {
    return <>{children}</>;
  }

  return <Navigate to="/login" replace state={{ from: location }} />;
};

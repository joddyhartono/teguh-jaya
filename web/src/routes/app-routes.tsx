import { Route, Routes } from "react-router";
import ProtectedRoute from "./protected-route";
import AdminLayout from "@/layouts/admin-layout";
import LoginPage from "@/pages/admin/login-page";
import ProductPage from "@/pages/public/product-page";
import PublicLayout from "@/layouts/public-layout";
import HomePage from "@/pages/public/home-page";
import DashboardPage from "@/pages/admin/dashboard-page";

const AppRoutes = () => {
  return (
    <Routes>
      <Route path="/admin/login" element={<LoginPage />} />

      <Route element={<ProtectedRoute />}>
        <Route element={<AdminLayout />}>
          <Route path="/admin/dashboard" element={<DashboardPage />} />
        </Route>
      </Route>

      <Route element={<PublicLayout />}>
        <Route index element={<HomePage />} />
        <Route path="/products" element={<ProductPage />} />
      </Route>
    </Routes>
  );
};

export default AppRoutes;

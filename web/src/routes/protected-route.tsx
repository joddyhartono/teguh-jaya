import { useAppSelector } from "@/store/hooks";
import { Navigate, Outlet } from "react-router";

const ProtectedRoute = () => {
  const token = useAppSelector((state) => state.auth.token);
  if (!token) {
    return <Navigate to="/admin/login" replace={true} />;
  }
  return <Outlet />;
};

export default ProtectedRoute;

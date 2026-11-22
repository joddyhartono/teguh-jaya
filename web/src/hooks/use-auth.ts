import type { LoginPayload } from "@/types/auth";
import * as authSlice from "@/store/slices/auth-slice";
import { useAppDispatch } from "@/store/hooks";

const useAuth = () => {
  const dispatch = useAppDispatch();

  const login = async (payload: LoginPayload) => {
    await dispatch(authSlice.login(payload)).unwrap();
  };

  const logout = () => {
    dispatch(authSlice.logout());
  };

  return {
    login,
    logout,
  };
};

export default useAuth;

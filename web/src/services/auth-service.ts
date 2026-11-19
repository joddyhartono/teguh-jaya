import instance from "@/lib/axios-instance";
import { AUTH } from "@/lib/endpoints";
import type { LoginPayload, LoginResponse } from "@/types/auth";

const login = async (payload: LoginPayload) => {
  const response = await instance.post<LoginResponse>(AUTH.LOGIN, payload);
  return response.data;
};

export { login };

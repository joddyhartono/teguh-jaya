import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";

import * as authService from "@/services/auth-service";
import axios from "axios";
import type { LoginPayload } from "@/types/auth";

type AuthState = {
  token: string | null;
};

const initialState: AuthState = {
  token: localStorage.getItem("token"),
};

export const login = createAsyncThunk(
  "auth/login",
  async (payload: LoginPayload) => {
    try {
      const data = await authService.login(payload);
      localStorage.setItem("token", data.token);
      return data.token;
    } catch (error) {
      if (axios.isAxiosError(error)) {
        throw new Error(error.response?.data.error);
      }
      throw error;
    }
  },
);

const authSlice = createSlice({
  name: "auth",
  initialState,
  reducers: {
    logout: (state) => {
      state.token = null;
    },
  },
  extraReducers: (builder) => {
    builder.addCase(login.fulfilled, (state, action) => {
      state.token = action.payload;
    });
  },
});

export const { logout } = authSlice.actions;
export default authSlice.reducer;

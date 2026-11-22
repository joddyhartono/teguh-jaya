import type { Category } from "@/types/category";
import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import axios from "axios";

import * as categoryService from "@/services/category-service";

type CategoryState = {
  categories: Category[];
  status: string;
  error?: string;
};

const initialState: CategoryState = {
  categories: [],
  status: "idle",
};

export const fetchCategories = createAsyncThunk(
  "category/fetchCategories",
  async () => {
    try {
      const data = await categoryService.getCategories();
      return data;
    } catch (error) {
      if (axios.isAxiosError(error)) {
        throw new Error(error.response?.data.error);
      }
      throw error;
    }
  },
);

const categorySlice = createSlice({
  name: "category",
  initialState: initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder.addCase(fetchCategories.pending, (state) => {
      state.status = "loading";
    });
    builder.addCase(fetchCategories.fulfilled, (state, action) => {
      state.categories = action.payload;
      state.status = "succeeded";
    });
    builder.addCase(fetchCategories.rejected, (state, action) => {
      state.status = "failed";
      state.error = action.error.message;
    });
  },
});

export default categorySlice.reducer;

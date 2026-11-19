import instance from "@/lib/axios-instance";
import { PRODUCTS } from "@/lib/endpoints";

const getProductsByCategoryId = async (categoryId: string) => {
  const response = await instance.get(PRODUCTS.BY_CATEGORY(categoryId));
  return response.data;
};

export { getProductsByCategoryId };

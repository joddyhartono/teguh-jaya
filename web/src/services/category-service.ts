import instance from "@/lib/axios-instance";
import { CATEGORIES } from "@/lib/endpoints";

const getCategories = async () => {
  const response = await instance.get(CATEGORIES.BASE);
  return response.data;
};

export { getCategories };
